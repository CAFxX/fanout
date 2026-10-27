package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint16) int8 {
    if (!false) {
        var v4 int32 = 1
        var v5 int16 = 42
        _ = v4
        _ = v5
    } else {
        var v6 uint64 = 1
        _ = v6
    }
    if true {
        if (!(-32768 > 1)) {
            v7 := GenS1{F0: 975, F1: -99, F2: -1}
            p0 += 255
            _ = v7
        } else {
            var v8 float64 = float64(0.0)
            _ = v8
        }
        p1 = p1
    }
    v9 := GenS1{F0: -2147483648, F1: -84, F2: -145}
    if (-46 > -1) {
        fmt.Println(p0)
        var v10 int8 = -3
        fmt.Println(p0, p1, p1)
        _ = v10
    }
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = v9
    return -22
}

func main() {
    var v1 uint32 = 944
    var v2 int64 = 1
    v3 := make(map[string]int)
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 bool = (false || (true && true))
        var v13 int = (fi11 ^ (fi11 >> ((2) & 63)))
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 -= (v2 >> ((18446744073709551615) & 63))
    for fi14 := 0; fi14 < 6; fi14++ {
        v1 -= (v1 - (v1 + (^(v1))))
        var v15 bool = (-823 != 2147483647)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 82, F2: -1}
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS1{F0: (fi18 * fi17), F1: 100, F2: -493}
            _ = fi18
            _ = v19
        }
        v20 := []int{fi17, fi17}
        v2 = (v2 - (v2 / (((-(v2))) | 1)))
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 int8 = -128
            var v24 bool = ((!true) && (-1 <= 1))
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int64 = v2
        switch 42 {
        case 127:
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 int16 = -32768
                _ = fi26
                _ = v27
            }
        case -136:
            var v28 string = (("foo bar" + "\n\t") + "\n\t")
            _ = v28
        case -2147483648:
            fmt.Println(v1)
        default:
            v29 := GenS2{F0: 23}
            _ = v29
        }
        _ = fi21
        _ = v25
    }
    var v30 int64 = (-(v2))
    v1 -= (^(v1))
    if (73 <= 105) {
        v31 := GenS1{F0: 4294967295, F1: 78, F2: -33}
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint64 = 609
            v1 = (v1 | v1)
            if (true && (255 < 42)) {
                var v34 uint32 = 65535
                for fi35 := 0; fi35 < 2; fi35++ {
                    fmt.Println(v2)
                    switch fi32 {
                    case -152:
                        var v36 int8 = -128
                        _ = v36
                    }
                    _ = fi35
                }
                v37 := GenS1{F0: 238, F1: 42, F2: v2}
                _ = v34
                _ = v37
            }
            _ = fi32
            _ = v33
        }
        v38 := GenS2{F0: 128}
        _ = v31
        _ = v38
    }
    _ = len(v3)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
}
