package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int8 = -88
    var v3 int64 = -9223372036854775808
    v1 = 1.5
    for fi4 := 0; fi4 < 8; fi4++ {
        if true {
            v2 = 42
        } else {
            if true {
                var v5 uint = 215
                fmt.Println(v3, v2, fi4)
                _ = v5
            }
        }
        _ = fi4
    }
    v6 := GenS1{F0: "a", F1: 32767, F2: ((!(0 > 896)) || (1000 < 0))}
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 = 1
        _ = fi7
    }
    var v8 uint64 = 42
    var v9 uint64 = v8
    var v10 float32 = 0.0
    for fi11 := 0; fi11 < 4; fi11++ {
        if true {
            fmt.Println(v8, v9, v3)
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(v1, v2)
                _ = fi12
            }
            var v13 int64 = v3
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            switch fi15 {
            case -248:
                var v16 int = int(v2)
                _ = v16
            case 0:
                if (0.1 != float64(v1)) {
                    fmt.Println(v1)
                    if (65535 != 47) {
                        v10 = (v1 - v10)
                    } else {
                        v17 := GenS1{F0: (("foo bar" + (("foo bar" + "") + ("" + "\n\t"))) + "a"), F1: 65535, F2: (!(528 <= 0))}
                        _ = v17
                    }
                }
            }
            var v18 uint32 = 2147483647
            var v19 float64 = 0.0
            fmt.Println(v1, v18, v19)
            _ = fi15
            _ = v18
            _ = v19
        }
        switch (fi14 / (((fi14 - (-(fi14)))) | 1)) {
        case 4294967295:
            v20 := []int{(fi14 / ((fi14) | 1))}
            _ = v20
        default:
            if ((v3 == v3) && false) {
                var v21 int32 = 2147483647
                _ = v21
            }
        }
        switch 578 {
        case -2147483648:
            switch fi14 {
            case -836:
                v22 := []int{fi14, (fi14 & (fi14 % ((786) | 1))), fi14, -2147483648}
                _ = v22
            case 726:
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := []int{}
                    _ = fi23
                    _ = v24
                }
            case 128:
                if ((-(v2)) <= (v2 >> ((4294967295) & 63))) {
                    if true {
                        v25 := []int{-513, fi14, fi14, fi14}
                        _ = v25
                    } else {
                        v26 := []int{38, 0}
                        var v27 uint8 = 0
                        _ = v26
                        _ = v27
                    }
                    switch "" {
                    case "\n\t":
                        v28 := GenS2{F0: 65535, F1: 42, F2: 175}
                        _ = v28
                    case "foo bar":
                        v1 = (v10 * (0.0 + v10))
                    default:
                        v29 := GenS1{F0: "a", F1: 442, F2: (((42 > 0) || (3.14159 <= (-2.25 * 0.0))) && true)}
                        _ = v29
                    }
                    v9 = v9
                } else {
                    var v30 int16 = 42
                    v31 := []int{(fi14 * -2147483648), -9223372036854775808, 127}
                    v3 += v3
                    _ = v30
                    _ = v31
                }
            default:
                if true {
                    v32 := GenS2{F0: 42, F1: 65535, F2: -1}
                    _ = v32
                }
            }
        case 32767:
            var v33 uint = 59
            _ = v33
        case 127:
            v10 = float32(-2.25)
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
