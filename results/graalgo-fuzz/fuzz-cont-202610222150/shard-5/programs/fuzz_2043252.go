package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    v4 := []int{}
    v5 := make(map[uint32]uint32)
    var v6 bool = true
    var v7 uint32 = 603
    v6 = (-668 <= -734)
    v7 = v7
    v7 = 42
    v8 := GenS1{F0: v7}
    if (156 >= 65535) {
        v2 += v2
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int64 = 0
            _ = fi9
            _ = v10
        }
        _ = len(v4)
    }
    switch "a" {
    case "\n\t":
        switch "a" {
        case "\n\t":
            switch 9223372036854775807 {
            case 128:
                if v6 {
                    if v6 {
                        v11 := v5[(v7 & (v7 - (v7 | 2147483647)))]
                        _ = v11
                    } else {
                        v12 := GenS2{F0: 712}
                        v13 := GenS2{F0: 0}
                        _ = v12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v2 = v2
                        var v15 bool = (float32(v2) != 3.14159)
                        var v16 uint16 = 32767
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v2 = ((v2 / 1.5) - v2)
                } else {
                    var v17 bool = v6
                    if ((0.0 <= 3.14159) && false) {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 int16 = 1
                            v4 = append(v4, fi18)
                            _ = fi18
                            _ = v19
                        }
                        var v20 int16 = 1
                        _ = v20
                    }
                    _ = v17
                }
            case -202:
                v4 = append(v4, -9223372036854775808)
            }
        }
    case "a":
        if true {
            var v21 uint64 = 0
            switch "" {
            case "":
                switch 128 {
                case 0:
                    if (v6 || ((v21 > (v21 >> ((2) & 63))) && (v7 > v7))) {
                        switch 4294967295 {
                        case -53:
                            switch 1 {
                            case 128:
                                if v6 {
                                    v22 := GenS2{F0: (^(v21))}
                                    var v23 int32 = -309
                                    var v24 bool = (!(!v6))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            case 1:
                                if (2 <= 255) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 += float64(-2.25)
                                    case "a":
                                        if (("\n\t" != "a") && ((v7 % (((v7 >> ((4294967295) & 63))) | 1)) >= 256)) {
                                            _ = len(v4)
                                        }
                                    case "hello":
                                        var v25 int32 = 2147483647
                                        _ = v25
                                    }
                                    v21 += v21
                                    var v26 float32 = (0.0 * 1e10)
                                    _ = v26
                                }
                            }
                        case -167:
                            fmt.Println(v7, v21)
                        case 1:
                            v27 := GenS1{F0: 638}
                            _ = v27
                        default:
                            if (674 > 32767) {
                                if v6 {
                                    v6 = (0 <= 1)
                                }
                                if ((v21 - 2) >= v21) {
                                    for fi28 := 0; fi28 < 4; fi28++ {
                                        var v29 int16 = 32767
                                        v21 += v21
                                        if (!v6) {
                                            var v30 int8 = 1
                                            var v31 float32 = -2.25
                                            v1[fi28] = fi28
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                    if v6 {
                                        var v32 float32 = 0.1
                                        _ = v32
                                    } else {
                                        v33 := v1[-9223372036854775808]
                                        _ = v33
                                    }
                                    var v34 uint64 = (^(v21))
                                    _ = v34
                                }
                            } else {
                                var v35 uint32 = (v7 & (v7 % ((2147483648) | 1)))
                                _ = v35
                            }
                        }
                    }
                case 1:
                    var v36 int32 = 2147483647
                    _ = v36
                }
            case "\n\t":
                v2 += 3.14159
            case "foo bar":
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 int = 32768
                    v4 = append(v4, v38)
                    _ = fi37
                    _ = v38
                }
            }
            if (("hello" + "hello") >= "") {
                v4 = append(v4, 4294967295)
            }
            _ = v21
        } else {
            v39 := GenS2{F0: 131}
            _ = v39
        }
    case "hello":
        for fi40 := 0; fi40 < 6; fi40++ {
            if (369 == 0) {
                v4 = append(v4, fi40)
                v6 = ("" < "\n\t")
                fmt.Println(fi40, v6, v6)
            } else {
                if v6 {
                    var v41 int32 = 42
                    _ = v41
                }
                if len(v4) > 0 {
                    v4[0] = (fi40 - 821)
                }
            }
            v42 := GenS3{F0: -501, F1: 0}
            var v43 uint8 = 83
            _ = fi40
            _ = v42
            _ = v43
        }
    default:
        v44 := GenS4{F0: 209}
        _ = v44
    }
    var v45 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
