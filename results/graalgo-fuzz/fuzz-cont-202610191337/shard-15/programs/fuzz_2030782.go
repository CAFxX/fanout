package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 int, p1 bool, p2 uint32) int32 {
    p1 = (!false)
    v4 := GenS1{F0: -2147483648, F1: false, F2: 169}
    switch ("" + "") {
    case "":
        v5 := []int{(p0 * (p0 >> ((18446744073709551615) & 63))), -789, 127, p0}
        _ = v5
    case "hello":
        fmt.Println(p0, p0)
    case "foo bar":
        v6 := GenS2{F0: 42, F1: -345}
        _ = v6
    default:
        var v7 int32 = 2147483647
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}
func genFn1(p0 int64, p1 bool, p2 int16) int64 {
    var v8 bool = p1
    var v9 int64 = 0
    v10 := []int{233, 32768}
    for fi11 := 0; fi11 < 11; fi11++ {
        _ = len(v10)
        p2 -= p2
        p2 -= 32767
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (v9 << ((18446744073709551615) & 63))
}

func main() {
    v1 := []int{565, 1}
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 943
    switch "hello" {
    case "foo bar":
        switch 128 {
        case 32768:
            var v12 int16 = -1
            _ = v12
        case 127:
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "a":
                v13 := GenS1{F0: -1, F1: true, F2: 18446744073709551615}
                _ = v13
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    case "a":
        if (("a" >= "a") || (v2 > 405)) {
            fmt.Println(v2, v3)
            var v14 int8 = 42
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
        }
    default:
        if true {
            if ((!(25 < 18446744073709551615)) || false) {
                v3 += (v3 * (v3 - v3))
            } else {
                v2 = v2
                v15 := GenS2{F0: 2, F1: -1}
                _ = v15
            }
            v3 = (v3 & (v3 << ((810) & 63)))
            if len(v1) > 0 {
                v1[0] = -65
            }
        }
    }
    v2 = v2
    for fi16 := 0; fi16 < 11; fi16++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v17 uint = 669
        v2 += v2
        _ = fi16
        _ = v17
    }
    switch 2147483647 {
    case -35:
        var v18 int64 = -1
        _ = v18
    }
    v2 = genFn1((v2 & (v2 * v2)), true, 42)
    fmt.Println(v3, v3)
    var v19 float32 = (float32(float64(float32(1.5))) * 3.14159)
    for fi20 := 0; fi20 < 6; fi20++ {
        v3 -= 491
        v1 = append(v1, (fi20 - fi20))
        _ = fi20
    }
    v3 -= v3
    for fi21 := 0; fi21 < 6; fi21++ {
        v2 = v2
        var v22 int16 = -1
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint = 0
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if false {
        if len(v1) > 0 {
            v1[0] = 579
        }
    } else {
        v26 := GenS3{F0: -1, F1: 42}
        for fi27 := 0; fi27 < 5; fi27++ {
            v19 += float32(((1e10 * 0.1) + 1e10))
            switch fi27 {
            case -1:
                _ = len(v1)
            case 925:
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint = 0
                    v30 := GenS1{F0: 0, F1: (!(2 <= 4294967295)), F2: 4294967295}
                    switch fi28 {
                    case 32768:
                        v31 := GenS2{F0: 18446744073709551615, F1: 1}
                        _ = v31
                    case -648:
                        var v32 bool = true
                        _ = v32
                    default:
                        v33 := GenS4{F0: ("a" + (("\n\t" + "foo bar") + ""))}
                        _ = v33
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            default:
                v34 := GenS5{F0: 0, F1: "hello"}
                _ = v34
            }
            v3 = (^(v3))
            _ = fi27
        }
        _ = v26
    }
    switch "hello" {
    case "foo bar":
        if (true && false) {
            v2 -= (^(v2))
        } else {
            fmt.Println(v19, v2)
            if true {
                if true {
                    v19 = (0.1 - -2.25)
                    v1 = append(v1, 1)
                    fmt.Println(v2, v2)
                }
                v2 = v2
            }
        }
    case "\n\t":
        v2 = (v2 % ((v2) | 1))
    case "":
        fmt.Println(v2, v3)
    default:
        fmt.Println(v2)
    }
    v3 += v3
    var v35 string = "\n\t"
    var v36 uint64 = 473
    var v37 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
