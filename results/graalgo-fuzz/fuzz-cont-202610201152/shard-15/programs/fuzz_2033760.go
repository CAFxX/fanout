package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    switch -594 {
    case 1:
        var v6 float32 = 0.0
        _ = v6
    case 4294967295:
        var v7 string = "a"
        _ = v7
    default:
        var v8 uint8 = 127
        _ = v8
    }
    _ = p0
    return 6
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 0.1
    v3 := []uint32{820, 2147483647, 2147483647, 4294967295}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    if ((42 > 42) && true) {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint16 = 108
            v2 = v2
            if true {
                var v11 uint64 = 2
                var v12 uint8 = 2
                _ = v11
                _ = v12
            } else {
                v1 -= 0
                switch "a" {
                case "foo bar":
                    v2 = (((3.14159 - ((v2 / -2.25) * v2)) / (v2 + v2)) * v2)
                case "\n\t":
                    v1 = (v1 + -1)
                case "a":
                    v2 = v2
                }
                switch 42 {
                case 1:
                    delete(v5, 42)
                default:
                    var v13 uint32 = 65535
                    _ = v13
                }
            }
            _ = fi9
            _ = v10
        }
        var v14 int16 = 32767
        _ = v14
    }
    v15 := GenS1{F0: ((-(v1)) >= v1)}
    for fi16 := 0; fi16 < 12; fi16++ {
        v1 += 0
        v1 -= v1
        fmt.Println(v2)
        _ = fi16
    }
    v2 = (float32(((((1e10 + 3.14159) + 0.0) - float64((1.5 / (1.5 - 3.14159)))) * (0.0 / 1.5))) + v2)
    switch -146 {
    case 1:
        var v17 int = -2147483648
        _ = v17
    case 0:
        v2 = float32((1.5 - -2.25))
    case 127:
        if false {
            fmt.Println(v2, v1)
        } else {
            var v18 float64 = -2.25
            var v19 int16 = -32768
            _ = v18
            _ = v19
        }
    default:
        fmt.Println(v1, v1)
    }
    if false {
        var v20 float32 = v2
        for fi21 := 0; fi21 < 5; fi21++ {
            v3 = append(v3, 0)
            if true {
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := v4[42]
                    v1 += -82
                    _ = fi22
                    _ = v23
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi21
        }
        _ = v20
    } else {
        delete(v5, 4294967295)
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint32 = 2147483647
            v25 = v25
            _ = fi24
            _ = v25
        }
    }
    v2 += 0.0
    var v26 int64 = 0
    v27 := v5[220]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
