package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0() uint16 {
    if true {
        v6 := []int{-13}
        v7 := GenS1{F0: 1, F1: 758, F2: 708}
        _ = v6
        _ = v7
    } else {
        v8 := GenS2{F0: -224, F1: 446}
        var v9 int32 = -1
        var v10 int64 = 42
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS3{F0: 255, F1: -197}
    _ = v11
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := []int{}
    var v3 int32 = 1
    v4 := []string{}
    var v5 uint8 = v1
    fmt.Println(v1, v3, v3)
    _ = len(v4)
    for fi12 := 0; fi12 < 18; fi12++ {
        switch "\n\t" {
        case "foo bar":
            v3 += v3
        case "\n\t":
            fmt.Println(v5, v5)
        }
        switch fi12 {
        case -976:
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 string = ((("hello" + "a") + ("" + "hello")) + ("hello" + "\n\t"))
                _ = fi13
                _ = v14
            }
        default:
            v15 := GenS1{F0: 32767, F1: fi12, F2: 0}
            _ = v15
        }
        fmt.Println(v3, v1)
        _ = fi12
    }
    v16 := GenS4{F0: 104, F1: 0}
    for fi17 := 0; fi17 < 7; fi17++ {
        switch fi17 {
        case 10:
            var v18 float64 = (1.5 - -2.25)
            _ = v18
        case 1:
            v19 := GenS2{F0: -759, F1: 2}
            _ = v19
        case -2147483648:
            fmt.Println(fi17)
        default:
            for fi20 := 0; fi20 < 6; fi20++ {
                if false {
                    v21 := GenS1{F0: 42, F1: 32768, F2: 333}
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi17, fi20, fi17)
                    case "\n\t":
                        fmt.Println(v5, fi17, v1)
                    case "":
                        if true {
                            v1 += v1
                            fmt.Println(fi20)
                            v3 = v3
                            var v22 bool = true
                            _ = v22
                        }
                    default:
                        var v23 int16 = -1
                        _ = v23
                    }
                    _ = v21
                } else {
                    v1 -= (^(v1))
                }
                v1 -= 1
                v3 -= (v3 & v3)
                _ = fi20
            }
        }
        _ = fi17
    }
    v24 := GenS4{F0: 1, F1: 127}
    for fi25 := 0; fi25 < 19; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            v1 = 1
            _ = fi26
        }
        if (-1 > -764) {
            v1 -= 128
            var v27 bool = true
            _ = v27
        } else {
            v28 := GenS5{F0: 255}
            v29 := GenS5{F0: 4294967295}
            switch "" {
            case "a":
                _ = len(v2)
            case "\n\t":
                v30 := GenS6{F0: -128}
                _ = v30
            default:
                switch fi25 {
                case 200:
                    var v31 int64 = 1
                    _ = v31
                case 128:
                    switch ("" + "") {
                    case "foo bar":
                        var v32 float64 = 0.1
                        _ = v32
                    case "\n\t":
                        for fi33 := 0; fi33 < 7; fi33++ {
                            if (false && true) {
                                var v34 string = "hello"
                                _ = len(v2)
                                _ = v34
                            }
                            var v35 uint8 = (v1 >> ((65535) & 63))
                            switch (fi25 + fi25) {
                            case -9223372036854775808:
                                fmt.Println(fi25)
                                v36 := GenS7{F0: (v3 / (((v3 % (((^(v3))) | 1))) | 1)), F1: genFn0()}
                                fmt.Println(v35)
                                v5 = 1
                                _ = v36
                            case 985:
                                v35 = v5
                            case 1:
                                var v37 int32 = v3
                                _ = v37
                            default:
                                v3 -= v3
                            }
                            _ = fi33
                            _ = v35
                        }
                    case "a":
                        if true {
                            switch ("a" + "hello") {
                            case "hello":
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    var v39 uint8 = v1
                                    switch (("" + "a") + "") {
                                    case "a":
                                        v40 := GenS2{F0: -1, F1: 255}
                                        _ = v40
                                    case "hello":
                                        if (!(false && ((v3 % (((v3 >> ((571) & 63))) | 1)) <= (v3 - -2147483648)))) {
                                            if true {
                                                _ = len(v2)
                                                v1 = (v5 % ((8) | 1))
                                                v1 += v5
                                            } else {
                                                v5 = (v5 << ((407) & 63))
                                                var v41 int32 = 0
                                                _ = len(v4)
                                                _ = v41
                                            }
                                            v42 := GenS4{F0: 32767, F1: -10}
                                            _ = v42
                                        } else {
                                            var v43 int64 = 1
                                            v44 := GenS2{F0: -740, F1: 2}
                                            v4 = append(v4, "a")
                                            _ = v43
                                            _ = v44
                                        }
                                    case "foo bar":
                                        var v45 int64 = 559
                                        _ = v45
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            if (!(("a" < ("foo bar" + "hello")) && (0 != 42))) {
                                _ = len(v4)
                                if (1 < -32768) {
                                    v46 := GenS3{F0: 18446744073709551615, F1: v3}
                                    _ = v46
                                } else {
                                    var v47 float64 = 0.0
                                    _ = v47
                                }
                            }
                            v1 -= v5
                        }
                    }
                default:
                    v48 := GenS8{F0: 42, F1: genFn0(), F2: (v3 % (((^(v3))) | 1))}
                    _ = v48
                }
            }
            _ = v28
            _ = v29
        }
        v5 -= (v5 << ((65535) & 63))
        _ = fi25
    }
    if true {
        v1 += (v5 & (v5 & (v1 & (v5 - v5))))
        var v49 int16 = -1
        _ = v49
    }
    fmt.Println(v3, v3, v1)
    var v50 string = ("a" + "")
    for fi51 := 0; fi51 < 3; fi51++ {
        v3 = (v3 << ((2) & 63))
        var v52 int = fi51
        _ = fi51
        _ = v52
    }
    v3 -= (v3 * -1)
    var v53 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
}
