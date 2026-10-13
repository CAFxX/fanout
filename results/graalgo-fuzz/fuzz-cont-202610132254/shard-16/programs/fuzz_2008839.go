package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint64) uint {
    if false {
        v6 := GenS1{F0: false, F1: (42 != 127), F2: true}
        var v7 int8 = 42
        _ = v6
        _ = v7
    } else {
        var v8 int16 = -32768
        v9 := []int{}
        _ = v8
        _ = v9
    }
    var v10 int8 = -82
    if true {
        fmt.Println(p0, v10)
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{(-(fi11)), fi11, fi11, fi11}
            _ = fi11
            _ = v12
        }
    } else {
        if ((255 != 133) || true) {
            var v13 uint = 795
            switch int(v13) {
            case 2147483647:
                v14 := GenS1{F0: true, F1: (p0 < p0), F2: false}
                _ = v14
            case 156:
                var v15 float32 = (1.5 - float32(3.14159))
                _ = v15
            }
            _ = v13
        }
    }
    _ = p0
    _ = v10
    return 0
}
func genFn1(p0 bool, p1 bool) uint32 {
    var v16 uint = 625
    for fi17 := 0; fi17 < 5; fi17++ {
        if p0 {
            var v18 int8 = -1
            var v19 bool = false
            _ = v18
            _ = v19
        } else {
            switch "" {
            case "foo bar":
                var v20 int16 = 42
                _ = v20
            case "":
                v16 = (v16 ^ (v16 | 791))
            case "\n\t":
                p1 = true
            default:
                var v21 int16 = 42
                _ = v21
            }
            if true {
                if (false || (false || (142 <= 255))) {
                    var v22 uint8 = 224
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: p0, F1: (!p1), F2: p1}
                p0 = p1
                _ = v23
            }
            switch "" {
            case "a":
                v24 := []int{(fi17 ^ fi17)}
                _ = v24
            case "foo bar":
                v25 := []int{fi17, fi17, (fi17 + fi17), -9223372036854775808}
                _ = v25
            default:
                v26 := []int{128, (^(fi17)), fi17, 32767}
                _ = v26
            }
        }
        _ = fi17
    }
    switch -488 {
    case 32767:
        var v27 uint8 = 2
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 bool = p1
            v30 = (v30 || (v30 && (p0 && (170 > -1))))
            var v31 int64 = -9223372036854775808
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    _ = p0
    _ = p1
    _ = v16
    return 534
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 128
    var v3 uint = 614
    v4 := make(map[string]uint32)
    var v5 uint64 = 1
    v3 += (v3 + v3)
    v3 -= v3
    v32 := GenS1{F0: (((^(v5)) > v5) && ((v3 >= genFn0(v5)) || (42 >= 0))), F1: (((1e10 + 1e10) * float32(3.14159)) <= ((-2.25 - 0.1) * (float32(1.5) * (0.1 / ((1e10 * (0.1 / 1.0)) - float32(float64(1.5))))))), F2: (32767 >= int16(v3))}
    var v33 uint8 = 127
    fmt.Println(v5, v33, v3)
    var v34 float32 = (1e10 * (1e10 + (((1e10 / (1e10 + (3.14159 * float32(1.5)))) - (0.0 * 0.0)) * (1e10 / 0.1))))
    v35 := []int{127, v2, v2, v2}
    for fi36 := 0; fi36 < 11; fi36++ {
        if (v34 >= (1e10 * (v34 + 3.14159))) {
            v37 := GenS1{F0: true, F1: false, F2: true}
            var v38 int32 = -714
            var v39 bool = (false || (309 < 2147483648))
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi36
    }
    v5 = (v5 * 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
