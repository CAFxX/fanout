package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) int {
    v6 := GenS1{F0: 0, F1: 67, F2: 157}
    v7 := []int{127, -974, -9223372036854775808, -129}
    var v8 uint32 = 2147483648
    var v9 uint = 65535
    v8 = 4294967295
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() string {
    var v10 uint8 = 220
    v11 := GenS2{F0: 1, F1: -9223372036854775808}
    if false {
        v10 = v10
    }
    _ = v10
    _ = v11
    return ("hello" + "")
}

func main() {
    var v1 int64 = -303
    v2 := []int{-1, 128, 1, -9223372036854775808}
    var v3 string = ("hello" + "\n\t")
    var v4 string = v3
    var v5 uint = 255
    switch 127 {
    case 2147483647:
        var v12 float64 = (float64(3.14159) / 1.5)
        _ = v12
    case 9223372036854775807:
        var v13 int8 = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (((0.0 + (1e10 / (1e10 / 1.0))) * 0.1) != float32(3.14159)) {
            switch (-(fi14)) {
            case 310:
                var v15 uint = v5
                _ = v15
            case 2147483647:
                var v16 int64 = 179
                _ = v16
            }
        }
        _ = fi14
    }
    v3 = v4
    var v17 int64 = -9223372036854775808
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS2{F0: 255, F1: v17}
        v2 = append(v2, fi18)
        _ = fi18
        _ = v19
    }
    v1 += 1
    v17 -= (v17 % ((-411) | 1))
    switch v3 {
    case "foo bar":
        var v20 int = -1
        _ = v20
    case "\n\t":
        var v21 uint16 = 32767
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v5 = v5
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 uint8 = 255
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    if (2 <= 1) {
        v2 = append(v2, 127)
    } else {
        v25 := GenS3{F0: (-306 >= -1), F1: 170, F2: 255}
        var v26 uint64 = 4294967295
        _ = v25
        _ = v26
    }
    _ = len(v2)
    v2 = append(v2, 32767)
    v2 = append(v2, -972)
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            fmt.Println(fi28, v17)
            _ = fi28
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi27
    }
    var v29 int32 = 0
    fmt.Println(v3, v29)
    v3 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
