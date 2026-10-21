package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint) int {
    var v4 int16 = 459
    fmt.Println(p0)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch fi5 {
        case -2147483648:
            var v6 uint8 = 243
            _ = v6
        case 32767:
            v4 = v4
        case 32768:
            p0 += (p0 << (((p0 / ((1) | 1))) & 63))
        default:
            if false {
                var v7 uint64 = 4294967295
                v4 = 42
                var v8 bool = (1e10 >= 1e10)
                _ = v7
                _ = v8
            }
        }
        var v9 uint64 = 932
        _ = fi5
        _ = v9
    }
    fmt.Println(p0, v4)
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 bool = (-9223372036854775808 <= 0)
    var v2 uint = 65535
    v3 := []uint32{992}
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 int64 = 9223372036854775807
    for fi11 := 0; fi11 < 13; fi11++ {
        v1 = true
        _ = fi11
    }
    v1 = false
    var v12 float64 = 3.14159
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int8 = 1
        var v15 int = int(v10)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint16 = 32767
    var v17 int32 = 0
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: v1}
        _ = fi18
        _ = v19
    }
    var v20 uint = v2
    switch 128 {
    case 32767:
        for fi21 := 0; fi21 < 6; fi21++ {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi21
        }
    default:
        v17 = (-(v17))
    }
    v10 = (v10 / (((v10 % ((-9223372036854775808) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
