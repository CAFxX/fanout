package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint; F1 string }
func genFn0() int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint8 = 255
        var v8 int8 = -45
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int = fi9
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        var v11 int = -1
        _ = v11
    }
    return 264
}

func main() {
    var v1 int8 = 42
    v2 := []int{406}
    v3 := []uint32{}
    v4 := make(map[string]string)
    var v5 int64 = -257
    var v12 int = 521
    if true {
        if (42 != 2147483648) {
            switch v12 {
            case 32768:
                var v13 uint = 4294967295
                _ = v13
            }
            v14 := v4["a"]
            _ = v14
        } else {
            var v15 uint64 = 184
            v12 -= (v12 / (((v12 ^ (v12 | (v12 & 32767)))) | 1))
            v16 := GenS1{F0: (v1 & v1), F1: (("\n\t" + "\n\t") + "foo bar"), F2: v12}
            _ = v15
            _ = v16
        }
        if (-1 <= v5) {
            if len(v2) > 0 {
                v2[0] = v12
            }
            if (4294967295 < 18446744073709551615) {
                switch "a" {
                case "foo bar":
                    v1 += -1
                default:
                    if true {
                        var v17 int64 = (v5 * (v5 | (v5 ^ (v5 | (v5 & -9223372036854775808)))))
                        _ = v17
                    }
                }
            }
            var v18 uint32 = 265
            _ = v18
        }
    }
    var v19 uint32 = 2
    v20 := GenS2{F0: -442, F1: 42}
    var v21 int32 = 42
    if (!(!(true || (0 < v21)))) {
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(fi22, v1)
            var v23 int8 = 0
            switch ((("\n\t" + "hello") + "") + "hello") {
            case "\n\t":
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 bool = false
                    _ = fi24
                    _ = v25
                }
            case "hello":
                var v26 uint64 = 1
                _ = v26
            case "":
                if true {
                    var v27 uint64 = 0
                    var v28 bool = ((fi22 | (fi22 >> ((0) & 63))) <= v12)
                    v23 = -44
                    _ = v27
                    _ = v28
                }
            default:
                if (!true) {
                    var v29 uint32 = 4294967295
                    v23 = 105
                    var v30 int = v12
                    _ = v29
                    _ = v30
                } else {
                    switch v12 {
                    case -1:
                        v31 := GenS3{F0: 1, F1: 1}
                        _ = v31
                    }
                    v12 -= (v12 % (((fi22 - (v12 / ((v12) | 1)))) | 1))
                    v21 = -887
                }
            }
            _ = fi22
            _ = v23
        }
        v21 = (v21 & v21)
    }
    fmt.Println(v5, v12)
    fmt.Println(v21, v1)
    v32 := GenS4{F0: 1, F1: ""}
    v19 += 885
    v33 := GenS2{F0: -370, F1: (v5 % ((v5) | 1))}
    var v34 int64 = v5
    if (true || (!(127 == 20))) {
        var v35 int16 = 1000
        _ = v35
    } else {
        v36 := GenS2{F0: -697, F1: v5}
        v2 = append(v2, v12)
        _ = v36
    }
    fmt.Println(v1, v12, v21)
    v37 := GenS2{F0: genFn0(), F1: v5}
    v1 = v1
    v38 := GenS1{F0: (^(v1)), F1: "\n\t", F2: v12}
    delete(v4, "")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
