package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 int }
func genFn0(p0 int32, p1 int64, p2 int16) uint64 {
    switch "hello" {
    case "hello":
        p1 -= -9223372036854775808
    default:
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int32 = (p0 % ((p0) | 1))
            v8 := []int{(^(fi6)), fi6, 4294967295, fi6}
            v9 := GenS1{F0: (0 < 243), F1: 111}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: true, F1: 4294967295}
    if true {
        v11 := []int{905, -577}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    v1 := []uint32{0, 328, 4294967295, 1}
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    v12 := GenS1{F0: false, F1: 18446744073709551615}
    var v13 bool = (946 < 42)
    var v14 float64 = 1e10
    switch "\n\t" {
    case "a":
        v15 := GenS1{F0: (1 > -871), F1: 1}
        _ = v15
    case "":
        var v16 uint16 = 1
        _ = v16
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v17 string = (((("hello" + "hello") + "a") + "\n\t") + "")
        _ = v17
    case 9223372036854775807:
        switch ("" + "a") {
        case "a":
            if (0 != 18446744073709551615) {
                var v18 uint8 = 127
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 float64 = (-2.25 / 0.1)
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                v1 = append(v1, 65535)
                v13 = ("foo bar" != ("hello" + "a"))
                v1 = append(v1, 2147483647)
            }
        case "foo bar":
            var v21 string = ""
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 18; fi22++ {
                v13 = (v13 || (65535 <= 65535))
                fmt.Println(fi22)
                v14 = ((((1e10 * (float64(float32(v14)) / -2.25)) * (-2.25 - float64(float32(v14)))) * 3.14159) * (1e10 / v14))
                _ = fi22
            }
        }
    }
    var v23 bool = (1e10 >= 1.5)
    v23 = (false || v23)
    if (!v23) {
        if (1e10 <= 1e10) {
            v14 += (3.14159 * v14)
            for fi24 := 0; fi24 < 16; fi24++ {
                v13 = ((-128 != 39) || (v13 && v23))
                for fi25 := 0; fi25 < 4; fi25++ {
                    delete(v5, fi24)
                    _ = fi25
                }
                _ = fi24
            }
        } else {
            if (9223372036854775807 > 32767) {
                var v26 uint = 824
                if len(v1) > 0 {
                    v1[0] = 219
                }
                var v27 uint = v26
                _ = v26
                _ = v27
            } else {
                v28 := GenS1{F0: (("a" + "hello") == "a"), F1: 545}
                _ = v28
            }
            v23 = (v14 >= float64(float32(v14)))
        }
    } else {
        var v29 int16 = 515
        v29 -= (v29 ^ v29)
        _ = v29
    }
    v30 := GenS1{F0: true, F1: 187}
    var v31 int64 = 42
    switch 421 {
    case 32768:
        v32 := v2["hello"]
        _ = v32
    }
    if (1 < 1) {
        v33 := GenS2{F0: (0 < 1), F1: "", F2: 0}
        fmt.Println(v31, v31)
        _ = v33
    }
    fmt.Println(v14)
    var v34 int16 = -32768
    v35 := GenS3{F0: 127}
    v31 += (v31 % ((42) | 1))
    v36 := GenS3{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
