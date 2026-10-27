package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
func genFn0(p0 uint16) int32 {
    switch 128 {
    case 4294967295:
        var v5 int8 = 0
        _ = v5
    case 32767:
        v6 := []int{127}
        _ = v6
    }
    var v7 bool = (!(float64((1e10 + 0.0)) <= 3.14159))
    _ = p0
    _ = v7
    return -760
}
func genFn1(p0 bool) bool {
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 177}
        _ = fi8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{1}
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float32 = 0.1
                v14 := GenS2{F0: -123, F1: 42, F2: (fi10 << ((65535) & 63))}
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 9; fi15++ {
                v11 = append(v11, fi10)
                p0 = ((0 >= 2) || ((42 != 964) && (((float32((((float64(0.1) - (1e10 + (float64((3.14159 * ((float32(((0.1 * 3.14159) - float64((3.14159 - ((float32(0.1) - 1e10) + 1.5))))) - (1.5 / 1e10)) - 0.0))) + 1.5))) + 0.1) * 1e10)) + 3.14159) >= (0.0 - ((0.1 * 0.0) / -2.25))) || ((!p0) || false))))
                var v16 float32 = 1.5
                _ = fi15
                _ = v16
            }
            _ = fi10
            _ = v11
        }
        p0 = p0
        p0 = (p0 && (4294967295 >= 2147483647))
    }
    _ = p0
    return true
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v17 float64 = float64(1e10)
    var v18 uint = 819
    v19 := []int{}
    fmt.Println(v17)
    switch -64 {
    case -1:
        v17 = (v17 - v17)
    default:
        if v1 {
            var v20 uint8 = 177
            _ = len(v19)
            _ = v20
        }
    }
    fmt.Println(v18, v17)
    switch "foo bar" {
    case "":
        v21 := GenS2{F0: -117, F1: 0, F2: 4294967295}
        _ = v21
    default:
        v22 := GenS2{F0: -112, F1: genFn0(334), F2: 128}
        _ = v22
    }
    if (42 != 0) {
        if v1 {
            if len(v19) > 0 {
                v19[0] = 789
            }
        } else {
            fmt.Println(v17, v1)
        }
        var v23 uint16 = 255
        _ = v23
    }
    var v24 bool = (-9223372036854775808 >= 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
