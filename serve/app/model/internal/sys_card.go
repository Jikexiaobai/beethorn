// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysCard is the golang structure for table sys_card.
type SysCard struct {
    CardId     int64       `orm:"card_id,primary" json:"cardId"`     //                         
    UsedId     int64       `orm:"used_id"         json:"usedId"`     // 使用者id                
    SecretKey  string      `orm:"secret_key"      json:"secretKey"`  //                         
    Money      float64     `orm:"money"           json:"money"`      //                         
    Status     int         `orm:"status"          json:"status"`     // 状态: 1未使用，2已使用  
    CreateTime *gtime.Time `orm:"create_time"     json:"createTime"` //                         
}