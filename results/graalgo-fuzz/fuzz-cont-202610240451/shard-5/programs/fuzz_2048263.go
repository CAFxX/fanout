package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 int16; F1 bool; F2 int }
func genFn0() bool {
    var v5 int8 = -128
    v6 := GenS1{F0: 65535, F1: "hello"}
    fmt.Println(v5)
    if false {
        if (-26 > 4294967295) {
            var v7 float64 = ((1.5 - float64((1.5 * -2.25))) - 3.14159)
            var v8 uint8 = 83
            _ = v7
            _ = v8
        }
        if (2 != 18446744073709551615) {
            v5 -= 127
            var v9 uint64 = 355
            _ = v9
        }
        switch ("hello" + "") {
        case "":
            v10 := GenS2{F0: 1, F1: 255}
            _ = v10
        case "\n\t":
            v5 = v5
        case "a":
            switch -1 {
            case 2147483647:
                if false {
                    v5 = v5
                } else {
                    v11 := GenS1{F0: 2147483648, F1: ""}
                    v12 := GenS1{F0: 65535, F1: ((("\n\t" + "a") + "a") + "")}
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 uint = 65535
                        _ = fi13
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
            }
        default:
            v5 -= (v5 * v5)
        }
    }
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint, p1 uint8) uint {
    fmt.Println(p1)
    for fi15 := 0; fi15 < 1; fi15++ {
        switch fi15 {
        case 2147483647:
            v16 := GenS2{F0: 18446744073709551615, F1: 42}
            _ = v16
        case 298:
            var v17 int64 = 42
            _ = v17
        case 1:
            v18 := []int{32768}
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS3{F0: 2, F1: (p1 ^ p1)}
            var v21 uint64 = 967
            var v22 uint8 = (p1 << ((p0) & 63))
            p0 = 534
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi15
    }
    p1 = (p1 | p1)
    p0 = p0
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 1
    var v3 int32 = -2147483648
    v4 := make(map[string]string)
    v23 := GenS1{F0: (v1 & (v1 ^ (v1 - (v1 ^ (^(v1)))))), F1: ("a" + "a")}
    var v24 int = 423
    var v25 string = ("\n\t" + "a")
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 float64 = -2.25
        var v28 string = "\n\t"
        fmt.Println(v25, v2, v27)
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 = (^(v3))
    if true {
        v29 := GenS4{F0: v25, F1: -9223372036854775808, F2: 18446744073709551615}
        switch v24 {
        case -659:
            if false {
                var v30 uint64 = 42
                for fi31 := 0; fi31 < 9; fi31++ {
                    v24 = fi31
                    v24 -= v24
                    v4[""] = v25
                    _ = fi31
                }
                var v32 int = 42
                _ = v30
                _ = v32
            }
        case -769:
            v33 := GenS5{F0: 0, F1: (^(v2)), F2: 1}
            _ = v33
        case 32768:
            var v34 int = v24
            _ = v34
        }
        var v35 int8 = -128
        _ = v29
        _ = v35
    }
    v3 = v3
    v25 = ((v25 + v25) + v25)
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(fi36)
        v25 = v25
        v1 += 2
        _ = fi36
    }
    v37 := GenS6{F0: 112, F1: 65535, F2: 100}
    if false {
        for fi38 := 0; fi38 < 3; fi38++ {
            for fi39 := 0; fi39 < 9; fi39++ {
                v40 := v4[v25]
                switch ((v25 + v25) + v40) {
                case "":
                    v3 = v3
                case "\n\t":
                    v3 = (v3 - v3)
                default:
                    if (true && (1 < -1)) {
                        if true {
                            if (-1 > -1) {
                                v3 += (v3 | (v3 * -1))
                                v41 := GenS7{F0: (v3 << ((genFn1(genFn1(2, v2), v2)) & 63)), F1: 0}
                                switch fi39 {
                                case 42:
                                    var v42 int32 = (v3 + v3)
                                    _ = v42
                                case 127:
                                    if (v3 <= (v3 & (v3 - (v3 ^ v3)))) {
                                        switch ((v40 + "a") + "hello") {
                                        case "hello":
                                            var v43 float64 = (3.14159 - 1.5)
                                            _ = v43
                                        case "a":
                                            var v44 string = "foo bar"
                                            _ = v44
                                        case "":
                                            v40 = "\n\t"
                                        default:
                                            var v45 int8 = 59
                                            _ = v45
                                        }
                                    }
                                }
                                switch (fi39 ^ (-(fi39))) {
                                case 2147483647:
                                    delete(v4, v40)
                                case 32768:
                                    v46 := GenS1{F0: (v1 - 2147483647), F1: v25}
                                    _ = v46
                                }
                                _ = v41
                            } else {
                                fmt.Println(v40, fi39)
                                switch v40 {
                                case "foo bar":
                                    var v47 float32 = (1.5 / (0.1 + (0.1 * 1.5)))
                                    _ = v47
                                case "":
                                    if true {
                                        fmt.Println(v3, v2)
                                        if (2 == 301) {
                                            var v48 int = 9223372036854775807
                                            _ = v48
                                        } else {
                                            var v49 string = (v25 + (("a" + ((v40 + v25) + ((v25 + ((v25 + "") + "hello")) + ("a" + "\n\t")))) + v40))
                                            fmt.Println(v2, fi38)
                                            switch (v49 + "") {
                                            case "\n\t":
                                                if (true || (v49 > "foo bar")) {
                                                    v50 := GenS8{F0: 9223372036854775807, F1: 42}
                                                    v1 += v1
                                                    v51 := GenS9{F0: (v3 + (v3 + 42)), F1: -9223372036854775808}
                                                    _ = v50
                                                    _ = v51
                                                }
                                            default:
                                                v52 := []int{}
                                                _ = v52
                                            }
                                            _ = v49
                                        }
                                        v3 += v3
                                    } else {
                                        var v53 uint64 = 0
                                        _ = v53
                                    }
                                case "a":
                                    v54 := GenS10{F0: 79}
                                    _ = v54
                                default:
                                    v25 = "\n\t"
                                }
                                var v55 uint32 = v1
                                _ = v55
                            }
                            var v56 int32 = 718
                            v57 := []int{fi38, fi39}
                            _ = v56
                            _ = v57
                        }
                        var v58 int64 = -746
                        v59 := []int{(-(fi38)), (fi38 >> ((255) & 63)), (fi38 * (fi38 % (((fi38 % ((4294967295) | 1))) | 1)))}
                        var v60 bool = (fi38 == fi39)
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                }
                _ = fi39
                _ = v40
            }
            v2 = 42
            _ = fi38
        }
    }
    var v61 uint16 = 545
    fmt.Println(v2, v24, v2)
    var v62 uint8 = 0
    if (18446744073709551615 < 18446744073709551615) {
        v63 := GenS11{F0: 591, F1: 255}
        for fi64 := 0; fi64 < 8; fi64++ {
            var v65 uint32 = v1
            _ = fi64
            _ = v65
        }
        v66 := GenS1{F0: v1, F1: "hello"}
        _ = v63
        _ = v66
    } else {
        for fi67 := 0; fi67 < 20; fi67++ {
            v68 := GenS7{F0: (^(v3)), F1: (v61 >> ((348) & 63))}
            _ = fi67
            _ = v68
        }
        switch (v25 + "hello") {
        case "foo bar":
            switch "foo bar" {
            case "":
                for fi69 := 0; fi69 < 15; fi69++ {
                    v70 := GenS3{F0: 1, F1: v62}
                    _ = fi69
                    _ = v70
                }
            default:
                delete(v4, ((("a" + "a") + ("" + (v25 + v25))) + v25))
            }
        case "\n\t":
            var v71 uint32 = 2147483647
            _ = v71
        case "":
            v72 := GenS12{F0: -1, F1: true, F2: v24}
            _ = v72
        default:
            fmt.Println(v2)
        }
        switch v24 {
        case 2147483647:
            switch v24 {
            case 128:
                var v73 uint16 = (v61 ^ (v61 & 1))
                _ = v73
            }
        default:
            v1 -= v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v61
    _ = v62
}
