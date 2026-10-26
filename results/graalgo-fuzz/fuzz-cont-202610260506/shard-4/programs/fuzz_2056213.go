package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18
    v3 := make(map[uint32]string)
    var v4 int8 = 116
    var v5 string = "\n\t"
    fmt.Println(v2, v5)
    var v6 int16 = 0
    v2 = (v2 ^ (v2 >> ((4294967295) & 63)))
    var v7 uint64 = 0
    fmt.Println(v4, v5, v5)
    var v8 float32 = (((0.0 * 0.0) + (((3.14159 * -2.25) - (-2.25 - 3.14159)) - -2.25)) + (float32((0.1 * 3.14159)) + 0.0))
    v4 -= v4
    var v9 string = "\n\t"
    v10 := GenS1{F0: (v6 / (((v6 << ((4294967295) & 63))) | 1)), F1: ("hello" < ("a" + ((((v5 + (v5 + v9)) + v9) + ((((v9 + v5) + v9) + ("hello" + v9)) + v5)) + (v9 + (v9 + v9))))), F2: (v7 ^ 0)}
    if false {
        var v11 int64 = -9223372036854775808
        for fi12 := 0; fi12 < 18; fi12++ {
            v7 = (v7 | v7)
            v13 := v1[65535]
            switch ((v5 + v5) + v5) {
            case "\n\t":
                var v14 float64 = (1e10 / -2.25)
                _ = v14
            case "a":
                var v15 float32 = 1.5
                _ = v15
            default:
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v11 -= int64(fi12)
                        _ = fi16
                    }
                }
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v1[65535] = 4294967295
    var v17 int16 = v6
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            var v19 float32 = v8
            v20 := []int{32768, fi18}
            _ = v19
            _ = v20
        }
        v7 = v2
        v21 := GenS1{F0: 909, F1: (!(2 <= uint32(v2))), F2: 0}
        _ = fi18
        _ = v21
    }
    var v22 float64 = ((0.0 + (-2.25 / 1e10)) / 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
