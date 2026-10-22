package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint; F2 bool }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool; F1 uint; F2 uint16 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int; F1 uint; F2 uint32 }

func main() {
    var v1 float64 = float64((0.0 + 3.14159))
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 int = -725
    var v5 uint = 18446744073709551615
    var v6 string = ("\n\t" + ("" + "a"))
    v4 = (v4 * v4)
    v2[9223372036854775807] = 2147483647
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v6 = "a"
            v9 := GenS1{F0: 32767}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 65535}
        var v11 float64 = (v1 + 0.0)
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 int32 = 532
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 = 373
    var v13 bool = true
    var v14 int32 = (v12 * (v12 + (v12 * v12)))
    v15 := GenS1{F0: 313}
    for fi16 := 0; fi16 < 13; fi16++ {
        delete(v2, fi16)
        for fi17 := 0; fi17 < 12; fi17++ {
            switch ((v6 + v6) + ((v6 + "\n\t") + v6)) {
            case "":
                v12 = (v14 ^ (v14 % (((v14 >> (((v5 | v5)) & 63))) | 1)))
            case "\n\t":
                var v18 uint64 = 2
                _ = v18
            }
            _ = fi17
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            v2[128] = 2147483647
            _ = fi19
        }
        _ = fi16
    }
    switch (-(v4)) {
    case 127:
        if (v13 && v13) {
            if v13 {
                for fi20 := 0; fi20 < 12; fi20++ {
                    if (543 <= 304) {
                        switch v6 {
                        case "foo bar":
                            v21 := GenS2{F0: v4}
                            _ = v21
                        case "hello":
                            if (!(70 > 42)) {
                                v5 = (v5 | 4294967295)
                            } else {
                                var v22 float32 = 1e10
                                switch ((((v6 + v6) + "\n\t") + ((v6 + v6) + (v6 + "a"))) + v6) {
                                case "\n\t":
                                    var v23 int16 = 0
                                    _ = v23
                                case "foo bar":
                                    _ = len(v3)
                                }
                                _ = v22
                            }
                        }
                        fmt.Println(v6, v12, v14)
                        if v13 {
                            v3 = append(v3, (fi20 << ((v5) & 63)))
                            v24 := GenS3{F0: 0, F1: v5, F2: false}
                            for fi25 := 0; fi25 < 8; fi25++ {
                                v4 = (fi20 | v4)
                                _ = fi25
                            }
                            _ = v24
                        } else {
                            v12 -= (-(v14))
                            v5 = v5
                        }
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        fmt.Println(v1, v1)
                        delete(v2, (v4 / (((fi26 - (v4 % (((fi20 ^ fi20)) | 1)))) | 1)))
                        _ = fi26
                    }
                    for fi27 := 0; fi27 < 5; fi27++ {
                        delete(v2, (fi27 / ((0) | 1)))
                        _ = fi27
                    }
                    _ = fi20
                }
                if v13 {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        fmt.Println(v4)
                        _ = fi28
                    }
                }
            }
            var v29 uint32 = 2147483648
            _ = v29
        } else {
            v30 := GenS4{F0: 0, F1: (^(v5))}
            var v31 string = v6
            switch (v31 + "") {
            case "hello":
                if (true && v13) {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        v33 := v2[fi32]
                        v34 := GenS4{F0: (fi32 << (((^(v5))) & 63)), F1: (v5 >> ((v5) & 63))}
                        switch (v4 * 0) {
                        case 1:
                            switch ((((v6 + "hello") + "\n\t") + v6) + ("foo bar" + v31)) {
                            case "hello":
                                if ((v14 % ((v12) | 1)) != (^(v14))) {
                                    v35 := GenS5{F0: false, F1: v5, F2: 2}
                                    _ = v35
                                } else {
                                    fmt.Println(v1, v6, v13)
                                    v36 := GenS6{F0: 127, F1: 0}
                                    _ = v36
                                }
                            case "a":
                                var v37 uint64 = 1
                                _ = v37
                            case "":
                                if true {
                                    var v38 uint = (v5 & (v5 & (v5 / ((860) | 1))))
                                    var v39 float64 = v1
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        if (("a" + v31) < v6) {
                                            v39 += 0.0
                                            var v41 uint16 = 65535
                                            var v42 uint = v5
                                            _ = v41
                                            _ = v42
                                        }
                                        v43 := GenS5{F0: (888 > 9223372036854775807), F1: (v38 & v38), F2: 737}
                                        v39 -= (v39 * ((1e10 * 1e10) / (((((v39 * v39) + -2.25) * v1) + (((0.1 - v39) + 3.14159) / 0.1)) + float64((3.14159 + 1.5)))))
                                        _ = fi40
                                        _ = v43
                                    }
                                    _ = v38
                                    _ = v39
                                } else {
                                    if (false || (3.14159 <= float64(1.5))) {
                                        v44 := GenS7{F0: 680, F1: v5, F2: 360}
                                        var v45 int16 = int16(fi32)
                                        _ = v44
                                        _ = v45
                                    }
                                    v3 = append(v3, (-(fi32)))
                                }
                            }
                        case -1:
                            v46 := GenS6{F0: (^(fi32)), F1: -106}
                            _ = v46
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        if (-1 != -26) {
                            v4 = v4
                        }
                    case "foo bar":
                        fmt.Println(v1)
                    }
                }
            default:
                for fi47 := 0; fi47 < 4; fi47++ {
                    v48 := GenS5{F0: false, F1: (^(v5)), F2: 32767}
                    v49 := GenS5{F0: v13, F1: v5, F2: 32767}
                    for fi50 := 0; fi50 < 14; fi50++ {
                        var v51 uint8 = 135
                        var v52 int = (fi47 | 346)
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
            _ = v30
            _ = v31
        }
    case 128:
        if (v13 && (-2.25 >= 0.1)) {
            switch v6 {
            case "foo bar":
                var v53 int16 = -32768
                _ = v53
            case "a":
                delete(v2, v4)
            default:
                v54 := GenS3{F0: 32767, F1: (v5 >> ((630) & 63)), F2: (((false && false) || (!v13)) || false)}
                _ = v54
            }
        }
    case 2147483647:
        var v55 bool = true
        _ = v55
    default:
        var v56 uint8 = 2
        _ = v56
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
