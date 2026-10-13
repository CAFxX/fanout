package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 uint64, p2 int16) int16 {
    var v4 int64 = -9223372036854775808
    v5 := GenS1{F0: -1}
    v6 := GenS2{F0: p0, F1: (false || (((v4 >> ((0) & 63)) < (v4 << ((65535) & 63))) || false)), F2: 65535}
    if (2147483647 != 354) {
        var v7 uint = 255
        _ = v7
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(p0, p0)
            var v9 int16 = -32768
            for fi10 := 0; fi10 < 4; fi10++ {
                v4 = int64(p1)
                var v11 float32 = -2.25
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v4, p1)
    }
    var v12 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 uint8 = 236
    v2 := make(map[string]string)
    var v3 bool = true
    if true {
        v13 := GenS1{F0: 0}
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    v14 += 1
    v3 = (0.1 < 3.14159)
    v15 := []int{32767, -1}
    fmt.Println(v1)
    switch ("a" + "foo bar") {
    case "hello":
        var v16 uint64 = 1
        _ = v16
    default:
        if false {
            for fi17 := 0; fi17 < 20; fi17++ {
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := GenS2{F0: -2147483648, F1: ("" != "foo bar"), F2: 32767}
                    var v20 int64 = 284
                    switch (^(fi18)) {
                    case 1:
                        v14 += (v20 - v14)
                    default:
                        v14 += v14
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v1)
                _ = fi17
            }
        } else {
            if len(v15) > 0 {
                v15[0] = -374
            }
            var v21 uint16 = 32767
            if v3 {
                v21 -= 727
            } else {
                delete(v2, ("a" + ("" + "hello")))
            }
            _ = v21
        }
    }
    var v22 bool = v3
    var v23 int32 = 828
    _ = len(v15)
    var v24 int64 = (v14 * -1)
    v23 = (v23 << ((4294967295) & 63))
    switch "foo bar" {
    case "hello":
        delete(v2, "")
    case "foo bar":
        switch 853 {
        case 127:
            v14 -= (^(v14))
        case -9223372036854775808:
            for fi25 := 0; fi25 < 16; fi25++ {
                v22 = v3
                _ = fi25
            }
        default:
            v22 = v22
        }
    case "\n\t":
        for fi26 := 0; fi26 < 4; fi26++ {
            if v22 {
                switch "\n\t" {
                case "\n\t":
                    for fi27 := 0; fi27 < 15; fi27++ {
                        switch "a" {
                        case "\n\t":
                            switch fi27 {
                            case 545:
                                switch fi27 {
                                case 348:
                                    if v22 {
                                        v23 = -2147483648
                                        v3 = v22
                                    }
                                case 127:
                                    v28 := v2[((("hello" + "") + "a") + (("hello" + ("\n\t" + ("\n\t" + "a"))) + "hello"))]
                                    _ = v28
                                case 211:
                                    v29 := GenS3{F0: 65535, F1: 1}
                                    _ = v29
                                }
                            }
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                v3 = true
                            case "foo bar":
                                v22 = v22
                            default:
                                v24 = v24
                            }
                        case "foo bar":
                            if false {
                                delete(v2, ("\n\t" + "a"))
                                var v30 uint64 = 153
                                _ = v30
                            } else {
                                v24 = v24
                                var v31 string = "\n\t"
                                if false {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    v32 := GenS1{F0: 42}
                                    v15 = append(v15, (fi27 | (fi27 - -1)))
                                    _ = v32
                                }
                                _ = v31
                            }
                        default:
                            fmt.Println(v3, fi26, fi27)
                        }
                        _ = fi27
                    }
                case "a":
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v22 = true
                        _ = fi33
                    }
                case "":
                    for fi34 := 0; fi34 < 11; fi34++ {
                        var v35 int16 = 1000
                        _ = fi34
                        _ = v35
                    }
                default:
                    v36 := GenS4{F0: ("" + "hello"), F1: (v23 * -2147483648), F2: 102}
                    _ = v36
                }
                var v37 string = "a"
                _ = v37
            } else {
                v22 = (v22 || v22)
            }
            _ = fi26
        }
    default:
        v22 = v22
    }
    v23 -= v23
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
