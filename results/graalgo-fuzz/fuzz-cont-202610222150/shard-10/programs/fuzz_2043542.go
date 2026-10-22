package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 int) uint {
    p0 += p0
    v4 := GenS1{F0: 18446744073709551615, F1: (false && true)}
    _ = p0
    _ = v4
    return 2
}
func genFn1(p0 uint16, p1 uint32) int32 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 70, F1: (255 != 0)}
        _ = fi5
        _ = v6
    }
    switch "\n\t" {
    case "hello":
        v7 := []int{9223372036854775807, 4294967295, 1}
        _ = v7
    case "foo bar":
        if (190 != 154) {
            var v8 uint8 = 128
            _ = v8
        }
    default:
        p1 += (p1 % ((p1) | 1))
    }
    var v9 int32 = 1
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    var v1 uint = 255
    var v2 int16 = 1000
    var v3 float32 = 1.5
    switch ("\n\t" + ("\n\t" + "a")) {
    case "hello":
        fmt.Println(v1)
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = 2
        v12 := GenS2{F0: fi10, F1: 42}
        _ = fi10
        _ = v11
        _ = v12
    }
    if (false || false) {
        var v13 uint32 = 228
        _ = v13
    }
    var v14 int = 1000
    var v15 uint8 = 1
    var v16 uint16 = 397
    v15 = (^(v15))
    if (24 < 18446744073709551615) {
        for fi17 := 0; fi17 < 8; fi17++ {
            switch "hello" {
            case "":
                var v18 float64 = 3.14159
                _ = v18
            case "hello":
                switch "\n\t" {
                case "\n\t":
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v20 := GenS1{F0: 18446744073709551615, F1: false}
                        fmt.Println(v14, fi19, v14)
                        var v21 int = 32767
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    switch ((("\n\t" + "") + ("foo bar" + "")) + (("hello" + "") + "\n\t")) {
                    case "foo bar":
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    default:
                        v23 := []int{(v14 / (((fi17 * 1)) | 1)), v14}
                        _ = v23
                    }
                }
            case "foo bar":
                v1 = (v1 % (((v1 % ((539) | 1))) | 1))
            }
            if false {
                if false {
                    switch "a" {
                    case "\n\t":
                        if ((v14 * (v14 + v14)) != fi17) {
                            v24 := []int{fi17, 42}
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            _ = v24
                        } else {
                            v25 := GenS1{F0: 65535, F1: true}
                            var v26 uint64 = 0
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        v1 = 611
                    case "":
                        v27 := GenS3{F0: ("" >= (("hello" + "a") + "\n\t"))}
                        _ = v27
                    default:
                        v28 := GenS1{F0: v1, F1: (100 <= 1)}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v14 += -187
                        _ = fi29
                    }
                    v15 = 113
                } else {
                    v16 += 1
                }
            } else {
                var v30 uint16 = (v16 - (v16 | 0))
                _ = v30
            }
            _ = fi17
        }
        var v31 int = v14
        _ = v31
    }
    v32 := GenS4{F0: -128}
    var v33 int32 = 21
    var v34 uint32 = 0
    if true {
        for fi35 := 0; fi35 < 13; fi35++ {
            v16 -= 525
            _ = fi35
        }
    }
    v36 := []int{(-(v14)), v14}
    var v37 int = v14
    var v38 int = (v37 - 42)
    var v39 uint = v1
    v33 -= v33
    v40 := GenS1{F0: uint(v16), F1: (!(320 != (v1 - (v39 + 255))))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
