-- Create a space --
box.schema.space.create('bands')
-- Specify field names and types --
box.space.MyDB:format({
    { name = 'key', type = 'string' },
    { name = 'value', type = 'string' },
    
})
