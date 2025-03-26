box.cfg{
    listen = 3301
}

if box.space.bands == nil then
    box.schema.space.create('bands')
end


-- Specify field names and types --
box.space.bands:format({
    { name = 'key', type = 'string' },
    { name = 'value', type = 'string' },
    
})
