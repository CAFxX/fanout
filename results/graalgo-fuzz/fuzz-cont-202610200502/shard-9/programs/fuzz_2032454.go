package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int32 {
    var v4 int16 = 32767
    v4 += (v4 << ((1) & 63))
    var v5 bool = (!false)
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int64) uint64 {
    v6 := []int{4294967295, -1}
    var v7 int8 = 100
    _ = p0
    _ = v6
    _ = v7
    return 554
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 132
    var v3 int8 = 42
    v2 = v2
    if false {
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 = (v2 - (v2 | (v2 ^ (^(v2)))))
            v9 := []int{fi8, fi8, (-(fi8))}
            var v10 uint8 = uint8(v2)
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint = 2
        v11 = v11
        _ = v11
    } else {
        var v12 uint16 = 255
        fmt.Println(v12)
        fmt.Println(v3)
        _ = v12
    }
    if ((2 < 0) || false) {
        var v13 float32 = ((3.14159 + float32((1e10 + 0.0))) - 0.1)
        _ = v13
    }
    fmt.Println(v3, v3)
    switch -9223372036854775808 {
    case -963:
        if false {
            v2 = v2
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int8 = (^(v3))
                _ = fi14
                _ = v15
            }
        }
    }
    v1[("foo bar" + "foo bar")] = 619
    if (-1 >= 0) {
        fmt.Println(v3, v3)
        switch "hello" {
        case "foo bar":
            _ = len(v1)
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            v3 = (v3 ^ v3)
            var v17 bool = true
            _ = fi16
            _ = v17
        }
    } else {
        var v18 float64 = 3.14159
        v2 = (^(v2))
        _ = v18
    }
    fmt.Println(v2)
    v2 = 115
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: 4294967295}
            v2 = (v2 ^ 1)
            _ = fi20
            _ = v21
        }
        var v22 uint64 = (v2 >> ((18446744073709551615) & 63))
        v2 = 0
        _ = fi19
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
