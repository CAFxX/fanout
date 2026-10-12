package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int; F1 bool; F2 uint32 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string; F1 int32 }

func main() {
    var v1 uint64 = 2
    var v2 uint = 42
    v3 := make(map[uint32]bool)
    var v4 uint = 65535
    var v5 uint8 = 196
    var v6 int8 = 15
    v2 = (v4 | (v2 / ((v2) | 1)))
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: false}
        _ = fi7
        _ = v8
    }
    v5 = uint8(v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v1)
        for fi10 := 0; fi10 < 13; fi10++ {
            switch "" {
            case "a":
                v11 := GenS2{F0: (v5 % (((^(v5))) | 1)), F1: 4294967295}
                _ = v11
            }
            _ = fi10
        }
        v12 := []int{fi9, fi9}
        _ = fi9
        _ = v12
    }
    if true {
        fmt.Println(v4)
        v6 += (v6 & (v6 << (((v4 & v4)) & 63)))
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: (fi13 > (fi13 << (((v4 | 255)) & 63)))}
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v4, fi15, v4)
                switch (fi13 ^ 32767) {
                case -177:
                    if (!false) {
                        v6 = (v6 & v6)
                        fmt.Println(v1)
                    }
                default:
                    fmt.Println(fi13, v2)
                }
                if true {
                    if false {
                        v16 := GenS2{F0: v5, F1: v1}
                        var v17 uint16 = 255
                        v18 := GenS3{F0: -426}
                        v19 := GenS3{F0: 235}
                        if (!(false || (v6 < 1))) {
                            var v20 int = (fi15 ^ (-(fi13)))
                            var v21 int16 = 0
                            v22 := []int{v20}
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            var v23 uint = v4
                            var v24 float32 = 3.14159
                            v25 := GenS3{F0: 1000}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v5 = v5
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    delete(v3, 608)
                } else {
                    var v26 string = ""
                    _ = v26
                }
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            switch "foo bar" {
            case "a":
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 int64 = 9223372036854775807
                    v30 := GenS3{F0: -400}
                    var v31 uint = v4
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            case "foo bar":
                v1 += (v1 << ((536) & 63))
            default:
                delete(v3, uint32(v1))
            }
            var v32 float32 = 0.1
            switch ("" + ("foo bar" + "hello")) {
            case "\n\t":
                var v33 int = (fi27 + fi27)
                _ = v33
            default:
                v34 := GenS1{F0: true}
                _ = v34
            }
            _ = fi27
            _ = v32
        }
        fmt.Println(v5, v6)
        if (1.5 <= 1e10) {
            var v35 uint32 = 65535
            _ = v35
        } else {
            for fi36 := 0; fi36 < 6; fi36++ {
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 bool = ((!(!(uint(v1) >= v4))) && true)
                    v4 = (v2 / ((1) | 1))
                    switch (fi37 * 0) {
                    case 127:
                        if true {
                            v4 += 255
                            v39 := []int{(fi36 % ((-2147483648) | 1)), (fi36 % (((fi37 >> ((0) & 63))) | 1))}
                            var v40 float64 = 0.1
                            _ = v39
                            _ = v40
                        }
                    case 32767:
                        v41 := GenS2{F0: v5, F1: uint64(v2)}
                        v42 := GenS4{F0: 127, F1: (v1 / ((v1) | 1)), F2: 1}
                        _ = v41
                        _ = v42
                    }
                    _ = fi37
                    _ = v38
                }
                _ = fi36
            }
            delete(v3, uint32(v1))
            v43 := []int{int(v6), 42, -450}
            _ = v43
        }
    }
    var v44 int16 = 1000
    switch (("hello" + ("hello" + "foo bar")) + "a") {
    case "\n\t":
        for fi45 := 0; fi45 < 9; fi45++ {
            if false {
                _ = len(v3)
                v4 = 901
            }
            _ = fi45
        }
    case "":
        v46 := []int{32768, -9223372036854775808}
        _ = v46
    default:
        v47 := v3[1]
        _ = v47
    }
    v4 -= (v4 + (v2 >> ((v2) & 63)))
    for fi48 := 0; fi48 < 15; fi48++ {
        v49 := GenS5{F0: 2147483647, F1: -2147483648}
        v50 := GenS1{F0: true}
        v3[65535] = ((v44 << ((272) & 63)) != int16(v2))
        _ = fi48
        _ = v49
        _ = v50
    }
    if false {
        var v51 int = 4294967295
        v6 += 42
        var v52 float64 = 1.5
        _ = v51
        _ = v52
    }
    v4 += (^(v4))
    var v53 string = (("\n\t" + ("a" + ((("" + (("\n\t" + "hello") + "hello")) + (("a" + ("\n\t" + ("\n\t" + "foo bar"))) + "a")) + "a"))) + "a")
    var v54 uint16 = 254
    v6 = -28
    for fi55 := 0; fi55 < 8; fi55++ {
        if ((!true) || ((fi55 - fi55) == (fi55 >> (((v4 ^ (^(v2)))) & 63)))) {
            v56 := v3[456]
            switch int(v44) {
            case 9223372036854775807:
                switch -429 {
                case -1:
                    fmt.Println(v56, v56, v5)
                default:
                    fmt.Println(fi55)
                }
            case 58:
                if ((float64(-2.25) / (1e10 * 0.1)) > 0.0) {
                    v57 := []int{fi55, (fi55 & (-(fi55)))}
                    if len(v57) > 0 {
                        v57[0] = fi55
                    }
                    if false {
                        fmt.Println(v6, v1, v53)
                        v56 = (true || (!false))
                    }
                    _ = v57
                } else {
                    switch v53 {
                    case "hello":
                        v5 = v5
                    case "foo bar":
                        v54 = 65535
                    case "a":
                        delete(v3, 753)
                    default:
                        for fi58 := 0; fi58 < 15; fi58++ {
                            v6 += (v6 % ((1) | 1))
                            v59 := GenS6{F0: (fi58 & 127), F1: false, F2: 2147483647}
                            _ = fi58
                            _ = v59
                        }
                    }
                    v60 := GenS1{F0: true}
                    if v56 {
                        v61 := GenS4{F0: 7, F1: v1, F2: 42}
                        _ = v61
                    } else {
                        for fi62 := 0; fi62 < 7; fi62++ {
                            v63 := v3[2147483648]
                            _ = fi62
                            _ = v63
                        }
                        v64 := GenS7{F0: (v44 / ((-677) | 1)), F1: int64(v54)}
                        _ = v64
                    }
                    _ = v60
                }
            case -1:
                v56 = ((128 == fi55) || ((uint(v1) != 0) || (!(((((0.1 * (0.0 + float32((((float64(1e10) - 0.0) - float64(0.1)) * 1e10)))) * 0.1) >= 1e10) && ((-2.25 > 0.0) || false)) || v56))))
            }
            v65 := GenS8{F0: 0}
            _ = v56
            _ = v65
        } else {
            v66 := GenS9{F0: "foo bar", F1: -2147483648}
            var v67 uint16 = 855
            v53 = v53
            _ = v66
            _ = v67
        }
        _ = fi55
    }
    v2 -= v4
    var v68 uint16 = (^(v54))
    v2 = (v2 / (((v2 << ((18446744073709551615) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v68
}
