package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint64, p1 uint) int64 {
    v5 := GenS1{F0: (false && true)}
    p1 = (p1 << ((p1) & 63))
    var v6 float64 = float64((1.5 * ((-2.25 - 0.0) * (3.14159 - 0.1))))
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 bool) int32 {
    if p1 {
        v8 := GenS2{F0: 2147483648, F1: 950, F2: 18446744073709551615}
        _ = v8
    }
    var v9 int = -2147483648
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{1, 65535, 2}
    var v2 uint = 212
    var v3 int = 448
    v4 := make(map[string]uint32)
    v10 := GenS2{F0: 2147483647, F1: 1, F2: 42}
    v11 := GenS1{F0: true}
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v3)
    }
    v1 = append(v1, 377)
    v12 := GenS1{F0: false}
    v13 := GenS1{F0: false}
    _ = len(v1)
    v14 := GenS1{F0: true}
    for fi15 := 0; fi15 < 14; fi15++ {
        v2 -= 2
        _ = fi15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = fi16
    }
    v4["a"] = 0
    if false {
        v17 := GenS1{F0: (true || (299 > 92))}
        v3 -= v3
        _ = v17
    } else {
        fmt.Println(v3)
    }
    var v18 float64 = (1e10 - 0.1)
    if false {
        fmt.Println(v3)
        if true {
            var v19 uint16 = 94
            _ = v19
        } else {
            if false {
                var v20 float32 = 0.0
                switch ("" + "foo bar") {
                case "foo bar":
                    switch ("hello" + ("foo bar" + "\n\t")) {
                    case "a":
                        switch v3 {
                        case 42:
                            v21 := GenS1{F0: (true || false)}
                            _ = v21
                        default:
                            v22 := GenS3{F0: 278, F1: int32(v3)}
                            _ = v22
                        }
                    }
                case "":
                    var v23 string = "a"
                    _ = v23
                case "hello":
                    fmt.Println(v2, v18)
                }
                _ = v20
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v2 -= v2
                    _ = fi24
                }
                v3 = v3
                fmt.Println(v18, v2, v2)
            }
        }
        v18 += v18
    }
    var v25 uint = v2
    var v26 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
