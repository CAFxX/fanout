package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0() int32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: -468, F1: (true || (1000 != 909))}
        for fi7 := 0; fi7 < 8; fi7++ {
            switch "\n\t" {
            case "a":
                v8 := GenS2{F0: -32768, F1: 2}
                _ = v8
            case "":
                v9 := []int{fi5}
                _ = v9
            default:
                var v10 uint64 = 627
                _ = v10
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    var v11 int64 = -777
    v11 += -1
    v11 += v11
    var v12 uint16 = 1
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 bool) int {
    for fi13 := 0; fi13 < 14; fi13++ {
        p0 = p0
        fmt.Println(fi13, fi13)
        if p0 {
            var v14 uint32 = 65535
            p0 = (-362 == 42)
            _ = v14
        }
        _ = fi13
    }
    if true {
        fmt.Println(p0)
        var v15 int = -9223372036854775808
        _ = v15
    }
    var v16 int16 = 1
    _ = p0
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 float32 = ((3.14159 / float32(-2.25)) + (float32(1.5) * (1.5 / ((1e10 / (1e10 / 1e10)) * (1e10 * 1.5)))))
    var v2 int8 = -128
    v3 := []string{"hello", "\n\t", "\n\t", ((("" + "\n\t") + "a") + ("" + "hello"))}
    v4 := []string{"hello", "", ""}
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 string = "hello"
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int32 = 2147483647
            var v21 uint16 = 255
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint = 366
        _ = fi17
        _ = v18
        _ = v22
    }
    fmt.Println(v1, v1)
    if true {
        var v23 int16 = -32768
        var v24 int16 = v23
        _ = len(v3)
        _ = v23
        _ = v24
    } else {
        fmt.Println(v2, v1)
        switch 42 {
        case -375:
            fmt.Println(v1)
        case 0:
            var v25 uint64 = 42
            _ = v25
        }
    }
    _ = len(v3)
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 int = 4294967295
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: 1, F1: 255, F2: 42}
    if (383 == 0) {
        switch genFn1(true) {
        case 9223372036854775807:
            v29 := GenS2{F0: 631, F1: 0}
            _ = v29
        case -163:
            v3 = append(v3, ("a" + ("hello" + "")))
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            if true {
                v4 = append(v4, "\n\t")
                fmt.Println(fi30)
                switch "a" {
                case "a":
                    for fi31 := 0; fi31 < 3; fi31++ {
                        switch "hello" {
                        case "hello":
                            var v32 uint64 = 250
                            _ = v32
                        }
                        v2 -= v2
                        _ = fi31
                    }
                default:
                    fmt.Println(v2, v2, v2)
                }
            }
            v1 -= (0.1 * 1.5)
            v1 -= 1.5
            _ = fi30
        }
    }
    var v33 int = 32767
    v3 = append(v3, "hello")
    if (((!(v1 >= (v1 - (v1 + v1)))) || false) || false) {
        var v34 int64 = -1
        switch v33 {
        case 282:
            var v35 int = (-(v33))
            _ = v35
        default:
            if (0.1 >= 3.14159) {
                switch (v33 | (v33 / ((v33) | 1))) {
                case 32768:
                    switch (v33 * 127) {
                    case -51:
                        for fi36 := 0; fi36 < 3; fi36++ {
                            for fi37 := 0; fi37 < 13; fi37++ {
                                v2 = (v2 & v2)
                                v38 := GenS2{F0: -314, F1: 2}
                                _ = fi37
                                _ = v38
                            }
                            v39 := GenS3{F0: (v34 << ((18446744073709551615) & 63)), F1: 0, F2: 4294967295}
                            _ = fi36
                            _ = v39
                        }
                    default:
                        v40 := GenS4{F0: 119, F1: ("foo bar" + ((("" + "") + (("\n\t" + "a") + "hello")) + (("hello" + "foo bar") + (("" + "hello") + "hello"))))}
                        _ = v40
                    }
                case 0:
                    var v41 uint = 42
                    _ = v41
                default:
                    v2 = v2
                }
            }
        }
        v42 := GenS2{F0: 42, F1: 59}
        _ = v34
        _ = v42
    }
    fmt.Println(v1, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
}
