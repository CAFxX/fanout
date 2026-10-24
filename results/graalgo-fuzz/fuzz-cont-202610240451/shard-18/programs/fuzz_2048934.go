package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 bool; F1 string; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
func genFn0() string {
    v4 := GenS1{F0: 2, F1: 9223372036854775807, F2: "a"}
    for fi5 := 0; fi5 < 1; fi5++ {
        switch (fi5 ^ (fi5 & (fi5 >> ((268) & 63)))) {
        case 954:
            switch "\n\t" {
            case "foo bar":
                var v6 float32 = 1e10
                _ = v6
            case "a":
                for fi7 := 0; fi7 < 1; fi7++ {
                    var v8 uint8 = 0
                    switch fi7 {
                    case -343:
                        fmt.Println(fi5)
                    case 128:
                        v9 := []int{(fi7 >> ((255) & 63))}
                        _ = v9
                    case -126:
                        fmt.Println(v8)
                    default:
                        v8 += 0
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v8 += (v8 | v8)
                    case "a":
                        fmt.Println(v8)
                    default:
                        var v10 int16 = -311
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
            }
        case -750:
            var v11 float64 = float64(3.14159)
            _ = v11
        default:
            fmt.Println(fi5)
        }
        switch "a" {
        case "\n\t":
            switch (fi5 ^ 43) {
            case 0:
                var v12 uint16 = 32767
                _ = v12
            }
        case "a":
            fmt.Println(fi5)
        case "hello":
            var v13 int64 = -827
            _ = v13
        default:
            fmt.Println(fi5)
        }
        fmt.Println(fi5)
        _ = fi5
    }
    switch 836 {
    case 9223372036854775807:
        var v14 uint32 = 1
        _ = v14
    case 0:
        var v15 uint32 = 65535
        _ = v15
    }
    var v16 uint64 = 2
    v17 := []int{-612}
    _ = v16
    _ = v17
    _ = v4
    return ("" + "\n\t")
}
func genFn1(p0 uint8, p1 int, p2 bool) uint64 {
    p0 += p0
    v18 := []int{}
    if (p2 && p2) {
        switch "\n\t" {
        case "":
            if p2 {
                var v19 int = (-(p1))
                p2 = true
                for fi20 := 0; fi20 < 10; fi20++ {
                    if p2 {
                        fmt.Println(p1, fi20, p0)
                        v18 = append(v18, p1)
                        var v21 int16 = 1000
                        _ = v21
                    } else {
                        var v22 uint = 541
                        fmt.Println(p0, fi20, p0)
                        _ = v22
                    }
                    _ = fi20
                }
                _ = v19
            } else {
                switch (p1 << ((1) & 63)) {
                case 128:
                    for fi23 := 0; fi23 < 10; fi23++ {
                        if (!p2) {
                            fmt.Println(p1, p2)
                        }
                        _ = fi23
                    }
                }
            }
        case "\n\t":
            for fi24 := 0; fi24 < 10; fi24++ {
                for fi25 := 0; fi25 < 20; fi25++ {
                    if len(v18) > 0 {
                        v18[0] = p1
                    }
                    _ = fi25
                }
                fmt.Println(p0)
                switch 32768 {
                case -9223372036854775808:
                    var v26 uint8 = 255
                    _ = v26
                case -2147483648:
                    if p2 {
                        p0 += (p0 % (((^(p0))) | 1))
                    } else {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            var v28 int32 = -1
                            if (65535 != uint32(p0)) {
                                if len(v18) > 0 {
                                    v18[0] = (fi27 + (fi27 % ((233) | 1)))
                                }
                            } else {
                                if p2 {
                                    var v29 int64 = -1
                                    _ = v29
                                } else {
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    p1 -= (fi27 / ((fi24) | 1))
                                    v30 := GenS2{F0: (1 < 65535), F1: ("foo bar" + "a"), F2: "\n\t"}
                                    _ = v30
                                }
                            }
                            p0 -= p0
                            if true {
                                p1 -= (fi27 + (p1 + fi27))
                                switch "foo bar" {
                                case "foo bar":
                                    v31 := GenS1{F0: 442, F1: 297, F2: "a"}
                                    _ = v31
                                }
                            } else {
                                var v32 float64 = 1.5
                                v33 := GenS3{F0: 127}
                                switch ("" + "") {
                                case "foo bar":
                                    v28 = (v28 | v28)
                                case "":
                                    p1 += (p1 % ((fi27) | 1))
                                }
                                v34 := GenS4{F0: -398, F1: 0}
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(p1, fi24, fi27)
                            var v35 uint64 = 42
                            v35 -= (v35 + v35)
                            var v36 uint8 = p0
                            _ = fi27
                            _ = v28
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS5{F0: p2, F1: 255, F2: 122}
                        _ = v37
                    }
                }
                _ = fi24
            }
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    p1 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 2
}

func main() {
    var v1 float64 = float64(((1.5 - float32(((1.5 * float64(1e10)) + (0.1 - 1e10)))) / 1.0))
    v2 := []string{"\n\t", "foo bar", "\n\t"}
    var v3 int8 = 42
    v1 = v1
    var v38 uint = 2
    v39 := GenS6{F0: false}
    if true {
        if ((439 <= 32767) || false) {
            var v40 float64 = ((0.0 / v1) + v1)
            _ = v40
        } else {
            for fi41 := 0; fi41 < 10; fi41++ {
                v42 := GenS4{F0: (fi41 % ((127) | 1)), F1: (v3 >> ((v38) & 63))}
                var v43 float64 = v1
                v2 = append(v2, ("hello" + (("hello" + "foo bar") + (("" + "\n\t") + ""))))
                _ = fi41
                _ = v42
                _ = v43
            }
            var v44 int = 953
            v1 = (float64((1.5 - float32((v1 + (v1 * -2.25))))) - 1.5)
            _ = v44
        }
    } else {
        var v45 float64 = 1e10
        var v46 float64 = -2.25
        v1 = 0.0
        _ = v45
        _ = v46
    }
    if (!(false || true)) {
        fmt.Println(v1)
        var v47 bool = true
        _ = v47
    }
    v38 = (v38 * (v38 + 286))
    v38 = v38
    if (0.1 <= 0.0) {
        switch 32768 {
        case 32767:
            if ((3.14159 - float32(v1)) < 1e10) {
                fmt.Println(v1, v3)
            }
        default:
            v48 := GenS3{F0: 1}
            _ = v48
        }
        v1 -= v1
    }
    v38 = 2
    _ = len(v2)
    v49 := GenS6{F0: (0.1 < 0.0)}
    var v50 int32 = 286
    v1 = v1
    v51 := GenS7{F0: (v50 <= (v50 % ((v50) | 1)))}
    v1 -= -2.25
    var v52 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
