package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint, p1 int, p2 bool) int8 {
    fmt.Println(p0)
    v7 := []int{p1, p1}
    v8 := GenS1{F0: -2147483648}
    p2 = (!true)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 bool, p1 uint8) bool {
    v9 := []int{321, 32767}
    v9 = append(v9, 485)
    fmt.Println(p0, p0)
    p1 += p1
    if (!(("a" + "") < "")) {
        var v10 uint = 4294967295
        _ = v10
    } else {
        var v11 uint8 = (p1 | p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return (uint32(p1) < 65535)
}

func main() {
    var v1 int32 = -42
    var v2 uint64 = 1
    v3 := []string{"a", (("\n\t" + ("a" + "")) + ("\n\t" + "a")), ("a" + ("" + "hello"))}
    var v4 uint = 0
    var v5 float64 = 3.14159
    v6 := make(map[string]uint32)
    switch 0 {
    case 128:
        if (v1 <= (^(v1))) {
            _ = len(v6)
        }
    default:
        v4 = v4
    }
    v4 = v4
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            if true {
                v13 := GenS1{F0: v1}
                _ = v13
            }
            var v14 bool = false
            _ = len(v6)
            _ = fi12
            _ = v14
        }
    } else {
        fmt.Println(v2, v5, v2)
    }
    v2 -= v2
    v1 -= 869
    var v15 int32 = (v1 + int32(v4))
    var v16 float64 = (((0.1 / v5) + v5) + v5)
    fmt.Println(v2)
    delete(v6, "foo bar")
    var v17 int16 = -32768
    var v18 int = 1
    v6["foo bar"] = 2147483648
    for fi19 := 0; fi19 < 5; fi19++ {
        if true {
            if false {
                _ = len(v3)
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch "foo bar" {
                    case "hello":
                        v1 = (v1 / (((v1 - v15)) | 1))
                    default:
                        if (int16(fi20) >= -1) {
                            if (0.1 > 3.14159) {
                                var v21 uint = 1
                                var v22 uint16 = uint16(v21)
                                _ = v21
                                _ = v22
                            }
                            v23 := GenS2{F0: -285}
                            _ = v23
                        }
                    }
                    fmt.Println(v16, fi19)
                    _ = fi20
                }
                var v24 int32 = v15
                delete(v6, "foo bar")
                _ = v24
            }
            v2 = v2
            switch ("foo bar" + ("foo bar" + (("\n\t" + "") + (("a" + ("foo bar" + ("" + (("" + ("\n\t" + (("a" + "a") + ("foo bar" + "a")))) + "\n\t")))) + "\n\t")))) {
            case "a":
                var v25 float64 = float64(1.5)
                _ = v25
            case "hello":
                switch (v18 / (((v18 >> ((v4) & 63))) | 1)) {
                case 32767:
                    var v26 uint8 = 255
                    _ = v26
                case 0:
                    if false {
                        switch 4294967295 {
                        case 1:
                            v27 := GenS1{F0: v15}
                            _ = v27
                        case -2147483648:
                            if false {
                                v17 = v17
                                v18 += 32768
                            } else {
                                v18 = (v18 * fi19)
                                if (v4 == (v4 << (((v4 & (v4 % ((v4) | 1)))) & 63))) {
                                    v28 := GenS1{F0: (v1 << (((v4 / (((v4 & (^(v4)))) | 1))) & 63))}
                                    var v29 int8 = 100
                                    _ = v28
                                    _ = v29
                                } else {
                                    v15 -= (^(v1))
                                }
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v1 -= v15
                                case "hello":
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        if true {
                                            v2 = 4294967295
                                        }
                                        switch ("hello" + "") {
                                        case "hello":
                                            var v31 uint16 = 648
                                            _ = v31
                                        case "":
                                            _ = len(v6)
                                        case "\n\t":
                                            v17 = v17
                                        default:
                                            v16 += 1e10
                                        }
                                        var v32 int32 = (v15 / ((v15) | 1))
                                        _ = fi30
                                        _ = v32
                                    }
                                case "a":
                                    v33 := GenS3{F0: 801, F1: (v5 <= v16)}
                                    _ = v33
                                }
                            }
                        default:
                            for fi34 := 0; fi34 < 5; fi34++ {
                                if ((v1 / (((v1 % ((1) | 1))) | 1)) > (^(v1))) {
                                    v15 = v1
                                }
                                switch (v18 / (((v18 / ((1) | 1))) | 1)) {
                                case 2147483647:
                                    var v35 uint64 = v2
                                    _ = v35
                                default:
                                    v1 = -543
                                }
                                v36 := GenS4{F0: v4, F1: (1 > v2), F2: v1}
                                switch fi19 {
                                case 32768:
                                    if true {
                                        v1 = 42
                                    } else {
                                        fmt.Println(v5, fi19)
                                    }
                                }
                                _ = fi34
                                _ = v36
                            }
                        }
                    }
                default:
                    fmt.Println(v18, v18, v2)
                }
            case "\n\t":
                v4 = 65535
            }
        } else {
            for fi37 := 0; fi37 < 6; fi37++ {
                switch "foo bar" {
                case "\n\t":
                    v38 := GenS5{F0: 127}
                    _ = v38
                case "a":
                    v39 := v6[(("a" + (("\n\t" + "a") + "foo bar")) + "\n\t")]
                    _ = v39
                case "foo bar":
                    v17 = 32767
                    var v40 int32 = v1
                    v15 += v15
                    _ = v40
                default:
                    v41 := GenS3{F0: 2, F1: false}
                    _ = v41
                }
                _ = fi37
            }
            v3 = append(v3, "hello")
        }
        _ = fi19
    }
    var v42 uint32 = 2
    var v43 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
