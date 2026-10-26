package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 uint16, p1 int32, p2 int16) uint16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 4294967295, F1: (p0 << ((255) & 63))}
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    if (!(!true)) {
        p1 = -2147483648
        fmt.Println(p2)
        v7 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch fi8 {
            case 521:
                v9 := []int{127, (fi8 | (fi8 >> ((444) & 63))), -1}
                _ = v9
            case -681:
                p1 += (^(p1))
            }
            _ = fi8
        }
        var v10 uint8 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 % ((p0) | 1))
}
func genFn1(p0 uint8) int8 {
    if true {
        var v11 uint64 = 615
        _ = v11
    }
    switch "hello" {
    case "hello":
        if (65535 == 2) {
            switch "foo bar" {
            case "a":
                var v12 int16 = -1
                _ = v12
            case "\n\t":
                p0 = p0
            }
            p0 = 81
        } else {
            if true {
                p0 = p0
            }
        }
    case "\n\t":
        if (597 <= 14) {
            v13 := []int{}
            p0 += p0
            _ = v13
        }
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int8 = -128
            _ = fi14
            _ = v15
        }
    }
    var v16 float32 = ((1e10 * -2.25) / 0.1)
    p0 = 128
    if false {
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := []int{-9223372036854775808}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi18
                _ = v19
            }
            var v20 int = fi17
            var v21 float32 = v16
            _ = fi17
            _ = v20
            _ = v21
        }
        var v22 uint8 = p0
        _ = v22
    }
    _ = p0
    _ = v16
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 uint16 = 32767
    switch ("foo bar" + (("a" + "a") + ("hello" + "hello"))) {
    case "":
        v4 = genFn0((v4 * (v4 - (v4 - genFn0((v4 << ((2) & 63)), 0, -922)))), 0, 390)
    case "\n\t":
        if false {
            var v23 uint32 = 1
            _ = v23
        }
    default:
        v24 := v2[1]
        _ = v24
    }
    v25 := GenS2{F0: 100, F1: -104}
    v4 -= v4
    var v26 uint8 = 2
    var v27 int64 = 291
    if false {
        fmt.Println(v26, v26, v26)
    }
    switch ("a" + "hello") {
    case "\n\t":
        var v28 uint16 = v4
        _ = v28
    }
    v29 := GenS2{F0: 100, F1: 2147483647}
    v26 -= v26
    for fi30 := 0; fi30 < 15; fi30++ {
        switch "foo bar" {
        case "":
            v27 = v27
        }
        fmt.Println(v4, fi30, fi30)
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v26, fi30, fi30)
            v32 := GenS1{F0: 1, F1: genFn0(v4, 152, 1)}
            fmt.Println(v27, v27)
            var v33 uint = 65535
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    v34 := []int{-2147483648, 2147483647, 42, 2147483647}
    if (!false) {
        switch ("foo bar" + ("" + (("hello" + "a") + "\n\t"))) {
        case "":
            v26 = v26
        case "a":
            v27 = v27
        default:
            if true {
                v35 := GenS1{F0: 18446744073709551615, F1: 42}
                if (!false) {
                    v26 -= v26
                    var v36 int16 = 751
                    for fi37 := 0; fi37 < 2; fi37++ {
                        var v38 string = "hello"
                        var v39 int = 322
                        for fi40 := 0; fi40 < 17; fi40++ {
                            if (false || false) {
                                var v41 bool = false
                                _ = len(v2)
                                fmt.Println(v27)
                                _ = v41
                            }
                            v26 -= v26
                            if true {
                                var v42 int8 = 100
                                fmt.Println(v26, v39)
                                _ = v42
                            }
                            _ = fi40
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v36
                }
                v43 := v3[0]
                _ = v35
                _ = v43
            } else {
                fmt.Println(v27)
            }
        }
        v44 := GenS1{F0: 4294967295, F1: v4}
        _ = v44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        for fi46 := 0; fi46 < 15; fi46++ {
            var v47 int64 = 9223372036854775807
            v48 := GenS2{F0: 6, F1: 488}
            var v49 int8 = 127
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        var v50 uint = 681
        _ = fi45
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
