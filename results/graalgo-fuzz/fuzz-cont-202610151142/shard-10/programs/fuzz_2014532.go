package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = ((false || false) || false)
    v3 := make(map[uint32]bool)
    var v4 string = ("a" + "foo bar")
    var v5 uint8 = 67
    var v6 uint32 = 100
    v2 = v2
    v2 = (!v2)
    _ = len(v3)
    v2 = ((v2 && v2) || (((3.14159 / ((0.1 * -2.25) / (1e10 * (1.5 / float32((float64(1.5) * (-2.25 / 0.1))))))) * 1.5) <= (1.5 + ((1e10 + 1.5) / float32(1e10)))))
    switch -1 {
    case 127:
        if ((v5 - 1) > v5) {
            switch v4 {
            case "foo bar":
                var v7 uint16 = uint16(v5)
                _ = v7
            case "a":
                v4 = "\n\t"
            case "hello":
                v5 -= (v5 - v5)
            }
            v2 = (42 < v1)
            v1 += v1
        } else {
            v8 := GenS1{F0: 42}
            _ = v8
        }
    case 2147483647:
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{(fi9 + fi9), fi9, (fi9 | 725)}
            v11 := GenS2{F0: v5}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        switch v4 {
        case "\n\t":
            if true {
                fmt.Println(v5)
                var v12 float64 = (float64(3.14159) + 3.14159)
                var v13 float32 = 0.0
                _ = v12
                _ = v13
            } else {
                v14 := []int{1, 32767, 9223372036854775807, 127}
                _ = v14
            }
        case "foo bar":
            v15 := v3[(v6 - 42)]
            _ = v15
        default:
            v2 = v2
        }
    }
    v3[v6] = true
    for fi16 := 0; fi16 < 7; fi16++ {
        delete(v3, v6)
        for fi17 := 0; fi17 < 5; fi17++ {
            _ = len(v3)
            v2 = true
            _ = fi17
        }
        _ = fi16
    }
    v4 = v4
    if (42 != 1000) {
        v4 = v4
        fmt.Println(v4)
    }
    var v18 int = 229
    v19 := GenS2{F0: (^(v5))}
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(fi20, v5)
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS3{F0: v5, F1: (v1 << ((1) & 63))}
            v1 += v1
            _ = fi21
            _ = v22
        }
        if (!(42 != -1)) {
            switch 32767 {
            case 32768:
                v2 = true
            case 800:
                var v23 int = v18
                _ = v23
            default:
                v3[uint32(v1)] = (int8(v6) >= 42)
            }
            _ = len(v3)
        }
        _ = fi20
    }
    v24 := GenS3{F0: (v5 * v5), F1: v1}
    for fi25 := 0; fi25 < 8; fi25++ {
        if v2 {
            switch v4 {
            case "foo bar":
                if v2 {
                    v1 = 792
                    switch v18 {
                    case 9223372036854775807:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            _ = len(v3)
                            v27 := v3[uint32(v1)]
                            fmt.Println(fi26, v6, v18)
                            v27 = (-9223372036854775808 <= -626)
                            _ = fi26
                            _ = v27
                        }
                    case 820:
                        v3[(v6 ^ 2147483648)] = false
                    case 857:
                        v1 = uint(v6)
                    default:
                        if true {
                            var v28 uint = v1
                            switch v4 {
                            case "hello":
                                fmt.Println(v4, v2, v2)
                            default:
                                var v29 int16 = -32768
                                _ = v29
                            }
                            _ = v28
                        } else {
                            v2 = false
                            if (!((v2 && ((v6 & (v6 - v6)) < 4294967295)) || (9223372036854775807 == 567))) {
                                v30 := GenS1{F0: 42}
                                _ = len(v3)
                                _ = v30
                            } else {
                                if true {
                                    v1 = (v1 | v1)
                                }
                            }
                            for fi31 := 0; fi31 < 2; fi31++ {
                                v3[(v6 >> (((v1 << ((v1) & 63))) & 63))] = false
                                _ = fi31
                            }
                        }
                    }
                } else {
                    var v32 float64 = float64((1.5 + -2.25))
                    var v33 int64 = 42
                    delete(v3, (v6 & 42))
                    _ = v32
                    _ = v33
                }
            case "a":
                if v2 {
                    v3[(v6 << ((255) & 63))] = v2
                } else {
                    var v34 float32 = -2.25
                    switch v4 {
                    case "hello":
                        switch (^(fi25)) {
                        case 732:
                            var v35 string = (v4 + "\n\t")
                            _ = v35
                        case 437:
                            switch v4 {
                            case "\n\t":
                                v36 := GenS2{F0: (v5 >> ((66) & 63))}
                                _ = v36
                            case "foo bar":
                                v37 := []int{}
                                _ = v37
                            }
                        default:
                            v2 = (42 < 100)
                        }
                    default:
                        v1 = v1
                    }
                    switch v18 {
                    case 813:
                        if v2 {
                            var v38 int16 = -1
                            v1 = v1
                            v18 += (v18 | (fi25 + (fi25 + 0)))
                            _ = v38
                        } else {
                            switch fi25 {
                            case 128:
                                v2 = (!v2)
                            case 4294967295:
                                if (false && true) {
                                    v39 := []int{32768}
                                    fmt.Println(v1, v1)
                                    _ = v39
                                } else {
                                    if (2 <= 65535) {
                                        v40 := GenS3{F0: v5, F1: (v1 + v1)}
                                        delete(v3, v6)
                                        v41 := GenS4{F0: (((v4 + v4) + (v4 + "foo bar")) + ("" + "\n\t")), F1: 125}
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            case 952:
                                var v42 int8 = 6
                                _ = v42
                            default:
                                v43 := v3[1]
                                _ = v43
                            }
                            var v44 uint64 = 1
                            _ = v44
                        }
                    }
                    _ = v34
                }
            }
        }
        if (false || v2) {
            v45 := []int{}
            v1 = v1
            _ = v45
        } else {
            var v46 int = (fi25 % ((-354) | 1))
            switch ((v4 + (v4 + v4)) + v4) {
            case "hello":
                var v47 int64 = -446
                _ = v47
            case "a":
                if v2 {
                    var v48 string = v4
                    var v49 bool = (127 < 100)
                    fmt.Println(v4)
                    _ = v48
                    _ = v49
                } else {
                    v5 += (v5 & v5)
                }
            case "foo bar":
                delete(v3, v6)
            default:
                var v50 uint8 = (^(v5))
                _ = v50
            }
            if (!(!true)) {
                var v51 uint64 = 2
                v52 := GenS1{F0: 32767}
                _ = v51
                _ = v52
            }
            _ = v46
        }
        var v53 uint64 = 42
        _ = fi25
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
