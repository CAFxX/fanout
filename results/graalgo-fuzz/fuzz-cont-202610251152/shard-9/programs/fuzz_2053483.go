package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 int64; F1 uint; F2 int64 }
func genFn0() int64 {
    var v7 int32 = -32
    v7 = v7
    _ = v7
    return -873
}
func genFn1() int8 {
    switch 590 {
    case 9223372036854775807:
        v8 := GenS1{F0: 2, F1: -53}
        _ = v8
    }
    v9 := GenS2{F0: 127, F1: -32768, F2: false}
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 uint = 65535
    v3 := make(map[string]string)
    var v4 bool = true
    v5 := []string{("hello" + ("" + "a"))}
    var v6 uint32 = 765
    v10 := GenS2{F0: 32767, F1: 42, F2: v4}
    if true {
        fmt.Println(v6, v4, v6)
        v6 = (v6 & v6)
        var v11 bool = (v4 || (32767 != -32768))
        _ = v11
    } else {
        if true {
            v12 := GenS2{F0: 9223372036854775807, F1: 1, F2: false}
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            if (((v4 || (!v4)) && v4) || (("a" + "a") == "\n\t")) {
                switch "a" {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                case "":
                    fmt.Println(v4, v6, v2)
                }
            }
            v6 += (v6 - v6)
            if (true && ((v2 <= v2) || false)) {
                if (v6 == 280) {
                    v2 -= (^(v2))
                }
                v6 = (v6 | (v6 * 42))
                v2 -= v2
            }
            _ = fi13
        }
        var v14 int8 = -128
        _ = v14
    }
    v4 = (v2 > (v2 << ((787) & 63)))
    var v15 uint32 = v6
    switch "foo bar" {
    case "\n\t":
        if v4 {
            var v16 uint64 = 289
            switch 666 {
            case 876:
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 uint64 = v16
                    v2 -= 415
                    v15 = (^(v6))
                    _ = fi17
                    _ = v18
                }
            case 2147483647:
                v19 := GenS1{F0: 29, F1: 100}
                _ = v19
            default:
                v16 = (v16 * (v16 / (((v16 - v16)) | 1)))
            }
            if true {
                v20 := GenS3{F0: -14, F1: 4294967295, F2: 9223372036854775807}
                _ = len(v5)
                _ = v20
            }
            _ = v16
        }
    }
    var v21 int64 = -414
    v21 = (v21 << ((255) & 63))
    _ = len(v1)
    if v4 {
        var v22 bool = (!(923 != 294))
        var v23 bool = v4
        v2 += (v2 - 180)
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
