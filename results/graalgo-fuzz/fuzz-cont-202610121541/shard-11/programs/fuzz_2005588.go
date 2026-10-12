package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 int8, p1 int32) bool {
    fmt.Println(p1)
    v5 := GenS1{F0: 3, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v5
    return true
}
func genFn1(p0 bool) uint {
    p0 = ((272 == 2) || false)
    fmt.Println(p0)
    _ = p0
    return 535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []string{"a", "foo bar", "a"}
    var v4 int = 1
    v4 = (^(v4))
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        if (168 < -527) {
            var v7 bool = false
            _ = v7
        } else {
            var v8 int64 = 287
            v8 = 154
            v8 = (-(v8))
            _ = v8
        }
        if (!(4294967295 != 74)) {
            v4 = 1
            v9 := GenS1{F0: 1, F1: 2147483647}
            v10 := GenS1{F0: 42, F1: 1}
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    fmt.Println(v4)
    v4 = (v4 % ((v4) | 1))
    switch (v4 >> ((65535) & 63)) {
    case 42:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: 42, F1: 1}
            if ("\n\t" <= ((("hello" + ("\n\t" + "foo bar")) + "foo bar") + "foo bar")) {
                var v13 float32 = -2.25
                _ = v13
            } else {
                var v14 int8 = -56
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    var v15 uint8 = 45
    v16 := GenS1{F0: 502, F1: 1}
    if ((891 > 2) || true) {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint8 = (v15 % ((2) | 1))
            _ = fi17
            _ = v18
        }
        v4 = (v4 - v4)
        fmt.Println(v15)
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            _ = len(v1)
            var v21 uint8 = v15
            var v22 uint16 = 928
            _ = fi20
            _ = v21
            _ = v22
        }
        switch "hello" {
        case "\n\t":
            v4 -= 2147483647
        default:
            for fi23 := 0; fi23 < 18; fi23++ {
                if false {
                    switch 9223372036854775807 {
                    case 2147483647:
                        if true {
                            var v24 int32 = 97
                            v15 -= v15
                            if (-9223372036854775808 > 42) {
                                switch "foo bar" {
                                case "a":
                                    v25 := GenS1{F0: 4294967295, F1: 2147483648}
                                    _ = v25
                                case "hello":
                                    v4 += 83
                                case "":
                                    v3 = append(v3, (("foo bar" + "hello") + "a"))
                                }
                                if (false && (42 != 575)) {
                                    v26 := GenS1{F0: 65535, F1: 2}
                                    _ = v26
                                } else {
                                    switch (-(fi19)) {
                                    case 9223372036854775807:
                                        v15 += (v15 | v15)
                                    default:
                                        if true {
                                            var v27 int8 = 100
                                            v24 = (v24 / ((v24) | 1))
                                            var v28 uint32 = 1
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(v4, v4)
                                        }
                                    }
                                    var v29 uint = 255
                                    _ = v29
                                }
                                v30 := GenS2{F0: 0}
                                _ = v30
                            } else {
                                fmt.Println(v4, v15, fi19)
                            }
                            var v31 int8 = 1
                            v32 := GenS3{F0: 32767}
                            _ = v24
                            _ = v31
                            _ = v32
                        }
                    case 127:
                        v33 := GenS1{F0: 964, F1: 4294967295}
                        _ = v33
                    default:
                        switch ("hello" + "") {
                        case "a":
                            switch fi23 {
                            case 42:
                                v4 = (fi23 * (-(v4)))
                            default:
                                if (0.0 > (0.0 + (0.1 - 0.1))) {
                                    v4 = -9223372036854775808
                                    v34 := GenS4{F0: -1, F1: 1, F2: 0}
                                    v4 = fi19
                                    _ = v34
                                } else {
                                    if true {
                                        var v35 float64 = 0.0
                                        switch (fi19 * int(v15)) {
                                        case 555:
                                            var v36 int16 = 1
                                            _ = v36
                                        case 32767:
                                            var v37 float64 = 0.0
                                            _ = v37
                                        }
                                        var v38 int16 = 42
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v35
                                        _ = v38
                                    } else {
                                        var v39 string = "\n\t"
                                        _ = v39
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        case "hello":
                            v2[2147483648] = v4
                        case "foo bar":
                            v4 = 128
                        }
                    }
                    if false {
                        _ = len(v2)
                        switch fi19 {
                        case 128:
                            var v40 float32 = (1.5 * 0.1)
                            _ = v40
                        case 4294967295:
                            if (42 == 2) {
                                v41 := v1[4294967295]
                                v41 = (("" + v41) + (v41 + v41))
                                _ = v41
                            }
                        }
                        var v42 int16 = 1
                        v2[501] = 32768
                        _ = v42
                    }
                } else {
                    var v43 int32 = -2147483648
                    _ = v43
                }
                switch ("" + "a") {
                case "a":
                    v4 = (fi23 - 757)
                    if true {
                        fmt.Println(fi23, v4)
                        var v44 uint64 = 1
                        _ = v44
                    } else {
                        switch (("foo bar" + "hello") + ("foo bar" + ("\n\t" + ("\n\t" + ((("a" + (("\n\t" + ("" + "hello")) + ("\n\t" + "foo bar"))) + "") + ("hello" + ("\n\t" + ("foo bar" + "")))))))) {
                        case "a":
                            v15 += v15
                        }
                    }
                    fmt.Println(fi19)
                }
                v4 -= -1
                _ = fi23
            }
        }
        if (42 > 32767) {
            v1[547] = "a"
            switch (fi19 >> ((4294967295) & 63)) {
            case 4294967295:
                v4 -= 42
            default:
                fmt.Println(v15, v15, fi19)
            }
            var v45 int64 = 1
            _ = v45
        } else {
            if (!false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v46 := GenS1{F0: 2, F1: 65535}
                v15 = 42
                _ = v46
            }
            var v47 uint32 = 2147483648
            if true {
                var v48 uint8 = (v15 ^ v15)
                if len(v3) > 0 {
                    v3[0] = ("a" + "\n\t")
                }
                v49 := GenS5{F0: 42, F1: v48, F2: "foo bar"}
                _ = v48
                _ = v49
            }
            _ = v47
        }
        _ = fi19
    }
    var v50 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
