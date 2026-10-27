package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0() uint64 {
    if false {
        var v6 int = 975
        if (true && false) {
            v6 = -133
            v6 = (-(v6))
        } else {
            switch v6 {
            case 4294967295:
                v6 -= (v6 - v6)
            case 9223372036854775807:
                if true {
                    switch (^(v6)) {
                    case -2147483648:
                        switch ("hello" + "foo bar") {
                        case "":
                            if (true || ((-9223372036854775808 == -1) || true)) {
                                for fi7 := 0; fi7 < 10; fi7++ {
                                    v8 := []int{(fi7 << ((1) & 63)), (v6 & (-(v6))), (v6 ^ (^(v6)))}
                                    _ = fi7
                                    _ = v8
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v9 := []int{(v6 ^ (v6 & (v6 << ((65535) & 63))))}
                                    _ = v9
                                case "a":
                                    v10 := GenS1{F0: -608}
                                    _ = v10
                                }
                            }
                        case "\n\t":
                            if true {
                                var v11 uint64 = 1
                                if false {
                                    var v12 float64 = 1e10
                                    switch (("" + "foo bar") + "a") {
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v6 = (v6 + (v6 | (v6 / (((v6 | v6)) | 1))))
                                        case "a":
                                            fmt.Println(v6, v6, v6)
                                        case "\n\t":
                                            if (958 <= 65535) {
                                                for fi13 := 0; fi13 < 1; fi13++ {
                                                    var v14 string = (("" + (("" + "") + "")) + ("foo bar" + ("foo bar" + ("\n\t" + ""))))
                                                    _ = fi13
                                                    _ = v14
                                                }
                                                v15 := []int{2147483647, (v6 * v6)}
                                                _ = v15
                                            } else {
                                                v16 := []int{0}
                                                _ = v16
                                            }
                                        default:
                                            fmt.Println(v6, v6)
                                        }
                                    case "":
                                        var v17 uint = 18446744073709551615
                                        _ = v17
                                    case "a":
                                        fmt.Println(v6, v11)
                                    default:
                                        switch v6 {
                                        case -393:
                                            var v18 float64 = (0.1 + ((v12 * 0.1) + float64((0.0 + -2.25))))
                                            _ = v18
                                        default:
                                            v19 := []int{v6, 32768}
                                            _ = v19
                                        }
                                    }
                                    _ = v12
                                }
                                _ = v11
                            }
                        }
                    case 32767:
                        var v20 uint8 = 128
                        _ = v20
                    case -1:
                        v21 := GenS1{F0: 1000}
                        _ = v21
                    }
                    v6 = (v6 | 32768)
                }
            case -2147483648:
                fmt.Println(v6)
            default:
                for fi22 := 0; fi22 < 7; fi22++ {
                    if (!(!false)) {
                        v6 -= (v6 % (((fi22 ^ (^(fi22)))) | 1))
                    }
                    v6 = v6
                    fmt.Println(fi22)
                    _ = fi22
                }
            }
            var v23 float64 = 3.14159
            v24 := []int{v6, (v6 / ((v6) | 1)), (^(v6))}
            _ = v23
            _ = v24
        }
        _ = v6
    }
    if false {
        for fi25 := 0; fi25 < 8; fi25++ {
            if true {
                switch ("a" + "foo bar") {
                case "a":
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v27 := GenS2{F0: 4294967295, F1: fi26}
                        var v28 int = (-(fi26))
                        var v29 uint8 = 2
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                default:
                    if false {
                        v30 := GenS1{F0: 272}
                        _ = v30
                    }
                }
                var v31 int = 836
                var v32 int8 = -120
                _ = v31
                _ = v32
            }
            if false {
                var v33 int64 = -1
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 int = (fi34 >> ((1) & 63))
                    v36 := GenS3{F0: (fi34 & 32767), F1: -128}
                    var v37 int64 = v33
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                var v38 float32 = (float32((0.0 + 0.0)) / 1.5)
                _ = v33
                _ = v38
            }
            _ = fi25
        }
    }
    var v39 uint16 = 482
    _ = v39
    return 0
}

func main() {
    var v1 bool = (!(!false))
    var v2 int32 = 0
    var v3 uint = 4294967295
    var v4 uint = (v3 | (v3 & v3))
    var v5 int64 = -9223372036854775808
    if v1 {
        fmt.Println(v2, v1)
    } else {
        fmt.Println(v1, v5, v5)
        v40 := GenS2{F0: uint32(v3), F1: -2147483648}
        _ = v40
    }
    fmt.Println(v5)
    fmt.Println(v1, v4)
    v41 := GenS4{F0: (v5 & v5)}
    v42 := GenS4{F0: 9223372036854775807}
    if v1 {
        switch "hello" {
        case "a":
            v43 := GenS5{F0: 65535, F1: 1}
            _ = v43
        case "":
            v44 := GenS2{F0: 4294967295, F1: 112}
            _ = v44
        case "\n\t":
            switch 0 {
            case -17:
                v5 = 783
            default:
                var v45 int32 = -1
                _ = v45
            }
        }
        var v46 float64 = float64(1.5)
        v47 := []int{-2147483648}
        _ = v46
        _ = v47
    } else {
        v48 := GenS4{F0: (v5 * 9223372036854775807)}
        v49 := []int{4294967295, -703}
        _ = v48
        _ = v49
    }
    var v50 int8 = 127
    for fi51 := 0; fi51 < 16; fi51++ {
        if v1 {
            if (!v1) {
                fmt.Println(v4)
            }
        }
        var v52 int64 = int64(fi51)
        _ = fi51
        _ = v52
    }
    v1 = v1
    if (2 == 1) {
        v2 -= int32(v50)
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v50
}
