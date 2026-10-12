package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0() uint8 {
    if ((688 < 42) && true) {
        v5 := GenS1{F0: -34}
        _ = v5
    }
    if false {
        var v6 uint8 = 127
        for fi7 := 0; fi7 < 7; fi7++ {
            v6 = v6
            v8 := []int{(fi7 | fi7), 510, (^(fi7))}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 string = ""
    v10 := []int{32, 150}
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int32, p2 uint64) int {
    if false {
        v11 := []int{int(p1), -300, 128, 0}
        var v12 bool = false
        switch (("" + "a") + "hello") {
        case "":
            for fi13 := 0; fi13 < 4; fi13++ {
                switch "a" {
                case "":
                    p1 -= int32(p0)
                case "a":
                    fmt.Println(v12)
                default:
                    if false {
                        v14 := GenS1{F0: -110}
                        for fi15 := 0; fi15 < 19; fi15++ {
                            switch 2147483647 {
                            case 4294967295:
                                if true {
                                    var v16 int16 = (p0 & (-(p0)))
                                    _ = v16
                                }
                            case 9223372036854775807:
                                if (p0 == p0) {
                                    var v17 string = ("" + "foo bar")
                                    v18 := GenS2{F0: (0 != 0)}
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 float64 = 0.0
                                    if (630 <= 2147483648) {
                                        if (4294967295 > (p2 + (^(p2)))) {
                                            p0 = -253
                                            p1 = 2147483647
                                        }
                                    }
                                    _ = v19
                                }
                            case -743:
                                var v20 int32 = 0
                                _ = v20
                            }
                            p0 = -625
                            _ = fi15
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v14
                    }
                }
                var v21 int16 = (p0 % (((p0 | p0)) | 1))
                _ = fi13
                _ = v21
            }
        default:
            if (!true) {
                if (21 != 100) {
                    _ = len(v11)
                }
            } else {
                fmt.Println(p1, p0)
                v12 = (v12 || v12)
                var v22 uint64 = p2
                _ = v22
            }
        }
        _ = v11
        _ = v12
    } else {
        p0 -= (p0 - p0)
        v23 := GenS2{F0: true}
        _ = v23
    }
    p1 = -2147483648
    for fi24 := 0; fi24 < 16; fi24++ {
        if (true || (530 > 18446744073709551615)) {
            v25 := []int{}
            _ = v25
        }
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    v3 := []string{"", ("\n\t" + "\n\t")}
    v4 := []uint32{}
    switch "hello" {
    case "hello":
        v2 += 0
    }
    if (42 < 2147483647) {
        v26 := GenS1{F0: 127}
        var v27 int32 = 0
        if (879 <= 42) {
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 bool = (!false)
                _ = fi28
                _ = v29
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                v2 -= (^(v2))
            }
        }
        _ = v26
        _ = v27
    }
    v2 += v2
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        if (519 == 86) {
            for fi30 := 0; fi30 < 5; fi30++ {
                if false {
                    v31 := GenS3{F0: 632, F1: uint64(v2)}
                    v2 = v2
                    _ = v31
                } else {
                    v2 += (v2 | (v2 - v2))
                    v3 = append(v3, "a")
                }
                fmt.Println(v2, fi30)
                if true {
                    v2 = 42
                    var v32 uint32 = 1
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    _ = v32
                }
                _ = fi30
            }
            switch "foo bar" {
            case "foo bar":
                var v33 uint64 = 18446744073709551615
                _ = v33
            default:
                v2 += 65535
            }
        }
    case "a":
        var v34 int = -339
        _ = v34
    default:
        if false {
            var v35 uint16 = 2
            var v36 int64 = 545
            for fi37 := 0; fi37 < 8; fi37++ {
                _ = len(v3)
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 uint32 = 65535
                    _ = fi38
                    _ = v39
                }
                v40 := GenS4{F0: v36, F1: (^(v2))}
                _ = fi37
                _ = v40
            }
            _ = v35
            _ = v36
        }
    }
    switch ((("foo bar" + (("\n\t" + ("hello" + ((("\n\t" + "a") + "") + ""))) + "a")) + "\n\t") + "\n\t") {
    case "a":
        v41 := GenS5{F0: 1, F1: 9223372036854775807}
        _ = v41
    case "foo bar":
        var v42 int16 = -1
        _ = v42
    case "\n\t":
        v2 = 65535
    }
    v43 := GenS2{F0: true}
    v2 = v2
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
}
