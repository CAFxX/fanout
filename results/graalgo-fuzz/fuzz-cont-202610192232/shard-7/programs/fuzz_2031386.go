package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 int8 }
func genFn0(p0 string) int32 {
    v5 := []int{}
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 18; fi7++ {
        switch (p0 + "") {
        case "\n\t":
            v8 := GenS1{F0: "hello"}
            _ = v8
        case "foo bar":
            p0 = (p0 + (p0 + (((("hello" + "\n\t") + p0) + p0) + p0)))
        case "a":
            switch fi7 {
            case -9223372036854775808:
                p0 = ("a" + ("a" + ("" + ("" + p0))))
            case -805:
                p0 = "\n\t"
            }
        }
        v9 := GenS2{F0: "hello"}
        for fi10 := 0; fi10 < 12; fi10++ {
            v5 = append(v5, (fi7 - (fi10 | -151)))
            _ = len(v5)
            if (9223372036854775807 < -382) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (!false) {
                    p0 = "hello"
                    p0 = ((p0 + (p0 + ("hello" + "hello"))) + p0)
                } else {
                    v11 := GenS2{F0: p0}
                    _ = v11
                }
                switch (fi10 % ((fi10) | 1)) {
                case -9223372036854775808:
                    var v12 float32 = 1e10
                    _ = v12
                case 42:
                    p0 = "\n\t"
                default:
                    fmt.Println(fi7, fi10, p0)
                }
                if false {
                    v5 = append(v5, (^(fi7)))
                }
            } else {
                fmt.Println(p0, fi7)
                p0 = (p0 + ("\n\t" + p0))
            }
            _ = fi10
        }
        _ = fi7
        _ = v9
    }
    v13 := GenS1{F0: p0}
    _ = p0
    _ = v13
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []int{}
    v2 := []int{-558, 1}
    v3 := []uint32{0, 483}
    var v4 bool = true
    switch "\n\t" {
    case "foo bar":
        v14 := GenS1{F0: ((("\n\t" + ("a" + "\n\t")) + "\n\t") + (("hello" + "a") + ""))}
        _ = v14
    case "a":
        var v15 int32 = genFn0("")
        _ = v15
    }
    fmt.Println(v4)
    switch "hello" {
    case "a":
        v2 = append(v2, 9223372036854775807)
    case "\n\t":
        var v16 uint8 = 255
        _ = v16
    case "foo bar":
        switch ("" + "\n\t") {
        case "foo bar":
            _ = len(v2)
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            if true {
                switch ((("a" + "a") + "foo bar") + "foo bar") {
                case "":
                    switch "hello" {
                    case "\n\t":
                        v4 = (!v4)
                    case "hello":
                        v4 = (42 >= 1)
                    }
                case "a":
                    var v19 bool = (((!v4) && true) || false)
                    v19 = v19
                    var v20 uint16 = 1
                    _ = v19
                    _ = v20
                case "\n\t":
                    fmt.Println(fi17, fi17)
                default:
                    if false {
                        v4 = false
                        v21 := GenS3{F0: -1, F1: 42}
                        _ = v21
                    }
                }
                var v22 int16 = 32767
                v1 = append(v1, fi18)
                _ = v22
            } else {
                if ((65535 < 848) || (-32768 < 740)) {
                    if (2 != 255) {
                        var v23 bool = ((0 == 2147483647) || v4)
                        _ = v23
                    }
                    if v4 {
                        var v24 int = (fi17 / (((fi17 >> ((1) & 63))) | 1))
                        v25 := GenS4{F0: 42, F1: 1}
                        _ = v24
                        _ = v25
                    } else {
                        if true {
                            var v26 bool = v4
                            var v27 float64 = 0.0
                            var v28 string = "foo bar"
                            v27 += v27
                            if v26 {
                                v27 -= (3.14159 + (float64((1e10 - (float32((3.14159 / float64(0.1))) * ((float32(float64(float32((v27 / 1e10)))) * (3.14159 - 3.14159)) + (0.0 + 1.5))))) - 3.14159))
                                v29 := GenS1{F0: ((v28 + "a") + ("\n\t" + (v28 + "a")))}
                                v4 = (0 <= -1)
                                _ = v29
                            } else {
                                if (2147483647 <= 745) {
                                    v30 := GenS5{F0: -128, F1: 194}
                                    if v26 {
                                        var v31 uint16 = 904
                                        var v32 uint = 0
                                        v31 -= v31
                                        _ = v31
                                        _ = v32
                                    }
                                    v26 = (-347 > 1)
                                    var v33 uint16 = 255
                                    _ = v30
                                    _ = v33
                                } else {
                                    var v34 uint32 = 1
                                    _ = v34
                                }
                                v26 = (-89 == -1)
                                if false {
                                    v35 := GenS6{F0: 32767}
                                    _ = v35
                                }
                            }
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    var v36 int16 = int16(fi18)
                    _ = v36
                }
                v4 = false
            }
            var v37 uint64 = 2
            v38 := GenS7{F0: 18446744073709551615}
            _ = fi18
            _ = v37
            _ = v38
        }
        switch 2147483647 {
        case 9223372036854775807:
            var v39 uint64 = 143
            _ = v39
        case -712:
            switch "" {
            case "\n\t":
                v4 = true
            case "hello":
                v40 := GenS8{F0: 18446744073709551615, F1: 1}
                _ = v40
            case "foo bar":
                if v4 {
                    fmt.Println(v4, v4)
                    var v41 uint16 = 214
                    var v42 int16 = -382
                    _ = v41
                    _ = v42
                }
            }
        default:
            v4 = v4
        }
        v43 := GenS2{F0: "\n\t"}
        _ = fi17
        _ = v43
    }
    var v44 int8 = 42
    fmt.Println(v44, v4)
    v45 := GenS7{F0: 910}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
