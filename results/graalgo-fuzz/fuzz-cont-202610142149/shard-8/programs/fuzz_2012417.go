package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 string) uint32 {
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "" {
        case "\n\t":
            var v6 uint16 = 581
            _ = v6
        case "hello":
            p0 = ((0.1 * 0.1) == (-2.25 - (1e10 / 1.5)))
        default:
            var v7 int64 = -9223372036854775808
            _ = v7
        }
        var v8 uint = 255
        v9 := GenS1{F0: 0, F1: 65535}
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 uint8 = 139
    var v11 uint16 = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 2
}
func genFn1(p0 string, p1 int32) uint16 {
    var v12 string = p0
    if false {
        p1 -= (^(p1))
        for fi13 := 0; fi13 < 2; fi13++ {
            switch (p0 + ("foo bar" + "\n\t")) {
            case "foo bar":
                p0 = v12
            default:
                fmt.Println(p1)
            }
            var v14 uint16 = 1
            _ = fi13
            _ = v14
        }
        v12 = v12
    }
    v15 := []int{2147483647, -2147483648, -189}
    for fi16 := 0; fi16 < 3; fi16++ {
        if len(v15) > 0 {
            v15[0] = (fi16 & (fi16 << ((4294967295) & 63)))
        }
        _ = fi16
    }
    if ((42 < -1) && ((-2.25 + 1e10) > 1.5)) {
        p1 = 86
        var v17 uint16 = 2
        _ = v17
    } else {
        v18 := GenS2{F0: 251, F1: 2147483648, F2: -814}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return 684
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -9223372036854775808
    var v3 int32 = -1
    v4 := []int{-778, 42, 1}
    v1 += 128
    if ((true || false) || true) {
        if true {
            var v19 uint16 = 255
            v1 = (^(v1))
            for fi20 := 0; fi20 < 6; fi20++ {
                switch "a" {
                case "a":
                    if (69 != 158) {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v4 = append(v4, fi20)
                            var v22 int32 = (v3 | 42)
                            _ = fi21
                            _ = v22
                        }
                    }
                default:
                    fmt.Println(fi20)
                }
                _ = fi20
            }
            _ = v19
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            v2 = (-(v2))
            var v24 uint16 = 483
            _ = fi23
            _ = v24
        }
    }
    switch "foo bar" {
    case "a":
        if (("a" + (("foo bar" + "foo bar") + "a")) != "a") {
            switch (("hello" + ("hello" + "a")) + "foo bar") {
            case "foo bar":
                fmt.Println(v1, v2, v2)
            case "a":
                var v25 uint16 = genFn1((("\n\t" + ("" + ((("" + "a") + "") + ("" + "a")))) + (((("\n\t" + "\n\t") + ("a" + "\n\t")) + "\n\t") + "hello")), v3)
                _ = v25
            default:
                v2 = (v2 / ((v2) | 1))
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: 2, F1: 65535}
                v3 += (v3 & 42)
                _ = fi26
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 19; fi28++ {
                switch fi28 {
                case 32768:
                    fmt.Println(v2)
                case -9223372036854775808:
                    switch int(v1) {
                    case 1:
                        v3 = v3
                    case 57:
                        if ((216 > 1) && (0.1 > 0.1)) {
                            v4 = append(v4, (fi28 >> ((4294967295) & 63)))
                            var v29 int = fi28
                            for fi30 := 0; fi30 < 11; fi30++ {
                                var v31 bool = true
                                var v32 float32 = (1e10 - 0.1)
                                var v33 int8 = -128
                                switch ((("a" + "hello") + "hello") + "a") {
                                case "hello":
                                    if (v2 > (v2 / (((v2 | (v2 >> ((1) & 63)))) | 1))) {
                                        v32 = float32(-2.25)
                                    }
                                default:
                                    v31 = v31
                                }
                                if len(v4) > 0 {
                                    v4[0] = v29
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            _ = v29
                        }
                    case 856:
                        v34 := GenS2{F0: 65535, F1: 2147483647, F2: fi28}
                        _ = v34
                    }
                case -690:
                    var v35 uint64 = 0
                    _ = v35
                default:
                    var v36 string = "a"
                    _ = v36
                }
                _ = fi28
            }
        }
    default:
        var v37 int16 = int16(v3)
        _ = v37
    }
    _ = len(v4)
    var v38 int64 = 831
    var v39 uint64 = 0
    v38 = (v38 / (((v38 - (v38 & (v38 & (v38 & v38))))) | 1))
    for fi40 := 0; fi40 < 10; fi40++ {
        v41 := GenS3{F0: (-(v38))}
        var v42 int = (fi40 - (fi40 ^ (fi40 - 9223372036854775807)))
        v42 += (fi40 ^ fi40)
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
