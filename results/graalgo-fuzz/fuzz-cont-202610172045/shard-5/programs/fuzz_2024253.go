package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 bool) int32 {
    var v4 float64 = -2.25
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}

func main() {
    v1 := []string{}
    var v2 uint8 = 127
    v3 := []int{}
    var v5 int32 = -110
    switch "hello" {
    case "":
        v2 = (^(v2))
    case "a":
        v2 = v2
    default:
        v2 = (v2 << ((312) & 63))
    }
    var v6 int64 = 0
    fmt.Println(v5, v5)
    if len(v3) > 0 {
        v3[0] = 113
    }
    var v7 int8 = 1
    v6 = (v6 << ((2) & 63))
    v5 = (v5 + 42)
    v6 += v6
    if false {
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                v7 = 127
                fmt.Println(v2, v6)
                _ = fi8
            }
        }
        var v9 int16 = 0
        _ = v9
    }
    var v10 string = ("\n\t" + ("" + ""))
    if ((v7 + v7) == 42) {
        v11 := GenS1{F0: true, F1: 0}
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 -= v5
            _ = fi12
        }
        var v13 bool = false
        _ = v11
        _ = v13
    } else {
        fmt.Println(v10, v2)
    }
    var v14 uint32 = uint32(v2)
    switch (v10 + v10) {
    case "foo bar":
        v15 := GenS2{F0: (-2.25 != 1.5)}
        _ = v15
    case "\n\t":
        v6 += 42
    case "a":
        v2 = v2
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch v10 {
            case "\n\t":
                v18 := GenS1{F0: false, F1: (v7 - v7)}
                _ = v18
            case "":
                fmt.Println(v2, fi17)
            default:
                if false {
                    var v19 uint32 = (^(v14))
                    _ = v19
                } else {
                    switch (fi16 * (fi17 * fi17)) {
                    case -817:
                        _ = len(v1)
                    case -9223372036854775808:
                        v20 := GenS2{F0: false}
                        _ = v20
                    default:
                        var v21 string = ""
                        _ = v21
                    }
                    v22 := GenS1{F0: true, F1: v7}
                    switch "" {
                    case "a":
                        var v23 float64 = (((3.14159 + 0.1) - 1.5) + 1e10)
                        _ = v23
                    case "":
                        if true {
                            if true {
                                switch (v10 + (((((((((v10 + (((v10 + "") + "hello") + ((v10 + (v10 + v10)) + (v10 + v10)))) + (((("a" + v10) + "foo bar") + v10) + ("hello" + v10))) + (v10 + ("a" + v10))) + ("\n\t" + ((v10 + (v10 + v10)) + v10))) + "hello") + v10) + v10) + v10) + "a")) {
                                case "hello":
                                    var v24 int = 32767
                                    _ = v24
                                }
                            }
                            v25 := GenS2{F0: (int64(v14) >= (v6 | (v6 % ((-1) | 1))))}
                            switch (-(fi16)) {
                            case 848:
                                v7 = 127
                            }
                            v10 = (v10 + v10)
                            _ = v25
                        }
                        if len(v3) > 0 {
                            v3[0] = 127
                        }
                    case "\n\t":
                        if ((fi16 > 0) || true) {
                            v26 := GenS1{F0: true, F1: 0}
                            var v27 uint = 1
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS2{F0: false}
                            if (!(v6 <= v6)) {
                                if true {
                                    if (!false) {
                                        v7 -= 0
                                    }
                                    switch v10 {
                                    case "a":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "":
                                        var v29 string = (v10 + "foo bar")
                                        v6 -= v6
                                        _ = v29
                                    }
                                }
                                var v30 uint64 = 42
                                v31 := GenS1{F0: ((v2 >> ((2) & 63)) >= v2), F1: v7}
                                fmt.Println(v5, v14)
                                var v32 int = int(v14)
                                var v33 uint16 = 571
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                v6 += (v6 << ((255) & 63))
                                var v34 bool = ((v7 << ((373) & 63)) >= -39)
                                _ = v34
                            }
                            v2 = v2
                            v6 = (v6 * (v6 / (((-(v6))) | 1)))
                            _ = v28
                        }
                    }
                    _ = v22
                }
            }
            _ = fi17
        }
        fmt.Println(v6, v2, v10)
        _ = fi16
    }
    v35 := GenS2{F0: (!true)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v5
    _ = v6
    _ = v7
}
