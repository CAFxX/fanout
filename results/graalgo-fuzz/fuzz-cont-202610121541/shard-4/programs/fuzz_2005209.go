package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2147483647}
    v3 := []string{"", ("foo bar" + "\n\t")}
    var v4 int8 = -128
    fmt.Println(v4)
    v4 = (v4 & (^(v4)))
    switch 686 {
    case 9223372036854775807:
        var v5 int8 = 21
        _ = v5
    }
    v4 = 100
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v4 = 1
            v8 := GenS1{F0: 367}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if true {
        v4 += v4
        fmt.Println(v4)
    }
    switch "foo bar" {
    case "hello":
        v9 := GenS1{F0: 0}
        _ = v9
    case "a":
        if (801 >= 42) {
            v10 := v1[2147483647]
            _ = v10
        }
    }
    var v11 uint32 = 2147483647
    v11 += (v11 - 718)
    var v12 string = ("a" + (("foo bar" + "foo bar") + "a"))
    switch -328 {
    case 32768:
        var v13 float64 = 3.14159
        _ = v13
    case 596:
        v14 := GenS2{F0: (false && ((42 != -2147483648) || true)), F1: -83, F2: 255}
        _ = v14
    case -158:
        if true {
            v11 += v11
            v15 := v1[(^(v11))]
            if len(v3) > 0 {
                v3[0] = v12
            }
            _ = v15
        } else {
            v12 = "foo bar"
            var v16 int32 = -501
            _ = v16
        }
    }
    var v17 int32 = 42
    if (v4 != -65) {
        for fi18 := 0; fi18 < 6; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint = 65535
                fmt.Println(v20, v12)
                var v21 int16 = -1
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                if true {
                    var v23 uint32 = 2147483647
                    _ = v23
                }
                switch v12 {
                case "":
                    v24 := GenS1{F0: 18446744073709551615}
                    _ = v24
                case "\n\t":
                    v25 := GenS2{F0: (false && false), F1: fi18, F2: 0}
                    _ = v25
                case "a":
                    v26 := GenS1{F0: 255}
                    _ = v26
                }
                _ = fi22
            }
            _ = fi18
        }
        fmt.Println(v12, v17)
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            if (!(false && (!(!true)))) {
                if (-2.25 != (float32(-2.25) - (0.0 + 1.5))) {
                    switch "hello" {
                    case "":
                        v29 := GenS2{F0: true, F1: -557, F2: 255}
                        _ = v29
                    case "\n\t":
                        v11 -= v11
                    default:
                        var v30 string = ((v12 + v12) + "")
                        _ = v30
                    }
                }
                if (!true) {
                    var v31 int16 = 1
                    if true {
                        v4 += (^(v4))
                        fmt.Println(v17, v17)
                        var v32 string = v12
                        var v33 uint8 = 42
                        _ = v32
                        _ = v33
                    } else {
                        var v34 float32 = 1.5
                        fmt.Println(fi27)
                        _ = v34
                    }
                    _ = v31
                } else {
                    v35 := GenS2{F0: true, F1: 127, F2: 48}
                    _ = v35
                }
            }
            switch 0 {
            case -1:
                if (175 >= 2) {
                    if false {
                        var v36 int = (^(fi28))
                        switch "\n\t" {
                        case "a":
                            var v37 bool = true
                            _ = v37
                        default:
                            v1[(v11 | 2)] = fi28
                        }
                        var v38 uint64 = 521
                        _ = v36
                        _ = v38
                    }
                    fmt.Println(v17, fi27)
                    if (467 <= 269) {
                        var v39 uint = 42
                        if false {
                            if ((fi28 + (fi28 >> ((v39) & 63))) < fi27) {
                                fmt.Println(v4, v39, v4)
                                v40 := GenS3{F0: 32767}
                                v39 -= (v39 ^ v39)
                                _ = v40
                            }
                            if true {
                                v4 = v4
                                v41 := GenS4{F0: 658, F1: 1}
                                var v42 int64 = 1
                                fmt.Println(v11)
                                if ("" >= (v12 + ("hello" + v12))) {
                                    var v43 uint = 562
                                    switch v12 {
                                    case "hello":
                                        if false {
                                            fmt.Println(v39, v11)
                                        } else {
                                            _ = len(v2)
                                            var v44 uint64 = 2
                                            v45 := GenS5{F0: 0, F1: (v42 / ((v42) | 1))}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v46 int32 = v17
                                            var v47 int8 = v4
                                            var v48 int16 = 1
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    case "\n\t":
                                        fmt.Println(v39, v43, fi27)
                                    case "foo bar":
                                        v1[(^(v11))] = fi27
                                    default:
                                        v4 += 100
                                    }
                                    fmt.Println(v12, fi28)
                                    _ = v43
                                }
                                _ = v41
                                _ = v42
                            }
                        }
                        _ = v39
                    }
                } else {
                    switch v12 {
                    case "a":
                        var v49 string = v12
                        _ = v49
                    case "":
                        var v50 int16 = 0
                        _ = v50
                    }
                    var v51 float32 = 0.1
                    _ = v51
                }
            default:
                switch "a" {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    var v52 uint64 = uint64(v11)
                    _ = v52
                }
            }
            var v53 int = fi27
            _ = fi28
            _ = v53
        }
        v54 := GenS3{F0: 253}
        _ = fi27
        _ = v54
    }
    for fi55 := 0; fi55 < 19; fi55++ {
        v56 := GenS3{F0: 32767}
        fmt.Println(fi55, v11)
        switch v12 {
        case "a":
            if (!true) {
                v11 = (v11 >> ((892) & 63))
            }
        case "foo bar":
            switch v12 {
            case "\n\t":
                v4 = (v4 & v4)
            case "a":
                var v57 uint16 = 518
                _ = v57
            case "":
                var v58 uint8 = 148
                _ = v58
            }
        default:
            v59 := GenS6{F0: v11, F1: (v4 >> ((87) & 63))}
            _ = v59
        }
        _ = fi55
        _ = v56
    }
    v11 = (v11 & v11)
    v60 := GenS1{F0: 79}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
