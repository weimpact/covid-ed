INSERT INTO facts (id, title, description)
    VALUES
    (1, 'Mask only prevents spreading from you', 'If you wear mask and sneeze you will not spread to others'),
    (2, 'Everyone gets corona', '85k+ deaths across world due to corona of all ages'),
    (3, 'Corona is not human made', 'corona virus is not made in lab by humans as it is not mismash of existing virus');


INSERT INTO myths (title, description, fact_id)
    VALUES
    ('Face masks protect against coronavirus', 'if you wear face mask then you wont get corona', 1),
    ('Only older adults and young gets corona', 'old and young are subject to disease and at risk', 2),
    ('Corona is human made and spread from china lab', 'china wuhans labe is the source', 3);


INSERT INTO articles (title, url, fact_id)
    VALUES
    ('covid is not human mande (with papers)', 'https://www.sciencenews.org/article/coronavirus-covid-19-not-human-made-lab-genetic-analysis-nature', 3),
    ('covid is not human mande (livescience blog)', 'https://www.livescience.com/coronavirus-not-human-made-in-lab.html', 3);



