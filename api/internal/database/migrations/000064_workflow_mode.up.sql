-- Add transition mode to workflows: 'ordered' keeps transition rules,
-- 'free' allows any status-to-status change among defined statuses.
ALTER TABLE workflows
    ADD COLUMN mode TEXT NOT NULL DEFAULT 'ordered'
    CHECK (mode IN ('ordered', 'free'));
