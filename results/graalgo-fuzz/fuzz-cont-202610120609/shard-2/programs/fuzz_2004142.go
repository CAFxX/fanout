package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 int }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int8; F1 uint32; F2 uint64 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    var v6 uint64 = 42
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 0, F1: 382}
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 -= p0
    for fi10 := 0; fi10 < 9; fi10++ {
        switch (fi10 - fi10) {
        case -100:
            v6 = v6
        case -163:
            v11 := GenS1{F0: 127, F1: -750}
            _ = v11
        default:
            switch p0 {
            case -908:
                v6 = (v6 + v6)
            case -1:
                p0 += p0
            case -9223372036854775808:
                v12 := []int{728}
                _ = v12
            }
        }
        switch 127 {
        case 42:
            if (("\n\t" + "foo bar") > "foo bar") {
                v13 := GenS2{F0: 0, F1: -63, F2: 2}
                switch (("" + "a") + "foo bar") {
                case "foo bar":
                    if false {
                        fmt.Println(fi10, v6)
                    }
                case "hello":
                    if false {
                        var v14 int = (p0 & (^(p0)))
                        v14 = p0
                        _ = v14
                    } else {
                        switch fi10 {
                        case -874:
                            v6 += (v6 - (v6 * v6))
                        default:
                            p0 = fi10
                        }
                    }
                }
                var v15 uint = 1
                _ = v13
                _ = v15
            }
        default:
            if (true && true) {
                var v16 bool = true
                var v17 uint64 = v6
                _ = v16
                _ = v17
            } else {
                var v18 int16 = 537
                v19 := GenS1{F0: 1, F1: (v18 * v18)}
                _ = v18
                _ = v19
            }
        }
        _ = fi10
    }
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint16 = 32767
    var v2 string = ("a" + "foo bar")
    var v3 int16 = 32767
    v4 := []string{"a", v2}
    v5 := []string{v2, v2}
    var v20 uint = 0
    var v21 float32 = 1.5
    var v22 uint32 = 4294967295
    for fi23 := 0; fi23 < 12; fi23++ {
        if true {
            fmt.Println(v3)
            fmt.Println(v22)
            v24 := GenS2{F0: v1, F1: 100, F2: uint32(v20)}
            _ = v24
        } else {
            var v25 bool = ((v20 >> ((v20) & 63)) == (^(v20)))
            _ = len(v5)
            _ = v25
        }
        var v26 string = (v2 + v2)
        _ = fi23
        _ = v26
    }
    v2 = v2
    v27 := GenS1{F0: 127, F1: 1}
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS1{F0: 1, F1: int16(fi28)}
        if (((v3 / ((v3) | 1)) > (v3 * 463)) && (true || ((0 == -384) || false))) {
            switch fi28 {
            case 2147483647:
                if true {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        switch (fi28 + fi28) {
                        case -464:
                            v31 := GenS2{F0: (v1 / ((v1) | 1)), F1: -1, F2: (v22 << ((0) & 63))}
                            _ = v31
                        case 32767:
                            if (65535 <= (v1 - v1)) {
                                v22 -= (v22 >> (((v20 / ((v20) | 1))) & 63))
                                v32 := GenS3{F0: (((v2 + (v2 + v2)) + "hello") + v2), F1: -128, F2: fi30}
                                _ = v32
                            } else {
                                v22 += (v22 % (((v22 & v22)) | 1))
                            }
                            v1 -= 7
                        case -376:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v33 uint64 = 1
                        _ = fi30
                        _ = v33
                    }
                    var v34 string = ("" + (("" + v2) + v2))
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v4 = append(v4, v2)
                        _ = fi35
                    }
                    v21 = (((v21 / (3.14159 - 1.5)) / (v21 * (1.5 / v21))) * -2.25)
                }
            default:
                if (!true) {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        if false {
                            v22 = v22
                            v21 -= v21
                            if len(v4) > 0 {
                                v4[0] = v2
                            }
                        }
                        v2 = v2
                        _ = fi36
                    }
                } else {
                    v2 = (v2 + v2)
                }
            }
        } else {
            v21 = (v21 / (v21 / (-2.25 / (v21 * (1e10 / float32((((3.14159 / 1.0) * float64(float32((0.1 - 0.0)))) / (1e10 + 0.0))))))))
            if false {
                var v37 uint16 = (^(v1))
                if (!true) {
                    if len(v5) > 0 {
                        v5[0] = v2
                    }
                    var v38 uint = 1
                    _ = v38
                } else {
                    v5 = append(v5, v2)
                }
                if true {
                    v39 := GenS2{F0: (v1 - v37), F1: 0, F2: v22}
                    if (!false) {
                        if (!(!(true || (0.0 == 1.5)))) {
                            if true {
                                if true {
                                    var v40 uint = v20
                                    var v41 int8 = 100
                                    _ = v40
                                    _ = v41
                                }
                            }
                        } else {
                            v1 -= 255
                            v42 := GenS4{F0: "hello", F1: fi28}
                            v43 := GenS3{F0: (v2 + "hello"), F1: 127, F2: fi28}
                            _ = v42
                            _ = v43
                        }
                        var v44 float64 = 3.14159
                        _ = v44
                    }
                    var v45 int32 = -2147483648
                    _ = v39
                    _ = v45
                } else {
                    v21 += 1e10
                    v46 := GenS2{F0: (v1 ^ v1), F1: int8(v22), F2: (^(v22))}
                    v3 = v3
                    _ = v46
                }
                _ = v37
            } else {
                for fi47 := 0; fi47 < 13; fi47++ {
                    var v48 string = ""
                    v2 = (v48 + v2)
                    var v49 uint8 = 1
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
            if true {
                switch "" {
                case "\n\t":
                    if (42 <= 0) {
                        fmt.Println(v3)
                        v50 := GenS5{F0: 100, F1: 0}
                        v4 = append(v4, "hello")
                        _ = v50
                    }
                case "hello":
                    var v51 uint = v20
                    _ = v51
                }
                v52 := GenS5{F0: 84, F1: (v22 * v22)}
                _ = v52
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v53 int64 = 1
                v3 += v3
                _ = v53
            }
        }
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v2)
    switch 1 {
    case 2147483647:
        fmt.Println(v20)
    case -422:
        _ = len(v4)
    case 9223372036854775807:
        for fi54 := 0; fi54 < 7; fi54++ {
            switch "foo bar" {
            case "hello":
                if true {
                    var v55 int16 = (v3 * v3)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v56 int64 = -9223372036854775808
                    _ = v55
                    _ = v56
                }
            case "foo bar":
                switch 32768 {
                case 955:
                    v57 := GenS6{F0: 18446744073709551615, F1: 1}
                    _ = v57
                default:
                    v58 := GenS3{F0: ("a" + ("foo bar" + v2)), F1: 50, F2: 32767}
                    _ = v58
                }
            }
            fmt.Println(fi54)
            _ = fi54
        }
    }
    v4 = append(v4, v2)
    fmt.Println(v1)
    if (1e10 <= 1e10) {
        if false {
            fmt.Println(v20, v22, v22)
        }
        v59 := GenS3{F0: v2, F1: 42, F2: -26}
        _ = v59
    } else {
        for fi60 := 0; fi60 < 12; fi60++ {
            v61 := GenS3{F0: "", F1: 0, F2: (fi60 >> ((v20) & 63))}
            v62 := GenS7{F0: 100, F1: 2, F2: 2}
            _ = fi60
            _ = v61
            _ = v62
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
