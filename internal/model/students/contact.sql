-- 联系人表
CREATE TABLE contact (
                          id INT AUTO_INCREMENT ,
                          student_id INT NOT NULL default 0 COMMENT '关联的学生ID',
                          type tinyint NOT NULL default 1 COMMENT '联系人类型 1 主联系人 2 第二联系人',
                          phone_number VARCHAR(20) NOT NULL default '' COMMENT '联系人手机号',
                          relationship VARCHAR(50) NOT NULL default '' COMMENT '与学生的关系',
                          name VARCHAR(100) NOT NULL default '' COMMENT '联系人姓名',
                          wechat_id VARCHAR(100) NOT NULL default '' COMMENT '联系人微信号',
                          created_at INT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
                          updated_at INT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
                          PRIMARY KEY (id)
) COMMENT '联系人信息表';