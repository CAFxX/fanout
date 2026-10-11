package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 string }
func genFn0(p0 bool) int64 {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint8 = 2
        _ = fi5
        _ = v6
    }
    switch 2147483647 {
    case 1:
        if p0 {
            switch 1 {
            case 128:
                var v7 int8 = -128
                _ = v7
            case 1:
                v8 := []int{507, 32767, 127}
                _ = v8
            case -9223372036854775808:
                switch 128 {
                case 127:
                    fmt.Println(p0)
                default:
                    v9 := []int{}
                    _ = v9
                }
            }
            if p0 {
                p0 = false
                v10 := GenS1{F0: "hello", F1: (1 == 0)}
                _ = v10
            }
        }
    case 2147483647:
        var v11 int64 = -1
        _ = v11
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint = 917
    fmt.Println(v1)
    if ("foo bar" == "a") {
        switch "" {
        case "hello":
            if (0.1 >= 1.5) {
                fmt.Println(v1, v1)
                var v12 uint16 = v1
                switch "\n\t" {
                case "hello":
                    if (-2147483648 == 0) {
                        var v13 int32 = -2147483648
                        _ = v13
                    } else {
                        fmt.Println(v4)
                    }
                case "foo bar":
                    var v14 uint32 = 57
                    _ = v14
                case "\n\t":
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS2{F0: "\n\t", F1: "a"}
                        v17 := GenS3{F0: ("" + "a")}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    switch 4294967295 {
                    case 32767:
                        v1 = 255
                    default:
                        _ = len(v3)
                    }
                }
                _ = v12
            } else {
                v4 = (v4 + (v4 % (((v4 & 0)) | 1)))
                v1 = 42
            }
        }
        _ = len(v2)
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    v4 = 896
    if (v1 < 2) {
        v4 -= v4
        var v18 int32 = -272
        var v19 int32 = (v18 + v18)
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: "a", F1: ("hello" + ("\n\t" + "\n\t"))}
    if false {
        var v21 uint16 = 1
        v3 = append(v3, (((((("\n\t" + "") + "a") + "foo bar") + "\n\t") + ("a" + ("a" + "\n\t"))) + ""))
        if false {
            switch 194 {
            case 32767:
                v1 = 2
            default:
                v1 += 0
            }
        }
        _ = v21
    }
    var v22 float32 = ((0.0 / (((-2.25 + 1.5) + float32(1e10)) / (3.14159 * -2.25))) - 0.0)
    _ = len(v3)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int = fi23
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
