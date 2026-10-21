package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int16; F2 int8 }
func genFn0(p0 int, p1 uint64) uint {
    var v6 uint64 = p1
    v7 := GenS1{F0: "a", F1: 1}
    v8 := []int{(p0 % ((p0) | 1)), (p0 << ((42) & 63)), 160, (p0 ^ 2147483647)}
    for fi9 := 0; fi9 < 13; fi9++ {
        switch ((("\n\t" + (("a" + "foo bar") + "hello")) + "\n\t") + "foo bar") {
        case "foo bar":
            v10 := GenS2{F0: 255, F1: 753, F2: 72}
            _ = v10
        case "":
            var v11 float32 = (0.0 * -2.25)
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v8 = append(v8, (p0 << ((4294967295) & 63)))
                if true {
                    v8 = append(v8, (fi9 * 796))
                }
                switch ("a" + "") {
                case "a":
                    var v13 int16 = 42
                    _ = v13
                case "hello":
                    p1 -= v6
                case "\n\t":
                    v14 := GenS3{F0: -1, F1: v6}
                    fmt.Println(p0)
                    var v15 int8 = 1
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = v14
                    _ = v15
                default:
                    var v16 uint = 0
                    _ = v16
                }
                _ = fi12
            }
        }
        if (42 > 434) {
            var v17 int64 = -1
            _ = v17
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    fmt.Println(p0)
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            if true {
                p1 = (p1 + p1)
                if (1 < -559) {
                    var v20 int64 = 0
                    p1 = p1
                    v21 := GenS3{F0: (v20 - -9223372036854775808), F1: 18446744073709551615}
                    _ = v20
                    _ = v21
                }
                if false {
                    var v22 uint8 = 42
                    p1 = (p1 | (v22 / (((p1 / ((p1) | 1))) | 1)))
                    _ = v22
                } else {
                    fmt.Println(p0, p0, fi18)
                }
            } else {
                if (0.1 >= 3.14159) {
                    switch (fi18 + (fi18 ^ (fi19 - 128))) {
                    case 181:
                        v23 := []int{-2147483648, (fi18 | (fi19 / ((658) | 1)))}
                        _ = v23
                    case 4294967295:
                        fmt.Println(p1, fi18)
                    case -954:
                        v24 := []int{-9223372036854775808}
                        _ = v24
                    }
                    if ("foo bar" <= ("" + ("a" + "\n\t"))) {
                        v25 := []int{int(p0), -1, fi18, 1}
                        if true {
                            var v26 int8 = 0
                            _ = v26
                        }
                        var v27 int16 = 32767
                        var v28 bool = ((int64(v27) <= -1) || (((1.5 + ((1e10 / 1.5) / 0.1)) * 1.5) < 0.0))
                        _ = v25
                        _ = v27
                        _ = v28
                    }
                }
                p1 = 128
                v29 := []int{127, fi18, 32767, 128}
                _ = v29
            }
            v30 := []int{}
            p0 = 2
            _ = fi19
            _ = v30
        }
        _ = fi18
    }
    if true {
        v31 := GenS3{F0: -243, F1: 18446744073709551615}
        _ = v31
    } else {
        var v32 bool = true
        p0 -= 65535
        if v32 {
            var v33 float32 = 3.14159
            p0 = p0
            _ = v33
        }
        _ = v32
    }
    var v34 uint32 = (p0 | 65535)
    var v35 bool = (true || (true && true))
    _ = p0
    _ = p1
    _ = v34
    _ = v35
    return (p1 | p1)
}

func main() {
    var v1 float64 = 0.0
    var v2 float64 = (1.5 / 1e10)
    var v3 int16 = 0
    var v4 int64 = -9223372036854775808
    var v5 uint16 = 32767
    var v36 float32 = (0.1 + float32(v1))
    if ((v36 != float32((v2 + 0.1))) || true) {
        v37 := GenS4{F0: 0, F1: 18446744073709551615, F2: (^(v5))}
        _ = v37
    } else {
        var v38 uint16 = (v5 & v5)
        _ = v38
    }
    if true {
        switch (("hello" + "a") + "\n\t") {
        case "hello":
            v1 -= ((v2 + v2) - (v2 * 0.0))
        case "foo bar":
            var v39 int16 = v3
            _ = v39
        case "":
            v2 += v2
        default:
            v2 = 0.0
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS2{F0: 255, F1: 1, F2: (^(v3))}
        fmt.Println(fi40, v4, v1)
        v1 += -2.25
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        v43 := GenS3{F0: v4, F1: 42}
        _ = fi42
        _ = v43
    }
    v44 := GenS5{F0: 2147483648, F1: (v3 >> ((955) & 63)), F2: 0}
    if true {
        v45 := GenS5{F0: 2, F1: (^(v3)), F2: -128}
        v46 := GenS1{F0: ("foo bar" + ""), F1: (v3 << ((uint(v5)) & 63))}
        _ = v45
        _ = v46
    }
    if ("\n\t" >= "foo bar") {
        v47 := []int{9223372036854775807, -635, 236, -9223372036854775808}
        _ = v47
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        switch ("foo bar" + ((("hello" + "foo bar") + ("foo bar" + "foo bar")) + ((("" + ("\n\t" + ("foo bar" + "hello"))) + ("foo bar" + "a")) + "\n\t"))) {
        case "a":
            v49 := []int{(fi48 << ((65535) & 63))}
            _ = v49
        case "\n\t":
            for fi50 := 0; fi50 < 11; fi50++ {
                fmt.Println(v5, v2, v3)
                _ = fi50
            }
        default:
            v51 := []int{0, fi48, (fi48 << ((2) & 63))}
            _ = v51
        }
        v52 := GenS3{F0: -188, F1: 42}
        _ = fi48
        _ = v52
    }
    var v53 uint = 348
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 19; fi55++ {
            v56 := []int{}
            fmt.Println(v4, fi55, fi55)
            _ = fi55
            _ = v56
        }
        fmt.Println(v53, v3, v5)
        v57 := GenS4{F0: 9223372036854775807, F1: v53, F2: 535}
        _ = fi54
        _ = v57
    }
    var v58 int32 = 0
    var v59 int32 = 93
    if true {
        switch ("" + "hello") {
        case "a":
            v60 := []int{32768}
            _ = v60
        }
        var v61 int8 = -41
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v5
    _ = v53
    _ = v58
    _ = v59
}
