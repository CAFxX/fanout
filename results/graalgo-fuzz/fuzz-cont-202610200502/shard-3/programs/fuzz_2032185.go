package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 int32) int32 {
    if true {
        var v5 uint16 = 255
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint64 = 1
                if (v5 > (v5 / ((32767) | 1))) {
                    var v9 uint = 255
                    var v10 int = (fi7 / ((fi7) | 1))
                    _ = v10
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            fmt.Println(p0, v5)
            _ = fi6
        }
        var v11 string = "hello"
        _ = v11
        _ = v5
    }
    v12 := []int{32768}
    _ = len(v12)
    p0 -= 0
    _ = p0
    _ = v12
    return (p0 & (p0 ^ p0))
}
func genFn1() uint {
    switch 32767 {
    case 528:
        if (-2147483648 <= 643) {
            v13 := []int{-864}
            if ((3.14159 + 1e10) >= 3.14159) {
                v13 = append(v13, 1)
                if true {
                    v14 := GenS1{F0: 2, F1: 22}
                    if (2 > 0) {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            var v16 string = ("hello" + (("foo bar" + (("" + "") + "\n\t")) + ""))
                            for fi17 := 0; fi17 < 12; fi17++ {
                                v13 = append(v13, (fi15 & 1))
                                v16 = "hello"
                                v16 = v16
                                _ = fi17
                            }
                            for fi18 := 0; fi18 < 4; fi18++ {
                                v16 = "hello"
                                var v19 uint8 = 1
                                v20 := GenS1{F0: 42, F1: 27}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        var v21 uint8 = 2
                        var v22 float32 = (((((0.0 * (3.14159 + 0.1)) - 0.1) / 1.5) / float32(1.5)) / 3.14159)
                        _ = v21
                        _ = v22
                    }
                    var v23 int32 = -1
                    _ = v14
                    _ = v23
                }
            } else {
                switch "" {
                case "a":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        fmt.Println(fi24)
                        for fi25 := 0; fi25 < 15; fi25++ {
                            if len(v13) > 0 {
                                v13[0] = fi24
                            }
                            _ = fi25
                        }
                        _ = fi24
                    }
                case "foo bar":
                    var v26 int16 = 42
                    _ = v26
                case "":
                    var v27 float32 = float32((((float64(1.5) / 1.0) - (1e10 + 0.0)) - ((float64(0.0) - (0.1 + -2.25)) * -2.25)))
                    _ = v27
                }
            }
            _ = v13
        }
    case 32768:
        var v28 bool = (false || (742 > 0))
        _ = v28
    case -9223372036854775808:
        var v29 int = 1
        _ = v29
    default:
        var v30 bool = false
        _ = v30
    }
    var v31 uint16 = 255
    _ = v31
    return 2
}

func main() {
    var v1 bool = (!(-829 <= 42))
    v2 := []string{"a"}
    var v3 uint = 798
    var v4 int16 = 32767
    v1 = (986 != 2147483647)
    var v32 uint = v3
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "foo bar")
    }
    var v33 uint8 = 0
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int = fi34
        v36 := GenS1{F0: (v32 - (v3 + genFn1())), F1: 114}
        _ = fi34
        _ = v35
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v37 uint8 = v33
    var v38 float64 = float64((((1e10 * ((0.0 * 1.5) * 1e10)) - (0.0 + 3.14159)) * 1.5))
    var v39 int = 127
    switch ("a" + "hello") {
    case "a":
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 int16 = v4
            if ("foo bar" <= "hello") {
                v42 := GenS2{F0: "a", F1: (v4 << (((^(v32))) & 63)), F2: 2}
                fmt.Println(v41, v4, v39)
                _ = v42
            }
            _ = fi40
            _ = v41
        }
    case "":
        var v43 float32 = (((0.0 - 0.1) + 3.14159) * 0.1)
        _ = v43
    case "foo bar":
        v4 -= int16(v39)
    }
    switch "foo bar" {
    case "foo bar":
        v44 := GenS1{F0: v32, F1: 31}
        _ = v44
    case "hello":
        v45 := GenS2{F0: ("a" + ""), F1: v4, F2: 65535}
        _ = v45
    default:
        if (1 != 1) {
            var v46 float32 = float32(((v38 * ((v38 * float64((0.1 / 1.0))) * (1e10 - (v38 * v38)))) / 3.14159))
            _ = v46
        }
    }
    v32 = genFn1()
    for fi47 := 0; fi47 < 2; fi47++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "" {
        case "a":
            for fi48 := 0; fi48 < 13; fi48++ {
                v49 := GenS1{F0: (v3 >> ((v32) & 63)), F1: 127}
                if false {
                    fmt.Println(v37)
                    fmt.Println(fi47)
                    var v50 uint8 = v33
                    _ = v50
                } else {
                    v39 = (v39 + (v39 | 127))
                }
                _ = fi48
                _ = v49
            }
        }
        v39 += fi47
        _ = fi47
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        v37 = v37
        v3 += 65535
        var v52 float32 = 1.5
        _ = fi51
        _ = v52
    }
    if v1 {
        if ((v37 << (((v3 / ((v32) | 1))) & 63)) == (v37 & 2)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
        }
        var v53 string = "foo bar"
        var v54 int16 = (v4 & v4)
        _ = v53
        _ = v54
    } else {
        v1 = v1
        var v55 uint64 = 4294967295
        if true {
            if (v1 || (0.0 <= 1.5)) {
                var v56 int8 = 42
                if ((v4 >> ((255) & 63)) >= 1000) {
                    _ = len(v2)
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
                _ = v56
            }
            v3 = (v32 + v3)
        } else {
            for fi57 := 0; fi57 < 19; fi57++ {
                v55 = (v55 % (((v55 - (^(v55)))) | 1))
                _ = fi57
            }
            var v58 bool = v1
            _ = v58
        }
        _ = v55
    }
    fmt.Println(v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
