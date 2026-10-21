package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint32 = 671
    v2 := []string{"a", ("a" + "")}
    var v3 string = "hello"
    v4 := []uint32{(v1 | 4294967295), (v1 / ((v1) | 1)), v1, v1}
    var v5 int32 = -1
    v6 := make(map[int]string)
    v7 := GenS1{F0: 0, F1: 94}
    v1 = (v1 >> ((459) & 63))
    _ = len(v2)
    var v8 int8 = -58
    v5 = v5
    v9 := GenS2{F0: (((v3 + v3) + v3) + "foo bar"), F1: 1, F2: v3}
    switch "a" {
    case "":
        _ = len(v6)
    default:
        var v10 bool = true
        _ = v10
    }
    var v11 uint8 = 127
    v8 = -9
    switch (v3 + ((v3 + "a") + "foo bar")) {
    case "\n\t":
        fmt.Println(v11)
    default:
        var v12 float64 = -2.25
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: "foo bar", F1: 888, F2: v3}
        _ = fi13
        _ = v14
    }
    fmt.Println(v5, v3, v5)
    if (((float64(0.1) + 3.14159) / (3.14159 + 3.14159)) > -2.25) {
        v5 += int32(v11)
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: 121, F1: (-(v8))}
            _ = fi15
            _ = v16
        }
        v8 += v8
    }
    switch -63 {
    case -1:
        if true {
            v17 := GenS1{F0: uint(v1), F1: v8}
            v18 := GenS3{F0: (1 > 2), F1: (v3 + (v3 + v3))}
            _ = v17
            _ = v18
        }
    case 0:
        if false {
            for fi19 := 0; fi19 < 9; fi19++ {
                for fi20 := 0; fi20 < 14; fi20++ {
                    if ((^(v8)) < v8) {
                        var v21 float32 = 1.5
                        v5 += v5
                        v5 -= (v5 << ((2) & 63))
                        _ = v21
                    }
                    if (false || (!false)) {
                        var v22 float32 = (1e10 * 1e10)
                        v3 = (v3 + v3)
                        v3 = ((v3 + v3) + (v3 + ("foo bar" + (v3 + v3))))
                        v23 := GenS3{F0: false, F1: "foo bar"}
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(fi19, v11, v11)
                        var v24 float32 = float32(1.5)
                        v8 = v8
                        _ = v24
                    }
                    v25 := v6[fi20]
                    _ = fi20
                    _ = v25
                }
                if true {
                    v26 := GenS2{F0: v3, F1: 264, F2: v3}
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                    _ = v26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    if ((^(v8)) != (v8 / ((v8) | 1))) {
                        v11 = v11
                        var v28 bool = (v5 == v5)
                        _ = v28
                    } else {
                        var v29 int32 = -972
                        if false {
                            var v30 bool = true
                            switch v3 {
                            case "foo bar":
                                switch ((v3 + v3) + "") {
                                case "foo bar":
                                    var v31 int8 = v8
                                    _ = v31
                                default:
                                    fmt.Println(v5)
                                }
                            case "\n\t":
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(fi19)
                                }
                            default:
                                var v32 bool = true
                                _ = v32
                            }
                            if ((0 < -1) && v30) {
                                v6[(fi27 * int(v11))] = ""
                                fmt.Println(v3, v11, v8)
                            }
                            fmt.Println(v1, fi19, v1)
                            switch 399 {
                            case -1:
                                v29 -= -1
                            }
                            _ = v30
                        }
                        _ = v29
                    }
                    v33 := GenS4{F0: (v8 - int8(v11))}
                    _ = fi27
                    _ = v33
                }
                _ = fi19
            }
            v5 = v5
            fmt.Println(v11, v1)
        }
    default:
        var v34 float32 = 1e10
        _ = v34
    }
    var v35 uint = 747
    if (v5 == v5) {
        if false {
            if (!((v11 > v11) && false)) {
                v36 := GenS4{F0: (v8 * (v8 >> ((v35) & 63)))}
                if len(v4) > 0 {
                    v4[0] = 2
                }
                for fi37 := 0; fi37 < 16; fi37++ {
                    if ((v5 != v5) || ((v1 + (^(v1))) <= 42)) {
                        var v38 uint32 = 65535
                        if (!(("\n\t" + "\n\t") == v3)) {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                var v40 int8 = (-(v8))
                                _ = fi39
                                _ = v40
                            }
                        } else {
                            var v41 int = 32768
                            switch (^(fi37)) {
                            case -9223372036854775808:
                                if len(v2) > 0 {
                                    v2[0] = v3
                                }
                            case -444:
                                v42 := GenS1{F0: uint(v11), F1: 74}
                                _ = v42
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (("a" + ((v3 + v3) + "hello")) + "foo bar")
                                }
                            }
                            _ = v41
                        }
                        switch "foo bar" {
                        case "hello":
                            v8 = int8(v1)
                        default:
                            v43 := v6[(fi37 >> (((v35 * uint(v11))) & 63))]
                            _ = v43
                        }
                        _ = v38
                    }
                    var v44 float64 = 0.1
                    _ = fi37
                    _ = v44
                }
                _ = v36
            } else {
                fmt.Println(v5)
            }
        } else {
            v45 := v6[901]
            var v46 int64 = -9223372036854775808
            _ = v45
            _ = v46
        }
        switch ((v3 + "a") + "a") {
        case "\n\t":
            if ((v11 & (v11 / ((v11) | 1))) < v11) {
                v47 := GenS2{F0: v3, F1: 823, F2: v3}
                _ = v47
            } else {
                var v48 string = v3
                fmt.Println(v48)
                for fi49 := 0; fi49 < 9; fi49++ {
                    v50 := GenS5{F0: (v5 / (((v5 << (((v35 ^ v35)) & 63))) | 1))}
                    v35 += v35
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
