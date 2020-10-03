use gorest;
create table articles (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, title varchar(255), content text);
insert into articles (title, content) values('lorem ipsum', 'lorem ipsum sit dolor amet')
