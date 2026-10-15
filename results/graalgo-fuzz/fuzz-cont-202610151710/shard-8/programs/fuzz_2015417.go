package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 uint16) uint32 {
    var v6 bool = true
    var v7 int32 = -455
    var v8 float64 = 1e10
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1() uint16 {
    switch ("hello" + ("\n\t" + "hello")) {
    case "a":
        var v9 uint32 = 65535
        _ = v9
    case "\n\t":
        var v10 int = 620
        _ = v10
    case "hello":
        if false {
            var v11 int64 = -611
            if true {
                var v12 float32 = 0.1
                if (((870 != 32767) || (42 >= -128)) || false) {
                    switch ((("\n\t" + "a") + "") + "\n\t") {
                    case "\n\t":
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := GenS1{F0: true}
                            var v15 uint16 = 65535
                            var v16 uint = 65535
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                }
                _ = v12
            }
            if (false || ("" < "a")) {
                switch ("foo bar" + "") {
                case "a":
                    for fi17 := 0; fi17 < 11; fi17++ {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if (false || (!(106 <= 101))) {
                                var v19 int32 = 734
                                fmt.Println(v19)
                                _ = v19
                            } else {
                                v20 := []int{fi17, (fi17 ^ -9223372036854775808), (-(fi18)), (fi17 & 1)}
                                _ = v20
                            }
                            _ = fi18
                        }
                        v21 := GenS1{F0: true}
                        _ = fi17
                        _ = v21
                    }
                default:
                    v11 -= (v11 & v11)
                }
                v11 -= v11
            }
            _ = v11
        } else {
            v22 := GenS2{F0: 878, F1: -113, F2: -32768}
            var v23 int16 = 183
            _ = v22
            _ = v23
        }
    }
    v24 := []int{}
    _ = v24
    return 32767
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    v2 := []string{"", "hello"}
    var v3 int8 = -128
    var v4 uint16 = 255
    v5 := []uint32{128, 1, 2147483647, 4294967295}
    if (0.0 != float32(0.1)) {
        fmt.Println(v3)
        var v25 uint64 = 700
        var v26 float32 = (1.5 + (0.0 * 1e10))
        _ = v25
        _ = v26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = append(v5, 682)
    v4 = 1
    var v27 bool = (!false)
    if len(v5) > 0 {
        v5[0] = genFn0(v4)
    }
    switch ("a" + "hello") {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "a"
        }
    case "a":
        fmt.Println(v27, v3, v3)
    case "":
        v3 += (v3 - v3)
    default:
        switch -594 {
        case 42:
            if v27 {
                v28 := GenS3{F0: (!((^(v4)) > (v4 | (v4 % (((v4 & v4)) | 1)))))}
                fmt.Println(v27, v4, v3)
                fmt.Println(v3, v4, v3)
                _ = v28
            } else {
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int = 2147483647
                    _ = fi29
                    _ = v30
                }
            }
        case 9223372036854775807:
            if (v27 || true) {
                v27 = (42 >= -492)
            }
        default:
            switch "" {
            case "\n\t":
                v31 := GenS3{F0: true}
                _ = v31
            case "a":
                for fi32 := 0; fi32 < 1; fi32++ {
                    var v33 int8 = (v3 ^ (v3 << ((201) & 63)))
                    _ = fi32
                    _ = v33
                }
            case "hello":
                switch "" {
                case "foo bar":
                    var v34 int8 = v3
                    _ = v34
                case "":
                    v35 := GenS3{F0: v27}
                    _ = v35
                case "hello":
                    v4 += genFn1()
                default:
                    var v36 string = ("foo bar" + "\n\t")
                    _ = v36
                }
            default:
                v27 = v27
            }
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v37 int = 656
    v38 := GenS4{F0: v37, F1: 2147483648}
    v37 = (v37 - 128)
    var v39 int8 = v3
    v40 := GenS4{F0: v37, F1: genFn0(0)}
    var v41 float64 = 0.0
    v27 = false
    fmt.Println(v41, v4, v4)
    var v42 uint32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
