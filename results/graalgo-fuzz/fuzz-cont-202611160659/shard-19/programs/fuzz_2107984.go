package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int }
func genFn0() uint8 {
    v7 := GenS1{F0: -150}
    v8 := []int{-660, -2147483648, 0}
    var v9 uint16 = 644
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 uint32 = 826
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 bool = true
    v5 := make(map[uint32]bool)
    var v6 float32 = (0.0 * (float32((0.0 * (0.0 - (0.0 + (-2.25 + 1.5))))) * 3.14159))
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v1, fi10, v4)
        _ = fi10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            switch fi11 {
            case 32767:
                v13 := []int{fi11}
                _ = v13
            }
            if false {
                v6 = -2.25
                if v4 {
                    v14 := []int{(fi12 | (fi11 * fi11))}
                    _ = v14
                }
            } else {
                v15 := GenS2{F0: 1, F1: 1}
                _ = v15
            }
            v16 := GenS3{F0: 2147483647}
            if false {
                v1 += v1
            } else {
                switch "" {
                case "hello":
                    fmt.Println(fi11, v6, v4)
                default:
                    var v17 int64 = -9223372036854775808
                    v18 := GenS3{F0: 201}
                    _ = v17
                    _ = v18
                }
            }
            v1 -= (^(v1))
            _ = fi12
            _ = v16
        }
        switch "foo bar" {
        case "a":
            var v19 int64 = 695
            _ = v19
        case "\n\t":
            if v4 {
                fmt.Println(v6, v1)
            } else {
                if (v4 || v4) {
                    v6 = v6
                } else {
                    switch 32767 {
                    case -9223372036854775808:
                        v20 := v3[((((((("a" + ("\n\t" + "")) + "a") + "a") + "a") + "hello") + ((((("hello" + "foo bar") + "\n\t") + "foo bar") + ((("\n\t" + ("\n\t" + "\n\t")) + ("a" + ("foo bar" + "foo bar"))) + "")) + "hello")) + "a")]
                        _ = v20
                    default:
                        fmt.Println(fi11, v1)
                    }
                }
                v4 = (357 != 1)
            }
        default:
            var v21 uint64 = 818
            _ = v21
        }
        _ = fi11
    }
    v4 = v4
    var v22 int8 = -95
    v23 := GenS1{F0: 42}
    v24 := GenS2{F0: 794, F1: -2147483648}
    if false {
        var v25 int32 = 42
        v26 := []int{-1}
        for fi27 := 0; fi27 < 5; fi27++ {
            v6 = v6
            v6 -= 0.0
            _ = fi27
        }
        _ = v25
        _ = v26
    } else {
        fmt.Println(v22, v22)
        var v28 int64 = 0
        fmt.Println(v6, v1, v4)
        _ = v28
    }
    v1 = (^(v1))
    var v29 int = 873
    if (!(!((true && v4) || v4))) {
        v1 -= 0
        var v30 float64 = 1e10
        v30 += 3.14159
        _ = v30
    }
    v6 = 1e10
    _ = len(v5)
    v4 = (!v4)
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint16 = 560
        var v33 uint16 = 255
        v34 := GenS3{F0: (^(fi31))}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    switch "hello" {
    case "a":
        delete(v2, (v1 >> ((904) & 63)))
    default:
        var v35 bool = (0 > 1)
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
