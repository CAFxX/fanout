package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int16; F1 uint8; F2 int }
func genFn0() bool {
    var v7 uint64 = 796
    v8 := GenS1{F0: 1}
    switch -1 {
    case 127:
        fmt.Println(v7)
    case 2147483647:
        fmt.Println(v7)
    case 9223372036854775807:
        var v9 float32 = float32((0.1 + -2.25))
        _ = v9
    }
    _ = v7
    _ = v8
    return (-32768 > -696)
}
func genFn1(p0 int64) uint16 {
    if (2 > 255) {
        p0 = p0
    }
    var v10 int16 = 32767
    if false {
        var v11 bool = true
        _ = v11
    }
    v10 += v10
    _ = p0
    _ = v10
    return 32767
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{}
    var v3 float32 = 0.1
    v4 := make(map[int]int)
    var v5 uint32 = 2147483648
    var v6 uint64 = 607
    fmt.Println(v1, v5)
    var v12 bool = true
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: -49}
        var v15 int32 = 42
        var v16 string = "a"
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := v4[-1]
    v17 = v17
    fmt.Println(v17)
    v18 := GenS2{F0: -77, F1: 270, F2: ("a" + "foo bar")}
    var v19 bool = (v12 && v12)
    if (-140 <= -767) {
        switch "\n\t" {
        case "a":
            var v20 uint16 = 255
            _ = v20
        }
        if v12 {
            v21 := GenS2{F0: 7, F1: (v5 - v5), F2: ""}
            v22 := v4[-65]
            var v23 uint16 = 467
            _ = v21
            _ = v22
            _ = v23
        }
    }
    fmt.Println(v6, v19)
    var v24 int32 = -2147483648
    v5 += 2
    switch "\n\t" {
    case "":
        v25 := GenS2{F0: int8(v1), F1: 410, F2: ("\n\t" + ("a" + "foo bar"))}
        _ = v25
    case "a":
        v26 := GenS3{F0: v17, F1: "\n\t"}
        _ = v26
    case "hello":
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := GenS3{F0: -2147483648, F1: (((((((("" + ("\n\t" + "")) + "\n\t") + ("" + (("" + "a") + ""))) + "") + "") + ("foo bar" + "a")) + ("foo bar" + "hello")) + ("hello" + "foo bar"))}
                _ = fi28
                _ = v29
            }
            if true {
                v6 = (v6 / (((v6 + v6)) | 1))
                if len(v2) > 0 {
                    v2[0] = -1
                }
                for fi30 := 0; fi30 < 2; fi30++ {
                    _ = len(v2)
                    if v12 {
                        v3 = v3
                        switch v17 {
                        case -197:
                            v31 := GenS4{F0: 1, F1: 0, F2: 32767}
                            _ = v31
                        case -504:
                            v4[(fi27 & 42)] = (v17 << ((4294967295) & 63))
                        default:
                            var v32 int16 = 0
                            _ = v32
                        }
                    } else {
                        var v33 uint32 = (^(v5))
                        _ = v33
                    }
                    _ = fi30
                }
            }
            var v34 bool = true
            _ = fi27
            _ = v34
        }
    default:
        var v35 uint32 = (v5 + v5)
        _ = v35
    }
    v1 -= int64(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
