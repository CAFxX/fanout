package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = -1
        var v7 int = fi5
        var v8 int16 = 0
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if false {
        var v9 int8 = 65
        fmt.Println(v9)
        _ = v9
    } else {
        var v10 int32 = 0
        _ = v10
    }
    if (!true) {
        var v11 uint64 = 4294967295
        v11 += (v11 | 18446744073709551615)
        _ = v11
    }
    return 42
}

func main() {
    v1 := []uint32{120, 230, 819, 567}
    var v2 int8 = 87
    var v3 int = 523
    var v4 string = "a"
    if true {
        fmt.Println(v3, v3, v2)
        if (v3 < 516) {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: 163}
                for fi14 := 0; fi14 < 9; fi14++ {
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    fmt.Println(v3, v4)
                    _ = len(v1)
                    if false {
                        var v15 uint16 = 32767
                        var v16 float64 = 1.5
                        fmt.Println(fi12)
                        v17 := GenS1{F0: 227}
                        var v18 bool = (!(float32(v16) >= 0.1))
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    if (3.14159 <= (float64((1e10 * 0.1)) - ((((0.1 + float64(1e10)) / 0.1) / (3.14159 / 3.14159)) / 0.1))) {
                        v19 := GenS2{F0: true, F1: 2}
                        _ = v19
                    }
                    _ = fi14
                }
                _ = fi12
                _ = v13
            }
        }
    } else {
        v3 = (v3 << ((604) & 63))
        var v20 uint32 = 42
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: 31}
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v23 bool = (false || (-1 != 0))
    for fi24 := 0; fi24 < 20; fi24++ {
        v4 = v4
        if v23 {
            v25 := GenS1{F0: 0}
            var v26 float32 = (1.5 + ((1.5 * -2.25) + (0.1 * -2.25)))
            _ = v25
            _ = v26
        } else {
            var v27 uint16 = 0
            v28 := GenS1{F0: 0}
            v2 -= v2
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v23 = (-1 >= -32768)
        _ = fi29
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v23 = v23
    var v30 uint8 = 1
    v23 = v23
    var v31 uint32 = 621
    var v32 string = "\n\t"
    v31 += 42
    var v33 int8 = (-(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
