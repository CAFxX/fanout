package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint8 }
func genFn0() bool {
    switch "" {
    case "hello":
        var v7 uint = 4294967295
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: "hello", F1: 195}
        _ = v8
    case "a":
        switch "" {
        case "a":
            var v9 int = -9223372036854775808
            _ = v9
        default:
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint16 = 32767
                _ = fi10
                _ = v11
            }
        }
    default:
        v12 := GenS1{F0: ("\n\t" + "foo bar"), F1: 255}
        _ = v12
    }
    var v13 int16 = 327
    v14 := GenS1{F0: "a", F1: 0}
    v15 := []int{9223372036854775807}
    _ = v13
    _ = v14
    _ = v15
    return true
}
func genFn1(p0 int8, p1 int, p2 int32) int {
    switch "\n\t" {
    case "":
        var v16 int64 = 42
        _ = v16
    case "\n\t":
        switch "hello" {
        case "":
            v17 := []int{}
            _ = v17
        case "hello":
            if (100 >= p0) {
                p1 += p1
            }
        default:
            switch (p1 * (p1 - int(p0))) {
            case 2147483647:
                v18 := GenS1{F0: "hello", F1: 881}
                _ = v18
            case -404:
                for fi19 := 0; fi19 < 12; fi19++ {
                    if true {
                        p0 -= (p0 * p0)
                        var v20 bool = (18446744073709551615 < 112)
                        _ = v20
                    }
                    var v21 uint8 = 127
                    if true {
                        v21 -= (v21 + 2)
                        switch (p1 * fi19) {
                        case -516:
                            p0 -= (^(p0))
                        case 0:
                            for fi22 := 0; fi22 < 7; fi22++ {
                                fmt.Println(p0, v21)
                                _ = fi22
                            }
                        }
                    }
                    _ = fi19
                    _ = v21
                }
            }
        }
    case "a":
        var v23 uint16 = 65535
        _ = v23
    }
    v24 := []int{127}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return (p1 | (p1 | p1))
}

func main() {
    var v1 string = (("\n\t" + ("a" + ("hello" + "hello"))) + ("foo bar" + ((("\n\t" + "") + ("hello" + "\n\t")) + "foo bar")))
    var v2 float64 = -2.25
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    var v5 uint64 = 194
    var v6 float64 = (float64((float32(v2) * ((1e10 / 3.14159) * 0.1))) - (v2 + ((((v2 * (((0.1 + 0.0) - v2) * v2)) * v2) / (float64(-2.25) * (v2 - float64(-2.25)))) + 0.0)))
    var v25 int16 = 32767
    v26 := GenS1{F0: "", F1: 166}
    v1 = "a"
    for fi27 := 0; fi27 < 4; fi27++ {
        switch (-(fi27)) {
        case 1:
            v28 := GenS2{F0: 127}
            _ = v28
        }
        _ = fi27
    }
    v29 := []int{-1}
    v4 = (v4 & v4)
    v5 -= 2
    var v30 int = 643
    switch (v30 >> ((255) & 63)) {
    case 0:
        switch (v1 + v1) {
        case "a":
            v31 := GenS2{F0: 42}
            _ = v31
        default:
            var v32 uint32 = 4294967295
            _ = v32
        }
    }
    if (!(431 == 994)) {
        for fi33 := 0; fi33 < 12; fi33++ {
            v25 += 1000
            _ = fi33
        }
        if false {
            switch v1 {
            case "foo bar":
                for fi34 := 0; fi34 < 20; fi34++ {
                    if (((v1 + v1) <= (((v1 + v1) + v1) + v1)) || false) {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            var v36 int = 32767
                            _ = fi35
                            _ = v36
                        }
                    }
                    for fi37 := 0; fi37 < 6; fi37++ {
                        fmt.Println(v2, fi34, fi37)
                        if len(v29) > 0 {
                            v29[0] = v30
                        }
                        _ = len(v3)
                        _ = fi37
                    }
                    _ = fi34
                }
            case "":
                switch (v30 << ((108) & 63)) {
                case -673:
                    fmt.Println(v25, v4, v30)
                default:
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
            case "a":
                var v38 string = ((((v1 + "\n\t") + v1) + (v1 + "")) + "a")
                _ = v38
            default:
                if (!(2147483647 < 469)) {
                    fmt.Println(v5)
                    for fi39 := 0; fi39 < 6; fi39++ {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            if len(v29) > 0 {
                                v29[0] = (fi40 << ((4294967295) & 63))
                            }
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            v41 := GenS3{F0: fi39, F1: 4294967295, F2: -2147483648}
                            _ = fi40
                            _ = v41
                        }
                        _ = fi39
                    }
                } else {
                    v25 = -477
                }
            }
            _ = len(v29)
            var v42 int32 = 0
            _ = v42
        }
        v43 := GenS3{F0: v30, F1: 2147483648, F2: 42}
        _ = v43
    }
    var v44 int32 = int32(v5)
    fmt.Println(v2, v44)
    var v45 uint64 = 42
    var v46 string = v1
    if (((v45 - (v45 & 2)) == (^(v45))) && true) {
        if true {
            v47 := GenS4{F0: false, F1: v5}
            if true {
                var v48 uint64 = v5
                _ = v48
            }
            _ = v47
        } else {
            if (true || true) {
                v49 := GenS5{F0: (v25 >> ((65535) & 63)), F1: 0}
                fmt.Println(v4, v30, v4)
                _ = v49
            }
            if (946 >= 42) {
                var v50 uint64 = (^(v5))
                v46 = v1
                var v51 uint64 = v5
                _ = v50
                _ = v51
            } else {
                switch v1 {
                case "a":
                    var v52 uint64 = 42
                    _ = v52
                case "":
                    switch -2147483648 {
                    case 2147483647:
                        var v53 uint = 764
                        _ = v53
                    case 942:
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                    case -520:
                        if false {
                            var v54 string = v46
                            v54 = ((((v54 + v54) + "a") + "foo bar") + v54)
                            _ = v54
                        }
                    default:
                        v55 := GenS2{F0: 0}
                        _ = v55
                    }
                case "\n\t":
                    if true {
                        v3[(v30 << ((255) & 63))] = 2147483647
                    }
                default:
                    var v56 int64 = 594
                    _ = v56
                }
                fmt.Println(v45, v45, v25)
            }
            v46 = v46
        }
    } else {
        v4 = (v4 - 0)
        for fi57 := 0; fi57 < 15; fi57++ {
            v58 := GenS3{F0: v30, F1: 285, F2: v44}
            var v59 bool = true
            var v60 uint64 = v5
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
        fmt.Println(v46, v25)
    }
    v61 := GenS1{F0: v1, F1: 0}
    var v62 int64 = 1
    v25 += v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
