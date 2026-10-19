package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int16, p1 int, p2 uint16) int32 {
    switch "hello" {
    case "a":
        var v4 int16 = p0
        _ = v4
    }
    v5 := GenS1{F0: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}
func genFn1(p0 int32, p1 uint16) int32 {
    p1 = (p1 % ((859) | 1))
    var v6 uint = 42
    var v7 bool = (!false)
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (p0 | p0)
}

func main() {
    var v1 int32 = 42
    v2 := []string{"a", "foo bar", "foo bar", "a"}
    var v3 uint16 = 42
    if (18 == 1) {
        var v9 int8 = -117
        for fi10 := 0; fi10 < 6; fi10++ {
            switch "" {
            case "\n\t":
                if false {
                    v1 = (v1 ^ (v1 % ((-1) | 1)))
                } else {
                    var v11 uint32 = 65535
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v11
                }
            case "a":
                switch (fi10 << ((0) & 63)) {
                case -882:
                    v12 := GenS2{F0: 9223372036854775807}
                    _ = v12
                case 967:
                    for fi13 := 0; fi13 < 3; fi13++ {
                        _ = len(v2)
                        v14 := GenS3{F0: 0, F1: 406, F2: 0}
                        _ = len(v2)
                        var v15 string = ("foo bar" + "\n\t")
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            default:
                if true {
                    var v16 int32 = v1
                    var v17 int = (fi10 | fi10)
                    _ = v16
                    _ = v17
                }
            }
            _ = fi10
        }
        _ = v9
    } else {
        v1 = (v1 | (v1 ^ 42))
        if (!(42 > 1000)) {
            var v18 int64 = 42
            _ = v18
        } else {
            v3 = 255
            var v19 int8 = 47
            v1 = 284
            _ = v19
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v3 = v3
            v21 := GenS4{F0: 339}
            if (22 != 0) {
                if false {
                    var v22 uint = 1
                    _ = v22
                }
            }
            _ = fi20
            _ = v21
        }
    }
    v3 = 491
    fmt.Println(v1)
    switch ("\n\t" + "") {
    case "\n\t":
        var v23 uint16 = (v3 - v3)
        _ = v23
    case "a":
        v3 = (v3 & (v3 ^ v3))
    case "":
        for fi24 := 0; fi24 < 8; fi24++ {
            v3 -= (v3 ^ v3)
            if (!false) {
                if true {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        if (2 >= (v3 - (v3 | (v3 % (((v3 / ((v3) | 1))) | 1))))) {
                            var v26 int64 = 9223372036854775807
                            if false {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    fmt.Println(v26)
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                default:
                                    v27 := GenS3{F0: 2, F1: 42, F2: 2147483647}
                                    _ = v27
                                }
                                v1 -= (v1 >> ((0) & 63))
                                fmt.Println(v1, fi24)
                            } else {
                                var v28 int8 = 100
                                v26 = int64(v3)
                                v29 := GenS4{F0: 443}
                                _ = v28
                                _ = v29
                            }
                            _ = v26
                        }
                        if (!(!((fi25 << ((1) & 63)) < 681))) {
                            v3 -= (v3 / ((v3) | 1))
                            if (!(31 >= 1)) {
                                v30 := GenS5{F0: 0, F1: (!(2 <= 635)), F2: ((("" + "hello") + "hello") + (((("\n\t" + ("foo bar" + "\n\t")) + ("\n\t" + "")) + "") + "a"))}
                                if (275 < 793) {
                                    if ((v1 + (v1 + v1)) != (-(v1))) {
                                        v31 := GenS2{F0: -1}
                                        v3 += (v3 & 0)
                                        fmt.Println(fi24, fi24, v1)
                                        _ = v31
                                    }
                                    v32 := GenS2{F0: 42}
                                    _ = v32
                                } else {
                                    var v33 uint8 = 2
                                    _ = v33
                                }
                                fmt.Println(fi24)
                                v3 += 255
                                v1 = 42
                                _ = v30
                            } else {
                                var v34 int64 = 42
                                _ = v34
                            }
                            if true {
                                switch (^(fi25)) {
                                case 0:
                                    v3 = (v3 << ((2) & 63))
                                case -9223372036854775808:
                                    var v35 uint16 = v3
                                    _ = v35
                                default:
                                    v1 = (-(v1))
                                }
                                v3 += v3
                            }
                        } else {
                            v36 := GenS6{F0: 42}
                            v37 := GenS1{F0: v3}
                            _ = v36
                            _ = v37
                        }
                        _ = fi25
                    }
                    if ((!(!true)) || true) {
                        if (!((100 <= 0) || false)) {
                            if true {
                                v3 = 255
                                for fi38 := 0; fi38 < 11; fi38++ {
                                    v2 = append(v2, ("hello" + "a"))
                                    fmt.Println(v1, v1, v1)
                                    if (359 == 616) {
                                        var v39 int = fi24
                                        switch "a" {
                                        case "a":
                                            v39 -= fi24
                                        case "\n\t":
                                            v40 := GenS6{F0: 42}
                                            _ = v40
                                        }
                                        _ = v39
                                    } else {
                                        var v41 int8 = 89
                                        v41 = v41
                                        _ = v41
                                    }
                                    _ = fi38
                                }
                            } else {
                                v42 := GenS3{F0: 0, F1: 42, F2: 2147483648}
                                _ = v42
                            }
                        }
                        fmt.Println(v3, v3)
                        v1 += v1
                    } else {
                        var v43 bool = true
                        var v44 bool = (!false)
                        v3 = 268
                        _ = v43
                        _ = v44
                    }
                } else {
                    var v45 float64 = 1.5
                    _ = v45
                }
                if true {
                    var v46 bool = (289 != 483)
                    v1 = (v1 * 2147483647)
                    _ = v46
                } else {
                    fmt.Println(fi24, fi24)
                    for fi47 := 0; fi47 < 8; fi47++ {
                        v1 += genFn0(-32768, -475, v3)
                        _ = fi47
                    }
                }
            }
            var v48 uint16 = v3
            _ = fi24
            _ = v48
        }
    default:
        var v49 uint64 = 1
        _ = v49
    }
    var v50 int8 = 42
    if false {
        for fi51 := 0; fi51 < 5; fi51++ {
            var v52 bool = (true && (-1 >= 1))
            for fi53 := 0; fi53 < 16; fi53++ {
                switch -108 {
                case -352:
                    var v54 float64 = ((-2.25 / 1.5) * ((((3.14159 + 0.1) - -2.25) + 3.14159) + 0.0))
                    _ = v54
                case -2147483648:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 1:
                    v55 := GenS5{F0: 1, F1: false, F2: ("foo bar" + "\n\t")}
                    fmt.Println(fi51)
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + "\n\t")
                    }
                    switch "hello" {
                    case "":
                        fmt.Println(v50, fi53)
                    case "a":
                        var v56 string = "hello"
                        _ = v56
                    case "foo bar":
                        switch ("" + "a") {
                        case "a":
                            if v52 {
                                switch "hello" {
                                case "hello":
                                    v57 := GenS1{F0: 32767}
                                    _ = v57
                                }
                            }
                        }
                    }
                    _ = v55
                }
                _ = fi53
            }
            _ = fi51
            _ = v52
        }
        v1 = v1
        fmt.Println(v1)
    } else {
        var v58 float32 = -2.25
        v58 = -2.25
        for fi59 := 0; fi59 < 1; fi59++ {
            var v60 uint32 = 260
            _ = fi59
            _ = v60
        }
        _ = v58
    }
    v1 += v1
    switch 1 {
    case -313:
        v61 := GenS7{F0: -451, F1: 1000}
        _ = v61
    case 629:
        for fi62 := 0; fi62 < 7; fi62++ {
            fmt.Println(v1, fi62)
            v1 -= (v1 % (((v1 + (v1 % (((v1 ^ (^(v1)))) | 1)))) | 1))
            _ = fi62
        }
    case -2147483648:
        v63 := GenS3{F0: 4294967295, F1: 32767, F2: 2147483648}
        _ = v63
    }
    v3 -= (v3 + (v3 + (v3 | 843)))
    for fi64 := 0; fi64 < 8; fi64++ {
        var v65 int16 = -606
        _ = fi64
        _ = v65
    }
    var v66 uint8 = 255
    _ = len(v2)
    fmt.Println(v3, v66)
    v67 := GenS8{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v66
    _ = v67
}
