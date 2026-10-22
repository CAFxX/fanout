package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint16; F1 bool; F2 int8 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    v1 += float32(-2.25)
    v1 = v1
    var v4 uint32 = 320
    v1 = v1
    if true {
        v5 := []int{42, 4294967295, 127}
        for fi6 := 0; fi6 < 3; fi6++ {
            if (false || ((((true && true) && (false && true)) || false) && (-59 > 100))) {
                v4 -= 59
            } else {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := GenS1{F0: "", F1: 18446744073709551615}
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := GenS2{F0: v4}
                    v4 -= (v4 % (((^(v4))) | 1))
                    var v11 uint32 = v4
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            _ = fi6
        }
        var v12 uint16 = 65535
        _ = v12
        _ = v5
    } else {
        v13 := []int{}
        var v14 uint64 = 0
        _ = v13
        _ = v14
    }
    v15 := []int{127, 9223372036854775807, 32767, 42}
    v16 := GenS3{F0: -1, F1: 32767}
    v17 := GenS4{F0: 891, F1: -1, F2: 0}
    var v18 uint8 = 0
    if (255 <= (v18 | 255)) {
        v15 = append(v15, 4294967295)
        v19 := GenS4{F0: 42, F1: int64(v18), F2: 2147483647}
        _ = v19
    } else {
        if false {
            v18 = 127
            v1 += ((float32(-2.25) - v1) * v1)
        } else {
            _ = len(v15)
        }
        if (1e10 != 0.0) {
            fmt.Println(v4)
            v18 = 115
            if false {
                if (42 >= 608) {
                    v18 = v18
                }
                var v20 uint32 = v4
                _ = v20
            } else {
                for fi21 := 0; fi21 < 15; fi21++ {
                    v18 = v18
                    _ = fi21
                }
            }
        }
    }
    v1 = v1
    switch "foo bar" {
    case "hello":
        v22 := GenS3{F0: -32768, F1: 0}
        _ = v22
    }
    _ = len(v3)
    var v23 int16 = 1
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS5{F0: v4, F1: int64(v18)}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        if (!(!true)) {
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS3{F0: -303, F1: (fi27 / ((fi26) | 1))}
                var v29 int32 = -335
                v30 := GenS4{F0: 2, F1: 42, F2: -2147483648}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        } else {
            if true {
                if true {
                    var v31 int = (fi26 ^ (fi26 * fi26))
                    switch "foo bar" {
                    case "hello":
                        var v32 float32 = 3.14159
                        _ = v32
                    }
                    _ = v31
                } else {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v34 := GenS4{F0: 2, F1: 1, F2: 46}
                        if true {
                            if (1 > 32767) {
                                var v35 int8 = -128
                                v36 := GenS6{F0: 2, F1: (v1 == v1), F2: v35}
                                fmt.Println(v23)
                                _ = v35
                                _ = v36
                            }
                            _ = len(v15)
                        }
                        _ = fi33
                        _ = v34
                    }
                    v23 = (-(v23))
                    var v37 uint = 2
                    _ = v37
                }
            }
        }
        _ = fi26
    }
    v38 := GenS6{F0: 0, F1: false, F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
}
