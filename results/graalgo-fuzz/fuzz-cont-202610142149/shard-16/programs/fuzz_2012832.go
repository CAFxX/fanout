package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int) uint8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 2, F1: false, F2: 18446744073709551615}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 18446744073709551615, F1: (!true), F2: 817}
    _ = p0
    _ = v6
    return 128
}
func genFn1(p0 int32, p1 int8) uint8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{-9223372036854775808, fi7, (fi7 - 127), (-(fi7))}
        v8 = append(v8, (fi7 + fi7))
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 539
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 663
    v2 := make(map[uint32]string)
    var v3 bool = true
    _ = len(v2)
    var v10 int32 = -1
    v11 := GenS1{F0: uint(v1), F1: (!true), F2: 42}
    v12 := GenS2{F0: 0}
    switch 822 {
    case 92:
        _ = len(v2)
    case -490:
        var v13 int16 = 42
        _ = v13
    case 42:
        var v14 int = 32767
        _ = v14
    default:
        v15 := v2[v1]
        _ = v15
    }
    switch "hello" {
    case "":
        v3 = false
    case "\n\t":
        var v16 int64 = 682
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 float64 = 0.1
        v19 := GenS1{F0: 18446744073709551615, F1: v3, F2: 42}
        switch (("foo bar" + "foo bar") + (("hello" + ("hello" + ("foo bar" + "a"))) + "foo bar")) {
        case "\n\t":
            fmt.Println(fi17, v10)
        case "hello":
            switch ("" + ("" + ("\n\t" + "foo bar"))) {
            case "\n\t":
                var v20 int8 = 42
                _ = v20
            default:
                switch "\n\t" {
                case "hello":
                    v18 = 0.0
                default:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v10 += v10
                        _ = fi21
                    }
                }
            }
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
