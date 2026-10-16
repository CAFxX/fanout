package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0() int64 {
    var v5 uint16 = 80
    switch "a" {
    case "\n\t":
        for fi6 := 0; fi6 < 12; fi6++ {
            v5 -= v5
            var v7 int = (fi6 & 778)
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v8 bool = true
        _ = v8
    case "hello":
        switch ("hello" + "hello") {
        case "foo bar":
            var v9 int64 = 9223372036854775807
            _ = v9
        case "a":
            switch "" {
            case "\n\t":
                fmt.Println(v5)
            case "foo bar":
                fmt.Println(v5)
            default:
                var v10 int16 = 1
                _ = v10
            }
        case "":
            fmt.Println(v5)
        default:
            fmt.Println(v5)
        }
    }
    v11 := []int{-518, 127, 32768}
    var v12 bool = true
    for fi13 := 0; fi13 < 5; fi13++ {
        if ((float32(0.0) <= 3.14159) || v12) {
            if (-229 < 385) {
                v5 += v5
                var v14 uint8 = 0
                _ = v14
            } else {
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch (("hello" + "foo bar") + ("hello" + "foo bar")) {
                    case "":
                        if v12 {
                            switch fi13 {
                            case 32768:
                                if v12 {
                                    var v16 float32 = 3.14159
                                    _ = v16
                                }
                            case -1:
                                v5 = (v5 | v5)
                            case 2147483647:
                                v5 = (v5 >> ((590) & 63))
                            }
                        } else {
                            if v12 {
                                v11 = append(v11, fi15)
                                switch "" {
                                case "a":
                                    if (2147483647 <= 4294967295) {
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        if (!(207 <= -1)) {
                                            if true {
                                                if len(v11) > 0 {
                                                    v11[0] = (fi13 * (fi15 & 32768))
                                                }
                                                _ = len(v11)
                                                v17 := GenS1{F0: 18446744073709551615, F1: 117, F2: (v5 << ((255) & 63))}
                                                var v18 int16 = -32768
                                                _ = v17
                                                _ = v18
                                            } else {
                                                var v19 uint16 = (^(v5))
                                                _ = v19
                                            }
                                            var v20 bool = true
                                            _ = v20
                                        } else {
                                            v12 = false
                                        }
                                    } else {
                                        v12 = (42 >= 18446744073709551615)
                                        v21 := GenS1{F0: 2, F1: int8(fi15), F2: v5}
                                        v12 = (false || true)
                                        _ = v21
                                    }
                                    v12 = v12
                                }
                            }
                            var v22 string = ("hello" + "foo bar")
                            v22 = ""
                            var v23 int = 0
                            v12 = v12
                            _ = v22
                            _ = v23
                        }
                    }
                    var v24 int16 = 42
                    _ = fi15
                    _ = v24
                }
                v25 := GenS2{F0: -9223372036854775808}
                var v26 float32 = 0.1
                _ = v25
                _ = v26
            }
            if len(v11) > 0 {
                v11[0] = 737
            }
        }
        _ = fi13
    }
    _ = v11
    _ = v12
    _ = v5
    return 700
}
func genFn1() bool {
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS3{F0: 589}
        _ = fi27
        _ = v28
    }
    var v29 float64 = 0.0
    _ = v29
    return ((v29 <= 1.5) && true)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -259
    var v3 int8 = 1
    var v4 int32 = (v2 / ((v2) | 1))
    if (610 > 494) {
        v4 -= v2
    }
    switch 0 {
    case 127:
        var v30 uint16 = 32767
        _ = v30
    case 103:
        v1[42] = false
    default:
        switch -551 {
        case 32768:
            switch "a" {
            case "":
                v4 -= (-(v2))
            case "\n\t":
                var v31 float32 = 1.5
                _ = v31
            default:
                v3 += v3
            }
        case -548:
            v4 -= v2
        }
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4)
    default:
        v32 := v1[2147483647]
        _ = v32
    }
    var v33 int = -1
    v34 := []int{}
    v35 := v1[2147483648]
    for fi36 := 0; fi36 < 15; fi36++ {
        for fi37 := 0; fi37 < 16; fi37++ {
            v33 -= v33
            _ = fi37
        }
        if (true || true) {
            if false {
                if (fi36 < fi36) {
                    v35 = v35
                    fmt.Println(v33, v35, v2)
                }
            }
        } else {
            var v38 int32 = (v2 - (v2 & int32(v33)))
            _ = v38
        }
        _ = fi36
    }
    v34 = append(v34, v33)
    v39 := GenS4{F0: 509, F1: 0}
    switch "hello" {
    case "\n\t":
        var v40 float64 = 3.14159
        _ = v40
    default:
        var v41 uint32 = 65535
        _ = v41
    }
    var v42 float64 = 1.5
    if v35 {
        v2 = (v2 | 0)
    } else {
        for fi43 := 0; fi43 < 20; fi43++ {
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            var v44 float64 = v42
            fmt.Println(v2, v35, v33)
            _ = fi43
            _ = v44
        }
        var v45 int8 = (v3 / (((^(v3))) | 1))
        _ = v45
    }
    fmt.Println(v33, v3, v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v42
}
