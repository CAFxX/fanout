package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int; F1 int8; F2 string }
func genFn0(p0 int64, p1 uint, p2 bool) bool {
    p0 = p0
    var v7 float64 = (1e10 + 3.14159)
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "" {
        case "":
            v9 := GenS1{F0: 65535}
            _ = v9
        case "hello":
            var v10 int16 = 1000
            _ = v10
        case "foo bar":
            if false {
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 float32 = (0.0 * 1e10)
                    _ = fi11
                    _ = v12
                }
                var v13 int = fi8
                var v14 uint = 0
                _ = v13
                _ = v14
            }
        default:
            v15 := []int{}
            _ = v15
        }
        var v16 bool = p2
        _ = fi8
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 int, p1 int8) uint16 {
    v17 := []int{-9223372036854775808, (-(p0)), 810, 826}
    var v18 float32 = (0.0 - float32(0.1))
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 566
}

func main() {
    var v1 int = 181
    var v2 uint8 = 42
    var v3 uint16 = 32767
    var v4 int16 = -1
    v5 := make(map[uint32]bool)
    v6 := []string{}
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint8 = v2
            v22 := GenS2{F0: 4294967295}
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            if (false || (42 > v2)) {
                v2 += (v2 / ((v2) | 1))
            }
            switch 42 {
            case -1:
                v6 = append(v6, "a")
            default:
                if false {
                    var v24 string = ""
                    _ = len(v5)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v25 int32 = -2147483648
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS2{F0: 0}
                    v2 += 11
                    _ = v26
                }
            }
            v3 -= (^(v3))
            _ = fi23
        }
        if len(v6) > 0 {
            v6[0] = (("foo bar" + "") + ("hello" + "foo bar"))
        }
        _ = fi19
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS2{F0: 42}
        v3 -= (v3 + (v3 << ((1) & 63)))
        var v29 uint16 = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    delete(v5, 1)
    fmt.Println(v3, v2)
    switch v1 {
    case 700:
        var v30 string = "\n\t"
        _ = v30
    case -495:
        v6 = append(v6, ("\n\t" + "\n\t"))
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 int8 = 42
        _ = fi31
        _ = v32
    }
    if (!true) {
        fmt.Println(v2)
        switch (("a" + "") + "a") {
        case "a":
            v33 := GenS3{F0: uint(v3), F1: ("" + ((("" + "foo bar") + "foo bar") + "foo bar"))}
            _ = v33
        case "hello":
            v34 := GenS4{F0: 128}
            _ = v34
        case "foo bar":
            v2 -= (v2 >> ((0) & 63))
        default:
            if (!(false || (4294967295 != uint32(v2)))) {
                delete(v5, 2147483648)
            }
        }
        if false {
            v35 := GenS4{F0: 128}
            var v36 uint64 = 362
            fmt.Println(v3, v36, v2)
            _ = v35
            _ = v36
        }
    }
    if true {
        v37 := GenS5{F0: (v3 ^ v3), F1: 2147483647}
        if false {
            v5[189] = (2147483647 >= 2)
            v2 += (v2 % (((v2 % (((v2 & v2)) | 1))) | 1))
        } else {
            for fi38 := 0; fi38 < 2; fi38++ {
                v3 -= (v3 - genFn1(fi38, -128))
                _ = fi38
            }
            var v39 int8 = -128
            var v40 uint8 = (v2 + (v2 % (((v2 >> ((2) & 63))) | 1)))
            _ = v39
            _ = v40
        }
        _ = v37
    }
    switch ("hello" + ("\n\t" + "\n\t")) {
    case "hello":
        var v41 uint64 = 0
        _ = v41
    case "a":
        switch "" {
        case "hello":
            v3 = v3
        case "\n\t":
            v2 += (^(v2))
        }
    case "":
        fmt.Println(v4)
    default:
        var v42 uint = 677
        _ = v42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        fmt.Println(fi43)
        if ((v1 & (v1 * 2147483647)) >= (fi43 | (v1 | (fi43 % ((v1) | 1))))) {
            v4 = v4
        } else {
            v44 := GenS1{F0: uint16(v2)}
            var v45 uint16 = uint16(v2)
            fmt.Println(v1, v2)
            _ = v44
            _ = v45
        }
        _ = fi43
    }
    v46 := GenS3{F0: 770, F1: "foo bar"}
    for fi47 := 0; fi47 < 3; fi47++ {
        var v48 uint64 = 18446744073709551615
        var v49 int8 = 1
        if (v4 > v4) {
            v50 := GenS2{F0: uint32(v3)}
            if (false || true) {
                if (140 <= v2) {
                    v51 := GenS2{F0: 0}
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 17; fi52++ {
                        if (true && true) {
                            var v53 uint32 = 0
                            fmt.Println(fi52, v48)
                            var v54 float64 = (-2.25 * 1e10)
                            _ = v53
                            _ = v54
                        }
                        var v55 int64 = -1
                        v5[648] = false
                        _ = fi52
                        _ = v55
                    }
                }
                v3 = v3
                fmt.Println(v3, v48, v49)
            } else {
                var v56 int16 = v4
                _ = v56
            }
            _ = v50
        } else {
            switch v1 {
            case 968:
                v57 := GenS3{F0: 0, F1: "a"}
                _ = v57
            case 923:
                var v58 uint16 = v3
                _ = v58
            }
            v59 := GenS6{F0: 4294967295, F1: v49, F2: (("foo bar" + "foo bar") + "")}
            _ = v59
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    switch ("hello" + "a") {
    case "":
        for fi60 := 0; fi60 < 5; fi60++ {
            for fi61 := 0; fi61 < 17; fi61++ {
                fmt.Println(v4, v2, fi61)
                _ = fi61
            }
            v3 = v3
            v3 += (v3 + v3)
            _ = fi60
        }
    case "\n\t":
        if false {
            fmt.Println(v2, v2)
        }
    case "hello":
        var v62 bool = true
        _ = v62
    }
    var v63 int = int(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v63
}
