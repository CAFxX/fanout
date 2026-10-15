package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
func genFn0(p0 bool) uint8 {
    switch 32767 {
    case 2147483647:
        fmt.Println(p0)
    default:
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := []int{fi5, (fi5 & fi5), fi5}
            var v7 int16 = -1
            var v8 uint64 = 1
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 42, F1: -57, F2: ((("\n\t" + (((("" + "foo bar") + ("a" + "\n\t")) + "") + ("hello" + "hello"))) + "a") + "")}
    _ = p0
    _ = v9
    return 255
}
func genFn1() int32 {
    v10 := []int{}
    var v11 bool = true
    _ = v10
    _ = v11
    return -117
}

func main() {
    v1 := []uint32{42, 2, 2147483648}
    var v2 int16 = 42
    v3 := []string{"\n\t", "", "a", "a"}
    v4 := make(map[string]int)
    v12 := GenS1{F0: 268, F1: 0, F2: "hello"}
    var v13 uint16 = 473
    fmt.Println(v2)
    v2 = -458
    for fi14 := 0; fi14 < 18; fi14++ {
        if false {
            v3 = append(v3, ("\n\t" + "foo bar"))
        }
        v15 := GenS2{F0: -990, F1: -511, F2: 111}
        v13 -= (^(v13))
        _ = fi14
        _ = v15
    }
    var v16 int8 = 46
    var v17 int = 9223372036854775807
    fmt.Println(v17)
    var v18 int16 = v2
    v16 = (v16 - (v16 >> ((18446744073709551615) & 63)))
    for fi19 := 0; fi19 < 16; fi19++ {
        v4["\n\t"] = fi19
        var v20 uint16 = 32767
        _ = fi19
        _ = v20
    }
    v18 -= int16(v13)
    var v21 int = v17
    var v22 float64 = (1e10 * -2.25)
    _ = len(v3)
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int = int(v13)
        _ = fi23
        _ = v24
    }
    fmt.Println(v22, v16, v16)
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
