package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int64) int32 {
    var v7 uint64 = 887
    switch -959 {
    case 32768:
        fmt.Println(v7, v7)
    case -330:
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: 0}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if ((false && true) && ((fi11 % ((fi11) | 1)) < 877)) {
            v7 = v7
            p0 = 588
        } else {
            if false {
                for fi12 := 0; fi12 < 2; fi12++ {
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
            }
            p0 -= (-(p0))
            if true {
                p0 = p0
                var v14 uint32 = 536
                p0 = p0
                _ = v14
            } else {
                var v15 uint8 = 128
                v16 := GenS2{F0: 0, F1: 2, F2: ("\n\t" + ("hello" + (("a" + ("hello" + "a")) + "a")))}
                v17 := []int{}
                _ = v15
                _ = v16
                _ = v17
            }
        }
        _ = fi11
    }
    var v18 uint64 = 1
    _ = p0
    _ = v18
    _ = v7
    return 280
}
func genFn1(p0 bool, p1 uint8) uint32 {
    if p0 {
        v19 := GenS3{F0: 811, F1: 446}
        switch -647 {
        case -1:
            v20 := GenS4{F0: 127}
            _ = v20
        }
        _ = v19
    }
    p1 = (p1 - 250)
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 float64 = (1.5 * ((-2.25 + 0.0) - 0.1))
    var v2 float32 = 3.14159
    var v3 uint32 = 2147483648
    v4 := []string{"foo bar"}
    var v5 uint64 = 4294967295
    var v6 int = 201
    var v21 float64 = (1.5 - v1)
    fmt.Println(v3)
    v1 -= v1
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v22 int16 = 1
        _ = v22
    case "foo bar":
        var v23 uint64 = (v5 & v5)
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := GenS1{F0: 283}
        if (9223372036854775807 < -9223372036854775808) {
            if (!false) {
                v21 -= 0.1
            } else {
                var v26 int16 = -32768
                v3 += (v3 % ((624) | 1))
                _ = v26
            }
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v27 float32 = float32(v1)
                _ = v27
            }
            switch ("" + ("" + "hello")) {
            case "\n\t":
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 uint = 65535
                    _ = fi28
                    _ = v29
                }
            case "hello":
                if true {
                    switch 32768 {
                    case 205:
                        var v30 uint8 = 128
                        _ = v30
                    case -50:
                        v31 := GenS3{F0: 0, F1: uint(v5)}
                        _ = v31
                    default:
                        if true {
                            var v32 uint8 = 78
                            switch ("foo bar" + "") {
                            case "\n\t":
                                switch "a" {
                                case "\n\t":
                                    var v33 int64 = 1
                                    _ = v33
                                default:
                                    v34 := GenS5{F0: v5}
                                    _ = v34
                                }
                            case "hello":
                                var v35 int8 = 61
                                _ = v35
                            case "":
                                var v36 uint = 549
                                _ = v36
                            }
                            _ = v32
                        }
                    }
                    _ = len(v4)
                }
            default:
                if (!(!(v3 <= 1))) {
                    if true {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v1 = v1
                            _ = fi37
                        }
                    }
                }
            }
        } else {
            v38 := GenS2{F0: 2, F1: 664, F2: "foo bar"}
            _ = v38
        }
        var v39 int16 = 0
        _ = fi24
        _ = v25
        _ = v39
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v40 := GenS6{F0: (v5 | (v5 & 583))}
    var v41 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
