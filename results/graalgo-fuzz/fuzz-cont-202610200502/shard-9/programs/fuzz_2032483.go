package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16; F1 string; F2 bool }
func genFn0() uint8 {
    switch "" {
    case "hello":
        if ((float64((-2.25 / 1.5)) > 0.0) || (!true)) {
            if false {
                v6 := GenS1{F0: ("hello" + ("foo bar" + "a")), F1: 42}
                v7 := GenS1{F0: "a", F1: 1}
                _ = v6
                _ = v7
            } else {
                var v8 int8 = 42
                _ = v8
            }
        }
    case "":
        if (1.5 <= ((0.1 - 1e10) / 1.0)) {
            var v9 uint16 = 129
            v9 = v9
            if true {
                v10 := []int{}
                _ = len(v10)
                _ = v10
            }
            _ = v9
        }
    default:
        v11 := GenS2{F0: 42}
        _ = v11
    }
    v12 := []int{0, -1, 32767, 4294967295}
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS3{F0: (!false)}
        if false {
            var v15 int8 = -41
            v16 := GenS3{F0: false}
            _ = v15
            _ = v16
        }
        switch "" {
        case "\n\t":
            if (2 == 330) {
                v17 := GenS3{F0: true}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint16 = 42
                    fmt.Println(fi18, v19, v19)
                    switch "foo bar" {
                    case "a":
                        v19 -= 42
                    case "\n\t":
                        var v20 int32 = 2147483647
                        _ = v20
                    }
                    _ = len(v12)
                    if (1 <= 122) {
                        var v21 string = "hello"
                        _ = len(v12)
                        var v22 uint8 = 224
                        _ = v21
                        _ = v22
                    } else {
                        if true {
                            fmt.Println(fi18, fi18, fi18)
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
        case "hello":
            v23 := GenS4{F0: true, F1: 255}
            _ = v23
        default:
            var v24 float32 = (((-2.25 / 1.0) + 0.0) - 0.1)
            _ = v24
        }
        _ = fi13
        _ = v14
    }
    v25 := GenS5{F0: 32767}
    _ = v12
    _ = v25
    return 255
}
func genFn1(p0 string) int {
    if (!((false && ((981 > 42) && true)) || ((!true) || false))) {
        p0 = ("hello" + (p0 + "foo bar"))
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := []int{-9223372036854775808, 9223372036854775807}
            _ = fi26
            _ = v27
        }
        v28 := []int{128, 770}
        _ = v28
    }
    switch p0 {
    case "foo bar":
        v29 := []int{323}
        _ = v29
    case "hello":
        switch (p0 + p0) {
        case "foo bar":
            switch ("\n\t" + (p0 + p0)) {
            case "\n\t":
                p0 = "a"
            case "a":
                p0 = p0
            default:
                switch "hello" {
                case "foo bar":
                    v30 := GenS6{F0: p0, F1: 0}
                    _ = v30
                case "hello":
                    var v31 uint8 = genFn0()
                    _ = v31
                default:
                    var v32 string = (p0 + p0)
                    _ = v32
                }
            }
        default:
            p0 = p0
        }
    case "a":
        var v33 uint16 = 2
        _ = v33
    default:
        var v34 uint8 = 1
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := []int{fi35, 42}
        _ = fi35
        _ = v36
    }
    var v37 int = -353
    for fi38 := 0; fi38 < 1; fi38++ {
        switch (fi38 - fi38) {
        case 0:
            var v39 int16 = 0
            _ = v39
        case -9223372036854775808:
            for fi40 := 0; fi40 < 10; fi40++ {
                fmt.Println(fi40)
                if true {
                    if true {
                        var v41 uint = 65535
                        var v42 int32 = 0
                        v37 = 0
                        _ = v41
                        _ = v42
                    } else {
                        fmt.Println(v37, fi38)
                        p0 = "hello"
                        fmt.Println(p0, fi38, p0)
                    }
                    var v43 int64 = 9223372036854775807
                    fmt.Println(v37)
                    _ = v43
                }
                _ = fi40
            }
        default:
            if true {
                switch fi38 {
                case -919:
                    switch (v37 + v37) {
                    case 127:
                        v44 := GenS2{F0: 718}
                        _ = v44
                    case 9223372036854775807:
                        v45 := []int{v37, (v37 >> ((2) & 63)), (fi38 * (^(fi38)))}
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 14; fi46++ {
                            v47 := []int{fi38, 377, (fi38 >> ((558) & 63)), v37}
                            _ = fi46
                            _ = v47
                        }
                    }
                case 32767:
                    fmt.Println(p0, fi38, v37)
                }
            } else {
                v48 := GenS7{F0: int8(v37)}
                if (!(true && true)) {
                    switch (v37 * -1) {
                    case 2147483647:
                        v37 = (fi38 - (fi38 % ((0) | 1)))
                    case 67:
                        v49 := []int{9223372036854775807, fi38, 127}
                        _ = v49
                    case 9223372036854775807:
                        var v50 uint32 = 65535
                        _ = v50
                    default:
                        if false {
                            for fi51 := 0; fi51 < 6; fi51++ {
                                var v52 float64 = 0.1
                                _ = fi51
                                _ = v52
                            }
                            v53 := GenS8{F0: -2147483648, F1: fi38}
                            p0 = p0
                            _ = v53
                        } else {
                            if (((1.5 + 1.5) / float32((float64(0.1) / 0.1))) >= float32(-2.25)) {
                                var v54 uint16 = 658
                                _ = v54
                            }
                        }
                    }
                    switch fi38 {
                    case 128:
                        var v55 uint8 = 255
                        _ = v55
                    case 32768:
                        switch p0 {
                        case "foo bar":
                            p0 = ((p0 + (("a" + (("foo bar" + "a") + p0)) + ("" + "hello"))) + ("" + (p0 + (p0 + (p0 + (p0 + p0))))))
                        case "":
                            v56 := GenS9{F0: 1, F1: 255, F2: 255}
                            _ = v56
                        }
                    }
                    v37 = (fi38 / ((0) | 1))
                } else {
                    for fi57 := 0; fi57 < 11; fi57++ {
                        var v58 int64 = -1
                        if true {
                            v59 := []int{32768, 1, (v37 >> ((65535) & 63)), -851}
                            _ = v59
                        } else {
                            v37 -= (fi38 ^ (fi38 ^ (v37 ^ (-(fi57)))))
                            v60 := GenS10{F0: 42}
                            v61 := GenS6{F0: p0, F1: v58}
                            _ = v60
                            _ = v61
                        }
                        _ = fi57
                        _ = v58
                    }
                }
                _ = v48
            }
        }
        switch (fi38 % (((v37 | fi38)) | 1)) {
        case 130:
            v62 := GenS11{F0: -813, F1: p0, F2: (0 != 18446744073709551615)}
            _ = v62
        case -380:
            var v63 uint16 = 255
            _ = v63
        default:
            if false {
                fmt.Println(fi38)
            } else {
                var v64 uint64 = 42
                _ = v64
            }
        }
        p0 = "a"
        _ = fi38
    }
    _ = p0
    _ = v37
    return (v37 - v37)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    var v3 uint32 = 42
    v4 := make(map[int]uint32)
    var v5 int = 128
    fmt.Println(v5, v5, v5)
    v3 = 0
    v65 := []int{(v5 & v5), genFn1(""), v5, 42}
    if len(v65) > 0 {
        v65[0] = v5
    }
    fmt.Println(v2)
    var v66 int64 = -9223372036854775808
    v1[v5] = false
    v3 += (v3 * 929)
    v5 = (v5 * v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
}
