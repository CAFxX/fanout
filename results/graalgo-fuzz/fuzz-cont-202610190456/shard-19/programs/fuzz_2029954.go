package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32 }
func genFn0() string {
    var v5 uint64 = 186
    v6 := []int{}
    v6 = append(v6, 127)
    fmt.Println(v5)
    var v7 uint64 = (v5 << ((4294967295) & 63))
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() uint8 {
    switch 0 {
    case 294:
        v8 := GenS1{F0: -71}
        _ = v8
    case 1:
        var v9 uint32 = 952
        _ = v9
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            switch ("hello" + (("\n\t" + "hello") + "\n\t")) {
            case "foo bar":
                switch "" {
                case "a":
                    v11 := []int{(fi10 / (((fi10 | fi10)) | 1)), (fi10 >> ((2) & 63))}
                    _ = v11
                }
            default:
                switch (("" + "foo bar") + "a") {
                case "a":
                    fmt.Println(fi10)
                case "hello":
                    if ((fi10 > 835) && false) {
                        fmt.Println(fi10)
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v13 := GenS2{F0: 128}
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        var v14 uint = 0
                        _ = v14
                    }
                case "\n\t":
                    var v15 string = "foo bar"
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                switch (("foo bar" + "a") + "") {
                case "a":
                    switch (fi10 << ((2) & 63)) {
                    case 9223372036854775807:
                        var v17 bool = (false && true)
                        _ = v17
                    case -2147483648:
                        var v18 string = "a"
                        _ = v18
                    default:
                        switch "foo bar" {
                        case "":
                            v19 := GenS3{F0: 212, F1: 42}
                            _ = v19
                        case "foo bar":
                            var v20 int64 = 0
                            v20 -= v20
                            _ = v20
                        }
                    }
                case "foo bar":
                    var v21 string = ""
                    v21 = ("\n\t" + "a")
                    v21 = ""
                    _ = v21
                case "\n\t":
                    var v22 float64 = (1.5 * -2.25)
                    _ = v22
                }
                _ = fi16
            }
            _ = fi10
        }
    }
    switch "foo bar" {
    case "hello":
        var v23 uint16 = 2
        _ = v23
    default:
        var v24 uint32 = 4294967295
        _ = v24
    }
    v25 := []int{789, -534, 1}
    _ = v25
    return 1
}

func main() {
    var v1 int64 = 57
    var v2 uint64 = 0
    var v3 bool = (0 < -52)
    v4 := []uint32{639}
    switch 32767 {
    case 0:
        v1 += (v1 ^ (v1 % ((308) | 1)))
    case 127:
        v1 += v1
    case -406:
        v26 := GenS3{F0: 1, F1: 0}
        _ = v26
    }
    switch "" {
    case "\n\t":
        switch "hello" {
        case "\n\t":
            var v27 string = "hello"
            _ = v27
        case "foo bar":
            v28 := GenS1{F0: -128}
            _ = v28
        case "a":
            v2 = v2
        default:
            for fi29 := 0; fi29 < 12; fi29++ {
                switch fi29 {
                case 32767:
                    v30 := GenS4{F0: 255}
                    _ = v30
                case -741:
                    for fi31 := 0; fi31 < 6; fi31++ {
                        var v32 int8 = 127
                        if (0 < 42) {
                            var v33 int16 = 32767
                            if (65535 < 65535) {
                                if v3 {
                                    var v34 int64 = (v1 % (((v1 >> ((65535) & 63))) | 1))
                                    if (-1 >= -814) {
                                        var v35 float64 = 0.0
                                        v32 = (v32 ^ -57)
                                        _ = v35
                                    }
                                    _ = v34
                                }
                            } else {
                                var v36 uint = 65535
                                _ = v36
                            }
                            v3 = ("" < ((("hello" + ("" + "hello")) + "hello") + ""))
                            _ = v33
                        } else {
                            v2 -= v2
                            if v3 {
                                fmt.Println(v3)
                                v1 = (-(v1))
                                if (!false) {
                                    v2 = v2
                                    if len(v4) > 0 {
                                        v4[0] = 0
                                    }
                                    v3 = v3
                                } else {
                                    v37 := GenS3{F0: uint8(v2), F1: 1}
                                    switch "hello" {
                                    case "":
                                        fmt.Println(fi29, v3)
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 0
                                    }
                                    _ = v37
                                }
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    var v38 string = "a"
                                    _ = v38
                                case "a":
                                    fmt.Println(fi31)
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                case 128:
                    v39 := GenS4{F0: 32767}
                    _ = v39
                default:
                    if v3 {
                        if v3 {
                            v1 = v1
                        } else {
                            var v40 int16 = -32768
                            var v41 bool = true
                            fmt.Println(v2, v40)
                            _ = v40
                            _ = v41
                        }
                    } else {
                        var v42 bool = (!false)
                        _ = v42
                    }
                }
                v1 -= 143
                _ = fi29
            }
        }
    case "":
        fmt.Println(v3, v3, v2)
    }
    v43 := GenS1{F0: 64}
    for fi44 := 0; fi44 < 18; fi44++ {
        if v3 {
            _ = len(v4)
            v45 := GenS5{F0: 2}
            _ = v45
        }
        v3 = (81 > (v1 + -1))
        _ = fi44
    }
    v1 = (^(v1))
    v4 = append(v4, 4294967295)
    if (!true) {
        var v46 int32 = 1
        var v47 int = 32768
        var v48 string = ((("foo bar" + "\n\t") + "hello") + ("a" + "hello"))
        _ = v46
        _ = v47
        _ = v48
    } else {
        v49 := GenS2{F0: 0}
        _ = v49
    }
    var v50 string = "\n\t"
    for fi51 := 0; fi51 < 8; fi51++ {
        v52 := GenS3{F0: 248, F1: 922}
        var v53 int64 = (v1 / (((-(v1))) | 1))
        _ = fi51
        _ = v52
        _ = v53
    }
    v54 := GenS6{F0: 2147483647}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v50
    _ = v54
}
