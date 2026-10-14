package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 bool = (!(-1 >= -204))
    var v4 uint32 = 292
    fmt.Println(v3, v3)
    for fi5 := 0; fi5 < 7; fi5++ {
        switch 9223372036854775807 {
        case 0:
            if v3 {
                v3 = (!v3)
            }
        case 9223372036854775807:
            if v3 {
                for fi6 := 0; fi6 < 12; fi6++ {
                    v3 = v3
                    var v7 uint64 = 1
                    if v3 {
                        v3 = (!v3)
                        fmt.Println(v4, v4)
                        v8 := []int{fi5, (fi5 + fi5)}
                        _ = v8
                    } else {
                        var v9 uint32 = (^(v4))
                        v9 -= (v4 ^ (v4 & (v9 & v4)))
                        var v10 string = ("" + ("foo bar" + "a"))
                        var v11 int64 = 0
                        var v12 uint32 = v9
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                    v3 = (2147483647 == -1)
                    _ = fi6
                    _ = v7
                }
                v4 -= v4
            }
        case 42:
            v4 += (^(v4))
        default:
            v3 = (65535 <= 42)
        }
        var v13 int16 = 1000
        _ = fi5
        _ = v13
    }
    fmt.Println(v3, v3)
    v3 = v3
    for fi14 := 0; fi14 < 10; fi14++ {
        if v3 {
            v4 = 2147483648
        }
        _ = fi14
    }
    if (!((0.0 / 0.1) <= 3.14159)) {
        var v15 float32 = (float32(1.5) + (3.14159 + (3.14159 + (1e10 * 0.1))))
        v2[(^(v4))] = 4294967295
        _ = v15
    }
    v16 := GenS1{F0: 42}
    if (!v3) {
        var v17 int64 = -9223372036854775808
        var v18 int16 = -32768
        v17 = (v17 * int64(v18))
        _ = v17
        _ = v18
    } else {
        v19 := []int{}
        for fi20 := 0; fi20 < 2; fi20++ {
            v19 = append(v19, (fi20 - (fi20 ^ -9223372036854775808)))
            var v21 bool = (65535 >= 0)
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    var v22 int = 9223372036854775807
    v23 := v2[2]
    var v24 int = -1
    v25 := v2[761]
    var v26 string = ("hello" + "a")
    if false {
        var v27 bool = v3
        v28 := GenS2{F0: 65535, F1: 1}
        v4 += v4
        _ = v27
        _ = v28
    } else {
        var v29 int16 = 0
        delete(v2, v4)
        if (!false) {
            for fi30 := 0; fi30 < 7; fi30++ {
                v22 = (fi30 + 1)
                _ = fi30
            }
        } else {
            v31 := GenS3{F0: -1, F1: -805}
            fmt.Println(v29)
            _ = v31
        }
        _ = v29
    }
    if (!(42 == 42)) {
        if true {
            for fi32 := 0; fi32 < 12; fi32++ {
                if v3 {
                    var v33 uint16 = 65535
                    _ = v33
                }
                _ = fi32
            }
        }
    }
    var v34 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}
