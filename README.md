# bladeTemplate for Yii2

在`config/web.php`中添加一下代码
```php
'components' => [
        //...other components

        'view'         => [
            'class'            => 'Wjunz\BladeView',
            'defaultExtension' => 'blade.php',
            'renderers'        => [
                'blade' => [
                    'class'     => 'Wjunz\ViewRenderer',
                    'cachePath' => '@runtime/cache',
                ],
            ],
        ],

        //...other components
]
```