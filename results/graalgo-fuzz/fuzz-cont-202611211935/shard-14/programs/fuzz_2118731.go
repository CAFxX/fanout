package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) int16 {
    v4 := GenS1{F0: 255}
    var v5 uint16 = 1
    switch -9223372036854775808 {
    case -2147483648:
        p0 = (p0 >> ((126) & 63))
    default:
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := []int{764, (^(fi6)), fi6, 1}
            v7 = append(v7, (-(fi6)))
            _ = fi6
            _ = v7
        }
    }
    switch "foo bar" {
    case "":
        var v8 uint64 = 4294967295
        _ = v8
    case "\n\t":
        var v9 float32 = (float32(float64(((0.1 + (1.5 / 1.5)) + 1e10))) / ((-2.25 - 0.1) + (-2.25 - -2.25)))
        _ = v9
    default:
        v10 := GenS1{F0: (v5 & 496)}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            p0 += 530
        }
        _ = fi11
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int8) bool {
    fmt.Println(p0)
    p0 = -1
    var v12 uint32 = 2147483648
    if true {
        var v13 uint8 = 255
        if (genFn0(0) != genFn0(383)) {
            var v14 uint = 65535
            _ = v14
        }
        _ = v13
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float64 = -2.25
        v17 := []int{(fi15 >> ((966) & 63)), (fi15 >> ((809) & 63)), fi15}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v12
    return (!(128 > 437))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 uint16 = 2
    v3 = (v3 + v3)
    fmt.Println(v3, v3)
    var v18 uint16 = (v3 ^ v3)
    switch "hello" {
    case "\n\t":
        var v19 uint64 = 0
        _ = v19
    case "hello":
        var v20 uint32 = 4294967295
        _ = v20
    case "":
        if false {
            var v21 uint64 = 125
            _ = len(v1)
            _ = v21
        }
    }
    var v22 int32 = 2147483647
    switch -2147483648 {
    case -2147483648:
        switch 42 {
        case -2147483648:
            if (1 == 42) {
                v3 = 718
            }
        case -1:
            v2 = v22
        default:
            v23 := GenS2{F0: 0}
            _ = v23
        }
    case 81:
        var v24 uint16 = 1
        _ = v24
    case 2147483647:
        var v25 int8 = 1
        _ = v25
    }
    fmt.Println(v2, v18, v18)
    var v26 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
}
