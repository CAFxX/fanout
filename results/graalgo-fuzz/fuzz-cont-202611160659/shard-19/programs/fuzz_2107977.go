package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int64) int64 {
    var v5 bool = true
    var v6 int16 = 1000
    switch "\n\t" {
    case "":
        v7 := GenS1{F0: 409, F1: -2147483648, F2: 18446744073709551615}
        _ = v7
    }
    v8 := []int{232, 4294967295}
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return p0
}
func genFn1(p0 uint8, p1 int8) uint8 {
    switch (("a" + "foo bar") + "foo bar") {
    case "":
        p0 = p0
    case "hello":
        var v9 uint16 = 1
        _ = v9
    }
    p1 += -1
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    var v1 uint = 178
    var v2 uint32 = 65535
    var v3 int64 = -9223372036854775808
    var v4 bool = false
    v10 := GenS1{F0: 127, F1: -2147483648, F2: 0}
    switch int(v3) {
    case 127:
        var v11 float32 = 3.14159
        _ = v11
    case 9223372036854775807:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint32 = v2
            if v4 {
                v14 := []int{-871, -1, fi12}
                switch (fi12 & (fi12 ^ 2147483647)) {
                case 2147483647:
                    v4 = true
                case 32767:
                    v15 := GenS1{F0: -1, F1: 2147483647, F2: 1}
                    _ = v15
                case 100:
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS1{F0: (fi12 / ((2147483647) | 1)), F1: 1, F2: 1}
                        if v4 {
                            fmt.Println(v1)
                            if (false || (1 > 1)) {
                                v18 := GenS1{F0: fi12, F1: 95, F2: 0}
                                var v19 int64 = 304
                                if ((!false) && v4) {
                                    v20 := GenS1{F0: fi12, F1: 1, F2: 1}
                                    fmt.Println(v3, fi16, fi16)
                                    var v21 int = (fi16 << (((v1 >> ((v1) & 63))) & 63))
                                    _ = v20
                                    _ = v21
                                } else {
                                    v22 := GenS2{F0: "foo bar", F1: 127}
                                    fmt.Println(v3, v1)
                                    _ = v22
                                }
                                _ = v18
                                _ = v19
                            } else {
                                var v23 uint = v1
                                if true {
                                    var v24 uint8 = 127
                                    _ = v24
                                } else {
                                    fmt.Println(v2)
                                    var v25 uint32 = v2
                                    _ = v25
                                }
                                _ = v23
                            }
                            var v26 int16 = -822
                            _ = v26
                        }
                        _ = fi16
                        _ = v17
                    }
                default:
                    var v27 uint32 = (v2 / ((0) | 1))
                    _ = v27
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    switch "hello" {
                    case "a":
                        v29 := GenS2{F0: "", F1: 215}
                        _ = v29
                    case "foo bar":
                        var v30 uint8 = 0
                        _ = v30
                    case "\n\t":
                        var v31 uint16 = 816
                        _ = v31
                    }
                    _ = fi28
                }
                _ = v14
            } else {
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 float32 = (1.5 * 1.5)
                    v34 := []int{347, (-(fi32)), fi32}
                    v34 = append(v34, fi32)
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v35 int8 = 100
                fmt.Println(v13, v2, v13)
                _ = v35
            }
            _ = fi12
            _ = v13
        }
    case -1:
        v36 := []int{}
        _ = v36
    }
    var v37 int16 = 32767
    for fi38 := 0; fi38 < 7; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            v40 := []int{-2147483648}
            v41 := GenS3{F0: 378, F1: (v37 / ((-820) | 1))}
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi38
    }
    v42 := []int{42, 1, -421, 224}
    var v43 float64 = 1.5
    v37 = v37
    var v44 int64 = genFn0(64)
    v45 := GenS1{F0: 41, F1: 1, F2: 1}
    _ = len(v42)
    v43 = float64((-2.25 / (3.14159 / (1e10 + 3.14159))))
    for fi46 := 0; fi46 < 9; fi46++ {
        fmt.Println(v43, v1, v1)
        v47 := GenS2{F0: "hello", F1: 128}
        v2 -= 0
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
