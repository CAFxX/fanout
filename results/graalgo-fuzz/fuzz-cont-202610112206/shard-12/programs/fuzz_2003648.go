package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int, p1 bool, p2 uint8) string {
    var v5 float32 = 1.5
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint16) int8 {
    var v7 int8 = -1
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = 1000
        v10 := GenS1{F0: 135}
        for fi11 := 0; fi11 < 13; fi11++ {
            p0 -= 32767
            v12 := []int{(^(fi8)), (fi8 ^ (fi11 / ((4294967295) | 1))), (fi11 * fi8), -1}
            var v13 int8 = (v7 - -29)
            var v14 int = (^(fi8))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v15 int64 = -482
    _ = p0
    _ = v15
    _ = v7
    return v7
}

func main() {
    var v1 uint = 255
    v2 := []int{int(v1), 4294967295}
    v3 := make(map[string]uint32)
    var v4 bool = ((float32(float64(3.14159)) <= 1.5) && (("a" + "a") <= ("a" + ("\n\t" + ("" + ("hello" + ("foo bar" + "hello")))))))
    if (359 >= 65535) {
        _ = len(v2)
    } else {
        switch "foo bar" {
        case "":
            v16 := GenS2{F0: 0, F1: -1, F2: 42}
            _ = v16
        case "hello":
            v1 += (v1 << (((v1 | 0)) & 63))
        case "\n\t":
            var v17 uint64 = 18446744073709551615
            _ = v17
        }
        var v18 uint64 = 35
        _ = v18
    }
    switch "a" {
    case "foo bar":
        v4 = (-157 >= 0)
    default:
        switch 422 {
        case -9223372036854775808:
            fmt.Println(v4, v4)
        case 128:
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint64 = 1
                v21 := GenS1{F0: -1}
                switch "\n\t" {
                case "a":
                    v22 := GenS3{F0: 153, F1: ("hello" + ("" + "\n\t"))}
                    _ = v22
                case "foo bar":
                    if v4 {
                        v1 += v1
                    }
                case "hello":
                    if v4 {
                        v23 := GenS3{F0: 128, F1: "a"}
                        v3["hello"] = 973
                        _ = v23
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            v24 := GenS4{F0: 127}
            _ = v24
        }
    }
    v25 := GenS4{F0: 130}
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        switch "a" {
        case "a":
            for fi27 := 0; fi27 < 12; fi27++ {
                var v28 float64 = 0.0
                if (42 == 42) {
                    _ = len(v2)
                    _ = len(v2)
                    var v29 uint = 65535
                    _ = v29
                } else {
                    if len(v2) > 0 {
                        v2[0] = (fi26 & (fi27 | fi27))
                    }
                    var v30 uint64 = 927
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        }
        _ = fi26
    }
    var v31 uint8 = 64
    v4 = (!(!v4))
    v1 = (v1 >> ((v1) & 63))
    var v32 int64 = 422
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
