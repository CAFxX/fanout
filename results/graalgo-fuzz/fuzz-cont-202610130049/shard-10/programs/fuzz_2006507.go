package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 string; F2 string }
func genFn0(p0 int, p1 uint64) bool {
    v5 := []int{9223372036854775807, -62}
    if len(v5) > 0 {
        v5[0] = (p0 - (p0 >> ((18446744073709551615) & 63)))
    }
    switch (p0 | (p0 - (-(p0)))) {
    case 42:
        var v6 int8 = -128
        _ = v6
    case 4294967295:
        var v7 uint32 = 2
        _ = v7
    default:
        if (!(false || false)) {
            p1 = (^(p1))
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return true
}
func genFn1(p0 int) uint {
    if true {
        v8 := GenS1{F0: ("\n\t" + "a"), F1: "a", F2: -9223372036854775808}
        p0 = (p0 ^ p0)
        p0 += 810
        _ = v8
    } else {
        fmt.Println(p0)
    }
    var v9 uint16 = 42
    var v10 uint8 = 248
    v11 := []int{p0}
    v12 := GenS1{F0: "\n\t", F1: "foo bar", F2: -92}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 2
}

func main() {
    v1 := []int{127}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar"}
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS1{F0: "a", F1: v13, F2: 42}
        _ = fi14
        _ = v15
    }
    switch "a" {
    case "foo bar":
        v16 := GenS2{F0: 4294967295, F1: 42, F2: 255}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS2{F0: 42, F1: 212, F2: 1}
                _ = fi18
                _ = v19
            }
            fmt.Println(v13)
            v20 := GenS2{F0: 0, F1: 42, F2: 255}
            _ = fi17
            _ = v20
        }
    }
    v13 = v13
    _ = len(v1)
    if true {
        var v21 uint8 = 127
        _ = v21
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v13 = v13
        if false {
            delete(v2, 255)
            if false {
                var v22 uint16 = 32767
                _ = v22
            }
            v1 = append(v1, -9223372036854775808)
        }
    }
    if (42 != 4294967295) {
        var v23 int64 = 42
        delete(v2, 42)
        if ((1 > 1) && ((true && ((float32(-2.25) > 0.1) && (!true))) || false)) {
            switch 2147483647 {
            case 1:
                switch "" {
                case "foo bar":
                    var v24 uint16 = 65535
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 int8 = 100
                        v2[1] = ((!(255 < 0)) || (((0.0 + (1.5 + 0.0)) > 0.1) || ((1.5 - float64(3.14159)) != -2.25)))
                        switch 32767 {
                        case 9223372036854775807:
                            for fi27 := 0; fi27 < 13; fi27++ {
                                if len(v1) > 0 {
                                    v1[0] = (fi25 + 32767)
                                }
                                _ = fi27
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            case -1:
                v13 = (((v13 + "hello") + (("hello" + ("\n\t" + v13)) + (v13 + v13))) + (v13 + (v13 + v13)))
            }
            if ((3.14159 / 3.14159) == 1.5) {
                var v28 uint32 = 1
                _ = len(v4)
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                fmt.Println(fi29)
                _ = fi29
            }
        }
        _ = v23
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v13 = ("a" + ("foo bar" + v13))
        _ = fi30
    }
    switch 42 {
    case 4294967295:
        v13 = "foo bar"
    case 32767:
        if false {
            v2[2147483648] = (9223372036854775807 == -930)
            fmt.Println(v13)
        } else {
            var v31 uint64 = 75
            v32 := GenS1{F0: ("a" + (v13 + v13)), F1: "hello", F2: 1}
            _ = v31
            _ = v32
        }
    case -1:
        switch 1 {
        case -671:
            v33 := GenS1{F0: v13, F1: (v13 + v13), F2: -1}
            _ = v33
        case 9223372036854775807:
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS3{F0: 128, F1: 42, F2: -1}
                _ = fi34
                _ = v35
            }
        case -1:
            v3[2147483647] = 583
        }
    }
    fmt.Println(v13)
    if true {
        v13 = (v13 + "\n\t")
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := GenS4{F0: 0, F1: 0}
            v13 = ""
            var v39 float32 = -2.25
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi36
    }
    fmt.Println(v13)
    switch v13 {
    case "foo bar":
        v40 := GenS2{F0: 42, F1: 122, F2: 0}
        _ = v40
    case "\n\t":
        v41 := GenS5{F0: 0, F1: "foo bar", F2: v13}
        _ = v41
    case "hello":
        fmt.Println(v13)
    }
    var v42 float64 = (((1.5 - 0.0) * 1e10) - 0.1)
    var v43 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
