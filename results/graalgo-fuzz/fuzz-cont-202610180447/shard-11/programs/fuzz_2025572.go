package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0(p0 string, p1 int32) uint64 {
    if true {
        switch p0 {
        case "":
            fmt.Println(p1, p1)
        case "hello":
            switch p0 {
            case "\n\t":
                if false {
                    for fi5 := 0; fi5 < 3; fi5++ {
                        v6 := []int{fi5, fi5}
                        fmt.Println(fi5)
                        _ = fi5
                        _ = v6
                    }
                } else {
                    p1 = (p1 + p1)
                }
            default:
                switch ((((p0 + p0) + "\n\t") + ("\n\t" + p0)) + "") {
                case "a":
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 uint64 = 4294967295
                        v9 := GenS1{F0: v8, F1: 18446744073709551615, F2: 4294967295}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                case "hello":
                    var v10 uint = 42
                    _ = v10
                default:
                    fmt.Println(p1, p0)
                }
            }
        case "\n\t":
            fmt.Println(p1, p1)
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int, p1 uint64) int8 {
    for fi11 := 0; fi11 < 9; fi11++ {
        switch 572 {
        case 32768:
            var v12 int32 = -2147483648
            _ = v12
        }
        _ = fi11
    }
    fmt.Println(p0, p1)
    var v13 float32 = 0.1
    v14 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -89
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{42, 2147483647, 0}
    var v3 float64 = 1e10
    var v4 int16 = -1
    fmt.Println(v4, v4)
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            v16 := GenS3{F0: 1}
            _ = v16
        } else {
            if false {
                var v17 uint16 = 65535
                _ = v17
            }
            v1 = (^(v1))
            var v18 uint8 = (v1 | 180)
            _ = v18
        }
        _ = fi15
    }
    var v19 uint8 = v1
    v3 = v3
    switch 0 {
    case -524:
        switch -187 {
        case 9223372036854775807:
            fmt.Println(v19)
        case -511:
            for fi20 := 0; fi20 < 7; fi20++ {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                var v21 uint64 = 1
                _ = fi20
                _ = v21
            }
        case -956:
            v1 += (v19 ^ v1)
        }
    default:
        fmt.Println(v1, v1, v4)
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        fmt.Println(v1, v4)
        var v25 int32 = 971
        _ = fi24
        _ = v25
    }
    fmt.Println(v1)
    var v26 string = "a"
    for fi27 := 0; fi27 < 10; fi27++ {
        switch "foo bar" {
        case "\n\t":
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 uint = 42
                fmt.Println(v29, v3)
                _ = fi28
                _ = v29
            }
        case "":
            if true {
                var v30 float64 = 0.0
                if (128 >= 0) {
                    var v31 string = v26
                    _ = v31
                } else {
                    v32 := GenS1{F0: 4294967295, F1: 800, F2: 728}
                    var v33 uint = uint(v19)
                    _ = v32
                    _ = v33
                }
                _ = v30
            }
        case "foo bar":
            fmt.Println(v19, v4)
        }
        if len(v2) > 0 {
            v2[0] = 2
        }
        var v34 float32 = 3.14159
        _ = fi27
        _ = v34
    }
    if (-603 > 1) {
        v26 = (v26 + v26)
        v3 += (-2.25 / v3)
        if ((false && ((27 > -106) && false)) && true) {
            v35 := GenS4{F0: 182, F1: 1}
            var v36 uint64 = 42
            _ = v35
            _ = v36
        } else {
            fmt.Println(v26, v19, v4)
        }
    } else {
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 uint8 = 128
            var v39 uint64 = 1
            var v40 uint64 = 893
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
