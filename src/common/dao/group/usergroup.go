// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package group

import (
	"fmt"

	"github.com/vmware/harbor/src/common/dao"
	"github.com/vmware/harbor/src/common/models"
	"github.com/vmware/harbor/src/common/utils/log"
)

const (
	// DupLDAPGroupMessage ...
	DupLDAPGroupMessage = "Found an existing LDAP group with DN"
)

// AddUserGroup - Add User Group
func AddUserGroup(userGroup models.UserGroup) (int, error) {
	o := dao.GetOrmer()
	query := models.UserGroup{GroupType: userGroup.GroupType, LdapGroupDN: userGroup.LdapGroupDN}
	result, err := QueryUserGroup(query)
	if err != nil {
		return 0, err
	}
	if len(result) > 0 {
		log.Debugf("Already exist same group type and group property")
		return result[0].ID, fmt.Errorf(DupLDAPGroupMessage+":%v", userGroup.LdapGroupDN)
	}
	id, err := o.Insert(&userGroup)

	if err != nil {
		return 0, err
	}
	return int(id), err
}

// QueryUserGroup - Query User Group
func QueryUserGroup(query models.UserGroup) ([]*models.UserGroup, error) {
	o := dao.GetOrmer()
	sql := `select id, group_name, group_type, ldap_group_dn from user_group where 1=1 `
	sqlParam := make([]interface{}, 1)
	groups := []*models.UserGroup{}
	if len(query.GroupName) != 0 {
		sql += ` and group_name like ? `
		sqlParam = append(sqlParam, `%`+dao.Escape(query.GroupName)+`%`)
	}

	if query.GroupType != 0 {
		sql += ` and group_type = ? `
		sqlParam = append(sqlParam, query.GroupType)
	}

	if len(query.LdapGroupDN) != 0 {
		sql += ` and ldap_group_dn = ? `
		sqlParam = append(sqlParam, query.LdapGroupDN)
	}
	_, err := o.Raw(sql, sqlParam).QueryRows(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// GetUserGroup ...
func GetUserGroup(id int) (*models.UserGroup, error) {
	userGroup := models.UserGroup{ID: id}
	o := dao.GetOrmer()
	err := o.Read(&userGroup)
	if err != nil {
		return nil, err
	}
	return &userGroup, nil
}

// DeleteUserGroup ...
func DeleteUserGroup(id int) error {
	userGroup := models.UserGroup{ID: id}
	o := dao.GetOrmer()
	_, err := o.Delete(&userGroup)
	if err == nil {
		//Delete all related project members
		sql := `delete from project_member where entity_id = ? and entity_type='g'`
		_, err := o.Raw(sql, id).Exec()
		if err != nil {
			return err
		}
	}
	return err
}

// UpdateUserGroupName ...
func UpdateUserGroupName(id int, groupName string) error {
	log.Debugf("Updating user_group with id:%v, name:%v", id, groupName)
	o := dao.GetOrmer()
	sql := "update user_group set group_name = ? where id =  ? "
	_, err := o.Raw(sql, groupName, id).Exec()
	return err
}
