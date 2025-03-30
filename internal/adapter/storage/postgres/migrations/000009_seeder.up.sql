INSERT INTO applicants (id, created_at, deleted_at, name, employment_status, marital_status, sex, date_of_birth)
VALUES ('b6c29c96-024b-4e70-834b-8e0dd2c66645', now(), NULL, 'Alice Cooper', 'employed', 'single', 'female',
        '1995-03-12'),
       ('c85062f2-e306-4ecd-b586-a3dcbb03deaf', now(), NULL, 'Bradley Cooper', 'unemployed', 'single', 'male',
        '1990-06-08'),
       ('2b73b011-9e5a-4d62-a645-2eccd739a75f', now(), NULL, 'Emma Watson', 'employed', 'married', 'female',
        '1987-04-15'),
       ('735b72e4-a0e8-4784-b3d3-b2caa90d6ac5', now(), NULL, 'Tom Hardy', 'employed', 'divorce', 'male',
        '1980-11-10'),
       ('6f47b906-bfe9-4c57-ad06-63e6e8871fcc', now(), NULL, 'Isabella Evans', 'unemployed', 'widowed', 'female',
        '1972-02-20'),
       ('55e83c42-97cf-461d-8565-5c4c742db5ee', now(), NULL, 'Chris Hemsworth', 'employed', 'married', 'male',
        '1983-08-11'),
       ('87992362-44ef-42e6-8a4c-d75c6080b71a', now(), NULL, 'Sophia Martinez', 'unemployed', 'single', 'female',
        '2001-07-18'),
       ('7ae93b68-62ff-4d76-8510-714ae8e2c257', now(), NULL, 'Liam Nelson', 'unemployed', 'single', 'male',
        '1992-12-05'),
       ('2b167178-46fc-494b-b17c-511c1ed385aa', now(), NULL, 'Grace White', 'employed', 'single', 'female',
        '1999-09-29');

INSERT INTO relationships (id, created_at, updated_at, applicant_a_id, applicant_b_id, relationship_type)
VALUES ('a12cf984-78f3-4eef-bb3c-66dae34a2fa1', now(), NULL, 'b6c29c96-024b-4e70-834b-8e0dd2c66645',
        'c85062f2-e306-4ecd-b586-a3dcbb03deaf', 'sibling'),
       ('ba76f9ac-b54a-4b27-9a73-3b9d6595c5ac', now(), NULL, '2b73b011-9e5a-4d62-a645-2eccd739a75f',
        '735b72e4-a0e8-4784-b3d3-b2caa90d6ac5', 'spouse'),
       ('3c98f267-dc08-43ee-bf55-90ee9d03ad26', now(), NULL, '6f47b906-bfe9-4c57-ad06-63e6e8871fcc',
        'c85062f2-e306-4ecd-b586-a3dcbb03deaf', 'parent'),
       ('4f83da47-cccd-4e81-8995-5e8d13e6f743', now(), NULL, '55e83c42-97cf-461d-8565-5c4c742db5ee',
        '87992362-44ef-42e6-8a4c-d75c6080b71a', 'parent'),
       ('0b9af7f5-98a0-4d52-bcf1-3b2fa2e6052e', now(), NULL, '7ae93b68-62ff-4d76-8510-714ae8e2c257',
        '2b167178-46fc-494b-b17c-511c1ed385aa', 'spouse'),
       ('a2436d69-fd62-47e5-bad3-a2c6148dfabb', now(), NULL, 'c85062f2-e306-4ecd-b586-a3dcbb03deaf',
        'b6c29c96-024b-4e70-834b-8e0dd2c66645', 'child'),
       ('e42935ec-ce5c-4f77-861d-1aebe3d91d42', now(), NULL, '2b167178-46fc-494b-b17c-511c1ed385aa',
        '6f47b906-bfe9-4c57-ad06-63e6e8871fcc', 'child');

INSERT INTO schemes (id, created_at, deleted_at, name)
VALUES ('7b555aaf-e824-4f75-a2dd-1ea59ec92310', now(), NULL, 'Elderly Support Scheme'),
       ('812b056b-8d12-472f-ac9b-419715a55b94', now(), NULL, 'Retrenchment Assistance Scheme'),
       ('c8296def-bd88-40d5-8071-b13ff147fb64', now(), NULL, 'Single Parent Support Scheme'),
       ('c8c699a7-8d59-40d7-8f9f-7f361804be40', now(), NULL, 'Retrenchment Assistance Scheme (families)');


INSERT INTO scheme_criteria (id, created_at, deleted_at, name, value, scheme_id)
VALUES ('1f399c51-335a-43c2-810f-9ee1f34e35b9', now(), NULL, 'age', '>=65', '7b555aaf-e824-4f75-a2dd-1ea59ec92310'),
       ('6579e01b-87ab-4e08-b4c5-86ecd50204fc', now(), NULL, 'employment_status', 'unemployed',
        '812b056b-8d12-472f-ac9b-419715a55b94'),
       ('6a04a785-71df-4fb5-9a65-ae0d7b007601', now(), NULL, 'has_children', 'true',
        'c8296def-bd88-40d5-8071-b13ff147fb64'),
       ('c391c72e-3bb7-4198-a436-039295b14468', now(), NULL, 'marital_status', 'single,widowed,divorce',
        'c8296def-bd88-40d5-8071-b13ff147fb64'),
       ('d87e9eae-b1a8-4cdc-8b96-5e0fa8c7297f', now(), NULL, 'employment_status', 'unemployed',
        'c8c699a7-8d59-40d7-8f9f-7f361804be40'),
       ('f257adf4-38ba-4fd3-a275-2ed4cc36ec7e', now(), NULL, 'has_children', 'true',
        'c8c699a7-8d59-40d7-8f9f-7f361804be40');

INSERT INTO benefits (id, created_at, deleted_at, scheme_id, name, amount)
VALUES ('0c4d2f12-fb86-4d8b-a165-4c6590e3decb', now(), NULL, '7b555aaf-e824-4f75-a2dd-1ea59ec92310',
        'Transport Allowance', 150),
       ('3352879d-1f0b-40f2-8611-c3985e750220', now(), NULL, '7b555aaf-e824-4f75-a2dd-1ea59ec92310', 'Medical Subsidy',
        300),
       ('39720afc-8f3a-46a7-9685-4be03cdaee9b', now(), NULL, 'c8c699a7-8d59-40d7-8f9f-7f361804be40', 'CDC Vouchers',
        300),
       ('585937e1-011a-42e6-b829-4c257a748874', now(), NULL, 'c8c699a7-8d59-40d7-8f9f-7f361804be40',
        'SkillsFuture Credits', 500),
       ('83886c2c-eab9-4abb-b313-02be0e4cd21e', now(), NULL, 'c8296def-bd88-40d5-8071-b13ff147fb64',
        'Housing Allowance', 600),
       ('91dd4122-6252-4db6-8430-582498ffe770', now(), NULL, '812b056b-8d12-472f-ac9b-419715a55b94',
        'SkillsFuture Credits', 500),
       ('ae72694b-d5a2-41de-b5f6-f2085ceaffce', now(), NULL, 'c8296def-bd88-40d5-8071-b13ff147fb64',
        'Childcare Subsidy', 400),
       ('c1643b56-a639-48dd-8743-e245141e179d', now(), NULL, 'c8c699a7-8d59-40d7-8f9f-7f361804be40',
        'School Meal Vouchers', 200),
       ('dfe6bef1-2d86-4b62-812b-72c09df1bbc3', now(), NULL, '812b056b-8d12-472f-ac9b-419715a55b94', 'CDC Vouchers',
        300);

INSERT INTO benefit_criteria (id, created_at, deleted_at, name, value, benefit_id)
VALUES ('1ec87355-dc32-4496-b902-4de5b5114a96', now(), NULL, 'has_primary_school_children', 'true',
        'c1643b56-a639-48dd-8743-e245141e179d');

INSERT INTO applications (id, created_at, applicant_id, scheme_id)
VALUES ('fe897b4f-568b-4ea1-8d95-99a91c97faf2', now(), 'c85062f2-e306-4ecd-b586-a3dcbb03deaf', '812b056b-8d12-472f-ac9b-419715a55b94')