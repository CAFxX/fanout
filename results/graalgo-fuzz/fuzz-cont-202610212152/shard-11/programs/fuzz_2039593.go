package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint64 }
func genFn0(p0 bool) string {
    switch 9223372036854775807 {
    case 0:
        switch 127 {
        case 32767:
            p0 = true
        case 4294967295:
            p0 = (255 > 11)
        case 32768:
            var v6 uint8 = 157
            _ = v6
        default:
            var v7 uint = 255
            _ = v7
        }
    }
    p0 = p0
    p0 = p0
    if p0 {
        v8 := GenS1{F0: -32768, F1: -1, F2: ((((("\n\t" + "hello") + "foo bar") + "hello") + (("foo bar" + "\n\t") + "hello")) + "foo bar")}
        p0 = ("\n\t" > "a")
        var v9 float64 = 3.14159
        _ = v8
        _ = v9
    }
    if (true && (!true)) {
        if true {
            v10 := []int{32768}
            v11 := GenS2{F0: 128, F1: 394, F2: (("foo bar" + ("foo bar" + "\n\t")) + "hello")}
            _ = v10
            _ = v11
        } else {
            v12 := GenS3{F0: 255}
            _ = v12
        }
        v13 := []int{67, 847}
        var v14 float64 = float64(float32(1.5))
        _ = v13
        _ = v14
    } else {
        p0 = (1 < 57)
    }
    _ = p0
    return "foo bar"
}
func genFn1() int {
    v15 := []int{127, -2147483648, 969}
    v15 = append(v15, -2147483648)
    _ = v15
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 2
    var v5 uint32 = v1
    var v16 uint32 = v1
    if false {
        for fi17 := 0; fi17 < 12; fi17++ {
            v5 -= v5
            _ = fi17
        }
    }
    v18 := GenS1{F0: 1000, F1: 169, F2: v2}
    v19 := []int{4294967295, genFn1()}
    switch (v2 + "a") {
    case "a":
        if ("\n\t" < "\n\t") {
            if (0 <= 4294967295) {
                fmt.Println(v16)
            }
        }
    case "":
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint8 = 128
            var v22 uint = 42
            var v23 int8 = 100
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    case "foo bar":
        fmt.Println(v2, v2, v2)
    }
    v24 := GenS3{F0: 2}
    var v25 int64 = 9223372036854775807
    for fi26 := 0; fi26 < 19; fi26++ {
        if (!(1 <= 679)) {
            fmt.Println(v25, v16, v16)
            v27 := GenS1{F0: -396, F1: int64(v4), F2: ""}
            _ = v27
        }
        _ = fi26
    }
    v1 = (^(v16))
    for fi28 := 0; fi28 < 6; fi28++ {
        v5 = v16
        delete(v3, v5)
        _ = len(v3)
        _ = fi28
    }
    v29 := GenS4{F0: 32767, F1: 83}
    if true {
        if len(v19) > 0 {
            v19[0] = 42
        }
    } else {
        for fi30 := 0; fi30 < 8; fi30++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = fi30
        }
        fmt.Println(v25)
        var v31 string = (v2 + "foo bar")
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v16 = (v5 + (^(v16)))
        v33 := GenS5{F0: 1000}
        switch "\n\t" {
        case "\n\t":
            switch fi32 {
            case 958:
                delete(v3, uint32(v4))
            case 127:
                v5 = v5
            default:
                if false {
                    v34 := GenS5{F0: 42}
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v36 := GenS6{F0: v2, F1: uint64(v1)}
                        if true {
                            fmt.Println(v5, v25)
                            if true {
                                v5 -= 2147483647
                                fmt.Println(fi35)
                            } else {
                                if (true || true) {
                                    switch (^(fi35)) {
                                    case 127:
                                        delete(v3, 2147483648)
                                    case 42:
                                        v1 -= 4294967295
                                    }
                                    v3[749] = v1
                                }
                            }
                        } else {
                            var v37 bool = true
                            _ = v37
                        }
                        if true {
                            v38 := GenS3{F0: 255}
                            var v39 uint64 = (v4 - 90)
                            _ = v38
                            _ = v39
                        } else {
                            var v40 uint64 = (v4 % ((1) | 1))
                            v41 := GenS6{F0: v2, F1: v4}
                            _ = v40
                            _ = v41
                        }
                        _ = fi35
                        _ = v36
                    }
                    _ = len(v19)
                    _ = v34
                }
            }
        }
        _ = fi32
        _ = v33
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        _ = len(v19)
        _ = len(v19)
        _ = fi42
    }
    var v43 uint16 = 1
    for fi44 := 0; fi44 < 15; fi44++ {
        v16 = 42
        if (true || true) {
            v4 = (v4 & (v4 + v4))
        }
        _ = fi44
    }
    v45 := GenS4{F0: (v43 << ((255) & 63)), F1: -23}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v5
}
