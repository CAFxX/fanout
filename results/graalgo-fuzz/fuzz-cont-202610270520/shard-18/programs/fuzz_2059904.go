package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 int8; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint8) int32 {
    p0 = p0
    p0 = 1
    p0 = 1
    var v6 uint = 2
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int8 = 42
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1() int {
    var v10 uint16 = 42
    v10 = v10
    if false {
        var v11 int32 = 97
        _ = v11
    } else {
        var v12 int8 = 70
        var v13 float32 = (0.0 / 1.0)
        fmt.Println(v10, v13)
        _ = v12
        _ = v13
    }
    _ = v10
    return 827
}

func main() {
    v1 := []uint32{2147483647, 0, 0, 0}
    var v2 int8 = 42
    v3 := []uint32{2147483647, 2147483647, 2147483647, 1}
    var v4 string = ""
    var v5 float32 = 3.14159
    if (9223372036854775807 < 1) {
        v14 := GenS1{F0: 42, F1: 1, F2: -1}
        fmt.Println(v2, v5, v4)
        _ = v14
    }
    v15 := GenS1{F0: 212, F1: -417, F2: 128}
    fmt.Println(v4)
    switch "a" {
    case "a":
        if false {
            v16 := GenS2{F0: -128, F1: "hello"}
            _ = v16
        } else {
            var v17 uint16 = 714
            _ = v17
        }
    case "hello":
        for fi18 := 0; fi18 < 10; fi18++ {
            switch (("a" + ("a" + v4)) + (((v4 + v4) + ("foo bar" + (v4 + v4))) + ("a" + v4))) {
            case "a":
                v19 := GenS1{F0: 2, F1: 1, F2: (fi18 + -9223372036854775808)}
                _ = v19
            case "foo bar":
                for fi20 := 0; fi20 < 12; fi20++ {
                    v3 = append(v3, 65535)
                    v5 += v5
                    var v21 uint32 = 154
                    fmt.Println(v21, v5, v4)
                    _ = fi20
                    _ = v21
                }
            case "hello":
                fmt.Println(fi18)
            }
            if false {
                if true {
                    fmt.Println(v4, fi18)
                    v5 = -2.25
                    v4 = "\n\t"
                } else {
                    if true {
                        v4 = v4
                    } else {
                        v4 = v4
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            if (true || false) {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int8 = (v2 - v2)
                    _ = fi22
                    _ = v23
                }
                v3 = append(v3, 65535)
            } else {
                var v24 int8 = (v2 - -1)
                _ = v24
            }
            _ = fi18
        }
    default:
        switch -9223372036854775808 {
        case 733:
            v25 := GenS3{F0: 18446744073709551615, F1: v2, F2: 9223372036854775807}
            _ = v25
        }
    }
    v2 += v2
    var v26 int8 = -97
    v27 := GenS4{F0: -2147483648}
    _ = len(v3)
    switch 9223372036854775807 {
    case -2147483648:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v28 string = v4
    v28 = ("" + (v28 + (v4 + "hello")))
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            v5 -= ((v5 + 1e10) * (v5 + 0.1))
            v2 = (v2 / ((v26) | 1))
            if len(v1) > 0 {
                v1[0] = 796
            }
            _ = fi30
        }
        if (915 < 1) {
            if true {
                var v31 int16 = int16(fi29)
                _ = v31
            } else {
                var v32 int8 = (v2 * (v2 % ((v26) | 1)))
                _ = v32
            }
            v28 = v4
            fmt.Println(v2, v26)
        } else {
            v33 := GenS2{F0: -128, F1: "a"}
            var v34 uint16 = 1
            if false {
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 uint = 1
                        _ = fi35
                        _ = v36
                    }
                }
                fmt.Println(v5, fi29, fi29)
            } else {
                v26 = (v2 | (v2 * v2))
                var v37 uint32 = 2
                _ = v37
            }
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    var v38 uint = 869
    for fi39 := 0; fi39 < 20; fi39++ {
        v5 -= float32((1e10 - 1.5))
        v4 = v4
        _ = fi39
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        if true {
            if true {
                switch v4 {
                case "hello":
                    switch genFn1() {
                    case 0:
                        if false {
                            v28 = "\n\t"
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        }
                    case 9223372036854775807:
                        v26 = (-(v2))
                    default:
                        v28 = ("hello" + "a")
                    }
                case "foo bar":
                    for fi41 := 0; fi41 < 13; fi41++ {
                        var v42 int16 = 506
                        var v43 int8 = v26
                        v5 += -2.25
                        var v44 string = v4
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                }
                v45 := GenS3{F0: 248, F1: v26, F2: genFn1()}
                v2 += v2
                _ = v45
            } else {
                if (!(127 <= (v26 << ((4294967295) & 63)))) {
                    fmt.Println(v4)
                    v26 = -110
                    v46 := GenS5{F0: 728}
                    _ = v46
                }
                v1 = append(v1, 4294967295)
            }
        }
        var v47 int8 = 37
        _ = fi40
        _ = v47
    }
    switch -851 {
    case -2147483648:
        switch 4294967295 {
        case 34:
            v48 := GenS2{F0: v2, F1: v28}
            _ = v48
        case 200:
            if len(v3) > 0 {
                v3[0] = 666
            }
        default:
            v3 = append(v3, 65535)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
