package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint = 0
    var v3 uint = v2
    v4 := make(map[uint32]uint32)
    var v5 uint32 = v1
    v5 -= uint32(v2)
    if true {
        var v6 uint = (^(v2))
        _ = v6
    }
    v2 -= 255
    fmt.Println(v5)
    fmt.Println(v3, v1, v5)
    v2 = (v2 / ((v3) | 1))
    var v7 int32 = 144
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := v4[(v1 >> ((v2) & 63))]
        for fi10 := 0; fi10 < 18; fi10++ {
            switch -2147483648 {
            case 0:
                var v11 float64 = (float64((1.5 + (3.14159 - 1.5))) + 1.5)
                _ = v11
            case 492:
                v12 := GenS1{F0: -61, F1: 255}
                _ = v12
            case 585:
                fmt.Println(fi8, v9)
                var v13 string = ""
                _ = v13
            default:
                fmt.Println(v1, v1, v1)
            }
            switch "" {
            case "a":
                fmt.Println(v5, v3)
            case "\n\t":
                fmt.Println(fi10)
            case "hello":
                v3 = (v2 >> (((v3 ^ 494)) & 63))
            default:
                fmt.Println(fi8)
            }
            var v14 uint = (v3 % ((65535) | 1))
            _ = fi10
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    v2 = v3
    var v15 string = ""
    var v16 string = "foo bar"
    if true {
        if false {
            if (65535 >= 65535) {
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch (fi17 * fi17) {
                    case 127:
                        if ((v5 ^ 1) <= (v1 + v5)) {
                            v16 = "hello"
                        } else {
                            v1 -= (v1 ^ v1)
                        }
                    case 1:
                        switch (^(fi17)) {
                        case 128:
                            var v18 string = "a"
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 18; fi19++ {
                                if (!false) {
                                    if (false || (32767 >= 32767)) {
                                        var v20 int8 = 0
                                        if (1 >= 128) {
                                            delete(v4, 2147483648)
                                            fmt.Println(v15)
                                        } else {
                                            if (((v5 << ((2) & 63)) <= (v5 + (^(v5)))) && false) {
                                                v21 := GenS1{F0: -106, F1: 1}
                                                switch v15 {
                                                case "":
                                                    switch -1 {
                                                    case 42:
                                                        _ = len(v4)
                                                    case 2147483647:
                                                        var v22 int = fi17
                                                        _ = v22
                                                    }
                                                case "hello":
                                                    v16 = ("hello" + v16)
                                                case "\n\t":
                                                    if false {
                                                        if ((float32(float64(-2.25)) * 1.5) == ((0.1 - float32(float64(0.0))) - float32(3.14159))) {
                                                            if true {
                                                                v23 := GenS1{F0: 1, F1: 255}
                                                                v24 := []int{}
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                            var v25 int32 = v7
                                                            _ = v25
                                                        }
                                                        switch (v15 + ((v15 + v16) + (v16 + v16))) {
                                                        case "hello":
                                                            v20 = (v20 / ((-128) | 1))
                                                        }
                                                        var v26 int64 = 773
                                                        _ = v26
                                                    } else {
                                                        v5 -= (v1 << (((v2 >> ((v3) & 63))) & 63))
                                                        v20 = 42
                                                        v27 := GenS1{F0: (v20 - v20), F1: 0}
                                                        if false {
                                                            var v28 string = ((v16 + v16) + "a")
                                                            _ = v28
                                                        } else {
                                                            var v29 uint64 = 684
                                                            var v30 float64 = 3.14159
                                                            v16 = (v15 + "")
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                        fmt.Println(v7, v5, fi19)
                                                        _ = v27
                                                    }
                                                }
                                                _ = v21
                                            } else {
                                                if (!(true || (true || ((v7 == (-(v7))) || false)))) {
                                                    v15 = v15
                                                    var v31 uint = (v2 / (((^(v2))) | 1))
                                                    _ = v31
                                                }
                                                fmt.Println(v3, fi19)
                                                if true {
                                                    var v32 float32 = (-2.25 / (0.1 + 1.5))
                                                    var v33 int16 = -74
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    v3 = 4294967295
                                                    v34 := GenS2{F0: true, F1: (v16 + "foo bar")}
                                                    if ((40 > 32767) && (!true)) {
                                                        var v35 uint16 = 2
                                                        fmt.Println(v20, v7, fi19)
                                                        _ = v35
                                                    } else {
                                                        var v36 string = v16
                                                        fmt.Println(v36)
                                                        _ = v36
                                                    }
                                                    _ = v34
                                                }
                                                var v37 int32 = (v7 - 911)
                                                v20 -= 100
                                                _ = v37
                                            }
                                        }
                                        _ = v20
                                    }
                                    v38 := GenS2{F0: (!true), F1: v16}
                                    v15 = "a"
                                    _ = v38
                                }
                                fmt.Println(v2)
                                v39 := GenS1{F0: -18, F1: 32767}
                                v16 = ("" + "hello")
                                _ = fi19
                                _ = v39
                            }
                        }
                    default:
                        for fi40 := 0; fi40 < 11; fi40++ {
                            _ = len(v4)
                            v5 -= (^(v5))
                            if (!(!false)) {
                                if true {
                                    var v41 int8 = -1
                                    if (32768 > -9223372036854775808) {
                                        if true {
                                            var v42 uint16 = 884
                                            if (!((^(v1)) >= v5)) {
                                                if (181 < 0) {
                                                    if true {
                                                        var v43 uint = v3
                                                        var v44 int64 = -1
                                                        v45 := GenS1{F0: 100, F1: 321}
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                } else {
                                                    v1 -= (^(v1))
                                                    if true {
                                                        fmt.Println(v1, v16, v15)
                                                        v46 := []int{}
                                                        v47 := GenS3{F0: -1}
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    if (true && false) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if (!((v15 + "hello") != "foo bar")) {
                                                                var v48 float64 = 0.1
                                                                v49 := GenS4{F0: (v3 & (v3 << (((^(v2))) & 63)))}
                                                                v15 = (v16 + "foo bar")
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        case "foo bar":
                                                            var v50 float64 = 1e10
                                                            fmt.Println(v3, fi17, v42)
                                                            _ = v50
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                        var v51 uint = v3
                                                        v15 = v15
                                                        fmt.Println(v3, v16)
                                                        _ = v51
                                                    } else {
                                                        v16 = v16
                                                    }
                                                }
                                                fmt.Println(v16, v15, v15)
                                                _ = len(v4)
                                            }
                                            v5 = v5
                                            v52 := GenS4{F0: 687}
                                            var v53 int64 = 636
                                            _ = v42
                                            _ = v52
                                            _ = v53
                                        }
                                        var v54 float64 = (1e10 - 1.5)
                                        var v55 string = "hello"
                                        _ = v54
                                        _ = v55
                                    }
                                    v1 += v1
                                    _ = v41
                                } else {
                                    var v56 float32 = 1e10
                                    _ = v56
                                }
                                var v57 int8 = 56
                                if true {
                                    fmt.Println(fi40, v57, v15)
                                } else {
                                    v1 -= 2
                                    fmt.Println(v2, v3)
                                    if (32767 == 655) {
                                        var v58 int64 = 0
                                        var v59 float32 = 3.14159
                                        v15 = v15
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v60 := GenS2{F0: true, F1: "\n\t"}
                                        _ = v60
                                    }
                                }
                                _ = v57
                            }
                            _ = fi40
                        }
                    }
                    v61 := GenS5{F0: 1}
                    _ = fi17
                    _ = v61
                }
                v62 := GenS6{F0: v3}
                _ = v62
            } else {
                if false {
                    for fi63 := 0; fi63 < 19; fi63++ {
                        var v64 bool = false
                        _ = fi63
                        _ = v64
                    }
                    var v65 uint = (v2 | v2)
                    _ = v65
                } else {
                    v66 := []int{2147483647, 1, -28}
                    var v67 string = "\n\t"
                    v16 = ""
                    _ = v66
                    _ = v67
                }
            }
        }
    }
    fmt.Println(v16)
    var v68 int8 = 38
    v69 := GenS2{F0: ((v2 << ((v3) & 63)) == v3), F1: v15}
    v70 := GenS1{F0: 0, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v7
    _ = v70
}
