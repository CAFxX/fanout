package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0() int32 {
    var v5 int32 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint16 = 255
        v8 := []int{fi6}
        v5 += 42
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = -2.25
    _ = v5
    _ = v9
    return (-(v5))
}
func genFn1() int64 {
    var v10 int = -9223372036854775808
    if ((true && true) || false) {
        var v11 bool = (!((!false) && ((32767 <= -32768) && false)))
        _ = v11
    }
    v12 := []int{v10, (v10 << ((18446744073709551615) & 63))}
    switch "\n\t" {
    case "a":
        var v13 float32 = 1.5
        _ = v13
    case "foo bar":
        v12 = append(v12, (v10 % ((v10) | 1)))
    case "\n\t":
        switch v10 {
        case -909:
            for fi14 := 0; fi14 < 10; fi14++ {
                v10 = v10
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi14
            }
        default:
            var v15 float32 = -2.25
            _ = v15
        }
    default:
        v10 += (v10 + 854)
    }
    v10 = v10
    _ = v10
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{("\n\t" + ("hello" + ("hello" + ""))), (("\n\t" + "foo bar") + ((("hello" + "\n\t") + ("hello" + "hello")) + (("" + "a") + ("" + "")))), "hello", "foo bar"}
    v3 := make(map[uint32]uint32)
    var v4 int16 = 42
    if true {
        if (42 == 65535) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        v16 := GenS1{F0: false}
        _ = v16
    }
    fmt.Println(v4, v1)
    v1 = 578
    var v17 int8 = 100
    v1 -= 2147483647
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS2{F0: 94, F1: 2147483647, F2: false}
        _ = fi18
        _ = v19
    }
    v1 = 0
    var v20 uint16 = 655
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
