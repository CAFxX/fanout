package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 uint) int8 {
    for fi4 := 0; fi4 < 3; fi4++ {
        switch "foo bar" {
        case "":
            switch "hello" {
            case "\n\t":
                if (874 >= 1) {
                    v5 := GenS1{F0: 18446744073709551615, F1: 100}
                    var v6 int8 = 39
                    _ = v5
                    _ = v6
                } else {
                    p1 = 567
                    fmt.Println(p0, p0, p0)
                }
            default:
                fmt.Println(p0, fi4, fi4)
            }
        case "a":
            var v7 int16 = 42
            _ = v7
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: 18446744073709551615, F1: -128}
                _ = fi8
                _ = v9
            }
        }
        _ = fi4
    }
    v10 := GenS1{F0: 18446744073709551615, F1: -128}
    v11 := GenS1{F0: 4294967295, F1: 0}
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -128
}

func main() {
    var v1 uint32 = 1
    var v2 string = (("foo bar" + ("\n\t" + ("hello" + ("" + "")))) + ("" + ((((("" + "a") + "a") + "a") + "") + "hello")))
    var v3 bool = false
    v12 := []int{-9223372036854775808, 32767, 128}
    fmt.Println(v3, v1, v1)
    var v13 float32 = (3.14159 / (0.0 - 1e10))
    var v14 int8 = 127
    var v15 bool = (110 == 211)
    if ((65535 == 65535) && v3) {
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(fi16, fi16)
                switch "hello" {
                case "a":
                    fmt.Println(v13, v2)
                case "foo bar":
                    v1 += v1
                case "":
                    v18 := GenS1{F0: 1, F1: 1}
                    v3 = v3
                    var v19 int64 = -1
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
            var v20 uint64 = 184
            v15 = (0 <= int64(fi16))
            _ = fi16
            _ = v20
        }
        v3 = (v3 && (556 < 42))
        switch v2 {
        case "hello":
            if v3 {
                for fi21 := 0; fi21 < 1; fi21++ {
                    switch (fi21 / ((fi21) | 1)) {
                    case -776:
                        v22 := GenS1{F0: 210, F1: v14}
                        _ = v22
                    case 42:
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if (!v3) {
                                var v24 bool = v3
                                _ = v24
                            }
                            fmt.Println(v3, v15)
                            var v25 uint32 = (^(v1))
                            v12 = append(v12, fi23)
                            _ = fi23
                            _ = v25
                        }
                    }
                    _ = fi21
                }
                if ((1.5 - (v13 - 0.0)) < 1.5) {
                    v12 = append(v12, -502)
                    switch 32767 {
                    case 100:
                        v3 = (!(!(v15 && false)))
                    case 4294967295:
                        var v26 float64 = ((((1.5 / 1.0) / ((0.1 - (((float64(-2.25) - float64(1e10)) + 0.0) + -2.25)) / 1e10)) / -2.25) - ((((0.1 + 1.5) - 0.0) / 1.0) + float64(float32(0.0))))
                        _ = v26
                    case 42:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            var v28 int8 = genFn0((!(!(775 > 1))), 18446744073709551615)
                            var v29 string = v2
                            _ = len(v12)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        v2 = "a"
                    }
                    var v30 uint16 = 720
                    _ = v30
                }
            }
        }
    } else {
        var v31 uint64 = 2
        v32 := GenS2{F0: genFn0((-1 >= 846), 42)}
        _ = v31
        _ = v32
    }
    var v33 uint32 = (v1 & (v1 + v1))
    var v34 uint = 193
    var v35 int8 = -62
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v3 = (!(190 != 128))
        v1 += (v1 - (v1 - v1))
        _ = fi36
    }
    var v37 int16 = 1
    var v38 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
}
