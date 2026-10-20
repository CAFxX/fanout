package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{42, 386, 282, -2147483648}
    var v3 int32 = -1
    v4 := make(map[string]string)
    var v5 int64 = 636
    var v6 int64 = v5
    fmt.Println(v5, v3)
    v7 := GenS1{F0: 255}
    switch 32767 {
    case 128:
        var v8 int64 = v5
        _ = v8
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            v6 = 9223372036854775807
            _ = fi9
        }
    }
    v5 = v5
    var v10 uint = 167
    v6 += (v6 >> ((v10) & 63))
    v11 := v1[441]
    v4["a"] = ""
    v12 := GenS2{F0: (v11 | v11), F1: 255, F2: -11}
    v3 = v3
    if (false && false) {
        var v13 uint64 = 285
        var v14 uint = v10
        _ = v13
        _ = v14
    }
    var v15 float64 = (((0.1 / 3.14159) - -2.25) - float64((1e10 / 1.0)))
    if false {
        if (((v15 / (v15 - (v15 - ((3.14159 - (((v15 - v15) / 3.14159) - (v15 * v15))) * v15)))) / ((float64(1.5) * v15) / -2.25)) != 1e10) {
            fmt.Println(v15, v10)
            v10 -= v10
        }
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 bool = (false && false)
            fmt.Println(v5, v6, v5)
            for fi18 := 0; fi18 < 17; fi18++ {
                switch ("\n\t" + "foo bar") {
                case "hello":
                    if v17 {
                        var v19 uint64 = 42
                        var v20 int = fi18
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS2{F0: 786, F1: 448, F2: 10}
                        _ = v21
                    }
                case "foo bar":
                    v22 := GenS2{F0: int(v5), F1: 801, F2: 1}
                    _ = v22
                case "":
                    if v17 {
                        if len(v2) > 0 {
                            v2[0] = (fi18 % ((42) | 1))
                        }
                        v3 += v3
                        v2 = append(v2, (fi18 ^ 32768))
                    }
                default:
                    switch ("\n\t" + ("\n\t" + ("hello" + ("a" + "a")))) {
                    case "a":
                        var v23 int8 = 0
                        _ = v23
                    default:
                        fmt.Println(v10)
                    }
                }
                var v24 bool = (!v17)
                v25 := GenS1{F0: 443}
                v26 := GenS2{F0: (v11 & v11), F1: 2, F2: 109}
                _ = fi18
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi16
            _ = v17
        }
        v27 := GenS3{F0: 2, F1: 104}
        _ = v27
    }
    fmt.Println(v15)
    if true {
        _ = len(v1)
    } else {
        for fi28 := 0; fi28 < 19; fi28++ {
            switch "\n\t" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                v5 = v5
            case "hello":
                v5 = v5
            default:
                fmt.Println(v15, v10, v5)
            }
            v29 := GenS3{F0: 233, F1: 2147483647}
            _ = fi28
            _ = v29
        }
        if ((0 == 4294967295) || ((("\n\t" + "") > "") && (true && true))) {
            v11 -= v11
        }
        if (false && (("a" == (("hello" + "foo bar") + ("" + "\n\t"))) && false)) {
            v5 -= (v5 | v5)
            if ((-1 >= int8(v3)) && (128 < 19)) {
                v6 = (v5 - -1)
                var v30 bool = (v10 >= (v10 / (((v10 + (^(v10)))) | 1)))
                v6 -= (v6 << (((v10 % (((^(v10))) | 1))) & 63))
                _ = v30
            } else {
                var v31 float32 = (3.14159 + 3.14159)
                switch ("\n\t" + "hello") {
                case "":
                    var v32 uint64 = 0
                    _ = v32
                case "hello":
                    var v33 uint64 = 918
                    _ = v33
                case "\n\t":
                    v5 = (-(v5))
                default:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        fmt.Println(v15, v3)
                        v3 -= v3
                        _ = fi34
                    }
                }
                for fi35 := 0; fi35 < 3; fi35++ {
                    if (v10 >= v10) {
                        v4["hello"] = ""
                        var v36 uint16 = 255
                        var v37 string = ""
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int16 = 0
                        var v39 uint8 = 255
                        _ = v38
                        _ = v39
                    }
                    var v40 uint = (^(v10))
                    _ = len(v2)
                    _ = fi35
                    _ = v40
                }
                _ = v31
            }
            v41 := GenS2{F0: v11, F1: 658, F2: -128}
            _ = v41
        }
    }
    v1[4294967295] = (v11 % (((-(v11))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
