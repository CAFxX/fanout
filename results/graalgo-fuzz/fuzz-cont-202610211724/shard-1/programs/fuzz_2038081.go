package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 int32, p1 int32, p2 uint16) uint8 {
    fmt.Println(p1, p2, p1)
    p2 = p2
    p2 = 1
    v4 := GenS1{F0: 369, F1: p2}
    v5 := GenS2{F0: -1, F1: p0, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 string = "hello"
    var v2 int = -292
    v3 := make(map[uint32]uint32)
    v3[1] = 42
    if true {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := v3[65535]
            _ = fi6
            _ = v7
        }
        switch (v2 & 42) {
        case 42:
            if false {
                switch v1 {
                case "hello":
                    v8 := GenS3{F0: 32767}
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 float64 = (0.0 / 1.5)
                        _ = fi9
                        _ = v10
                    }
                }
                v2 = v2
            } else {
                if (1e10 != 0.0) {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v3[28] = 42
                        v1 = v1
                        _ = fi11
                    }
                }
                var v12 int32 = -204
                switch "foo bar" {
                case "":
                    v2 = (v2 >> ((4294967295) & 63))
                case "a":
                    delete(v3, 65535)
                case "hello":
                    delete(v3, 395)
                }
                _ = v12
            }
        case 127:
            v1 = (v1 + v1)
        case -1:
            var v13 int32 = -712
            _ = v13
        }
        v1 = ""
    }
    switch v2 {
    case -545:
        var v14 uint16 = 540
        _ = v14
    case 127:
        fmt.Println(v1, v1)
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        if (false || true) {
            if (18446744073709551615 != 354) {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v1 = (v1 + "\n\t")
                    _ = fi16
                }
            }
        } else {
            fmt.Println(v1, v2)
            v17 := GenS3{F0: 255}
            _ = v17
        }
        _ = fi15
    }
    var v18 int64 = 42
    switch (v2 % ((9223372036854775807) | 1)) {
    case 2147483647:
        var v19 uint = 18446744073709551615
        _ = v19
    case -9223372036854775808:
        v20 := GenS4{F0: 565, F1: 42}
        _ = v20
    case -190:
        switch ("foo bar" + "\n\t") {
        case "":
            var v21 uint8 = 2
            _ = v21
        case "\n\t":
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v2, v18)
                switch v1 {
                case "":
                    var v23 bool = false
                    _ = v23
                case "foo bar":
                    fmt.Println(v1, v2, v18)
                case "hello":
                    var v24 float64 = 0.0
                    _ = v24
                default:
                    v25 := v3[65535]
                    _ = v25
                }
                _ = fi22
            }
        }
    }
    v3[2] = 1
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 string = (("a" + "a") + v1)
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: 2, F1: 242}
    var v29 string = ("hello" + "a")
    if false {
        v29 = (((v1 + v1) + ("\n\t" + ("" + (v1 + v1)))) + v29)
        v1 = v29
        var v30 uint32 = 129
        _ = v30
    } else {
        v2 -= (v2 - (v2 ^ (v2 / (((v2 - v2)) | 1))))
        v31 := []int{v2, 1}
        v18 += v18
        _ = v31
    }
    var v32 int32 = 0
    v1 = (v29 + "foo bar")
    var v33 float64 = (((0.0 + (1.5 - (0.1 - (1e10 - -2.25)))) - 0.1) - (((0.0 + 0.0) / 0.1) - (-2.25 / 3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
