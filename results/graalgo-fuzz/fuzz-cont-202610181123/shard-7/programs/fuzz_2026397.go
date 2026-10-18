package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 int) string {
    fmt.Println(p0)
    v4 := GenS1{F0: 1, F1: -362, F2: 42}
    _ = p0
    _ = v4
    return "\n\t"
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 4294967295
    var v3 int = 2147483647
    v2 += v2
    v5 := GenS1{F0: 42, F1: 0, F2: 0}
    v1 += (v1 + (((v1 / ((0.0 - 1.5) + (v1 * 1e10))) / v1) + 1.5))
    var v6 int8 = 99
    var v7 uint32 = 2147483648
    fmt.Println(v6, v3, v3)
    v3 -= (^(v3))
    fmt.Println(v2, v2)
    v7 -= (v7 ^ (v7 & (^(v7))))
    if ((v6 | (v6 << ((v2) & 63))) <= int8(v2)) {
        v6 = v6
        v8 := GenS1{F0: 500, F1: 0, F2: 331}
        _ = v8
    }
    switch (^(v3)) {
    case 876:
        v9 := GenS1{F0: 65535, F1: 631, F2: v7}
        _ = v9
    case 570:
        var v10 bool = (v6 > (v6 | -65))
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    switch (v3 >> ((v2) & 63)) {
    case 32768:
        fmt.Println(v6, v7)
    case 32767:
        if ((!true) || ((42 <= 222) || (false || (157 == 2)))) {
            var v12 uint = v2
            if ((!((!(!false)) && false)) || false) {
                fmt.Println(v7, v6, v12)
                if (((v6 != int8(v7)) || (-32768 < -32768)) || ((4294967295 == 0) || false)) {
                    v7 += 1
                    var v13 uint = (v2 & v12)
                    v13 = v12
                    _ = v13
                } else {
                    var v14 int64 = 42
                    v12 += (v2 | 921)
                    _ = v14
                }
            }
            v3 = v3
            _ = v12
        }
    case 1:
        v15 := []int{v3, v3, int(v6), 186}
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: -1, F2: v7}
    var v17 bool = (true || false)
    v18 := GenS2{F0: (v3 + -17)}
    v6 = 127
    v1 += (float32((0.1 / 3.14159)) * v1)
    for fi19 := 0; fi19 < 5; fi19++ {
        v1 += v1
        v20 := []int{(fi19 - (v3 >> ((2) & 63)))}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
