package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 float64 = 1.5
    var v3 uint = 0
    v2 += v1
    v1 = 1e10
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            v2 = ((3.14159 * v2) - 1e10)
            _ = fi5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v3, v3, v2)
        var v7 uint64 = uint64(v3)
        v3 += 4294967295
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 2
    if false {
        var v9 uint16 = 42
        _ = v9
    }
    v10 := []int{1}
    v8 = v8
    _ = len(v10)
    for fi11 := 0; fi11 < 7; fi11++ {
        if true {
            v1 -= ((v2 * 3.14159) * 0.1)
        } else {
            fmt.Println(fi11)
        }
        _ = fi11
    }
    v2 = ((0.0 + (v1 - 1.5)) * v2)
    v1 += v2
    if (0.0 >= -2.25) {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int64 = 626
            v14 := GenS1{F0: false}
            for fi15 := 0; fi15 < 7; fi15++ {
                if true {
                    if (v8 > v8) {
                        var v16 uint8 = 107
                        switch (fi12 / ((fi12) | 1)) {
                        case 1:
                            if len(v10) > 0 {
                                v10[0] = fi12
                            }
                        }
                        _ = v16
                    }
                    v2 -= ((float64(1.5) * float64((((3.14159 * (1.5 + 3.14159)) * 0.0) - (1.5 - ((3.14159 - 0.0) * 0.0))))) + ((1e10 * (1e10 * float64(0.1))) - 0.1))
                } else {
                    v17 := GenS1{F0: true}
                    switch (fi12 * 496) {
                    case 9223372036854775807:
                        var v18 int = fi12
                        _ = v18
                    case 32767:
                        v3 = v3
                    }
                    _ = v17
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        if (1 != 0) {
            if false {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v3 = 152
                    _ = fi19
                }
                if (242 >= 65535) {
                    var v20 int8 = -128
                    v21 := GenS2{F0: ("a" + "hello"), F1: "a"}
                    v20 = v20
                    _ = v20
                    _ = v21
                } else {
                    var v22 float64 = v2
                    if true {
                        v1 += (float64(float32(0.0)) * ((((v22 + (float64((0.1 + float32((((v1 + ((((float64(0.1) + 3.14159) / 3.14159) - v22) / v1)) + float64(1.5)) - v1)))) - (v2 - (v22 - (0.1 + 3.14159))))) - (v22 + -2.25)) * (v2 + (v1 / (1e10 + v1)))) / (((0.1 * (v22 / 1e10)) - (v1 - float64(0.1))) + 1.5)))
                        var v23 uint8 = v8
                        var v24 int8 = -128
                        _ = v23
                        _ = v24
                    } else {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        for fi25 := 0; fi25 < 7; fi25++ {
                            fmt.Println(v3)
                            v2 = v1
                            if (((42 > -1) || ((2147483648 == uint32(v8)) && ("a" >= ("\n\t" + ("\n\t" + ("foo bar" + "a")))))) || true) {
                                v8 += v8
                            }
                            _ = fi25
                        }
                    case "a":
                        if false {
                            v26 := GenS1{F0: true}
                            v10 = append(v10, 9223372036854775807)
                            for fi27 := 0; fi27 < 20; fi27++ {
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    var v29 float64 = v1
                                    v30 := GenS2{F0: ("" + "foo bar"), F1: ("hello" + (("hello" + (("hello" + "hello") + "\n\t")) + "foo bar"))}
                                    v10 = append(v10, (-(fi28)))
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi27
                            }
                            _ = v26
                        } else {
                            switch "\n\t" {
                            case "a":
                                for fi31 := 0; fi31 < 2; fi31++ {
                                    v32 := GenS3{F0: 42, F1: 242}
                                    _ = fi31
                                    _ = v32
                                }
                            case "hello":
                                var v33 uint16 = 0
                                _ = v33
                            case "":
                                if true {
                                    v34 := GenS4{F0: 787}
                                    v35 := GenS5{F0: 960, F1: v8}
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 float64 = (0.0 - 1e10)
                                    switch -2147483648 {
                                    case -570:
                                        v37 := GenS5{F0: 936, F1: 238}
                                        _ = v37
                                    case 698:
                                        v2 = (float64(-2.25) * 1e10)
                                    case -814:
                                        switch -1 {
                                        case 0:
                                            var v38 float64 = v22
                                            _ = v38
                                        case 32768:
                                            v8 = (v8 / ((v8) | 1))
                                        default:
                                            v22 -= v36
                                        }
                                    default:
                                        v3 += v3
                                    }
                                    for fi39 := 0; fi39 < 12; fi39++ {
                                        for fi40 := 0; fi40 < 6; fi40++ {
                                            _ = len(v10)
                                            v41 := GenS3{F0: -334, F1: 240}
                                            if true {
                                                var v42 uint8 = 2
                                                v43 := GenS4{F0: 1}
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v44 := GenS6{F0: (1 != -9223372036854775808), F1: 0}
                                                v2 = v22
                                                var v45 int64 = 42
                                                _ = v44
                                                _ = v45
                                            }
                                            v46 := GenS7{F0: 2}
                                            fmt.Println(v8, v2)
                                            _ = fi40
                                            _ = v41
                                            _ = v46
                                        }
                                        _ = fi39
                                    }
                                    _ = v36
                                }
                            }
                            switch (("a" + "\n\t") + "\n\t") {
                            case "hello":
                                if len(v10) > 0 {
                                    v10[0] = 9223372036854775807
                                }
                            }
                        }
                    case "":
                        var v47 int8 = int8(v8)
                        _ = v47
                    default:
                        var v48 uint64 = 0
                        _ = v48
                    }
                    _ = v22
                }
            }
        }
    } else {
        var v49 float32 = 3.14159
        _ = len(v10)
        var v50 int = 32768
        _ = v49
        _ = v50
    }
    var v51 uint16 = 42
    var v52 uint16 = uint16(v8)
    var v53 uint8 = v8
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v8
}
