package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 int8; F2 bool }
type GenS4 struct { F0 int32; F1 string; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS8 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0(p0 uint) int8 {
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            p0 += 647
            v7 = v7
        }
        if ((1 == 1) && true) {
            var v9 int64 = -799
            var v10 int = (fi8 - fi8)
            if true {
                var v11 uint8 = 41
                _ = v11
            } else {
                v12 := GenS1{F0: 2, F1: 42}
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            v13 := []int{(fi8 % ((fi8) | 1)), (fi8 | 205)}
            _ = v13
        }
        if (201 >= 117) {
            v7 = (v7 ^ v7)
            if false {
                var v14 int64 = -18
                v15 := []int{(fi8 ^ 2147483647), (^(fi8)), (^(fi8)), fi8}
                if false {
                    v16 := GenS2{F0: "hello", F1: 1, F2: (224 != 104)}
                    _ = v16
                } else {
                    _ = len(v15)
                }
                _ = v14
                _ = v15
            } else {
                var v17 uint = (^(p0))
                var v18 string = "foo bar"
                _ = v17
                _ = v18
            }
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return -30
}
func genFn1() int32 {
    var v19 uint16 = 255
    v19 += v19
    for fi20 := 0; fi20 < 5; fi20++ {
        if (((true || true) || false) || false) {
            v19 -= (^(v19))
            v19 += 42
            var v21 uint64 = 481
            _ = v21
        } else {
            v22 := GenS3{F0: 42, F1: genFn0(194), F2: (55 < 0)}
            v23 := GenS1{F0: (v19 - v19), F1: 1000}
            v19 = v19
            _ = v22
            _ = v23
        }
        v24 := GenS4{F0: -2147483648, F1: "\n\t", F2: genFn0(255)}
        _ = fi20
        _ = v24
    }
    v25 := GenS5{F0: "hello"}
    _ = v19
    _ = v25
    return 216
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 uint8 = 255
    var v4 int8 = 100
    v5 := []int{4294967295, 42}
    var v6 int = 127
    v26 := GenS3{F0: v4, F1: v4, F2: (true && (true || ((3.14159 + -2.25) < float32(-2.25))))}
    fmt.Println(v3)
    var v27 uint16 = 1
    v28 := GenS3{F0: -89, F1: v4, F2: false}
    var v29 uint32 = 2147483647
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 -= 32768
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 uint8 = 2
        _ = fi30
        _ = v31
    }
    if (!((-2.25 > 3.14159) && true)) {
        v32 := GenS6{F0: (v4 < -128)}
        _ = v32
    } else {
        v33 := GenS2{F0: "hello", F1: (v3 - v3), F2: true}
        fmt.Println(v29, v4)
        if ((!(v3 > v3)) || (false && (0.1 < (-2.25 - float64((1.5 - float32(-2.25))))))) {
            if ((v29 == uint32(v3)) && false) {
                v29 = (v29 | 255)
            } else {
                v34 := GenS3{F0: 21, F1: v4, F2: false}
                v35 := GenS7{F0: v3, F1: 42, F2: (v4 - (v4 | int8(v29)))}
                _ = v34
                _ = v35
            }
            var v36 float64 = 1.5
            _ = v36
        }
        _ = v33
    }
    if (("\n\t" + ("foo bar" + "hello")) == "a") {
        v4 -= genFn0(1)
        var v37 string = (("foo bar" + "\n\t") + "hello")
        _ = v37
    } else {
        for fi38 := 0; fi38 < 6; fi38++ {
            v29 = 1
            fmt.Println(v29, v27)
            _ = fi38
        }
        switch v6 {
        case -1:
            _ = len(v5)
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 uint16 = 32767
            v6 = v6
            v41 := GenS8{F0: true, F1: v40, F2: -1}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    var v42 bool = (239 <= 148)
    v4 = (v4 + (v4 & (v4 + (v4 * (v4 + 1)))))
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 string = "\n\t"
        fmt.Println(v29, v44)
        if (!(true || (!v42))) {
            var v45 int32 = v2
            if v42 {
                for fi46 := 0; fi46 < 2; fi46++ {
                    delete(v1, v29)
                    switch v44 {
                    case "":
                        var v47 float64 = 3.14159
                        _ = v47
                    }
                    _ = fi46
                }
                v48 := v1[(v29 - (v29 - (v29 - v29)))]
                _ = v48
            }
            v27 = 255
            _ = v45
        }
        _ = fi43
        _ = v44
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
