package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int64; F1 uint16; F2 int }
func genFn0() int64 {
    v7 := []int{128, 127}
    switch 9223372036854775807 {
    case -9223372036854775808:
        if true {
            if (2 <= 42) {
                v7 = append(v7, 9223372036854775807)
            } else {
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 int32 = 0
                    _ = fi8
                    _ = v9
                }
            }
        }
    case 1:
        var v10 float64 = (1e10 - 1.5)
        _ = v10
    case -596:
        v11 := GenS1{F0: 1, F1: ("\n\t" + "\n\t")}
        _ = v11
    default:
        var v12 int8 = 42
        _ = v12
    }
    var v13 uint8 = 2
    if (!false) {
        for fi14 := 0; fi14 < 20; fi14++ {
            if (840 > 1) {
                if (false || (!(("foo bar" != ("" + "\n\t")) || false))) {
                    v15 := GenS2{F0: 391, F1: -963}
                    fmt.Println(fi14, v13)
                    _ = v15
                }
            }
            switch (fi14 << ((42) & 63)) {
            case 0:
                var v16 int64 = 0
                _ = v16
            case -9223372036854775808:
                var v17 int64 = 661
                _ = v17
            case 511:
                v13 = 41
            }
            _ = fi14
        }
        var v18 uint32 = 911
        _ = v18
    } else {
        v13 = v13
        v13 = 1
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    if false {
        _ = len(v7)
        _ = len(v7)
        v13 += 191
    }
    _ = v13
    _ = v7
    return 1
}

func main() {
    var v1 bool = true
    v2 := []string{("hello" + "\n\t"), "foo bar", "hello"}
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    v5 := make(map[string]uint32)
    v6 := []uint32{42}
    _ = len(v6)
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 string = "foo bar"
        switch ((v20 + ("a" + "foo bar")) + "\n\t") {
        case "foo bar":
            v21 := GenS3{F0: (^(v4))}
            _ = v21
        case "\n\t":
            v22 := v5["hello"]
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v23 := GenS4{F0: 108, F1: 9223372036854775807, F2: (v4 % ((v4) | 1))}
    for fi24 := 0; fi24 < 19; fi24++ {
        switch (fi24 - (fi24 | (fi24 << ((18446744073709551615) & 63)))) {
        case -739:
            switch "foo bar" {
            case "a":
                v1 = v1
            }
        case -2147483648:
            v25 := GenS2{F0: 2147483647, F1: 2147483647}
            _ = v25
        default:
            for fi26 := 0; fi26 < 2; fi26++ {
                if (42 <= 42) {
                    v4 -= v4
                    v1 = (v1 || (0 <= 0))
                }
                v4 += v4
                _ = fi26
            }
        }
        _ = fi24
    }
    v3["a"] = ((2 <= 87) && (9223372036854775807 < 697))
    v4 = 654
    var v27 uint8 = 1
    v28 := v5["hello"]
    v29 := GenS5{F0: 42, F1: "a"}
    v30 := GenS4{F0: (v4 * (v4 * (v4 - v4))), F1: -52, F2: (^(v4))}
    var v31 int32 = -109
    v31 = (v31 % ((v31) | 1))
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS6{F0: 9223372036854775807, F1: 252, F2: (fi32 | (fi32 ^ (fi32 | -105)))}
        for fi34 := 0; fi34 < 2; fi34++ {
            switch fi32 {
            case -2147483648:
                var v35 int16 = 407
                _ = v35
            case 9223372036854775807:
                var v36 uint = 255
                _ = v36
            default:
                var v37 int = fi32
                _ = v37
            }
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    v27 = v27
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
