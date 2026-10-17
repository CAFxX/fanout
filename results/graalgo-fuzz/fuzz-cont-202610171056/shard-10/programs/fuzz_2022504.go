package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 2147483647
    var v3 uint16 = v1
    fmt.Println(v2)
    v4 := GenS1{F0: 2, F1: 4294967295}
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{(fi5 >> ((1) & 63)), (fi5 >> ((1) & 63))}
        v3 -= v3
        _ = fi5
        _ = v6
    }
    if (0 < 0) {
        v2 = -2147483648
        v7 := []int{128, 1}
        v1 -= (v1 << ((2) & 63))
        _ = v7
    } else {
        v2 = v2
        var v8 uint = 2
        _ = v8
    }
    fmt.Println(v2, v2, v1)
    var v9 bool = false
    v10 := []int{755, 1, 32768, 1}
    var v11 float32 = 3.14159
    for fi12 := 0; fi12 < 12; fi12++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if ((v2 * (v2 - v2)) == (v2 % ((v2) | 1))) {
            v13 := GenS1{F0: 127, F1: 170}
            _ = v13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v9)
        v10 = append(v10, fi14)
        v15 := GenS2{F0: (v2 + v2), F1: 18446744073709551615}
        _ = fi14
        _ = v15
    }
    fmt.Println(v9)
    v9 = v9
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    v16 := GenS3{F0: (v3 << ((255) & 63)), F1: v2}
    var v17 uint64 = 550
    fmt.Println(v17, v9)
    if (!(v2 < v2)) {
        for fi18 := 0; fi18 < 1; fi18++ {
            v10 = append(v10, (fi18 | fi18))
            v11 += 1e10
            v1 -= (v3 << ((255) & 63))
            _ = fi18
        }
    } else {
        if len(v10) > 0 {
            v10[0] = 128
        }
    }
    if false {
        fmt.Println(v1, v3, v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
