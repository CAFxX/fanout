package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint16; F2 bool }
type GenS5 struct { F0 int8; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 987
    v3 := make(map[uint32]int)
    fmt.Println(v2)
    v2 += 897
    switch "hello" {
    case "\n\t":
        v4 := GenS1{F0: 23}
        _ = v4
    }
    _ = len(v1)
    v2 = v2
    v5 := v1[42]
    v6 := GenS1{F0: 1}
    switch (^(v5)) {
    case -1:
        if (630 != 42) {
            if true {
                fmt.Println(v2, v5)
                v5 += (-(v5))
                v2 -= -132
            }
        }
    default:
        if (255 <= 2) {
            for fi7 := 0; fi7 < 6; fi7++ {
                if (-191 > 42) {
                    if true {
                        v8 := GenS1{F0: 32767}
                        v5 += 808
                        switch 2147483647 {
                        case -891:
                            var v9 int8 = -101
                            _ = v9
                        }
                        _ = v8
                    }
                }
                if (1 == 255) {
                    var v10 float32 = (0.0 / -2.25)
                    var v11 int = -1
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(fi12, fi12, fi12)
                        delete(v3, 2)
                        _ = fi12
                    }
                    if true {
                        if false {
                            var v13 bool = (true || (328 < 42))
                            if v13 {
                                delete(v3, 2)
                                switch ("hello" + ("hello" + "hello")) {
                                case "foo bar":
                                    var v14 int16 = -293
                                    _ = v14
                                case "":
                                    for fi15 := 0; fi15 < 7; fi15++ {
                                        var v16 string = ("\n\t" + "a")
                                        _ = fi15
                                        _ = v16
                                    }
                                case "hello":
                                    v2 -= (-(fi7))
                                }
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    v5 = v5
                                    _ = fi17
                                }
                            } else {
                                var v18 int32 = -2147483648
                                _ = v18
                            }
                            _ = v13
                        } else {
                            v2 = v5
                        }
                        v19 := []int{32768, (v2 / ((-449) | 1)), v5, v5}
                        _ = v19
                    } else {
                        v2 = (v2 + (fi7 ^ v2))
                        for fi20 := 0; fi20 < 18; fi20++ {
                            switch "hello" {
                            case "":
                                fmt.Println(fi20, fi20, v2)
                            case "foo bar":
                                var v21 uint8 = 127
                                v21 = v21
                                _ = v21
                            case "hello":
                                var v22 int = (v5 / ((fi7) | 1))
                                _ = v22
                            default:
                                v23 := GenS1{F0: -923}
                                _ = v23
                            }
                            v24 := GenS1{F0: -581}
                            _ = fi20
                            _ = v24
                        }
                        if true {
                            v5 = 186
                        } else {
                            for fi25 := 0; fi25 < 18; fi25++ {
                                v2 -= 1
                                switch "a" {
                                case "":
                                    fmt.Println(v2, v5, v5)
                                case "foo bar":
                                    fmt.Println(fi25)
                                }
                                _ = fi25
                            }
                        }
                    }
                    v26 := GenS1{F0: -248}
                    _ = v26
                }
                _ = fi7
            }
            if false {
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := v1[2147483648]
                    var v29 float64 = (float64((float32(0.0) + 3.14159)) + 1.5)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch (v5 % ((v5) | 1)) {
                case 9223372036854775807:
                    var v30 uint16 = 862
                    _ = v30
                case -979:
                    var v31 uint64 = 4294967295
                    _ = v31
                case 275:
                    var v32 uint64 = 1
                    _ = v32
                default:
                    v33 := GenS1{F0: 32767}
                    _ = v33
                }
            } else {
                v5 = 2147483647
                v34 := GenS2{F0: 832}
                var v35 uint8 = 15
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := []int{(v5 & 9223372036854775807)}
                var v38 uint = 1
                switch fi36 {
                case -957:
                    v39 := GenS2{F0: 32767}
                    _ = v39
                case 1:
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 uint = v38
                        _ = fi40
                        _ = v41
                    }
                case 475:
                    v42 := GenS3{F0: (("a" + "") + ("foo bar" + "foo bar"))}
                    _ = v42
                }
                _ = fi36
                _ = v37
                _ = v38
            }
        } else {
            switch "" {
            case "hello":
                var v43 int64 = 9223372036854775807
                _ = v43
            case "a":
                var v44 int8 = -128
                _ = v44
            case "":
                var v45 int32 = -586
                _ = v45
            }
            var v46 float32 = 3.14159
            for fi47 := 0; fi47 < 15; fi47++ {
                var v48 int64 = 42
                _ = fi47
                _ = v48
            }
            _ = v46
        }
    }
    v5 = 127
    if (0.1 >= (1.5 * -2.25)) {
        v49 := GenS2{F0: 0}
        v1[425] = (v5 % (((v5 >> ((296) & 63))) | 1))
        _ = v49
    } else {
        switch "\n\t" {
        case "a":
            v50 := GenS4{F0: 551, F1: 469, F2: ((0.0 / 0.1) != -2.25)}
            _ = v50
        case "foo bar":
            v5 += (-(v5))
        case "":
            fmt.Println(v5)
        default:
            v5 = (v2 ^ 1)
        }
    }
    fmt.Println(v2, v5)
    if false {
        v5 = (v5 << ((4294967295) & 63))
        for fi51 := 0; fi51 < 6; fi51++ {
            for fi52 := 0; fi52 < 2; fi52++ {
                v53 := v1[65535]
                v54 := []int{v5}
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = fi51
        }
    } else {
        if true {
            v5 = (v5 % (((v2 * -2147483648)) | 1))
            v5 += 529
            var v55 bool = (!true)
            _ = v55
        } else {
            var v56 int16 = -432
            _ = v56
        }
    }
    v57 := GenS2{F0: -1}
    if false {
        v58 := []int{}
        _ = v58
    } else {
        var v59 int32 = 1
        _ = v59
    }
    v60 := GenS5{F0: 0, F1: 126, F2: -575}
    fmt.Println(v5, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v57
    _ = v6
    _ = v60
}
