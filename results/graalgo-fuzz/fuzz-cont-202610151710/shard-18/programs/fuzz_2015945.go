package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint }

func main() {
    var v1 int32 = 863
    v2 := make(map[int]uint32)
    var v3 int8 = -3
    var v4 uint8 = 0
    var v5 int64 = -1
    var v6 string = "a"
    if true {
        var v7 uint = 2
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi8, v7, v3)
            _ = fi8
        }
        _ = v7
    } else {
        v4 -= 2
    }
    v6 = v6
    var v9 string = v6
    v4 += (v4 << ((180) & 63))
    var v10 string = v9
    v11 := v2[1]
    v1 -= (v1 << ((18446744073709551615) & 63))
    var v12 uint16 = 0
    var v13 int32 = 622
    v14 := []int{32767}
    v4 += v4
    switch 266 {
    case -9223372036854775808:
        var v15 int8 = v3
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: (v3 - -71), F2: (true && false)}
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: 18446744073709551615, F1: 25, F2: false}
        var v19 uint32 = (v11 >> ((42) & 63))
        switch (fi17 * fi17) {
        case 537:
            var v20 bool = ((0.0 == 0.1) || false)
            _ = v20
        case 9223372036854775807:
            if false {
                var v21 bool = false
                v3 -= 100
                v5 = v5
                _ = v21
            }
        case -819:
            var v22 float64 = (1.5 / -2.25)
            _ = v22
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v23 int64 = 1
    v24 := GenS2{F0: 43}
    var v25 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
