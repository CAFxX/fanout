package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 int16) int8 {
    p1 = (-(p1))
    var v5 int = int(p1)
    _ = p0
    _ = p1
    _ = v5
    return 127
}
func genFn1(p0 uint8, p1 int8, p2 int8) int64 {
    p0 -= (p0 * 114)
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = 42
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 503
}

func main() {
    var v1 int16 = 42
    var v2 float32 = 0.0
    var v3 float32 = 3.14159
    v4 := []int{-9223372036854775808}
    switch "foo bar" {
    case "hello":
        fmt.Println(v2, v1, v2)
    }
    v4 = append(v4, 581)
    if (true || false) {
        v7 := GenS1{F0: 678, F1: v1, F2: 18446744073709551615}
        _ = v7
    }
    switch ("\n\t" + ("\n\t" + ((("" + "") + "foo bar") + "hello"))) {
    case "a":
        if true {
            v1 = 1000
        } else {
            var v8 int64 = 1
            for fi9 := 0; fi9 < 20; fi9++ {
                if len(v4) > 0 {
                    v4[0] = (fi9 >> ((184) & 63))
                }
                if (805 > 786) {
                    var v10 float32 = v3
                    if true {
                        v11 := GenS1{F0: 255, F1: (v1 >> ((65535) & 63)), F2: 65535}
                        var v12 uint16 = 496
                        v8 = (v8 | genFn1(2, -128, 1))
                        _ = v11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 18; fi13++ {
                        if false {
                            v3 -= 0.1
                            if false {
                                v2 += -2.25
                            }
                        }
                        _ = len(v4)
                        _ = fi13
                    }
                    _ = v10
                }
                _ = fi9
            }
            _ = v8
        }
    case "hello":
        v2 = v3
    default:
        fmt.Println(v2)
    }
    v2 += v3
    fmt.Println(v2, v1, v1)
    if (127 != -1) {
        _ = len(v4)
        v14 := GenS2{F0: 42}
        _ = v14
    }
    if true {
        if (42 > 0) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v4)
            _ = fi15
        }
    } else {
        var v16 uint64 = 2
        var v17 uint = 1
        var v18 uint32 = 1
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
