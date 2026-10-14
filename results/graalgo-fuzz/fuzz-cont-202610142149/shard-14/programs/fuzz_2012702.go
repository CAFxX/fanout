package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint) uint32 {
    v4 := []int{}
    v5 := GenS1{F0: -1, F1: p0, F2: 32767}
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -1
    v3 := []string{v1, v1}
    var v6 uint = 1
    v2 -= v2
    v3 = append(v3, "")
    v7 := GenS1{F0: 187, F1: 319, F2: 1000}
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v6, v1, v2)
        for fi9 := 0; fi9 < 9; fi9++ {
            v2 += (v2 % (((v2 >> (((v6 >> ((2) & 63))) & 63))) | 1))
            if (false && (v6 == (v6 >> ((v6) & 63)))) {
                v1 = v1
                fmt.Println(v1, v2)
            }
            _ = fi9
        }
        _ = len(v3)
        _ = fi8
    }
    v2 += v2
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 = append(v3, (v1 + "hello"))
        _ = fi10
    }
    var v11 bool = false
    v6 -= v6
    switch 9223372036854775807 {
    case 4294967295:
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v2)
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(v11, v2)
                _ = len(v3)
                _ = fi13
            }
            _ = fi12
        }
    case 32768:
        fmt.Println(v6)
    case 9223372036854775807:
        var v14 int = -369
        _ = v14
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (1 > -57) {
        var v15 uint64 = 625
        var v16 int16 = 1000
        _ = v15
        _ = v16
    } else {
        v6 -= (v6 % ((1) | 1))
        if true {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int = 9223372036854775807
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS1{F0: fi17, F1: v6, F2: 42}
                    _ = fi19
                    _ = v20
                }
                switch v1 {
                case "":
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v22 := GenS1{F0: v18, F1: v6, F2: 1000}
                        _ = fi21
                        _ = v22
                    }
                case "\n\t":
                    v2 -= v2
                }
                _ = fi17
                _ = v18
            }
            v23 := GenS1{F0: 9223372036854775807, F1: v6, F2: 1000}
            if (false && (-9223372036854775808 < -889)) {
                switch 0 {
                case 4294967295:
                    var v24 uint = (v6 | 1)
                    _ = v24
                }
                v25 := GenS2{F0: ("\n\t" + "\n\t")}
                _ = v25
            }
            _ = v23
        }
    }
    if true {
        if len(v3) > 0 {
            v3[0] = v1
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            v11 = false
            _ = fi26
        }
        v6 -= v6
    }
    fmt.Println(v11)
    v1 = "\n\t"
    _ = len(v3)
    var v27 uint16 = 42
    v28 := GenS1{F0: -9223372036854775808, F1: v6, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v7
}
