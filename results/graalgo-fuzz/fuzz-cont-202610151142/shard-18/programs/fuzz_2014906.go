package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int8; F1 uint64; F2 uint8 }
func genFn0() int16 {
    var v4 uint64 = 18446744073709551615
    v4 = v4
    v5 := GenS1{F0: 65535, F1: v4, F2: 356}
    v6 := []int{124, 2147483647, 9223372036854775807}
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch fi8 {
            case 0:
                var v9 uint32 = 2147483648
                _ = v9
            }
            _ = fi8
        }
        var v10 string = "hello"
        _ = fi7
        _ = v10
    }
    _ = v4
    _ = v5
    _ = v6
    return -32768
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = (((1e10 / 0.1) - 1.5) * 0.0)
    var v3 int64 = -172
    if (65535 == 42) {
        v11 := []int{-9223372036854775808, 32767, 757, 127}
        v1 = 290
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if ("a" >= "foo bar") {
            for fi13 := 0; fi13 < 13; fi13++ {
                v2 += (v2 / v2)
                _ = fi13
            }
        }
        v14 := []int{}
        switch "foo bar" {
        case "hello":
            v3 -= (v3 >> ((42) & 63))
        }
        _ = fi12
        _ = v14
    }
    var v15 int = 34
    if (!((!true) && false)) {
        if true {
            var v16 int8 = 127
            v17 := GenS2{F0: 1, F1: "hello"}
            if (v1 > 608) {
                fmt.Println(v15, v15, v1)
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    fmt.Println(v15)
                    switch (^(fi18)) {
                    case -600:
                        var v19 bool = (2147483648 > 2147483648)
                        _ = v19
                    }
                    v20 := GenS1{F0: 937, F1: v1, F2: 0}
                    _ = fi18
                    _ = v20
                }
                v21 := GenS1{F0: 255, F1: (^(v1)), F2: 412}
                _ = v21
            }
            _ = v16
            _ = v17
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                if true {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint16 = 32767
                        var v25 int8 = 1
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                v26 := GenS1{F0: 680, F1: (v1 ^ v1), F2: 237}
                v15 = 1
                _ = fi22
                _ = v26
            }
            var v27 int64 = v3
            var v28 uint64 = (v1 >> ((1) & 63))
            _ = v27
            _ = v28
        }
        if false {
            v29 := GenS3{F0: v1, F1: 1000}
            _ = v29
        }
    } else {
        switch 738 {
        case 2147483647:
            var v30 string = ""
            _ = v30
        case -2147483648:
            fmt.Println(v1, v15)
        }
        if ("\n\t" <= ("a" + "a")) {
            fmt.Println(v15, v2, v3)
            v1 = 998
            var v31 uint8 = 72
            _ = v31
        } else {
            var v32 uint = 400
            if true {
                v33 := GenS3{F0: v1, F1: genFn0()}
                var v34 float64 = v2
                _ = v33
                _ = v34
            } else {
                v32 = v32
                v2 += (v2 / 0.1)
            }
            if true {
                switch (("hello" + "hello") + "hello") {
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        if false {
                            v1 = v1
                        }
                    case "":
                        v35 := []int{0, 470, (v15 % (((v15 + v15)) | 1)), (v15 & v15)}
                        _ = v35
                    case "a":
                        if ((v1 & v1) >= (v1 & (^(v1)))) {
                            var v36 int8 = 100
                            v37 := GenS4{F0: 2, F1: ("foo bar" + "foo bar")}
                            var v38 bool = (2147483647 > 42)
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    default:
                        var v39 float64 = v2
                        _ = v39
                    }
                case "a":
                    var v40 uint32 = 2
                    _ = v40
                case "":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v42 := GenS5{F0: (genFn0() < 659), F1: v15}
                        _ = fi41
                        _ = v42
                    }
                }
                v43 := GenS6{F0: true}
                _ = v43
            } else {
                var v44 int16 = genFn0()
                _ = v44
            }
            _ = v32
        }
    }
    if (-2147483648 < -647) {
        if false {
            for fi45 := 0; fi45 < 20; fi45++ {
                v1 = v1
                var v46 uint16 = 0
                _ = fi45
                _ = v46
            }
            var v47 float32 = 0.1
            var v48 uint8 = 207
            _ = v47
            _ = v48
        }
        var v49 uint = 810
        _ = v49
    } else {
        var v50 int = (^(v15))
        if (("foo bar" + ("" + "\n\t")) <= "hello") {
            var v51 int64 = -170
            switch "foo bar" {
            case "":
                var v52 int64 = v3
                _ = v52
            case "hello":
                v1 += (v1 & 4294967295)
            default:
                for fi53 := 0; fi53 < 7; fi53++ {
                    v15 = (v50 + (-(v15)))
                    _ = fi53
                }
            }
            if (!((-(v15)) != int(v1))) {
                v54 := GenS7{F0: 2, F1: -9223372036854775808}
                switch (v15 * v15) {
                case 32768:
                    switch "a" {
                    case "\n\t":
                        for fi55 := 0; fi55 < 17; fi55++ {
                            v1 += (v1 - (v1 & (v1 << ((2) & 63))))
                            v56 := []int{}
                            _ = fi55
                            _ = v56
                        }
                    case "foo bar":
                        if false {
                            var v57 float64 = v2
                            for fi58 := 0; fi58 < 16; fi58++ {
                                switch fi58 {
                                case 260:
                                    var v59 float32 = (3.14159 / (0.1 - (0.0 / 0.1)))
                                    _ = v59
                                case -2147483648:
                                    v60 := []int{}
                                    _ = v60
                                }
                                v61 := []int{v50}
                                var v62 float64 = (3.14159 / 0.1)
                                _ = fi58
                                _ = v61
                                _ = v62
                            }
                            var v63 float64 = (((v57 / v2) + (v2 + (v2 * (v2 / v2)))) - ((v2 + 3.14159) * 1e10))
                            _ = v57
                            _ = v63
                        }
                    }
                case 32767:
                    var v64 uint32 = 564
                    _ = v64
                case -2147483648:
                    fmt.Println(v50, v50, v51)
                default:
                    switch "a" {
                    case "\n\t":
                        var v65 uint16 = 709
                        _ = v65
                    case "":
                        fmt.Println(v3)
                    default:
                        for fi66 := 0; fi66 < 2; fi66++ {
                            v15 -= -830
                            _ = fi66
                        }
                    }
                }
                _ = v54
            }
            _ = v51
        }
        var v67 string = "foo bar"
        _ = v50
        _ = v67
    }
    v68 := GenS8{F0: -128, F1: (^(v1)), F2: 127}
    var v69 float64 = v2
    var v70 bool = (4294967295 > 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v68
    _ = v69
    _ = v70
}
