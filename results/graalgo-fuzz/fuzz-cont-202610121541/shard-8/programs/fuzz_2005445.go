package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int; F1 uint16; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 string; F1 int16; F2 int64 }
func genFn0() uint16 {
    if (518 > 128) {
        v7 := GenS1{F0: 0, F1: 255, F2: 42}
        var v8 int = 42
        fmt.Println(v8)
        _ = v7
        _ = v8
    } else {
        var v9 uint = 4294967295
        _ = v9
    }
    v10 := GenS2{F0: 0}
    _ = v10
    return 465
}

func main() {
    var v1 uint64 = 889
    v2 := []uint32{0, 67, 2147483647, 90}
    var v3 int16 = -210
    v4 := make(map[uint32]int)
    var v5 uint = 255
    v6 := []string{(((("" + "\n\t") + ("hello" + "a")) + "foo bar") + ("" + ("hello" + ""))), "foo bar", "a", ("a" + "foo bar")}
    v11 := GenS3{F0: 0}
    v3 += v3
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS3{F0: 1}
            var v15 bool = (645 > (^(fi12)))
            v5 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        delete(v4, 65535)
        switch ("a" + "hello") {
        case "\n\t":
            switch "\n\t" {
            case "a":
                v16 := GenS2{F0: v5}
                _ = v16
            default:
                var v17 bool = (!(18446744073709551615 > v5))
                _ = v17
            }
        default:
            v3 -= v3
        }
        _ = fi12
    }
    fmt.Println(v5, v1, v3)
    var v18 uint8 = 128
    delete(v4, uint32(v1))
    switch 185 {
    case 32767:
        var v19 int8 = 13
        _ = v19
    case 170:
        for fi20 := 0; fi20 < 4; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                switch 9223372036854775807 {
                case 2147483647:
                    var v22 int64 = 171
                    _ = v22
                case -2147483648:
                    if (("foo bar" + "a") >= ((("" + ("" + "foo bar")) + (("hello" + ("a" + "")) + "")) + "a")) {
                        if (-757 == 1) {
                            v23 := GenS4{F0: v18, F1: (v3 * 32767)}
                            _ = v23
                        }
                    } else {
                        switch (fi20 - 42) {
                        case 128:
                            v24 := GenS4{F0: (v18 % ((2) | 1)), F1: (^(v3))}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v24
                        case -9223372036854775808:
                            delete(v4, 4294967295)
                        }
                        var v25 int8 = 1
                        _ = v25
                    }
                default:
                    v18 -= 119
                }
                v5 += v5
                _ = fi21
            }
            fmt.Println(v3, v18, v1)
            _ = fi20
        }
    }
    _ = len(v6)
    var v26 int8 = 1
    var v27 string = "a"
    v27 = v27
    if false {
        delete(v4, 819)
        v28 := GenS5{F0: -9223372036854775808, F1: uint16(v5), F2: v3}
        if (true || false) {
            switch ("\n\t" + v27) {
            case "foo bar":
                v3 -= v3
            case "hello":
                v18 = v18
            }
        } else {
            if (!false) {
                v29 := GenS6{F0: (0.1 < 1.5)}
                v5 = (v5 << ((v5) & 63))
                if false {
                    v3 = v3
                    for fi30 := 0; fi30 < 7; fi30++ {
                        if false {
                            v26 -= -68
                            var v31 string = (v27 + "hello")
                            for fi32 := 0; fi32 < 14; fi32++ {
                                delete(v4, 65535)
                                _ = fi32
                            }
                            _ = v31
                        }
                        fmt.Println(v26, v18)
                        fmt.Println(v1, v27)
                        _ = fi30
                    }
                }
                _ = v29
            }
            for fi33 := 0; fi33 < 6; fi33++ {
                v3 -= (v3 ^ 0)
                for fi34 := 0; fi34 < 10; fi34++ {
                    v35 := GenS1{F0: v3, F1: genFn0(), F2: 0}
                    _ = len(v4)
                    var v36 int16 = 778
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            if true {
                var v37 int64 = int64(v3)
                _ = v37
            } else {
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 int16 = (v3 + v3)
                    _ = fi38
                    _ = v39
                }
                if true {
                    var v40 uint8 = (v18 % (((v18 >> (((v5 >> (((^(v5))) & 63))) & 63))) | 1))
                    _ = v40
                }
                for fi41 := 0; fi41 < 6; fi41++ {
                    v27 = (v27 + ("\n\t" + ("" + v27)))
                    v42 := GenS7{F0: uint16(v1), F1: (v27 + ("a" + (v27 + (v27 + "hello"))))}
                    _ = fi41
                    _ = v42
                }
            }
        }
        _ = v28
    } else {
        v43 := GenS8{F0: v27, F1: (v3 - v3), F2: -380}
        _ = v43
    }
    var v44 int32 = 0
    if true {
        if (float64(float32(float64(-2.25))) == 3.14159) {
            v45 := GenS3{F0: 65535}
            _ = v45
        }
        var v46 uint32 = 55
        var v47 uint = v5
        _ = v46
        _ = v47
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v48 uint16 = 1
    v49 := GenS2{F0: v5}
    for fi50 := 0; fi50 < 4; fi50++ {
        v48 = 2
        fmt.Println(v3, v27)
        v48 -= (v48 - genFn0())
        _ = fi50
    }
    for fi51 := 0; fi51 < 7; fi51++ {
        switch "\n\t" {
        case "":
            var v52 int64 = -1
            _ = v52
        case "hello":
            v53 := GenS6{F0: (int(v44) <= fi51)}
            _ = v53
        case "\n\t":
            if (!true) {
                var v54 int64 = 9223372036854775807
                _ = v54
            }
        }
        if len(v2) > 0 {
            v2[0] = 1
        }
        var v55 string = v27
        _ = fi51
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
