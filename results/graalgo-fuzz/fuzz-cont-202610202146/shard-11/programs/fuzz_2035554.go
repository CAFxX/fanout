package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32; F1 int8; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS6 struct { F0 bool; F1 string; F2 uint32 }
func genFn0() uint64 {
    v6 := []int{-1, 42, -926, 127}
    var v7 uint8 = 1
    if false {
        v8 := GenS1{F0: (("" + "hello") + "\n\t"), F1: 126, F2: -9223372036854775808}
        var v9 string = ""
        v10 := GenS2{F0: 18446744073709551615, F1: 1}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v7 = v7
        if false {
            var v12 int = (fi11 >> ((2) & 63))
            switch "" {
            case "a":
                if false {
                    _ = len(v6)
                } else {
                    _ = len(v6)
                    switch 547 {
                    case 2147483647:
                        v13 := GenS1{F0: "\n\t", F1: -1, F2: 42}
                        _ = v13
                    case -2147483648:
                        _ = len(v6)
                    default:
                        var v14 int16 = -479
                        _ = v14
                    }
                    v12 -= v12
                }
            case "":
                v7 = v7
            case "\n\t":
                if false {
                    var v15 uint16 = 255
                    var v16 uint16 = (v15 - 236)
                    v15 -= (v15 & 32767)
                    _ = v15
                    _ = v16
                } else {
                    switch "hello" {
                    case "hello":
                        switch ("" + "a") {
                        case "foo bar":
                            if false {
                                v12 -= (fi11 - 0)
                                switch "" {
                                case "hello":
                                    v7 = (^(v7))
                                case "":
                                    var v17 string = "\n\t"
                                    _ = v17
                                }
                                var v18 uint64 = 2
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    switch (fi19 * fi19) {
                                    case 950:
                                        switch 639 {
                                        case 128:
                                            fmt.Println(fi19)
                                        case 363:
                                            if true {
                                                switch 9223372036854775807 {
                                                case 0:
                                                    v12 = (fi19 - 32768)
                                                }
                                            }
                                            fmt.Println(v7, v7)
                                        case -9223372036854775808:
                                            if (true && (42 > 858)) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v7 += 128
                                                case "a":
                                                    var v20 int8 = int8(v7)
                                                    _ = v20
                                                case "foo bar":
                                                    v12 -= fi11
                                                default:
                                                    fmt.Println(v12, v7, fi11)
                                                    switch (v12 >> ((255) & 63)) {
                                                    case -9223372036854775808:
                                                        var v21 int = 32768
                                                        _ = v21
                                                    case -2147483648:
                                                        var v22 float64 = 0.1
                                                        v22 = ((v22 - float64(1.5)) - v22)
                                                        _ = len(v6)
                                                        _ = v22
                                                    default:
                                                        switch (-(v12)) {
                                                        case -9223372036854775808:
                                                            if true {
                                                                v23 := GenS3{F0: 2147483647, F1: 42, F2: ((!false) || false)}
                                                                v24 := GenS3{F0: -2147483648, F1: 14, F2: (!false)}
                                                                var v25 uint32 = 65535
                                                                _ = v23
                                                                _ = v24
                                                                _ = v25
                                                            } else {
                                                                v12 = -1
                                                                v12 += fi11
                                                            }
                                                        case 30:
                                                            var v26 uint16 = 42
                                                            _ = v26
                                                        case -1:
                                                            v7 += (^(v7))
                                                        }
                                                    }
                                                }
                                                v7 = v7
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v7 = (v7 | 0)
                                            }
                                        default:
                                            var v27 string = (("\n\t" + "") + ("hello" + ""))
                                            _ = v27
                                        }
                                    case -9223372036854775808:
                                        if false {
                                            v28 := GenS4{F0: 127}
                                            fmt.Println(fi19)
                                            var v29 uint16 = 5
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                    var v30 int = (v12 + -313)
                                    _ = fi19
                                    _ = v30
                                }
                                fmt.Println(v12, fi11, v7)
                                for fi31 := 0; fi31 < 1; fi31++ {
                                    if len(v6) > 0 {
                                        v6[0] = 118
                                    }
                                    v32 := GenS5{F0: v7, F1: 0, F2: 1}
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        case "a":
                            v12 += v12
                        default:
                            v12 = (v12 & v12)
                        }
                    }
                    v33 := GenS6{F0: ((true && false) || (true && true)), F1: "a", F2: 2147483648}
                    _ = v33
                }
            default:
                for fi34 := 0; fi34 < 9; fi34++ {
                    v7 = v7
                    if true {
                        fmt.Println(fi34, fi11, v7)
                        fmt.Println(fi11)
                        v7 -= v7
                    }
                    var v35 int8 = -1
                    _ = fi34
                    _ = v35
                }
            }
            _ = v12
        }
        _ = fi11
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8) int16 {
    var v36 int16 = 1000
    v37 := []int{}
    for fi38 := 0; fi38 < 9; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 int8 = 100
            if len(v37) > 0 {
                v37[0] = fi39
            }
            _ = fi39
            _ = v40
        }
        if false {
            var v41 string = ("\n\t" + "a")
            _ = v41
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            _ = len(v37)
            _ = fi42
        }
        _ = fi38
    }
    _ = p0
    _ = v36
    _ = v37
    return v36
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-1, -652}
    var v3 int32 = -2147483648
    var v4 uint64 = 18446744073709551615
    v5 := []string{}
    switch "\n\t" {
    case "a":
        v3 = 1
    case "foo bar":
        v43 := v1[725]
        _ = v43
    default:
        v4 += v4
    }
    switch 2147483647 {
    case -291:
        v44 := GenS6{F0: true, F1: "a", F2: 192}
        _ = v44
    case 32768:
        v4 = (^(v4))
    default:
        v45 := GenS3{F0: 2147483647, F1: 19, F2: ((-2.25 / (float32((1.5 + 3.14159)) + ((float32(-2.25) + (0.0 - 0.0)) + ((1.5 - 1.5) * float32(1.5))))) == 1e10)}
        _ = v45
    }
    v4 += (v4 | v4)
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int8 = 126
        _ = fi46
        _ = v47
    }
    fmt.Println(v4, v4)
    v3 += (v3 << ((979) & 63))
    if len(v5) > 0 {
        v5[0] = ("\n\t" + "foo bar")
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        if false {
            for fi48 := 0; fi48 < 13; fi48++ {
                var v49 bool = (true && ((fi48 >> ((4294967295) & 63)) <= (fi48 >> ((1) & 63))))
                _ = fi48
                _ = v49
            }
            _ = len(v1)
        }
    case "hello":
        var v50 uint16 = 42
        _ = v50
    case "a":
        fmt.Println(v3, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
