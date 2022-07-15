--default data for users
INSERT INTO users (id_user, username, password, id_line, owner_validation)
VALUES (1, 'patara', 'testpassword', 'Ue5308cc32ee5ca607c596e87877715b6', true);
INSERT INTO users (id_user, username, password, id_line, owner_validation)
VALUES (2, 'testUserWizane', 'testpassword2', 'U0d8f460619ae601aed723b4cab9c856b', true);

alter table wallet add wallet_name varchar default null;

--default data for wallet
INSERT INTO wallet (wallet_id, metamask_wallet_id,follow_wallet, id_user, wallet_name, id_chain, last_block_number)
VALUES (1, '0x891B68D6B21c64d56dB262D066B38Ea76B6468f6',false, 1, 'patara', 1, 0);
INSERT INTO wallet (wallet_id, metamask_wallet_id,follow_wallet, id_user, wallet_name, id_chain, last_block_number)
VALUES (2, '0x891B68D6B21c64d56dB262D066B38Ea76B6468f6',true, 1, 'freedom', 1, 0);