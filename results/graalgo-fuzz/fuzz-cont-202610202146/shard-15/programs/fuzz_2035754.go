package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 string; F1 int64; F2 bool }
type GenS6 struct { F0 int8; F1 bool; F2 uint16 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 int; F1 uint16 }
func genFn0(p0 uint16, p1 int, p2 uint16) uint16 {
    var v7 int = 32767
    p0 = 319
    if (-727 >= -887) {
        switch (-(v7)) {
        case 305:
            v8 := []int{(p1 & p1), v7, (-(p1))}
            _ = v8
        case 149:
            var v9 bool = false
            _ = v9
        default:
            p1 = p1
        }
        p0 += p0
        var v10 string = ("foo bar" + "\n\t")
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 uint, p1 int16, p2 int) uint64 {
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: false}
        _ = fi11
        _ = v12
    }
    if ((float64(3.14159) - float64(((3.14159 - (1.5 - 3.14159)) + ((1e10 - 1.5) / (3.14159 + (0.1 - -2.25)))))) >= ((0.1 + (((1.5 + -2.25) - 0.1) / float64(0.1))) / 1.0)) {
        p0 -= (p0 >> (((p0 ^ p0)) & 63))
        p1 = (p1 * p1)
    } else {
        v13 := []int{32768}
        fmt.Println(p0, p2, p2)
        v14 := GenS2{F0: 128, F1: int(p1)}
        _ = v13
        _ = v14
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 118
    var v2 int = 128
    v3 := []int{(v2 - (v2 % (((-(v2))) | 1)))}
    var v4 uint64 = 4294967295
    var v5 uint64 = (v4 / (((^(v4))) | 1))
    var v6 int16 = 42
    fmt.Println(v4)
    if (!(!false)) {
        switch "a" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = -652
            }
        case "a":
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v2)
                var v16 uint8 = 128
                if ((v6 / (((v6 - 755)) | 1)) > (v6 | (v6 | v6))) {
                    v5 = 522
                } else {
                    var v17 uint = 0
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        case "foo bar":
            if (128 > 127) {
                v18 := GenS3{F0: 65535}
                v19 := GenS3{F0: 281}
                _ = v18
                _ = v19
            }
        default:
            v4 = 18446744073709551615
        }
        v20 := GenS4{F0: (v1 >> ((42) & 63)), F1: (v2 | (v2 << ((295) & 63)))}
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS5{F0: "hello", F1: 248, F2: (!(uint16(v5) < (^(v1))))}
                    var v24 float64 = float64(0.1)
                    var v25 uint32 = 2147483648
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi21
            }
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v20
    } else {
        if true {
            for fi26 := 0; fi26 < 19; fi26++ {
                if len(v3) > 0 {
                    v3[0] = v2
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 uint = 255
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            var v29 uint64 = v4
            _ = v29
        } else {
            switch v2 {
            case 1:
                for fi30 := 0; fi30 < 12; fi30++ {
                    var v31 int64 = 9223372036854775807
                    v32 := GenS1{F0: (v31 == (v31 ^ (v31 & v31)))}
                    v1 = v1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case 4294967295:
                if (1 != 2147483648) {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        _ = len(v3)
                        switch -2147483648 {
                        case -458:
                            v6 -= -1
                        case 42:
                            for fi34 := 0; fi34 < 2; fi34++ {
                                v35 := GenS3{F0: 43}
                                switch v2 {
                                case -68:
                                    if ("\n\t" > "a") {
                                        v36 := GenS5{F0: "hello", F1: 790, F2: false}
                                        _ = v36
                                    } else {
                                        v4 = (v5 / ((v4) | 1))
                                        v3 = append(v3, (fi34 << ((uint(v4)) & 63)))
                                    }
                                case -447:
                                    if (!false) {
                                        v37 := GenS6{F0: -1, F1: (((!true) && (4294967295 >= 0)) && true), F2: (v1 & v1)}
                                        var v38 bool = (2147483647 > 2147483647)
                                        var v39 uint16 = (v1 ^ (v1 >> ((65535) & 63)))
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    } else {
                                        v1 = (^(v1))
                                    }
                                    v40 := GenS5{F0: "", F1: -9223372036854775808, F2: true}
                                    _ = v40
                                }
                                _ = fi34
                                _ = v35
                            }
                        case 1:
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 int16 = 1
                                var v43 uint = 702
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                        default:
                            var v44 string = ("hello" + "")
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 13; fi45++ {
                            v3 = append(v3, 127)
                            v4 -= v5
                            v6 = 592
                            _ = fi45
                        }
                        _ = fi33
                    }
                }
            }
        }
        if true {
            fmt.Println(v2, v2)
            v5 = (v5 | (v4 % (((^(v5))) | 1)))
            var v46 uint = 4294967295
            _ = v46
        } else {
            v47 := GenS5{F0: "\n\t", F1: -183, F2: (1.5 >= 1e10)}
            var v48 int = (^(v2))
            fmt.Println(v2)
            _ = v47
            _ = v48
        }
        var v49 int64 = -1
        _ = v49
    }
    v4 = v5
    v50 := GenS6{F0: 37, F1: false, F2: genFn0(uint16(v4), v2, (v1 % ((v1) | 1)))}
    v4 -= 875
    var v51 uint16 = uint16(v4)
    var v52 int16 = 32767
    for fi53 := 0; fi53 < 11; fi53++ {
        fmt.Println(v2, v52, v5)
        v54 := GenS7{F0: 100, F1: fi53}
        _ = fi53
        _ = v54
    }
    switch "hello" {
    case "foo bar":
        if true {
            v2 = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v55 int8 = 42
            _ = v55
        } else {
            v56 := GenS8{F0: (v2 ^ -143), F1: 9223372036854775807}
            _ = v56
        }
    case "hello":
        v57 := GenS9{F0: 42, F1: false}
        _ = v57
    }
    v1 = v51
    v58 := GenS7{F0: 1, F1: (v2 >> ((1) & 63))}
    var v59 bool = false
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(!(1.5 >= ((((((0.1 + ((-2.25 - (1e10 + 1e10)) * 1.5)) + ((-2.25 / 1.0) / 0.1)) + 1e10) * -2.25) / (3.14159 * 1e10)) + 0.0)))) {
        fmt.Println(v6, v51, v2)
    } else {
        for fi60 := 0; fi60 < 15; fi60++ {
            var v61 float64 = 0.0
            v62 := GenS10{F0: v2, F1: v1}
            if v59 {
                fmt.Println(v52, v5)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v63 := GenS3{F0: (v51 - 32767)}
                _ = v63
            }
            _ = fi60
            _ = v61
            _ = v62
        }
        v6 = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v58
    _ = v59
    _ = v6
}
