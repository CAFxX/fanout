package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint32 = 1
    var v2 int = -1
    v3 := []string{"foo bar"}
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 8; fi5++ {
        if true {
            var v6 float32 = (3.14159 + (-2.25 * 0.0))
            var v7 int16 = -32768
            _ = v6
            _ = v7
        } else {
            switch "hello" {
            case "\n\t":
                if false {
                    if (true && ((false && false) && false)) {
                        var v8 uint = 2
                        v2 -= v2
                        v9 := GenS1{F0: 1, F1: 2}
                        _ = v8
                        _ = v9
                    } else {
                        var v10 uint64 = 803
                        for fi11 := 0; fi11 < 12; fi11++ {
                            if (!false) {
                                v10 += v10
                            }
                            v1 += (v1 / (((v1 >> ((576) & 63))) | 1))
                            v10 = v10
                            _ = fi11
                        }
                        _ = v10
                    }
                    if (!(2147483647 >= 587)) {
                        v2 += 697
                        for fi12 := 0; fi12 < 2; fi12++ {
                            var v13 uint32 = (v1 % ((v1) | 1))
                            _ = fi12
                            _ = v13
                        }
                        var v14 int16 = -32768
                        _ = v14
                    } else {
                        var v15 float32 = 3.14159
                        var v16 string = "hello"
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v15 = 1.5
                            _ = fi17
                        }
                        _ = v15
                        _ = v16
                    }
                    v2 = fi5
                } else {
                    v1 += 2
                    var v18 int = v2
                    fmt.Println(v2)
                    _ = v18
                }
            case "":
                v4[0] = fi5
            case "a":
                var v19 uint16 = 32767
                _ = v19
            default:
                _ = len(v4)
            }
            if (!(!(100 > -124))) {
                var v20 float32 = 3.14159
                v20 += (v20 + v20)
                var v21 string = ("hello" + ((("foo bar" + (("a" + "hello") + "foo bar")) + "foo bar") + "foo bar"))
                _ = v20
                _ = v21
            }
            var v22 bool = true
            _ = v22
        }
        _ = len(v3)
        _ = fi5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := v4[v1]
        v2 = (-(v24))
        var v25 uint = 65535
        _ = fi23
        _ = v24
        _ = v25
    }
    v2 += -2147483648
    switch (v2 << ((689) & 63)) {
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v1 += 641
    }
    _ = len(v3)
    v4[4294967295] = 993
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
