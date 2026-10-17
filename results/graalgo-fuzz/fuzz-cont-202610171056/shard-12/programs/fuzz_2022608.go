package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint8 = 192
    v2 := []int{32767}
    var v3 uint64 = 4294967295
    var v4 uint16 = uint16(v3)
    v5 := make(map[string]int)
    v6 := GenS1{F0: 0, F1: (0 <= 1), F2: false}
    v4 = 2
    if true {
        v7 := GenS1{F0: 827, F1: (!false), F2: true}
        _ = v7
    }
    v8 := GenS2{F0: -1, F1: ("" + (("foo bar" + "hello") + (("a" + "\n\t") + ("\n\t" + ("hello" + "")))))}
    var v9 uint16 = (v4 - v4)
    for fi10 := 0; fi10 < 2; fi10++ {
        switch 428 {
        case 127:
            if (0.1 != 3.14159) {
                if true {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v1 = (v1 + v1)
                        fmt.Println(v4)
                        var v12 uint16 = (v4 / ((v4) | 1))
                        _ = fi11
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        fmt.Println(v4, v4, v1)
                        v14 := GenS2{F0: 2147483647, F1: ("\n\t" + "")}
                        _ = fi13
                        _ = v14
                    }
                    var v15 bool = ((fi10 ^ 290) == (^(fi10)))
                    v16 := v5["\n\t"]
                    _ = v15
                    _ = v16
                }
                switch 128 {
                case -92:
                    if (1000 == -32768) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            v2 = append(v2, (^(fi10)))
                            v18 := GenS3{F0: 65535}
                            _ = fi17
                            _ = v18
                        }
                        v19 := GenS4{F0: "\n\t"}
                        _ = v19
                    } else {
                        delete(v5, "a")
                        switch "a" {
                        case "a":
                            v1 = (v1 * 127)
                        case "foo bar":
                            v9 -= 633
                        case "":
                            var v20 int64 = 513
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 2; fi21++ {
                                var v22 float64 = 3.14159
                                switch fi10 {
                                case -2147483648:
                                    if (true && ((true && false) && (724 >= uint(v4)))) {
                                        switch "foo bar" {
                                        case "a":
                                            var v23 float64 = (v22 + 0.0)
                                            _ = v23
                                        case "":
                                            v24 := GenS5{F0: 2147483648, F1: v1, F2: 1}
                                            _ = v24
                                        case "foo bar":
                                            if (fi10 <= fi21) {
                                                if (("" + "") >= "a") {
                                                    fmt.Println(fi21, v4, fi21)
                                                } else {
                                                    var v25 bool = ((v3 > 42) || (true && (2 < 618)))
                                                    fmt.Println(v4, v4, fi21)
                                                    var v26 string = ("\n\t" + (("" + "a") + "a"))
                                                    _ = v25
                                                    _ = v26
                                                }
                                                v27 := GenS2{F0: 42, F1: ("foo bar" + "hello")}
                                                v22 = (v22 + (v22 / v22))
                                                var v28 bool = (0 <= 42)
                                                _ = v27
                                                _ = v28
                                            } else {
                                                var v29 uint16 = v9
                                                _ = v29
                                            }
                                        default:
                                            if true {
                                                var v30 float64 = v22
                                                _ = v30
                                            } else {
                                                if ((v1 << ((65535) & 63)) > (v1 & v1)) {
                                                    v4 = (v4 / ((v4) | 1))
                                                }
                                            }
                                        }
                                    }
                                case 32768:
                                    switch "\n\t" {
                                    case "a":
                                        v1 += v1
                                    case "hello":
                                        v3 = 0
                                    }
                                }
                                switch "" {
                                case "":
                                    v9 += (^(v9))
                                    switch 32768 {
                                    case 4294967295:
                                        var v31 uint = 255
                                        _ = v31
                                    case -115:
                                        v32 := GenS2{F0: 1, F1: ((("hello" + "\n\t") + ("" + "a")) + "hello")}
                                        _ = v32
                                    case -2147483648:
                                        if true {
                                            switch fi10 {
                                            case -940:
                                                v22 = v22
                                            case 32768:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v33 uint64 = (v3 >> ((44) & 63))
                                                    _ = v33
                                                case "":
                                                    var v34 uint32 = 574
                                                    _ = v34
                                                case "foo bar":
                                                    v4 = (v9 % (((v9 / ((v9) | 1))) | 1))
                                                default:
                                                    v9 += (v4 - (v4 % (((v9 << ((0) & 63))) | 1)))
                                                    var v35 float64 = v22
                                                    v36 := GenS5{F0: 42, F1: (v1 / ((v1) | 1)), F2: uint(v3)}
                                                    _ = v35
                                                    _ = v36
                                                }
                                            default:
                                                v22 = 1e10
                                            }
                                            if (-889 != -32768) {
                                                v5[(("a" + "\n\t") + "")] = 9223372036854775807
                                            }
                                        } else {
                                            v37 := GenS3{F0: 42}
                                            _ = v37
                                        }
                                    default:
                                        if (1 >= (^(v1))) {
                                            var v38 int64 = 0
                                            v3 += v3
                                            _ = v38
                                        }
                                    }
                                    v4 -= v4
                                case "\n\t":
                                    fmt.Println(v1, fi21, fi21)
                                }
                                _ = fi21
                                _ = v22
                            }
                        }
                        v39 := GenS2{F0: 337, F1: ""}
                        _ = v39
                    }
                case 42:
                    switch fi10 {
                    case -858:
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v9 = v9
                            switch ("\n\t" + ("hello" + "\n\t")) {
                            case "a":
                                fmt.Println(fi40)
                            case "hello":
                                v4 = v4
                            }
                            v41 := GenS6{F0: false}
                            var v42 int16 = 32767
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    case 42:
                        for fi43 := 0; fi43 < 2; fi43++ {
                            var v44 bool = (((((fi10 * fi10) > fi43) || (false || (v1 != 2))) && true) && false)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v45 int = (fi43 | (fi43 | fi43))
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    case 1:
                        var v46 uint16 = (v9 * (^(v4)))
                        _ = v46
                    }
                default:
                    var v47 int64 = -663
                    _ = v47
                }
            }
        default:
            delete(v5, ("hello" + ""))
        }
        if len(v2) > 0 {
            v2[0] = (fi10 | -293)
        }
        v48 := GenS7{F0: -655}
        _ = fi10
        _ = v48
    }
    delete(v5, "foo bar")
    var v49 uint8 = v1
    fmt.Println(v9, v3)
    var v50 uint16 = v9
    v4 = (v4 >> ((uint(v1)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v8
    _ = v9
}
