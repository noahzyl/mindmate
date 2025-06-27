package com.example.test001;

import androidx.appcompat.app.AppCompatActivity;

import android.annotation.SuppressLint;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;

public class aaMainActivity extends AppCompatActivity {

    public static final String ACCOUNT = "account";
    public static final String PASSWORD = "password";

    private static final int DEFAULT_DIALER_REQUEST_ID = 83;
    public static final String TAG = "MADARA";
    String  name="";
    String  code="";

    private ImageView bt_see_title_img;
    private ImageView bt_chosed_title_img;
    private ImageView add_information_img;
    private ImageView bt_back_img;


    @SuppressLint("MissingInflatedId")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.aaactivity_main);

        Intent intent = getIntent();
        name = intent.getStringExtra(ACCOUNT);
        code = intent.getStringExtra(PASSWORD);

        bt_see_title_img = (ImageView) findViewById(R.id.see_title_img);
        bt_chosed_title_img = (ImageView) findViewById(R.id.chosed_title_img);
        add_information_img = (ImageView) findViewById(R.id.information_img);
        bt_back_img = (ImageView) findViewById(R.id.back_img);

        bt_see_title_img.setOnClickListener(new View.OnClickListener() {

            public void onClick(View v) {
                Intent intent1 = new Intent(aaMainActivity.this, aaMainActivity.class);
                startActivity(intent1);

            }
        });

        bt_chosed_title_img.setOnClickListener(new View.OnClickListener() {

            public void onClick(View v) {
                Intent intent1 = new Intent(aaMainActivity.this, aaMainActivity.class);
                startActivity(intent1);

            }
        });


        add_information_img.setOnClickListener(new View.OnClickListener() {

            public void onClick(View v) {
                Intent intent2 = new Intent(aaMainActivity.this, aaMainActivity.class);
                intent2.putExtra("NAME", name);
                intent2.putExtra("CODE", code);
                startActivity(intent2);

            }
        });

        bt_back_img.setOnClickListener(new View.OnClickListener() {

            public void onClick(View v) {
                Intent intent2 = new Intent(aaMainActivity.this, LoginActivity.class);
                startActivity(intent2);
            }
        });

    }

}

