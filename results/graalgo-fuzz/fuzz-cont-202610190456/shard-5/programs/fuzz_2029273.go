package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    v4 := make(map[string]uint32)
    var v5 uint = 2
    v5 -= v5
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 = 18446744073709551615
        v7 := GenS1{F0: 4294967295}
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint32 = 699
            v10 := GenS2{F0: false}
            v11 := GenS3{F0: 2}
            var v12 uint16 = 255
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v3 = (v3 % ((-1) | 1))
    v4[""] = 0
    delete(v4, ((("" + "hello") + "a") + "hello"))
    v13 := GenS4{F0: 2147483648, F1: 127, F2: 982}
    v4[""] = 852
    v5 += v5
    for fi14 := 0; fi14 < 15; fi14++ {
        v5 += v5
        v5 = v5
        if (42 > 65535) {
            switch "foo bar" {
            case "a":
                var v15 int32 = -1
                _ = v15
            case "":
                for fi16 := 0; fi16 < 11; fi16++ {
                    v5 -= (v5 + (v5 * (v5 << ((v5) & 63))))
                    _ = fi16
                }
            default:
                if false {
                    var v17 int8 = 127
                    fmt.Println(v3)
                    _ = v17
                }
            }
            var v18 float64 = 1.5
            fmt.Println(v5, v18)
            _ = v18
        }
        _ = fi14
    }
    v19 := GenS4{F0: 401, F1: 127, F2: 262}
    for fi20 := 0; fi20 < 8; fi20++ {
        v5 = 2
        var v21 uint8 = 128
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
