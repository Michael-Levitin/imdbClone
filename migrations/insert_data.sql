INSERT INTO movies (movie, description, release, rating, removed, created_at)
VALUES ('BraveHeart',
        'BraveHeart is a 1995 American epic historical drama film directed by, produced by, and starring Mel Gibson. Gibson portrays Sir William Wallace, a late-13th century Scottish warrior who led the Scots in the First War of Scottish Independence against King Edward I of England. The film also stars Sophie Marceau, Patrick McGoohan and Catherine McCormack. The story is inspired by Blind Harry''s 15th century epic poem The Actes and Deidis of the Illustre and Vallyeant Campioun Schir William Wallace and was adapted for the screen by Randall Wallace.',
        now(), 8.3, false, now()),
       ('Edge of Tomorrow',
        'Edge of Tomorrow[a] is a 2014 American science fiction action film directed by Doug Liman and written by Christopher McQuarrie and the writing team of Jez and John-Henry Butterworth, loosely based on the Japanese novel All You Need Is Kill by Hiroshi Sakurazaka. Starring Tom Cruise and Emily Blunt, the film takes place in a future where most of Europe is occupied by an alien race. Major William Cage (Cruise), a public relations officer with limited combat experience, is forced by his superiors to join a landing operation against the aliens, only to find himself experiencing a time loop as he tries to find a way to defeat the invaders. Bill Paxton and Brendan Gleeson also cast as supporting roles.',
        now(), 7.9, false, now()),
       ('DragonHeart',
        'Dragonheart (stylized as DragonHeart) is a 1996 fantasy adventure film directed by Rob Cohen and written by Charles Edward Pogue, based on a story created by him and Patrick Read Johnson. The film stars Dennis Quaid, David Thewlis, Pete Postlethwaite, Dina Meyer, and Sean Connery as the voice of Draco. It was nominated for the Academy Award for Best Visual Effects and various other awards in 1996 and 1997. The film received mixed reviews, with critics praising the premise, visual effects, and character development but panning the script as confusing and clichéd. It was a box-office success, earning $115 million worldwide. It was dedicated to the memory of Steve Price and Irwin Cohen.',
        now(), 6.4, false, now()),
       ('Brave',
        'Brave is a 2012 American animated fantasy film produced by Pixar Animation Studios and released by Walt Disney Pictures. The film was directed by Mark Andrews and Brenda Chapman (in the former''s feature directorial debut), co-directed by Steve Purcell, and produced by Katherine Sarafian, with John Lasseter, Andrew Stanton, and Pete Docter serving as executive producers. The story was written by Chapman, who also co-wrote the film''s screenplay with Andrews, Purcell, and Irene Mecchi. The film stars the voices of Kelly Macdonald, Billy Connolly, Emma Thompson, Julie Walters, Robbie Coltrane, Kevin McKidd, and Craig Ferguson. Set in the Scottish Highlands, the film tells the story of Princess Merida of DunBroch (Macdonald) who defies an age-old custom, causing chaos in the kingdom by expressing the desire not to be betrothed.',
        now(), 7.1, false, now());


INSERT INTO actors (name, dob, removed, created_at)
VALUES ('Mel Gibson', now(), false, now()),
       ('Sophie Marceau', now(), false, now()),
       ('Patrick McGoohan', now(), false, now()),
       ('Catherine McCormack', now(), false, now()),
       ('Tom Cruise', now(), false, now()),
       ('Emily Blunt', now(), false, now()),
       ('Tom Hanks', now(), false, now()),
       ('Sharon Stone', now(), false, now());


INSERT INTO parts (movie_id, actor_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (1, 4),
       (2, 5),
       (2, 6);


