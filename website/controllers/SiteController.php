<?php

namespace app\controllers;

use Yii;
use yii\filters\AccessControl;
use yii\helpers\ArrayHelper;
use yii\web\Controller;
use yii\web\Response;
use yii\filters\VerbFilter;
use app\models\LoginForm;
use app\models\ContactForm;

class SiteController extends Controller
{
    /**
     * {@inheritdoc}
     */
    public function behaviors()
    {
        return [
            'access' => [
                'class' => AccessControl::className(),
                'only' => ['logout'],
                'rules' => [
                    [
                        'actions' => ['logout'],
                        'allow' => true,
                        'roles' => ['@'],
                    ],
                ],
            ],
            'verbs' => [
                'class' => VerbFilter::className(),
                'actions' => [
                    'logout' => ['post'],
                ],
            ],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function actions()
    {
        return [
            'error' => [
                'class' => 'yii\web\ErrorAction',
            ],
            'captcha' => [
                'class' => 'yii\captcha\CaptchaAction',
                'fixedVerifyCode' => YII_ENV_TEST ? 'testme' : null,
            ],
        ];
    }

    /**
     * Displays homepage.
     *
     * @return string
     */
    public function actionIndex()
    {

        return $this->render('index');
    }

    public function actionApi(){

        $client = \Yii::$app->clickhouse;
        $sql = 'select * from data order by EventTime desc limit 25';
        $data = $client->createCommand($sql )->queryAll();

        $dt['Pressure'] = ArrayHelper::getColumn($data,'Pressure');
       $dt['Humidity'] = ArrayHelper::getColumn($data,'Humidity');
        $dt['TemperatureR'] = ArrayHelper::getColumn($data,'TemperatureR');
        $dt['TemperatureA'] = ArrayHelper::getColumn($data,'TemperatureA');
        $dt['pH'] = ArrayHelper::getColumn($data,'pH');
        $dt['FlowRate'] = ArrayHelper::getColumn($data,'FlowRate');
        $dt['CO'] = ArrayHelper::getColumn($data,'CO');
        $dt['EventTime'] = ArrayHelper::getColumn($data,'EventTime');
        return json_encode($dt);
    }


}
