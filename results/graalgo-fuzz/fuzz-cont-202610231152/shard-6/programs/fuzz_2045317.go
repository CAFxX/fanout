package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    var v7 int32 = -1
    if false {
        var v8 int = 127
        v9 := []int{}
        _ = v8
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{65535, 2}
    var v2 uint64 = 18446744073709551615
    var v3 string = ""
    var v4 bool = (!true)
    v5 := []uint32{2147483648, 42, 0}
    v6 := []int{}
    fmt.Println(v2, v4, v2)
    v10 := GenS1{F0: -32768}
    v4 = (!v4)
    v11 := GenS1{F0: 42}
    v2 = v2
    var v12 uint = 4294967295
    var v13 bool = v4
    switch 2147483647 {
    case 0:
        var v14 float32 = float32(((float64(float32(3.14159)) * -2.25) + ((-2.25 + float64(float32(((((((0.1 * (float64((3.14159 * (1.5 / 1.0))) * float64(0.1))) * 0.1) * float64(3.14159)) * (0.1 + (3.14159 * -2.25))) / 1e10) * 1e10)))) / (0.1 - 1e10))))
        _ = v14
    case 128:
        if (v12 < v12) {
            _ = len(v5)
            var v15 int = -953
            _ = v15
        }
    case -935:
        v2 = (v2 + (v2 / ((v2) | 1)))
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v4 = true
            _ = fi16
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v4 = (4294967295 >= 2)
        _ = fi17
    }
    v2 += 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
