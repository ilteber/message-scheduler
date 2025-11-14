-- Comprehensive seed data for testing with clear status indicators
-- This includes multiple scenarios for thorough testing

-- ============================================================================
-- PHONE NUMBER 1: +905551234567 (65 messages - PAGINATION TEST)
-- Mix of sent, pending, and failed messages with various dates
-- ============================================================================

-- Sent messages from Nov 10, 2025 (25 messages)
INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905551234567', 'This message is supposed to be successfully sent (1/55)', 'sent', 'msg-001-' || gen_random_uuid(), '2025-11-10 08:00:00', '2025-11-10 07:55:00', '2025-11-10 08:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (2/55)', 'sent', 'msg-002-' || gen_random_uuid(), '2025-11-10 08:30:00', '2025-11-10 08:25:00', '2025-11-10 08:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (3/55)', 'sent', 'msg-003-' || gen_random_uuid(), '2025-11-10 09:00:00', '2025-11-10 08:55:00', '2025-11-10 09:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (4/55)', 'sent', 'msg-004-' || gen_random_uuid(), '2025-11-10 09:30:00', '2025-11-10 09:25:00', '2025-11-10 09:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (5/55)', 'sent', 'msg-005-' || gen_random_uuid(), '2025-11-10 10:00:00', '2025-11-10 09:55:00', '2025-11-10 10:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (6/55)', 'sent', 'msg-006-' || gen_random_uuid(), '2025-11-10 10:30:00', '2025-11-10 10:25:00', '2025-11-10 10:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (7/55)', 'sent', 'msg-007-' || gen_random_uuid(), '2025-11-10 11:00:00', '2025-11-10 10:55:00', '2025-11-10 11:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (8/55)', 'sent', 'msg-008-' || gen_random_uuid(), '2025-11-10 11:30:00', '2025-11-10 11:25:00', '2025-11-10 11:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (9/55)', 'sent', 'msg-009-' || gen_random_uuid(), '2025-11-10 12:00:00', '2025-11-10 11:55:00', '2025-11-10 12:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (10/55)', 'sent', 'msg-010-' || gen_random_uuid(), '2025-11-10 12:30:00', '2025-11-10 12:25:00', '2025-11-10 12:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (11/55)', 'sent', 'msg-011-' || gen_random_uuid(), '2025-11-10 13:00:00', '2025-11-10 12:55:00', '2025-11-10 13:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (12/55)', 'sent', 'msg-012-' || gen_random_uuid(), '2025-11-10 13:30:00', '2025-11-10 13:25:00', '2025-11-10 13:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (13/55)', 'sent', 'msg-013-' || gen_random_uuid(), '2025-11-10 14:00:00', '2025-11-10 13:55:00', '2025-11-10 14:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (14/55)', 'sent', 'msg-014-' || gen_random_uuid(), '2025-11-10 14:30:00', '2025-11-10 14:25:00', '2025-11-10 14:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (15/55)', 'sent', 'msg-015-' || gen_random_uuid(), '2025-11-10 15:00:00', '2025-11-10 14:55:00', '2025-11-10 15:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (16/55)', 'sent', 'msg-016-' || gen_random_uuid(), '2025-11-10 15:30:00', '2025-11-10 15:25:00', '2025-11-10 15:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (17/55)', 'sent', 'msg-017-' || gen_random_uuid(), '2025-11-10 16:00:00', '2025-11-10 15:55:00', '2025-11-10 16:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (18/55)', 'sent', 'msg-018-' || gen_random_uuid(), '2025-11-10 16:30:00', '2025-11-10 16:25:00', '2025-11-10 16:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (19/55)', 'sent', 'msg-019-' || gen_random_uuid(), '2025-11-10 17:00:00', '2025-11-10 16:55:00', '2025-11-10 17:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (20/55)', 'sent', 'msg-020-' || gen_random_uuid(), '2025-11-10 17:30:00', '2025-11-10 17:25:00', '2025-11-10 17:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (21/55)', 'sent', 'msg-021-' || gen_random_uuid(), '2025-11-10 18:00:00', '2025-11-10 17:55:00', '2025-11-10 18:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (22/55)', 'sent', 'msg-022-' || gen_random_uuid(), '2025-11-10 18:30:00', '2025-11-10 18:25:00', '2025-11-10 18:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (23/55)', 'sent', 'msg-023-' || gen_random_uuid(), '2025-11-10 19:00:00', '2025-11-10 18:55:00', '2025-11-10 19:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (24/55)', 'sent', 'msg-024-' || gen_random_uuid(), '2025-11-10 19:30:00', '2025-11-10 19:25:00', '2025-11-10 19:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (25/55)', 'sent', 'msg-025-' || gen_random_uuid(), '2025-11-10 20:00:00', '2025-11-10 19:55:00', '2025-11-10 20:00:00');

-- Sent messages from Nov 11, 2025 (20 messages)
INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905551234567', 'This message is supposed to be successfully sent (26/55)', 'sent', 'msg-026-' || gen_random_uuid(), '2025-11-11 07:00:00', '2025-11-11 06:55:00', '2025-11-11 07:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (27/55)', 'sent', 'msg-027-' || gen_random_uuid(), '2025-11-11 08:00:00', '2025-11-11 07:55:00', '2025-11-11 08:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (28/55)', 'sent', 'msg-028-' || gen_random_uuid(), '2025-11-11 09:00:00', '2025-11-11 08:55:00', '2025-11-11 09:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (29/55)', 'sent', 'msg-029-' || gen_random_uuid(), '2025-11-11 10:00:00', '2025-11-11 09:55:00', '2025-11-11 10:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (30/55)', 'sent', 'msg-030-' || gen_random_uuid(), '2025-11-11 11:00:00', '2025-11-11 10:55:00', '2025-11-11 11:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (31/55)', 'sent', 'msg-031-' || gen_random_uuid(), '2025-11-11 12:00:00', '2025-11-11 11:55:00', '2025-11-11 12:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (32/55)', 'sent', 'msg-032-' || gen_random_uuid(), '2025-11-11 13:00:00', '2025-11-11 12:55:00', '2025-11-11 13:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (33/55)', 'sent', 'msg-033-' || gen_random_uuid(), '2025-11-11 14:00:00', '2025-11-11 13:55:00', '2025-11-11 14:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (34/55)', 'sent', 'msg-034-' || gen_random_uuid(), '2025-11-11 15:00:00', '2025-11-11 14:55:00', '2025-11-11 15:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (35/55)', 'sent', 'msg-035-' || gen_random_uuid(), '2025-11-11 16:00:00', '2025-11-11 15:55:00', '2025-11-11 16:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (36/55)', 'sent', 'msg-036-' || gen_random_uuid(), '2025-11-11 17:00:00', '2025-11-11 16:55:00', '2025-11-11 17:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (37/55)', 'sent', 'msg-037-' || gen_random_uuid(), '2025-11-11 18:00:00', '2025-11-11 17:55:00', '2025-11-11 18:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (38/55)', 'sent', 'msg-038-' || gen_random_uuid(), '2025-11-11 19:00:00', '2025-11-11 18:55:00', '2025-11-11 19:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (39/55)', 'sent', 'msg-039-' || gen_random_uuid(), '2025-11-11 20:00:00', '2025-11-11 19:55:00', '2025-11-11 20:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (40/55)', 'sent', 'msg-040-' || gen_random_uuid(), '2025-11-11 21:00:00', '2025-11-11 20:55:00', '2025-11-11 21:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (41/55)', 'sent', 'msg-041-' || gen_random_uuid(), '2025-11-11 22:00:00', '2025-11-11 21:55:00', '2025-11-11 22:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (42/55)', 'sent', 'msg-042-' || gen_random_uuid(), '2025-11-11 22:30:00', '2025-11-11 22:25:00', '2025-11-11 22:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (43/55)', 'sent', 'msg-043-' || gen_random_uuid(), '2025-11-11 23:00:00', '2025-11-11 22:55:00', '2025-11-11 23:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (44/55)', 'sent', 'msg-044-' || gen_random_uuid(), '2025-11-11 23:30:00', '2025-11-11 23:25:00', '2025-11-11 23:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (45/55)', 'sent', 'msg-045-' || gen_random_uuid(), '2025-11-11 23:59:00', '2025-11-11 23:54:00', '2025-11-11 23:59:00');

-- Sent messages from Nov 12, 2025 (10 messages)
INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905551234567', 'This message is supposed to be successfully sent (46/55)', 'sent', 'msg-046-' || gen_random_uuid(), '2025-11-12 06:00:00', '2025-11-12 05:55:00', '2025-11-12 06:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (47/55)', 'sent', 'msg-047-' || gen_random_uuid(), '2025-11-12 08:00:00', '2025-11-12 07:55:00', '2025-11-12 08:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (48/55)', 'sent', 'msg-048-' || gen_random_uuid(), '2025-11-12 10:00:00', '2025-11-12 09:55:00', '2025-11-12 10:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (49/55)', 'sent', 'msg-049-' || gen_random_uuid(), '2025-11-12 12:00:00', '2025-11-12 11:55:00', '2025-11-12 12:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (50/55)', 'sent', 'msg-050-' || gen_random_uuid(), '2025-11-12 12:30:00', '2025-11-12 12:25:00', '2025-11-12 12:30:00'),
('+905551234567', 'This message is supposed to be successfully sent (51/55)', 'sent', 'msg-051-' || gen_random_uuid(), '2025-11-12 14:00:00', '2025-11-12 13:55:00', '2025-11-12 14:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (52/55)', 'sent', 'msg-052-' || gen_random_uuid(), '2025-11-12 16:00:00', '2025-11-12 15:55:00', '2025-11-12 16:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (53/55)', 'sent', 'msg-053-' || gen_random_uuid(), '2025-11-12 18:00:00', '2025-11-12 17:55:00', '2025-11-12 18:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (54/55)', 'sent', 'msg-054-' || gen_random_uuid(), '2025-11-12 20:00:00', '2025-11-12 19:55:00', '2025-11-12 20:00:00'),
('+905551234567', 'This message is supposed to be successfully sent (55/55)', 'sent', 'msg-055-' || gen_random_uuid(), '2025-11-12 20:30:00', '2025-11-12 20:25:00', '2025-11-12 20:30:00');

-- Failed messages (5 messages)
INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905551234567', 'This message is supposed to fail (1/5) - testing error handling', 'failed', '2025-11-13 08:00:00', '2025-11-13 08:00:30'),
('+905551234567', 'This message is supposed to fail (2/5) - testing error handling', 'failed', '2025-11-13 10:00:00', '2025-11-13 10:00:30'),
('+905551234567', 'This message is supposed to fail (3/5) - testing error handling', 'failed', '2025-11-13 12:00:00', '2025-11-13 12:00:30'),
('+905551234567', 'This message is supposed to fail (4/5) - testing error handling', 'failed', '2025-11-13 14:00:00', '2025-11-13 14:00:30'),
('+905551234567', 'This message is supposed to fail (5/5) - testing error handling', 'failed', '2025-11-13 16:00:00', '2025-11-13 16:00:30');

-- Pending messages for +905551234567 (10 messages)
INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905551234567', 'This message is supposed to be pending (1/31) - waiting for scheduler', 'pending', '2025-11-14 08:00:00', '2025-11-14 08:00:00'),
('+905551234567', 'This message is supposed to be pending (2/31) - waiting for scheduler', 'pending', '2025-11-14 08:30:00', '2025-11-14 08:30:00'),
('+905551234567', 'This message is supposed to be pending (3/31) - waiting for scheduler', 'pending', '2025-11-14 09:00:00', '2025-11-14 09:00:00'),
('+905551234567', 'This message is supposed to be pending (4/31) - waiting for scheduler', 'pending', '2025-11-14 09:30:00', '2025-11-14 09:30:00'),
('+905551234567', 'This message is supposed to be pending (5/31) - waiting for scheduler', 'pending', '2025-11-14 10:00:00', '2025-11-14 10:00:00'),
('+905551234567', 'This message is supposed to be pending (6/31) - waiting for scheduler', 'pending', '2025-11-14 10:30:00', '2025-11-14 10:30:00'),
('+905551234567', 'This message is supposed to be pending (7/31) - waiting for scheduler', 'pending', '2025-11-14 11:00:00', '2025-11-14 11:00:00'),
('+905551234567', 'This message is supposed to be pending (8/31) - waiting for scheduler', 'pending', '2025-11-14 11:30:00', '2025-11-14 11:30:00'),
('+905551234567', 'This message is supposed to be pending (9/31) - waiting for scheduler', 'pending', '2025-11-14 12:00:00', '2025-11-14 12:00:00'),
('+905551234567', 'This message is supposed to be pending (10/31) - waiting for scheduler', 'pending', '2025-11-14 12:30:00', '2025-11-14 12:30:00');

-- ============================================================================
-- PHONE NUMBER 2: +905559876543 (15 messages - Mixed statuses)
-- ============================================================================

INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905559876543', 'This message is supposed to be successfully sent (1/5)', 'sent', 'msg-100-' || gen_random_uuid(), '2025-11-10 09:00:00', '2025-11-10 08:55:00', '2025-11-10 09:00:00'),
('+905559876543', 'This message is supposed to be successfully sent (2/5)', 'sent', 'msg-101-' || gen_random_uuid(), '2025-11-10 12:00:00', '2025-11-10 11:55:00', '2025-11-10 12:00:00'),
('+905559876543', 'This message is supposed to be successfully sent (3/5)', 'sent', 'msg-102-' || gen_random_uuid(), '2025-11-11 10:00:00', '2025-11-11 09:55:00', '2025-11-11 10:00:00'),
('+905559876543', 'This message is supposed to be successfully sent (4/5)', 'sent', 'msg-103-' || gen_random_uuid(), '2025-11-11 14:00:00', '2025-11-11 13:55:00', '2025-11-11 14:00:00'),
('+905559876543', 'This message is supposed to be successfully sent (5/5)', 'sent', 'msg-104-' || gen_random_uuid(), '2025-11-12 08:00:00', '2025-11-12 07:55:00', '2025-11-12 08:00:00');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905559876543', 'This message is supposed to fail (1/2) - testing error handling', 'failed', '2025-11-13 10:00:00', '2025-11-13 10:00:30'),
('+905559876543', 'This message is supposed to fail (2/2) - testing error handling', 'failed', '2025-11-13 14:00:00', '2025-11-13 14:00:30');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905559876543', 'This message is supposed to be pending (11/31) - waiting for scheduler', 'pending', '2025-11-14 08:00:00', '2025-11-14 08:00:00'),
('+905559876543', 'This message is supposed to be pending (12/31) - waiting for scheduler', 'pending', '2025-11-14 09:00:00', '2025-11-14 09:00:00'),
('+905559876543', 'This message is supposed to be pending (13/31) - waiting for scheduler', 'pending', '2025-11-14 10:00:00', '2025-11-14 10:00:00'),
('+905559876543', 'This message is supposed to be pending (14/31) - waiting for scheduler', 'pending', '2025-11-14 11:00:00', '2025-11-14 11:00:00'),
('+905559876543', 'This message is supposed to be pending (15/31) - waiting for scheduler', 'pending', '2025-11-14 12:00:00', '2025-11-14 12:00:00'),
('+905559876543', 'This message is supposed to be pending (16/31) - waiting for scheduler', 'pending', '2025-11-14 13:00:00', '2025-11-14 13:00:00'),
('+905559876543', 'This message is supposed to be pending (17/31) - waiting for scheduler', 'pending', '2025-11-14 14:00:00', '2025-11-14 14:00:00'),
('+905559876543', 'This message is supposed to be pending (18/31) - waiting for scheduler', 'pending', '2025-11-14 15:00:00', '2025-11-14 15:00:00');

-- ============================================================================
-- PHONE NUMBER 3: +905557778888 (12 messages - Mostly successful)
-- ============================================================================

INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905557778888', 'This message is supposed to be successfully sent (1/10)', 'sent', 'msg-200-' || gen_random_uuid(), '2025-11-09 10:00:00', '2025-11-09 09:55:00', '2025-11-09 10:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (2/10)', 'sent', 'msg-201-' || gen_random_uuid(), '2025-11-09 14:00:00', '2025-11-09 13:55:00', '2025-11-09 14:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (3/10)', 'sent', 'msg-202-' || gen_random_uuid(), '2025-11-10 11:00:00', '2025-11-10 10:55:00', '2025-11-10 11:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (4/10)', 'sent', 'msg-203-' || gen_random_uuid(), '2025-11-10 11:05:00', '2025-11-10 11:00:00', '2025-11-10 11:05:00'),
('+905557778888', 'This message is supposed to be successfully sent (5/10)', 'sent', 'msg-204-' || gen_random_uuid(), '2025-11-10 15:00:00', '2025-11-10 14:55:00', '2025-11-10 15:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (6/10)', 'sent', 'msg-205-' || gen_random_uuid(), '2025-11-11 09:00:00', '2025-11-11 08:55:00', '2025-11-11 09:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (7/10)', 'sent', 'msg-206-' || gen_random_uuid(), '2025-11-12 10:00:00', '2025-11-12 09:55:00', '2025-11-12 10:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (8/10)', 'sent', 'msg-207-' || gen_random_uuid(), '2025-11-13 08:00:00', '2025-11-13 07:55:00', '2025-11-13 08:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (9/10)', 'sent', 'msg-208-' || gen_random_uuid(), '2025-11-13 15:00:00', '2025-11-13 14:55:00', '2025-11-13 15:00:00'),
('+905557778888', 'This message is supposed to be successfully sent (10/10)', 'sent', 'msg-209-' || gen_random_uuid(), '2025-11-13 18:00:00', '2025-11-13 17:55:00', '2025-11-13 18:00:00');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905557778888', 'This message is supposed to be pending (19/31) - waiting for scheduler', 'pending', '2025-11-14 10:00:00', '2025-11-14 10:00:00'),
('+905557778888', 'This message is supposed to be pending (20/31) - waiting for scheduler', 'pending', '2025-11-14 11:00:00', '2025-11-14 11:00:00');

-- ============================================================================
-- PHONE NUMBER 4: +905553334444 (8 messages - Some failed)
-- ============================================================================

INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905553334444', 'This message is supposed to be successfully sent (1/3)', 'sent', 'msg-300-' || gen_random_uuid(), '2025-11-12 09:00:00', '2025-11-12 08:55:00', '2025-11-12 09:00:00'),
('+905553334444', 'This message is supposed to be successfully sent (2/3)', 'sent', 'msg-301-' || gen_random_uuid(), '2025-11-12 09:15:00', '2025-11-12 09:10:00', '2025-11-12 09:15:00'),
('+905553334444', 'This message is supposed to be successfully sent (3/3)', 'sent', 'msg-302-' || gen_random_uuid(), '2025-11-12 14:00:00', '2025-11-12 13:55:00', '2025-11-12 14:00:00');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905553334444', 'This message is supposed to fail (1/3) - testing error handling', 'failed', '2025-11-13 14:00:00', '2025-11-13 14:00:30'),
('+905553334444', 'This message is supposed to fail (2/3) - testing error handling', 'failed', '2025-11-13 15:00:00', '2025-11-13 15:00:30'),
('+905553334444', 'This message is supposed to fail (3/3) - testing error handling', 'failed', '2025-11-13 16:00:00', '2025-11-13 16:00:30');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905553334444', 'This message is supposed to be pending (21/31) - waiting for scheduler', 'pending', '2025-11-14 09:00:00', '2025-11-14 09:00:00'),
('+905553334444', 'This message is supposed to be pending (22/31) - waiting for scheduler', 'pending', '2025-11-14 12:00:00', '2025-11-14 12:00:00');

-- ============================================================================
-- PHONE NUMBER 5: +905556667777 (10 messages - Recent activity)
-- ============================================================================

INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905556667777', 'This message is supposed to be successfully sent (1/5)', 'sent', 'msg-400-' || gen_random_uuid(), '2025-11-13 08:00:00', '2025-11-13 07:55:00', '2025-11-13 08:00:00'),
('+905556667777', 'This message is supposed to be successfully sent (2/5)', 'sent', 'msg-401-' || gen_random_uuid(), '2025-11-13 10:00:00', '2025-11-13 09:55:00', '2025-11-13 10:00:00'),
('+905556667777', 'This message is supposed to be successfully sent (3/5)', 'sent', 'msg-402-' || gen_random_uuid(), '2025-11-13 14:00:00', '2025-11-13 13:55:00', '2025-11-13 14:00:00'),
('+905556667777', 'This message is supposed to be successfully sent (4/5)', 'sent', 'msg-403-' || gen_random_uuid(), '2025-11-13 16:00:00', '2025-11-13 15:55:00', '2025-11-13 16:00:00'),
('+905556667777', 'This message is supposed to be successfully sent (5/5)', 'sent', 'msg-404-' || gen_random_uuid(), '2025-11-13 18:00:00', '2025-11-13 17:55:00', '2025-11-13 18:00:00');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905556667777', 'This message is supposed to be pending (23/31) - waiting for scheduler', 'pending', '2025-11-14 07:00:00', '2025-11-14 07:00:00'),
('+905556667777', 'This message is supposed to be pending (24/31) - waiting for scheduler', 'pending', '2025-11-14 10:00:00', '2025-11-14 10:00:00'),
('+905556667777', 'This message is supposed to be pending (25/31) - waiting for scheduler', 'pending', '2025-11-14 12:00:00', '2025-11-14 12:00:00'),
('+905556667777', 'This message is supposed to be pending (26/31) - waiting for scheduler', 'pending', '2025-11-14 14:00:00', '2025-11-14 14:00:00'),
('+905556667777', 'This message is supposed to be pending (27/31) - waiting for scheduler', 'pending', '2025-11-14 16:00:00', '2025-11-14 16:00:00');

-- ============================================================================
-- PHONE NUMBER 6: +905552221111 (7 messages - Minimal activity)
-- ============================================================================

INSERT INTO messages (phone_number, content, status, message_id, sent_at, created_at, updated_at) VALUES
('+905552221111', 'This message is supposed to be successfully sent (1/3)', 'sent', 'msg-500-' || gen_random_uuid(), '2025-11-11 10:00:00', '2025-11-11 09:55:00', '2025-11-11 10:00:00'),
('+905552221111', 'This message is supposed to be successfully sent (2/3)', 'sent', 'msg-501-' || gen_random_uuid(), '2025-11-11 11:00:00', '2025-11-11 10:55:00', '2025-11-11 11:00:00'),
('+905552221111', 'This message is supposed to be successfully sent (3/3)', 'sent', 'msg-502-' || gen_random_uuid(), '2025-11-12 09:00:00', '2025-11-12 08:55:00', '2025-11-12 09:00:00');

INSERT INTO messages (phone_number, content, status, created_at, updated_at) VALUES
('+905552221111', 'This message is supposed to be pending (28/31) - waiting for scheduler', 'pending', '2025-11-14 09:00:00', '2025-11-14 09:00:00'),
('+905552221111', 'This message is supposed to be pending (29/31) - waiting for scheduler', 'pending', '2025-11-14 13:00:00', '2025-11-14 13:00:00'),
('+905552221111', 'This message is supposed to be pending (30/31) - waiting for scheduler', 'pending', '2025-11-14 15:00:00', '2025-11-14 15:00:00'),
('+905552221111', 'This message is supposed to be pending (31/31) - waiting for scheduler', 'pending', '2025-11-14 17:00:00', '2025-11-14 17:00:00');

-- ============================================================================
-- SUMMARY STATS:
-- Phone +905551234567: 70 messages (55 sent, 5 failed, 10 pending)
-- Phone +905559876543: 15 messages (5 sent, 2 failed, 8 pending)
-- Phone +905557778888: 12 messages (10 sent, 0 failed, 2 pending)
-- Phone +905553334444: 8 messages (3 sent, 3 failed, 2 pending)
-- Phone +905556667777: 10 messages (5 sent, 0 failed, 5 pending)
-- Phone +905552221111: 7 messages (3 sent, 0 failed, 4 pending)
-- 
-- TOTAL: 122 messages across 6 phone numbers
-- SENT: 81 messages | FAILED: 10 messages | PENDING: 31 messages
-- ============================================================================
