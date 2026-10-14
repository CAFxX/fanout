package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint; F1 uint64; F2 bool }
func genFn0(p0 bool) uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: p0, F1: -32768}
        v9 := GenS1{F0: (!p0), F1: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = "foo bar"
    v11 := GenS1{F0: (42 != 0), F1: 715}
    v12 := []int{-415, 128}
    switch v10 {
    case "a":
        switch 2147483647 {
        case 1:
            var v13 uint64 = 18446744073709551615
            _ = v13
        case 114:
            if false {
                switch 9223372036854775807 {
                case 32767:
                    v14 := GenS1{F0: ((false && (p0 && false)) && (-372 != 1000)), F1: -1}
                    _ = v14
                case -670:
                    fmt.Println(p0, v10)
                case -2147483648:
                    var v15 uint8 = 128
                    _ = v15
                default:
                    v16 := GenS1{F0: (564 == 1), F1: -32768}
                    _ = v16
                }
            }
        case 9223372036854775807:
            p0 = p0
        default:
            v10 = "a"
        }
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 42
}
func genFn1() uint8 {
    v17 := GenS1{F0: true, F1: 1}
    var v18 int = -761
    _ = v17
    _ = v18
    return 128
}

func main() {
    var v1 int8 = -1
    var v2 float32 = (0.1 / 1.0)
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    var v6 bool = (true && (42 == 0))
    v19 := []int{}
    fmt.Println(v2, v2)
    if v6 {
        var v20 int8 = (v1 - (^(v1)))
        fmt.Println(v20, v20)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS2{F0: (v1 * (v1 + (v1 - 127))), F1: (v1 >> ((32) & 63)), F2: 791}
            if ((2 > 65535) && v6) {
                var v23 int8 = (v1 - v1)
                v23 += (-(v1))
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 float64 = 1.5
                    v26 := GenS3{F0: "hello", F1: 1}
                    var v27 int32 = 0
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v23
            } else {
                fmt.Println(fi21, v2)
                v1 += v1
                switch "\n\t" {
                case "foo bar":
                    for fi28 := 0; fi28 < 15; fi28++ {
                        v2 += -2.25
                        v2 = ((v2 * v2) + float32((((1.5 + (0.1 - (1.5 - 0.0))) + 1e10) + float64(1.5))))
                        if v6 {
                            v29 := GenS2{F0: (v1 | v1), F1: v1, F2: 489}
                            _ = v29
                        }
                        _ = fi28
                    }
                case "\n\t":
                    var v30 float32 = ((v2 / v2) + (float32(1.5) / 3.14159))
                    _ = v30
                case "hello":
                    var v31 uint32 = 425
                    _ = v31
                default:
                    fmt.Println(v6, v2, v2)
                }
            }
            _ = fi21
            _ = v22
        }
        switch -802 {
        case 0:
            var v32 bool = v6
            _ = v32
        case 22:
            for fi33 := 0; fi33 < 12; fi33++ {
                v34 := GenS2{F0: v1, F1: (v1 - (-(v1))), F2: 2}
                v2 = (1e10 + (1e10 - ((((1.5 * (1.5 + v2)) + 1e10) / v2) / v2)))
                fmt.Println(v2, v6)
                _ = fi33
                _ = v34
            }
        default:
            switch 4294967295 {
            case -113:
                v6 = false
            default:
                var v35 float64 = (0.0 * -2.25)
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 float32 = v2
            fmt.Println(v6, v6, v6)
            _ = fi36
            _ = v37
        }
    }
    var v38 int16 = -786
    v39 := GenS2{F0: -57, F1: -100, F2: 2}
    var v40 uint64 = 42
    v41 := GenS4{F0: ("\n\t" + ""), F1: 0}
    for fi42 := 0; fi42 < 9; fi42++ {
        if ("foo bar" > ("hello" + "hello")) {
            v43 := GenS5{F0: 407, F1: v40, F2: true}
            for fi44 := 0; fi44 < 15; fi44++ {
                switch (fi42 - (fi42 ^ fi42)) {
                case 32767:
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case 2147483647:
                    fmt.Println(fi42)
                default:
                    if v6 {
                        v40 = (v40 / (((^(v40))) | 1))
                        v1 = (-(v1))
                        var v45 int = fi42
                        _ = v45
                    }
                }
                var v46 bool = v6
                var v47 uint32 = 187
                _ = fi44
                _ = v46
                _ = v47
            }
            _ = v43
        }
        var v48 int = (fi42 - fi42)
        fmt.Println(v1, v1)
        _ = fi42
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
