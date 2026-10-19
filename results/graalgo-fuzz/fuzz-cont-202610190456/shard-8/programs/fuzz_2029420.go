package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int32; F2 bool }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0(p0 uint32, p1 uint8) int16 {
    var v5 uint8 = 1
    var v6 int64 = 42
    if (!true) {
        v5 = p1
    } else {
        v7 := []int{32767, 57, 1, -2147483648}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -32768
}
func genFn1(p0 uint32) uint64 {
    var v8 int64 = 9223372036854775807
    switch 650 {
    case 32768:
        v8 += (v8 - v8)
    }
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := []int{555, 9223372036854775807, 128, -9223372036854775808}
    v3 := []string{("" + "hello"), "", "\n\t", ""}
    v4 := []uint32{}
    if false {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                switch (^(fi10)) {
                case 2147483647:
                    var v11 int16 = 578
                    _ = v11
                case 206:
                    v1 = (v1 - 255)
                }
                _ = fi10
            }
            _ = fi9
        }
        v12 := GenS1{F0: 100, F1: 0, F2: 42}
        _ = v12
    }
    if ((!true) && false) {
        for fi13 := 0; fi13 < 1; fi13++ {
            v1 -= 255
            v1 -= v1
            _ = fi13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS2{F0: 1}
                switch (fi15 + -1) {
                case 42:
                    var v17 int16 = 1
                    _ = v17
                default:
                    v18 := GenS1{F0: -1, F1: -740, F2: -1}
                    _ = v18
                }
                v2 = append(v2, fi14)
                v3 = append(v3, "\n\t")
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v19 := GenS3{F0: "foo bar"}
    var v20 int32 = -849
    for fi21 := 0; fi21 < 6; fi21++ {
        v20 = 0
        var v22 int8 = int8(fi21)
        _ = fi21
        _ = v22
    }
    var v23 uint8 = v1
    for fi24 := 0; fi24 < 2; fi24++ {
        v3 = append(v3, "a")
        _ = fi24
    }
    var v25 int64 = 213
    v26 := GenS4{F0: 1000}
    for fi27 := 0; fi27 < 18; fi27++ {
        if false {
            var v28 int64 = (-(v25))
            fmt.Println(v23)
            v3 = append(v3, "\n\t")
            _ = v28
        }
        var v29 uint16 = 255
        _ = fi27
        _ = v29
    }
    fmt.Println(v1, v23)
    v23 += v1
    var v30 int8 = 127
    fmt.Println(v25)
    switch 32768 {
    case 32768:
        if true {
            v25 = (v25 - (^(v25)))
            if false {
                v20 += int32(v30)
                v30 = -1
            }
        }
    case -356:
        var v31 float32 = -2.25
        _ = v31
    case 838:
        for fi32 := 0; fi32 < 12; fi32++ {
            if false {
                v33 := GenS5{F0: v1, F1: v20, F2: true}
                _ = v33
            } else {
                v30 = int8(fi32)
                for fi34 := 0; fi34 < 20; fi34++ {
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "")
                    }
                    _ = fi34
                }
                fmt.Println(v20, v30, v25)
            }
            switch "a" {
            case "foo bar":
                if ((true && true) || ("foo bar" < ("\n\t" + ""))) {
                    if false {
                        _ = len(v3)
                        if true {
                            var v35 float64 = 0.0
                            if (true || (!((v20 != v20) && (v20 < (v20 >> ((4294967295) & 63)))))) {
                                v30 = v30
                                var v36 int16 = genFn0(2147483647, v23)
                                v1 = (v1 - v23)
                                _ = v36
                            }
                            _ = v35
                        }
                    }
                }
            default:
                var v37 string = (("foo bar" + ("\n\t" + "foo bar")) + "hello")
                _ = v37
            }
            _ = fi32
        }
    }
    switch "foo bar" {
    case "":
        var v38 int64 = v25
        _ = v38
    case "\n\t":
        v3 = append(v3, "foo bar")
    default:
        var v39 string = (("a" + "a") + "hello")
        _ = v39
    }
    v40 := GenS6{F0: false, F1: (v30 >> ((4294967295) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v40
}
