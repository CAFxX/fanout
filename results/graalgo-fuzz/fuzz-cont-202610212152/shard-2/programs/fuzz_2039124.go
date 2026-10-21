package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32, p1 uint64) uint8 {
    v5 := []int{}
    p0 += p0
    if true {
        v6 := GenS1{F0: (^(p0))}
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 string = "foo bar"
            p1 -= p1
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 float64 = (float64(((1.5 + 1e10) * float32(1.5))) - float64(0.0))
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 1
}

func main() {
    v1 := []int{-599, 2147483647, 4294967295}
    var v2 uint8 = 66
    var v3 bool = (2 == 65535)
    var v4 int8 = 100
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = v2
        if (3.14159 < float64((1e10 / 0.1))) {
            switch "a" {
            case "\n\t":
                fmt.Println(v2)
            }
        } else {
            v11 := GenS1{F0: 310}
            for fi12 := 0; fi12 < 1; fi12++ {
                switch (^(fi10)) {
                case -9223372036854775808:
                    v1 = append(v1, fi12)
                case 128:
                    var v13 int32 = -2147483648
                    _ = v13
                case -238:
                    var v14 uint = 4294967295
                    _ = v14
                default:
                    var v15 float64 = 1.5
                    _ = v15
                }
                v16 := GenS1{F0: 42}
                if v3 {
                    switch "" {
                    case "":
                        v17 := GenS1{F0: -352}
                        _ = v17
                    case "foo bar":
                        v4 += v4
                    default:
                        if false {
                            if false {
                                v18 := GenS1{F0: 0}
                                if (1 >= 255) {
                                    var v19 int64 = -9223372036854775808
                                    v2 = (v2 * genFn0(202, 1))
                                    v4 -= v4
                                    _ = v19
                                } else {
                                    v1 = append(v1, (fi10 >> ((0) & 63)))
                                    v20 := GenS1{F0: 13}
                                    v2 -= (v2 & (v2 + (v2 * (v2 >> ((759) & 63)))))
                                    v2 = 2
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                v3 = (585 <= 1)
                                var v21 uint8 = v2
                                _ = len(v1)
                                _ = v21
                            }
                            switch fi10 {
                            case 127:
                                v3 = ((0 < 127) || v3)
                            case 4294967295:
                                v22 := GenS1{F0: 520}
                                _ = v22
                            case -572:
                                v23 := GenS1{F0: 2147483647}
                                _ = v23
                            }
                            if len(v1) > 0 {
                                v1[0] = fi10
                            }
                        }
                    }
                    var v24 bool = true
                    var v25 int32 = 42
                    _ = v24
                    _ = v25
                } else {
                    v2 = (v2 % ((2) | 1))
                    v3 = v3
                    if false {
                        _ = len(v1)
                        fmt.Println(v2, fi12)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v26 uint8 = v2
                        fmt.Println(v4, fi10, v4)
                        var v27 int16 = 260
                        if v3 {
                            var v28 int8 = -128
                            fmt.Println(v2, v2)
                            _ = v28
                        } else {
                            var v29 uint8 = (v2 - (v26 >> ((42) & 63)))
                            var v30 uint16 = 2
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS1{F0: 1}
                        _ = v26
                        _ = v27
                        _ = v31
                    } else {
                        fmt.Println(v2, fi12)
                    }
                }
                var v32 int = -9223372036854775808
                v33 := GenS2{F0: 776, F1: 1, F2: 0}
                _ = fi12
                _ = v16
                _ = v32
                _ = v33
            }
            _ = v11
        }
        _ = fi10
    }
    var v34 float64 = 0.1
    v35 := GenS1{F0: -1}
    fmt.Println(v3, v34, v3)
    v36 := GenS1{F0: 0}
    switch "foo bar" {
    case "a":
        switch -894 {
        case -505:
            var v37 int32 = 42
            _ = v37
        }
    case "":
        if (("foo bar" <= "a") || ((int64(v2) >= 0) && v3)) {
            switch "a" {
            case "a":
                v34 += v34
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if (("foo bar" + "\n\t") <= "hello") {
                    var v38 uint64 = 42
                    v34 = v34
                    var v39 float32 = 0.1
                    _ = v38
                    _ = v39
                }
            }
        } else {
            v40 := GenS3{F0: 42}
            _ = v40
        }
    case "hello":
        v4 += (v4 - (-(v4)))
    }
    v2 = (v2 - genFn0(-1, 4294967295))
    _ = len(v1)
    var v41 uint64 = 201
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
}
