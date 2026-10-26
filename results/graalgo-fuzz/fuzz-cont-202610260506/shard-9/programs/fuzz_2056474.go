package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    var v1 uint16 = 479
    v2 := make(map[uint32]uint32)
    var v3 bool = (false && false)
    var v4 uint16 = v1
    fmt.Println(v1, v3, v3)
    v1 = v1
    var v5 uint8 = 71
    if (973 == 0) {
        var v6 float64 = (float64((0.1 * -2.25)) / 1.5)
        switch "\n\t" {
        case "hello":
            switch 80 {
            case 42:
                if (!v3) {
                    for fi7 := 0; fi7 < 7; fi7++ {
                        for fi8 := 0; fi8 < 6; fi8++ {
                            v6 = v6
                            switch ("" + "a") {
                            case "hello":
                                switch 93 {
                                case 438:
                                    v2[4294967295] = 4294967295
                                }
                            case "":
                                v9 := []int{(fi8 | (^(fi8)))}
                                _ = v9
                            case "\n\t":
                                if true {
                                    var v10 int = (fi7 ^ 4294967295)
                                    fmt.Println(v5)
                                    var v11 int = (v10 >> ((65535) & 63))
                                    _ = v10
                                    _ = v11
                                }
                            default:
                                if true {
                                    v12 := GenS1{F0: "hello", F1: 4294967295, F2: 128}
                                    v4 += (v1 * v4)
                                    _ = v12
                                }
                            }
                            _ = fi8
                        }
                        var v13 int = -1
                        if v3 {
                            v14 := GenS1{F0: ("" + ""), F1: 4294967295, F2: 127}
                            for fi15 := 0; fi15 < 9; fi15++ {
                                var v16 uint16 = (v4 % (((^(v1))) | 1))
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    var v17 uint = 652
                                    _ = v17
                                case "foo bar":
                                    if v3 {
                                        switch ("" + ("a" + "foo bar")) {
                                        case "":
                                            var v18 uint = 18446744073709551615
                                            _ = v18
                                        case "hello":
                                            fmt.Println(fi7)
                                        case "foo bar":
                                            v13 -= 216
                                        }
                                        fmt.Println(v4)
                                    }
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v3 = false
                                    case "\n\t":
                                        v19 := GenS2{F0: v1, F1: "a", F2: -56}
                                        _ = v19
                                    }
                                }
                                _ = fi15
                                _ = v16
                            }
                            v20 := v2[883]
                            _ = v14
                            _ = v20
                        }
                        _ = fi7
                        _ = v13
                    }
                }
            }
        case "foo bar":
            v6 = (1.5 * v6)
        }
        _ = v6
    } else {
        v21 := v2[4294967295]
        var v22 string = (("\n\t" + "foo bar") + ((("" + "foo bar") + "foo bar") + ""))
        var v23 uint = 5
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v2[2147483647] = 2147483648
        _ = fi24
    }
    v1 -= (v1 * (v1 - (v1 / ((v4) | 1))))
    if (0 < 569) {
        var v25 int8 = 127
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 int = fi26
            v4 = (v1 & (v1 - v4))
            for fi28 := 0; fi28 < 6; fi28++ {
                if v3 {
                    var v29 uint8 = uint8(v4)
                    v25 -= (-(v25))
                    _ = v29
                } else {
                    v1 = (v1 ^ uint16(v5))
                }
                if ((1.5 + 0.0) != 1e10) {
                    v4 = v4
                    fmt.Println(v1, v25)
                } else {
                    var v30 float32 = (0.1 - -2.25)
                    _ = v30
                }
                var v31 float32 = 0.0
                var v32 int16 = 42
                var v33 uint32 = 42
                _ = fi28
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            for fi35 := 0; fi35 < 17; fi35++ {
                v5 -= 2
                if (!(250 <= 528)) {
                    var v36 uint16 = (v4 | v1)
                    _ = v36
                }
                v3 = v3
                _ = fi35
            }
            _ = fi34
        }
        var v37 string = "hello"
        _ = v37
    }
    v38 := GenS2{F0: (v4 & (v4 / ((2) | 1))), F1: "foo bar", F2: -1}
    switch ("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + (("a" + ("" + "foo bar")) + ""))) {
    case "a":
        var v39 float32 = (1e10 * (0.1 * -2.25))
        _ = v39
    case "":
        fmt.Println(v5, v4, v5)
    case "foo bar":
        v4 -= 42
    }
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
