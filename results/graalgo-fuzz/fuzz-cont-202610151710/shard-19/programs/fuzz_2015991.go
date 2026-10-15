package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int64; F2 int16 }

func main() {
    var v1 uint8 = 101
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 uint64 = 2
    var v5 string = "foo bar"
    switch 1 {
    case 42:
        if false {
            if (v5 > (("" + v5) + "foo bar")) {
                v6 := GenS1{F0: 1}
                _ = v6
            } else {
                v5 = "\n\t"
            }
            if false {
                fmt.Println(v3, v5)
            }
            fmt.Println(v4, v5, v5)
        } else {
            var v7 uint8 = (v1 << ((65535) & 63))
            _ = v7
        }
    case 2147483647:
        var v8 float32 = 3.14159
        _ = v8
    case -9223372036854775808:
        v9 := GenS1{F0: 1}
        _ = v9
    }
    var v10 bool = (!(!(((0.0 + float64(1.5)) != 0.1) && false)))
    var v11 bool = v10
    if (v5 == ("foo bar" + v5)) {
        v5 = v5
        if v11 {
            v12 := GenS1{F0: -128}
            v10 = true
            switch ("foo bar" + ("a" + v5)) {
            case "":
                v5 = "hello"
            }
            _ = v12
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := []int{(-(fi13)), (fi14 % (((fi13 ^ (fi14 * (fi13 * (fi13 << ((65535) & 63)))))) | 1))}
                    _ = fi14
                    _ = v15
                }
                if v11 {
                    v16 := GenS1{F0: 42}
                    if false {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            fmt.Println(v10, v4)
                            _ = fi17
                        }
                        var v18 uint64 = (v4 >> ((2) & 63))
                        if v10 {
                            if (v10 || (2147483647 == -1)) {
                                var v19 int32 = -494
                                v10 = true
                                if (v1 >= v1) {
                                    for fi20 := 0; fi20 < 20; fi20++ {
                                        if v11 {
                                            v21 := v2[(fi20 + fi20)]
                                            _ = v21
                                        }
                                        if (int(v3) >= 0) {
                                            switch v5 {
                                            case "hello":
                                                v22 := GenS1{F0: -1}
                                                _ = v22
                                            case "\n\t":
                                                if v10 {
                                                    var v23 string = ((((v5 + v5) + v5) + "a") + v5)
                                                    v24 := GenS2{F0: (-(fi20)), F1: v11}
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                        } else {
                                            var v25 int16 = (v3 << ((544) & 63))
                                            fmt.Println(v25, v4)
                                            _ = v25
                                        }
                                        _ = fi20
                                    }
                                    v10 = v11
                                    var v26 uint = 2
                                    _ = v26
                                }
                                _ = v19
                            } else {
                                fmt.Println(v3)
                                if v10 {
                                    if (42 > (^(v4))) {
                                        fmt.Println(fi13, fi13, v18)
                                        var v27 float32 = 3.14159
                                        v28 := []int{(fi13 | (fi13 - 4294967295)), fi13, 127}
                                        _ = v27
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: fi13, F1: (((^(v1)) <= (v1 | v1)) || v10)}
                                    _ = v29
                                } else {
                                    v30 := GenS3{F0: true, F1: (^(v3)), F2: 768}
                                    if (v11 || v11) {
                                        v31 := GenS3{F0: true, F1: (v3 | v3), F2: 271}
                                        var v32 uint8 = (v1 * v1)
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v30
                                }
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    v18 = (v18 + (v4 | v4))
                                    _ = fi33
                                }
                            }
                        } else {
                            v11 = true
                        }
                        _ = v18
                    } else {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            var v35 uint8 = 128
                            _ = fi34
                            _ = v35
                        }
                        v36 := GenS4{F0: 474}
                        _ = v36
                    }
                    v1 = 1
                    _ = v16
                }
                v37 := []int{(fi13 ^ (fi13 * 42)), (fi13 / (((fi13 | fi13)) | 1)), int(v4), fi13}
                _ = fi13
                _ = v37
            }
            _ = len(v2)
            v1 -= 170
        }
    } else {
        v4 = 18446744073709551615
    }
    var v38 int32 = 2147483647
    v39 := []int{}
    fmt.Println(v5, v3, v3)
    if (978 >= v3) {
        var v40 uint8 = v1
        v41 := GenS1{F0: -107}
        fmt.Println(v10, v38)
        _ = v40
        _ = v41
    }
    _ = len(v39)
    switch v5 {
    case "hello":
        v4 = (v4 & (v4 & v4))
    case "":
        if v10 {
            for fi42 := 0; fi42 < 20; fi42++ {
                fmt.Println(v11)
                var v43 int64 = -1
                _ = fi42
                _ = v43
            }
            switch "\n\t" {
            case "hello":
                var v44 string = v5
                _ = v44
            case "a":
                if v10 {
                    if (!v11) {
                        var v45 uint = 52
                        fmt.Println(v3)
                        for fi46 := 0; fi46 < 11; fi46++ {
                            fmt.Println(v10)
                            _ = fi46
                        }
                        _ = v45
                    }
                }
            case "foo bar":
                v4 += v4
            default:
                var v47 uint8 = v1
                _ = v47
            }
        }
    }
    var v48 float64 = (1e10 * 0.1)
    for fi49 := 0; fi49 < 6; fi49++ {
        _ = len(v39)
        v50 := GenS5{F0: v1, F1: 1, F2: v3}
        _ = fi49
        _ = v50
    }
    if true {
        var v51 int32 = v38
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v5
}
