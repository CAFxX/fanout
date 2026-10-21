package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    var v1 uint = 2
    var v2 float64 = 0.0
    var v3 int = 1
    var v4 int64 = -751
    var v5 int16 = 1
    if (42 <= 73) {
        if (2 >= 0) {
            var v6 uint8 = 127
            var v7 string = ("hello" + ("" + "\n\t"))
            _ = v6
            _ = v7
        }
        var v8 int32 = -2147483648
        _ = v8
    }
    var v9 uint64 = 2
    var v10 uint32 = 2147483647
    switch (-(v3)) {
    case -792:
        v10 = (v10 - (v10 + v10))
    case 868:
        v5 = v5
    case 42:
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch fi13 {
        case -9223372036854775808:
            v14 := GenS1{F0: -129, F1: 2, F2: 94}
            _ = v14
        case -68:
            switch (((("a" + "a") + "foo bar") + ((("a" + "hello") + "hello") + "foo bar")) + "foo bar") {
            case "hello":
                var v15 float64 = (3.14159 * (-2.25 / v2))
                _ = v15
            default:
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int64 = v4
                    if false {
                        v18 := []int{}
                        _ = v18
                    } else {
                        v19 := []int{}
                        if (true || ((^(v3)) <= 635)) {
                            var v20 uint64 = v9
                            v20 += (v9 | 773)
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v3)
                            default:
                                v21 := GenS2{F0: 255, F1: 0, F2: int8(v10)}
                                _ = v21
                            }
                            if (2 < uint16(v1)) {
                                v22 := GenS3{F0: 0, F1: 423, F2: (v17 - (v17 << ((v1) & 63)))}
                                var v23 uint64 = (v20 * 4294967295)
                                v1 += 65535
                                _ = v22
                                _ = v23
                            } else {
                                v3 = (-(v3))
                            }
                            var v24 int8 = 21
                            if (-2147483648 != 1) {
                                if ("\n\t" == "hello") {
                                    v20 = v9
                                    var v25 string = "hello"
                                    _ = v25
                                }
                                v19 = append(v19, 42)
                                var v26 int = (fi16 ^ fi16)
                                _ = len(v19)
                                if true {
                                    v27 := GenS4{F0: false}
                                    _ = v27
                                }
                                _ = v26
                            }
                            if (v3 > (fi13 << ((v1) & 63))) {
                                v2 -= (v2 + v2)
                            } else {
                                fmt.Println(v3)
                                v2 -= 1.5
                                var v28 float64 = v2
                                _ = v28
                            }
                            var v29 string = "a"
                            _ = v20
                            _ = v24
                            _ = v29
                        }
                        if (!true) {
                            v2 = 1.5
                        } else {
                            var v30 float32 = 0.1
                            _ = v30
                        }
                        _ = v19
                    }
                    var v31 float32 = 1e10
                    var v32 uint16 = 1
                    _ = fi16
                    _ = v17
                    _ = v31
                    _ = v32
                }
            }
        case 4294967295:
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 string = "\n\t"
                var v35 int = 130
                var v36 uint32 = 718
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
        default:
            var v37 float32 = 3.14159
            _ = v37
        }
        _ = fi13
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 int32 = 42
        _ = fi38
        _ = v39
    }
    v10 = v10
    var v40 bool = ((((-2.25 > 0.1) && true) && (true && (!(!(true || false))))) || (2 <= uint16(v9)))
    v41 := GenS5{F0: 128, F1: v40}
    if false {
        var v42 uint8 = 2
        v43 := []int{(v3 | -1)}
        _ = v42
        _ = v43
    }
    switch v3 {
    case 42:
        v44 := GenS3{F0: v10, F1: 671, F2: (v4 | 42)}
        _ = v44
    case 1:
        if true {
            var v45 uint32 = v10
            v46 := []int{128, v3}
            var v47 float64 = ((v2 - v2) + v2)
            _ = v45
            _ = v46
            _ = v47
        } else {
            for fi48 := 0; fi48 < 17; fi48++ {
                for fi49 := 0; fi49 < 11; fi49++ {
                    var v50 int32 = 42
                    var v51 float64 = -2.25
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                for fi52 := 0; fi52 < 10; fi52++ {
                    var v53 float32 = 1.5
                    if ((605 == (v9 ^ 42)) && v40) {
                        v54 := []int{int(v10), (fi52 | (^(fi48)))}
                        v54 = append(v54, -226)
                        _ = v54
                    } else {
                        var v55 float64 = (v2 / 1e10)
                        switch (("a" + ("hello" + "hello")) + ("\n\t" + "hello")) {
                        case "a":
                            switch "hello" {
                            case "foo bar":
                                var v56 float64 = (v55 + v2)
                                _ = v56
                            case "a":
                                v5 += int16(v9)
                            default:
                                if v40 {
                                    if (!v40) {
                                        if (!true) {
                                            var v57 int32 = -2147483648
                                            v58 := GenS2{F0: 166, F1: 847, F2: 1}
                                            v59 := []int{fi48, (v3 | fi52), (fi52 % ((fi48) | 1))}
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v60 := GenS4{F0: v40}
                                            v3 = int(v1)
                                            switch "" {
                                            case "\n\t":
                                                v61 := []int{9223372036854775807, -103, 9223372036854775807, (fi52 + (fi52 << (((^(v1))) & 63)))}
                                                _ = v61
                                            case "foo bar":
                                                v2 -= (float64(0.1) / (v55 / -2.25))
                                            case "a":
                                                v2 += (v2 - ((float64((v53 - ((v53 * v53) * 0.1))) + ((v2 + (float64(v53) + (v2 + v55))) + v55)) * float64(float32(v55))))
                                            default:
                                                if ((v2 + 0.0) != (float64(v53) / v55)) {
                                                    var v62 string = "foo bar"
                                                    v63 := GenS4{F0: true}
                                                    v64 := GenS3{F0: v10, F1: 1, F2: (v4 << ((v1) & 63))}
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    v65 := []int{(fi48 << ((401) & 63)), (v3 * fi52), fi48}
                                                    _ = v65
                                                }
                                            }
                                            _ = v60
                                        }
                                    }
                                    fmt.Println(v5, fi48, v1)
                                }
                            }
                        case "":
                            v40 = (!v40)
                        }
                        if (70 < 42) {
                            switch "a" {
                            case "\n\t":
                                var v66 float64 = (1e10 - float64(((v53 / v53) + v53)))
                                _ = v66
                            case "hello":
                                v9 -= 1
                            default:
                                var v67 int8 = 47
                                _ = v67
                            }
                            fmt.Println(v40, v53, fi48)
                        }
                        _ = v55
                    }
                    _ = fi52
                    _ = v53
                }
                _ = fi48
            }
        }
    case -9223372036854775808:
        if (v40 || (v3 <= (v3 % (((v3 | int(v9))) | 1)))) {
            var v68 float32 = 0.1
            var v69 int32 = -2147483648
            fmt.Println(v2, v9)
            _ = v68
            _ = v69
        }
    default:
        if (255 > 255) {
            v70 := GenS6{F0: ((v4 >> ((v1) & 63)) != 0), F1: 1000}
            var v71 float64 = (v2 / 1e10)
            _ = v70
            _ = v71
        }
    }
    v10 += v10
    v4 = v4
    fmt.Println(v40, v3, v3)
    v3 -= v3
    v72 := GenS7{F0: 0, F1: 152, F2: false}
    if true {
        if (!(!((v9 >> (((v1 * (v1 * (^(v1))))) & 63)) != v9))) {
            for fi73 := 0; fi73 < 16; fi73++ {
                fmt.Println(v3, v1)
                fmt.Println(v4, v2, v1)
                _ = fi73
            }
            v40 = v40
        } else {
            v40 = v40
        }
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v72
    _ = v9
}
