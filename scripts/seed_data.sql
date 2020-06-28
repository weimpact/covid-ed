INSERT INTO facts (id, locale, title, description)
    VALUES (1, 'en-US', 'Mask only prevents spreading from you', 'If you wear mask and sneeze you will not spread to others'),
    (2, 'en-US', 'Everyone gets corona', 'more than 4.5 million deaths across world due to corona, and affected people of all ages'),
    (3, 'en-US', 'Corona is not human made', 'corona virus is not made in lab by humans as it is not mismash of existing virus'),
    (4, 'ta', 'எந்த வயதினரும் கொரோனாவைப் பெறலாம்', 'சீனாவில் பெரும்பான்மையான வழக்குகள் - 87% - 30 முதல் 79 வயதிற்குட்பட்டவர்களில் உள்ளன, பிப்ரவரி 11 ஆம் தேதி நிலவரப்படி கோவிட் -19 நோயால் கண்டறியப்பட்டவர்களில் 72,314 பேரின் தரவுகளின் அடிப்படையில் சீனாவின் நோய் கட்டுப்பாட்டு மையம் கடந்த மாதம் அறிக்கை செய்தது. மற்றவர்களுடன் அடிக்கடி தொடர்புகொள்வது போன்ற வாழ்க்கை முறையை விட உயிரியலைப் பற்றி ஏதாவது. பதின்ம வயதினரும், 20 வயதிற்குட்பட்டவர்களும் பள்ளி, வேலை மற்றும் பொதுப் போக்குவரத்தில் பலரைச் சந்திக்கிறார்கள், ஆனாலும் அவர்கள் குறிப்பிடத்தக்க விகிதத்தில் நோயைக் குறைப்பதாகத் தெரியவில்லை: 8.1% வழக்குகள் மட்டுமே 20-சில விஷயங்கள், 1.2% பதின்ம வயதினர்கள், மற்றும் 0.9% 9 அல்லது அதற்கு மேற்பட்டவர்கள். பிப்ரவரி 20 ஆம் தேதி நிலவரப்படி 78% வழக்குகள் 30 முதல் 69 வயதுடையவர்களில் இருப்பதாக சீனாவிற்கான உலக சுகாதார அமைப்பு பணி கண்டறிந்துள்ளது.'),
    (5, 'ta', 'தும்மல் அல்லது இருமல் போது உமிழக்கூடிய துளிகள் வழியாக கொரோனா பரவுகிறது', 'ஆதாரங்களின் அடிப்படையில், COVID-19 நோயாளிகளை கவனித்துக்கொள்பவர்களுக்கு துளி மற்றும் தொடர்பு முன்னெச்சரிக்கை நடவடிக்கைகளை WHO தொடர்ந்து பரிந்துரைத்து வருகிறது. ஆபத்து மதிப்பீட்டின்படி, ஏரோசல் உருவாக்கும் நடைமுறைகள் மற்றும் ஆதரவு சிகிச்சை மேற்கொள்ளப்படும் சூழ்நிலைகள் மற்றும் அமைப்புகளுக்கான வான்வழி முன்னெச்சரிக்கை நடவடிக்கைகளை WHO தொடர்ந்து பரிந்துரைக்கிறது'),
    (6, 'en-US', 'There is no official licensed drug which can be used in corona, use of Hydroxychloroquine could be harmful', 'earlier may lancet published article which could not find any benefits with 96 thousand poeople test, WHO halted its trial after lancet articles and resumed in june, and there is no evidence as of now that the drug helps.'),
    (7, 'ta', 'சமூக தூரத்தின் மூலம் மட்டுமே கொரோனாவைத் தடுக்க முடியும்', 'சமூக தூரத்தின் மூலம் மட்டுமே கொரோனாவைத் தடுக்க முடியும்');


INSERT INTO myths (locale, title, description, fact_id)
    VALUES
    ('en-US', 'Face masks protect against coronavirus', 'if you wear face mask then you wont get corona', 1),
    ('en-US', 'Only older adults and young gets corona', 'old and young are subject to disease and at risk', 2),
    ('en-US', 'Corona is human made and spread from china lab', 'china wuhans lab is the source', 3),
    ('ta', 'குழந்தைகள் கொரோனா வைரஸைப் பெற முடியாது', 'ஆரம்ப அறிக்கைகள் குழந்தைகள் கொரோனாவுக்கு குறைவாகவே பாதிக்கப்படுகின்றன என்று பரிந்துரைத்தன, ஆனால் பின்னர் குழந்தைகள் பெரியவர்களைப் போலவே பாதிக்கப்படுகின்றனர் என்று அறியப்படுகிறது.', 4),
    ('ta', 'புதிய கொரோனா வைரஸ் கொசு கடித்தால் பரவுகிறது', 'புதிய கொரோனா வைரஸ் கொசுக்களால் பரவக்கூடும் என்பதற்கான எந்த தகவலும் ஆதாரமும் இன்றுவரை இல்லை. புதிய கொரோனா வைரஸ் என்பது ஒரு சுவாச வைரஸ் ஆகும், இது முதன்மையாக பாதிக்கப்பட்ட நபர் இருமும்போது அல்லது தும்மும்போது உருவாகும் நீர்த்துளிகள் மூலமாகவோ அல்லது உமிழ்நீர் துளிகள் மூலமாகவோ அல்லது மூக்கிலிருந்து வெளியேறும் மூலமாகவோ பரவுகிறது. உங்களைப் பாதுகாத்துக் கொள்ள, ஆல்கஹால் சார்ந்த கை தடவினால் உங்கள் கைகளை அடிக்கடி சுத்தம் செய்யுங்கள் அல்லது சோப்பு மற்றும் தண்ணீரில் கழுவவும். மேலும், இருமல் மற்றும் தும்மக்கூடிய எவருடனும் நெருங்கிய தொடர்பைத் தவிர்க்கவும்.', 5),
    ('en-US', 'Hydroxychloroquine is official drug for corona treatment', 'There were early small studies in china and france which claimed it can benefit, and the drug was advocated in USA as well, ', 6);


INSERT INTO articles (locale, title, url, fact_id)
    VALUES
    ('en-US', 'covid is not human mande (with papers)', 'https://www.sciencenews.org/article/coronavirus-covid-19-not-human-made-lab-genetic-analysis-nature', 3),
    ('en-US', 'covid is not human mande (livescience blog)', 'https://www.livescience.com/coronavirus-not-human-made-in-lab.html', 3),
    ('ta', 'எந்த வயதினரும் கொரோனாவைப் பெறலாம் (who) கட்டுரை', 'https://www1.nyc.gov/assets/doh/downloads/pdf/imm/covid-19-daily-data-summary-deaths-04152020-1.pdf', 4),
    ('en-US', 'lancet: usage of Hydroxychloroquine could be harmful', 'https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)31174-0/fulltext', 6),
    ('en-US', 'Hydroxychloroquine does not prevent covid', 'https://www.statnews.com/2020/06/03/hydroxychloroquine-does-not-prevent-covid-19-infection-in-people-who-have-been-exposed-study-says/', 6),
    ('en-US', 'WHO resumes hydroxychloroquine study after reviewing safety concerns', 'https://www.statnews.com/2020/06/03/who-resuming-hydroxychloroquine-study-for-covid-19/', 6);


INSERT INTO funds (website, donate_url, title, description, image_url)
    VALUES 
    ('https://www.feedingindia.org', 'https://www.feedingindia.org/donate', 'Zomato FeedingIndia', 'Zomato initiated the "Feed the Daily Wager" project to provide food support to such families and to help them have a reliable supply of meals in the absence of employment opportunities', 'feedingindia.png'),
    ('https://indiafightscorona.giveindia.org', 'https://fundraisers.giveindia.org/donate/LC5ead82e1ae804#!/login', 'India Fights Corona', 'hygiene kits and meals to families hit by covid, site holds multiple fund campaigns', 'indiafightscorona.png'),
    ('https://www.pmcares.gov.in/', 'https://www.pmcares.gov.in/en/web/contribution/donate_india', 'PM Cares Fund', 'public charitable trust under the name of Prime Minister and emergency relief fund to provide relief to the affected', 'pmcares.png');

INSERT into medias (kind, category, title, description, url)
    VALUES
    ('image','awareness', 'How to stop the spread', 'protect others by using mask so you wont spread by sneeze, cough or while talking', 'https://www.who.int/images/default-source/health-topics/coronavirus/risk-communications/general-public/protect-yourself/blue-3.tmb-1920v.png?sfvrsn=b1ef6d45_5'),
    ('image','awareness', 'How to stop the spread', 'wash your hands thoroughly', 'https://www.who.int/images/default-source/health-topics/coronavirus/risk-communications/general-public/protect-yourself/blue-2.tmb-1920v.png?sfvrsn=2bc43de1_5'),
    ('image', 'awareness', 'Stay Home', 'Stay Home', 'https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/1169297c-9269-4549-9486-1e81bb85d263/ddu3ic0-d3fb4055-2468-4749-af84-41412668f31f.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvMTE2OTI5N2MtOTI2OS00NTQ5LTk0ODYtMWU4MWJiODVkMjYzXC9kZHUzaWMwLWQzZmI0MDU1LTI0NjgtNDc0OS1hZjg0LTQxNDEyNjY4ZjMxZi5qcGcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ._si3AzATZTAEbQ-ZX1KRTIX8tzqpAAOdzpg2DgSzZF4'),
    ('image', 'awareness', 'Stay Home', 'stay at home prescription', 'https://images.unsplash.com/photo-1588613254520-9d722c39aad5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1600&q=80'),
    ('image', 'awareness', 'Prevention', 'Wash your Hands', 'https://images.unsplash.com/photo-1588780456980-3032256f910a?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2167&q=80'),
    ('image', 'awareness', 'Information', 'Stay away from rumours','https://images.unsplash.com/photo-1588775566283-40d96c9d0e20?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1600&q=80'),
    ('image', 'awareness', 'Prevention', 'avoid touching your face', 'https://images.unsplash.com/photo-1588773922373-d8225f401dd1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=954&q=80')
    ;
