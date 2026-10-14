package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool, p1 uint32, p2 uint32) int8 {
    switch 32768 {
    case 32767:
        p0 = (!true)
    }
    var v7 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint = 117
    v2 := make(map[string]string)
    var v3 float32 = (-2.25 + 1.5)
    v4 := make(map[uint32]int)
    var v5 float64 = ((3.14159 + 1.5) - float64((float32(((1.5 - 0.1) * (float64(1e10) * (0.0 * float64(v3))))) + 1.5)))
    var v6 int8 = -128
    var v8 uint = v1
    v6 -= 21
    v6 += v6
    if false {
        v9 := []int{106, 9223372036854775807, 300}
        _ = v9
    }
    v10 := GenS1{F0: 743, F1: -2147483648}
    if (473 != 0) {
        if ((v3 + (float32(v5) + v3)) <= v3) {
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := []int{}
                var v13 float64 = (3.14159 / 1e10)
                for fi14 := 0; fi14 < 17; fi14++ {
                    v6 -= v6
                    _ = fi14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            v15 := GenS1{F0: 678, F1: -217}
            switch -2147483648 {
            case -41:
                v6 += (^(v6))
            default:
                v16 := v4[0]
                _ = v16
            }
            _ = v15
        }
        var v17 float64 = (1e10 / v5)
        var v18 int8 = 0
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = (fi19 ^ 2147483647)
                _ = len(v4)
                v5 = v5
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        var v22 uint16 = 42
        v23 := GenS1{F0: 1, F1: -854}
        _ = v22
        _ = v23
    }
    var v24 uint = v8
    var v25 int64 = -1
    var v26 uint32 = 913
    v27 := GenS1{F0: -32768, F1: 349}
    if false {
        var v28 int8 = (v6 * (v6 - genFn0(false, (v26 | (v26 >> ((0) & 63))), (v26 + 4294967295))))
        v4[(^(v26))] = 9223372036854775807
        switch ("a" + "") {
        case "":
            v5 = v5
        case "\n\t":
            v29 := GenS2{F0: 842, F1: 2}
            _ = v29
        case "a":
            var v30 uint8 = 202
            _ = v30
        default:
            if (!(int16(v26) == 0)) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := []int{570, fi31, fi31, (fi31 * -2147483648)}
                    var v33 int32 = -1
                    var v34 int32 = (-(v33))
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                fmt.Println(v24, v3, v8)
            }
        }
        _ = v28
    }
    v35 := GenS2{F0: 859, F1: 152}
    if true {
        for fi36 := 0; fi36 < 8; fi36++ {
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 float64 = (-2.25 * v5)
                _ = fi37
                _ = v38
            }
            v1 = v1
            _ = len(v4)
            _ = fi36
        }
        v39 := []int{-9223372036854775808, 986, -2147483648}
        v39 = append(v39, 128)
        _ = v39
    } else {
        v24 = 821
    }
    v5 = float64(0.0)
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 13; fi41++ {
            delete(v2, "hello")
            v42 := []int{}
            var v43 uint64 = 4294967295
            var v44 float32 = v3
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
