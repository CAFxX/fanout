package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483648, 0, 0}
    v1 += v1
    if false {
        var v4 bool = false
        var v5 uint8 = 0
        _ = v4
        _ = v5
    } else {
        var v6 int32 = v1
        _ = v6
    }
    fmt.Println(v1)
    switch -124 {
    case 32767:
        fmt.Println(v1)
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 -= v1
            _ = fi8
        }
        v9 := GenS1{F0: 0, F1: "\n\t", F2: -557}
        _ = fi7
        _ = v9
    }
    var v10 int8 = 100
    var v11 int8 = v10
    for fi12 := 0; fi12 < 16; fi12++ {
        switch ("" + "foo bar") {
        case "hello":
            var v13 bool = (179 <= 0)
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint16 = 32767
                if true {
                    var v16 int16 = 1000
                    fmt.Println(fi12)
                    v17 := GenS1{F0: (^(v15)), F1: "\n\t", F2: -1}
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: 52, F1: 708}
                var v19 int64 = -124
                v1 = (v1 / ((v1) | 1))
                _ = fi14
                _ = v15
                _ = v18
                _ = v19
            }
        case "\n\t":
            switch "hello" {
            case "foo bar":
                _ = len(v3)
            case "":
                v11 += (v11 * -128)
            case "hello":
                fmt.Println(v1)
            default:
                v1 += v1
            }
        }
        _ = fi12
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := v2[2]
        _ = fi20
        _ = v21
    }
    if false {
        var v22 uint = 841
        for fi23 := 0; fi23 < 13; fi23++ {
            if (false && false) {
                var v24 int32 = v1
                v25 := GenS1{F0: 239, F1: (("hello" + "\n\t") + "hello"), F2: 0}
                if (("hello" + "") == "") {
                    var v26 bool = false
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(fi23)
                    default:
                        var v27 uint = v22
                        _ = v27
                    }
                    v28 := GenS2{F0: 184, F1: 0}
                    _ = v26
                    _ = v28
                }
                _ = v24
                _ = v25
            }
            _ = fi23
        }
        _ = v22
    } else {
        if true {
            switch 2147483647 {
            case -2147483648:
                if false {
                    var v29 uint8 = 255
                    for fi30 := 0; fi30 < 7; fi30++ {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v10 -= (-(v11))
                            var v32 float64 = -2.25
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 2; fi33++ {
                            var v34 float64 = 1e10
                            v35 := v2[4294967295]
                            v34 -= v34
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS2{F0: 42, F1: 42}
                        _ = fi30
                        _ = v36
                    }
                    var v37 bool = true
                    _ = v29
                    _ = v37
                }
            case 1:
                v11 += v10
            }
            v38 := GenS1{F0: 2, F1: "a", F2: -835}
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 float64 = 0.0
                _ = fi39
                _ = v40
            }
            _ = v38
        }
    }
    var v41 int = 42
    switch ("a" + (((("a" + "\n\t") + ("a" + ("foo bar" + "hello"))) + ("a" + (("a" + "") + ("\n\t" + ("" + "\n\t"))))) + "foo bar")) {
    case "hello":
        fmt.Println(v10)
    case "a":
        var v42 uint16 = 193
        _ = v42
    case "foo bar":
        v1 = (^(v1))
    default:
        var v43 uint8 = 181
        _ = v43
    }
    fmt.Println(v41, v41)
    v44 := GenS2{F0: 128, F1: 0}
    var v45 string = "\n\t"
    if len(v3) > 0 {
        v3[0] = 712
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        v41 -= (fi46 << ((255) & 63))
        v10 = 127
        v41 -= v41
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v41
    _ = v44
    _ = v45
}
