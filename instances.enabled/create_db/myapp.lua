box.cfg{}

if not box.space.bands then
    local bands = box.schema.space.create('bands', { if_not_exists = true })

    bands:format({
	{ name = 'id', type = 'int' },
        { name = 'key', type = 'string' },
        { name = 'value', type = 'string' }
    })

    bands:create_index('primary', { parts = {1, 'string'}, unique = true })
end

