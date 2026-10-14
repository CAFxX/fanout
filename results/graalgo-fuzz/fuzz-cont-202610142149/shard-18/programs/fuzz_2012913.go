package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 bool) int16 {
    v5 := []int{128, 4294967295}
    if (!true) {
        if len(v5) > 0 {
            v5[0] = 32767
        }
        var v6 int8 = -1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v6
    }
    var v7 uint = 81
    _ = p0
    _ = v5
    _ = v7
    return -307
}
func genFn1(p0 int16, p1 int, p2 uint64) string {
    switch "foo bar" {
    case "foo bar":
        var v8 int64 = 679
        _ = v8
    case "\n\t":
        v9 := []int{(p1 ^ p1), p1, 32767, p1}
        _ = v9
    case "a":
        v10 := []int{(p1 >> ((255) & 63)), -154, (p1 | -800)}
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint16 = 42
        var v13 bool = (p2 > 18446744073709551615)
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + ((((("" + "a") + "foo bar") + ((("" + ("" + "\n\t")) + "\n\t") + (("\n\t" + ("a" + "foo bar")) + ("a" + "")))) + "a") + "foo bar"))
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{42, 9223372036854775807, 32768, 1}
    var v3 uint = 65535
    v4 := make(map[uint32]uint32)
    delete(v4, 766)
    v14 := GenS1{F0: "a", F1: 42}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 int64 = int64(fi15)
            v2 = append(v2, 0)
            if len(v2) > 0 {
                v2[0] = -537
            }
            _ = fi16
            _ = v17
        }
        var v18 uint16 = 1
        v19 := v4[2147483647]
        _ = fi15
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v3 = (^(v3))
            switch v1 {
            case "":
                switch v1 {
                case "a":
                    var v22 bool = true
                    _ = v22
                case "foo bar":
                    if (-32768 > 276) {
                        if false {
                            var v23 int8 = 29
                            fmt.Println(v3, fi21, v3)
                            _ = len(v2)
                            _ = v23
                        }
                        v24 := GenS2{F0: "\n\t", F1: 126, F2: 197}
                        _ = v24
                    }
                }
            case "\n\t":
                v25 := GenS2{F0: (v1 + "foo bar"), F1: 2, F2: 255}
                _ = v25
            default:
                fmt.Println(v3)
            }
            switch ((("" + (v1 + (v1 + v1))) + ("hello" + v1)) + (("hello" + ("" + ("foo bar" + (("foo bar" + "\n\t") + v1)))) + (((("a" + "") + "foo bar") + v1) + v1))) {
            case "\n\t":
                v26 := GenS1{F0: v1, F1: 2147483647}
                _ = v26
            case "foo bar":
                var v27 bool = (100 == -96)
                if (v27 && (4294967295 == 1)) {
                    v28 := GenS1{F0: v1, F1: 588}
                    _ = v28
                }
                _ = v27
            default:
                switch fi21 {
                case -60:
                    if false {
                        var v29 uint = v3
                        fmt.Println(v1, v1, v29)
                        v29 = v29
                        _ = v29
                    } else {
                        v30 := GenS3{F0: (!(2 > 854)), F1: 1000, F2: ((v1 + v1) + v1)}
                        var v31 uint8 = 255
                        _ = v30
                        _ = v31
                    }
                case 127:
                    var v32 bool = (!true)
                    v33 := GenS4{F0: genFn0((284 != 255)), F1: 1}
                    _ = v32
                    _ = v33
                case 4294967295:
                    v3 += (v3 - (v3 - (v3 * (v3 + (v3 - v3)))))
                }
            }
            fmt.Println(v3)
            if (!false) {
                switch "\n\t" {
                case "hello":
                    var v34 int8 = 42
                    _ = v34
                case "\n\t":
                    if (32767 != 1000) {
                        v35 := GenS1{F0: v1, F1: 0}
                        _ = v35
                    }
                }
                var v36 int16 = 1000
                _ = v36
            }
            _ = fi21
        }
        if true {
            v3 = 0
        } else {
            v1 = "a"
            v1 = v1
        }
        _ = fi20
    }
    if false {
        var v37 int16 = 1000
        v3 -= v3
        for fi38 := 0; fi38 < 8; fi38++ {
            if len(v2) > 0 {
                v2[0] = (fi38 + -69)
            }
            v37 = v37
            v4[651] = 2
            _ = fi38
        }
        _ = v37
    }
    v39 := GenS1{F0: v1, F1: 42}
    v40 := GenS2{F0: ((v1 + v1) + v1), F1: 42, F2: 54}
    v41 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
