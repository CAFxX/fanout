package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) bool {
    v5 := GenS1{F0: 2, F1: p0, F2: -1}
    p0 += (p0 >> ((950) & 63))
    v6 := []int{9223372036854775807, 32767}
    p0 = (p0 % ((980) | 1))
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint16) uint {
    fmt.Println(p0)
    if false {
        if true {
            if ((("hello" + "hello") + "hello") == "\n\t") {
                v7 := []int{4294967295, 1}
                fmt.Println(p0)
                if true {
                    if (1e10 > (0.1 / 1.0)) {
                        var v8 uint32 = 2
                        _ = v8
                    }
                }
                _ = v7
            }
            if (!false) {
                p0 -= p0
                var v9 bool = (100 != 117)
                _ = v9
            }
            var v10 int8 = -1
            _ = v10
        } else {
            v11 := []int{42, 4294967295, -287}
            fmt.Println(p0)
            _ = v11
        }
    }
    if (1 > -177) {
        v12 := []int{}
        if (true || false) {
            fmt.Println(p0)
            p0 -= 1
            p0 = 346
        } else {
            switch 9223372036854775807 {
            case 953:
                if (!(("\n\t" + (("foo bar" + "a") + "hello")) > "a")) {
                    _ = len(v12)
                    var v13 float32 = 3.14159
                    _ = v13
                }
            case 32768:
                p0 -= p0
            default:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            v14 := GenS2{F0: 128, F1: uint8(p0), F2: -32768}
            p0 = 414
            _ = v14
        }
        p0 = 65535
        _ = v12
    }
    _ = p0
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 9223372036854775807
    v3 := []int{726, 4294967295, 128, -2147483648}
    var v4 uint32 = 901
    v15 := GenS2{F0: 1, F1: 1, F2: -32768}
    var v16 int64 = -107
    if true {
        v2 = (v2 ^ v2)
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint32 = 101
                v19 = v19
                var v20 uint64 = 42
                v16 += v2
                if false {
                    var v21 uint = 2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(fi17)
                    _ = v21
                }
                v4 += v4
                _ = fi18
                _ = v19
                _ = v20
            }
            v22 := GenS3{F0: -2147483648, F1: 42}
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 float64 = 3.14159
                v2 -= v16
                v25 := v1["hello"]
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi17
            _ = v22
        }
        v4 -= v4
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            if true {
                v27 := GenS3{F0: -2147483648, F1: 1}
                switch "" {
                case "hello":
                    var v28 int32 = 2147483647
                    _ = v28
                case "":
                    fmt.Println(fi26)
                }
                switch -2147483648 {
                case 9223372036854775807:
                    if true {
                        fmt.Println(fi26, v4)
                    }
                case 0:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v3 = append(v3, fi26)
                        v2 = 725
                        _ = fi29
                    }
                default:
                    v1[""] = "hello"
                }
                _ = v27
            }
            _ = fi26
        }
        _ = len(v1)
    }
    v1["a"] = "hello"
    switch ("hello" + ("hello" + "hello")) {
    case "":
        _ = len(v1)
    }
    v30 := GenS2{F0: 42, F1: 42, F2: 0}
    var v31 int64 = (v2 / ((v2) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
