package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int8; F1 uint64; F2 int8 }
func genFn0(p0 uint) int {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: "a", F1: 0}
        fmt.Println(fi4, fi4)
        _ = fi4
        _ = v5
    }
    var v6 int = 4294967295
    v6 = v6
    _ = p0
    _ = v6
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    var v3 string = "a"
    v7 := []int{}
    fmt.Println(v3, v2)
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: v3, F1: 84}
        if true {
            switch fi8 {
            case 32767:
                switch v3 {
                case "foo bar":
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 uint = 18446744073709551615
                        if true {
                            if ((42 > 0) || (42 == 1)) {
                                v1[v3] = (-128 != -7)
                                v12 := GenS1{F0: v3, F1: 1}
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = v12
                            } else {
                                if (false && ((((v3 + v3) >= v3) && (413 <= v11)) || false)) {
                                    v13 := GenS1{F0: "", F1: 69}
                                    var v14 float64 = 1.5
                                    switch v3 {
                                    case "foo bar":
                                        v15 := GenS1{F0: v3, F1: -1}
                                        if true {
                                            var v16 uint8 = 1
                                            if false {
                                                v17 := GenS1{F0: v3, F1: -61}
                                                fmt.Println(v16)
                                                _ = v17
                                            }
                                            v18 := GenS2{F0: ""}
                                            v16 -= (v16 + (v16 ^ (v16 / ((4) | 1))))
                                            var v19 float32 = 1e10
                                            _ = v16
                                            _ = v18
                                            _ = v19
                                        } else {
                                            v3 = v3
                                            fmt.Println(v14)
                                        }
                                        if true {
                                            v20 := GenS1{F0: ("\n\t" + ((v3 + v3) + v3)), F1: 0}
                                            _ = v20
                                        } else {
                                            delete(v1, (v3 + ""))
                                            v2 -= v2
                                        }
                                        _ = v15
                                    case "a":
                                        fmt.Println(v14, v3)
                                    }
                                    _ = v13
                                    _ = v14
                                }
                                v21 := GenS3{F0: -1}
                                var v22 float64 = (-2.25 / 1.5)
                                _ = v21
                                _ = v22
                            }
                        }
                        switch genFn0((v11 & 255)) {
                        case 769:
                            var v23 int32 = -1
                            _ = v23
                        case -191:
                            v7 = append(v7, (fi10 + fi8))
                        }
                        _ = fi10
                        _ = v11
                    }
                case "\n\t":
                    for fi24 := 0; fi24 < 14; fi24++ {
                        var v25 uint64 = 2
                        var v26 int8 = -42
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                case "hello":
                    var v27 float32 = ((1e10 * (3.14159 - (1.5 + float32(float64(1e10))))) - float32(((float64(float32(3.14159)) + 1.5) + 0.0)))
                    _ = v27
                default:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        }
        switch fi8 {
        case 1:
            switch "hello" {
            case "a":
                var v28 float64 = -2.25
                _ = v28
            case "\n\t":
                for fi29 := 0; fi29 < 10; fi29++ {
                    switch "hello" {
                    case "":
                        v30 := GenS4{F0: 42, F1: 42}
                        _ = v30
                    case "\n\t":
                        v31 := GenS5{F0: -1, F1: -9223372036854775808}
                        _ = v31
                    case "foo bar":
                        var v32 int64 = v2
                        _ = v32
                    }
                    _ = fi29
                }
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        case -1:
            for fi33 := 0; fi33 < 7; fi33++ {
                if false {
                    fmt.Println(fi33, fi33, v3)
                    v7 = append(v7, (fi33 ^ 32768))
                }
                v34 := GenS2{F0: (v3 + (v3 + (v3 + ((((v3 + "\n\t") + v3) + v3) + "a"))))}
                v35 := GenS5{F0: -1, F1: -1}
                _ = fi33
                _ = v34
                _ = v35
            }
        case 127:
            v2 -= v2
        }
        _ = fi8
        _ = v9
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v37 := GenS4{F0: 42, F1: 0}
        v3 = ("" + "")
        _ = fi36
        _ = v37
    }
    v2 = (v2 * v2)
    switch -182 {
    case 102:
        fmt.Println(v3, v3)
    case 191:
        v2 = (v2 * (v2 << ((2) & 63)))
    default:
        v3 = "hello"
    }
    v3 = ("hello" + ((((v3 + v3) + "\n\t") + "\n\t") + (v3 + v3)))
    switch "hello" {
    case "a":
        if false {
            var v38 uint16 = 236
            v7 = append(v7, -404)
            _ = v38
        } else {
            var v39 int32 = -1
            _ = v39
        }
    case "foo bar":
        for fi40 := 0; fi40 < 3; fi40++ {
            fmt.Println(fi40, v3)
            for fi41 := 0; fi41 < 7; fi41++ {
                v2 = -1
                fmt.Println(fi40, fi40)
                if true {
                    v42 := GenS6{F0: 42}
                    v3 = v3
                    var v43 float64 = ((0.1 * (float64(1.5) + 0.1)) - 1e10)
                    switch (fi41 % (((fi41 | fi40)) | 1)) {
                    case 9223372036854775807:
                        fmt.Println(v43, fi41, fi41)
                        var v44 float64 = -2.25
                        _ = v44
                    }
                    fmt.Println(fi40, fi41)
                    if (3.14159 <= ((float32(-2.25) - 0.1) / ((((((float32(float64((3.14159 * (float32(v43) * (3.14159 - ((1e10 / 1.5) * ((3.14159 * (1e10 + 3.14159)) + 1e10))))))) / 1.0) - ((((0.0 * -2.25) + 3.14159) / 1e10) + 1e10)) + 1.5) - 3.14159) / 1.0) / 3.14159))) {
                        var v45 int = fi40
                        v46 := v1["a"]
                        _ = v45
                        _ = v46
                    }
                    _ = v42
                    _ = v43
                } else {
                    var v47 bool = (-774 > -9223372036854775808)
                    v2 += (v2 << ((255) & 63))
                    var v48 int16 = 511
                    _ = v47
                    _ = v48
                }
                _ = fi41
            }
            _ = fi40
        }
    case "":
        var v49 uint32 = 2147483648
        _ = v49
    default:
        v50 := GenS3{F0: 100}
        _ = v50
    }
    v2 = (^(v2))
    v2 = v2
    for fi51 := 0; fi51 < 5; fi51++ {
        if (111 != 190) {
            switch v3 {
            case "foo bar":
                if false {
                    for fi52 := 0; fi52 < 17; fi52++ {
                        v53 := v1[v3]
                        switch v3 {
                        case "":
                            var v54 uint = 42
                            _ = v54
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                switch (fi52 ^ -2147483648) {
                                case -873:
                                    var v55 uint32 = 1
                                    _ = v55
                                default:
                                    var v56 bool = v53
                                    _ = v56
                                }
                            case "a":
                                v57 := GenS7{F0: 42, F1: v53}
                                _ = v57
                            }
                        case "hello":
                            var v58 int32 = 0
                            _ = v58
                        }
                        _ = fi52
                        _ = v53
                    }
                    v59 := GenS8{F0: -28, F1: 619, F2: -18}
                    _ = v59
                } else {
                    if false {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                }
            case "\n\t":
                var v60 int8 = -104
                _ = v60
            default:
                v61 := GenS5{F0: 1, F1: v2}
                _ = v61
            }
        }
        var v62 int16 = -32768
        fmt.Println(v3, v2, v3)
        _ = fi51
        _ = v62
    }
    if len(v7) > 0 {
        v7[0] = -739
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
