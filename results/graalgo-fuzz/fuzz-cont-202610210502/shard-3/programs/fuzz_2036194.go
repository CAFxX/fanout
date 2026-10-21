package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 int }
func genFn0() int8 {
    if true {
        var v6 int64 = 9223372036854775807
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 int64 = 1
    v10 := GenS1{F0: 0, F1: false, F2: 3}
    _ = v10
    _ = v9
    return -1
}
func genFn1(p0 uint) int8 {
    v11 := GenS1{F0: 2, F1: (true || true), F2: -793}
    switch "hello" {
    case "a":
        if false {
            if true {
                if (!(!true)) {
                    var v12 float64 = float64(1e10)
                    _ = v12
                } else {
                    fmt.Println(p0)
                    var v13 int = 42
                    _ = v13
                }
                p0 = (p0 | (p0 - (p0 & 255)))
                fmt.Println(p0)
            }
            fmt.Println(p0)
        }
    case "":
        p0 -= p0
    case "\n\t":
        fmt.Println(p0)
    }
    _ = p0
    _ = v11
    return genFn0()
}

func main() {
    var v1 bool = false
    var v2 int16 = -1
    var v3 int16 = v2
    var v4 int32 = 0
    var v5 uint32 = 4294967295
    fmt.Println(v5, v3)
    if true {
        switch 571 {
        case 805:
            fmt.Println(v4, v1)
        default:
            fmt.Println(v1, v3, v2)
        }
        v1 = v1
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint32 = v5
            var v16 int32 = int32(v3)
            v2 += (v3 % ((v2) | 1))
            _ = fi14
            _ = v15
            _ = v16
        }
        v5 -= v5
    }
    var v17 bool = v1
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{fi18, (fi18 + 9223372036854775807)}
        fmt.Println(v17, fi18, v5)
        var v20 uint16 = 0
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: -401}
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := []int{(fi22 & (fi22 - 32767))}
        if len(v23) > 0 {
            v23[0] = fi22
        }
        v24 := GenS1{F0: 968, F1: (fi22 != fi22), F2: (v4 >> ((18446744073709551615) & 63))}
        _ = fi22
        _ = v23
        _ = v24
    }
    v3 = v3
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS1{F0: 165, F1: ((271 != uint64(v5)) && ((!v17) || false)), F2: (^(v4))}
        var v27 string = "foo bar"
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int32 = (v4 & (v4 ^ (v4 ^ v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
