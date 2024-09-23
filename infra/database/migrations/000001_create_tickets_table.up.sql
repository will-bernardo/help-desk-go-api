CREATE TABLE tickets (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID,
  circuit_id UUID,
  suport_id UUID,
  technician_id UUID,
  subject VARCHAR(60) NOT NULL,
  status VARCHAR(50) NOT NULL,
  description VARCHAR(1000) NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  created_at TIMESTAMPTZ DEFAULT NOW()
);
