package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS7 struct { F0 int64; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 bool) int16 {
    p0 = (("" <= "a") && (0 < 2))
    for fi4 := 0; fi4 < 2; fi4++ {
        if (!p0) {
            var v5 uint64 = 0
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint64 = v5
                switch fi6 {
                case 42:
                    v8 := []int{fi6, 32767}
                    _ = v8
                case -9223372036854775808:
                    var v9 int = int(v5)
                    _ = v9
                }
                if (92 != 91) {
                    if ((-9223372036854775808 < -1) || (2 < 121)) {
                        v7 = v7
                        v10 := GenS1{F0: 1, F1: 553}
                        _ = v10
                    } else {
                        v5 = v7
                    }
                    v11 := GenS2{F0: 42, F1: 148, F2: 2}
                    v12 := []int{}
                    if len(v12) > 0 {
                        v12[0] = fi4
                    }
                    v13 := GenS3{F0: (!p0), F1: -905}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 float32 = (float32(1e10) * (3.14159 - 1e10))
                var v15 int = -1
                _ = fi6
                _ = v14
                _ = v15
                _ = v7
            }
            _ = v5
        } else {
            p0 = (0 == -991)
            fmt.Println(fi4, fi4)
            var v16 uint16 = 42
            _ = v16
        }
        switch fi4 {
        case 42:
            var v17 float32 = -2.25
            _ = v17
        default:
            p0 = (p0 || p0)
        }
        _ = fi4
    }
    _ = p0
    return 42
}
func genFn1(p0 int64, p1 uint8) int {
    p0 = p0
    var v18 int16 = genFn0(false)
    _ = p0
    _ = p1
    _ = v18
    return 1
}

func main() {
    var v1 bool = (4294967295 < 32768)
    var v2 int16 = -32768
    v3 := []string{("foo bar" + "foo bar"), "a"}
    if false {
        var v19 float64 = float64(((3.14159 - 3.14159) + 1.5))
        _ = v19
    }
    var v20 uint16 = 257
    var v21 int64 = 1
    if ((v2 & (v2 * v2)) == genFn0(v1)) {
        v22 := GenS4{F0: 4294967295, F1: (v2 <= genFn0((v21 > (-(v21)))))}
        _ = v22
    }
    v1 = ((("\n\t" + "a") + "hello") >= ("\n\t" + "foo bar"))
    fmt.Println(v20, v2, v20)
    switch "foo bar" {
    case "foo bar":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 float64 = 0.1
            if (2 == 2) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "hello" {
                case "":
                    if (v21 <= (v21 >> ((65535) & 63))) {
                        switch fi23 {
                        case 128:
                            v21 = v21
                        case 2147483647:
                            if (!((fi23 * -485) != (fi23 + (fi23 >> ((0) & 63))))) {
                                if (1e10 > 0.1) {
                                    v21 -= (v21 % ((v21) | 1))
                                    switch genFn1(v21, 2) {
                                    case 127:
                                        switch (fi23 | 239) {
                                        case -380:
                                            if (true && v1) {
                                                fmt.Println(v24)
                                                for fi25 := 0; fi25 < 11; fi25++ {
                                                    var v26 int8 = 127
                                                    var v27 int = fi23
                                                    v1 = v1
                                                    _ = fi25
                                                    _ = v26
                                                    _ = v27
                                                }
                                            }
                                        case 32767:
                                            for fi28 := 0; fi28 < 17; fi28++ {
                                                v24 = v24
                                                if v1 {
                                                    var v29 string = "\n\t"
                                                    var v30 uint8 = 37
                                                    var v31 float32 = (-2.25 * 1.5)
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v1 = (-2147483648 >= 1)
                                                    fmt.Println(v21, fi23)
                                                }
                                                _ = fi28
                                            }
                                        case 42:
                                            if true {
                                                v20 += (v20 << ((42) & 63))
                                                v32 := GenS3{F0: v1, F1: v21}
                                                for fi33 := 0; fi33 < 20; fi33++ {
                                                    if (58 != 0) {
                                                        var v34 float32 = 3.14159
                                                        _ = v34
                                                    } else {
                                                        v2 = genFn0(v1)
                                                        switch ("\n\t" + "\n\t") {
                                                        case "\n\t":
                                                            v35 := GenS5{F0: 380}
                                                            _ = v35
                                                        case "":
                                                            v24 = (v24 / ((v24 - float64(1e10)) - v24))
                                                        case "a":
                                                            if true {
                                                                var v36 int16 = v2
                                                                _ = v36
                                                            } else {
                                                                var v37 int16 = v2
                                                                if (!true) {
                                                                    v37 = genFn0((true || (!("hello" >= "foo bar"))))
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v38 int = fi33
                                                                        _ = v38
                                                                    case "":
                                                                        var v39 int8 = -128
                                                                        _ = v39
                                                                    case "\n\t":
                                                                        if (!((2147483648 < 42) && v1)) {
                                                                            var v40 int = fi33
                                                                            v41 := GenS6{F0: 1, F1: 4294967295, F2: -455}
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                        var v42 int32 = 917
                                                                        _ = v42
                                                                    }
                                                                    var v43 int64 = v21
                                                                    _ = v43
                                                                }
                                                                var v44 uint8 = 0
                                                                _ = v37
                                                                _ = v44
                                                            }
                                                        }
                                                        fmt.Println(v24, v20, fi33)
                                                    }
                                                    var v45 float64 = float64(1.5)
                                                    _ = len(v3)
                                                    _ = fi33
                                                    _ = v45
                                                }
                                                _ = v32
                                            } else {
                                                var v46 int = 42
                                                _ = v46
                                            }
                                        }
                                    default:
                                        if true {
                                            _ = len(v3)
                                        }
                                    }
                                }
                                if v1 {
                                    v47 := GenS2{F0: 0, F1: 841, F2: 255}
                                    var v48 float64 = (v24 + ((1.5 / v24) / v24))
                                    v1 = v1
                                    _ = v47
                                    _ = v48
                                }
                            }
                        case 32768:
                            switch genFn1((v21 & v21), 250) {
                            case 2147483647:
                                if false {
                                    v1 = (2 <= 4294967295)
                                } else {
                                    fmt.Println(v21)
                                    v49 := GenS4{F0: 0, F1: false}
                                    _ = v49
                                }
                            case -164:
                                var v50 uint = 18446744073709551615
                                _ = v50
                            default:
                                for fi51 := 0; fi51 < 15; fi51++ {
                                    v52 := GenS1{F0: 110, F1: 2147483648}
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        default:
                            var v53 int32 = 2147483647
                            _ = v53
                        }
                    } else {
                        _ = len(v3)
                        switch (fi23 | genFn1(v21, 1)) {
                        case 128:
                            var v54 uint32 = 829
                            _ = v54
                        case 142:
                            if (v20 >= 1) {
                                v55 := GenS4{F0: 65535, F1: v1}
                                var v56 float64 = (v24 * 1.5)
                                v57 := GenS7{F0: (v21 >> ((4294967295) & 63)), F1: 255, F2: 0}
                                _ = v55
                                _ = v56
                                _ = v57
                            } else {
                                for fi58 := 0; fi58 < 18; fi58++ {
                                    var v59 uint32 = 65535
                                    var v60 uint32 = 2
                                    v61 := GenS8{F0: (v20 ^ (v20 << ((4294967295) & 63)))}
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                            }
                        case -9223372036854775808:
                            v20 = 65535
                        }
                        for fi62 := 0; fi62 < 17; fi62++ {
                            if false {
                                v24 = float64(0.0)
                            }
                            _ = fi62
                        }
                    }
                default:
                    var v63 int8 = 100
                    _ = v63
                }
                fmt.Println(v24, v2, fi23)
            } else {
                var v64 int = (fi23 / (((fi23 * fi23)) | 1))
                var v65 uint = 42
                _ = v64
                _ = v65
            }
            _ = fi23
            _ = v24
        }
    case "\n\t":
        var v66 bool = (v1 || (!(true || v1)))
        _ = v66
    default:
        v2 = (v2 >> ((2) & 63))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v67 float32 = (1.5 / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v67
}
