package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := []string{(("\n\t" + "\n\t") + ""), "hello", "\n\t", "\n\t"}
    v2 := make(map[uint32]bool)
    var v3 int64 = 0
    var v4 int32 = 0
    fmt.Println(v4, v4)
    var v5 int64 = v3
    v2[218] = false
    fmt.Println(v5, v3)
    if (!false) {
        var v6 uint8 = 64
        if ((("\n\t" + ("a" + (("foo bar" + "a") + ("hello" + "\n\t")))) + "") != "hello") {
            fmt.Println(v6)
        } else {
            var v7 uint8 = v6
            v8 := GenS1{F0: 430, F1: 0, F2: 255}
            v4 = 1
            _ = v7
            _ = v8
        }
        if (!false) {
            var v9 float32 = -2.25
            v2[529] = (!true)
            v2[206] = (("a" + "foo bar") <= ("foo bar" + "foo bar"))
            _ = v9
        } else {
            if true {
                v5 = v3
                delete(v2, 1)
            } else {
                v3 -= v5
            }
            fmt.Println(v4, v5, v3)
        }
        _ = v6
    }
    v5 += 9223372036854775807
    var v10 uint32 = 626
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v4, v10, v4)
    v11 := GenS1{F0: 255, F1: 65535, F2: 2}
    if true {
        var v12 uint16 = 0
        if false {
            v13 := GenS2{F0: "foo bar", F1: 0, F2: "foo bar"}
            var v14 float32 = float32(((0.0 + (1.5 * 3.14159)) / ((float64(1.5) * -2.25) / 1e10)))
            _ = v13
            _ = v14
        } else {
            switch 32768 {
            case 2147483647:
                var v15 int = -73
                _ = v15
            case 9223372036854775807:
                var v16 int64 = 963
                _ = v16
            case 127:
                var v17 uint8 = 255
                _ = v17
            }
            var v18 int32 = (v4 << ((1) & 63))
            v1 = append(v1, "a")
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 uint8 = 206
                if true {
                    if ((((v21 << ((22) & 63)) >= (v21 - 2)) || (256 >= 546)) || ("hello" <= "foo bar")) {
                        v4 = (v4 | v4)
                        var v22 bool = (false || false)
                        _ = v22
                    } else {
                        if true {
                            fmt.Println(v4, v3)
                            var v23 float32 = 1.5
                            _ = v23
                        } else {
                            var v24 int32 = v4
                            _ = v24
                        }
                        var v25 int16 = 239
                        _ = v25
                    }
                    var v26 uint32 = (v10 << ((65535) & 63))
                    if (!true) {
                        v3 = 9223372036854775807
                    } else {
                        v26 -= (v10 / ((v26) | 1))
                        var v27 int = (fi20 | (fi19 << ((42) & 63)))
                        _ = v27
                    }
                    _ = v26
                } else {
                    v5 -= 526
                    v5 += (v3 % (((v3 * v3)) | 1))
                }
                v1 = append(v1, "foo bar")
                v28 := GenS2{F0: ((("foo bar" + "hello") + "hello") + ("\n\t" + ("a" + ("foo bar" + "a")))), F1: (^(v21)), F2: "\n\t"}
                _ = fi20
                _ = v21
                _ = v28
            }
            _ = fi19
        }
        _ = v12
    } else {
        var v29 string = "hello"
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v31 := GenS3{F0: (v3 << ((2) & 63)), F1: 42}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            delete(v2, v10)
            var v34 int32 = 877
            fmt.Println(fi32, v3)
            _ = fi33
            _ = v34
        }
        switch fi32 {
        case -9223372036854775808:
            v4 = (^(v4))
        case 512:
            v35 := GenS3{F0: (v5 / ((0) | 1)), F1: -74}
            _ = v35
        }
        fmt.Println(v5, v4, v5)
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
