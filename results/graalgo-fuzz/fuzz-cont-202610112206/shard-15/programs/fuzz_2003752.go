package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint = 1
            if (42 >= 1) {
                v10 := []int{fi7, (^(fi8))}
                _ = v10
            }
            var v11 uint64 = 4294967295
            _ = fi8
            _ = v11
            _ = v9
        }
        fmt.Println(fi7)
        v12 := GenS1{F0: 4294967295, F1: 255}
        _ = fi7
        _ = v12
    }
    var v13 uint = 472
    _ = v13
    return 2
}
func genFn1(p0 string, p1 int, p2 uint16) uint64 {
    v14 := GenS1{F0: 0, F1: 32767}
    var v15 float32 = 1e10
    var v16 float64 = -2.25
    var v17 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 int16 = -42
    v2 := make(map[string]string)
    var v3 bool = (2 != 42)
    var v4 uint32 = 65535
    v5 := make(map[int]string)
    var v6 uint = 255
    var v18 float32 = 0.1
    var v19 uint64 = 629
    v20 := GenS1{F0: v6, F1: 1}
    var v21 int8 = 0
    for fi22 := 0; fi22 < 11; fi22++ {
        _ = len(v5)
        fmt.Println(v4, v19)
        var v23 uint8 = 128
        _ = fi22
        _ = v23
    }
    if v3 {
        var v24 uint32 = v4
        switch 1 {
        case -655:
            if (v3 || v3) {
                var v25 int32 = 0
                var v26 int64 = -9223372036854775808
                v26 -= (v26 ^ (-(v26)))
                _ = v25
                _ = v26
            } else {
                var v27 uint = 114
                fmt.Println(v18, v18, v1)
                _ = v27
            }
        }
        v28 := GenS2{F0: "\n\t"}
        _ = v24
        _ = v28
    } else {
        switch "foo bar" {
        case "foo bar":
            var v29 float64 = 3.14159
            _ = v29
        case "a":
            var v30 int32 = 0
            _ = v30
        case "\n\t":
            switch 9223372036854775807 {
            case -825:
                var v31 bool = v3
                _ = v31
            case 127:
                v32 := []int{424}
                _ = v32
            default:
                v3 = v3
            }
        default:
            v18 = 3.14159
        }
        v33 := GenS3{F0: v21, F1: "\n\t"}
        _ = v33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if (uint32(v6) == (v4 & 0)) {
            if ((((("a" + "foo bar") + ("a" + "a")) + "\n\t") + "hello") >= ((((("\n\t" + "foo bar") + "") + ("foo bar" + (("" + "a") + (("" + "\n\t") + "hello")))) + "hello") + "a")) {
                v35 := GenS4{F0: 906}
                _ = v35
            }
            if (v6 >= uint(v4)) {
                for fi36 := 0; fi36 < 17; fi36++ {
                    v6 += v6
                    _ = len(v5)
                    if false {
                        if true {
                            v37 := GenS1{F0: (v6 >> ((18446744073709551615) & 63)), F1: 32767}
                            delete(v2, "hello")
                            if (v3 || ((^(v19)) != (v19 + 2))) {
                                v38 := GenS2{F0: ""}
                                var v39 float64 = (-2.25 + -2.25)
                                _ = v38
                                _ = v39
                            } else {
                                switch (fi36 << (((v6 << ((v6) & 63))) & 63)) {
                                case -235:
                                    fmt.Println(v6, v19)
                                }
                            }
                            v3 = v3
                            v18 = v18
                            _ = v37
                        }
                    } else {
                        var v40 int16 = v1
                        v41 := []int{}
                        _ = v40
                        _ = v41
                    }
                    _ = fi36
                }
                for fi42 := 0; fi42 < 11; fi42++ {
                    v3 = v3
                    v18 += v18
                    _ = fi42
                }
            } else {
                fmt.Println(v6, v1, v6)
                v18 -= v18
            }
        } else {
            v43 := []int{467, fi34}
            _ = v43
        }
        _ = fi34
    }
    v44 := GenS5{F0: 2147483647}
    for fi45 := 0; fi45 < 15; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 uint16 = 1
            _ = fi46
            _ = v47
        }
        var v48 bool = (2 == 128)
        _ = fi45
        _ = v48
    }
    v18 += v18
    switch ("\n\t" + "a") {
    case "\n\t":
        v4 = uint32(v6)
    case "foo bar":
        delete(v2, ("" + ""))
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
