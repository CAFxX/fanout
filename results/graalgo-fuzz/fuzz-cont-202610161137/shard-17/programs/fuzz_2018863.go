package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8 }
func genFn0() int16 {
    if ((float32((0.1 - 1e10)) >= (-2.25 * (1.5 * 3.14159))) || (0 <= 1)) {
        for fi6 := 0; fi6 < 11; fi6++ {
            if (!true) {
                fmt.Println(fi6)
            } else {
                switch (("foo bar" + ("\n\t" + ((("" + "") + "hello") + ("a" + "foo bar")))) + "a") {
                case "":
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi6)
                    }
                default:
                    var v7 uint8 = 240
                    _ = v7
                }
            }
            _ = fi6
        }
    }
    v8 := []int{4294967295}
    var v9 uint32 = 2147483647
    if len(v8) > 0 {
        v8[0] = 128
    }
    var v10 int = 1
    _ = v10
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 uint32) string {
    p0 = (p0 % ((p0) | 1))
    p0 += (p0 | (p0 * (p0 << ((45) & 63))))
    p0 -= (p0 & p0)
    fmt.Println(p0)
    _ = p0
    return (("foo bar" + "") + "foo bar")
}

func main() {
    var v1 float32 = 3.14159
    var v2 float64 = -2.25
    var v3 bool = true
    var v4 uint16 = 255
    v5 := []int{}
    if true {
        var v11 uint = 416
        v11 = (v11 >> ((v11) & 63))
        v12 := GenS1{F0: (v11 | v11), F1: (("hello" + "a") + "a")}
        _ = v11
        _ = v12
    } else {
        _ = len(v5)
    }
    v13 := GenS1{F0: 42, F1: ("a" + ("foo bar" + (("a" + ("a" + ("hello" + ("hello" + "foo bar")))) + "foo bar")))}
    v14 := GenS2{F0: -1}
    switch 216 {
    case 9223372036854775807:
        v15 := GenS2{F0: 784}
        _ = v15
    case 483:
        switch -389 {
        case -9223372036854775808:
            switch "foo bar" {
            case "hello":
                v2 = float64((((v1 + ((v1 + (0.1 + ((v1 - (v1 * ((0.1 * (v1 / v1)) - 1e10))) * v1))) + v1)) + -2.25) / (float32(float64(((0.0 / v1) - (3.14159 + 0.0)))) / float32(float64(float32(((0.0 / float64(v1)) + 0.1)))))))
            case "":
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint = uint(v4)
                    v3 = v3
                    v18 := GenS2{F0: -2147483648}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        case 32768:
            v19 := GenS2{F0: 124}
            _ = v19
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case 1:
        var v20 uint32 = 42
        _ = v20
    default:
        var v21 uint64 = 4294967295
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        switch "\n\t" {
        case "hello":
            for fi23 := 0; fi23 < 9; fi23++ {
                _ = len(v5)
                _ = fi23
            }
        case "foo bar":
            v1 = (v1 * 1.5)
        default:
            for fi24 := 0; fi24 < 4; fi24++ {
                fmt.Println(v4)
                v25 := GenS3{F0: v3, F1: 0, F2: -32768}
                _ = fi24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            v3 = v3
            if true {
                v27 := GenS1{F0: 18446744073709551615, F1: ("hello" + "\n\t")}
                _ = v27
            } else {
                if len(v5) > 0 {
                    v5[0] = (fi22 % (((fi26 / ((fi26) | 1))) | 1))
                }
                var v28 float32 = v1
                _ = v28
            }
            _ = fi26
        }
        _ = fi22
    }
    switch 127 {
    case -1:
        switch 548 {
        case 2147483647:
            for fi29 := 0; fi29 < 10; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (-1 <= 0) {
                        v4 = 37
                        v2 -= 0.1
                        switch ((("\n\t" + "a") + ("" + "\n\t")) + "\n\t") {
                        case "":
                            if (59 <= 199) {
                                if (false || v3) {
                                    _ = len(v5)
                                    v31 := GenS2{F0: 361}
                                    var v32 int16 = genFn0()
                                    _ = v31
                                    _ = v32
                                } else {
                                    switch -386 {
                                    case 4294967295:
                                        var v33 uint32 = 611
                                        _ = len(v5)
                                        _ = v33
                                    case -866:
                                        switch -793 {
                                        case 2147483647:
                                            switch (fi29 + (^(fi30))) {
                                            case 32767:
                                                var v34 int = fi30
                                                _ = v34
                                            case 127:
                                                if (v2 < v2) {
                                                    var v35 int32 = 777
                                                    v36 := GenS3{F0: (v1 == float32(v2)), F1: 812, F2: -1}
                                                    _ = v35
                                                    _ = v36
                                                }
                                                if false {
                                                    v3 = (false && v3)
                                                } else {
                                                    var v37 int32 = int32(fi29)
                                                    _ = v37
                                                }
                                                if v3 {
                                                    var v38 int16 = -32768
                                                    _ = len(v5)
                                                    _ = v38
                                                } else {
                                                    v39 := GenS4{F0: v4}
                                                    v40 := GenS2{F0: 640}
                                                    v41 := GenS2{F0: (fi29 << ((0) & 63))}
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                            default:
                                                v42 := GenS5{F0: 4294967295, F1: v3}
                                                _ = v42
                                            }
                                        default:
                                            var v43 string = "foo bar"
                                            _ = v43
                                        }
                                    }
                                }
                                v1 = v1
                                fmt.Println(v3)
                            }
                            var v44 uint32 = 909
                            v45 := GenS1{F0: 42, F1: ""}
                            _ = v44
                            _ = v45
                        }
                    } else {
                        var v46 float64 = v2
                        if ("a" == "foo bar") {
                            if ("\n\t" <= "\n\t") {
                                v46 = 0.0
                            } else {
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    v47 := GenS2{F0: fi29}
                                    _ = v47
                                case "foo bar":
                                    v1 -= (float32(((0.1 + v2) * 1.5)) * (0.1 / v1))
                                default:
                                    v48 := GenS4{F0: v4}
                                    _ = v48
                                }
                            }
                            v49 := GenS6{F0: 42, F1: "a"}
                            _ = v49
                        } else {
                            v50 := GenS7{F0: false}
                            if false {
                                var v51 int8 = 100
                                var v52 float32 = v1
                                v52 += v52
                                _ = v51
                                _ = v52
                            }
                            _ = v50
                        }
                        _ = v46
                    }
                    v4 += 255
                    _ = fi30
                }
                _ = fi29
            }
        case -9223372036854775808:
            v53 := GenS7{F0: ((!(!(1 != 4294967295))) && (!v3))}
            _ = v53
        }
    case 128:
        v3 = ((false || (42 > 42)) || ((32768 <= 128) && (!(((v2 * float64(1.5)) * float64(v1)) == v2))))
    default:
        for fi54 := 0; fi54 < 9; fi54++ {
            var v55 int16 = -1
            fmt.Println(v1, v3, v2)
            v56 := GenS6{F0: 2, F1: ("foo bar" + ((("\n\t" + "hello") + ("\n\t" + ("\n\t" + ("\n\t" + "foo bar")))) + "hello"))}
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    var v57 int16 = genFn0()
    v58 := GenS3{F0: (!(v2 <= 1e10)), F1: 1, F2: v57}
    if (0 == 42) {
        switch "" {
        case "foo bar":
            var v59 float32 = v1
            _ = v59
        default:
            v60 := GenS8{F0: 42}
            _ = v60
        }
        fmt.Println(v4)
    }
    fmt.Println(v2, v57)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
