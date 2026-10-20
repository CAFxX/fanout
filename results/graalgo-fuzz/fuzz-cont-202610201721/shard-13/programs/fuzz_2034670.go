package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    p0 += (p0 + 2147483647)
    if false {
        p0 += p0
        switch "foo bar" {
        case "hello":
            v7 := GenS1{F0: 1}
            _ = v7
        case "foo bar":
            fmt.Println(p0)
        default:
            var v8 int16 = 1
            _ = v8
        }
    }
    _ = p0
    return -494
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 0
    v3 := []int{-1, 1}
    v4 := []string{}
    var v5 float32 = 3.14159
    var v6 int = 42
    var v9 string = ("" + "a")
    if len(v4) > 0 {
        v4[0] = "a"
    }
    switch v9 {
    case "\n\t":
        if true {
            var v10 string = (v9 + v9)
            _ = v10
        } else {
            v5 += v5
            var v11 uint = 18446744073709551615
            _ = v11
        }
    case "foo bar":
        v1 += 18446744073709551615
    }
    switch "hello" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        v5 = -2.25
    case "a":
        var v12 string = v9
        _ = v12
    default:
        v13 := GenS2{F0: 42}
        _ = v13
    }
    v14 := GenS2{F0: 21}
    v15 := GenS1{F0: (v2 ^ (v2 ^ v2))}
    for fi16 := 0; fi16 < 4; fi16++ {
        switch fi16 {
        case -730:
            if true {
                switch v9 {
                case "hello":
                    if true {
                        var v17 float64 = 0.0
                        var v18 int8 = 0
                        _ = v17
                        _ = v18
                    }
                case "a":
                    v4 = append(v4, v9)
                default:
                    if false {
                        v5 += (float32(0.1) - v5)
                    } else {
                        _ = len(v3)
                        v6 += int(v1)
                        if false {
                            v1 -= (v1 % ((v1) | 1))
                        }
                    }
                }
                fmt.Println(v1, v1)
                for fi19 := 0; fi19 < 20; fi19++ {
                    v6 -= int(v1)
                    _ = fi19
                }
            }
        case 128:
            if (0 <= 0) {
                var v20 uint64 = v1
                _ = v20
            }
        }
        var v21 uint64 = (v1 + (v1 % ((18446744073709551615) | 1)))
        _ = fi16
        _ = v21
    }
    var v22 uint8 = (v2 - 127)
    var v23 uint = 18446744073709551615
    if (!false) {
        if (int(v1) > 127) {
            v2 = 42
            var v24 float32 = v5
            var v25 int32 = 42
            _ = v24
            _ = v25
        } else {
            v4 = append(v4, "foo bar")
            fmt.Println(v1, v9)
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int8 = 121
                v6 -= (fi26 * v6)
                _ = fi26
                _ = v27
            }
        }
    }
    var v28 uint64 = (v1 - (v1 + v1))
    var v29 int = (v6 + (v6 + (-(v6))))
    v6 += (v6 * int(v23))
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int16 = int16(v6)
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 116}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
