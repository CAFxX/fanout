package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0() bool {
    if true {
        switch 747 {
        case -2147483648:
            var v4 int32 = 2147483647
            _ = v4
        case -91:
            var v5 uint32 = 13
            _ = v5
        }
        var v6 int32 = -953
        _ = v6
    }
    var v7 float64 = float64(3.14159)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint32 = 2147483648
            v11 := []int{(^(fi9)), fi9}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v7, v7)
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := []int{fi8, (fi12 | fi12)}
            var v14 uint16 = 65535
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    v15 := []int{-1, 127, 32768, 929}
    v16 := GenS1{F0: 8, F1: 18446744073709551615}
    _ = v15
    _ = v16
    _ = v7
    return (!false)
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 bool = false
    switch ("foo bar" + ("\n\t" + ("foo bar" + (("foo bar" + ("foo bar" + "hello")) + "")))) {
    case "a":
        var v17 int = 859
        _ = v17
    }
    switch -9223372036854775808 {
    case 32768:
        v3 = false
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: 42, F1: 656}
        fmt.Println(v3, fi18)
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: 638, F2: 9223372036854775807}
    v1 = append(v1, 32768)
    switch -44 {
    case 519:
        v21 := v2[128]
        _ = v21
    }
    v3 = v3
    v22 := GenS2{F0: 32767, F1: 1000, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
