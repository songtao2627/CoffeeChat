import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

/// 主动呼叫
class PageAVChatCallerStatefulWidget extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _PageAVChatCallerWidgetState();
}

class _PageAVChatCallerWidgetState extends State<PageAVChatCallerStatefulWidget> {
  @override
  Widget build(BuildContext context) {
    return _buildCaller();
  }

  /// 呼叫
  Widget _buildCaller() {
    return Scaffold(
      body: Container(
        decoration: BoxDecoration(
            // background
            gradient: LinearGradient(
          colors: [Color(0xFF442923), Color(0xFF2b2723), Color(0xFF442923)],
          begin: Alignment.topCenter,
          end: Alignment.bottomCenter,
        )),
        child: Column(children: <Widget>[
//          Padding(
//            padding: EdgeInsets.only(top: 60, left: 20),
//            child: Row(
//              children: <Widget>[
//                Container(
//                  width: 55,
//                  child: RaisedButton(
//                    child: Icon(
//                      Icons.call_received,
//                      color: Colors.white70,
//                    ),
//                    color: Colors.transparent,
//                    onPressed: _onMinimum,
//                    shape: RoundedRectangleBorder(
//                        side: BorderSide(color: Colors.white), borderRadius: BorderRadius.all(Radius.circular(10))),
//                  ),
//                ),
//              ],
//            ),
//          ),
          Container(height: 90),
          Column(
            children: <Widget>[
              Padding(padding: EdgeInsets.only(top: 50)),
              Image.asset(
                "assets/default_avatar.png",
                width: 128,
                height: 128,
                fit: BoxFit.fill,
              ),
              Padding(padding: EdgeInsets.only(top: 20)),
              Text(
                "三生三世十里桃花",
                style: new TextStyle(fontSize: 20, color: Colors.white),
              ),
              Padding(padding: EdgeInsets.only(top: 5)),
              Text("正在等待对方接受邀请...", style: new TextStyle(fontSize: 16, color: Colors.white70))
            ],
          ),
          Expanded(child: Container()),
          Container(
            height: 200,
            child: Row(
              children: <Widget>[
                Expanded(
                  child: Column(
                    children: <Widget>[
                      Container(
                        height: 80,
                        child: RaisedButton(
                          child: Icon(
                            Icons.mic_off,
                            color: Colors.white,
                            size: 30,
                          ),
                          color: Colors.transparent,
                          elevation: 60,
                          shape: CircleBorder(side: BorderSide(color: Colors.white)),
                          onPressed: _onMute,
                        ),
                      ),
                      Padding(
                        child: Text("静 音", style: new TextStyle(fontSize: 14, color: Colors.white)),
                        padding: EdgeInsets.only(top: 10),
                      )
                    ],
                  ),
                ),
                Expanded(
                  child: Column(
                    children: <Widget>[
                      Container(
                        height: 80,
                        child: RaisedButton(
                          child: Icon(
                            Icons.call_end,
                            color: Colors.white,
                            size: 30,
                          ),
                          color: Colors.red,
                          elevation: 60,
                          shape: CircleBorder(),
                          onPressed: _onCancel,
                        ),
                      ),
                      Padding(
                        child: Text("取 消", style: new TextStyle(fontSize: 14, color: Colors.white)),
                        padding: EdgeInsets.only(top: 10),
                      )
                    ],
                  ),
                ),
                Expanded(
                  child: Column(
                    children: <Widget>[
                      Container(
                        height: 80,
                        child: RaisedButton(
                          child: Icon(
                            Icons.volume_up,
                            color: Colors.white,
                            size: 30,
                          ),
                          color: Colors.transparent,
                          elevation: 60,
                          shape: CircleBorder(side: BorderSide(color: Colors.white)),
                          onPressed: _onHandsFree,
                        ),
                      ),
                      Padding(
                        child: Text("免 提", style: new TextStyle(fontSize: 14, color: Colors.white)),
                        padding: EdgeInsets.only(top: 10),
                      )
                    ],
                  ),
                ),
              ],
            ),
          ),
        ]),
      ),
    );
  }

  void _onCancel() {
    Navigator.of(this.context).pop();
  }

  void _onMute() {}

  void _onMinimum() {}

  void _onHandsFree() {}
}