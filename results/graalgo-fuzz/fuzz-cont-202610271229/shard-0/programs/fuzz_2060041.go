package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint64; F1 uint64 }
func genFn0(p0 bool) uint32 {
    if (1 > 0) {
        if p0 {
            if p0 {
                p0 = (128 == 94)
                v4 := []int{-9223372036854775808, 1, 127}
                _ = v4
            } else {
                fmt.Println(p0)
                v5 := []int{}
                _ = v5
            }
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: 2147483647, F1: -823}
                _ = fi6
                _ = v7
            }
        }
        v8 := []int{-907}
        p0 = p0
        _ = v8
    }
    if (0 <= 2147483647) {
        fmt.Println(p0)
        v9 := GenS1{F0: -1, F1: 1}
        p0 = (p0 && true)
        _ = v9
    } else {
        v10 := GenS1{F0: -2147483648, F1: 42}
        v11 := []int{}
        if true {
            p0 = p0
            p0 = (1 < -23)
            fmt.Println(p0)
        } else {
            fmt.Println(p0)
            if p0 {
                if p0 {
                    if len(v11) > 0 {
                        v11[0] = 9223372036854775807
                    }
                    for fi12 := 0; fi12 < 17; fi12++ {
                        if true {
                            if (("a" != (("\n\t" + ("hello" + "\n\t")) + ("hello" + ((("foo bar" + "a") + (((("hello" + "hello") + "\n\t") + "\n\t") + ("hello" + "a"))) + ("foo bar" + ("hello" + "")))))) || (float32(-2.25) < 3.14159)) {
                                p0 = (65535 != 65535)
                                p0 = (0 >= 212)
                            }
                            p0 = p0
                        }
                        _ = fi12
                    }
                    var v13 int64 = 0
                    _ = v13
                } else {
                    var v14 int16 = 1000
                    _ = v14
                }
            } else {
                p0 = (p0 || (true || (65535 != 2)))
            }
        }
        _ = v10
        _ = v11
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint16 = 255
            v18 := GenS2{F0: 52}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint16 = 1
        switch (fi15 << ((2) & 63)) {
        case -659:
            v19 -= 145
        case 350:
            var v20 uint32 = 2147483647
            _ = v20
        default:
            v21 := []int{(fi15 - (-(fi15))), fi15, fi15}
            _ = v21
        }
        _ = fi15
        _ = v19
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 string, p1 uint, p2 uint64) uint {
    p1 = 1
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS3{F0: 2147483648, F1: (p1 * (p1 / (((p1 / ((2) | 1))) | 1)))}
        _ = fi22
        _ = v23
    }
    var v24 uint8 = 0
    var v25 int8 = -56
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 1
}

func main() {
    var v1 string = ""
    v2 := []string{(v1 + "a"), (v1 + v1), v1}
    v3 := []string{(("" + "hello") + (((v1 + v1) + v1) + v1)), v1, "\n\t", "\n\t"}
    v1 = (v1 + "hello")
    switch -1 {
    case 657:
        for fi26 := 0; fi26 < 19; fi26++ {
            if len(v2) > 0 {
                v2[0] = v1
            }
            if ("hello" > v1) {
                v1 = v1
            } else {
                for fi27 := 0; fi27 < 13; fi27++ {
                    fmt.Println(fi26)
                    var v28 uint8 = 255
                    switch v1 {
                    case "":
                        switch (fi27 ^ fi26) {
                        case -607:
                            switch fi27 {
                            case 615:
                                fmt.Println(fi27)
                                v1 = v1
                            case 9223372036854775807:
                                v1 = (v1 + ("hello" + v1))
                            }
                        case 9223372036854775807:
                            _ = len(v3)
                        case 2147483647:
                            switch 776 {
                            case 127:
                                var v29 float64 = (1.5 * 0.1)
                                _ = v29
                            }
                        default:
                            if false {
                                switch v1 {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        }
                    }
                    v1 = v1
                    _ = fi27
                    _ = v28
                }
            }
            var v30 float32 = 0.1
            _ = fi26
            _ = v30
        }
    case 32768:
        var v31 int64 = 0
        _ = v31
    case 0:
        v32 := GenS2{F0: 255}
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi33
    }
    var v34 uint32 = 535
    var v35 int32 = -1
    var v36 int = 32768
    v37 := GenS4{F0: (v36 | -33), F1: 2}
    v38 := GenS5{F0: -2147483648}
    v39 := GenS6{F0: 18446744073709551615, F1: 42}
    fmt.Println(v36, v34, v34)
    v40 := GenS7{F0: 42, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
