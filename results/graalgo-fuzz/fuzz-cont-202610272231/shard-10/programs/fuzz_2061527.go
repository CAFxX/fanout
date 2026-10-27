package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 int16 }
func genFn0(p0 int64) uint16 {
    v5 := []int{128}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v5
    return 139
}
func genFn1(p0 int64) uint8 {
    var v6 uint32 = 42
    var v7 int = -2147483648
    v6 = 42
    fmt.Println(v6)
    if (0 == -128) {
        var v8 uint16 = 2
        switch "a" {
        case "\n\t":
            v6 -= v6
        case "foo bar":
            var v9 string = "\n\t"
            _ = v9
        default:
            v10 := GenS1{F0: 1}
            _ = v10
        }
        if true {
            if (100 <= -128) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v8 = (v8 << ((790) & 63))
                    var v12 int8 = 42
                    _ = fi11
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := []int{fi13}
                    _ = fi13
                    _ = v14
                }
            }
            fmt.Println(v8)
        }
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 2
    v3 := []string{("\n\t" + "\n\t"), "\n\t", (("" + ("hello" + ("" + ((("foo bar" + ("foo bar" + (((("foo bar" + "a") + "") + "a") + "foo bar"))) + "a") + "foo bar")))) + "\n\t"), "\n\t"}
    var v4 int64 = 0
    for fi15 := 0; fi15 < 17; fi15++ {
        v1 = v1
        var v16 bool = false
        _ = fi15
        _ = v16
    }
    v4 -= (v4 ^ (v4 % (((v4 >> ((300) & 63))) | 1)))
    v17 := GenS1{F0: (v2 * (^(v2)))}
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v1, fi18, v4)
        v2 = v2
        v1 += (-(v1))
        _ = fi18
    }
    if (-1 == 575) {
        fmt.Println(v1, v4, v2)
        fmt.Println(v2)
    }
    var v19 float64 = 0.1
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "a")
    }
    v20 := GenS2{F0: "hello", F1: 42, F2: v1}
    v1 += (v1 * (v1 * -1))
    v21 := GenS1{F0: 1}
    var v22 uint32 = 4294967295
    v23 := GenS2{F0: "\n\t", F1: genFn1(v4), F2: v1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch int(v2) {
    case 0:
        v3 = append(v3, (("hello" + ("hello" + "")) + "a"))
    case -1:
        fmt.Println(v4, v4)
    case 2147483647:
        fmt.Println(v19, v2)
    default:
        var v24 string = ("" + ("foo bar" + ("\n\t" + "\n\t")))
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
