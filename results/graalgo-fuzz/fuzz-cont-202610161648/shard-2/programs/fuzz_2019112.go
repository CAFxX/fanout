package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool; F1 int }

func main() {
    var v1 uint = 65535
    var v2 float32 = 0.0
    v3 := []uint32{}
    var v4 float64 = 1.5
    if ("" <= "") {
        switch "hello" {
        case "\n\t":
            var v5 uint = (v1 / (((v1 * 4294967295)) | 1))
            _ = v5
        case "":
            for fi6 := 0; fi6 < 17; fi6++ {
                if (true && true) {
                    switch (fi6 * fi6) {
                    case 32768:
                        for fi7 := 0; fi7 < 11; fi7++ {
                            var v8 int32 = 42
                            if (!false) {
                                if (true && (100 > -14)) {
                                    var v9 int16 = 42
                                    _ = v9
                                }
                                fmt.Println(v4)
                                var v10 bool = true
                                _ = v10
                            }
                            if (104 > 1) {
                                switch 2147483647 {
                                case 9223372036854775807:
                                    var v11 bool = (959 < 1)
                                    _ = v11
                                default:
                                    var v12 int16 = -388
                                    _ = v12
                                }
                                var v13 int64 = 9223372036854775807
                                var v14 uint8 = 150
                                _ = v13
                                _ = v14
                            }
                            _ = fi7
                            _ = v8
                        }
                    case -520:
                        var v15 uint8 = 222
                        _ = v15
                    }
                }
                _ = fi6
            }
        case "hello":
            for fi16 := 0; fi16 < 15; fi16++ {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 float32 = 1e10
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        default:
            if false {
                v1 += (v1 - (v1 & (v1 | (^(v1)))))
            }
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            switch (fi19 / (((fi19 ^ fi19)) | 1)) {
            case 32767:
                v2 -= 3.14159
            default:
                switch (fi19 % (((fi19 + 987)) | 1)) {
                case -2147483648:
                    switch "" {
                    case "hello":
                        v2 += 0.0
                    default:
                        if (1 == 0) {
                            var v20 float64 = v4
                            if len(v3) > 0 {
                                v3[0] = 683
                            }
                            _ = v20
                        }
                    }
                case 127:
                    var v21 string = ("" + (("\n\t" + "") + ("foo bar" + (((("foo bar" + "a") + ("foo bar" + ("\n\t" + "foo bar"))) + "a") + ("hello" + "a")))))
                    _ = v21
                default:
                    var v22 int64 = 0
                    _ = v22
                }
            }
            _ = fi19
        }
        if false {
            switch 2147483647 {
            case -1:
                v23 := GenS1{F0: 336}
                _ = v23
            case 127:
                if (!(2 == 127)) {
                    var v24 uint = (v1 % ((0) | 1))
                    for fi25 := 0; fi25 < 18; fi25++ {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v4, v2)
                        default:
                            v26 := GenS1{F0: 1}
                            _ = v26
                        }
                        v27 := GenS2{F0: 0, F1: 2147483647}
                        _ = fi25
                        _ = v27
                    }
                    var v28 string = "a"
                    _ = v24
                    _ = v28
                }
            case 863:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
            fmt.Println(v4, v4)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v3 = append(v3, 341)
    var v29 uint32 = 570
    v30 := GenS1{F0: -1}
    var v31 uint8 = 128
    v32 := GenS1{F0: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ((("foo bar" + "") + ("\n\t" + "\n\t")) + ("a" + ((("a" + "hello") + ("hello" + (("hello" + "hello") + ""))) + "foo bar"))) {
    case "a":
        if len(v3) > 0 {
            v3[0] = (v29 % (((v29 >> ((4294967295) & 63))) | 1))
        }
    case "":
        if true {
            fmt.Println(v1, v31, v29)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v2, v29, v1)
        } else {
            var v33 uint32 = v29
            switch 128 {
            case -2147483648:
                for fi34 := 0; fi34 < 4; fi34++ {
                    if true {
                        v1 += v1
                        var v35 uint32 = 42
                        _ = v35
                    }
                    if (0 > 42) {
                        var v36 int32 = 134
                        v37 := GenS1{F0: 921}
                        _ = v36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 10; fi38++ {
                        var v39 int = fi38
                        v2 += -2.25
                        v40 := GenS1{F0: 1}
                        v41 := GenS3{F0: false, F1: v39}
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = fi34
                }
            default:
                v33 += uint32(v1)
            }
            _ = v33
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
