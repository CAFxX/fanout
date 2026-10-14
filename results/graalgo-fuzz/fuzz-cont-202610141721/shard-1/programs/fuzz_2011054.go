package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int16; F1 int8; F2 string }
func genFn0(p0 int8) uint16 {
    p0 -= (p0 | (p0 ^ -1))
    switch -123 {
    case -966:
        fmt.Println(p0)
    }
    switch "a" {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        var v6 string = ("hello" + "hello")
        _ = v6
    case "":
        if false {
            var v7 uint32 = 65535
            _ = v7
        } else {
            if ((42 > 1) && true) {
                v8 := []int{32768, -804, 9223372036854775807, 4294967295}
                _ = v8
            }
        }
    default:
        if true {
            switch "hello" {
            case "":
                var v9 bool = ((false || (42 < 1)) || ("\n\t" > (("\n\t" + (("" + ("a" + ("a" + "foo bar"))) + "hello")) + "hello")))
                _ = v9
            default:
                v10 := []int{}
                _ = v10
            }
            if false {
                var v11 string = "foo bar"
                var v12 int8 = p0
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: -686, F1: false}
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        switch (fi15 & fi15) {
                        case -968:
                            fmt.Println(p0, p0)
                        }
                        v16 := GenS2{F0: 65535, F1: 100, F2: 171}
                        var v17 uint8 = 231
                        fmt.Println(fi15)
                        p0 = -128
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS2{F0: 65535, F1: p0, F2: 144}
                    var v19 int = (fi14 * fi14)
                    _ = fi14
                    _ = v18
                    _ = v19
                }
                switch ("\n\t" + "foo bar") {
                case "hello":
                    p0 -= p0
                }
                _ = v13
            }
            var v20 uint = 450
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(fi21)
                p0 = (p0 ^ (p0 - p0))
                _ = fi21
            }
        }
    }
    if true {
        p0 = (p0 ^ 100)
        var v22 int16 = 42
        _ = v22
    } else {
        v23 := GenS3{F0: ("" + ("hello" + "")), F1: 0, F2: 1}
        _ = v23
    }
    var v24 int8 = 30
    _ = p0
    _ = v24
    return 361
}
func genFn1(p0 int64, p1 int8) uint16 {
    p0 = (-(p0))
    var v25 uint32 = 2147483648
    var v26 string = "\n\t"
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (18446744073709551615 >= 18446744073709551615)
    v3 := []int{-652, 615}
    var v4 uint64 = 4294967295
    var v5 float64 = (0.0 / 1e10)
    v27 := GenS3{F0: "a", F1: 2, F2: 997}
    var v28 int = 32768
    v29 := GenS4{F0: 127, F1: 32767}
    v2 = false
    v2 = v2
    if (!false) {
        v30 := GenS3{F0: "foo bar", F1: 127, F2: 42}
        _ = v30
    } else {
        var v31 float64 = v5
        _ = len(v1)
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 3; fi33++ {
                v34 := GenS1{F0: 42, F1: (!(!true))}
                v2 = ((v28 & 2147483647) <= (fi32 | (fi32 * fi33)))
                var v35 uint32 = 0
                var v36 int = 32768
                v35 -= (v35 / ((v35) | 1))
                fmt.Println(v31, v4)
                var v37 int16 = 32767
                v38 := GenS5{F0: (v37 % (((v37 ^ 1000)) | 1)), F1: 1, F2: ("foo bar" + "foo bar")}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi32
        }
        _ = v31
    }
    _ = len(v3)
    v3 = append(v3, -1)
    v4 -= 830
    var v39 uint = 724
    var v40 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
