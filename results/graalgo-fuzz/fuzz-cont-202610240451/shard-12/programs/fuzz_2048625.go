package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32; F1 uint8; F2 string }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 int8 = 42
    var v4 string = ""
    var v5 int32 = 2147483647
    v3 = v3
    switch 127 {
    case -892:
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                v5 = 42
                _ = fi7
            }
            var v8 int8 = 100
            _ = fi6
            _ = v8
        }
    case 42:
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: v5, F1: (fi9 - 162)}
                if (1 != 329) {
                    var v12 bool = false
                    v5 += (v5 + -2147483648)
                    _ = v12
                } else {
                    v4 = v4
                    var v13 int = (fi9 % (((fi10 | 9223372036854775807)) | 1))
                    _ = v13
                }
                v3 = (v3 - (^(v3)))
                v4 = "foo bar"
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case 590:
        if true {
            v14 := GenS2{F0: v5, F1: 1, F2: (((v4 + v4) + (("hello" + ("foo bar" + v4)) + v4)) + (v4 + (v4 + v4)))}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int64 = 325
        if (6 == 42) {
            v17 := []int{}
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := v2[fi15]
                var v20 int16 = 1000
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 uint32 = 141
            _ = v17
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v5)
    fmt.Println(v4)
    v4 = ("hello" + (((v4 + ((v4 + v4) + (v4 + (v4 + "")))) + "\n\t") + "hello"))
    if true {
        if false {
            var v22 int8 = (v3 | (-(v3)))
            v23 := []int{838, 32768}
            if (((v4 + v4) + v4) >= (v4 + ((v4 + v4) + (v4 + v4)))) {
                fmt.Println(v22, v3)
                if true {
                    v22 = v3
                    v24 := GenS3{F0: 1000, F1: (true && ((679 > 415) || (!true))), F2: (v22 | v3)}
                    var v25 uint64 = 0
                    _ = v24
                    _ = v25
                } else {
                    v5 += (v5 | v5)
                    _ = len(v1)
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v23 = append(v23, (fi26 - 243))
                        switch fi26 {
                        case 707:
                            for fi27 := 0; fi27 < 4; fi27++ {
                                _ = len(v23)
                                var v28 uint32 = 8
                                _ = fi27
                                _ = v28
                            }
                        case -488:
                            fmt.Println(v5, v5, v3)
                        case 32768:
                            for fi29 := 0; fi29 < 1; fi29++ {
                                _ = len(v23)
                                _ = fi29
                            }
                        default:
                            if false {
                                v30 := GenS3{F0: -1, F1: false, F2: -1}
                                for fi31 := 0; fi31 < 1; fi31++ {
                                    fmt.Println(v3)
                                    switch (v4 + ((v4 + (v4 + v4)) + v4)) {
                                    case "":
                                        v32 := GenS4{F0: (0.0 == (1.5 - 0.1))}
                                        _ = v32
                                    case "\n\t":
                                        if ((v5 <= (v5 | (-(v5)))) || false) {
                                            v33 := GenS2{F0: (v5 ^ v5), F1: 0, F2: "hello"}
                                            var v34 int = fi26
                                            _ = v33
                                            _ = v34
                                        }
                                    case "a":
                                        var v35 int64 = 307
                                        var v36 uint16 = 65535
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := GenS1{F0: (v5 - v5), F1: fi31}
                                    _ = fi31
                                    _ = v37
                                }
                                var v38 float32 = (1.5 * (1e10 + float32(0.0)))
                                _ = v30
                                _ = v38
                            } else {
                                var v39 uint32 = 1
                                switch (-(fi26)) {
                                case 32767:
                                    if (v22 <= v22) {
                                        switch fi26 {
                                        case -2147483648:
                                            switch (v4 + (v4 + (v4 + "\n\t"))) {
                                            case "\n\t":
                                                fmt.Println(v39)
                                            case "foo bar":
                                                fmt.Println(v5, v3)
                                            }
                                        case 1:
                                            switch int(v22) {
                                            case 127:
                                                v4 = v4
                                            default:
                                                if len(v23) >= 2 {
                                                    v23 = v23[1:len(v23)]
                                                }
                                            }
                                        case 311:
                                            v39 -= (^(v39))
                                        }
                                        v40 := GenS4{F0: (v5 != v5)}
                                        v39 += (v39 + v39)
                                        _ = v40
                                    }
                                case 42:
                                    if true {
                                        v22 += v3
                                        v23 = append(v23, 32767)
                                        if (4294967295 >= 18446744073709551615) {
                                            fmt.Println(v22, v5)
                                            v41 := GenS2{F0: v5, F1: 235, F2: ""}
                                            v5 = v5
                                            _ = v41
                                        }
                                    } else {
                                        v42 := GenS4{F0: false}
                                        v43 := v2[(fi26 ^ -229)]
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                v44 := GenS5{F0: int32(fi26), F1: (v5 ^ (v5 ^ (v5 >> ((255) & 63))))}
                                _ = v39
                                _ = v44
                            }
                        }
                        _ = fi26
                    }
                }
            }
            _ = v22
            _ = v23
        } else {
            var v45 int64 = 944
            _ = v45
        }
        var v46 uint = 1
        _ = v46
    } else {
        v3 = 127
        v5 += (v5 << ((65535) & 63))
        fmt.Println(v3, v5)
    }
    fmt.Println(v5)
    var v47 int8 = (^(v3))
    fmt.Println(v4, v47, v3)
    v48 := GenS5{F0: (v5 - v5), F1: v5}
    delete(v2, 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
