package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 string = "a"
    var v2 uint16 = 65535
    var v3 bool = true
    var v4 float32 = 1e10
    if ((v3 && (42 != 42)) && v3) {
        var v5 float64 = 0.0
        v3 = false
        _ = v5
    } else {
        v6 := GenS1{F0: 2, F1: 9223372036854775807, F2: 747}
        _ = v6
    }
    if ((-9223372036854775808 == 780) || (9223372036854775807 < 32767)) {
        v7 := GenS2{F0: 540}
        _ = v7
    }
    v2 = v2
    v3 = v3
    if (42 == 0) {
        if v3 {
            v8 := []int{-783, 128}
            var v9 uint32 = 49
            _ = v8
            _ = v9
        } else {
            switch (("hello" + "foo bar") + v1) {
            case "foo bar":
                v10 := []int{592, -1, 128}
                _ = v10
            }
            v3 = v3
        }
        if (-1 == 1) {
            switch "a" {
            case "":
                v11 := GenS1{F0: 1, F1: 42, F2: 1}
                _ = v11
            case "hello":
                if v3 {
                    fmt.Println(v3, v3, v2)
                } else {
                    if (-921 >= -1) {
                        v12 := GenS3{F0: 20, F1: 1, F2: 65535}
                        v13 := GenS4{F0: 2147483647, F1: (v1 + ""), F2: 42}
                        _ = v12
                        _ = v13
                    }
                }
            }
            var v14 float32 = ((v4 * v4) * -2.25)
            v1 = "\n\t"
            _ = v14
        }
        v2 = (v2 - 542)
    } else {
        v15 := []int{-2147483648, -839}
        if (-1 > 32767) {
            switch (("a" + v1) + (v1 + v1)) {
            case "a":
                if len(v15) > 0 {
                    v15[0] = 9223372036854775807
                }
            case "hello":
                switch "foo bar" {
                case "\n\t":
                    v1 = "\n\t"
                default:
                    if (true && true) {
                        v16 := GenS3{F0: -69, F1: 32767, F2: (v2 >> ((0) & 63))}
                        _ = v16
                    }
                }
            default:
                v1 = v1
            }
        }
        if false {
            if (-1 < 0) {
                v3 = (42 == -89)
                v15 = append(v15, 32768)
            } else {
                var v17 uint64 = 1
                if (v3 || (((v1 + "foo bar") + "a") >= v1)) {
                    v4 += -2.25
                }
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 float32 = v4
                    _ = fi19
                    _ = v20
                }
                if v3 {
                    if v3 {
                        fmt.Println(v1, v2, fi18)
                        fmt.Println(v2, v3)
                        var v21 float32 = 0.0
                        _ = v21
                    }
                    var v22 int32 = 2147483647
                    if v3 {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 int = fi18
                            if len(v15) > 0 {
                                v15[0] = fi23
                            }
                            _ = fi23
                            _ = v24
                        }
                        v1 = v1
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                    _ = v22
                } else {
                    var v25 int = fi18
                    for fi26 := 0; fi26 < 14; fi26++ {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v27 int = 127
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
                var v28 uint16 = 255
                _ = fi18
                _ = v28
            }
            fmt.Println(v2)
        }
        _ = v15
    }
    v3 = v3
    v29 := []int{638, 32767, 127, 437}
    switch "a" {
    case "foo bar":
        if (18446744073709551615 != 807) {
            v2 = 42
            v30 := GenS5{F0: v2}
            _ = v30
        }
    case "\n\t":
        v31 := GenS6{F0: 14, F1: (v2 + v2)}
        _ = v31
    case "":
        if v3 {
            v3 = v3
        }
    }
    var v32 float64 = -2.25
    if true {
        if ((-1 != 2147483647) && true) {
            v33 := GenS3{F0: 42, F1: 32768, F2: v2}
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint16 = 32767
                v29 = append(v29, fi34)
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        v32 = v32
        v36 := GenS4{F0: 1, F1: v1, F2: -1}
        _ = v36
    }
    v29 = append(v29, 42)
    var v37 uint64 = 42
    v1 = v1
    v2 += (v2 + (v2 & (v2 << ((1) & 63))))
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v37
    _ = v4
}
