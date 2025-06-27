package com.example.test001;

import android.app.Activity;
import android.content.Intent;
import android.view.View;
import android.widget.Button;
import android.os.Bundle;


public class MainActivity extends Activity {
    private Button btn_startnewchat,btn_startoldchat,btn_startsetting;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main_liao);
        init();

        //跳转建立新聊天
        btn_startnewchat.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent=new Intent(MainActivity.this, ChatActivity.class);
                startActivity(intent);
                MainActivity.this.finish();
            }
        });

    }

    /**进行初始化操作
     */
    private void init(){
        //与xml的控件绑定
        btn_startnewchat=findViewById(R.id.btn_startnewchat);
//        btn_startoldchat=findViewById(R.id.btn_startoldchat);
//        btn_startsetting=findViewById(R.id.btn_startsetting);
    }

}