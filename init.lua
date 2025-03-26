box.cfg{}

-- Создаем пользователя, если его нет
if not box.schema.user.exists('sampleuser') then
    box.schema.user.create('sampleuser', { password = '123456' })
    box.schema.user.grant('sampleuser', 'read,write,execute', 'universe')
end

-- Создаем пространство bands
bands = box.schema.space.create('bands', {
    if_not_exists = true
})

bands:format({
    { name = 'id', type = 'string' },
    { name = 'name', type = 'string' }
})

bands:create_index('primary', {
    parts = { 'id' },
    if_not_exists = true
})
