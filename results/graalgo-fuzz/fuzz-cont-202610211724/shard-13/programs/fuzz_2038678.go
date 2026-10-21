package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
func genFn0(p0 int8, p1 uint32) uint32 {
    switch (("" + ("\n\t" + "foo bar")) + "hello") {
    case "\n\t":
        var v7 uint8 = 42
        _ = v7
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{(fi8 ^ (fi8 | 2147483647)), (fi8 + fi8), (fi8 + fi8), (fi8 & fi8)}
            if len(v9) > 0 {
                v9[0] = int(p1)
            }
            if (true || true) {
                var v10 uint32 = 2147483648
                _ = v10
            } else {
                var v11 float64 = (0.1 * (0.0 * -2.25))
                p0 -= p0
                v12 := GenS1{F0: 42, F1: 1}
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    }
    v13 := []int{-870}
    var v14 bool = true
    if v14 {
        var v15 string = (("hello" + "a") + "a")
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int32 = 565
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        if ((!(p0 > (^(p0)))) && v14) {
            if ((3.14159 * 1.5) > 1.5) {
                switch "a" {
                case "\n\t":
                    _ = len(v13)
                case "foo bar":
                    p1 = 694
                case "":
                    for fi18 := 0; fi18 < 13; fi18++ {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi18
                    }
                default:
                    if len(v13) > 0 {
                        v13[0] = 9223372036854775807
                    }
                }
                var v19 uint = 2
                _ = v19
            }
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return p1
}
func genFn1() int16 {
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS1{F0: 1, F1: 1}
        var v22 int = (fi20 * (^(fi20)))
        if false {
            if (-866 > -1) {
                v23 := GenS2{F0: 65535, F1: 42, F2: 2}
                if true {
                    var v24 string = (("hello" + "") + ("foo bar" + "a"))
                    _ = v24
                } else {
                    v25 := []int{(fi20 / ((v22) | 1))}
                    if false {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            fmt.Println(v22, fi20)
                            _ = fi26
                        }
                        fmt.Println(fi20, v22)
                        fmt.Println(v22, v22)
                    } else {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            var v28 string = ("foo bar" + "foo bar")
                            _ = fi27
                            _ = v28
                        }
                        v22 = v22
                    }
                    v29 := GenS3{F0: 118}
                    _ = v25
                    _ = v29
                }
                _ = v23
            }
            v22 = fi20
            v22 = fi20
        } else {
            v22 = v22
            fmt.Println(fi20, fi20)
            fmt.Println(v22)
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    v30 := []int{-2147483648, 32768, -2147483648}
    v31 := GenS4{F0: 2147483647, F1: 1}
    _ = v30
    _ = v31
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 897
    var v3 bool = false
    var v4 uint8 = 42
    var v5 uint16 = 65535
    var v6 int32 = 42
    v32 := GenS5{F0: "foo bar", F1: (v4 & v4), F2: true}
    var v33 string = "hello"
    var v34 int64 = (v1 - (v1 & v1))
    v3 = (0 >= 0)
    if v3 {
        v35 := []int{127, -1}
        if v3 {
            var v36 uint = 42
            _ = v36
        }
        v37 := GenS3{F0: (v4 | v4)}
        _ = v35
        _ = v37
    } else {
        var v38 uint8 = 127
        if v3 {
            v33 = "a"
            v2 += v2
            var v39 int64 = (^(v34))
            _ = v39
        } else {
            fmt.Println(v33, v4)
            v40 := []int{0, 453, 2147483647}
            if v3 {
                v41 := GenS1{F0: (^(v5)), F1: 841}
                for fi42 := 0; fi42 < 8; fi42++ {
                    fmt.Println(v4)
                    fmt.Println(v3, v1, v4)
                    _ = fi42
                }
                _ = v41
            }
            _ = v40
        }
        _ = v38
    }
    v43 := []int{9223372036854775807, 525, 1, 32768}
    if v3 {
        var v44 uint8 = v4
        _ = v44
    }
    v1 -= (v34 / ((v1) | 1))
    var v45 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v6
}
