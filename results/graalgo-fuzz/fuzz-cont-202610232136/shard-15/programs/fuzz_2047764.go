package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = 42
    var v2 int8 = 127
    v3 := []int{-624, 0}
    v4 := GenS1{F0: 255}
    var v5 int8 = -101
    v1 -= v1
    switch "hello" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 int8 = (v2 - (v5 | (v2 % ((v2) | 1))))
            _ = fi6
            _ = v7
        }
    default:
        var v8 uint64 = 988
        _ = v8
    }
    var v9 float64 = 0.1
    v10 := GenS1{F0: 109}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint32 = 65535
        if false {
            v5 = -118
            v5 = v5
        } else {
            fmt.Println(v1, v5)
            if false {
                var v13 int = (fi11 + fi11)
                v9 = v9
                var v14 float64 = float64((1.5 / ((-2.25 + 0.1) * 0.1)))
                _ = v13
                _ = v14
            } else {
                var v15 float64 = 0.1
                v9 += v9
                if false {
                    var v16 float32 = 3.14159
                    v1 = (v1 & v1)
                    _ = v16
                }
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v9)
    switch "" {
    case "":
        v17 := GenS2{F0: 65535}
        _ = v17
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            v5 -= 42
            _ = fi18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v3 = append(v3, -1)
        var v20 int64 = 1
        _ = len(v3)
        _ = fi19
        _ = v20
    }
    fmt.Println(v9, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v21 string = ""
        switch ((v21 + v21) + v21) {
        case "foo bar":
            if true {
                fmt.Println(v9)
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 float32 = (1e10 - (3.14159 * (((0.1 / 1.5) + 0.0) * 3.14159)))
                    _ = fi22
                    _ = v23
                }
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "\n\t":
            var v24 int8 = (^(v5))
            _ = v24
        case "a":
            fmt.Println(v1)
        default:
            fmt.Println(v9, v1)
        }
        if (1 != 65535) {
            v2 += (v5 & 127)
        }
        _ = v21
    } else {
        if len(v3) > 0 {
            v3[0] = 191
        }
        switch -9223372036854775808 {
        case 111:
            if false {
                v1 += (v1 * (-(v1)))
            }
        case -9223372036854775808:
            v9 = (float64((1e10 / -2.25)) + -2.25)
        case 42:
            var v25 int = -1
            _ = v25
        default:
            fmt.Println(v5, v9, v1)
        }
        v3 = append(v3, -305)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
