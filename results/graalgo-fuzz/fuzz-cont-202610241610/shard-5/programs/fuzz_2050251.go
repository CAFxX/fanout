package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint16 }

func main() {
    var v1 int = 1
    var v2 uint32 = 2147483648
    var v3 int16 = -377
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{fi5, 127, (v1 % ((v1) | 1)), (fi5 * int(v2))}
        _ = len(v6)
        _ = fi5
        _ = v6
    }
    v4 = "a"
    var v7 uint = 255
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: uint32(v7), F1: 91}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 float64 = 1.5
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: (v2 << ((v7) & 63)), F1: 25}
        _ = fi12
        _ = v13
    }
    var v14 int = (v1 | v1)
    v15 := GenS2{F0: 1}
    if (0 > 573) {
        v1 = (v1 ^ (v14 << (((v7 / ((v7) | 1))) & 63)))
        for fi16 := 0; fi16 < 2; fi16++ {
            v1 = (-(v14))
            _ = fi16
        }
    } else {
        v17 := []int{}
        _ = v17
    }
    if false {
        if (1.5 > ((0.0 - 1e10) * (1.5 - -2.25))) {
            fmt.Println(v7)
            var v18 uint8 = 0
            var v19 uint64 = 0
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: (v2 + v2), F1: 2}
    v7 -= v7
    v21 := GenS3{F0: 65535, F1: 127, F2: 2}
    var v22 int = v14
    if (1 != 100) {
        v1 = -205
        v23 := GenS3{F0: 255, F1: 0, F2: 65535}
        _ = v23
    } else {
        v24 := GenS1{F0: (v2 & 42), F1: 4294967295}
        v2 = v2
        v25 := []int{}
        _ = v24
        _ = v25
    }
    var v26 int8 = -1
    var v27 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
