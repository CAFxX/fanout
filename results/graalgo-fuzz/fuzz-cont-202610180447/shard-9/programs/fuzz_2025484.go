package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 bool, p1 bool) int16 {
    p1 = (("\n\t" + "hello") < "\n\t")
    v6 := GenS1{F0: 42, F1: 0}
    v7 := GenS2{F0: 9223372036854775807}
    if p0 {
        if (true || false) {
            p1 = false
            p0 = (128 != 9223372036854775807)
        } else {
            p0 = (0 < 928)
            v8 := []int{127}
            v9 := GenS2{F0: -1}
            _ = v8
            _ = v9
        }
        v10 := []int{-191, 32768, -2147483648, 32767}
        v11 := GenS3{F0: 0, F1: 2}
        _ = v10
        _ = v11
    } else {
        p0 = ((25 < 71) || true)
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1000
}
func genFn1() uint64 {
    switch "\n\t" {
    case "foo bar":
        var v12 uint = 0
        _ = v12
    case "hello":
        var v13 int32 = -996
        _ = v13
    case "":
        var v14 string = "foo bar"
        _ = v14
    default:
        var v15 int = 127
        _ = v15
    }
    v16 := []int{-2147483648, 42}
    _ = v16
    return 42
}

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 167
    var v3 float32 = 0.1
    var v4 bool = false
    var v5 int64 = -1
    v17 := []int{int(v1), 0, -281, -543}
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 bool = false
        for fi20 := 0; fi20 < 1; fi20++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v2)
    fmt.Println(v5, v2, v1)
    var v21 uint16 = 65535
    switch -1 {
    case -818:
        var v22 string = "foo bar"
        _ = v22
    }
    switch int(v1) {
    case -602:
        for fi23 := 0; fi23 < 1; fi23++ {
            _ = len(v17)
            v24 := GenS4{F0: v21, F1: v1, F2: -128}
            for fi25 := 0; fi25 < 1; fi25++ {
                if false {
                    var v26 float64 = float64((0.0 / ((1e10 * ((0.0 + v3) + 0.0)) * (3.14159 + 0.1))))
                    _ = len(v17)
                    _ = v26
                }
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    }
    switch int(v2) {
    case 42:
        for fi27 := 0; fi27 < 1; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                switch "\n\t" {
                case "a":
                    v3 -= float32((0.0 + (0.0 + (0.1 / -2.25))))
                case "hello":
                    switch 0 {
                    case 1:
                        if (v4 || (!v4)) {
                            var v29 bool = ("a" > (("hello" + "a") + "\n\t"))
                            if len(v17) > 0 {
                                v17[0] = fi27
                            }
                            if v4 {
                                var v30 string = "a"
                                _ = v30
                            }
                            _ = v29
                        } else {
                            _ = len(v17)
                            v5 += v5
                            v1 -= 18446744073709551615
                        }
                    case 128:
                        switch fi27 {
                        case 128:
                            if (2 > 612) {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            } else {
                                switch "a" {
                                case "a":
                                    v2 = (v2 + uint8(v1))
                                }
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                        default:
                            var v31 int32 = -2147483648
                            _ = v31
                        }
                    }
                }
                _ = fi28
            }
            _ = fi27
        }
    }
    v4 = v4
    v3 = v3
    v21 += 65535
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 13; fi33++ {
            if false {
                v34 := GenS5{F0: 2147483647}
                fmt.Println(v4, v3, v21)
                fmt.Println(fi33, v5)
                _ = v34
            }
            v35 := GenS3{F0: 215, F1: 252}
            _ = fi33
            _ = v35
        }
        v1 = v1
        v36 := GenS1{F0: (v2 << ((65535) & 63)), F1: (fi32 + -9223372036854775808)}
        _ = fi32
        _ = v36
    }
    v3 += v3
    for fi37 := 0; fi37 < 13; fi37++ {
        switch (fi37 ^ fi37) {
        case -9223372036854775808:
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 uint8 = 42
                var v40 float64 = 1e10
                var v41 int = fi38
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        case 853:
            fmt.Println(v4, v5)
        case 0:
            for fi42 := 0; fi42 < 14; fi42++ {
                switch (fi42 | (fi42 * fi37)) {
                case 391:
                    fmt.Println(v3)
                case 9223372036854775807:
                    if len(v17) > 0 {
                        v17[0] = fi37
                    }
                }
                _ = fi42
            }
        }
        var v43 uint64 = (v1 & (v1 >> ((18446744073709551615) & 63)))
        _ = fi37
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
