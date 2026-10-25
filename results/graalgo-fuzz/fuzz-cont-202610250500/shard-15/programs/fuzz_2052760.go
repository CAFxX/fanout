package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0() string {
    var v5 uint64 = 1
    v5 = v5
    _ = v5
    return "hello"
}
func genFn1(p0 int64, p1 int32) int {
    var v6 float64 = 3.14159
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 707
}

func main() {
    var v1 int = -1
    var v2 uint16 = 255
    var v3 float64 = float64(1.5)
    var v4 int64 = -1
    v8 := GenS1{F0: -1, F1: v2, F2: ""}
    var v9 bool = (561 != 18446744073709551615)
    var v10 uint = 565
    var v11 uint8 = 2
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "hello":
        switch ("hello" + "\n\t") {
        case "a":
            v12 := []int{-9223372036854775808, (v1 * -2147483648)}
            _ = v12
        case "hello":
            var v13 float32 = 0.0
            _ = v13
        default:
            var v14 uint64 = uint64(v2)
            _ = v14
        }
    case "":
        v10 = v10
    default:
        v15 := []int{}
        _ = v15
    }
    fmt.Println(v4)
    v16 := []int{(v1 - v1), v1, v1, v1}
    var v17 uint32 = 2147483648
    var v18 int64 = v4
    v19 := GenS1{F0: -2147483648, F1: (v2 + v2), F2: ""}
    for fi20 := 0; fi20 < 16; fi20++ {
        switch ("\n\t" + ("foo bar" + (("foo bar" + "foo bar") + ""))) {
        case "\n\t":
            var v21 int64 = v18
            _ = v21
        case "a":
            var v22 float32 = 0.0
            _ = v22
        case "foo bar":
            if v9 {
                v4 += -1
                v2 -= 807
            } else {
                v18 -= v4
                if len(v16) > 0 {
                    v16[0] = (fi20 % ((v1) | 1))
                }
                var v23 int16 = 0
                _ = v23
            }
        }
        if (!v9) {
            v16 = append(v16, (fi20 >> ((65535) & 63)))
        } else {
            var v24 uint = (v10 * (v10 & (v10 & v10)))
            if v9 {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    if v9 {
                        v9 = false
                    } else {
                        var v26 float64 = v3
                        var v27 int16 = -1
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
            }
            var v28 int8 = -59
            _ = v24
            _ = v28
        }
        _ = fi20
    }
    v29 := GenS2{F0: 42, F1: "hello"}
    v1 = 25
    switch "\n\t" {
    case "\n\t":
        v30 := GenS1{F0: 0, F1: 2, F2: (("foo bar" + (("hello" + ("" + ("" + "foo bar"))) + ("hello" + (("a" + (("" + "a") + "")) + ("foo bar" + "a"))))) + "")}
        _ = v30
    case "foo bar":
        v31 := GenS2{F0: (v11 >> ((v10) & 63)), F1: "hello"}
        _ = v31
    case "":
        if ((537 >= (v10 * v10)) || v9) {
            _ = len(v16)
            for fi32 := 0; fi32 < 7; fi32++ {
                _ = len(v16)
                v3 -= v3
                _ = fi32
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                v34 := GenS1{F0: int32(v1), F1: (v2 << (((v10 & 65535)) & 63)), F2: ""}
                _ = fi33
                _ = v34
            }
        }
    }
    v16 = append(v16, v1)
    if false {
        for fi35 := 0; fi35 < 14; fi35++ {
            if v9 {
                var v36 float64 = float64(float32((-2.25 - v3)))
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                if (1 == v4) {
                    if len(v16) > 0 {
                        v16[0] = (fi35 - -915)
                    }
                    _ = len(v16)
                    var v37 uint16 = (v2 + (v2 % ((65535) | 1)))
                    _ = v37
                } else {
                    v16 = append(v16, fi35)
                    v16 = append(v16, genFn1(9223372036854775807, -1))
                }
                _ = v36
            } else {
                var v38 int32 = 42
                v10 += v10
                switch int(v18) {
                case -9223372036854775808:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
                _ = v38
            }
            if v9 {
                if v9 {
                    var v39 int8 = 1
                    v16 = append(v16, (v1 % ((fi35) | 1)))
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 4; fi40++ {
                        v3 = (float64(1.5) - (0.0 * v3))
                        _ = len(v16)
                        _ = fi40
                    }
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v42 := GenS3{F0: -128, F1: 1000}
                        v43 := GenS3{F0: 127, F1: -369}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                }
                v11 += (v11 << ((1) & 63))
                switch "hello" {
                case "foo bar":
                    var v44 int8 = 1
                    _ = v44
                case "":
                    if (!(18446744073709551615 <= 4294967295)) {
                        for fi45 := 0; fi45 < 5; fi45++ {
                            v1 = genFn1(v4, -82)
                            _ = fi45
                        }
                        v17 = 137
                        v4 = (v4 + v4)
                    }
                case "\n\t":
                    if v9 {
                        fmt.Println(v1, v2, fi35)
                        v3 = v3
                    } else {
                        var v46 uint = (v10 + (v10 * v10))
                        _ = v46
                    }
                }
            } else {
                v11 = (v11 * v11)
                v47 := GenS2{F0: (v11 << (((v10 >> ((1) & 63))) & 63)), F1: ("\n\t" + "foo bar")}
                var v48 uint16 = (v2 + (v2 + 42))
                _ = v47
                _ = v48
            }
            _ = fi35
        }
        var v49 string = "hello"
        fmt.Println(v49, v49)
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
