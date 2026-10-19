package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int; F2 bool }
type GenS6 struct { F0 int32; F1 uint32; F2 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int64; F1 int; F2 string }
type GenS9 struct { F0 uint32; F1 int64; F2 int }
type GenS10 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS11 struct { F0 int; F1 uint16; F2 uint16 }
type GenS12 struct { F0 int64; F1 bool; F2 int32 }
type GenS13 struct { F0 uint; F1 uint; F2 int16 }
type GenS14 struct { F0 uint32; F1 int; F2 uint32 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 bool; F2 int32 }
type GenS17 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8) int8 {
    v4 := []int{-1, 32767}
    for fi5 := 0; fi5 < 16; fi5++ {
        if len(v4) > 0 {
            v4[0] = fi5
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1() int32 {
    if (4294967295 == 687) {
        v6 := GenS1{F0: 42}
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint64 = 137
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    if (214 <= 97) {
        var v9 int16 = 42
        var v10 uint = 375
        _ = v10
        _ = v9
    }
    var v11 uint32 = 2
    var v12 string = "foo bar"
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 64
    v3 := make(map[uint32]uint32)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 bool = true
        if ((fi13 + fi13) > (^(fi13))) {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 uint16 = 32767
                _ = fi15
                _ = v16
            }
        } else {
            if false {
                v17 := []int{0, (fi13 % ((fi13) | 1)), 127, fi13}
                _ = v17
            } else {
                if v14 {
                    delete(v3, v2)
                } else {
                    v18 := GenS2{F0: 0, F1: 128, F2: 127}
                    if (!false) {
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v20 := GenS1{F0: 923}
                            v2 += (v2 * v2)
                            _ = fi19
                            _ = v20
                        }
                        v2 -= v2
                    }
                    if v14 {
                        fmt.Println(v2, fi13)
                        v21 := GenS1{F0: v2}
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v23 := GenS1{F0: (v2 << ((2) & 63))}
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    } else {
                        fmt.Println(v14, fi13)
                    }
                    _ = v18
                }
                var v24 uint = 1
                v25 := GenS3{F0: 1, F1: v14, F2: 138}
                _ = v24
                _ = v25
            }
        }
        _ = len(v1)
        _ = fi13
        _ = v14
    }
    v26 := GenS4{F0: 548}
    switch 2147483647 {
    case 0:
        if (18446744073709551615 > 1) {
            for fi27 := 0; fi27 < 2; fi27++ {
                v2 -= 2
                v2 = 2147483647
                _ = fi27
            }
            if (!true) {
                fmt.Println(v2)
                switch ((("foo bar" + (("hello" + "\n\t") + "\n\t")) + "a") + "a") {
                case "":
                    v28 := GenS5{F0: 1, F1: 1, F2: true}
                    _ = v28
                }
            } else {
                if (!true) {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v2)
                    case "foo bar":
                        for fi29 := 0; fi29 < 13; fi29++ {
                            v30 := []int{}
                            var v31 uint32 = (v2 >> ((8) & 63))
                            var v32 bool = true
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    case "hello":
                        v2 = v2
                    default:
                        if true {
                            if (!false) {
                                v33 := GenS3{F0: 1, F1: (39 > 42), F2: 42}
                                if ((v2 ^ (^(v2))) <= v2) {
                                    switch -417 {
                                    case 32768:
                                        v34 := GenS6{F0: -676, F1: 4294967295, F2: 1}
                                        _ = v34
                                    case 516:
                                        v35 := []int{2147483647, -1}
                                        _ = v35
                                    default:
                                        v2 = (v2 + 466)
                                    }
                                    var v36 int64 = -9223372036854775808
                                    _ = v36
                                }
                                _ = v33
                            }
                            v2 -= 339
                        } else {
                            if ((4 == 255) && (genFn1() < 2147483647)) {
                                var v37 uint8 = 8
                                var v38 uint16 = 65535
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                    var v39 string = "foo bar"
                    _ = v39
                }
                for fi40 := 0; fi40 < 19; fi40++ {
                    if false {
                        var v41 bool = (4294967295 == 560)
                        v2 += (^(v2))
                        if v41 {
                            var v42 string = "foo bar"
                            v41 = true
                            var v43 float64 = 1e10
                            _ = v42
                            _ = v43
                        } else {
                            v41 = false
                            v41 = (!true)
                        }
                        _ = v41
                    }
                    v3[v2] = (v2 << ((42) & 63))
                    var v44 int64 = 1
                    _ = fi40
                    _ = v44
                }
            }
            v45 := GenS7{F0: 42, F1: ("\n\t" + "foo bar")}
            _ = v45
        } else {
            if (!false) {
                switch -14 {
                case 42:
                    v2 = 2
                case 32768:
                    for fi46 := 0; fi46 < 10; fi46++ {
                        v2 = v2
                        for fi47 := 0; fi47 < 1; fi47++ {
                            fmt.Println(v2, fi47)
                            _ = fi47
                        }
                        fmt.Println(fi46, v2)
                        _ = fi46
                    }
                case 986:
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v2)
                    case "a":
                        v48 := GenS8{F0: -14, F1: -556, F2: ""}
                        _ = v48
                    case "hello":
                        v2 += v2
                    default:
                        v49 := GenS9{F0: 1, F1: 42, F2: 54}
                        _ = v49
                    }
                default:
                    if false {
                        if (((0.1 - (0.1 - float64(((float32((1e10 + float64((1e10 - (-2.25 * (0.1 + (0.1 / (float32(0.1) - 1.5)))))))) + 1.5) + (0.1 - 0.0))))) * (3.14159 * 1.5)) == 1e10) {
                            for fi50 := 0; fi50 < 11; fi50++ {
                                v51 := GenS10{F0: 2, F1: 4294967295, F2: 1}
                                v52 := GenS11{F0: 956, F1: 0, F2: 1}
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v53 := []int{32768, -936, 32767}
                            _ = v53
                        }
                        fmt.Println(v2)
                        v2 -= (v2 | v2)
                    } else {
                        if false {
                            delete(v1, "hello")
                        } else {
                            v54 := []int{}
                            _ = v54
                        }
                    }
                }
                fmt.Println(v2)
            }
            for fi55 := 0; fi55 < 14; fi55++ {
                if true {
                    var v56 int = (fi55 - fi55)
                    v57 := []int{}
                    _ = v56
                    _ = v57
                } else {
                    v2 = 0
                    for fi58 := 0; fi58 < 3; fi58++ {
                        v59 := GenS12{F0: 837, F1: ("" == ("a" + "hello")), F2: 42}
                        if false {
                            v2 = (v2 >> ((2) & 63))
                        }
                        _ = fi58
                        _ = v59
                    }
                    v2 -= (v2 % ((v2) | 1))
                }
                if true {
                    for fi60 := 0; fi60 < 8; fi60++ {
                        fmt.Println(fi55)
                        fmt.Println(fi55, fi55, fi60)
                        v61 := GenS13{F0: 777, F1: 18446744073709551615, F2: -1}
                        _ = fi60
                        _ = v61
                    }
                }
                var v62 float32 = 3.14159
                _ = fi55
                _ = v62
            }
            switch "a" {
            case "a":
                v63 := []int{-1, 9223372036854775807, 0, -486}
                _ = v63
            case "\n\t":
                if false {
                    var v64 string = ("hello" + "\n\t")
                    _ = v64
                }
            }
        }
    case -2147483648:
        if true {
            var v65 uint32 = (v2 << ((18446744073709551615) & 63))
            _ = v65
        }
    case 4294967295:
        var v66 float64 = 0.0
        _ = v66
    }
    var v67 uint = 255
    var v68 int8 = int8(v67)
    for fi69 := 0; fi69 < 8; fi69++ {
        if (!((((true && true) || (fi69 >= 128)) || false) || false)) {
            v67 = (v67 * uint(v2))
            v70 := v3[(v2 * (v2 % ((v2) | 1)))]
            _ = v70
        } else {
            v2 = (v2 >> ((v67) & 63))
            v71 := GenS11{F0: fi69, F1: 42, F2: 65535}
            if false {
                v72 := GenS14{F0: (v2 - (v2 * 2147483647)), F1: (fi69 - (fi69 | int(v68))), F2: uint32(v67)}
                _ = v72
            } else {
                if (32767 <= 32767) {
                    for fi73 := 0; fi73 < 7; fi73++ {
                        if false {
                            switch ("hello" + ("a" + "\n\t")) {
                            case "":
                                v74 := GenS3{F0: 874, F1: false, F2: 42}
                                _ = v74
                            case "hello":
                                v75 := []int{}
                                _ = v75
                            case "foo bar":
                                var v76 bool = (!false)
                                _ = v76
                            }
                            v77 := GenS15{F0: "hello"}
                            _ = v77
                        }
                        v2 = 42
                        _ = fi73
                    }
                    for fi78 := 0; fi78 < 5; fi78++ {
                        switch (-(fi78)) {
                        case 42:
                            v3[907] = v2
                            _ = len(v3)
                            v2 = 4294967295
                        default:
                            v79 := GenS16{F0: 55, F1: ((^(v67)) == v67), F2: -2147483648}
                            _ = v79
                        }
                        v67 -= (v67 * v67)
                        _ = fi78
                    }
                }
            }
            _ = v71
        }
        v80 := []int{}
        switch ("hello" + ("foo bar" + "a")) {
        case "hello":
            v3[(v2 * 2147483647)] = 42
        default:
            for fi81 := 0; fi81 < 19; fi81++ {
                var v82 uint32 = 823
                _ = fi81
                _ = v82
            }
        }
        _ = fi69
        _ = v80
    }
    var v83 int8 = v68
    if (0.1 > (1e10 + (float32(3.14159) + -2.25))) {
        var v84 int8 = genFn0(128)
        delete(v1, "foo bar")
        var v85 int32 = 42
        _ = v84
        _ = v85
    } else {
        var v86 uint = v67
        _ = v86
    }
    v87 := GenS17{F0: (v83 * v83), F1: 65535}
    fmt.Println(v2)
    var v88 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v83:", v83)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v67
    _ = v68
    _ = v83
    _ = v87
    _ = v88
}
