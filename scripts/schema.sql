-- Drop table if exists
DROP TABLE IF EXISTS messages;

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20) NOT NULL,
    content VARCHAR(500) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' NOT NULL,
    message_id VARCHAR(100),
    sent_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create indexes for better query performance
CREATE INDEX idx_messages_status ON messages(status);
CREATE INDEX idx_messages_sent_at ON messages(sent_at);
CREATE INDEX idx_messages_created_at ON messages(created_at);

-- Add comment to table
COMMENT ON TABLE messages IS 'Stores messages to be sent via webhook';
COMMENT ON COLUMN messages.status IS 'Message status: pending, sent, or failed';
COMMENT ON COLUMN messages.content IS 'Message content - maximum 500 characters';
COMMENT ON COLUMN messages.message_id IS 'MessageID returned from webhook after successful send';

