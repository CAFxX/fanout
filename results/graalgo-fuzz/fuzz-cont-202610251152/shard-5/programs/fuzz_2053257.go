package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 int64) int64 {
    v6 := []int{128, -2147483648, 264}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 2}
        p0 = 0
        p0 = 9223372036854775807
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := []int{}
    var v2 int = 2147483647
    v3 := make(map[string]uint32)
    var v4 uint64 = 1
    v5 := make(map[string]int)
    var v9 uint16 = 65535
    if (-1 == (v2 * v2)) {
        if len(v1) > 0 {
            v1[0] = (v2 | v2)
        }
        var v10 int64 = -847
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 2}
            fmt.Println(v9, v2)
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := GenS2{F0: 100, F1: -1}
        var v14 int8 = 100
        v15 := GenS1{F0: 224}
        _ = v13
        _ = v14
        _ = v15
    }
    if ((v2 * (v2 * (v2 >> ((182) & 63)))) != (v2 ^ (v2 | (-(v2))))) {
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS2{F0: 31, F1: 2147483647}
            if (!(-32768 != -529)) {
                v9 += (v9 % (((v9 % ((823) | 1))) | 1))
                var v18 uint16 = (^(v9))
                if ((522 != 448) && false) {
                    if false {
                        var v19 int32 = 42
                        v20 := GenS2{F0: -128, F1: int32(v2)}
                        _ = v19
                        _ = v20
                    }
                    v4 -= v4
                } else {
                    v21 := GenS3{F0: (v4 / ((2) | 1)), F1: 0, F2: (true && false)}
                    v2 -= (^(v2))
                    _ = v21
                }
                _ = v18
            }
            var v22 float32 = 3.14159
            _ = fi16
            _ = v17
            _ = v22
        }
        var v23 float64 = (1.5 + float64(0.0))
        _ = v23
    }
    var v24 int16 = 1
    v4 += v4
    v3[""] = 1
    for fi25 := 0; fi25 < 19; fi25++ {
        if ((-2147483648 < -2147483648) || (127 != 1)) {
            v4 = 399
            var v26 int8 = 41
            _ = v26
        } else {
            v5["hello"] = v2
            switch "a" {
            case "\n\t":
                v27 := GenS4{F0: -46, F1: 729, F2: (v4 >> ((153) & 63))}
                _ = v27
            }
        }
        var v28 int64 = 275
        v29 := GenS4{F0: -1, F1: 607, F2: 0}
        _ = fi25
        _ = v28
        _ = v29
    }
    switch "foo bar" {
    case "a":
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := GenS3{F0: v4, F1: -333, F2: false}
            _ = fi30
            _ = v31
        }
    default:
        v2 = v2
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            v2 += fi33
            _ = fi33
        }
        fmt.Println(v24, v24)
        fmt.Println(v2, v24, v4)
        _ = fi32
    }
    var v34 int32 = 42
    var v35 uint64 = v4
    v35 -= (v35 * (v4 << ((654) & 63)))
    switch (("a" + ("hello" + ("" + "foo bar"))) + ("\n\t" + "hello")) {
    case "hello":
        delete(v5, ("hello" + ((("hello" + "a") + ("hello" + "a")) + ("" + (("foo bar" + "") + "")))))
    case "a":
        switch (("a" + "") + ("a" + "")) {
        case "foo bar":
            if false {
                v36 := GenS3{F0: 18446744073709551615, F1: 42, F2: (!true)}
                var v37 float64 = (((0.1 * (-2.25 * 1.5)) * (-2.25 * 0.1)) + float64((1.5 * (((3.14159 * 0.0) + ((1.5 - 0.0) - 0.0)) / 1.0))))
                _ = v36
                _ = v37
            } else {
                var v38 string = ""
                v35 = v35
                v9 -= v9
                _ = v38
            }
        case "a":
            for fi39 := 0; fi39 < 9; fi39++ {
                for fi40 := 0; fi40 < 17; fi40++ {
                    var v41 uint8 = 1
                    if false {
                        if false {
                            v34 -= -839
                        }
                    } else {
                        _ = len(v5)
                    }
                    _ = fi40
                    _ = v41
                }
                fmt.Println(v34, v24)
                _ = fi39
            }
        }
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        if false {
            switch (v2 + (^(fi42))) {
            case -1:
                v9 = v9
            case -758:
                v43 := GenS4{F0: 42, F1: 4294967295, F2: 42}
                _ = v43
            default:
                var v44 int = v2
                _ = v44
            }
        } else {
            var v45 uint16 = v9
            for fi46 := 0; fi46 < 4; fi46++ {
                _ = len(v3)
                v24 += int16(v34)
                v9 = v45
                _ = fi46
            }
            _ = v45
        }
        _ = fi42
    }
    var v47 bool = (v24 != v24)
    fmt.Println(v24, v9, v34)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v47
    _ = v5
    _ = v9
}
