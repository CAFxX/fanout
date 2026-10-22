package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
func genFn0(p0 int16) uint {
    var v4 string = ""
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 1, F1: ("hello" + (v4 + (("\n\t" + ("\n\t" + v4)) + ("hello" + (("a" + v4) + v4))))), F2: (fi5 >> ((65535) & 63))}
        for fi7 := 0; fi7 < 15; fi7++ {
            if true {
                v8 := []int{fi5, fi7}
                v9 := GenS1{F0: 9223372036854775807, F1: v4, F2: (fi7 << ((85) & 63))}
                _ = v8
                _ = v9
            } else {
                var v10 uint8 = 102
                _ = v10
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 128
    var v2 int = 32768
    var v3 int32 = 118
    var v11 uint8 = v1
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 1, F1: ("foo bar" + (("\n\t" + "hello") + "hello")), F2: 4294967295}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int64 = 848
        _ = fi14
        _ = v15
    }
    var v16 float64 = 3.14159
    v17 := []int{v2, (v2 >> ((genFn0(1)) & 63))}
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int16 = -457
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int8 = 62
            if false {
                v17 = append(v17, 4294967295)
            }
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    if false {
        var v22 uint8 = (v1 << ((uint(v1)) & 63))
        fmt.Println(v2)
        _ = v22
    }
    v17 = append(v17, int(v3))
    v16 = v16
    v23 := GenS1{F0: -720, F1: ("foo bar" + ""), F2: v2}
    if len(v17) > 0 {
        v17[0] = v2
    }
    var v24 int8 = int8(v3)
    var v25 bool = (float32((-2.25 + 1.5)) >= (3.14159 + ((0.0 * ((((float32(-2.25) / 1.0) / -2.25) / 3.14159) + (0.1 - -2.25))) + (1e10 - 0.0))))
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(v25)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
