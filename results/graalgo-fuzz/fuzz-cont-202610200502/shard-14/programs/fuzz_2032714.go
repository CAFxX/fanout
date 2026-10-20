package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int, p1 int) uint8 {
    p0 = p1
    v4 := []int{}
    var v5 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1() uint {
    var v6 int16 = 1
    v6 = (^(v6))
    v7 := []int{-2147483648, 9223372036854775807, 54}
    _ = v6
    _ = v7
    return 521
}

func main() {
    var v1 int64 = 1
    var v2 float32 = 1e10
    var v3 uint32 = 350
    switch 128 {
    case -665:
        for fi8 := 0; fi8 < 14; fi8++ {
            v1 = v1
            var v9 bool = false
            if (42 < 18446744073709551615) {
                v2 = 3.14159
                v1 = (v1 * 677)
            }
            _ = fi8
            _ = v9
        }
    case 736:
        v10 := []int{42, -900}
        _ = v10
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    var v12 float32 = v2
    var v13 string = "\n\t"
    fmt.Println(v1, v13, v12)
    v14 := GenS1{F0: 4294967295}
    var v15 int64 = v1
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            if ((-(fi16)) >= 32768) {
                v17 := []int{fi16, (fi16 + fi16), fi16}
                _ = v17
            } else {
                v1 = (v1 ^ (^(v1)))
            }
        }
        var v18 uint32 = (v3 - (v3 / (((v3 + (v3 & (v3 * (^(v3)))))) | 1)))
        _ = fi16
        _ = v18
    }
    fmt.Println(v12)
    v2 += 0.1
    var v19 float64 = (((((((((1.5 + 0.0) * ((float64(v12) / 0.1) - 1.5)) - ((1e10 + 3.14159) / 1.5)) - (float64(v2) - 1.5)) - -2.25) - ((0.1 + 0.0) * (0.0 / 1e10))) - (1.5 / (float64(float32((((((1e10 * (0.0 / 1.5)) * -2.25) - 0.1) * float64((-2.25 - float32(((0.1 + (0.1 / float64(v12))) + 0.0))))) + 1.5))) + 1e10))) / 1.0) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}
