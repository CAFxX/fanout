package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 uint16) uint64 {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint32 = 329
            v7 += (v7 ^ (v7 << ((255) & 63)))
            p0 += fi5
            var v8 uint32 = (v7 - (v7 / (((v7 + (v7 ^ (v7 * (v7 & 0))))) | 1)))
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    v9 := []int{4294967295}
    if (!(false || true)) {
        p2 = 984
    }
    var v10 uint64 = 952
    if false {
        var v11 bool = true
        var v12 int = (p0 & (p0 % ((p0) | 1)))
        _ = v11
        _ = v12
    } else {
        if true {
            v13 := GenS1{F0: 127, F1: -128, F2: 255}
            _ = v13
        } else {
            var v14 int8 = -128
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 uint32 = 0
    var v2 string = ("a" + ("foo bar" + ""))
    var v3 int64 = -108
    v4 := []int{0, 272, 9223372036854775807}
    if (!((!(1 > 1)) || ((v1 <= (v1 / ((v1) | 1))) && true))) {
        var v15 uint16 = 629
        v2 = ""
        fmt.Println(v2, v3, v1)
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: -68, F2: 18446744073709551615}
    if false {
        var v17 uint = 65535
        if ((0.1 * 3.14159) < -2.25) {
            switch v2 {
            case "a":
                var v18 uint8 = 127
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 16; fi19++ {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        if true {
                            var v21 bool = true
                            if (!(727 != 2)) {
                                v2 = v2
                            } else {
                                if (!(127 == 210)) {
                                    v1 = (^(v1))
                                    if (v2 >= (("\n\t" + ("" + v2)) + (v2 + (v2 + "")))) {
                                        v4 = append(v4, 32768)
                                        switch fi20 {
                                        case 42:
                                            v3 = v3
                                        default:
                                            fmt.Println(v3, v17, v17)
                                        }
                                        var v22 uint8 = 127
                                        _ = v22
                                    } else {
                                        switch fi19 {
                                        case -496:
                                            switch (^(fi19)) {
                                            case -1:
                                                v2 = (v2 + (("hello" + v2) + ((((v2 + (v2 + ("foo bar" + ""))) + v2) + (v2 + v2)) + ("a" + ("" + ("" + "foo bar"))))))
                                            case 12:
                                                var v23 int8 = 67
                                                _ = v23
                                            default:
                                                v24 := GenS1{F0: -1, F1: -17, F2: (^(v17))}
                                                _ = v24
                                            }
                                        }
                                        switch int(v17) {
                                        case 0:
                                            var v25 uint64 = 1
                                            v26 := GenS2{F0: v3}
                                            fmt.Println(v3, v1)
                                            _ = v25
                                            _ = v26
                                        case 378:
                                            v21 = ((v1 * (v1 << ((0) & 63))) < v1)
                                        default:
                                            v27 := GenS2{F0: v3}
                                            _ = v27
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v4 = append(v4, (fi19 / ((fi19) | 1)))
                                }
                                v2 = "\n\t"
                                var v28 int64 = int64(fi19)
                                _ = v28
                            }
                            switch (v2 + (((v2 + v2) + ((v2 + "foo bar") + (v2 + v2))) + "hello")) {
                            case "\n\t":
                                var v29 uint16 = 65535
                                _ = v29
                            case "foo bar":
                                fmt.Println(v2, v2, v1)
                                if true {
                                    var v30 bool = false
                                    v3 = (v3 ^ 0)
                                    _ = v30
                                } else {
                                    switch 42 {
                                    case -1:
                                        var v31 float32 = (((1e10 * 1.5) + 1e10) / float32(1.5))
                                        _ = v31
                                    }
                                }
                                v2 = "\n\t"
                            }
                            if len(v4) > 0 {
                                v4[0] = fi19
                            }
                            _ = v21
                        }
                        _ = fi20
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v17 = v17
                    _ = fi19
                }
            case "":
                for fi32 := 0; fi32 < 20; fi32++ {
                    v17 = v17
                    v3 = (v3 * -322)
                    var v33 int8 = 1
                    _ = fi32
                    _ = v33
                }
            default:
                if ((^(v1)) > (v1 - 2147483648)) {
                    switch 127 {
                    case 32767:
                        v2 = "a"
                    }
                }
            }
            v2 = "a"
        } else {
            fmt.Println(v17, v1)
            v1 -= v1
        }
        _ = v17
    } else {
        v34 := GenS2{F0: (v3 >> ((671) & 63))}
        if false {
            var v35 float64 = 3.14159
            _ = v35
        }
        _ = v34
    }
    var v36 int16 = 32767
    var v37 uint8 = 128
    v2 = ("hello" + ("\n\t" + ""))
    if true {
        for fi38 := 0; fi38 < 14; fi38++ {
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 uint8 = (v37 | uint8(v1))
                v41 := GenS3{F0: fi38}
                var v42 uint64 = 2
                v43 := GenS1{F0: int8(fi38), F1: 127, F2: 2}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
                _ = v43
            }
            _ = fi38
        }
        var v44 uint8 = v37
        if (true && false) {
            var v45 int64 = (v3 ^ v3)
            switch 1 {
            case -9223372036854775808:
                var v46 float64 = 1.5
                _ = v46
            }
            v47 := GenS1{F0: -128, F1: 100, F2: 4294967295}
            _ = v45
            _ = v47
        }
        _ = v44
    } else {
        if (!(!false)) {
            var v48 int = -863
            if len(v4) > 0 {
                v4[0] = (v48 * 128)
            }
            v1 -= (v1 | 1)
            _ = v48
        }
        v4 = append(v4, 127)
        v2 = ((("foo bar" + v2) + "foo bar") + v2)
    }
    v49 := GenS4{F0: 678}
    v1 -= (v1 << ((1) & 63))
    for fi50 := 0; fi50 < 18; fi50++ {
        v51 := GenS1{F0: 0, F1: int8(v3), F2: 1}
        var v52 uint = 542
        _ = fi50
        _ = v51
        _ = v52
    }
    v53 := GenS5{F0: ""}
    var v54 int16 = (^(v36))
    v54 -= v36
    var v55 int = 42
    var v56 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v49
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
