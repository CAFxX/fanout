package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int8; F1 uint32; F2 uint32 }

func main() {
    var v1 int32 = 903
    v2 := make(map[int]string)
    var v3 int = 849
    var v4 uint8 = 72
    for fi5 := 0; fi5 < 10; fi5++ {
        if (2 > 42) {
            if (2 > 4294967295) {
                if (444 != 0) {
                    v6 := GenS1{F0: (fi5 | 127), F1: -933}
                    v7 := GenS1{F0: (v3 >> ((5) & 63)), F1: 1}
                    _ = v6
                    _ = v7
                } else {
                    fmt.Println(fi5, v3)
                    var v8 bool = false
                    _ = v8
                }
            }
            if true {
                if ((18446744073709551615 <= 69) && (2 != 377)) {
                    v9 := []int{}
                    v1 += (-(v1))
                    _ = v9
                } else {
                    var v10 float64 = ((float64(1.5) + 1.5) + 0.0)
                    v11 := v2[(fi5 + 4294967295)]
                    _ = v10
                    _ = v11
                }
            } else {
                var v12 uint32 = 65535
                var v13 uint16 = 608
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                v1 -= (v1 + (v1 + 242))
                _ = fi14
            }
        }
        v15 := GenS1{F0: v3, F1: 1000}
        var v16 uint32 = 4294967295
        _ = fi5
        _ = v15
        _ = v16
    }
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            switch ("hello" + "hello") {
            case "foo bar":
                v19 := []int{(fi18 & (fi18 * 0))}
                _ = v19
            case "hello":
                if false {
                    fmt.Println(v3, v3, v4)
                    v2[fi18] = "hello"
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v1 = -2147483648
                        if ((1e10 + 1e10) != 1e10) {
                            v3 -= fi18
                        }
                        v4 += (v4 ^ (v4 << ((1) & 63)))
                        _ = fi20
                    }
                    v21 := GenS2{F0: 0}
                    switch 0 {
                    case 32767:
                        v17 -= v17
                    }
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 19; fi22++ {
                    v1 -= (v1 ^ (^(v1)))
                    _ = fi22
                }
            }
            v23 := []int{0, (v3 - (v3 - 850)), 127}
            for fi24 := 0; fi24 < 10; fi24++ {
                v2[9223372036854775807] = ""
                v25 := GenS3{F0: 32767, F1: 127, F2: v1}
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        var v26 int8 = 127
        switch "hello" {
        case "foo bar":
            v27 := GenS1{F0: 914, F1: 1}
            _ = v27
        case "a":
            v28 := GenS2{F0: v26}
            _ = v28
        case "":
            v2[4294967295] = "hello"
        default:
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v4, v3)
                var v30 int = 0
                v31 := v2[(^(fi29))]
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = fi18
        _ = v26
    }
    v2[-325] = "foo bar"
    v32 := GenS2{F0: 127}
    var v33 bool = true
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 string = ("a" + "hello")
            if (!v33) {
                switch (fi34 << ((350) & 63)) {
                case -164:
                    fmt.Println(v36, v33)
                }
            }
            if v33 {
                var v37 float32 = (((3.14159 + ((float32(1.5) + float32(3.14159)) * 1.5)) * (-2.25 / 1.5)) * 0.0)
                fmt.Println(v4, v3)
                fmt.Println(v36, v3, v37)
                _ = v37
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        switch 128 {
        case 32767:
            v39 := GenS4{F0: "hello", F1: (v1 & 748), F2: 66}
            _ = v39
        case -1:
            if v33 {
                v2[(fi38 << ((1) & 63))] = ""
                v40 := []int{fi38, (fi38 ^ (v3 ^ v3)), (fi38 | (fi38 - fi38)), 128}
                _ = v40
            } else {
                v17 = (v17 + v17)
                for fi41 := 0; fi41 < 7; fi41++ {
                    var v42 string = "foo bar"
                    var v43 int16 = 32767
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                var v44 int32 = v1
                _ = v44
            }
        }
        _ = fi38
    }
    fmt.Println(v17, v3)
    v45 := GenS5{F0: v33}
    v46 := GenS6{F0: 2, F1: 42}
    switch 972 {
    case 2147483647:
        v47 := GenS7{F0: 747, F1: -901}
        _ = v47
    case 128:
        if (-1 != 100) {
            var v48 float64 = float64(-2.25)
            _ = v48
        } else {
            v49 := v2[-2147483648]
            for fi50 := 0; fi50 < 17; fi50++ {
                fmt.Println(v17, v17, fi50)
                _ = fi50
            }
            v17 -= 1
            _ = v49
        }
    case 1:
        for fi51 := 0; fi51 < 17; fi51++ {
            v1 += v1
            _ = fi51
        }
    }
    var v52 int8 = int8(v1)
    delete(v2, v3)
    for fi53 := 0; fi53 < 1; fi53++ {
        for fi54 := 0; fi54 < 12; fi54++ {
            if ("hello" <= ("foo bar" + "\n\t")) {
                v55 := GenS8{F0: v52, F1: 65535, F2: 2}
                if true {
                    fmt.Println(v33)
                } else {
                    fmt.Println(fi54)
                    fmt.Println(v3, fi54)
                }
                _ = v55
            } else {
                v56 := GenS3{F0: (v3 >> ((0) & 63)), F1: (v52 % (((v52 * (v52 * 42))) | 1)), F2: (v1 << ((906) & 63))}
                v3 = (fi53 * fi53)
                var v57 float64 = float64(1.5)
                _ = v56
                _ = v57
            }
            _ = fi54
        }
        switch -1 {
        case 9223372036854775807:
            var v58 float64 = 0.1
            _ = v58
        case 0:
            for fi59 := 0; fi59 < 1; fi59++ {
                v60 := GenS7{F0: -9223372036854775808, F1: -32768}
                if (v33 && true) {
                    v61 := []int{128}
                    v17 += v17
                    _ = v61
                }
                _ = fi59
                _ = v60
            }
        case 42:
            v52 = (v52 / (((^(v52))) | 1))
        }
        switch -700 {
        case 127:
            var v62 int = (fi53 << ((42) & 63))
            _ = v62
        case 1:
            fmt.Println(v3)
        default:
            for fi63 := 0; fi63 < 2; fi63++ {
                v64 := v2[-2147483648]
                switch int(v17) {
                case 752:
                    var v65 string = (("\n\t" + v64) + ("a" + v64))
                    _ = v65
                case -9223372036854775808:
                    v66 := GenS4{F0: "a", F1: v1, F2: 100}
                    _ = v66
                default:
                    v1 -= v1
                }
                _ = fi63
                _ = v64
            }
        }
        _ = fi53
    }
    fmt.Println(v1, v33, v1)
    fmt.Println(v17, v33, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v46
    _ = v52
}
