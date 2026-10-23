package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0() int64 {
    var v7 uint32 = 407
    v7 = (v7 - 42)
    var v8 uint = 4294967295
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64, p1 string, p2 int16) int32 {
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint16 = 410
        _ = fi9
        _ = v10
    }
    var v11 bool = (656 <= 18446744073709551615)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2147483647
}

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 uint = 1
    var v4 uint8 = 78
    v5 := []string{"\n\t", "hello", "foo bar", "hello"}
    var v6 int8 = 100
    if ((!true) && v2) {
        v6 = -91
        v1 = append(v1, "\n\t")
        v12 := GenS1{F0: 65535}
        _ = v12
    }
    if false {
        if (2147483647 < 32768) {
            fmt.Println(v3, v2)
        } else {
            var v13 uint = 255
            _ = len(v1)
            _ = v13
        }
        switch 0 {
        case -2147483648:
            switch "hello" {
            case "":
                if len(v5) > 0 {
                    v5[0] = "a"
                }
            }
        case 32767:
            v14 := GenS1{F0: 775}
            _ = v14
        }
        if (2147483647 > 2147483647) {
            fmt.Println(v2, v2, v2)
            if ((v6 >> ((v3) & 63)) != v6) {
                var v15 string = ("" + (((("a" + "hello") + "\n\t") + "foo bar") + "foo bar"))
                var v16 int32 = -1
                _ = v15
                _ = v16
            } else {
                fmt.Println(v6, v3)
                v17 := GenS1{F0: 255}
                _ = v17
            }
            fmt.Println(v3, v3)
        }
    }
    v4 = 0
    fmt.Println(v3, v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v2 {
        if ((v2 && v2) || (v2 && (!true))) {
            v3 -= (v3 * v3)
            var v18 uint16 = 255
            v2 = (!v2)
            _ = v18
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                switch "" {
                case "a":
                    var v20 uint64 = 186
                    _ = v20
                case "":
                    var v21 int = 32767
                    _ = v21
                case "foo bar":
                    fmt.Println(v2, v6)
                default:
                    var v22 int16 = 481
                    _ = v22
                }
                _ = fi19
            }
            if false {
                v2 = v2
            }
            v6 = v6
        }
        if true {
            v3 = (v3 | (v3 - v3))
            v5 = append(v5, "")
        }
        v23 := GenS2{F0: 1000, F1: 498, F2: v4}
        _ = v23
    }
    var v24 uint16 = 922
    fmt.Println(v6, v24)
    var v25 float32 = (float32(1e10) / 1.0)
    var v26 int32 = 0
    var v27 int8 = v6
    v28 := GenS2{F0: -866, F1: 2, F2: 42}
    var v29 int64 = -1
    fmt.Println(v27, v24, v29)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := GenS3{F0: int16(v26), F1: 4294967295, F2: genFn1(2, "a", 1)}
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 uint16 = (v24 & v24)
            fmt.Println(v29)
            if (-2.25 < 3.14159) {
                v27 = (v6 >> ((42) & 63))
                v34 := GenS4{F0: 42, F1: -9223372036854775808, F2: v26}
                var v35 float32 = v25
                _ = v34
                _ = v35
            }
            v36 := GenS5{F0: 770, F1: (v33 % (((v24 * uint16(v4))) | 1))}
            _ = fi32
            _ = v33
            _ = v36
        }
        _ = fi30
        _ = v31
    }
    if len(v5) > 0 {
        v5[0] = ((("" + "\n\t") + "a") + "hello")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
