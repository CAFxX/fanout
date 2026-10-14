package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int16; F1 string; F2 uint8 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 0
    v3 := make(map[uint32]string)
    if true {
        fmt.Println(v1, v1)
        v1 += v1
        if false {
            var v4 float32 = (0.1 * float32(0.0))
            if true {
                v5 := []int{}
                _ = v5
            } else {
                if true {
                    var v6 uint8 = 2
                    if (42 <= 18446744073709551615) {
                        v7 := GenS1{F0: uint16(v2)}
                        fmt.Println(v4, v1, v4)
                        var v8 int8 = 1
                        _ = v7
                        _ = v8
                    } else {
                        if ((^(v2)) == v2) {
                            var v9 uint32 = (^(v2))
                            _ = v9
                        } else {
                            v10 := GenS2{F0: 1}
                            var v11 uint = 0
                            fmt.Println(v1, v11, v2)
                            _ = v10
                            _ = v11
                        }
                        v12 := GenS1{F0: 65535}
                        if ((true || true) && (true && ((false && true) || (!false)))) {
                            var v13 int8 = -1
                            _ = v13
                        }
                        _ = v12
                    }
                    _ = v6
                } else {
                    switch 127 {
                    case 104:
                        var v14 int8 = 0
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int32 = 0
                    fmt.Println(fi15, v2, v4)
                    _ = fi15
                    _ = v16
                }
            }
            _ = v4
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int32 = -2147483648
                switch "" {
                case "hello":
                    fmt.Println(v2, v18)
                case "":
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        v19 := GenS1{F0: 255}
                        _ = v19
                    default:
                        var v20 bool = (1e10 != 1e10)
                        _ = v20
                    }
                }
                _ = fi17
                _ = v18
            }
        }
    } else {
        var v21 uint16 = 0
        fmt.Println(v21, v1)
        v2 = v2
        _ = v21
    }
    if false {
        var v22 uint64 = 440
        fmt.Println(v22, v2)
        var v23 bool = (86 > -128)
        _ = v22
        _ = v23
    }
    if true {
        v24 := GenS3{F0: "hello"}
        switch "a" {
        case "hello":
            for fi25 := 0; fi25 < 19; fi25++ {
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := GenS2{F0: 127}
                    if (2 != (v1 * uint8(v2))) {
                        v28 := GenS2{F0: -128}
                        var v29 int8 = 1
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        default:
            if (2147483648 != v2) {
                var v30 uint16 = 954
                _ = v30
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v1 = (v1 & 127)
                    _ = fi31
                }
            }
        }
        _ = v24
    } else {
        for fi32 := 0; fi32 < 3; fi32++ {
            if (!false) {
                for fi33 := 0; fi33 < 15; fi33++ {
                    switch ("hello" + "foo bar") {
                    case "":
                        v34 := GenS1{F0: uint16(v2)}
                        _ = v34
                    case "hello":
                        if false {
                            v35 := GenS4{F0: 0}
                            v36 := []int{(fi32 >> ((1) & 63)), (fi32 / (((fi32 * fi33)) | 1))}
                            _ = v35
                            _ = v36
                        } else {
                            v1 -= 128
                            _ = len(v3)
                            v2 -= (v2 << ((4294967295) & 63))
                            var v37 int16 = 1000
                            _ = v37
                        }
                    default:
                        v1 = v1
                    }
                    _ = fi33
                }
            } else {
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 int = fi32
                    v40 := GenS5{F0: (v1 + v1)}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                switch (fi32 >> ((4294967295) & 63)) {
                case 127:
                    v2 = v2
                case -532:
                    v41 := GenS2{F0: -64}
                    _ = v41
                case -9223372036854775808:
                    switch "\n\t" {
                    case "":
                        v42 := []int{fi32, fi32, fi32, fi32}
                        _ = v42
                    }
                }
            }
            _ = fi32
        }
    }
    var v43 int8 = -128
    v44 := GenS5{F0: (^(v1))}
    v45 := GenS1{F0: 42}
    v46 := v3[v2]
    v46 = v46
    v1 = 1
    v47 := GenS4{F0: 255}
    for fi48 := 0; fi48 < 17; fi48++ {
        v49 := v3[v2]
        v50 := GenS3{F0: ("" + ("\n\t" + v46))}
        for fi51 := 0; fi51 < 5; fi51++ {
            if (((true && ((2 > v1) && true)) && false) && false) {
                var v52 int8 = (-(v43))
                v53 := GenS6{F0: 179, F1: v43}
                var v54 uint32 = (^(v2))
                _ = v52
                _ = v53
                _ = v54
            }
            switch (^(fi48)) {
            case 618:
                var v55 int = (fi51 * fi51)
                _ = v55
            case -2147483648:
                if (((1000 != 32767) || false) || (!((fi48 - fi51) == (-(fi48))))) {
                    if (0.0 < (0.1 / 3.14159)) {
                        v49 = v49
                        var v56 uint = 42
                        fmt.Println(v43, fi51, v43)
                        _ = v56
                    }
                    if (uint16(v1) != 255) {
                        var v57 int8 = 1
                        _ = v57
                    } else {
                        v46 = ("hello" + v49)
                        fmt.Println(v49)
                        v58 := GenS5{F0: v1}
                        _ = v58
                    }
                }
            }
            if (v2 < 2) {
                v3[1] = ""
                v46 = ("" + "\n\t")
                v46 = v49
            }
            _ = fi51
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    v59 := GenS7{F0: -32768, F1: ((((v46 + v46) + "a") + v46) + ("hello" + (v46 + v46))), F2: (v1 << ((2) & 63))}
    if ((false && ((v2 * v2) < (v2 % ((v2) | 1)))) && (!true)) {
        fmt.Println(v46)
        if true {
            for fi60 := 0; fi60 < 8; fi60++ {
                if false {
                    v2 = (v2 & (v2 & uint32(v1)))
                    v61 := []int{-1, 736, fi60}
                    _ = v61
                } else {
                    var v62 int = fi60
                    _ = v62
                }
                fmt.Println(v46, v43)
                _ = fi60
            }
            if false {
                for fi63 := 0; fi63 < 11; fi63++ {
                    var v64 int8 = (^(v43))
                    v43 = -128
                    for fi65 := 0; fi65 < 3; fi65++ {
                        v46 = "hello"
                        _ = fi65
                    }
                    _ = fi63
                    _ = v64
                }
                var v66 string = v46
                _ = v66
            }
        }
        var v67 int64 = 1
        _ = v67
    } else {
        v68 := []int{1, 32768}
        var v69 string = ("hello" + v46)
        _ = v68
        _ = v69
    }
    fmt.Println(v43, v1)
    if false {
        v46 = ((((v46 + (v46 + v46)) + v46) + ("a" + v46)) + v46)
        for fi70 := 0; fi70 < 15; fi70++ {
            var v71 uint64 = 202
            _ = fi70
            _ = v71
        }
        v72 := GenS8{F0: (v2 >> ((624) & 63))}
        _ = v72
    }
    v3[v2] = (("hello" + v46) + v46)
    v73 := GenS6{F0: 761, F1: (v43 + v43)}
    v74 := GenS2{F0: (v43 - 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v59
    _ = v73
    _ = v74
}
