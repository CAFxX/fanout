package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 int32) bool {
    v6 := []int{}
    var v7 uint8 = 136
    fmt.Println(v7)
    _ = p0
    _ = v6
    _ = v7
    return (!true)
}
func genFn1() uint {
    var v8 float32 = (((3.14159 / -2.25) * ((0.0 / (1.5 + 3.14159)) * (3.14159 - 1.5))) - -2.25)
    switch ("a" + ((("foo bar" + "foo bar") + ((("a" + "a") + "a") + ("hello" + "hello"))) + "a")) {
    case "":
        var v9 int8 = -128
        _ = v9
    case "a":
        v8 -= 1.5
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{}
            var v13 int = (-(fi11))
            var v14 uint8 = 181
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    var v15 uint8 = 1
    _ = v15
    _ = v8
    return 187
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = (128 != 475)
    var v3 float64 = 0.1
    var v4 bool = (!v2)
    var v5 string = "\n\t"
    v16 := []int{}
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int16 = 924
        _ = fi17
        _ = v18
    }
    var v19 int64 = -568
    if (!v2) {
        var v20 int16 = 1
        _ = v20
    } else {
        switch ((("\n\t" + v5) + ((v5 + (("a" + "\n\t") + ((v5 + v5) + "\n\t"))) + (((((("a" + "\n\t") + "") + v5) + v5) + v5) + "foo bar"))) + (v5 + v5)) {
        case "foo bar":
            if (32767 == -9223372036854775808) {
                if (!(-9223372036854775808 <= 983)) {
                    v5 = ""
                    var v21 float32 = 1.5
                    if v4 {
                        switch 535 {
                        case 136:
                            v22 := GenS1{F0: 1, F1: 1}
                            _ = v22
                        case 6:
                            v3 = 0.1
                        }
                    } else {
                        v1 = (v1 + (v1 << ((347) & 63)))
                    }
                    _ = v21
                } else {
                    switch v5 {
                    case "\n\t":
                        if (!(478 > v1)) {
                            fmt.Println(v3)
                            for fi23 := 0; fi23 < 2; fi23++ {
                                if v2 {
                                    v4 = v4
                                }
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    v3 += (((((v3 * v3) - v3) * 1.5) - v3) * v3)
                                    v25 := GenS2{F0: 32767, F1: 2}
                                    var v26 int16 = 32767
                                    if (false || (1 > 872)) {
                                        if false {
                                            fmt.Println(v26, fi23, fi23)
                                            v26 = -1
                                            v27 := GenS1{F0: 4294967295, F1: 255}
                                            _ = v27
                                        }
                                    } else {
                                        if (!false) {
                                            fmt.Println(v19, v5, v2)
                                        }
                                        switch v5 {
                                        case "foo bar":
                                            if (((!v2) || (true && ((v1 << ((4294967295) & 63)) <= v1))) && false) {
                                                v28 := GenS2{F0: v26, F1: 18446744073709551615}
                                                fmt.Println(fi23, v19, v26)
                                                _ = v28
                                            } else {
                                                var v29 int16 = (v26 * (-(v26)))
                                                v5 = ("\n\t" + v5)
                                                _ = v29
                                            }
                                        case "":
                                            v19 = (^(v19))
                                        default:
                                            var v30 string = "\n\t"
                                            _ = v30
                                        }
                                        var v31 bool = ((v1 == (v1 ^ 2147483647)) || v4)
                                        _ = v31
                                    }
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi23
                            }
                            v32 := GenS3{F0: 67, F1: uint64(v1), F2: 32767}
                            _ = v32
                        } else {
                            if v4 {
                                var v33 int16 = 32767
                                v34 := GenS1{F0: 2, F1: 127}
                                _ = v33
                                _ = v34
                            } else {
                                for fi35 := 0; fi35 < 10; fi35++ {
                                    v1 += 2147483647
                                    if len(v16) > 0 {
                                        v16[0] = (fi35 | 1)
                                    }
                                    _ = fi35
                                }
                                v36 := GenS3{F0: 1, F1: 42, F2: 2}
                                _ = v36
                            }
                            var v37 uint64 = 646
                            _ = v37
                        }
                    }
                    var v38 float64 = (3.14159 + float64((3.14159 + 1e10)))
                    _ = v38
                }
            } else {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                v39 := GenS2{F0: 0, F1: uint(v1)}
                _ = v39
            }
        case "hello":
            v19 += 9223372036854775807
        default:
            for fi40 := 0; fi40 < 7; fi40++ {
                var v41 bool = false
                _ = fi40
                _ = v41
            }
        }
    }
    v42 := GenS2{F0: int16(v19), F1: 18446744073709551615}
    fmt.Println(v4, v4)
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if true {
        v19 = 42
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    } else {
        v5 = v5
        switch v5 {
        case "\n\t":
            switch int(v1) {
            case 127:
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 uint32 = 2147483648
                    var v45 uint8 = 113
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            default:
                _ = len(v16)
            }
        case "hello":
            v46 := GenS4{F0: 92, F1: 1000}
            _ = v46
        case "a":
            for fi47 := 0; fi47 < 12; fi47++ {
                for fi48 := 0; fi48 < 16; fi48++ {
                    var v49 float32 = (float32(v3) + -2.25)
                    _ = fi48
                    _ = v49
                }
                var v50 string = v5
                _ = fi47
                _ = v50
            }
        }
    }
    _ = len(v16)
    var v51 string = v5
    fmt.Println(v2)
    v52 := GenS2{F0: 902, F1: 65535}
    for fi53 := 0; fi53 < 6; fi53++ {
        v2 = (v2 && (65535 <= 693))
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v51
    _ = v52
}
