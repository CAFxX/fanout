package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint32; F1 int32; F2 bool }
func genFn0() int16 {
    var v7 int8 = 1
    v7 = (v7 / ((110) | 1))
    v7 -= v7
    fmt.Println(v7)
    _ = v7
    return 0
}
func genFn1() uint16 {
    var v8 int8 = -128
    var v9 int8 = v8
    for fi10 := 0; fi10 < 1; fi10++ {
        if true {
            fmt.Println(v9)
            v11 := GenS1{F0: 42}
            _ = v11
        } else {
            v9 = 1
        }
        v12 := []int{fi10, (fi10 % ((334) | 1)), (fi10 * fi10)}
        _ = fi10
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 9223372036854775807}
        _ = fi13
        _ = v14
    }
    _ = v8
    _ = v9
    return 160
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]string)
    var v3 uint16 = 702
    v4 := make(map[uint32]bool)
    var v5 int64 = 42
    var v6 string = ("" + (("a" + "foo bar") + "a"))
    var v15 uint32 = 2147483648
    v4[v15] = (!(true || true))
    v15 -= uint32(v3)
    delete(v2, 199)
    var v16 string = ("\n\t" + v6)
    switch 522 {
    case -1:
        var v17 int32 = 978
        _ = v17
    case 4294967295:
        if true {
            if false {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v6 = v16
                    _ = fi18
                }
                fmt.Println(v15, v6)
                var v19 int16 = -454
                _ = v19
            } else {
                if true {
                    var v20 uint8 = 0
                    v3 = 0
                    _ = v20
                } else {
                    if false {
                        v21 := GenS1{F0: v1}
                        _ = v21
                    } else {
                        switch v6 {
                        case "a":
                            switch 32768 {
                            case -759:
                                v3 = 65535
                            case 9223372036854775807:
                                if (false || false) {
                                    v16 = (v6 + (v6 + (v6 + (v16 + v6))))
                                    v22 := GenS2{F0: (v15 * (v15 >> ((18446744073709551615) & 63))), F1: -1}
                                    v1 = -1
                                    _ = v22
                                }
                            }
                        }
                    }
                    switch 9223372036854775807 {
                    case -60:
                        if false {
                            switch "hello" {
                            case "":
                                v15 -= 222
                            case "hello":
                                v16 = v6
                            case "foo bar":
                                var v23 uint32 = 280
                                _ = v23
                            }
                            fmt.Println(v15, v15)
                        } else {
                            v6 = (v6 + v6)
                        }
                    default:
                        if true {
                            _ = len(v2)
                            v5 = v1
                            v24 := []int{159, 127, -21}
                            _ = v24
                        }
                    }
                    if (v16 <= v16) {
                        for fi25 := 0; fi25 < 1; fi25++ {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                v27 := GenS3{F0: int(v5)}
                                if false {
                                    v16 = v6
                                    v2[fi26] = v16
                                    switch "foo bar" {
                                    case "hello":
                                        var v28 uint8 = 42
                                        _ = v28
                                    default:
                                        if (1 <= 738) {
                                            if (((!((!(false && ((int32(v3) > int32(v3)) && false))) || (false || false))) || ((28 <= 2) || false)) || (float32(float64(3.14159)) == float32(((-2.25 / (-2.25 * 3.14159)) / 0.1)))) {
                                                var v29 uint8 = 127
                                                _ = v29
                                            } else {
                                                var v30 int = (fi26 >> ((4294967295) & 63))
                                                if (775 <= 505) {
                                                    var v31 int = int(v5)
                                                    _ = v31
                                                } else {
                                                    if (1 <= 42) {
                                                        var v32 int = fi25
                                                        v30 -= fi26
                                                        v5 -= (-(v1))
                                                        _ = v32
                                                    }
                                                    v1 -= (v1 - -1)
                                                    v33 := GenS4{F0: v5, F1: -2147483648}
                                                    v34 := GenS5{F0: 1}
                                                    _ = v33
                                                    _ = v34
                                                }
                                                var v35 bool = false
                                                if (1 >= 1) {
                                                    var v36 uint64 = 42
                                                    v37 := GenS1{F0: (v5 >> ((823) & 63))}
                                                    _ = v36
                                                    _ = v37
                                                }
                                                switch 2147483647 {
                                                case 808:
                                                    var v38 bool = (v35 || v35)
                                                    _ = v38
                                                case -840:
                                                    v6 = "\n\t"
                                                default:
                                                    var v39 string = v16
                                                    _ = v39
                                                }
                                                _ = v30
                                                _ = v35
                                            }
                                        }
                                    }
                                    v3 = (v3 & (^(v3)))
                                }
                                _ = fi26
                                _ = v27
                            }
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v41 := GenS3{F0: fi25}
                                switch v6 {
                                case "foo bar":
                                    if (961 <= 42) {
                                        if false {
                                            var v42 float64 = 1e10
                                            _ = v42
                                        } else {
                                            v3 = v3
                                            switch (fi40 >> ((4294967295) & 63)) {
                                            case 358:
                                                if true {
                                                    delete(v4, 2)
                                                    var v43 int = (fi40 | (fi25 / ((fi25) | 1)))
                                                    if (((v6 + "hello") >= "foo bar") && (false || (((3.14159 / 0.1) / (1.5 * -2.25)) != 0.0))) {
                                                        switch (("hello" + ("hello" + (("foo bar" + "\n\t") + v16))) + v16) {
                                                        case "":
                                                            delete(v4, 2)
                                                        default:
                                                            var v44 int = fi40
                                                            _ = v44
                                                        }
                                                    } else {
                                                        v45 := []int{}
                                                        _ = v45
                                                    }
                                                    _ = v43
                                                }
                                            case 4294967295:
                                                var v46 uint32 = (^(v15))
                                                _ = v46
                                            case 32768:
                                                switch (fi40 * int(v3)) {
                                                case 32767:
                                                    v47 := GenS6{F0: v5}
                                                    _ = v47
                                                case -2147483648:
                                                    var v48 int32 = 1
                                                    _ = v48
                                                }
                                                var v49 uint8 = 0
                                                fmt.Println(v3, v5, fi25)
                                                _ = v49
                                            default:
                                                fmt.Println(v1, fi25, v16)
                                            }
                                        }
                                    }
                                }
                                _ = fi40
                                _ = v41
                            }
                            _ = fi25
                        }
                        if (true && false) {
                            v2[2147483647] = v16
                            for fi50 := 0; fi50 < 6; fi50++ {
                                if (v16 < "") {
                                    v6 = "\n\t"
                                } else {
                                    if false {
                                        for fi51 := 0; fi51 < 9; fi51++ {
                                            v52 := GenS1{F0: -9223372036854775808}
                                            switch "foo bar" {
                                            case "hello":
                                                v53 := []int{fi50, 42}
                                                _ = v53
                                            }
                                            _ = fi51
                                            _ = v52
                                        }
                                    }
                                }
                                _ = fi50
                            }
                        }
                    }
                }
                for fi54 := 0; fi54 < 2; fi54++ {
                    v3 = 65535
                    if true {
                        for fi55 := 0; fi55 < 1; fi55++ {
                            v3 -= (v3 / ((v3) | 1))
                            _ = fi55
                        }
                    } else {
                        var v56 int8 = -1
                        delete(v4, v15)
                        _ = v56
                    }
                    delete(v2, fi54)
                    _ = fi54
                }
            }
        }
    case 2147483647:
        fmt.Println(v15, v3, v15)
    default:
        v57 := v4[(^(v15))]
        _ = v57
    }
    var v58 uint = 18446744073709551615
    if (0 < -1) {
        switch 599 {
        case -9223372036854775808:
            var v59 uint16 = genFn1()
            _ = v59
        case 1:
            var v60 bool = true
            _ = v60
        }
    } else {
        delete(v4, (v15 ^ (v15 << (((v58 / ((v58) | 1))) & 63))))
        var v61 float32 = (((float32((1e10 * 0.0)) / 1e10) + 0.0) + (1.5 + 1.5))
        for fi62 := 0; fi62 < 8; fi62++ {
            var v63 uint64 = 2
            for fi64 := 0; fi64 < 17; fi64++ {
                if true {
                    if false {
                        fmt.Println(v58, v15, v58)
                        var v65 string = v16
                        var v66 string = v16
                        _ = v65
                        _ = v66
                    }
                    var v67 uint8 = 0
                    fmt.Println(v58, v15)
                    _ = v67
                } else {
                    if false {
                        var v68 float64 = ((1.5 / 1.5) - 1.5)
                        fmt.Println(v6, v68, v58)
                        v1 = (v5 >> (((v58 / ((uint(v63)) | 1))) & 63))
                        _ = v68
                    } else {
                        switch (v16 + v16) {
                        case "hello":
                            switch v16 {
                            case "a":
                                fmt.Println(v63)
                            }
                        case "foo bar":
                            var v69 float64 = 0.1
                            _ = v69
                        }
                    }
                    v70 := []int{int(v3), fi64, (fi62 - int(v1)), int(v3)}
                    var v71 uint8 = 1
                    _ = v70
                    _ = v71
                }
                v72 := GenS7{F0: 1, F1: (v16 + "hello")}
                _ = fi64
                _ = v72
            }
            var v73 uint64 = v63
            _ = fi62
            _ = v63
            _ = v73
        }
        _ = v61
    }
    if (-1 > 268) {
        if true {
            var v74 int32 = 188
            var v75 uint32 = 42
            _ = v74
            _ = v75
        }
        v4[v15] = (false && ((1 != 18446744073709551615) || true))
        v76 := GenS1{F0: (-(v1))}
        _ = v76
    }
    v77 := v2[192]
    var v78 int64 = v5
    for fi79 := 0; fi79 < 6; fi79++ {
        fmt.Println(v77)
        v80 := GenS7{F0: 1, F1: v6}
        _ = fi79
        _ = v80
    }
    fmt.Println(v58, v58)
    if false {
        v81 := v4[(v15 - (v15 & v15))]
        v15 = 0
        _ = v81
    } else {
        if true {
            var v82 uint64 = 18446744073709551615
            fmt.Println(v15)
            v83 := GenS8{F0: 4294967295, F1: -409, F2: false}
            _ = v82
            _ = v83
        }
    }
    v78 = v78
    var v84 int = -2147483648
    v85 := GenS5{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v77
    _ = v78
    _ = v84
    _ = v85
}
