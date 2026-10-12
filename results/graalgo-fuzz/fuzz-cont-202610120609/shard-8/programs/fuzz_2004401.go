package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string, p1 uint8) uint {
    v4 := GenS1{F0: -1}
    var v5 uint16 = 42
    v5 -= (v5 | v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 809
}
func genFn1() uint {
    var v6 uint16 = 32767
    v6 += v6
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v6, v6)
        _ = fi7
    }
    var v8 int8 = 1
    v9 := []int{-1}
    _ = v6
    _ = v8
    _ = v9
    return 510
}

func main() {
    var v1 int64 = -533
    var v2 int32 = -848
    var v3 float32 = (1e10 + ((((1e10 - float32(((1e10 - 0.1) / 0.1))) * float32((((1e10 - 3.14159) * 1e10) + 0.0))) / 0.1) - -2.25))
    var v10 int64 = v1
    for fi11 := 0; fi11 < 14; fi11++ {
        v10 = -455
        v12 := GenS1{F0: (v2 | (v2 >> ((1) & 63)))}
        fmt.Println(v10)
        _ = fi11
        _ = v12
    }
    var v13 string = "foo bar"
    if (true && true) {
        v10 = int64(v2)
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v1, v3)
        var v15 uint8 = 127
        fmt.Println(fi14, v3)
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v1, v2)
    switch v13 {
    case "hello":
        v1 -= v10
    case "a":
        var v16 int8 = -1
        _ = v16
    }
    if (0.0 <= 1e10) {
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := []int{fi17, fi17, 32767, (fi17 - -328)}
            _ = fi17
            _ = v18
        }
        fmt.Println(v2, v10, v1)
        v19 := []int{806, 0, 0, 4294967295}
        _ = v19
    } else {
        if (1 != 678) {
            v20 := []int{}
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(v10, v10)
                var v22 int16 = -32768
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 uint8 = 2
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: (v2 | (v2 ^ 424))}
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
}
