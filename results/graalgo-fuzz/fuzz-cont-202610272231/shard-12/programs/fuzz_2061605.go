package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int16; F1 int32; F2 int32 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int; F2 uint32 }
type GenS8 struct { F0 string; F1 int32 }
func genFn0(p0 int16, p1 bool) uint32 {
    p1 = p1
    switch 9223372036854775807 {
    case 42:
        for fi5 := 0; fi5 < 14; fi5++ {
            fmt.Println(p1, fi5, p0)
            _ = fi5
        }
    case 128:
        var v6 float64 = (3.14159 / (1.5 + 1e10))
        _ = v6
    default:
        p0 = 32767
    }
    _ = p0
    _ = p1
    return 367
}
func genFn1(p0 bool) uint {
    v7 := GenS1{F0: -863, F1: 301, F2: 2}
    v8 := GenS1{F0: -1, F1: 42, F2: 42}
    if false {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS2{F0: true, F1: (true && p0), F2: 9223372036854775807}
            var v11 uint16 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
        if (575 >= 18446744073709551615) {
            v12 := GenS2{F0: (42 < 127), F1: p0, F2: 42}
            fmt.Println(p0)
            switch 128 {
            case -167:
                var v13 uint64 = 767
                _ = v13
            case 9223372036854775807:
                v14 := []int{-121, 882, 42}
                _ = v14
            case 42:
                fmt.Println(p0)
            }
            _ = v12
        }
        if (p0 || (128 < 41)) {
            v15 := []int{128}
            p0 = true
            _ = v15
        }
    } else {
        v16 := []int{9223372036854775807, 9223372036854775807, -1}
        p0 = true
        p0 = (4294967295 != 0)
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS3{F0: 32767, F1: (p0 || (2147483648 != 42))}
            p0 = ((-47 > -1) && (1e10 > (0.1 - float32(float64(-2.25)))))
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    _ = p0
    _ = v7
    _ = v8
    return 812
}

func main() {
    var v1 uint64 = 583
    var v2 int32 = 0
    var v3 uint64 = 1
    v4 := []int{1, -967}
    var v20 uint = 486
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 := GenS4{F0: -515, F1: 1, F2: v2}
    v1 -= (v3 << ((v20) & 63))
    v20 += (v20 + 1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch -1 {
    case 575:
        if (((0.0 - 0.1) + 3.14159) == 1e10) {
            v22 := GenS2{F0: true, F1: false, F2: 42}
            _ = v22
        } else {
            if false {
                if false {
                    var v23 uint = v20
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v23
                }
            } else {
                if (-1 > 127) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v24 int16 = 42
                    _ = v24
                } else {
                    _ = len(v4)
                    v4 = append(v4, 32768)
                    if len(v4) > 0 {
                        v4[0] = 128
                    }
                }
            }
        }
    case 32767:
        var v25 uint = v20
        _ = v25
    case -2147483648:
        for fi26 := 0; fi26 < 6; fi26++ {
            if len(v4) > 0 {
                v4[0] = fi26
            }
            v3 = v3
            if false {
                v2 = 2147483647
                if true {
                    v4 = append(v4, int(v1))
                    var v27 bool = (true && false)
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS4{F0: -32768, F1: (^(v2)), F2: v2}
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                } else {
                    var v30 uint8 = 210
                    var v31 int = fi26
                    switch fi26 {
                    case 2147483647:
                        for fi32 := 0; fi32 < 2; fi32++ {
                            v33 := GenS5{F0: 272, F1: v30}
                            v31 -= 747
                            var v34 float32 = 0.0
                            if (!true) {
                                if true {
                                    v20 = (v20 * 255)
                                    var v35 uint = (v20 * (v20 / ((v20) | 1)))
                                    switch "hello" {
                                    case "":
                                        switch (("" + "a") + "foo bar") {
                                        case "foo bar":
                                            v1 = (v3 - 1)
                                        }
                                    case "\n\t":
                                        if (true || true) {
                                            if len(v4) > 0 {
                                                v4[0] = (fi32 ^ (v31 - (fi26 * (-(fi32)))))
                                            }
                                        }
                                    default:
                                        v31 = 32767
                                    }
                                    _ = len(v4)
                                    _ = v35
                                }
                                var v36 string = ("hello" + "foo bar")
                                v20 -= v20
                                var v37 uint64 = 791
                                _ = v36
                                _ = v37
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (fi26 * (v31 - int(v30)))
                                }
                                switch ("\n\t" + (((("hello" + "") + "hello") + "\n\t") + "\n\t")) {
                                case "hello":
                                    var v38 uint16 = 644
                                    _ = v38
                                }
                                var v39 float64 = 1e10
                                _ = v39
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    case 32767:
                        v40 := GenS4{F0: 42, F1: (v2 ^ -735), F2: v2}
                        _ = v40
                    }
                    _ = v30
                    _ = v31
                }
            } else {
                fmt.Println(v20)
            }
            _ = fi26
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v41 float64 = ((float64(float32(1e10)) * 0.1) / 3.14159)
    v2 = (v2 - 0)
    var v42 bool = true
    for fi43 := 0; fi43 < 2; fi43++ {
        if (("\n\t" + "foo bar") != "foo bar") {
            switch (-(fi43)) {
            case -791:
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
            case 844:
                v44 := GenS6{F0: "foo bar"}
                _ = v44
            case -898:
                for fi45 := 0; fi45 < 17; fi45++ {
                    v46 := GenS6{F0: "foo bar"}
                    _ = len(v4)
                    var v47 uint16 = 42
                    var v48 int8 = 127
                    var v49 uint = 65535
                    var v50 int16 = -863
                    var v51 int16 = v50
                    fmt.Println(fi43, v48, v51)
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                    _ = v49
                    _ = v50
                    _ = v51
                }
            }
            if len(v4) > 0 {
                v4[0] = (fi43 | 127)
            }
        }
        _ = fi43
    }
    v3 = (v3 >> (((v20 * 65)) & 63))
    var v52 float64 = (((((v41 / ((float64(((((-2.25 * float32(float64(-2.25))) - 3.14159) * 1e10) + 1.5)) / -2.25) - (0.0 - 1e10))) * (v41 - v41)) / -2.25) * 0.1) / ((3.14159 / 3.14159) + 0.1))
    v53 := GenS6{F0: "hello"}
    if (1 == genFn0(42, ("a" == "a"))) {
        v3 = v3
        var v54 int8 = 24
        if (0 != v20) {
            if v42 {
                v3 = v3
                if (true && (57 < 144)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v42 = v42
                    var v55 int = 0
                    _ = v55
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        _ = v54
    } else {
        switch 0 {
        case -2147483648:
            v56 := GenS7{F0: 42, F1: -827, F2: 42}
            _ = v56
        }
        fmt.Println(v20, v1)
        var v57 uint32 = 2147483648
        _ = v57
    }
    var v58 int = -9223372036854775808
    v59 := GenS8{F0: ("\n\t" + (("foo bar" + "a") + "a")), F1: (v2 - (v2 % (((v2 * -1)) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v52
    _ = v53
    _ = v58
    _ = v59
}
