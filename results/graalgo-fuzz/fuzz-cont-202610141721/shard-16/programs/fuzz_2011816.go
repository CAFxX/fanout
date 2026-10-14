package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16, p1 int32) uint64 {
    var v4 bool = (!(2147483647 < -879))
    p1 = p1
    if false {
        p1 += (p1 & p1)
    } else {
        var v5 bool = false
        v6 := GenS1{F0: p1}
        v7 := GenS2{F0: 1, F1: 856}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: -1, F1: 96}
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 388
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    v3 := make(map[string]uint32)
    v2 = v2
    var v9 uint = 584
    fmt.Println(v2, v9)
    v9 = (v9 - 18446744073709551615)
    v10 := []int{32768, -300, 683}
    v2 = (0.0 >= (((-2.25 * float32(0.0)) - float32(((float64(0.0) - -2.25) - 3.14159))) + 0.1))
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            switch "\n\t" {
            case "a":
                var v12 uint = (v9 - v9)
                _ = v12
            case "":
                v9 = (v9 << ((v9) & 63))
            case "hello":
                if (((v2 && ((float32((float64((float32(3.14159) / 3.14159)) - 0.0)) == float32(0.0)) || (!(2 < 899)))) && (1 <= 2147483647)) || (!(!(-2.25 >= ((((1.5 + float32((3.14159 - 1.5))) + 3.14159) * 1e10) * 0.0))))) {
                    var v13 int = fi11
                    var v14 int16 = 1
                    var v15 uint32 = 561
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v2 = (!(0 <= 1))
                }
            }
            v9 = 255
            var v16 bool = (0 == 42)
            _ = v16
        }
        if (v2 && v2) {
            var v17 uint32 = 42
            if v2 {
                var v18 uint64 = 909
                fmt.Println(v2, v9)
                _ = v18
            }
            _ = v17
        }
        _ = fi11
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v9, v9, fi19)
        if v2 {
            v9 += v9
            v9 = v9
            v2 = true
        }
        var v20 int8 = 66
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
