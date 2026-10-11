package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 uint16, p1 int) uint8 {
    switch ("a" + ("a" + "")) {
    case "hello":
        fmt.Println(p1, p0)
    default:
        v6 := GenS1{F0: 1, F1: "hello", F2: 2}
        _ = v6
    }
    var v7 int8 = -125
    fmt.Println(p0, p1, v7)
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            var v9 uint8 = 128
            v10 := []int{(fi8 << ((620) & 63))}
            _ = v10
            _ = v9
        } else {
            v11 := []int{fi8, -1}
            v12 := GenS2{F0: 2147483647, F1: "\n\t"}
            _ = v11
            _ = v12
        }
        v13 := []int{32767}
        if (32767 != 1000) {
            var v14 uint32 = 0
            _ = len(v13)
            _ = v14
        } else {
            p1 = 1
        }
        _ = fi8
        _ = v13
    }
    var v15 float64 = 0.0
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 int = -2147483648
    v3 := []uint32{2, 42, 2147483648, 0}
    var v4 uint16 = v1
    var v5 string = "a"
    fmt.Println(v4, v2)
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v16 string = "a"
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v19 uint32 = 2147483647
                _ = v19
            case 0:
                fmt.Println(v1)
            }
            var v20 bool = true
            _ = fi18
            _ = v20
        }
        _ = len(v3)
        var v21 bool = ((45 != 2) && true)
        _ = fi17
        _ = v21
    }
    if (v2 < (^(v2))) {
        for fi22 := 0; fi22 < 19; fi22++ {
            _ = len(v3)
            v23 := GenS2{F0: 2147483648, F1: v5}
            _ = len(v3)
            _ = fi22
            _ = v23
        }
        if len(v3) > 0 {
            v3[0] = 0
        }
        v1 -= (^(v1))
    } else {
        switch ("foo bar" + ((v5 + v5) + (v5 + ("" + "foo bar")))) {
        case "\n\t":
            var v24 int32 = -1
            _ = v24
        }
    }
    v5 = (v5 + v5)
    fmt.Println(v2)
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
