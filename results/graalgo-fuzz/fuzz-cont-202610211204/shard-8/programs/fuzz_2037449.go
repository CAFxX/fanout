package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 bool) int8 {
    var v5 uint64 = 42
    if (9223372036854775807 >= 32768) {
        if false {
            switch "" {
            case "hello":
                p0 = false
            }
            var v6 int32 = 0
            p0 = (!("" >= ((("" + (("hello" + "hello") + "foo bar")) + "a") + (("" + (((("\n\t" + "\n\t") + (("hello" + "") + "foo bar")) + ("a" + "\n\t")) + "a")) + "foo bar"))))
            _ = v6
        } else {
            v7 := []int{-9223372036854775808}
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: -2147483648}
                var v10 int32 = 833
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        var v11 uint32 = 2
        _ = v11
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return 100
}

func main() {
    var v1 int32 = 0
    var v2 int = -1
    v3 := make(map[uint32]bool)
    v4 := []int{(v2 - (v2 % ((32768) | 1))), v2}
    switch (v2 - 42) {
    case 42:
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint = 4294967295
                _ = fi12
                _ = v13
            }
        }
    case -9223372036854775808:
        fmt.Println(v2)
    case -953:
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint64 = 1
            _ = fi14
            _ = v15
        }
    default:
        switch "hello" {
        case "a":
            v16 := GenS1{F0: (^(v1))}
            _ = v16
        case "hello":
            switch (((("hello" + "\n\t") + (("hello" + ("\n\t" + "\n\t")) + ("foo bar" + "foo bar"))) + "hello") + (("a" + "a") + "foo bar")) {
            case "a":
                fmt.Println(v2)
            case "\n\t":
                v17 := GenS2{F0: -505, F1: 142}
                _ = v17
            default:
                if false {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int = (v2 + v2)
                        _ = fi18
                        _ = v19
                    }
                    v1 -= (^(v1))
                    switch v2 {
                    case 127:
                        fmt.Println(v2, v1)
                    case -1:
                        if true {
                            for fi20 := 0; fi20 < 11; fi20++ {
                                fmt.Println(v1, v2, v2)
                                _ = fi20
                            }
                            var v21 float64 = 1e10
                            _ = v21
                        }
                    case 42:
                        if (!false) {
                            if (0.0 >= 0.1) {
                                v2 += v2
                                v3[4294967295] = false
                            } else {
                                v4 = append(v4, 488)
                            }
                            v1 = (-(v1))
                        } else {
                            delete(v3, 900)
                            v2 -= v2
                        }
                    }
                }
            }
        default:
            var v22 uint64 = 570
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 bool = (((!false) && (!(171 > 255))) || ((42 < 42) || false))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 bool = ((v1 ^ 42) >= (v1 / ((0) | 1)))
        for fi27 := 0; fi27 < 17; fi27++ {
            v1 = v1
            v2 = fi27
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    if (0 <= 192) {
        var v28 int64 = -9223372036854775808
        v2 = (v2 / ((-561) | 1))
        _ = v28
    } else {
        var v29 bool = ((42 <= 861) && (!false))
        _ = v29
    }
    if true {
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 int16 = 32767
            switch (v2 % ((-1) | 1)) {
            case 0:
                fmt.Println(v31, v1)
            }
            _ = fi30
            _ = v31
        }
        v3[183] = (!((true && false) || false))
        for fi32 := 0; fi32 < 4; fi32++ {
            switch (fi32 & 996) {
            case 32767:
                v1 -= (v1 << ((94) & 63))
            case -770:
                v33 := GenS1{F0: (v1 / ((42) | 1))}
                _ = v33
            default:
                var v34 float64 = 0.1
                _ = v34
            }
            v35 := v3[585]
            v2 = v2
            _ = fi32
            _ = v35
        }
    }
    var v36 int64 = 42
    if (((v2 <= v2) || ((v1 | v1) > v1)) && true) {
        v37 := GenS2{F0: (v36 | int64(v1)), F1: 1}
        v2 -= v2
        v2 = 128
        _ = v37
    } else {
        delete(v3, 0)
    }
    var v38 int8 = -128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
}
