package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8, p1 int64, p2 uint) string {
    fmt.Println(p1)
    v7 := []int{-291, -171}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("a" + ("a" + "\n\t"))
}
func genFn1(p0 uint64) uint32 {
    p0 = p0
    p0 = p0
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{707, 651, 2147483647}
    var v2 string = "foo bar"
    var v3 uint = 732
    v4 := make(map[string]string)
    var v5 float32 = -2.25
    v6 := make(map[int]uint32)
    _ = len(v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint16 = 624
    if true {
        fmt.Println(v2)
        v8 = 65535
        v5 = v5
    } else {
        var v9 string = ""
        v3 -= 1
        v2 = ""
        _ = v9
    }
    v10 := GenS1{F0: false}
    v8 -= (^(v8))
    switch "a" {
    case "":
        var v11 uint32 = 65535
        _ = v11
    case "a":
        switch v2 {
        case "":
            var v12 bool = (113 != 1)
            _ = v12
        case "hello":
            v13 := GenS1{F0: ((v5 / v5) < v5)}
            _ = v13
        default:
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        }
    }
    v3 = v3
    switch -9223372036854775808 {
    case 42:
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(v3, v3, v8)
            v5 = (0.0 - (1e10 + v5))
            v2 = "\n\t"
            _ = fi14
        }
    case 128:
        v15 := GenS1{F0: false}
        _ = v15
    default:
        var v16 uint8 = 127
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v3 = 341
        fmt.Println(fi17, v8, v2)
        _ = fi17
    }
    var v18 uint = 18446744073709551615
    switch 1 {
    case 128:
        fmt.Println(v18, v2)
    case 9223372036854775807:
        v5 = v5
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v2 = ((("\n\t" + v2) + ("hello" + v2)) + (v2 + v2))
        _ = fi19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 float64 = 0.1
        if true {
            if false {
                v22 := GenS1{F0: true}
                for fi23 := 0; fi23 < 20; fi23++ {
                    v4["\n\t"] = ""
                    v18 = v18
                    v5 += (1.5 * v5)
                    _ = fi23
                }
                _ = v22
            }
        } else {
            if (!false) {
                v5 = v5
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint16 = (v8 << ((v18) & 63))
                    v3 += 0
                    switch v2 {
                    case "hello":
                        var v26 int8 = -128
                        _ = v26
                    }
                    if true {
                        var v27 float32 = 1.5
                        _ = len(v4)
                        v21 = 0.0
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v8, v5, fi20)
            }
        }
        var v28 float32 = -2.25
        _ = fi20
        _ = v21
        _ = v28
    }
    var v29 int64 = 42
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
