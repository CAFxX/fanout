package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 bool; F1 uint8; F2 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16; F1 string; F2 uint32 }
type GenS6 struct { F0 bool; F1 int64; F2 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 int16) uint16 {
    p0 -= p0
    var v5 float64 = 3.14159
    var v6 int64 = 1
    p0 += (^(p0))
    if (764 == 640) {
        var v7 uint64 = 512
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            p1 = 894
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 uint32) uint {
    var v9 uint16 = 2
    var v10 string = "\n\t"
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return uint(p0)
}

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    var v3 int64 = -1
    var v4 int8 = 1
    for fi11 := 0; fi11 < 2; fi11++ {
        v1 = append(v1, -772)
        v12 := GenS1{F0: 9223372036854775807, F1: -541, F2: "hello"}
        if (9223372036854775807 == v3) {
            v4 -= v4
            v13 := GenS2{F0: 1, F1: "a"}
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: 179, F1: ""}
    var v15 uint32 = 2
    v16 := GenS3{F0: false, F1: uint8(v15), F2: 0}
    for fi17 := 0; fi17 < 11; fi17++ {
        switch (fi17 & fi17) {
        case 9223372036854775807:
            for fi18 := 0; fi18 < 13; fi18++ {
                _ = len(v1)
                fmt.Println(v2, v2, fi17)
                _ = fi18
            }
        case 127:
            v19 := GenS3{F0: false, F1: 255, F2: 255}
            _ = v19
        case 4294967295:
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS1{F0: (fi20 & -2147483648), F1: int16(fi17), F2: "foo bar"}
                _ = fi20
                _ = v21
            }
        }
        if len(v1) > 0 {
            v1[0] = fi17
        }
        switch "foo bar" {
        case "foo bar":
            _ = len(v1)
        case "\n\t":
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := GenS2{F0: 4294967295, F1: "hello"}
                fmt.Println(fi17, v2)
                if (v3 != -9223372036854775808) {
                    if len(v1) > 0 {
                        v1[0] = fi17
                    }
                    switch fi22 {
                    case -55:
                        var v24 string = "foo bar"
                        _ = v24
                    }
                }
                v25 := GenS4{F0: 1, F1: 127}
                v26 := GenS5{F0: 13, F1: (("a" + "\n\t") + "hello"), F2: 42}
                _ = fi22
                _ = v23
                _ = v25
                _ = v26
            }
        case "":
            if true {
                switch ("\n\t" + "a") {
                case "\n\t":
                    switch "\n\t" {
                    case "":
                        v27 := GenS6{F0: false, F1: v3, F2: int8(v3)}
                        _ = v27
                    }
                default:
                    v28 := GenS7{F0: (false || (0 != -533))}
                    _ = v28
                }
                if false {
                    fmt.Println(v2, fi17)
                    if true {
                        if (4294967295 >= 4294967295) {
                            var v29 string = "foo bar"
                            _ = v29
                        }
                        _ = len(v1)
                    } else {
                        var v30 uint = 1
                        fmt.Println(fi17)
                        v31 := GenS8{F0: 851, F1: 4294967295}
                        _ = v30
                        _ = v31
                    }
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    if false {
                        v4 = v4
                    }
                    v4 = -32
                    _ = fi32
                }
            }
        }
        _ = fi17
    }
    v4 = v4
    var v33 int64 = v3
    var v34 bool = (4294967295 >= v15)
    var v35 uint8 = 128
    var v36 float32 = (0.0 + 1.5)
    for fi37 := 0; fi37 < 20; fi37++ {
        _ = len(v1)
        var v38 int16 = 32767
        _ = fi37
        _ = v38
    }
    v2 -= v2
    fmt.Println(v2)
    if (!((int8(v3) > int8(v33)) && v34)) {
        v39 := GenS9{F0: v33, F1: 255, F2: 0}
        _ = len(v1)
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
