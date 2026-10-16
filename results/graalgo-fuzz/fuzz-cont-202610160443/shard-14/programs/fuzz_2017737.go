package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int8; F2 uint64 }

func main() {
    var v1 float64 = (3.14159 - 0.0)
    var v2 int16 = -372
    var v3 uint64 = 495
    var v4 uint64 = v3
    v5 := []int{-225, -2147483648, 32768, -84}
    if (true || (!false)) {
        fmt.Println(v3)
        v3 += v4
    }
    v4 = (v3 | v3)
    switch 550 {
    case 128:
        var v6 int8 = -25
        _ = v6
    case 32768:
        v3 = 4294967295
    }
    _ = len(v5)
    if (32768 < -904) {
        if (2 > 4294967295) {
            v7 := GenS1{F0: 1, F1: v2}
            _ = v7
        } else {
            var v8 uint16 = 266
            v3 -= v4
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 bool = (false || false)
                var v11 float64 = v1
                if (692 <= 42) {
                    switch "hello" {
                    case "a":
                        switch (fi9 * fi9) {
                        case 124:
                            for fi12 := 0; fi12 < 6; fi12++ {
                                if false {
                                    switch -9223372036854775808 {
                                    case 32768:
                                        _ = len(v5)
                                    }
                                } else {
                                    if v10 {
                                        v5 = append(v5, fi9)
                                    } else {
                                        if false {
                                            var v13 uint8 = 47
                                            switch fi12 {
                                            case 42:
                                                fmt.Println(v10, v11, v8)
                                            case 1:
                                                v13 -= (^(v13))
                                            }
                                            v14 := GenS1{F0: (^(v13)), F1: 42}
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v8 -= v8
                                            _ = v13
                                            _ = v14
                                        } else {
                                            _ = len(v5)
                                        }
                                        v15 := GenS2{F0: -45}
                                        _ = v15
                                    }
                                    v16 := GenS1{F0: 42, F1: (^(v2))}
                                    var v17 int64 = -885
                                    _ = v16
                                    _ = v17
                                }
                                v8 = 621
                                _ = fi12
                            }
                        }
                    case "\n\t":
                        v18 := GenS2{F0: 8}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 int8 = -16
                        v1 = (v11 + -2.25)
                        var v21 uint8 = 194
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                } else {
                    if false {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            fmt.Println(fi9)
                            _ = fi22
                        }
                        v23 := GenS2{F0: 100}
                        _ = v23
                    }
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            v4 += (v4 + v3)
            v25 := GenS3{F0: ("" + "hello"), F1: int8(v2), F2: v3}
            _ = fi24
            _ = v25
        }
    } else {
        v5 = append(v5, 2147483647)
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS1{F0: 0, F1: -310}
        var v28 bool = true
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 += 42
    var v29 int8 = 127
    var v30 uint = 1
    switch "foo bar" {
    case "\n\t":
        v29 = 0
    }
    v29 -= 103
    var v31 int16 = v2
    v2 += 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
