package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }

func main() {
    var v1 bool = true
    var v2 int64 = 4
    var v3 uint = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    if v1 {
        var v6 bool = false
        if v6 {
            var v7 int64 = 1
            _ = v7
        }
        switch 705 {
        case 127:
            v8 := GenS1{F0: 113, F1: -1, F2: 11}
            _ = v8
        case -1:
            v6 = v6
        case 4294967295:
            v6 = (-44 != 44)
        default:
            for fi9 := 0; fi9 < 18; fi9++ {
                v2 = v2
                v3 = (v3 << ((v3) & 63))
                if ("hello" == "") {
                    v2 -= (v2 / ((v2) | 1))
                }
                _ = fi9
            }
        }
        _ = v6
    } else {
        v10 := GenS1{F0: 42, F1: -1, F2: 78}
        if (9223372036854775807 != -2147483648) {
            v1 = v1
            v2 -= 1
        } else {
            if ((929 <= 2) || v1) {
                v2 += v2
                var v11 string = ""
                _ = v11
            }
        }
        switch "" {
        case "foo bar":
            v1 = (42 < 42)
        case "\n\t":
            v1 = (1.5 > 0.0)
        case "hello":
            var v12 uint64 = 0
            _ = v12
        default:
            v3 += v3
        }
        _ = v10
    }
    v1 = true
    if v1 {
        var v13 float32 = 1e10
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            delete(v5, fi14)
            _ = fi14
        }
        var v15 uint32 = 42
        switch -244 {
        case 0:
            switch "a" {
            case "\n\t":
                var v16 uint16 = uint16(v3)
                _ = v16
            }
        case 412:
            if false {
                v17 := v4[(v15 & 4294967295)]
                _ = v17
            }
        }
        _ = v15
    }
    switch 128 {
    case 1:
        if v1 {
            delete(v4, 2147483648)
            var v18 float32 = (1.5 - (1e10 + (-2.25 - 3.14159)))
            _ = v18
        } else {
            var v19 uint8 = 128
            v2 = int64(v19)
            _ = len(v4)
            _ = v19
        }
    case 51:
        fmt.Println(v3, v1)
    default:
        var v20 uint32 = 65535
        _ = v20
    }
    v21 := GenS1{F0: uint16(v3), F1: 2, F2: 127}
    v2 -= 173
    switch -2147483648 {
    case 42:
        if (1 <= 4294967295) {
            for fi22 := 0; fi22 < 13; fi22++ {
                delete(v4, 2)
                _ = fi22
            }
            fmt.Println(v2)
            v23 := GenS1{F0: 65535, F1: 0, F2: 197}
            _ = v23
        }
    case 1:
        if v1 {
            var v24 float64 = (float64((((1e10 / (1.5 + 1.5)) + -2.25) - 0.0)) / 1.5)
            v25 := []int{9223372036854775807, 128, -280}
            _ = v24
            _ = v25
        }
    case 20:
        var v26 int8 = 42
        _ = v26
    default:
        if (-944 < 646) {
            v2 = v2
            v2 -= 754
        } else {
            var v27 uint16 = 42
            var v28 uint8 = 127
            _ = v27
            _ = v28
        }
    }
    switch 0 {
    case 333:
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint64 = 0
            fmt.Println(v30)
            v5[-698] = ("" + ("" + "a"))
            _ = fi29
            _ = v30
        }
    default:
        v3 -= (v3 % (((v3 << (((v3 << ((231) & 63))) & 63))) | 1))
    }
    var v31 int32 = 1
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := []int{fi32, -320}
        _ = fi32
        _ = v33
    }
    var v34 uint8 = 0
    var v35 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
