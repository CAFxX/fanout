package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint8; F2 uint8 }

func main() {
    v1 := []string{("foo bar" + ("" + ""))}
    var v2 int16 = -1
    var v3 float64 = 0.0
    var v4 uint8 = 1
    var v5 uint16 = 255
    var v6 int32 = 2147483647
    var v7 float32 = -2.25
    if true {
        var v8 float64 = v3
        _ = v8
    } else {
        if (1 != 127) {
            v9 := GenS1{F0: v4, F1: ((v2 >> ((65535) & 63)) > 1000), F2: (^(v2))}
            switch -472 {
            case 128:
                var v10 uint64 = 18446744073709551615
                _ = v10
            case -2147483648:
                v7 += v7
            }
            v11 := GenS1{F0: (v4 & (v4 ^ v4)), F1: false, F2: (^(v2))}
            _ = v11
            _ = v9
        } else {
            switch 128 {
            case 0:
                v3 = (v3 * (1.5 / v3))
            case -2147483648:
                _ = len(v1)
            case 79:
                var v12 uint32 = uint32(v5)
                _ = v12
            }
        }
        fmt.Println(v5)
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint64 = 2
            _ = fi14
            _ = v15
        }
        var v16 int16 = v2
        _ = fi13
        _ = v16
    }
    var v17 int8 = -128
    fmt.Println(v7, v17, v7)
    if (((v2 * v2) != (v2 - -412)) || true) {
        if true {
            if (32768 != 731) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    if len(v1) > 0 {
                        v1[0] = ("" + "hello")
                    }
                    v1 = append(v1, "a")
                    _ = fi18
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v4 -= (v4 | uint8(v5))
            switch 21 {
            case 128:
                for fi19 := 0; fi19 < 14; fi19++ {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        _ = fi20
                    }
                    if true {
                        v6 = 0
                        v21 := GenS1{F0: 0, F1: (v6 < -60), F2: v2}
                        v22 := GenS2{F0: v17}
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 8; fi23++ {
                        if false {
                            var v24 uint8 = 127
                            v5 -= (v5 << ((4294967295) & 63))
                            _ = v24
                        }
                        _ = fi23
                    }
                    _ = fi19
                }
            case 32768:
                switch -137 {
                case 4294967295:
                    switch ("hello" + "\n\t") {
                    case "":
                        v25 := GenS1{F0: 255, F1: (("hello" == "") && true), F2: v2}
                        _ = v25
                    default:
                        v26 := GenS3{F0: 0}
                        _ = v26
                    }
                case 32768:
                    switch 828 {
                    case -9223372036854775808:
                        v4 = v4
                    case 127:
                        v4 = (v4 ^ (v4 + v4))
                    case 32767:
                        v4 = v4
                    }
                case -741:
                    fmt.Println(v7, v6, v5)
                default:
                    var v27 uint = 4294967295
                    _ = v27
                }
            case 828:
                if (false || (42 < (^(v6)))) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v6 += (v6 % ((2147483647) | 1))
                }
            default:
                var v28 bool = (670 == 1)
                _ = v28
            }
        }
        v29 := GenS1{F0: (^(v4)), F1: false, F2: -1}
        _ = v29
    }
    v30 := GenS3{F0: 4294967295}
    var v31 int32 = v6
    var v32 float32 = float32(0.0)
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint64 = 1
            _ = fi34
            _ = v35
        }
        if (!false) {
            var v36 int = fi33
            var v37 int64 = 642
            if true {
                if (108 < 2147483648) {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                }
            }
            _ = v36
            _ = v37
        } else {
            _ = len(v1)
            var v38 int = fi33
            _ = v38
        }
        v39 := GenS4{F0: (fi33 >> ((389) & 63)), F1: v4, F2: 127}
        _ = fi33
        _ = v39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
