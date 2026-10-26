package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint = 4294967295
    v2 := []string{}
    var v3 int = 920
    v4 := make(map[int]uint32)
    var v5 bool = (113 >= 65535)
    var v6 int32 = 595
    if v5 {
        var v7 int = (v3 ^ -983)
        v5 = true
        _ = v7
    }
    v1 = v1
    v6 = (v6 | (v6 - v6))
    for fi8 := 0; fi8 < 15; fi8++ {
        v1 -= (v1 * 42)
        var v9 int32 = 1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        switch "\n\t" {
        case "foo bar":
            var v11 uint = v1
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            v1 = v1
            _ = fi12
        }
        v1 += 740
        _ = fi10
    }
    var v13 uint64 = 462
    v5 = false
    var v14 int8 = 0
    if ((v6 ^ (-(v6))) >= -2147483648) {
        var v15 float64 = float64(((-2.25 * 0.1) + -2.25))
        for fi16 := 0; fi16 < 1; fi16++ {
            switch "foo bar" {
            case "a":
                _ = len(v2)
            case "":
                if v5 {
                    v17 := GenS1{F0: (^(v13)), F1: uint16(v13), F2: 255}
                    _ = v17
                } else {
                    var v18 uint = 18446744073709551615
                    var v19 int64 = -375
                    switch -479 {
                    case 259:
                        var v20 float64 = 1.5
                        _ = v20
                    case 2147483647:
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v18 = v1
                            switch (^(v3)) {
                            case 32767:
                                var v22 float64 = v15
                                _ = v22
                            case -701:
                                v23 := GenS1{F0: v13, F1: 115, F2: 128}
                                _ = v23
                            case 1:
                                var v24 uint16 = 846
                                _ = v24
                            }
                            _ = fi21
                        }
                    default:
                        switch 32768 {
                        case -541:
                            if (!true) {
                                var v25 int8 = (v14 - 116)
                                fmt.Println(v13)
                                _ = v25
                            } else {
                                v4[fi16] = 2147483648
                                var v26 uint8 = 189
                                var v27 int64 = v19
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                var v28 int16 = 428
                _ = v28
            }
            _ = len(v4)
            switch "a" {
            case "\n\t":
                fmt.Println(fi16, v3)
            case "a":
                for fi29 := 0; fi29 < 15; fi29++ {
                    v4[(fi16 & 127)] = 1
                    var v30 int = fi16
                    _ = fi29
                    _ = v30
                }
            default:
                var v31 string = ("a" + (("" + (("a" + "foo bar") + ("a" + "foo bar"))) + ("" + "a")))
                _ = v31
            }
            _ = fi16
        }
        _ = v15
    } else {
        v32 := GenS1{F0: (^(v13)), F1: 2, F2: 124}
        if true {
            switch (-(v3)) {
            case 1:
                fmt.Println(v3, v6)
            default:
                if ((v1 - 255) <= v1) {
                    v33 := GenS1{F0: 42, F1: 32767, F2: 128}
                    _ = v33
                }
            }
        }
        v1 = v1
        _ = v32
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 int = int(v14)
        _ = fi34
        _ = v35
    }
    var v36 int8 = v14
    v1 += v1
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS2{F0: int64(v36)}
        _ = fi37
        _ = v38
    }
    var v39 uint = 2
    var v40 bool = true
    var v41 uint8 = 129
    if ((!v40) || v5) {
        _ = len(v2)
    } else {
        fmt.Println(v3, v41, v14)
    }
    v42 := GenS3{F0: -431}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
