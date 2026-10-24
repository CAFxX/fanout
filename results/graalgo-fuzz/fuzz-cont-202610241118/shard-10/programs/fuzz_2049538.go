package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
func genFn0() string {
    switch "\n\t" {
    case "\n\t":
        v7 := GenS1{F0: 2, F1: -799}
        _ = v7
    case "foo bar":
        if true {
            switch -2147483648 {
            case 252:
                var v8 bool = (127 < 96)
                _ = v8
            default:
                var v9 uint64 = 491
                _ = v9
            }
        } else {
            if (((0.0 / ((1e10 / (1e10 * 3.14159)) * -2.25)) / -2.25) == float64(-2.25)) {
                v10 := []int{-2147483648, 687, 127}
                _ = v10
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint16 = 255
                v13 := GenS2{F0: -1, F1: 248, F2: (v12 ^ v12)}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    default:
        var v14 uint64 = 2
        _ = v14
    }
    var v15 bool = false
    v16 := GenS2{F0: -98, F1: 1, F2: 969}
    _ = v15
    _ = v16
    return "foo bar"
}

func main() {
    var v1 int8 = 42
    v2 := []string{"a", ((("foo bar" + ("foo bar" + (("hello" + ("a" + ("\n\t" + ("a" + "\n\t")))) + ""))) + "hello") + "\n\t"), ("\n\t" + (("hello" + "") + "a")), "foo bar"}
    v3 := make(map[string]string)
    v4 := []int{300, -383, 4294967295}
    v5 := make(map[int]int)
    var v6 string = "foo bar"
    v17 := GenS3{F0: 435, F1: 127}
    v1 -= v1
    if (1 != 542) {
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                if true {
                    v20 := v5[fi18]
                    _ = v20
                }
                var v21 int32 = 0
                _ = fi19
                _ = v21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 int8 = (v1 | v1)
                var v24 bool = (!false)
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 910, F1: fi18}
            _ = fi18
            _ = v25
        }
    }
    var v26 uint64 = 785
    v1 -= v1
    switch ("a" + v6) {
    case "hello":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "":
        v28 := GenS4{F0: 1, F1: v6, F2: "\n\t"}
        _ = v28
    default:
        v29 := GenS2{F0: 0, F1: 127, F2: 0}
        _ = v29
    }
    _ = len(v2)
    var v30 uint32 = 363
    var v31 int = 56
    if ((3.14159 - (0.1 - 0.0)) >= 1e10) {
        v6 = v6
        var v32 int32 = 0
        if false {
            var v33 int32 = -38
            if false {
                for fi34 := 0; fi34 < 7; fi34++ {
                    v35 := GenS3{F0: 624, F1: 249}
                    _ = fi34
                    _ = v35
                }
            }
            var v36 uint8 = 198
            _ = v33
            _ = v36
        }
        _ = v32
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint64 = (^(v26))
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
