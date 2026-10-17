package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int16; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS7 struct { F0 bool; F1 int; F2 bool }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 bool; F1 uint16; F2 int64 }
type GenS12 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() uint16 {
    var v7 uint32 = 2147483647
    var v8 int = 553
    var v9 uint8 = 128
    if true {
        var v10 uint16 = 32767
        if ((v8 << ((18446744073709551615) & 63)) >= v8) {
            if true {
                if ((352 != 533) && true) {
                    v8 = v8
                    switch "hello" {
                    case "hello":
                        v11 := GenS1{F0: 100}
                        _ = v11
                    case "foo bar":
                        v12 := []int{(v8 ^ 0), (v8 / (((v8 | v8)) | 1)), (v8 * (v8 ^ (v8 >> ((42) & 63))))}
                        _ = v12
                    case "a":
                        for fi13 := 0; fi13 < 2; fi13++ {
                            v14 := GenS2{F0: 1, F1: 1}
                            var v15 float64 = 3.14159
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                } else {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v8 -= fi16
                        _ = fi16
                    }
                    fmt.Println(v10, v8)
                }
            } else {
                var v17 uint64 = 4294967295
                v17 += 133
                var v18 int16 = 0
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                if (1 == 37) {
                    var v20 string = ""
                    v21 := GenS1{F0: 2}
                    _ = v20
                    _ = v21
                }
                v10 = 0
                _ = fi19
            }
        }
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 0
    var v2 string = ""
    var v3 int8 = -128
    v4 := make(map[uint32]uint32)
    var v5 int16 = 32767
    var v6 uint8 = 127
    v3 -= v3
    if true {
        if true {
            var v22 uint8 = v6
            v3 += (v3 / ((v3) | 1))
            switch "hello" {
            case "\n\t":
                var v23 uint64 = 1
                _ = v23
            }
            _ = v22
        }
        fmt.Println(v1, v2, v6)
    }
    delete(v4, 275)
    v3 -= 14
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS1{F0: (^(v3))}
        var v26 uint16 = v1
        var v27 uint8 = (^(v6))
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 int64 = int64(v3)
    fmt.Println(v1, v3, v5)
    if true {
        var v29 int32 = 42
        v30 := GenS1{F0: v3}
        v31 := []int{}
        _ = v29
        _ = v30
        _ = v31
    } else {
        v1 = genFn0()
        v1 = (v1 * 0)
        switch -9223372036854775808 {
        case 42:
            if ((0.0 <= (((3.14159 / float64(1.5)) / (float64(1e10) - 1.5)) + float64(float32(1e10)))) || (("foo bar" + "a") >= v2)) {
                v5 += (v5 ^ (v5 >> ((0) & 63)))
                v28 -= 9223372036854775807
                if true {
                    fmt.Println(v3, v1, v5)
                } else {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        var v33 bool = (-254 >= -2147483648)
                        if v33 {
                            v4[uint32(v1)] = 1
                            v4[2147483647] = 2147483648
                        } else {
                            switch v2 {
                            case "hello":
                                fmt.Println(fi32)
                            }
                        }
                        if v33 {
                            v34 := GenS3{F0: (v28 - (v28 ^ (v28 + v28))), F1: v2, F2: (v6 + (v6 & 8))}
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 7; fi35++ {
                                v1 = 65535
                                var v36 float64 = (((float64(0.0) / (0.0 - (1.5 - (1e10 / 1.5)))) / 1.5) * -2.25)
                                v37 := GenS4{F0: (v5 / (((v5 ^ v5)) | 1))}
                                if (!v33) {
                                    var v38 uint32 = 699
                                    v39 := []int{(^(fi35)), 433, (fi35 % ((fi35) | 1)), 32767}
                                    var v40 string = v2
                                    _ = len(v4)
                                    if (true || (-32768 >= (v5 - v5))) {
                                        v6 = (^(v6))
                                        if ((("\n\t" > v40) || ((^(v3)) < (v3 << ((1) & 63)))) || (int8(fi35) > 0)) {
                                            var v41 uint32 = (v38 << ((42) & 63))
                                            v42 := GenS5{F0: (v3 / (((v3 % (((v3 >> ((2) & 63))) | 1))) | 1)), F1: (v5 % ((v5) | 1)), F2: (-(v28))}
                                            v43 := GenS3{F0: (v28 >> ((1) & 63)), F1: v2, F2: (^(v6))}
                                            v38 += uint32(v6)
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        } else {
                                            switch (v2 + (v40 + v40)) {
                                            case "":
                                                v44 := GenS6{F0: v38, F1: (v6 / ((255) | 1)), F2: 1}
                                                _ = v44
                                            case "a":
                                                v45 := GenS1{F0: (v3 << ((595) & 63))}
                                                _ = v45
                                            case "hello":
                                                v6 -= (v6 / (((v6 >> ((42) & 63))) | 1))
                                            default:
                                                fmt.Println(v1, v40, v36)
                                                v3 -= int8(fi35)
                                            }
                                            fmt.Println(v28, v36)
                                            fmt.Println(v3)
                                            var v46 uint64 = 18
                                            _ = v46
                                        }
                                    } else {
                                        var v47 uint64 = 18446744073709551615
                                        if (v33 || ((v40 + v2) > v40)) {
                                            v4[v38] = 234
                                            v48 := GenS7{F0: v33, F1: 688, F2: true}
                                            var v49 int = 4294967295
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 float64 = float64(3.14159)
                                            v38 = (v38 % (((v38 << ((255) & 63))) | 1))
                                            _ = v50
                                        }
                                        v51 := GenS8{F0: (v5 % ((v5) | 1)), F1: 146}
                                        _ = v47
                                        _ = v51
                                    }
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                switch (v2 + "a") {
                                case "\n\t":
                                    switch ("a" + v2) {
                                    case "foo bar":
                                        fmt.Println(v28, fi35, fi35)
                                        if v33 {
                                            var v52 float32 = 0.0
                                            _ = v52
                                        } else {
                                            _ = len(v4)
                                            fmt.Println(v36, v2)
                                            v2 = v2
                                        }
                                        v53 := GenS9{F0: 1, F1: 2147483647, F2: 4294967295}
                                        _ = v53
                                    case "hello":
                                        var v54 float32 = 3.14159
                                        _ = v54
                                    case "\n\t":
                                        var v55 uint = 255
                                        _ = v55
                                    default:
                                        switch v2 {
                                        case "\n\t":
                                            fmt.Println(fi35)
                                        case "":
                                            v3 = v3
                                        default:
                                            v33 = ((v1 - (v1 | v1)) <= (^(v1)))
                                        }
                                    }
                                default:
                                    var v56 uint32 = 725
                                    _ = v56
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            var v57 uint32 = 2147483648
                            _ = v57
                        }
                        _ = fi32
                        _ = v33
                    }
                    if (true && (true && ((-32768 < v5) && true))) {
                        var v58 float32 = (3.14159 * 0.1)
                        _ = v58
                    }
                }
            } else {
                v59 := []int{-282, 9223372036854775807, int(v6)}
                v2 = v2
                v1 += 0
                _ = v59
            }
        case 127:
            for fi60 := 0; fi60 < 18; fi60++ {
                var v61 uint = 0
                switch v2 {
                case "":
                    for fi62 := 0; fi62 < 19; fi62++ {
                        v63 := GenS3{F0: v28, F1: (v2 + ("a" + v2)), F2: 128}
                        v64 := GenS10{F0: uint8(v61), F1: (^(v5))}
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                case "a":
                    var v65 int8 = v3
                    _ = v65
                default:
                    switch (fi60 ^ 1) {
                    case -493:
                        v28 = (v28 & v28)
                    case -2147483648:
                        v66 := []int{(fi60 & 842), (fi60 % (((fi60 ^ fi60)) | 1)), (fi60 - fi60), fi60}
                        _ = v66
                    case 42:
                        v3 = (^(v3))
                    default:
                        if false {
                            var v67 uint32 = 1
                            if ((v1 + (v1 >> ((v61) & 63))) > (v1 & v1)) {
                                fmt.Println(v3)
                                switch ((("a" + v2) + ("hello" + v2)) + (v2 + v2)) {
                                case "hello":
                                    fmt.Println(v6, v3, v67)
                                case "\n\t":
                                    v68 := []int{fi60, (fi60 - (fi60 & (^(fi60))))}
                                    _ = v68
                                case "":
                                    v69 := GenS10{F0: v6, F1: v5}
                                    _ = v69
                                }
                                v70 := GenS11{F0: false, F1: uint16(v61), F2: v28}
                                _ = v70
                            }
                            if (true || true) {
                                if (42 != 1) {
                                    fmt.Println(v61, v1)
                                    fmt.Println(v1, v2)
                                }
                                var v71 int64 = v28
                                var v72 uint32 = (v67 << (((v61 + 42)) & 63))
                                _ = v71
                                _ = v72
                            }
                            _ = v67
                        } else {
                            v73 := GenS7{F0: ((65535 >= 2147483647) || false), F1: int(v6), F2: true}
                            v74 := []int{(fi60 >> ((v61) & 63)), fi60, int(v61)}
                            v4[409] = 2147483648
                            _ = v73
                            _ = v74
                        }
                    }
                }
                _ = fi60
                _ = v61
            }
        }
    }
    var v75 float32 = 3.14159
    var v76 uint8 = v6
    v77 := GenS9{F0: -744, F1: 2147483647, F2: 4294967295}
    var v78 float32 = v75
    v79 := GenS8{F0: v5, F1: 18446744073709551615}
    var v80 uint32 = 4294967295
    fmt.Println(v6)
    for fi81 := 0; fi81 < 10; fi81++ {
        v4[v80] = v80
        _ = fi81
    }
    if false {
        v82 := []int{}
        v2 = v2
        v83 := GenS12{F0: -1, F1: v76, F2: 42}
        _ = v82
        _ = v83
    }
    for fi84 := 0; fi84 < 18; fi84++ {
        switch fi84 {
        case 42:
            switch v2 {
            case "a":
                v85 := []int{32768, (fi84 >> ((19) & 63)), 32767, (fi84 | fi84)}
                _ = v85
            case "hello":
                var v86 int8 = -127
                _ = v86
            case "":
                var v87 int32 = 1
                _ = v87
            default:
                for fi88 := 0; fi88 < 17; fi88++ {
                    v80 += v80
                    v89 := []int{}
                    v28 = (v28 << ((0) & 63))
                    _ = fi88
                    _ = v89
                }
            }
        case 127:
            for fi90 := 0; fi90 < 16; fi90++ {
                fmt.Println(v2)
                _ = fi90
            }
        }
        _ = fi84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
