package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 int32; F2 string }
type GenS4 struct { F0 int8 }
func genFn0(p0 string, p1 int16) string {
    if false {
        p0 = p0
        v5 := []int{49, 32768}
        _ = v5
    } else {
        var v6 int64 = 0
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{fi7}
            switch p0 {
            case "hello":
                v8 = append(v8, 9223372036854775807)
            case "":
                var v9 uint = 509
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: p0, F1: (v6 >> ((255) & 63)), F2: (!true)}
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v11 uint8 = 1
    p1 = p1
    v12 := []int{9223372036854775807, -2147483648, -1}
    var v13 int64 = int64(v11)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return p0
}
func genFn1(p0 uint16, p1 int, p2 uint16) uint8 {
    var v14 string = "\n\t"
    v15 := GenS1{F0: v14, F1: -811, F2: false}
    var v16 uint8 = 255
    var v17 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return v16
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 55
    var v3 int8 = -1
    var v4 int = 824
    fmt.Println(v1, v4, v2)
    v3 += (v3 % ((42) | 1))
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch "hello" {
        case "foo bar":
            var v18 uint8 = genFn1(1, (v4 * (v4 * v4)), 312)
            _ = v18
        case "":
            var v19 bool = true
            _ = v19
        }
    case 127:
        var v20 uint8 = 2
        _ = v20
    }
    if false {
        v21 := GenS1{F0: "\n\t", F1: 638, F2: ((!true) || true)}
        if false {
            if false {
                v4 = (v4 / ((-9223372036854775808) | 1))
            }
            v4 += -561
            switch ("a" + ("" + (("\n\t" + "foo bar") + ("a" + ("a" + ("a" + "hello")))))) {
            case "\n\t":
                if false {
                    fmt.Println(v3)
                } else {
                    fmt.Println(v3, v4, v1)
                }
            }
        } else {
            for fi22 := 0; fi22 < 5; fi22++ {
                v2 = (v2 / ((v2) | 1))
                v23 := []int{v4, 2147483647, v4, -9223372036854775808}
                var v24 float64 = -2.25
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        var v25 int16 = 42
        _ = v21
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint = v2
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS1{F0: (("a" + "") + ("" + ("foo bar" + ""))), F1: 42, F2: true}
            if (true && true) {
                switch v4 {
                case -477:
                    if true {
                        v30 := []int{}
                        _ = v30
                    } else {
                        if false {
                            v27 = 65535
                            var v31 string = ("\n\t" + ("" + ("a" + ("" + ("" + "a")))))
                            _ = v31
                        }
                        fmt.Println(v1)
                    }
                case 0:
                    var v32 float32 = float32(0.1)
                    _ = v32
                }
                switch "a" {
                case "\n\t":
                    var v33 int16 = 834
                    _ = v33
                default:
                    var v34 string = ""
                    _ = v34
                }
                if true {
                    var v35 int = fi28
                    v36 := GenS1{F0: "foo bar", F1: 0, F2: ((v2 ^ 675) < (v27 << (((v2 & (v27 >> ((255) & 63)))) & 63)))}
                    fmt.Println(v1)
                    _ = v35
                    _ = v36
                }
            }
            v2 -= 0
            _ = fi28
            _ = v29
        }
        v4 = fi26
        _ = fi26
        _ = v27
    }
    v4 = v4
    v2 = v2
    switch (("hello" + "\n\t") + "hello") {
    case "\n\t":
        var v37 uint = v2
        _ = v37
    case "hello":
        for fi38 := 0; fi38 < 5; fi38++ {
            if true {
                v39 := GenS2{F0: -1, F1: (v4 % (((fi38 << ((v2) & 63))) | 1))}
                for fi40 := 0; fi40 < 10; fi40++ {
                    fmt.Println(fi38)
                    v1 += (v1 & (v1 * v1))
                    _ = fi40
                }
                _ = v39
            } else {
                switch "\n\t" {
                case "foo bar":
                    if (-2.25 != 0.1) {
                        fmt.Println(v2, v1)
                    }
                case "hello":
                    var v41 int16 = 42
                    _ = v41
                default:
                    if (v3 > v3) {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            fmt.Println(v3, v2, v3)
                            var v43 float32 = (0.1 - (0.0 / (1e10 - float32(3.14159))))
                            fmt.Println(v4, v43, v1)
                            var v44 int16 = 769
                            var v45 float64 = (0.0 / (3.14159 + (1e10 / (0.1 - 1.5))))
                            _ = fi42
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        if (!false) {
                            switch ("foo bar" + "a") {
                            case "a":
                                if false {
                                    var v46 int8 = 127
                                    v47 := []int{v4}
                                    for fi48 := 0; fi48 < 11; fi48++ {
                                        v2 -= 255
                                        v4 = (fi48 + fi38)
                                        var v49 uint8 = 42
                                        v4 = 127
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v46
                                    _ = v47
                                }
                            case "hello":
                                var v50 bool = (0 < 65535)
                                _ = v50
                            }
                            v2 -= (v2 >> (((v2 & (v2 ^ (v2 + 547)))) & 63))
                        }
                        v3 -= (v3 % (((v3 % (((v3 - (v3 & (-(v3))))) | 1))) | 1))
                    } else {
                        switch -2147483648 {
                        case 42:
                            v51 := GenS3{F0: (fi38 >> (((v2 / ((v2) | 1))) & 63)), F1: -1, F2: ("" + "")}
                            _ = v51
                        case 207:
                            fmt.Println(v4)
                        case -203:
                            fmt.Println(fi38)
                        }
                        var v52 int = 42
                        v53 := []int{fi38, 42, fi38}
                        _ = v52
                        _ = v53
                    }
                }
                var v54 uint64 = 0
                var v55 int = -275
                _ = v54
                _ = v55
            }
            v3 -= (v3 >> (((v2 ^ v2)) & 63))
            _ = fi38
        }
    case "a":
        v4 += (v4 / ((-202) | 1))
    }
    v3 -= (v3 / ((v3) | 1))
    if false {
        v4 = v4
        var v56 int16 = int16(v3)
        v57 := []int{(v4 >> ((v2) & 63)), -9223372036854775808, -2147483648, v4}
        _ = v56
        _ = v57
    } else {
        v58 := []int{-9223372036854775808}
        v1 += v1
        _ = v58
    }
    v59 := GenS4{F0: -1}
    v4 = (v4 / (((-(v4))) | 1))
    v3 = v3
    v2 -= 932
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
}
