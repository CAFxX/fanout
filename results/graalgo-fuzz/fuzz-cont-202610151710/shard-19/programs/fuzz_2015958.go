package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint }
func genFn0(p0 uint32) int64 {
    var v5 uint16 = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = 18446744073709551615
        var v8 int64 = 0
        if true {
            v9 := GenS1{F0: 2}
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v10 int = 1
    v11 := []int{32768, v10, 4294967295, v10}
    fmt.Println(v5)
    _ = p0
    _ = v10
    _ = v11
    _ = v5
    return -1
}
func genFn1() uint8 {
    var v12 float32 = ((((float32(-2.25) * 0.1) * float32(1e10)) - 1e10) * -2.25)
    var v13 int32 = -2147483648
    if (false || false) {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint32 = 42
            v12 = float32(0.1)
            if (!(!true)) {
                switch ("" + ((("\n\t" + "hello") + "foo bar") + "hello")) {
                case "a":
                    if false {
                        v16 := []int{fi14, (fi14 / ((fi14) | 1))}
                        fmt.Println(v12, fi14, v13)
                        _ = v16
                    }
                case "\n\t":
                    v17 := []int{(fi14 ^ int(v15)), -627, (fi14 << ((2) & 63)), 244}
                    _ = v17
                case "hello":
                    v18 := GenS2{F0: (((("foo bar" + "\n\t") + "hello") + "\n\t") + "\n\t")}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v20 := GenS1{F0: 0}
                        var v21 uint64 = 0
                        v12 += v12
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        switch 1 {
        case -916:
            var v22 float64 = 3.14159
            _ = v22
        default:
            switch ("a" + (("hello" + "foo bar") + "a")) {
            case "a":
                if (0 > 1) {
                    switch -1 {
                    case -2147483648:
                        v23 := GenS1{F0: 2}
                        _ = v23
                    case -1:
                        var v24 float32 = float32(3.14159)
                        _ = v24
                    case 440:
                        for fi25 := 0; fi25 < 20; fi25++ {
                            if ((!(1.5 != 0.1)) && (50 == 9223372036854775807)) {
                                v26 := GenS3{F0: 127}
                                var v27 uint32 = 65535
                                _ = v26
                                _ = v27
                            } else {
                                fmt.Println(v12, v12)
                            }
                            _ = fi25
                        }
                    default:
                        switch "" {
                        case "\n\t":
                            for fi28 := 0; fi28 < 3; fi28++ {
                                v12 += 0.0
                                var v29 uint16 = 42
                                if ((1e10 * 3.14159) > 3.14159) {
                                    v13 += -466
                                    switch ((("hello" + "a") + "hello") + (("" + (("foo bar" + ("hello" + (("hello" + "\n\t") + ("a" + "")))) + "\n\t")) + (("hello" + ("a" + "a")) + "\n\t"))) {
                                    case "a":
                                        v30 := GenS4{F0: 361, F1: (-(v13)), F2: 679}
                                        _ = v30
                                    case "hello":
                                        if true {
                                            var v31 uint = 1
                                            _ = v31
                                        }
                                    }
                                    v13 = (v13 << ((550) & 63))
                                }
                                _ = fi28
                                _ = v29
                            }
                        case "foo bar":
                            var v32 float32 = (((((v12 - -2.25) / v12) / -2.25) - float32((3.14159 * -2.25))) * (v12 / (v12 + v12)))
                            _ = v32
                        default:
                            v33 := []int{2147483647}
                            _ = v33
                        }
                    }
                } else {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        if (-32768 > 342) {
                            v35 := GenS1{F0: 2}
                            _ = v35
                        } else {
                            fmt.Println(v12, fi34, fi34)
                        }
                        _ = fi34
                    }
                    v36 := []int{9223372036854775807, 42, -808, -745}
                    _ = v36
                }
            case "hello":
                v37 := GenS1{F0: 0}
                _ = v37
            case "foo bar":
                v12 += 1.5
            default:
                v12 = 1.5
            }
        }
        v12 = 3.14159
    }
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 int32 = -379
    var v3 int16 = -1
    var v4 int64 = 42
    if false {
        v1 -= 775
    } else {
        v38 := []int{42}
        _ = v38
    }
    v39 := GenS2{F0: ""}
    v3 += 1000
    v40 := []int{127, 127, -9223372036854775808, 363}
    v41 := GenS4{F0: 922, F1: (v2 >> ((478) & 63)), F2: 2}
    var v42 float64 = (((1e10 * 1e10) / 1e10) + 0.1)
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    v3 = (v1 % ((v3) | 1))
    var v43 float32 = 1.5
    if ((v4 ^ (v4 >> ((145) & 63))) < (v4 ^ (v4 & (v4 + 428)))) {
        if (((!true) && true) && (2 < 1)) {
            v44 := GenS5{F0: 65535, F1: int16(v2)}
            _ = v44
        } else {
            var v45 uint64 = 53
            _ = v45
        }
        for fi46 := 0; fi46 < 1; fi46++ {
            fmt.Println(v2, v1)
            if (fi46 == fi46) {
                v43 = (3.14159 - v43)
            }
            if len(v40) > 0 {
                v40[0] = (-(fi46))
            }
            _ = fi46
        }
        if (0 <= genFn1()) {
            fmt.Println(v4, v4, v43)
            v47 := GenS3{F0: 42}
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = v47
        } else {
            var v48 float64 = (float64(v43) - 3.14159)
            var v49 int8 = 1
            for fi50 := 0; fi50 < 10; fi50++ {
                v2 -= (v2 * (v2 << ((883) & 63)))
                _ = fi50
            }
            _ = v48
            _ = v49
        }
    } else {
        for fi51 := 0; fi51 < 11; fi51++ {
            var v52 int = (fi51 | -9223372036854775808)
            if true {
                var v53 uint16 = 255
                var v54 uint8 = 1
                switch (-(fi51)) {
                case 32767:
                    v42 = -2.25
                case -9223372036854775808:
                    fmt.Println(v3, v2, v43)
                case 582:
                    for fi55 := 0; fi55 < 6; fi55++ {
                        fmt.Println(fi55, fi55)
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                        if (v2 >= v2) {
                            v54 = 128
                            var v56 uint = 65535
                            v52 += fi55
                            _ = v56
                        }
                        _ = fi55
                    }
                default:
                    v53 = v53
                }
                _ = v53
                _ = v54
            }
            _ = fi51
            _ = v52
        }
    }
    if len(v40) > 0 {
        v40[0] = 2147483647
    }
    fmt.Println(v43, v1, v4)
    v43 = ((v43 * v43) - (v43 - ((v43 * ((v43 / v43) / (float32(float64(1.5)) / (v43 / (v43 + 3.14159))))) / 1.5)))
    v57 := GenS6{F0: 1, F1: 634}
    fmt.Println(v3, v43)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v57
}
