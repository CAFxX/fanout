package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 bool = (0 >= 451)
    var v2 uint8 = 255
    v3 := []string{("a" + "hello"), "foo bar"}
    var v4 int8 = 42
    v5 := []string{("a" + (("foo bar" + "") + "foo bar"))}
    switch 42 {
    case 32767:
        var v6 int8 = (v4 % ((43) | 1))
        _ = v6
    case 1:
        var v7 uint = 395
        _ = v7
    case 490:
        for fi8 := 0; fi8 < 8; fi8++ {
            v1 = ((1 < -588) && (v1 || (65535 < 378)))
            if (276 != 858) {
                v9 := GenS1{F0: 1, F1: 65535}
                for fi10 := 0; fi10 < 17; fi10++ {
                    if v1 {
                        switch "foo bar" {
                        case "\n\t":
                            v11 := GenS1{F0: -1, F1: 2}
                            _ = v11
                        }
                        if v1 {
                            if (791 != 42) {
                                fmt.Println(v2, v1)
                                var v12 uint16 = 42
                                _ = v12
                            }
                            fmt.Println(v2)
                        }
                    }
                    _ = fi10
                }
                if true {
                    v2 = (v2 - 255)
                    fmt.Println(v4, v1)
                }
                _ = v9
            } else {
                var v13 int64 = -1
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 int8 = v4
                    var v16 int = -9223372036854775808
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            _ = len(v5)
            _ = fi8
        }
    default:
        for fi17 := 0; fi17 < 20; fi17++ {
            if v1 {
                fmt.Println(v2)
                v18 := GenS2{F0: 373, F1: v2, F2: 2147483648}
                fmt.Println(v4, v2)
                _ = v18
            } else {
                v19 := GenS2{F0: 2, F1: 127, F2: 65535}
                _ = v19
            }
            v4 -= (v4 ^ (v4 & 5))
            v4 -= (v4 & v4)
            _ = fi17
        }
    }
    switch "foo bar" {
    case "hello":
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    case "":
        var v20 uint8 = (v2 & 95)
        _ = v20
    default:
        v21 := GenS2{F0: 32767, F1: (v2 ^ v2), F2: 4294967295}
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS1{F0: int64(v4), F1: 112}
        switch (fi22 << ((65535) & 63)) {
        case 32767:
            v24 := GenS2{F0: 33, F1: 168, F2: 318}
            _ = v24
        case 42:
            v25 := GenS1{F0: 0, F1: 0}
            _ = v25
        case 924:
            var v26 string = "a"
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    v27 := GenS3{F0: "\n\t", F1: (v4 & (v4 / ((v4) | 1)))}
    v28 := GenS1{F0: -1, F1: 276}
    for fi29 := 0; fi29 < 18; fi29++ {
        v1 = ((65535 < 4294967295) || false)
        if (435 == 1000) {
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint32 = 65535
                var v32 uint64 = 4294967295
                v1 = true
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        var v33 int32 = 1
        _ = fi29
        _ = v33
    }
    fmt.Println(v4, v2)
    switch ("hello" + "foo bar") {
    case "":
        switch "a" {
        case "":
            fmt.Println(v1)
        default:
            var v34 int8 = v4
            _ = v34
        }
    case "a":
        v35 := GenS4{F0: (v4 - (v4 + -128))}
        _ = v35
    }
    v36 := GenS5{F0: -1}
    fmt.Println(v1, v1, v2)
    if v1 {
        if v1 {
            v1 = ((!(380 < 2)) && (float32(3.14159) != ((0.0 + 1.5) + -2.25)))
            if len(v3) > 0 {
                v3[0] = (("hello" + "foo bar") + ("a" + (("" + ((("a" + "hello") + "foo bar") + "\n\t")) + ("hello" + ""))))
            }
            v5 = append(v5, (("" + ((((((("hello" + ("foo bar" + "\n\t")) + "a") + ("" + "foo bar")) + "a") + (("\n\t" + "") + "a")) + "\n\t") + ("foo bar" + ("foo bar" + "a")))) + "\n\t"))
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        v4 += 42
        _ = fi37
    }
    fmt.Println(v4, v4)
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (v2 / (((v2 >> ((780) & 63))) | 1))
    var v38 int32 = -2147483648
    if true {
        v38 = 1
        v2 = 127
    } else {
        v39 := GenS1{F0: -1, F1: 0}
        v40 := GenS5{F0: -9223372036854775808}
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
