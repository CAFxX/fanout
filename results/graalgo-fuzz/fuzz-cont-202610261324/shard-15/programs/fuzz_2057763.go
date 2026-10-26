package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8) int64 {
    p0 -= (p0 << ((1) & 63))
    p0 += p0
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", "hello"}
    var v3 float32 = (0.0 - (float32(float64((3.14159 - 1e10))) * 1e10))
    delete(v1, 2147483648)
    v3 = v3
    switch 614 {
    case 32768:
        if true {
            var v4 string = "foo bar"
            var v5 uint16 = 54
            fmt.Println(v4, v3)
            _ = v4
            _ = v5
        } else {
            var v6 int32 = 2147483647
            var v7 int16 = 724
            switch -671 {
            case -9223372036854775808:
                v6 += (-(v6))
            default:
                v8 := GenS1{F0: 42, F1: 677, F2: 65535}
                _ = v8
            }
            _ = v6
            _ = v7
        }
    case 127:
        v3 = ((v3 * v3) / v3)
    case 2147483647:
        v3 = float32((3.14159 + (3.14159 - float64(v3))))
    }
    if (!false) {
        var v9 int32 = -2147483648
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v1[42] = ("\n\t" + "hello")
            _ = fi10
        }
        v3 = 0.0
        var v11 bool = true
        _ = v11
    }
    v12 := GenS2{F0: 1000, F1: 42, F2: -1}
    var v13 uint8 = 255
    v3 += (v3 - v3)
    for fi14 := 0; fi14 < 16; fi14++ {
        switch "a" {
        case "\n\t":
            v15 := GenS3{F0: 497}
            _ = v15
        case "hello":
            _ = len(v1)
        case "a":
            var v16 int32 = 0
            _ = v16
        }
        v17 := v1[2]
        _ = fi14
        _ = v17
    }
    switch "hello" {
    case "foo bar":
        switch "\n\t" {
        case "":
            v13 += 177
        case "hello":
            if false {
                v3 = v3
                v3 = v3
            }
        default:
            var v18 float32 = 0.0
            _ = v18
        }
    default:
        var v19 string = ("" + "\n\t")
        _ = v19
    }
    fmt.Println(v3)
    switch 1 {
    case 32768:
        fmt.Println(v13, v3)
    }
    var v20 bool = (-103 == 42)
    v3 -= v3
    var v21 uint16 = 1
    v22 := v1[1]
    if (0 >= 0) {
        v21 = (v21 | 1)
    }
    switch v22 {
    case "foo bar":
        v23 := GenS2{F0: -1, F1: 476, F2: -1}
        _ = v23
    case "":
        var v24 uint8 = (v13 << ((4294967295) & 63))
        _ = v24
    case "a":
        var v25 int = 128
        _ = v25
    default:
        if v20 {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int = (fi26 >> ((206) & 63))
                for fi28 := 0; fi28 < 2; fi28++ {
                    v29 := GenS1{F0: 4294967295, F1: 32767, F2: uint(v21)}
                    _ = fi28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                _ = len(v2)
                _ = fi30
            }
            v3 = v3
            var v31 int32 = 1
            _ = v31
        }
    }
    v13 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
