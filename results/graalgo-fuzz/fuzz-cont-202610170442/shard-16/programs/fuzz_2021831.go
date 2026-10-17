package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string; F2 int8 }
func genFn0(p0 int, p1 bool, p2 uint64) int32 {
    p1 = p1
    p0 = (p0 ^ (p0 * (p0 * 0)))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return -64
}

func main() {
    var v1 int64 = -835
    var v2 uint = 350
    var v3 uint16 = 2
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    var v6 uint8 = 0
    v7 := []int{459, 9223372036854775807, 0}
    v8 := GenS1{F0: 127}
    var v9 int16 = -437
    switch "hello" {
    case "":
        switch "hello" {
        case "foo bar":
            if false {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v4[-2] = 1
            } else {
                var v10 float32 = ((0.0 + ((float32(-2.25) / 0.1) + (1e10 + (1.5 - (0.0 / float32(3.14159)))))) + 1e10)
                _ = v10
            }
        case "hello":
            var v11 int = -106
            _ = v11
        case "\n\t":
            v2 += v2
        }
    case "foo bar":
        _ = len(v7)
    case "a":
        v12 := GenS1{F0: -855}
        _ = v12
    default:
        var v13 float32 = (0.1 - (3.14159 * float32(1.5)))
        _ = v13
    }
    v5[""] = "foo bar"
    if true {
        if true {
            _ = len(v5)
            v1 = v1
            v2 += (^(v2))
        }
        var v14 int8 = 100
        v1 = (v1 | (v1 & v1))
        _ = v14
    } else {
        var v15 int8 = 42
        v16 := GenS1{F0: 32768}
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 127}
    switch "a" {
    case "hello":
        switch 347 {
        case -9223372036854775808:
            v7 = append(v7, 42)
        case 32767:
            v18 := GenS1{F0: -910}
            _ = v18
        case -815:
            for fi19 := 0; fi19 < 4; fi19++ {
                if ((v2 << ((1) & 63)) == (v2 & 2)) {
                    switch "foo bar" {
                    case "a":
                        var v20 int16 = v9
                        _ = v20
                    case "foo bar":
                        if false {
                            if true {
                                var v21 uint16 = (v3 - 255)
                                _ = v21
                            } else {
                                v6 = v6
                            }
                            v9 = (v9 * v9)
                            var v22 uint32 = 0
                            _ = v22
                        } else {
                            v5["hello"] = "hello"
                            v23 := GenS2{F0: fi19, F1: ("\n\t" + "\n\t"), F2: 0}
                            v9 += (^(v9))
                            _ = v23
                        }
                    case "hello":
                        var v24 uint32 = 2
                        _ = v24
                    }
                }
                _ = fi19
            }
        }
    case "a":
        v25 := GenS2{F0: 816, F1: ("a" + "\n\t"), F2: -22}
        _ = v25
    case "":
        var v26 int32 = -2147483648
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
