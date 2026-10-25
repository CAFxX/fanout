package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 int32 }
func genFn0(p0 int8, p1 int, p2 bool) int16 {
    if p2 {
        p2 = p2
        for fi7 := 0; fi7 < 7; fi7++ {
            p1 -= 128
            var v8 uint64 = 42
            v9 := GenS1{F0: -434, F1: -802}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 string = ("hello" + "foo bar")
        _ = v10
    } else {
        switch "a" {
        case "hello":
            switch ("" + ((("\n\t" + "a") + "hello") + "")) {
            case "\n\t":
                p0 = (^(p0))
            default:
                var v11 int32 = 1
                _ = v11
            }
        default:
            if (!(1 < -126)) {
                v12 := []int{}
                _ = v12
            }
        }
    }
    var v13 float32 = ((1.5 - -2.25) - 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}
func genFn1(p0 int16, p1 int16) int {
    var v14 int64 = 9223372036854775807
    if ("hello" != "a") {
        p1 -= 42
        v15 := []int{}
        _ = v15
    }
    var v16 uint16 = 1
    if (!(!false)) {
        var v17 int32 = 42
        if true {
            v18 := GenS2{F0: v16}
            for fi19 := 0; fi19 < 4; fi19++ {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 uint64 = 725
                    var v22 int8 = 42
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS2{F0: v16}
                v24 := []int{}
                _ = fi19
                _ = v23
                _ = v24
            }
            _ = v18
        }
        _ = v17
    } else {
        fmt.Println(v14)
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v14)
            _ = fi25
        }
    }
    if (!false) {
        for fi26 := 0; fi26 < 11; fi26++ {
            if false {
                fmt.Println(v16, p1, p0)
                switch (fi26 * fi26) {
                case 32768:
                    p1 += -216
                case 127:
                    v14 = v14
                default:
                    v14 += (v14 % ((-82) | 1))
                }
                if ((65535 == 690) && false) {
                    v16 += 65535
                    fmt.Println(v14, fi26)
                    p1 = -32768
                } else {
                    v27 := GenS1{F0: 42, F1: 42}
                    if (int64(p0) < (v14 & (^(v14)))) {
                        v28 := []int{}
                        var v29 float32 = 0.0
                        if (!true) {
                            var v30 float64 = (0.0 - (((0.0 + 1e10) * 1.5) / float64(v29)))
                            var v31 uint = 1
                            v32 := GenS2{F0: 12}
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS2{F0: (v16 & v16)}
                            v29 -= 0.1
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                fmt.Println(v29, v16)
                            case "hello":
                                fmt.Println(fi26, v16, p0)
                            case "a":
                                v34 := GenS3{F0: 18446744073709551615, F1: 2}
                                _ = v34
                            }
                            _ = v33
                        }
                        _ = v28
                        _ = v29
                    }
                    var v35 int16 = p0
                    _ = v27
                    _ = v35
                }
            } else {
                var v36 uint64 = 0
                _ = v36
            }
            _ = fi26
        }
        var v37 float32 = (1e10 + 1e10)
        _ = v37
    } else {
        for fi38 := 0; fi38 < 16; fi38++ {
            switch 9223372036854775807 {
            case -2147483648:
                fmt.Println(v14, p1)
            case 0:
                var v39 uint64 = 0
                _ = v39
            }
            _ = fi38
        }
        var v40 int16 = (p0 << ((294) & 63))
        _ = v40
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v16
    return 2147483647
}

func main() {
    var v1 uint64 = 951
    v2 := []int{32767, -576, -9223372036854775808, -1}
    v3 := make(map[uint32]uint32)
    var v4 int8 = -88
    v5 := make(map[string]int)
    var v6 int = 1
    v41 := v3[1]
    if false {
        for fi42 := 0; fi42 < 19; fi42++ {
            _ = len(v3)
            _ = fi42
        }
        v43 := GenS1{F0: 0, F1: 11}
        v6 = (-(v6))
        _ = v43
    }
    switch v6 {
    case 2147483647:
        v44 := GenS4{F0: 366, F1: 0, F2: (("a" + ("foo bar" + "")) + "")}
        _ = v44
    case -1:
        v45 := GenS5{F0: -1, F1: 2147483647, F2: 1}
        _ = v45
    default:
        fmt.Println(v4, v4)
    }
    v46 := GenS2{F0: 255}
    fmt.Println(v6)
    var v47 string = ("" + "hello")
    for fi48 := 0; fi48 < 8; fi48++ {
        var v49 uint32 = v41
        var v50 int64 = 1
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        var v52 uint16 = 1
        v53 := GenS4{F0: (v52 ^ 70), F1: 192, F2: "a"}
        _ = fi51
        _ = v52
        _ = v53
    }
    var v54 uint16 = 0
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v47
    _ = v5
    _ = v54
    _ = v6
}
