package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 int64) uint8 {
    fmt.Println(p1)
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    return 41
}
func genFn1() uint {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 float32 = 0.1
            v9 = v9
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int = (fi10 % ((fi7) | 1))
            var v12 uint32 = 1
            v12 = (v12 >> ((2) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint = 65535
        _ = fi7
        _ = v13
    }
    var v14 int16 = 42
    var v15 uint = 255
    switch ("a" + ("" + ("foo bar" + "hello"))) {
    case "hello":
        v14 += v14
    case "\n\t":
        if (!false) {
            var v16 int8 = 1
            v17 := GenS1{F0: (((0 < -271) || true) || false)}
            _ = v16
            _ = v17
        } else {
            var v18 float32 = (float32(3.14159) + (float32(1.5) * (-2.25 + 3.14159)))
            if (2147483647 <= -2147483648) {
                v19 := []int{}
                switch 209 {
                case 32768:
                    fmt.Println(v14)
                case 2147483647:
                    if (2147483647 < -2147483648) {
                        v18 += v18
                    } else {
                        var v20 uint8 = 1
                        switch (("hello" + "foo bar") + ((("foo bar" + "a") + ("\n\t" + "a")) + (("hello" + "hello") + ""))) {
                        case "\n\t":
                            var v21 float32 = v18
                            _ = v21
                        case "hello":
                            if true {
                                if len(v19) > 0 {
                                    v19[0] = 32768
                                }
                                v22 := GenS1{F0: false}
                                v20 += (v20 & v20)
                                _ = v22
                            } else {
                                var v23 int32 = -2147483648
                                var v24 int8 = -1
                                v24 += v24
                                _ = v23
                                _ = v24
                            }
                        case "a":
                            for fi25 := 0; fi25 < 3; fi25++ {
                                v26 := GenS2{F0: v15}
                                v15 = v15
                                v27 := GenS3{F0: 42, F1: 42}
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        default:
                            v15 = 2
                        }
                        _ = v20
                    }
                case 933:
                    var v28 int8 = 1
                    _ = v28
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    if len(v19) > 0 {
                        v19[0] = 128
                    }
                    var v30 int = fi29
                    var v31 uint = (v15 >> ((v15) & 63))
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v19
            } else {
                if false {
                    fmt.Println(v15, v15, v15)
                    v18 = v18
                    v32 := GenS3{F0: 42, F1: 1}
                    _ = v32
                }
            }
            var v33 uint16 = 1
            _ = v18
            _ = v33
        }
    case "a":
        switch "a" {
        case "foo bar":
            for fi34 := 0; fi34 < 12; fi34++ {
                v15 = (v15 * 0)
                _ = fi34
            }
        case "a":
            for fi35 := 0; fi35 < 18; fi35++ {
                v14 = (v14 & v14)
                _ = fi35
            }
        case "\n\t":
            v14 -= v14
        }
    }
    if true {
        v14 -= v14
    }
    _ = v14
    _ = v15
    return (v15 / (((v15 % ((4294967295) | 1))) | 1))
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 100
    var v3 int32 = -708
    var v4 uint8 = 255
    var v5 int = 0
    var v6 float32 = 0.0
    v36 := []int{(v5 + (v5 & 32768)), v5, 2147483647}
    if (!false) {
        for fi37 := 0; fi37 < 1; fi37++ {
            v1 -= (v1 >> ((848) & 63))
            _ = fi37
        }
        switch v5 {
        case 2147483647:
            if true {
                var v38 int32 = (v3 - 2147483647)
                var v39 int8 = (v2 ^ 1)
                _ = v38
                _ = v39
            }
        case 0:
            v6 = (float32((3.14159 * -2.25)) - float32(0.1))
        default:
            v3 += v3
        }
    } else {
        var v40 int64 = v1
        switch (^(v5)) {
        case -351:
            if ((v6 / v6) != ((v6 - ((1e10 - 3.14159) * v6)) / (v6 * v6))) {
                switch (v5 >> ((18446744073709551615) & 63)) {
                case 2147483647:
                    var v41 float64 = 0.0
                    _ = v41
                case 32768:
                    v42 := GenS4{F0: int8(v1), F1: (v4 + v4), F2: 907}
                    _ = v42
                case -1:
                    _ = len(v36)
                default:
                    var v43 bool = (!true)
                    _ = v43
                }
            } else {
                v6 += float32(0.0)
            }
        }
        _ = v40
    }
    switch (-(v5)) {
    case 9223372036854775807:
        if len(v36) > 0 {
            v36[0] = -9223372036854775808
        }
    case -7:
        var v44 int = -1
        _ = v44
    default:
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
    }
    v36 = append(v36, v5)
    _ = len(v36)
    var v45 uint8 = v4
    var v46 uint16 = 255
    var v47 string = ("foo bar" + ("foo bar" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
