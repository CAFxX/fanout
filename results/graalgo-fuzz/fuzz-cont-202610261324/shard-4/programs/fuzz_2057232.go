package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 bool, p2 int32) uint {
    for fi7 := 0; fi7 < 20; fi7++ {
        p1 = p1
        p0 = 4294967295
        _ = fi7
    }
    v8 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 121
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{66, 65535}
    v3 := []string{"foo bar", "\n\t", ("foo bar" + ("foo bar" + "a"))}
    v4 := make(map[string]int)
    var v5 uint16 = 2
    var v6 float64 = -2.25
    delete(v1, ("foo bar" + ("\n\t" + "a")))
    var v9 float64 = ((1.5 - v6) * (1e10 / (float64(1e10) + -2.25)))
    switch 1 {
    case -1:
        switch 42 {
        case -74:
            var v10 float64 = v6
            _ = v10
        case -1:
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint16 = v5
                _ = fi11
                _ = v12
            }
        case -105:
            v4["foo bar"] = 32768
        default:
            _ = len(v1)
        }
    case 32768:
        for fi13 := 0; fi13 < 2; fi13++ {
            switch (fi13 >> ((2) & 63)) {
            case 2147483647:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v9 = (1.5 - float64(1e10))
                    _ = fi14
                }
            case 0:
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    _ = len(v4)
                case "hello":
                    var v15 int8 = 100
                    _ = v15
                }
            case 32767:
                fmt.Println(fi13)
            default:
                v16 := GenS1{F0: -1}
                _ = v16
            }
            _ = fi13
        }
    default:
        var v17 int32 = 0
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int = fi18
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS2{F0: 42, F1: true, F2: 42}
            _ = fi20
            _ = v21
        }
        switch 32768 {
        case -2147483648:
            v22 := GenS1{F0: 384}
            _ = v22
        case 712:
            if false {
                for fi23 := 0; fi23 < 15; fi23++ {
                    switch "" {
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        var v24 int = fi18
                        _ = v24
                    }
                    if true {
                        v5 = v5
                        if (0.0 <= 1.5) {
                            if (569 < 145) {
                                fmt.Println(v9)
                                v5 = v5
                            } else {
                                v19 = v19
                                fmt.Println(v19)
                            }
                        }
                        v5 = 1
                    } else {
                        if false {
                            v9 = 0.1
                            var v25 uint16 = v5
                            _ = v25
                        } else {
                            v5 -= v5
                        }
                    }
                    _ = fi23
                }
                fmt.Println(fi18, v6, v5)
            }
        case 588:
            if (v6 <= (v9 / v9)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = fi18
        _ = v19
    }
    v6 = 1e10
    v26 := GenS3{F0: 0}
    _ = len(v3)
    var v27 bool = true
    if false {
        for fi28 := 0; fi28 < 11; fi28++ {
            _ = len(v1)
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 int64 = -9223372036854775808
                switch (^(fi29)) {
                case 128:
                    fmt.Println(fi29)
                case 32768:
                    v6 = v9
                default:
                    var v31 uint64 = 4294967295
                    _ = v31
                }
                if v27 {
                    fmt.Println(fi29, v30)
                } else {
                    v5 = 952
                    v32 := GenS4{F0: "foo bar"}
                    v27 = true
                    v9 += v6
                    fmt.Println(v9, v5, v30)
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    fmt.Println(v9, v27, v5)
    v27 = v27
    if (-32768 >= -828) {
        switch "a" {
        case "hello":
            v33 := GenS1{F0: 42}
            _ = v33
        case "":
            var v34 int16 = -1
            _ = v34
        default:
            v35 := GenS2{F0: -1, F1: v27, F2: -128}
            _ = v35
        }
        v36 := GenS2{F0: 2, F1: (v6 <= 1e10), F2: -128}
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
