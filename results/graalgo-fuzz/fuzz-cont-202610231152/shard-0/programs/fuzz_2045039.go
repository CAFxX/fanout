package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16; F1 int; F2 string }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 string; F1 string; F2 uint16 }
type GenS6 struct { F0 uint }

func main() {
    var v1 string = (("hello" + "hello") + "foo bar")
    v2 := []string{"\n\t", (("hello" + v1) + (v1 + "")), (v1 + (v1 + "foo bar"))}
    var v3 int = 438
    var v4 int16 = 324
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: -426}
        fmt.Println(v3, v1, fi5)
        v7 := GenS1{F0: v4}
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v2) > 0 {
                v2[0] = v1
            }
            var v10 string = v1
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v3 -= v3
    var v11 bool = false
    v1 = "foo bar"
    switch (v3 - (-(v3))) {
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int32 = 2147483647
            v14 := GenS2{F0: 202, F1: (v13 - -2147483648)}
            switch v1 {
            case "hello":
                if ((v3 & v3) > fi12) {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v1 = "hello"
                        v16 := GenS3{F0: v4, F1: v3, F2: v1}
                        var v17 int16 = (v4 >> ((390) & 63))
                        fmt.Println(fi15)
                        var v18 uint32 = 2147483648
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v4 += (v4 * -32768)
                }
            case "":
                if (-20 > -1) {
                    var v19 string = (v1 + "hello")
                    _ = v19
                } else {
                    v4 = (v4 + (v4 ^ (v4 << ((482) & 63))))
                }
            case "\n\t":
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v1, v1, v4)
    var v21 bool = false
    if ((!v21) || false) {
        if v21 {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint32 = 2147483647
                v3 = fi22
                _ = fi22
                _ = v23
            }
            var v24 int = (^(v3))
            _ = v24
        } else {
            _ = len(v2)
            v2 = append(v2, v1)
            if (126 > 33) {
                v21 = (1 <= -9223372036854775808)
            }
        }
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        _ = len(v2)
        if v11 {
            if v21 {
                if true {
                    fmt.Println(v1, v3, fi25)
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v21 = v21
                        _ = fi26
                    }
                    v11 = false
                }
            } else {
                v27 := GenS3{F0: (v4 << ((328) & 63)), F1: v3, F2: v1}
                _ = v27
            }
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int = 190
                if v21 {
                    if (!true) {
                        if false {
                            var v30 uint64 = 42
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                v30 += 18446744073709551615
                                v29 -= v3
                            }
                            v11 = (2147483648 >= 2)
                            if true {
                                switch -212 {
                                case 127:
                                    v30 = 317
                                case 128:
                                    if (-77 < 2) {
                                        if true {
                                            if (196 >= 42) {
                                                _ = len(v2)
                                            } else {
                                                v2 = append(v2, v1)
                                            }
                                        } else {
                                            v31 := GenS4{F0: ((v1 + v1) + ""), F1: (v4 ^ (v4 + (v4 & (v4 - v4))))}
                                            _ = v31
                                        }
                                    }
                                }
                                v32 := GenS5{F0: "", F1: (v1 + v1), F2: 411}
                                v2 = append(v2, v1)
                                v29 = 996
                                v21 = v11
                                switch (v3 | -534) {
                                case -1:
                                    fmt.Println(v1)
                                case 352:
                                    v1 = ("" + v1)
                                case 128:
                                    switch ("foo bar" + "a") {
                                    case "\n\t":
                                        v33 := GenS6{F0: 18446744073709551615}
                                        _ = v33
                                    }
                                }
                                if true {
                                    if (!v11) {
                                        var v34 int32 = 578
                                        v35 := GenS1{F0: (v4 & v4)}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v29 = (v29 | fi25)
                                    }
                                    var v36 int = (v29 & 42)
                                    _ = v36
                                }
                                _ = v32
                            }
                            _ = v30
                        } else {
                            var v37 int64 = 0
                            v1 = (v1 + v1)
                            var v38 uint = 1
                            _ = v37
                            _ = v38
                        }
                        v21 = (!true)
                    } else {
                        fmt.Println(v1, fi25, fi25)
                    }
                    var v39 int16 = v4
                    switch (fi25 / ((fi25) | 1)) {
                    case 4294967295:
                        if (((v39 & (v4 % ((0) | 1))) != v39) && v21) {
                            v40 := GenS1{F0: -983}
                            _ = v40
                        }
                    case 32767:
                        var v41 int16 = v39
                        _ = v41
                    case 2147483647:
                        if v11 {
                            fmt.Println(fi25, v21)
                            v4 += v4
                            fmt.Println(v11, v11, v1)
                        }
                    }
                    v42 := GenS4{F0: ((v1 + v1) + v1), F1: v4}
                    var v43 int16 = (-(v4))
                    v11 = false
                    var v44 int8 = -124
                    v45 := GenS1{F0: (v4 + v43)}
                    v46 := GenS2{F0: 42, F1: 998}
                    var v47 uint16 = 12
                    _ = v39
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                    _ = v47
                }
                _ = fi28
                _ = v29
            }
        } else {
            v2 = append(v2, (v1 + v1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, "a")
        }
        var v48 uint16 = 1
        _ = fi25
        _ = v48
    }
    fmt.Println(v1, v1)
    for fi49 := 0; fi49 < 6; fi49++ {
        v4 = 0
        v50 := GenS1{F0: (-(v4))}
        switch (-(fi49)) {
        case 4294967295:
            v11 = (!false)
        case 802:
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 bool = ("hello" >= v1)
                _ = fi51
                _ = v52
            }
        case 394:
            switch v1 {
            case "a":
                var v53 int8 = -128
                _ = v53
            case "\n\t":
                v11 = (v4 != v4)
            case "foo bar":
                var v54 int32 = 0
                _ = v54
            }
        default:
            for fi55 := 0; fi55 < 20; fi55++ {
                v21 = true
                v56 := GenS5{F0: ("a" + v1), F1: "hello", F2: 680}
                _ = fi55
                _ = v56
            }
        }
        _ = fi49
        _ = v50
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        if len(v2) > 0 {
            v2[0] = (v1 + (v1 + v1))
        }
        _ = len(v2)
        v4 = v4
        _ = fi57
    }
    var v58 float32 = 0.1
    var v59 string = (("hello" + "hello") + ("hello" + (("foo bar" + "a") + (v1 + ("foo bar" + "")))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v58
    _ = v59
}
