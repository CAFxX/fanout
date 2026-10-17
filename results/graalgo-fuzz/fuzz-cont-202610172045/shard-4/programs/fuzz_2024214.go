package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int; F1 uint; F2 int64 }
type GenS5 struct { F0 int16; F1 uint; F2 string }
func genFn0() bool {
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint = 0
        v7 := []int{}
        v4 = (v4 ^ 255)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return (!false)
}
func genFn1(p0 bool, p1 string) uint32 {
    p0 = (212 > 167)
    for fi8 := 0; fi8 < 8; fi8++ {
        if p0 {
            fmt.Println(p1, fi8, p1)
            var v9 uint16 = 2
            _ = v9
        }
        var v10 uint = 2
        v11 := []int{(^(fi8)), fi8, (fi8 >> ((65535) & 63)), 516}
        _ = fi8
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{}
        if ((!((386 > 130) || (42 <= -32768))) || false) {
            for fi14 := 0; fi14 < 7; fi14++ {
                if ((fi12 * fi14) > fi12) {
                    switch fi12 {
                    case 32768:
                        p1 = ("hello" + "\n\t")
                    case 127:
                        v15 := GenS1{F0: -1, F1: 60}
                        var v16 int32 = 0
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = v15
                        _ = v16
                    }
                }
                fmt.Println(p0)
                v17 := GenS2{F0: (269 == 1)}
                _ = fi14
                _ = v17
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                switch (((("hello" + p1) + (p1 + "\n\t")) + "foo bar") + p1) {
                case "\n\t":
                    _ = len(v13)
                }
                p1 = p1
                if (171 > 2147483647) {
                    if len(v13) > 0 {
                        v13[0] = fi12
                    }
                } else {
                    var v19 uint32 = 42
                    v13 = append(v13, (fi18 % ((182) | 1)))
                    _ = v19
                }
                _ = fi18
            }
        } else {
            if (1 < 1) {
                var v20 int8 = 1
                _ = len(v13)
                _ = v20
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: -1, F1: 643}
                    p1 = (p1 + (p1 + ((p1 + p1) + "a")))
                    _ = fi21
                    _ = v22
                }
                p0 = (32767 != -1)
            }
            v23 := GenS3{F0: 42, F1: 42}
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            _ = len(v13)
            if (-292 != -1) {
                p1 = p1
                if false {
                    var v25 float64 = 1e10
                    v26 := GenS4{F0: (fi12 ^ 42), F1: 42, F2: -9223372036854775808}
                    switch (((p1 + p1) + "hello") + "hello") {
                    case "hello":
                        if true {
                            if len(v13) > 0 {
                                v13[0] = -85
                            }
                            p0 = p0
                        }
                    case "":
                        var v27 string = "a"
                        _ = v27
                    default:
                        fmt.Println(p1, v25)
                    }
                    if false {
                        var v28 int64 = -9223372036854775808
                        v28 -= (v28 << ((4294967295) & 63))
                        if true {
                            var v29 uint32 = 2147483647
                            if len(v13) > 0 {
                                v13[0] = 128
                            }
                            var v30 uint8 = 2
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(v25, v25, v25)
                            v31 := GenS3{F0: -1, F1: 255}
                            _ = v31
                        }
                        _ = v28
                    } else {
                        p0 = (!(200 == 9223372036854775807))
                    }
                    _ = v25
                    _ = v26
                }
            } else {
                p0 = (255 > 128)
                var v32 int8 = 42
                _ = v32
            }
            p1 = p1
            _ = fi24
        }
        _ = fi12
        _ = v13
    }
    var v33 string = p1
    p0 = p0
    _ = p0
    _ = p1
    _ = v33
    return 2147483647
}

func main() {
    v1 := []int{-261}
    var v2 int8 = 113
    var v3 uint = 2
    _ = len(v1)
    v2 = v2
    for fi34 := 0; fi34 < 10; fi34++ {
        v2 = 113
        _ = fi34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v36 := GenS3{F0: v2, F1: 127}
        v37 := GenS5{F0: 1, F1: v3, F2: "foo bar"}
        v2 = 100
        _ = fi35
        _ = v36
        _ = v37
    }
    if len(v1) > 0 {
        v1[0] = 602
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 uint32 = 107
        fmt.Println(fi38, v39)
        _ = fi38
        _ = v39
    }
    v2 += v2
    if false {
        v2 += v2
        v2 = v2
        v3 -= 536
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 float32 = (1.5 - ((((1.5 / -2.25) / (0.1 / 1.0)) / 1e10) / 1.5))
        for fi42 := 0; fi42 < 13; fi42++ {
            fmt.Println(v2, fi40, v41)
            switch (("hello" + "") + (("" + (("\n\t" + (("hello" + "hello") + "hello")) + "foo bar")) + "hello")) {
            case "hello":
                var v43 uint16 = 42
                _ = v43
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v41 += v41
            default:
                var v44 int = fi42
                _ = v44
            }
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    if len(v1) > 0 {
        v1[0] = -140
    }
    v45 := GenS1{F0: v2, F1: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
}
