package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
func genFn0() int {
    var v6 uint8 = 161
    v6 += 255
    _ = v6
    return -2147483648
}

func main() {
    var v1 float64 = ((0.0 - 1e10) + 3.14159)
    var v2 uint = 65535
    var v3 uint16 = 1
    var v4 bool = (0 < -777)
    var v5 uint16 = (v3 & 255)
    if (128 != 127) {
        var v7 uint8 = 226
        var v8 float32 = (((1.5 / (0.0 + (0.1 + float32(v1)))) - 3.14159) + float32((-2.25 + float64(3.14159))))
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: 18446744073709551615}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v5 -= uint16(v2)
        v13 := []int{fi12}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v5 -= (v5 * 799)
        _ = fi14
    }
    var v15 uint16 = 32767
    for fi16 := 0; fi16 < 2; fi16++ {
        v4 = (!true)
        switch fi16 {
        case 2147483647:
            fmt.Println(v3)
        }
        fmt.Println(v3, v1, v3)
        _ = fi16
    }
    var v17 float64 = float64((0.0 - (float32(v1) - (float32(float64(1.5)) / 3.14159))))
    var v18 uint8 = 27
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
