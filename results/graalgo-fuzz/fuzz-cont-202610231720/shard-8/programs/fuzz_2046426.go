package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }

func main() {
    var v1 float64 = (3.14159 + (-2.25 + 0.1))
    v2 := make(map[uint32]int)
    var v3 float32 = 1e10
    var v4 int64 = -1
    fmt.Println(v3, v1, v3)
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: 496}
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v5, v5)
        var v8 int16 = 827
        _ = fi7
        _ = v8
    }
    if (true && true) {
        if ("foo bar" > "foo bar") {
            switch 0 {
            case 4294967295:
                var v9 int16 = 0
                _ = v9
            case -591:
                v10 := GenS2{F0: (42 <= 42)}
                _ = v10
            case 127:
                v5 = (v5 ^ v5)
            }
            if true {
                if (true && false) {
                    fmt.Println(v5, v4, v1)
                    v4 -= -9223372036854775808
                } else {
                    fmt.Println(v3, v1)
                    v5 = (^(v5))
                    if true {
                        _ = len(v2)
                    } else {
                        fmt.Println(v3, v1, v3)
                    }
                }
                v5 = 1
            } else {
                var v11 uint16 = 255
                if false {
                    v5 -= v5
                    fmt.Println(v11, v3, v1)
                }
                if (!false) {
                    v5 -= (v5 - v5)
                    var v12 int64 = v4
                    _ = v12
                } else {
                    var v13 bool = false
                    var v14 int = 42
                    for fi15 := 0; fi15 < 18; fi15++ {
                        _ = len(v2)
                        v16 := []int{(v14 >> ((4294967295) & 63)), v14}
                        if v13 {
                            fmt.Println(v4, v14)
                            fmt.Println(v13, fi15)
                            var v17 bool = (((4294967295 > fi15) && true) && v13)
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
                _ = v11
            }
            fmt.Println(v1, v4, v5)
        } else {
            fmt.Println(v4, v1)
            for fi18 := 0; fi18 < 8; fi18++ {
                if ((100 == -128) && true) {
                    fmt.Println(v1, v4, v4)
                    v4 = (v4 + int64(fi18))
                } else {
                    var v19 float32 = 3.14159
                    v20 := GenS3{F0: 0}
                    _ = v19
                    _ = v20
                }
                delete(v2, v5)
                v21 := []int{}
                _ = fi18
                _ = v21
            }
            if false {
                v22 := []int{128, -2147483648, -2147483648}
                var v23 string = ""
                _ = v22
                _ = v23
            }
        }
        var v24 uint32 = (v5 * v5)
        _ = v24
    }
    var v25 uint64 = 1
    switch 128 {
    case 1:
        v5 += v5
    case 127:
        v25 -= (v25 ^ (v25 + v25))
    case 42:
        delete(v2, (^(v5)))
    }
    if true {
        for fi26 := 0; fi26 < 3; fi26++ {
            if (7 == 732) {
                fmt.Println(v25, v1, v25)
            }
            var v27 float32 = ((v3 + ((1e10 * (v3 - 0.1)) / -2.25)) + 0.1)
            v28 := GenS2{F0: true}
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        if false {
            var v29 float32 = v3
            _ = v29
        } else {
            if (false || (true || (32767 == 255))) {
                v30 := []int{0}
                if (3.14159 <= (v3 / v3)) {
                    var v31 int8 = 6
                    if true {
                        var v32 uint32 = (v5 % (((v5 << ((4294967295) & 63))) | 1))
                        v3 = 1e10
                        for fi33 := 0; fi33 < 3; fi33++ {
                            if (false && (32767 != 0)) {
                                var v34 int = fi33
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    switch (fi33 >> ((1) & 63)) {
                                    case 1:
                                        fmt.Println(v4, v1)
                                    case 4294967295:
                                        if true {
                                            var v36 int16 = 32767
                                            fmt.Println(fi35, v1)
                                            _ = v36
                                        }
                                    case -2147483648:
                                        if false {
                                            switch ("foo bar" + "") {
                                            case "a":
                                                if (-2147483648 == 621) {
                                                    v37 := v2[v5]
                                                    fmt.Println(fi33)
                                                    _ = v37
                                                }
                                            case "":
                                                var v38 string = ("\n\t" + "")
                                                _ = v38
                                            case "\n\t":
                                                v25 += (^(v25))
                                            }
                                            if (true || true) {
                                                v30 = append(v30, fi33)
                                                if (!true) {
                                                    switch "foo bar" {
                                                    case "":
                                                        v4 = v4
                                                    case "a":
                                                        v2[v5] = (-(fi33))
                                                    default:
                                                        if (45 >= 1) {
                                                            fmt.Println(v32)
                                                        }
                                                    }
                                                    var v39 int8 = 0
                                                    _ = v39
                                                }
                                                var v40 uint16 = 0
                                                _ = v40
                                            } else {
                                                var v41 int32 = 1
                                                _ = v41
                                            }
                                            _ = len(v2)
                                        } else {
                                            fmt.Println(v5, v3, v32)
                                            var v42 int8 = (v31 ^ -1)
                                            var v43 bool = true
                                            _ = v42
                                            _ = v43
                                        }
                                    }
                                    var v44 uint16 = 42
                                    _ = fi35
                                    _ = v44
                                }
                                _ = v34
                            } else {
                                var v45 uint8 = 220
                                _ = v45
                            }
                            if false {
                                v4 = (v4 | (v4 & v4))
                                if true {
                                    if len(v30) >= 2 {
                                        v30 = v30[1:len(v30)]
                                    }
                                    switch (fi33 * fi33) {
                                    case 361:
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            if false {
                                                fmt.Println(v25, v31, v31)
                                            }
                                        }
                                    case -674:
                                        var v46 int64 = 1
                                        _ = v46
                                    case 904:
                                        if false {
                                            if true {
                                                fmt.Println(v31, v5)
                                                switch ("" + ((("" + (("\n\t" + ((((("a" + "a") + (("\n\t" + ("\n\t" + "a")) + "")) + (("" + ("hello" + "\n\t")) + "hello")) + "a") + "\n\t")) + ("hello" + "\n\t"))) + ("\n\t" + ("hello" + "a"))) + "a")) {
                                                case "foo bar":
                                                    var v47 bool = false
                                                    _ = v47
                                                }
                                                if (32767 == -32768) {
                                                    v1 = 1.5
                                                }
                                            } else {
                                                var v48 uint16 = 32767
                                                var v49 int = (fi33 << ((128) & 63))
                                                _ = v48
                                                _ = v49
                                            }
                                        } else {
                                            v4 += (v4 | v4)
                                        }
                                    default:
                                        switch (fi33 & 32768) {
                                        case -616:
                                            for fi50 := 0; fi50 < 2; fi50++ {
                                                var v51 int16 = 42
                                                _ = fi50
                                                _ = v51
                                            }
                                        }
                                    }
                                    var v52 bool = (fi33 >= fi33)
                                    _ = v52
                                } else {
                                    fmt.Println(fi33)
                                }
                                for fi53 := 0; fi53 < 17; fi53++ {
                                    v4 -= 498
                                    _ = fi53
                                }
                            }
                            _ = fi33
                        }
                        _ = v32
                    } else {
                        if (false || false) {
                            if (3.14159 <= (-2.25 / (3.14159 + v1))) {
                                v25 += (^(v25))
                                var v54 uint64 = 150
                                _ = v54
                            } else {
                                switch "a" {
                                case "\n\t":
                                    _ = len(v2)
                                case "a":
                                    v55 := GenS4{F0: (v4 % (((v4 & v4)) | 1)), F1: v25}
                                    _ = v55
                                case "":
                                    switch 4294967295 {
                                    case 4294967295:
                                        var v56 int16 = int16(v5)
                                        _ = v56
                                    case -171:
                                        var v57 bool = ((-905 < -924) && true)
                                        _ = v57
                                    }
                                }
                                v4 = v4
                                for fi58 := 0; fi58 < 14; fi58++ {
                                    for fi59 := 0; fi59 < 12; fi59++ {
                                        v60 := GenS5{F0: fi59}
                                        v4 -= (v4 - (v4 | v4))
                                        _ = fi59
                                        _ = v60
                                    }
                                    delete(v2, v5)
                                    _ = fi58
                                }
                            }
                        }
                    }
                    _ = v31
                } else {
                    if (!true) {
                        var v61 uint64 = 1
                        fmt.Println(v25, v1)
                        _ = v61
                    }
                }
                _ = v30
            }
        }
        switch ("foo bar" + ("\n\t" + "foo bar")) {
        case "hello":
            v62 := GenS5{F0: 9223372036854775807}
            _ = v62
        case "foo bar":
            for fi63 := 0; fi63 < 3; fi63++ {
                v64 := GenS6{F0: -1}
                var v65 int8 = -83
                _ = fi63
                _ = v64
                _ = v65
            }
        }
    }
    v25 += uint64(v5)
    for fi66 := 0; fi66 < 18; fi66++ {
        v5 -= v5
        if ((3.14159 - 3.14159) == (v1 * 0.1)) {
            var v67 int = (fi66 & (fi66 * (fi66 & fi66)))
            v68 := GenS1{F0: 0}
            v69 := GenS7{F0: ("\n\t" + (("a" + "a") + "a"))}
            _ = v67
            _ = v68
            _ = v69
        }
        _ = fi66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
