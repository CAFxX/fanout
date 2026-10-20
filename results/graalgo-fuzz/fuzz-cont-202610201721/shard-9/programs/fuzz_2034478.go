package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 string) int8 {
    var v5 float64 = 0.1
    v6 := GenS1{F0: -86}
    if (p0 != "foo bar") {
        v5 += v5
    } else {
        var v7 uint16 = 65535
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 = ""
            v7 = v7
            var v9 int16 = 0
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return -71
}

func main() {
    v1 := []int{-9223372036854775808, 127, 0}
    var v2 int32 = 2147483647
    var v3 uint16 = 1
    v4 := make(map[int]bool)
    v2 -= (v2 << ((18446744073709551615) & 63))
    v10 := GenS2{F0: -9223372036854775808}
    v11 := GenS2{F0: 128}
    for fi12 := 0; fi12 < 1; fi12++ {
        delete(v4, 817)
        _ = fi12
    }
    switch "foo bar" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    }
    v13 := GenS2{F0: -1}
    fmt.Println(v3)
    v3 -= (v3 & 65535)
    var v14 int32 = (^(v2))
    v1 = append(v1, 4294967295)
    switch "\n\t" {
    case "":
        for fi15 := 0; fi15 < 17; fi15++ {
            if (0 < 42) {
                fmt.Println(v2)
                _ = len(v4)
                if ((!true) && true) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            fmt.Println(v2)
                        case "hello":
                            if (3.14159 >= 1.5) {
                                var v17 int = 0
                                var v18 uint8 = 113
                                _ = v17
                                _ = v18
                            }
                        }
                        v2 = (v14 | (v14 << ((42) & 63)))
                        _ = fi16
                    }
                    if true {
                        if ((0.0 / (1.5 / (0.1 * float64(1.5)))) == (float64((((1.5 + -2.25) - 1.5) + -2.25)) - 3.14159)) {
                            var v19 int64 = -932
                            _ = v19
                        }
                        var v20 uint16 = 2
                        fmt.Println(v3, v20, fi15)
                        _ = v20
                    } else {
                        v2 -= (v14 - v14)
                    }
                }
            } else {
                var v21 float32 = 0.1
                if (42 == 4294967295) {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 int8 = -128
                        _ = len(v4)
                        if false {
                            if (927 <= 4294967295) {
                                var v24 int64 = 699
                                v1 = append(v1, fi15)
                                switch "a" {
                                case "\n\t":
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v25 int64 = (v24 & 104)
                                        v1 = append(v1, -641)
                                        _ = v25
                                    }
                                default:
                                    var v26 int = (fi22 | fi22)
                                    v27 := GenS2{F0: fi15}
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                            } else {
                                fmt.Println(v3, v3)
                            }
                            if (v3 == (v3 ^ v3)) {
                                var v28 int64 = 42
                                fmt.Println(fi22)
                                _ = v28
                            }
                            var v29 string = "foo bar"
                            var v30 int32 = (v14 << ((239) & 63))
                            _ = v29
                            _ = v30
                        }
                        _ = len(v4)
                        _ = len(v4)
                        _ = fi22
                        _ = v23
                    }
                }
                _ = v21
            }
            if false {
                var v31 int8 = 17
                if true {
                    fmt.Println(fi15, v2, v3)
                    delete(v4, fi15)
                } else {
                    _ = len(v1)
                    v3 = (^(v3))
                }
                _ = v31
            } else {
                if true {
                    var v32 float32 = -2.25
                    _ = v32
                } else {
                    v33 := GenS3{F0: 359, F1: false, F2: 4294967295}
                    var v34 string = ""
                    if false {
                        var v35 float64 = float64(0.0)
                        var v36 uint16 = 0
                        v36 += 0
                        _ = v35
                        _ = v36
                    }
                    _ = v33
                    _ = v34
                }
            }
            var v37 float32 = float32(1e10)
            _ = fi15
            _ = v37
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
