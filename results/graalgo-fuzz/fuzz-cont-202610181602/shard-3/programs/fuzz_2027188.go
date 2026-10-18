package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 uint }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int64 }
func genFn0(p0 int64) uint32 {
    switch -776 {
    case 42:
        v6 := GenS1{F0: 32768, F1: "foo bar", F2: (2147483648 >= 619)}
        _ = v6
    case 2147483647:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint16 = 759
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 128, F1: ("foo bar" + "foo bar"), F2: false}
    var v10 bool = (128 >= 195)
    v11 := GenS1{F0: 4294967295, F1: "foo bar", F2: v10}
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := []int{fi12, (fi12 << ((4294967295) & 63)), (fi12 % (((fi12 >> ((506) & 63))) | 1)), (fi12 & fi12)}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 65535
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "a"
    var v3 uint16 = 0
    var v4 float64 = 1e10
    v5 := []int{32767}
    v14 := GenS2{F0: 578, F1: -2147483648, F2: 95}
    v15 := GenS2{F0: genFn0(-777), F1: 42, F2: 358}
    v16 := GenS1{F0: 4294967295, F1: v2, F2: true}
    var v17 uint = 255
    switch 2147483647 {
    case 1:
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int = fi18
            v20 := GenS3{F0: 42, F1: v3}
            _ = fi18
            _ = v19
            _ = v20
        }
    case -1:
        if false {
            v21 := GenS4{F0: 42}
            _ = v21
        }
    default:
        var v22 int16 = 1
        _ = v22
    }
    v23 := GenS5{F0: 4294967295, F1: 128, F2: -1}
    _ = len(v5)
    v1 -= 0.1
    var v24 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
