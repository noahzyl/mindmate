package com.example.test001;

import android.content.Context;
import android.speech.tts.TextToSpeech;
import android.util.Log;
import android.view.View;

import java.util.Locale;

/**
这个文件没用
 */
public class AutoReader {
    private TextToSpeech tts;
    private boolean isInitialized = false;
    public AutoReader(Context context){
        // 初始化 TextToSpeech 实例
        tts = new TextToSpeech(context,  new TextToSpeech.OnInitListener(){
            @Override
            public void onInit(int status) {
                if (status == TextToSpeech.SUCCESS) {
                    int result = tts.setLanguage(Locale.US);

                    if (result == TextToSpeech.LANG_MISSING_DATA
                            || result == TextToSpeech.LANG_NOT_SUPPORTED) {
                        Log.e("TextSpeaker", "This Language is not supported");
                    } else {
                        isInitialized = true;
                    }
                } else {
                    Log.e("TextSpeaker", "Initilization Failed!");
                }
            }
        });
    }
    public void ReadOut(String text){
        if (isInitialized && !text.isEmpty()) {
            tts.speak(text, TextToSpeech.QUEUE_FLUSH, null, null);
        } else {
            Log.w("TextSpeaker", "TextSpeaker is not initialized or text is empty.");
        }
    }
}
