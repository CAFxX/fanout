package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 uint16, p1 int8) int {
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            p1 -= p1
            _ = fi7
        }
        if (!false) {
            var v8 float64 = ((float64(-2.25) / 0.1) - float64((float32((1e10 * 0.0)) / 3.14159)))
            var v9 float64 = 0.0
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 string = ("foo bar" + (("a" + (((("" + "") + ("a" + "a")) + ("\n\t" + "a")) + "a")) + "a"))
    _ = p0
    _ = p1
    _ = v10
    return 0
}
func genFn1() uint16 {
    v11 := GenS1{F0: -1, F1: 653, F2: 1}
    v12 := GenS1{F0: 40, F1: 42, F2: 241}
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := []int{(fi13 + 2147483647), (fi13 * (-(fi13))), (fi13 & -9223372036854775808)}
        switch 50 {
        case -178:
            if false {
                if (true || (true && (((0 <= 331) || (18446744073709551615 <= 134)) && (false || false)))) {
                    v15 := GenS1{F0: 1, F1: 4, F2: 0}
                    _ = v15
                } else {
                    switch ("\n\t" + ("hello" + "\n\t")) {
                    case "":
                        var v16 float32 = -2.25
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        switch (fi13 ^ fi13) {
                        case 1:
                            var v18 string = (("hello" + ("hello" + "\n\t")) + ("hello" + ("foo bar" + ("foo bar" + "a"))))
                            switch (fi17 / ((0) | 1)) {
                            case -1:
                                var v19 int8 = 100
                                _ = v19
                            default:
                                v18 = (("hello" + ("hello" + "a")) + v18)
                            }
                            var v20 float32 = 3.14159
                            _ = v18
                            _ = v20
                        case -203:
                            switch "foo bar" {
                            case "":
                                var v21 uint64 = 4294967295
                                _ = v21
                            case "\n\t":
                                if (133 == 85) {
                                    fmt.Println(fi17)
                                    if len(v14) > 0 {
                                        v14[0] = 703
                                    }
                                } else {
                                    v14 = append(v14, (fi13 ^ 6))
                                }
                            }
                        }
                        fmt.Println(fi13)
                        var v22 string = "foo bar"
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = fi17
                        _ = v22
                    }
                    if true {
                        var v23 int16 = -1
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v23
                    } else {
                        var v24 int8 = 65
                        _ = v24
                    }
                }
            } else {
                var v25 uint8 = 0
                switch "a" {
                case "foo bar":
                    if false {
                        v26 := GenS2{F0: 0}
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            _ = fi27
                        }
                        for fi28 := 0; fi28 < 8; fi28++ {
                            var v29 int64 = 42
                            var v30 float64 = 1.5
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = v26
                    } else {
                        var v31 int16 = -32768
                        fmt.Println(v25, fi13, v31)
                        var v32 int16 = (v31 / (((-(v31))) | 1))
                        _ = v31
                        _ = v32
                    }
                case "\n\t":
                    var v33 uint32 = 2
                    _ = v33
                default:
                    var v34 uint64 = 42
                    _ = v34
                }
                v25 += (v25 | (v25 - v25))
                _ = v25
            }
        case -409:
            var v35 float32 = (float32(1.5) / 1.0)
            _ = v35
        }
        _ = fi13
        _ = v14
    }
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := []string{("\n\t" + ("foo bar" + ("foo bar" + ""))), "foo bar", (("a" + "\n\t") + ("" + "a")), ("\n\t" + "a")}
    v3 := make(map[int]bool)
    v4 := []int{388, 32768}
    var v5 uint8 = (v1 & (v1 ^ v1))
    var v36 uint32 = 2147483648
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v37 := v3[684]
    var v38 int16 = 42
    v39 := GenS1{F0: -2147483648, F1: 510, F2: -179}
    var v40 int32 = 0
    if v37 {
        var v41 uint = 2
        _ = v41
    }
    switch -169 {
    case -948:
        var v42 int16 = (v38 << ((uint(v36)) & 63))
        _ = v42
    }
    switch "\n\t" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        v40 += (v40 * (v40 + v40))
        var v44 uint64 = 1
        v45 := GenS3{F0: (false && false), F1: 42}
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
