package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 bool; F2 int64 }
func genFn0() uint32 {
    v7 := []int{-831}
    var v8 uint64 = 2
    if true {
        v8 += v8
    } else {
        var v9 uint32 = 2
        _ = v9
    }
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int64, p1 string) uint {
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint8 = 128
        _ = fi10
        _ = v11
    }
    fmt.Println(p1, p1)
    if (42 <= 2) {
        p0 = 42
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int8 = 112
    var v2 uint16 = 32767
    v3 := []uint32{901, 2147483648}
    v4 := []string{""}
    v5 := make(map[int]bool)
    var v6 uint64 = uint64(v2)
    for fi12 := 0; fi12 < 4; fi12++ {
        v2 = (v2 - (v2 & 65535))
        var v13 int16 = 32767
        v14 := GenS1{F0: ((661 > 9223372036854775807) || (true || (3.14159 >= ((((1e10 + (0.0 - 1.5)) * 0.1) * (-2.25 * 1e10)) / float32(0.1)))))}
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 += v1
    fmt.Println(v6, v6)
    fmt.Println(v1)
    if true {
        var v15 int32 = 78
        _ = v15
    }
    v2 += (v2 ^ 2)
    v1 = (v1 ^ v1)
    v1 -= v1
    fmt.Println(v2, v6)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint = 65535
        var v18 uint64 = 18446744073709551615
        var v19 uint32 = 1
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v6 += (v6 / ((v6) | 1))
        if true {
            var v21 int64 = 42
            _ = v21
        } else {
            if (((18446744073709551615 > 0) || true) && ((false && false) && false)) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    switch -1 {
                    case 0:
                        v23 := GenS1{F0: false}
                        _ = v23
                    }
                    _ = len(v3)
                    _ = fi22
                }
                if len(v4) > 0 {
                    v4[0] = ("\n\t" + "\n\t")
                }
            }
            switch "" {
            case "foo bar":
                for fi24 := 0; fi24 < 16; fi24++ {
                    switch (fi20 << ((761) & 63)) {
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v1)
                            fmt.Println(v6)
                        } else {
                            switch 9223372036854775807 {
                            case -1:
                                fmt.Println(fi20)
                            default:
                                switch fi20 {
                                case 69:
                                    var v25 string = "foo bar"
                                    _ = v25
                                case 965:
                                    var v26 uint16 = (v2 & (v2 << ((42) & 63)))
                                    _ = v26
                                case 9223372036854775807:
                                    v2 = (v2 >> ((18446744073709551615) & 63))
                                    var v27 int32 = 0
                                    switch "\n\t" {
                                    case "a":
                                        switch int(v1) {
                                        case 4294967295:
                                            switch "hello" {
                                            case "foo bar":
                                                var v28 int16 = 667
                                                _ = v28
                                            case "a":
                                                var v29 int = (fi20 >> ((genFn1(-9223372036854775808, ("a" + ""))) & 63))
                                                _ = v29
                                            default:
                                                v30 := GenS1{F0: (!true)}
                                                _ = v30
                                            }
                                        case -286:
                                            if ("\n\t" == "") {
                                                fmt.Println(v1, v1, fi20)
                                                if ((1.5 - -2.25) != (1e10 * (0.0 * 1e10))) {
                                                    v31 := GenS2{F0: 4294967295, F1: 596, F2: -261}
                                                    var v32 int32 = v27
                                                    _ = v31
                                                    _ = v32
                                                }
                                            }
                                        case 32768:
                                            v33 := GenS3{F0: 2}
                                            _ = v33
                                        default:
                                            var v34 uint64 = v6
                                            _ = v34
                                        }
                                    case "foo bar":
                                        v35 := GenS4{F0: 42, F1: true, F2: 42}
                                        _ = v35
                                    case "\n\t":
                                        v27 = (v27 ^ -2147483648)
                                    }
                                    _ = v27
                                }
                            }
                        }
                    }
                    switch -690 {
                    case 1:
                        v6 -= 702
                    case -544:
                        _ = len(v5)
                    default:
                        v36 := v5[fi20]
                        _ = v36
                    }
                    _ = fi24
                }
            case "\n\t":
                v37 := GenS4{F0: 655, F1: (800 != (v6 | v6)), F2: -1}
                _ = v37
            }
        }
        switch fi20 {
        case 9223372036854775807:
            v6 -= 42
        case 127:
            v6 = 955
        default:
            switch fi20 {
            case 32768:
                var v38 uint = 255
                _ = v38
            case 9223372036854775807:
                v6 = 598
            default:
                for fi39 := 0; fi39 < 7; fi39++ {
                    v40 := GenS3{F0: 65535}
                    _ = fi39
                    _ = v40
                }
            }
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
