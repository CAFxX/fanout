package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0() uint64 {
    switch ("hello" + "") {
    case "a":
        var v7 float32 = ((1e10 * 1e10) * (-2.25 + 0.1))
        _ = v7
    case "hello":
        v8 := []int{}
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case -965:
            var v10 uint8 = 9
            _ = v10
        }
        var v11 int = (fi9 ^ (-(fi9)))
        _ = fi9
        _ = v11
    }
    return 721
}
func genFn1(p0 uint32) uint8 {
    var v12 string = "a"
    var v13 int = 32768
    switch v12 {
    case "a":
        v13 = (v13 | v13)
    case "foo bar":
        var v14 int = v13
        _ = v14
    case "hello":
        if (!(2 < 42)) {
            fmt.Println(p0)
            var v15 int = (v13 >> ((18446744073709551615) & 63))
            v15 += 9223372036854775807
            _ = v15
        }
    }
    v13 = 482
    _ = p0
    _ = v12
    _ = v13
    return 88
}

func main() {
    var v1 int32 = -64
    var v2 int = 1
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    var v5 uint32 = 460
    var v6 float64 = float64(0.1)
    switch "hello" {
    case "hello":
        switch ("\n\t" + (("a" + "") + "foo bar")) {
        case "\n\t":
            v16 := GenS1{F0: 4294967295, F1: 285, F2: v5}
            _ = v16
        case "foo bar":
            switch -2147483648 {
            case 2147483647:
                v5 -= v5
            default:
                var v17 int = v2
                _ = v17
            }
        default:
            v6 = float64(3.14159)
        }
    case "foo bar":
        v18 := GenS1{F0: 42, F1: 179, F2: v5}
        _ = v18
    case "\n\t":
        if true {
            v19 := GenS1{F0: 65535, F1: 18446744073709551615, F2: v5}
            _ = v19
        }
    default:
        var v20 uint32 = 350
        _ = v20
    }
    if false {
        var v21 int8 = 73
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int = 127
        _ = fi22
        _ = v23
    }
    v2 -= v2
    fmt.Println(v5)
    var v24 uint16 = 1
    v25 := GenS1{F0: uint(v5), F1: 42, F2: (v5 << ((1) & 63))}
    v4[v5] = 0
    fmt.Println(v2)
    v2 -= (v2 + v2)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
