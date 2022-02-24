CREATE TABLE IF NOT EXISTS fishes (
    id int AUTO_INCREMENT,
    name varchar(100) NOT NULL COMMENT "名前",
    classification tinyint NOT NULL DEFAULT 0 COMMENT "分類 0=該当なし、不明, 1=海水, 2=淡水",
    description TEXT NOT NULL COMMENT "説明",
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
) COMMENT = "魚";
