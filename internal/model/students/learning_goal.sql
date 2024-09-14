-- 学习目标表
CREATE TABLE learning_goal (
                                id INT AUTO_INCREMENT,
                                student_id INT NOT NULL default 0 COMMENT '关联的学生ID',
                                target_score FLOAT NOT NULL default 0 COMMENT '目标分数',
                                entry_score FLOAT NOT NULL default 0 COMMENT '入学分数',
                                target_course VARCHAR(100) NOT NULL default '' COMMENT '咨询课程',
                                target_country VARCHAR(100) NOT NULL default '' COMMENT '目标国家',
                                created_at INT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
                                updated_at INT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
                                PRIMARY KEY (id)
) COMMENT '学习目标表';