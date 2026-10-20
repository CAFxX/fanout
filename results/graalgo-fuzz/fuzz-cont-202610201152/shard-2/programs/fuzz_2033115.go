package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int32; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int }
func genFn0(p0 uint8, p1 int16) string {
    switch "foo bar" {
    case "\n\t":
        var v7 string = "foo bar"
        _ = v7
    case "hello":
        if (0.0 != (0.0 / 1.0)) {
            var v8 int32 = 555
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "a":
                var v9 uint8 = p0
                _ = v9
            case "":
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint8 = p0
                    var v12 uint64 = 42
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "\n\t":
                v13 := []int{-2147483648, 128, 127, -2147483648}
                _ = v13
            }
            _ = v8
        } else {
            p1 = (p1 | 0)
        }
    }
    v14 := []int{int(p1), -985, 127}
    _ = p0
    _ = p1
    _ = v14
    return "\n\t"
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 uint64 = 42
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    v5 := []string{"", "\n\t", "", ""}
    var v6 int64 = -1
    if true {
        delete(v3, 2147483648)
    }
    if (2 > 511) {
        var v15 int32 = 2147483647
        var v16 int16 = -32768
        var v17 int32 = v15
        _ = v15
        _ = v16
        _ = v17
    }
    delete(v3, 433)
    for fi18 := 0; fi18 < 10; fi18++ {
        v3[42] = 42
        v4 = (1e10 > v1)
        _ = fi18
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 42 {
    case 1:
        v19 := GenS1{F0: 145, F1: 32768}
        _ = v19
    }
    var v20 uint64 = v2
    var v21 string = ((("hello" + (("a" + "foo bar") + "hello")) + "hello") + "\n\t")
    v22 := GenS2{F0: 4294967295, F1: (v6 * int64(v2))}
    if v4 {
        var v23 uint16 = 32767
        _ = v23
    } else {
        if (v2 > 4294967295) {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint = 1
                for fi26 := 0; fi26 < 19; fi26++ {
                    switch fi24 {
                    case 127:
                        var v27 uint8 = 40
                        _ = v27
                    }
                    if false {
                        v4 = (1 <= int64(v25))
                        var v28 uint32 = 65535
                        v21 = ""
                        _ = v28
                    } else {
                        var v29 uint64 = 770
                        if len(v5) > 0 {
                            v5[0] = "foo bar"
                        }
                        _ = v29
                    }
                    _ = fi26
                }
                var v30 uint32 = 4294967295
                _ = fi24
                _ = v25
                _ = v30
            }
        } else {
            var v31 uint16 = 1
            for fi32 := 0; fi32 < 17; fi32++ {
                v21 = v21
                v3[47] = 42
                _ = fi32
            }
            _ = v31
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            v34 := GenS3{F0: 42}
            if (-1 == 42) {
                switch -2147483648 {
                case 0:
                    _ = len(v5)
                case 1:
                    switch (^(fi33)) {
                    case 2147483647:
                        v35 := GenS1{F0: 147, F1: fi33}
                        _ = v35
                    }
                default:
                    if false {
                        v36 := GenS4{F0: v21, F1: 0, F2: 42}
                        delete(v3, 2)
                        var v37 int16 = 1
                        _ = v36
                        _ = v37
                    }
                }
                var v38 int32 = 0
                _ = v38
            } else {
                var v39 uint16 = 255
                for fi40 := 0; fi40 < 14; fi40++ {
                    var v41 uint8 = 127
                    if v4 {
                        var v42 uint32 = 65535
                        if true {
                            fmt.Println(fi33)
                            var v43 uint32 = (^(v42))
                            v20 = v2
                            _ = v43
                        }
                        _ = v42
                    } else {
                        fmt.Println(v1)
                        switch int(v41) {
                        case 1:
                            v2 += (v20 * 739)
                        case -1:
                            if v4 {
                                if ("foo bar" <= ("\n\t" + "a")) {
                                    if v4 {
                                        var v44 float64 = float64((v1 / float32(1.5)))
                                        v4 = v4
                                        var v45 int32 = 2147483647
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch ((v21 + v21) + "a") {
                                        case "\n\t":
                                            var v46 float64 = (((1e10 * (0.0 * 1.5)) - (3.14159 / -2.25)) - (3.14159 + 3.14159))
                                            _ = v46
                                        case "":
                                            if ((v6 * 9223372036854775807) < (v6 - (v6 >> ((722) & 63)))) {
                                                v41 += (v41 | (v41 % ((uint8(v39)) | 1)))
                                            }
                                        default:
                                            fmt.Println(v1, v41)
                                        }
                                        fmt.Println(v2, fi40, v39)
                                        if (1 <= 359) {
                                            switch "\n\t" {
                                            case "hello":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            v2 += (v2 + v20)
                                            v4 = (!v4)
                                        } else {
                                            fmt.Println(v21, v20, v20)
                                            switch (("\n\t" + v21) + v21) {
                                            case "foo bar":
                                                if v4 {
                                                    v20 += 0
                                                    fmt.Println(v6, v39)
                                                } else {
                                                    v39 -= uint16(v20)
                                                    switch fi40 {
                                                    case 128:
                                                        if true {
                                                            v47 := v3[2]
                                                            var v48 uint16 = 998
                                                            v4 = (((255 >= 0) && (int64(v41) > 0)) || v4)
                                                            _ = v47
                                                            _ = v48
                                                        } else {
                                                            v20 = v2
                                                            var v49 uint8 = 0
                                                            _ = v49
                                                        }
                                                    case 42:
                                                        var v50 int16 = -32768
                                                        _ = v50
                                                    case -949:
                                                        v51 := GenS5{F0: 0}
                                                        _ = v51
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        case 892:
                            var v52 uint64 = v20
                            _ = v52
                        default:
                            v5 = append(v5, v21)
                        }
                        v21 = "\n\t"
                    }
                    delete(v3, 65535)
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
            _ = fi33
            _ = v34
        }
    }
    if v4 {
        if v4 {
            fmt.Println(v21, v21)
            v3[808] = 28
            var v53 int8 = 100
            _ = v53
        }
    } else {
        switch 15 {
        case -368:
            var v54 float64 = 0.1
            _ = v54
        case -832:
            v55 := GenS6{F0: 128, F1: 127}
            _ = v55
        case 32768:
            var v56 uint8 = 127
            _ = v56
        default:
            v20 -= 1
        }
    }
    v20 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
