package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }
func genFn0() int64 {
    var v4 float64 = ((1.5 / 0.1) / ((-2.25 * 3.14159) + 1e10))
    v5 := GenS1{F0: 127, F1: -102, F2: -109}
    var v6 uint8 = 8
    _ = v4
    _ = v5
    _ = v6
    return 999
}
func genFn1(p0 bool, p1 uint32, p2 int16) uint {
    var v7 int = 2147483647
    p1 = (^(p1))
    switch ("\n\t" + "a") {
    case "a":
        var v8 int64 = 414
        _ = v8
    case "foo bar":
        v9 := []int{42, 32768, v7, (v7 / (((v7 - v7)) | 1))}
        _ = v9
    }
    p1 += 477
    var v10 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return uint(p1)
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 614
    var v3 int32 = 336
    var v11 int16 = 0
    switch "hello" {
    case "":
        v2 = (v2 / ((42) | 1))
    case "\n\t":
        if true {
            _ = len(v1)
            if true {
                var v12 uint32 = v2
                if true {
                    switch 1 {
                    case 32768:
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v2, v2)
                            v13 := GenS2{F0: -2147483648, F1: 331}
                            _ = v13
                        }
                    case 128:
                        var v14 uint64 = 4294967295
                        _ = v14
                    }
                } else {
                    var v15 float32 = 3.14159
                    fmt.Println(v3, v12, v11)
                    _ = v15
                }
                var v16 uint32 = 168
                _ = v12
                _ = v16
            }
        } else {
            v11 = (v11 | 0)
            var v17 float32 = 1e10
            var v18 int8 = 34
            _ = v17
            _ = v18
        }
    default:
        v2 -= 4294967295
    }
    if false {
        v2 = v2
        switch ((("\n\t" + ("hello" + (("a" + ("hello" + "\n\t")) + (("foo bar" + "hello") + "\n\t")))) + ("" + "\n\t")) + "a") {
        case "hello":
            v2 = 40
        case "foo bar":
            var v19 int8 = 25
            _ = v19
        }
    } else {
        fmt.Println(v2, v11, v2)
    }
    var v20 float64 = 3.14159
    if false {
        v1 = append(v1, v2)
        switch -2147483648 {
        case 9223372036854775807:
            for fi21 := 0; fi21 < 18; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := GenS3{F0: (3.14159 != 0.1)}
                    v24 := GenS4{F0: 42}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                var v25 uint16 = 0
                _ = fi21
                _ = v25
            }
        case -1:
            v3 -= (v3 - -1)
        default:
            if (false || (127 == 127)) {
                var v26 uint32 = 167
                _ = v26
            }
        }
    } else {
        var v27 int = 0
        switch -9223372036854775808 {
        case 1:
            v28 := GenS5{F0: (v27 % (((^(v27))) | 1))}
            _ = v28
        }
        v3 -= (v3 >> ((18446744073709551615) & 63))
        _ = v27
    }
    if (!false) {
        v1 = append(v1, 2)
    }
    var v29 int64 = -1
    fmt.Println(v2, v11, v29)
    var v30 uint32 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
}
