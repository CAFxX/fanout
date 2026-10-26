package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]string)
    var v3 uint64 = 932
    var v4 int8 = 0
    v5 := []uint32{308, 65535}
    for fi6 := 0; fi6 < 2; fi6++ {
        v2[(fi6 / (((-(fi6))) | 1))] = ("\n\t" + "foo bar")
        _ = fi6
    }
    var v7 int8 = -84
    if (-1 != -1) {
        var v8 float32 = (1e10 - (0.0 - ((3.14159 - 1.5) - -2.25)))
        _ = v8
    }
    if (!false) {
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            var v9 int16 = v1
            _ = v9
        case "":
            switch (("\n\t" + "\n\t") + ("hello" + "hello")) {
            case "a":
                fmt.Println(v7, v3)
            case "":
                if false {
                    var v10 uint64 = 217
                    _ = v10
                }
            default:
                if true {
                    v4 = v4
                }
            }
        case "hello":
            if true {
                var v11 bool = (0 > 0)
                var v12 uint32 = 2147483647
                if true {
                    var v13 uint32 = v12
                    if true {
                        var v14 int = -97
                        if (0.0 == (3.14159 + 3.14159)) {
                            fmt.Println(v11)
                        }
                        switch (v14 + (v14 % (((-(v14))) | 1))) {
                        case -2147483648:
                            var v15 int16 = v1
                            _ = v15
                        case 128:
                            _ = len(v5)
                        default:
                            v3 = v3
                        }
                        _ = v14
                    } else {
                        fmt.Println(v3, v13)
                    }
                    switch 42 {
                    case 0:
                        switch -9223372036854775808 {
                        case -604:
                            if (1e10 != 1.5) {
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v17 := GenS1{F0: v11, F1: (^(v1)), F2: -2147483648}
                                    _ = fi16
                                    _ = v17
                                }
                            } else {
                                v1 += (v1 << ((0) & 63))
                                v4 -= (v4 >> ((607) & 63))
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    for fi19 := 0; fi19 < 14; fi19++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v20 := GenS1{F0: (!v11), F1: -1, F2: -565}
                                            _ = v20
                                        case "hello":
                                            v4 = v7
                                        case "":
                                            fmt.Println(v11)
                                        }
                                        fmt.Println(v11, v3, v7)
                                        _ = fi19
                                    }
                                    var v21 uint64 = v3
                                    _ = fi18
                                    _ = v21
                                }
                            }
                        }
                    case -233:
                        v22 := GenS1{F0: true, F1: v1, F2: 0}
                        _ = v22
                    default:
                        if (1e10 <= 1.5) {
                            var v23 string = ""
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch -608 {
                            case 42:
                                var v24 uint32 = (v12 ^ (v13 | (v12 ^ (v12 % ((v13) | 1)))))
                                _ = v24
                            case 1:
                                var v25 int = 127
                                _ = v25
                            case -53:
                                if v11 {
                                    v4 = v4
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if (true || (!v11)) {
                                        if (1 > 32768) {
                                            v12 = (v13 - 2)
                                            v26 := GenS1{F0: (0.0 >= (float64((0.0 - float32(0.0))) + 1.5)), F1: v1, F2: -745}
                                            _ = v26
                                        }
                                    }
                                }
                            default:
                                switch v23 {
                                case "a":
                                    v27 := GenS1{F0: v11, F1: (v1 ^ 215), F2: 1}
                                    _ = v27
                                default:
                                    fmt.Println(v11)
                                }
                            }
                            _ = v23
                        } else {
                            v11 = v11
                        }
                    }
                    _ = v13
                } else {
                    v3 += v3
                }
                _ = v11
                _ = v12
            }
        }
        v4 = (v4 / (((v4 | v4)) | 1))
    }
    v7 -= v4
    _ = len(v2)
    var v28 uint16 = 65535
    switch "foo bar" {
    case "foo bar":
        if false {
            v29 := GenS1{F0: false, F1: v1, F2: 0}
            v1 = v1
            _ = v29
        } else {
            v30 := GenS1{F0: (!true), F1: (v1 - (v1 / ((v1) | 1))), F2: 1}
            v31 := GenS1{F0: (!(0.0 >= 1.5)), F1: (v1 & (v1 << ((uint(v28)) & 63))), F2: -1}
            v32 := GenS1{F0: false, F1: 0, F2: -2147483648}
            _ = v30
            _ = v31
            _ = v32
        }
    case "hello":
        v33 := GenS1{F0: (921 == 65535), F1: 0, F2: 709}
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v28 += v28
        var v35 uint = 433
        var v36 int64 = -1
        _ = fi34
        _ = v35
        _ = v36
    }
    v28 += 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
