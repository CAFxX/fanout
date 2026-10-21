package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 bool }
type GenS4 struct { F0 int16; F1 int64; F2 uint32 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 bool; F1 int64; F2 string }
type GenS7 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint32 = 0
    v2 := []int{-9223372036854775808, -496, 520}
    v3 := make(map[int]int)
    v4 := []uint32{}
    v5 := []int{0, -398, 2147483647, -988}
    var v6 int8 = 42
    var v7 bool = true
    var v8 int64 = -895
    switch 9223372036854775807 {
    case 0:
        _ = len(v3)
    }
    var v9 bool = (42 <= 2)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int64 = 0
        var v12 uint16 = 32767
        if false {
            v5 = append(v5, 127)
            fmt.Println(v8)
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: false, F1: -1, F2: 42}
    v14 := GenS1{F0: v7, F1: (v6 - (^(v6))), F2: 377}
    switch ("foo bar" + "hello") {
    case "\n\t":
        fmt.Println(v6, v9)
    case "a":
        var v15 int64 = v8
        _ = v15
    case "hello":
        if ("" < "\n\t") {
            for fi16 := 0; fi16 < 4; fi16++ {
                if (621 == 65535) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if (!false) {
                            switch (-(fi17)) {
                            case 32768:
                                v7 = (v8 > (v8 % ((v8) | 1)))
                            case 4294967295:
                                switch "hello" {
                                case "\n\t":
                                    var v18 uint = 42
                                    _ = v18
                                default:
                                    _ = len(v2)
                                }
                            case 141:
                                fmt.Println(v7, v6)
                            default:
                                var v19 uint = 272
                                _ = v19
                            }
                            if true {
                                delete(v3, fi16)
                                var v20 uint = 57
                                _ = v20
                            } else {
                                fmt.Println(v8)
                                v8 = (v8 >> ((4294967295) & 63))
                            }
                        } else {
                            if v7 {
                                v21 := GenS2{F0: 0}
                                v8 = 1
                                var v22 float64 = 0.0
                                _ = v21
                                _ = v22
                            } else {
                                if false {
                                    var v23 int64 = v8
                                    v24 := GenS3{F0: 88, F1: "foo bar", F2: false}
                                    var v25 uint32 = 945
                                    switch "" {
                                    case "foo bar":
                                        var v26 string = ("foo bar" + "")
                                        _ = v26
                                    case "hello":
                                        var v27 float64 = (float64(1e10) + ((float64(float32(0.1)) - 0.1) - ((1e10 - 0.0) + 0.0)))
                                        _ = v27
                                    }
                                    v8 = v23
                                    fmt.Println(v8, v25)
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                            if (1e10 != (-2.25 + 1.5)) {
                                if v9 {
                                    v28 := GenS2{F0: 255}
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 457, F1: v8, F2: (v1 % (((v1 | v1)) | 1))}
                                    delete(v3, fi17)
                                    _ = v29
                                }
                            }
                            var v30 int64 = v8
                            _ = v30
                        }
                        v31 := GenS3{F0: 4294967295, F1: "\n\t", F2: ((fi16 | fi16) > fi16)}
                        _ = fi17
                        _ = v31
                    }
                    v9 = (true && v9)
                    var v32 uint64 = 4294967295
                    _ = v32
                } else {
                    switch (fi16 % ((fi16) | 1)) {
                    case 32768:
                        var v33 int8 = -1
                        _ = v33
                    default:
                        v34 := GenS4{F0: 32767, F1: -628, F2: (v1 << ((4294967295) & 63))}
                        _ = v34
                    }
                }
                v5 = append(v5, fi16)
                v8 -= v8
                _ = fi16
            }
            if ((float32(((1e10 / 1.0) * 1.5)) / ((-2.25 / 1e10) / 0.1)) == (-2.25 - ((-2.25 * 1e10) * (1e10 * float32((0.0 + ((0.0 - 1e10) * ((float64(-2.25) + 0.1) - 1e10)))))))) {
                for fi35 := 0; fi35 < 13; fi35++ {
                    v8 = v8
                    switch "a" {
                    case "hello":
                        v7 = ((fi35 - (fi35 >> ((666) & 63))) >= (fi35 << ((777) & 63)))
                    case "":
                        var v36 string = "hello"
                        _ = v36
                    }
                    _ = fi35
                }
                if len(v2) > 0 {
                    v2[0] = 291
                }
                var v37 uint16 = 255
                _ = v37
            }
        }
    }
    delete(v3, -9223372036854775808)
    if (v9 || false) {
        if v9 {
            for fi38 := 0; fi38 < 1; fi38++ {
                if ((v6 >> ((65535) & 63)) >= (^(v6))) {
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if (-479 > -2147483648) {
                            var v40 int16 = 42
                            var v41 int = (fi38 - fi39)
                            v42 := v3[fi38]
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = fi39
                    }
                    v7 = (v9 && (((0.1 * float64(1.5)) - 1e10) != ((float64(1e10) * 0.1) - -2.25)))
                    if v9 {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v1 -= 239
                            v44 := GenS1{F0: (4294967295 == 0), F1: (v6 | v6), F2: 617}
                            fmt.Println(v7, fi38, v8)
                            _ = fi43
                            _ = v44
                        }
                    } else {
                        fmt.Println(fi38, v6)
                        var v45 string = "a"
                        _ = v45
                    }
                }
                _ = fi38
            }
            v1 += (v1 >> ((996) & 63))
        } else {
            v46 := GenS4{F0: -359, F1: v8, F2: 4294967295}
            v7 = v9
            var v47 uint16 = 2
            _ = v46
            _ = v47
        }
    }
    var v48 string = "a"
    fmt.Println(v9)
    if true {
        var v49 uint = 4294967295
        _ = v49
    }
    if v7 {
        switch (("\n\t" + (v48 + v48)) + "") {
        case "foo bar":
            for fi50 := 0; fi50 < 14; fi50++ {
                fmt.Println(v48)
                for fi51 := 0; fi51 < 17; fi51++ {
                    switch v48 {
                    case "a":
                        if v9 {
                            _ = len(v2)
                        }
                    case "foo bar":
                        v52 := v3[fi51]
                        _ = v52
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                        v1 = v1
                    default:
                        v1 = 1
                    }
                    var v53 int16 = -453
                    switch (v48 + v48) {
                    case "\n\t":
                        v6 = (-(v6))
                        fmt.Println(v7, v48, v9)
                    }
                    var v54 float64 = -2.25
                    v8 -= v8
                    _ = fi51
                    _ = v53
                    _ = v54
                }
                if v7 {
                    switch int(v6) {
                    case -873:
                        v55 := GenS5{F0: v8, F1: -128}
                        _ = v55
                    case 1:
                        v56 := GenS6{F0: true, F1: (v8 ^ (-(v8))), F2: ("\n\t" + ((("" + v48) + v48) + "\n\t"))}
                        _ = v56
                    }
                    if v7 {
                        var v57 uint32 = v1
                        _ = v57
                    } else {
                        v58 := GenS2{F0: 127}
                        v59 := v3[32768]
                        if (v7 && (0 < 0)) {
                            v60 := GenS6{F0: ((v9 || true) && v7), F1: v8, F2: ""}
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            var v61 float32 = 0.0
                            _ = v60
                            _ = v61
                        }
                        _ = v58
                        _ = v59
                    }
                } else {
                    v6 = -81
                }
                _ = fi50
            }
        default:
            switch v48 {
            case "a":
                v7 = false
            }
        }
        var v62 int16 = 746
        v63 := GenS1{F0: ((v7 || v9) && true), F1: v6, F2: 1}
        _ = v62
        _ = v63
    } else {
        v64 := GenS6{F0: (v7 || (!(v7 && (54 >= 974)))), F1: 0, F2: (v48 + v48)}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v64
    }
    v65 := GenS7{F0: 0, F1: -1}
    switch (v48 + "hello") {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v65
    _ = v7
    _ = v8
    _ = v9
}
