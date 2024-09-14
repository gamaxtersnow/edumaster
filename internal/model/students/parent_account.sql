-- 家长账号表
CREATE TABLE parent_account (
                                 id INT AUTO_INCREMENT,
                                 student_id INT NOT NULL DEFAULT 0 COMMENT '关联的学生ID',
                                 name VARCHAR(100) NOT NULL DEFAULT '' COMMENT '家长姓名',
                                 relationship VARCHAR(50) NOT NULL DEFAULT '' COMMENT '与学生的关系',
                                 phone_number VARCHAR(20) NOT NULL DEFAULT '' COMMENT '家长手机号',
                                 login_password VARCHAR(255) NOT NULL DEFAULT '' COMMENT '家长登录密码',
                                 created_at INT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
                                 updated_at INT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
                                 PRIMARY KEY (id)
) COMMENT '家长账号信息表';