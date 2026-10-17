package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    var v4 string = ("\n\t" + ("a" + "foo bar"))
    v5 := []int{42}
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 bool) int8 {
    if p0 {
        switch "foo bar" {
        case "\n\t":
            v6 := []int{42, -9223372036854775808, 1}
            _ = v6
        case "foo bar":
            var v7 string = ("\n\t" + "")
            _ = v7
        }
        var v8 uint16 = 0
        var v9 int64 = -191
        _ = v8
        _ = v9
    } else {
        if p0 {
            switch 128 {
            case 9223372036854775807:
                var v10 int16 = 0
                _ = v10
            case 32767:
                var v11 int32 = 42
                _ = v11
            }
        } else {
            var v12 int8 = -73
            _ = v12
        }
        p0 = (-941 >= -1)
    }
    var v13 int = 198
    _ = p0
    _ = v13
    return 113
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{4294967295}
    v3 := make(map[int]string)
    switch -1 {
    case 32768:
        var v14 float32 = (0.1 / float32((((3.14159 * -2.25) * 1e10) - 0.0)))
        _ = v14
    case 4294967295:
        if (968 < 333) {
            var v15 int32 = -2147483648
            switch ("a" + "") {
            case "hello":
                v3[189] = "hello"
            case "a":
                switch 1 {
                case -188:
                    if true {
                        var v16 float64 = (3.14159 * (3.14159 + 1e10))
                        _ = v16
                    } else {
                        v15 = (v15 ^ (-(v15)))
                        fmt.Println(v15)
                        fmt.Println(v15)
                    }
                default:
                    v15 = (v15 ^ v15)
                }
            case "foo bar":
                if false {
                    var v17 uint64 = 254
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS1{F0: 1, F1: (!true), F2: 42}
                        var v20 bool = (("foo bar" + ("" + "")) != ((("" + "foo bar") + ((("a" + ("\n\t" + "\n\t")) + "\n\t") + ("foo bar" + ("hello" + "hello")))) + ("a" + ("hello" + ((("hello" + ("a" + "\n\t")) + "hello") + "\n\t")))))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 14; fi21++ {
                        fmt.Println(fi21, fi21)
                        _ = fi21
                    }
                    if true {
                        v1[851] = true
                    }
                }
            }
            _ = v15
        } else {
            var v22 int8 = genFn1((true && false))
            if (65535 > 42) {
                v22 -= v22
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v23 int16 = 562
                v23 += v23
                _ = v23
            }
            if (0 < 65535) {
                var v24 int32 = 2147483647
                var v25 uint64 = 605
                _ = v24
                _ = v25
            } else {
                v22 += genFn1((!(2 != 255)))
            }
            _ = v22
        }
    case 0:
        var v26 uint8 = 224
        _ = v26
    }
    var v27 int64 = 144
    v27 += v27
    v28 := GenS2{F0: "hello"}
    for fi29 := 0; fi29 < 1; fi29++ {
        v2 = append(v2, 620)
        fmt.Println(v27, v27)
        var v30 bool = ("foo bar" > "hello")
        _ = fi29
        _ = v30
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v27)
    } else {
        v1[1] = ((0.1 + (3.14159 - 1e10)) == 0.0)
    }
    if false {
        var v31 uint32 = 2147483648
        _ = v31
    }
    fmt.Println(v27)
    v32 := GenS2{F0: ""}
    var v33 uint32 = 745
    if ((float64((1e10 * 0.0)) <= (-2.25 + 3.14159)) && (0 > 127)) {
        if ((2 == 218) || (42 <= -2147483648)) {
            if (32767 <= uint16(v33)) {
                v27 = v27
                switch "a" {
                case "foo bar":
                    v33 = v33
                }
                v33 = 2147483648
            }
            var v34 int8 = 81
            _ = v34
        } else {
            v35 := GenS3{F0: 1}
            v36 := v3[-2147483648]
            if len(v2) > 0 {
                v2[0] = 65535
            }
            _ = v35
            _ = v36
        }
        if (true && (((3.14159 + float64(float32(3.14159))) - float64(1e10)) == 1.5)) {
            fmt.Println(v27, v33)
            delete(v3, 32768)
        } else {
            if (genFn0() < 42) {
                switch -989 {
                case 128:
                    var v37 float32 = (1e10 + 1e10)
                    _ = v37
                case 663:
                    v27 = 42
                }
            }
            var v38 uint64 = 18446744073709551615
            _ = v38
        }
    } else {
        _ = len(v2)
    }
    if false {
        var v39 uint64 = 2
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
