package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint16) int64 {
    v5 := []int{-2147483648, 710, 0}
    if (2 < 1) {
        p0 = p0
        v6 := GenS1{F0: 39}
        _ = len(v5)
        _ = v6
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1() int8 {
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float64 = 0.0
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint = 4294967295
                v10 = (^(v10))
                switch (^(fi7)) {
                case 1:
                    v8 -= v8
                default:
                    v11 := []int{}
                    _ = v11
                }
                var v12 int = 32767
                var v13 int = -2147483648
                v10 = (^(v10))
                _ = fi9
                _ = v10
                _ = v12
                _ = v13
            }
            if (0.1 <= 0.1) {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := []int{fi7, fi7, fi7}
                    v16 := GenS2{F0: 42}
                    v8 = v8
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                if true {
                    var v17 int16 = 1
                    var v18 int64 = 0
                    v19 := GenS2{F0: 841}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v8 += v8
            } else {
                if (v8 < ((v8 * ((-2.25 * (((v8 / v8) * v8) - v8)) - (3.14159 - (1.5 + v8)))) - (v8 / 0.1))) {
                    v20 := []int{(fi7 - fi7), (fi7 ^ fi7), fi7, (fi7 - -9223372036854775808)}
                    if true {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        for fi21 := 0; fi21 < 9; fi21++ {
                            var v22 uint16 = 1
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = fi21
                            _ = v22
                        }
                        v8 += 0.1
                    } else {
                        _ = len(v20)
                    }
                    _ = v20
                }
            }
            _ = fi7
            _ = v8
        }
        var v23 float32 = (1e10 + (((float32(0.1) * 1e10) * 0.1) + 1e10))
        v23 += 0.0
        _ = v23
    }
    var v24 uint16 = 255
    _ = v24
    return -100
}

func main() {
    var v1 bool = (!(389 != 159))
    var v2 int64 = 1
    v3 := make(map[int]bool)
    var v4 int = 128
    v25 := GenS3{F0: -456, F1: -128, F2: 1}
    fmt.Println(v4, v1, v2)
    var v26 uint8 = 127
    switch ("" + ("a" + "")) {
    case "":
        if false {
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS1{F0: 0}
                switch ("hello" + "") {
                case "a":
                    if ((!(32767 > 717)) || ((-(v4)) == (fi27 << ((4294967295) & 63)))) {
                        _ = len(v3)
                        var v29 string = "foo bar"
                        _ = v29
                    } else {
                        var v30 uint64 = 559
                        var v31 int = (fi27 << ((452) & 63))
                        fmt.Println(v26, v26, v30)
                        _ = v30
                        _ = v31
                    }
                case "hello":
                    v26 = (v26 % ((v26) | 1))
                case "\n\t":
                    if v1 {
                        switch v4 {
                        case 42:
                            v32 := GenS2{F0: 1}
                            _ = v32
                        case 32768:
                            v26 = 239
                        case -839:
                            v2 += (v2 & (^(v2)))
                        default:
                            var v33 bool = true
                            _ = v33
                        }
                        v34 := GenS4{F0: 2, F1: 172, F2: 40}
                        _ = v34
                    } else {
                        switch "" {
                        case "":
                            v1 = v1
                        case "hello":
                            var v35 int16 = 42
                            _ = v35
                        case "a":
                            for fi36 := 0; fi36 < 2; fi36++ {
                                fmt.Println(v4)
                                _ = fi36
                            }
                        }
                    }
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    v38 := GenS5{F0: 65535, F1: (-(v2))}
                    if (!v1) {
                        v3[(-(v4))] = (("a" + "foo bar") >= ("a" + ""))
                        if (!v1) {
                            v39 := GenS6{F0: 4294967295}
                            var v40 uint16 = 42
                            var v41 float32 = (1e10 + float32(1e10))
                            _ = v39
                            _ = v40
                            _ = v41
                        } else {
                            v1 = (v26 <= (v26 / ((144) | 1)))
                            v26 += 127
                        }
                    }
                    _ = fi37
                    _ = v38
                }
                _ = fi27
                _ = v28
            }
        }
    default:
        var v42 float32 = 1e10
        _ = v42
    }
    v43 := GenS2{F0: 32767}
    switch ("hello" + "a") {
    case "\n\t":
        v44 := GenS5{F0: 1, F1: (v2 % (((v2 & (-(v2)))) | 1))}
        _ = v44
    default:
        switch -9223372036854775808 {
        case 739:
            var v45 int16 = 589
            _ = v45
        case 42:
            var v46 uint8 = v26
            _ = v46
        default:
            for fi47 := 0; fi47 < 15; fi47++ {
                var v48 uint8 = 42
                _ = fi47
                _ = v48
            }
        }
    }
    fmt.Println(v1, v26)
    v49 := GenS2{F0: 595}
    fmt.Println(v4, v1, v1)
    var v50 float64 = 1.5
    var v51 float64 = (1e10 - v50)
    if (489 < 42) {
        v52 := GenS2{F0: 32767}
        v51 -= float64(1.5)
        v2 = -1
        _ = v52
    }
    v2 = v2
    for fi53 := 0; fi53 < 5; fi53++ {
        if (!(true && (!(0 <= 0)))) {
            v50 += float64((0.1 - ((1e10 - 0.1) - float32(-2.25))))
        }
        var v54 int32 = 51
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        v26 -= 2
        v50 = v50
        _ = fi55
    }
    var v56 int64 = (v2 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v50
    _ = v51
    _ = v56
}
