package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0() uint8 {
    if (42 <= 2147483648) {
        v6 := GenS1{F0: 2}
        _ = v6
    }
    var v7 uint32 = 4294967295
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 bool = (true || false)
        v10 := GenS1{F0: 1}
        v11 := GenS2{F0: 406, F1: 255}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v7)
        v13 := []int{(fi12 >> ((1) & 63))}
        v7 += (v7 ^ (v7 & (^(v7))))
        _ = fi12
        _ = v13
    }
    v7 = v7
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 uint8 = 3
    var v4 uint = 30
    var v5 uint16 = 226
    var v14 int16 = int16(v5)
    v3 = uint8(v5)
    var v15 int8 = 100
    _ = len(v1)
    fmt.Println(v4, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v16 int16 = (v14 / ((-32768) | 1))
    switch ("hello" + ("a" + ("hello" + "foo bar"))) {
    case "":
        for fi17 := 0; fi17 < 18; fi17++ {
            switch "a" {
            case "":
                var v18 uint64 = 18446744073709551615
                _ = v18
            case "a":
                if true {
                    switch (-(fi17)) {
                    case 32768:
                        v5 = (v5 * (v5 - 42))
                    case -1:
                        if false {
                            v19 := GenS1{F0: genFn0()}
                            if (v15 >= v15) {
                                v14 += (v16 ^ v16)
                            }
                            for fi20 := 0; fi20 < 10; fi20++ {
                                if (v15 <= -128) {
                                    v3 = (v3 / ((v3) | 1))
                                    v21 := GenS2{F0: v5, F1: uint16(v3)}
                                    _ = v21
                                } else {
                                    var v22 string = "a"
                                    switch v22 {
                                    case "\n\t":
                                        fmt.Println(v3, fi20, fi20)
                                    case "":
                                        v23 := GenS1{F0: (v3 | (^(v3)))}
                                        _ = v23
                                    case "a":
                                        var v24 int32 = 2147483647
                                        _ = v24
                                    }
                                    var v25 uint8 = genFn0()
                                    _ = v22
                                    _ = v25
                                }
                                v14 += v14
                                _ = fi20
                            }
                            _ = v19
                        }
                    }
                }
            case "\n\t":
                if (true || (false || (true && false))) {
                    v14 = v14
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    var v26 int8 = (v15 - (v15 << ((65535) & 63)))
                    _ = v26
                } else {
                    switch int(v4) {
                    case 42:
                        v16 += 42
                    case 32767:
                        v14 -= v14
                    case -9223372036854775808:
                        var v27 uint64 = 538
                        _ = v27
                    default:
                        var v28 uint16 = (v5 & (v5 >> ((uint(v5)) & 63)))
                        _ = v28
                    }
                }
            default:
                fmt.Println(v15, v16)
            }
            var v29 bool = false
            for fi30 := 0; fi30 < 8; fi30++ {
                _ = len(v1)
                _ = len(v1)
                switch 128 {
                case 845:
                    v15 -= v15
                case 32768:
                    if v29 {
                        v5 -= uint16(v4)
                        if v29 {
                            v16 = v16
                            var v31 uint16 = (v5 - 565)
                            v2 = append(v2, "\n\t")
                            _ = v31
                        }
                    } else {
                        v32 := GenS3{F0: 993, F1: (^(v5))}
                        v2 = append(v2, "\n\t")
                        switch "" {
                        case "a":
                            v15 = (v15 & v15)
                        case "\n\t":
                            v3 = genFn0()
                        default:
                            fmt.Println(v4, fi30, v29)
                        }
                        _ = v32
                    }
                }
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + ("foo bar" + ("\n\t" + (("hello" + ("hello" + ((("" + "a") + "a") + ""))) + ""))))
                }
                _ = fi30
            }
            _ = fi17
            _ = v29
        }
    case "foo bar":
        switch ("\n\t" + ((((("" + "hello") + "") + "") + ("a" + ((("\n\t" + ("a" + "foo bar")) + (((("" + "") + "") + "hello") + "foo bar")) + "\n\t"))) + (("hello" + "\n\t") + "foo bar"))) {
        case "":
            v16 += 42
        default:
            var v33 string = "\n\t"
            _ = v33
        }
    case "hello":
        var v34 int16 = v16
        _ = v34
    default:
        fmt.Println(v14)
    }
    v35 := GenS1{F0: (v3 ^ (v3 % ((v3) | 1)))}
    v4 -= (v4 - v4)
    for fi36 := 0; fi36 < 20; fi36++ {
        fmt.Println(v14, v14)
        switch -76 {
        case 897:
            v14 = (v16 % ((-1) | 1))
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            fmt.Println(v14)
            var v38 int16 = 1000
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    if true {
        var v39 int8 = v15
        if true {
            v14 = v16
        } else {
            switch "\n\t" {
            case "foo bar":
                var v40 int64 = 711
                _ = v40
            case "":
                v39 += (^(v39))
            case "\n\t":
                var v41 float64 = 1.5
                _ = v41
            default:
                if ((v3 + v3) <= (v3 & 253)) {
                    if (29 >= 76) {
                        v2 = append(v2, "a")
                    }
                } else {
                    v14 = 32767
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            if true {
                var v42 uint32 = 1
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = v42
            } else {
                fmt.Println(v39, v15)
            }
        }
        _ = v39
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        fmt.Println(v14)
        v44 := GenS1{F0: (^(v3))}
        v3 = (v3 ^ v3)
        _ = fi43
        _ = v44
    }
    if ((^(v3)) <= (v3 ^ genFn0())) {
        var v45 int64 = 9223372036854775807
        for fi46 := 0; fi46 < 8; fi46++ {
            _ = len(v2)
            v15 += (v15 & v15)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi46
        }
        _ = v45
    } else {
        switch "" {
        case "":
            if len(v1) > 0 {
                v1[0] = ("a" + ("a" + "a"))
            }
        case "a":
            var v47 int64 = -1
            _ = v47
        case "foo bar":
            var v48 string = ("" + "hello")
            _ = v48
        }
        v49 := GenS2{F0: (^(v5)), F1: v5}
        for fi50 := 0; fi50 < 1; fi50++ {
            if false {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                fmt.Println(v16, v14, fi50)
                if true {
                    if (!(true && (true && ("a" >= ("a" + "a"))))) {
                        for fi51 := 0; fi51 < 3; fi51++ {
                            v16 += (v16 - (v14 | 0))
                            _ = fi51
                        }
                        v52 := GenS1{F0: (v3 >> ((v4) & 63))}
                        var v53 float64 = ((1.5 - 0.1) - (1e10 * -2.25))
                        _ = v52
                        _ = v53
                    }
                    v5 += (v5 >> (((v4 - (^(v4)))) & 63))
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v4, v3, v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi50
        }
        _ = v49
    }
    var v54 uint16 = (v5 & (v5 | v5))
    v5 += v54
    fmt.Println(v5, v15, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v54
}
