CREATE TABLE wineDetails (
                        id INT NOT NULL AUTO_INCREMENT,
                         country VARCHAR(256),
                         details LONGTEXT,
                         designation VARCHAR(256),
                         price VARCHAR(256),
                         title VARCHAR(256), 
                         winery VARCHAR(256),
                         taster_name VARCHAR(256),
                         PRIMARY KEY (id));


CREATE TABLE reviewer (
                        id INT NOT NULL AUTO_INCREMENT,
                        taster_twitter_handle VARCHAR(256) NOT NULL,
                        taster_name VARCHAR(256),
                        PRIMARY KEY (id)
                        );