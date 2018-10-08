# 春歌服务程序

用于发送一些静态资源：如表情包、语音、固定的模板语句等。

## 命令列表

| 命令 | 参数 | 说明 |
| --- | --- | --- |
| haruka:echo | ... | 发出参数中的话 |
| haruka:hitokoto | 1 个选项，表示一言的分类，不指定则随机。-a (Anime), -c (Comic), -g (Game), -n (Novel), -i (Internet), -o (Other), -m (Nya) | 说一句话（基于[一言网](https://hitokoto.cn/)的 API |
| haruka:sticker | 1 个参数，表示贴纸的路径，如果是目录则从目录中随机 | 发送 `./data/sticker/` 路径下的贴纸 |
