package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
func genFn0() uint {
    var v6 float64 = float64(1e10)
    v7 := GenS1{F0: 18446744073709551615, F1: (-1 >= 127), F2: 18446744073709551615}
    fmt.Println(v6)
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{}
        fmt.Println(v6, v6)
        _ = fi8
        _ = v9
    }
    if ((1 >= 255) && (-9223372036854775808 <= 32767)) {
        v10 := GenS1{F0: 21, F1: false, F2: 839}
        if false {
            if false {
                var v11 bool = (!true)
                var v12 int64 = 1
                _ = v11
                _ = v12
            } else {
                v6 = 1e10
            }
            v13 := []int{}
            v6 += v6
            _ = v13
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int8 = 127
                _ = fi14
                _ = v15
            }
            fmt.Println(v6)
            if true {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int = -2147483648
                    _ = fi16
                    _ = v17
                }
            } else {
                if true {
                    v18 := GenS1{F0: 559, F1: true, F2: 2}
                    _ = v18
                }
            }
        }
        switch "a" {
        case "foo bar":
            var v19 int = 1
            _ = v19
        default:
            v20 := GenS1{F0: 546, F1: (!false), F2: 1}
            _ = v20
        }
        _ = v10
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int16) uint64 {
    v21 := GenS2{F0: 1}
    p0 = 597
    if true {
        var v22 float64 = (float64(-2.25) + 0.0)
        if ((42 < 0) || (p0 < p0)) {
            v22 -= v22
            var v23 int64 = 794
            v24 := GenS2{F0: 65535}
            _ = v23
            _ = v24
        }
        _ = v22
    } else {
        var v25 int32 = 644
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := []int{}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    if true {
        v28 := GenS2{F0: 65535}
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := []int{}
            if ((true || false) && true) {
                var v31 float32 = 0.1
                var v32 uint = 592
                _ = v31
                _ = v32
            }
            if (24 > 2147483647) {
                p0 = (p0 * p0)
            } else {
                v33 := GenS1{F0: genFn0(), F1: ((fi29 * fi29) != fi29), F2: 451}
                p0 = -60
                _ = v33
            }
            _ = fi29
            _ = v30
        }
        fmt.Println(p0)
        _ = v28
    } else {
        v34 := GenS1{F0: 18446744073709551615, F1: false, F2: 18446744073709551615}
        _ = v34
    }
    v35 := GenS2{F0: 775}
    _ = p0
    _ = v21
    _ = v35
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1
    var v3 int16 = 1
    v4 := make(map[string]int)
    var v5 uint8 = 91
    var v36 uint32 = 2
    var v37 uint16 = 255
    fmt.Println(v5)
    v5 = v5
    if (!(v2 == 0)) {
        var v38 uint16 = 0
        v39 := GenS1{F0: 1, F1: ((false && true) && true), F2: genFn1(v2)}
        _ = v38
        _ = v39
    } else {
        if (!false) {
            v4[""] = 396
            switch "hello" {
            case "a":
                switch -1 {
                case 32768:
                    v3 -= (v3 % ((v3) | 1))
                }
            case "":
                v2 = 1000
            default:
                for fi40 := 0; fi40 < 15; fi40++ {
                    var v41 uint8 = 48
                    var v42 int16 = (v2 * -32768)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            var v43 float32 = 0.0
            _ = v43
        }
        v3 = 0
        v44 := GenS1{F0: 740, F1: ((v37 <= 42) && (genFn1(42) >= 0)), F2: 42}
        _ = v44
    }
    v2 = -1
    v36 = v36
    for fi45 := 0; fi45 < 13; fi45++ {
        v5 = (v5 * (v5 + v5))
        v5 = (v5 & v5)
        _ = fi45
    }
    var v46 float64 = -2.25
    var v47 int8 = -128
    for fi48 := 0; fi48 < 17; fi48++ {
        for fi49 := 0; fi49 < 2; fi49++ {
            _ = len(v4)
            v50 := GenS2{F0: (v37 - v37)}
            v51 := GenS2{F0: (v37 * v37)}
            _ = fi49
            _ = v50
            _ = v51
        }
        v37 += uint16(v5)
        var v52 uint8 = (^(v5))
        _ = fi48
        _ = v52
    }
    fmt.Println(v5, v36, v37)
    for fi53 := 0; fi53 < 17; fi53++ {
        v54 := GenS3{F0: 9223372036854775807}
        v55 := []int{(^(fi53))}
        var v56 int8 = 42
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        v58 := GenS2{F0: 42}
        _ = fi57
        _ = v58
    }
    v59 := v4[("foo bar" + ("a" + "a"))]
    var v60 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v59
    _ = v60
}
