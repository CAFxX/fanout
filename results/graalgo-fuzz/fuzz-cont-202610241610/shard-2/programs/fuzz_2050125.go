package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 bool; F1 int8; F2 int }
type GenS5 struct { F0 uint64; F1 uint16; F2 int64 }
func genFn0() int16 {
    v7 := []int{-517}
    if ((((1.5 + 1e10) / 3.14159) + float64((((((1e10 + (3.14159 / 1.0)) + ((float32((((1e10 + (0.1 / 1.5)) / 1e10) * 3.14159)) - (((3.14159 - 0.1) - float32((float64(3.14159) * 3.14159))) + 0.0)) * 1.5)) * 3.14159) / 3.14159) / 1.0))) <= float64(((3.14159 + -2.25) - 1.5))) {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 string = "foo bar"
            fmt.Println(fi8, fi8)
            var v10 uint8 = 255
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 float64 = 1.5
    _ = v11
    _ = v7
    return 1
}
func genFn1() uint32 {
    v12 := []int{-1, 32767}
    _ = len(v12)
    for fi13 := 0; fi13 < 1; fi13++ {
        switch "foo bar" {
        case "a":
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int = fi13
                switch ((("hello" + ((("foo bar" + "") + ("hello" + "")) + "\n\t")) + "a") + "") {
                case "hello":
                    switch (^(fi14)) {
                    case 721:
                        if (793 > 42) {
                            var v16 int8 = 0
                            var v17 bool = ("\n\t" == "\n\t")
                            _ = v16
                            _ = v17
                        } else {
                            if ((true || (!((("" + "\n\t") + "a") != ("\n\t" + "foo bar")))) || (0 <= 1)) {
                                var v18 uint = 55
                                v12 = append(v12, (fi13 >> ((v18) & 63)))
                                _ = v18
                            } else {
                                fmt.Println(v15)
                            }
                            v15 = (-(v15))
                            v19 := GenS1{F0: 42, F1: -778}
                            _ = v19
                        }
                    case 392:
                        v20 := GenS1{F0: 1, F1: 1}
                        _ = v20
                    }
                case "":
                    v21 := GenS1{F0: 0, F1: -1}
                    _ = v21
                default:
                    v15 -= fi13
                }
                _ = fi14
                _ = v15
            }
        }
        switch 42 {
        case 9223372036854775807:
            var v22 int64 = 42
            _ = v22
        case -629:
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 bool = (!true)
                _ = fi23
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint = 4294967295
                _ = fi25
                _ = v26
            }
        }
        _ = fi13
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int64 = 42
    var v4 int64 = (v3 >> ((0) & 63))
    var v5 string = ""
    var v6 bool = (!true)
    if (((float32(3.14159) == -2.25) && v6) && (0.0 != 1e10)) {
        var v27 uint16 = 65535
        _ = v27
    } else {
        v28 := GenS2{F0: (v5 + v5), F1: 396}
        _ = v28
    }
    var v29 int64 = (v3 / ((9223372036854775807) | 1))
    v30 := GenS1{F0: 358, F1: v29}
    delete(v1, -1)
    var v31 uint64 = 824
    var v32 int32 = 0
    switch "foo bar" {
    case "a":
        for fi33 := 0; fi33 < 2; fi33++ {
            switch -430 {
            case 4294967295:
                if (v5 >= "\n\t") {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        v29 = v4
                        v35 := []int{(-(fi33)), (fi34 ^ 715), (fi33 * fi34)}
                        fmt.Println(v32)
                        var v36 int8 = 64
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                } else {
                    if v6 {
                        v37 := v2[v5]
                        _ = v37
                    } else {
                        if v6 {
                            var v38 int8 = -128
                            v4 += -575
                            var v39 string = v5
                            _ = v38
                            _ = v39
                        }
                    }
                    if (-80 == 1) {
                        v3 += (v29 ^ -1)
                        var v40 uint = 4294967295
                        for fi41 := 0; fi41 < 11; fi41++ {
                            fmt.Println(fi41)
                            _ = fi41
                        }
                        _ = v40
                    } else {
                        v4 -= (v4 | (^(v3)))
                    }
                    if v6 {
                        var v42 uint32 = 42
                        _ = v42
                    }
                }
            }
            _ = fi33
        }
    case "\n\t":
        v43 := GenS3{F0: 0, F1: 32767, F2: 127}
        _ = v43
    }
    v44 := []int{}
    if v6 {
        switch "a" {
        case "foo bar":
            var v45 string = ("hello" + v5)
            _ = v45
        }
    }
    var v46 int16 = 1000
    var v47 int = -237
    v48 := GenS4{F0: v6, F1: 127, F2: v47}
    var v49 uint8 = 31
    if v6 {
        for fi50 := 0; fi50 < 16; fi50++ {
            v51 := GenS4{F0: (!(1e10 >= 0.1)), F1: 1, F2: (v47 ^ (v47 % (((v47 & (^(fi50)))) | 1)))}
            _ = fi50
            _ = v51
        }
    } else {
        delete(v2, "a")
        switch 128 {
        case 994:
            fmt.Println(v31, v31)
        case 0:
            v44 = append(v44, v47)
        case 226:
            delete(v1, (v47 - 9223372036854775807))
        }
    }
    var v52 uint8 = v49
    fmt.Println(v52, v5, v46)
    v53 := GenS5{F0: v31, F1: 255, F2: (v29 & (^(v4)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
