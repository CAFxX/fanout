package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int8; F1 int64; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int16; F2 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint; F1 string; F2 uint64 }
func genFn0() uint32 {
    var v5 int32 = -753
    var v6 string = "\n\t"
    var v7 float32 = 0.0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint8 {
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v8, fi9)
        for fi10 := 0; fi10 < 1; fi10++ {
            if (int8(v8) <= 61) {
                var v11 string = "a"
                var v12 int32 = -2147483648
                _ = v11
                _ = v12
            } else {
                var v13 int64 = 486
                v14 := GenS1{F0: false, F1: 1}
                _ = v13
                _ = v14
            }
            switch "\n\t" {
            case "foo bar":
                var v15 uint = 1
                _ = v15
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int16 = -32768
        _ = fi16
        _ = v17
    }
    v18 := []int{32767}
    _ = v18
    _ = v8
    return 255
}

func main() {
    v1 := []int{912}
    var v2 int8 = 100
    var v3 uint64 = 42
    var v4 uint32 = 396
    v19 := GenS1{F0: true, F1: 68}
    switch 32767 {
    case -661:
        var v20 int = 476
        _ = v20
    }
    v2 = v2
    switch "" {
    case "a":
        v4 -= (v4 % (((^(v4))) | 1))
    }
    var v21 float32 = (3.14159 * -2.25)
    if true {
        var v22 uint = 817
        v2 += (v2 | 127)
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS2{F0: 65535}
            if false {
                v2 += (v2 | -128)
                v4 = (^(v4))
                if ((fi23 >= (fi23 / (((fi23 % (((^(fi23))) | 1))) | 1))) && (!(((fi23 ^ (fi23 * (fi23 >> (((v22 + 1)) & 63)))) <= fi23) || ((v4 - (v4 >> (((v22 ^ 1)) & 63))) == (v4 & 42))))) {
                    switch "a" {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch "" {
                    case "\n\t":
                        v22 -= v22
                    case "hello":
                        var v25 uint32 = v4
                        _ = v25
                    }
                    v26 := GenS3{F0: 65535, F1: (((("\n\t" + "a") + "hello") + "a") + ((("\n\t" + "foo bar") + "") + "hello"))}
                    _ = v26
                }
            } else {
                if false {
                    var v27 int = 495
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v29 := GenS4{F0: -32, F1: 291, F2: (fi23 | (fi28 | fi23))}
                        _ = fi28
                        _ = v29
                    }
                    var v30 int = 42
                    _ = v27
                    _ = v30
                }
                var v31 uint64 = (v3 ^ 4294967295)
                _ = v31
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v3 -= 167
    switch ("foo bar" + "") {
    case "foo bar":
        v32 := GenS5{F0: -41}
        _ = v32
    default:
        for fi33 := 0; fi33 < 16; fi33++ {
            v34 := GenS6{F0: 128, F1: 327, F2: 0}
            v2 += 127
            switch (^(fi33)) {
            case -366:
                for fi35 := 0; fi35 < 14; fi35++ {
                    fmt.Println(v3, fi35, fi33)
                    v2 -= v2
                    _ = fi35
                }
            }
            _ = fi33
            _ = v34
        }
    }
    var v36 uint16 = 934
    _ = len(v1)
    var v37 uint16 = 2
    v1 = append(v1, 32768)
    var v38 uint64 = (v3 % ((v3) | 1))
    var v39 uint = 2
    fmt.Println(v4, v39)
    if false {
        v3 -= 1
    } else {
        v2 = (v2 % ((v2) | 1))
        if (((v2 < 42) && false) && (v2 == (v2 ^ v2))) {
            for fi40 := 0; fi40 < 2; fi40++ {
                v41 := GenS7{F0: (^(v36))}
                var v42 float64 = (0.1 + (-2.25 - (3.14159 - (1e10 * 1.5))))
                var v43 int64 = 42
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
            v44 := GenS8{F0: 1, F1: "foo bar"}
            v2 = v2
            _ = v44
        } else {
            v45 := GenS7{F0: (v36 ^ (v37 - v37))}
            var v46 string = ("\n\t" + "\n\t")
            _ = v45
            _ = v46
        }
    }
    if false {
        v47 := GenS9{F0: v39, F1: "", F2: 4294967295}
        _ = v47
    } else {
        fmt.Println(v37)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
