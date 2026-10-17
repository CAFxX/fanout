package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint32 = 231
    var v2 int32 = 1
    var v3 int16 = -1
    v4 := []uint32{(^(v1)), v1, (v1 - (^(v1)))}
    var v5 uint64 = 311
    v6 := []int{}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            if ((42 <= -67) || ("foo bar" == "")) {
                v3 = v3
                var v9 float32 = (0.1 + 3.14159)
                if (!(v5 < (v5 % (((v5 - (v5 + 443))) | 1)))) {
                    if (1.5 < 1.5) {
                        var v10 string = ("foo bar" + ("foo bar" + "hello"))
                        _ = v10
                    }
                } else {
                    _ = len(v4)
                }
                v4 = append(v4, v1)
                _ = v9
            }
            v3 += (v3 << ((1) & 63))
            _ = fi8
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(fi11, v1, v2)
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: 65535}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint8 = 42
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi14
            _ = v15
        }
        _ = fi11
    }
    v5 -= 929
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 float32 = 3.14159
        v18 := GenS1{F0: 4294967295}
        _ = fi16
        _ = v17
        _ = v18
    }
    switch -858 {
    case 32767:
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint = 18446744073709551615
            _ = fi19
            _ = v20
        }
    default:
        var v21 int32 = (v2 * (v2 * (v2 * v2)))
        _ = v21
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v22 uint64 = (v5 + (v5 >> ((390) & 63)))
    v3 += (v3 << ((42) & 63))
    _ = len(v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    v5 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
