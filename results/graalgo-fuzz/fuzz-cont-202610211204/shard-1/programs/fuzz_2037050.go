package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint32 = 454
    var v2 string = "hello"
    var v3 float64 = -2.25
    v4 := GenS1{F0: v1}
    fmt.Println(v2, v3)
    v5 := GenS1{F0: (v1 - v1)}
    switch (v2 + v2) {
    case "a":
        v6 := GenS1{F0: (^(v1))}
        _ = v6
    case "":
        var v7 int8 = 0
        _ = v7
    default:
        v8 := GenS1{F0: (v1 % ((945) | 1))}
        _ = v8
    }
    if (!false) {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v2, fi9, v1)
            _ = fi9
        }
        switch (v2 + ("foo bar" + ("" + v2))) {
        case "foo bar":
            var v10 uint8 = 2
            _ = v10
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: v1}
                if (!(-32768 <= -346)) {
                    var v13 int16 = 32767
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
        }
        switch 9223372036854775807 {
        case 32767:
            switch 4294967295 {
            case 2147483647:
                if true {
                    v14 := GenS2{F0: v2, F1: 127}
                    if (false && (255 <= 42)) {
                        v1 = v1
                        switch -2147483648 {
                        case 495:
                            for fi15 := 0; fi15 < 12; fi15++ {
                                if (-977 == -288) {
                                    v3 += 3.14159
                                    v16 := []int{}
                                    if false {
                                        v16 = append(v16, fi15)
                                        if (850 <= 4294967295) {
                                            for fi17 := 0; fi17 < 3; fi17++ {
                                                switch "a" {
                                                case "a":
                                                    v2 = v2
                                                default:
                                                    v2 = "foo bar"
                                                }
                                                _ = fi17
                                            }
                                        } else {
                                            var v18 bool = false
                                            for fi19 := 0; fi19 < 16; fi19++ {
                                                if (33 == 255) {
                                                    var v20 int = (fi15 << ((2) & 63))
                                                    _ = v20
                                                } else {
                                                    var v21 float32 = -2.25
                                                    switch fi19 {
                                                    case -2147483648:
                                                        v22 := GenS1{F0: (v1 / ((v1) | 1))}
                                                        v2 = ((v2 + ((v2 + v2) + v2)) + ((("" + "foo bar") + ("a" + v2)) + v2))
                                                        _ = v22
                                                    case 0:
                                                        var v23 int = fi15
                                                        _ = v23
                                                    default:
                                                        v18 = v18
                                                        if (v18 && (0 <= 417)) {
                                                            v3 = (v3 / float64(-2.25))
                                                        } else {
                                                            v1 = v1
                                                            v2 = v2
                                                            v16 = append(v16, 209)
                                                        }
                                                        v24 := GenS3{F0: (^(v1)), F1: 127}
                                                        _ = v24
                                                    }
                                                    _ = v21
                                                }
                                                switch (fi19 << ((2) & 63)) {
                                                case 507:
                                                    fmt.Println(fi15, v1)
                                                case -1:
                                                    if (191 <= 9223372036854775807) {
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                        fmt.Println(v3, fi15, v2)
                                                        switch "hello" {
                                                        case "hello":
                                                            var v25 int8 = 52
                                                            _ = v25
                                                        case "\n\t":
                                                            var v26 int8 = 0
                                                            _ = v26
                                                        default:
                                                            var v27 string = v2
                                                            _ = v27
                                                        }
                                                    } else {
                                                        v16 = append(v16, (fi15 | (fi19 & fi19)))
                                                        v16 = append(v16, (fi15 | 128))
                                                    }
                                                }
                                                _ = fi19
                                            }
                                            switch v2 {
                                            case "hello":
                                                fmt.Println(v1, fi15, v1)
                                            }
                                            _ = v18
                                        }
                                        v28 := GenS4{F0: uint8(v1)}
                                        _ = v28
                                    }
                                    _ = v16
                                }
                                var v29 string = ((v2 + v2) + "\n\t")
                                var v30 uint = 512
                                _ = fi15
                                _ = v29
                                _ = v30
                            }
                        case 120:
                            switch -2147483648 {
                            case -1:
                                v31 := []int{2147483647}
                                _ = v31
                            case -168:
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    var v33 int32 = 1
                                    var v34 uint32 = v1
                                    v35 := []int{}
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            default:
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    v37 := []int{9223372036854775807, fi36, fi36}
                                    _ = fi36
                                    _ = v37
                                }
                            }
                        case 32767:
                            v2 = v2
                        }
                        v38 := []int{127, 811, 32767, 42}
                        _ = v38
                    } else {
                        v39 := []int{-719, 127, 123}
                        var v40 uint16 = 593
                        _ = v39
                        _ = v40
                    }
                    _ = v14
                } else {
                    if true {
                        var v41 uint16 = 1
                        var v42 uint16 = v41
                        v2 = v2
                        _ = v41
                        _ = v42
                    } else {
                        for fi43 := 0; fi43 < 20; fi43++ {
                            var v44 uint64 = 1
                            _ = fi43
                            _ = v44
                        }
                    }
                }
            }
        default:
            v1 += v1
        }
    }
    var v45 float64 = 1e10
    v46 := GenS5{F0: v1}
    fmt.Println(v2)
    for fi47 := 0; fi47 < 9; fi47++ {
        v1 += v1
        v48 := GenS6{F0: 18446744073709551615, F1: 4294967295}
        var v49 int16 = -130
        _ = fi47
        _ = v48
        _ = v49
    }
    var v50 int8 = -128
    for fi51 := 0; fi51 < 19; fi51++ {
        switch "hello" {
        case "hello":
            v45 -= v3
        case "":
            v52 := GenS2{F0: "", F1: (fi51 / ((-1) | 1))}
            _ = v52
        case "foo bar":
            v53 := GenS7{F0: 1, F1: (v50 >> ((1) & 63))}
            _ = v53
        }
        _ = fi51
    }
    var v54 uint64 = 0
    switch (v2 + v2) {
    case "a":
        fmt.Println(v2)
    case "hello":
        v55 := []int{466, -364}
        _ = v55
    case "foo bar":
        for fi56 := 0; fi56 < 13; fi56++ {
            v2 = "hello"
            v54 = (^(v54))
            _ = fi56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v50
    _ = v54
}
