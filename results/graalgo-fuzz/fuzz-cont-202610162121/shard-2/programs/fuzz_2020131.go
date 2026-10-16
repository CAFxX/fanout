package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
func genFn0(p0 uint16) int {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 100, F1: (("hello" + "a") + "a")}
        _ = fi7
        _ = v8
    }
    if ((0.1 == 1.5) || false) {
        v9 := []int{}
        if (!false) {
            var v10 int32 = 42
            for fi11 := 0; fi11 < 16; fi11++ {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    var v12 uint8 = 103
                    _ = v12
                }
                if ((false && true) && ((("\n\t" + "\n\t") + "hello") == ("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))))) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 int64 = 231
                        v15 := GenS1{F0: 100, F1: ""}
                        v16 := GenS1{F0: 0, F1: ("" + "hello")}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS1{F0: 3, F1: ""}
                        v10 += (v10 | (v10 >> ((4294967295) & 63)))
                        _ = fi17
                        _ = v18
                    }
                    _ = len(v9)
                }
                _ = fi11
            }
            _ = v10
        }
        p0 = (p0 * p0)
        _ = v9
    }
    _ = p0
    return 42
}
func genFn1(p0 int8, p1 uint) uint32 {
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v19 float64 = -2.25
        _ = v19
    case "a":
        var v20 uint32 = 42
        _ = v20
    case "\n\t":
        if ((p0 % (((-(p0))) | 1)) != p0) {
            v21 := []int{-9223372036854775808}
            _ = v21
        }
    default:
        var v22 float64 = 0.1
        _ = v22
    }
    var v23 uint16 = 0
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := GenS1{F0: 100, F1: "a"}
        _ = fi24
        _ = v25
    }
    var v26 int16 = 0
    _ = p0
    _ = p1
    _ = v23
    _ = v26
    return 65535
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647}
    v3 := make(map[int]string)
    var v4 int32 = -2147483648
    var v5 int32 = v4
    var v6 string = ""
    fmt.Println(v6, v4, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v3)
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS1{F0: 127, F1: ""}
        _ = fi27
        _ = v28
    }
    var v29 bool = (2 < 1)
    switch v6 {
    case "\n\t":
        var v30 float64 = (float64(1e10) * float64(float32(((((-2.25 / 1e10) + 1e10) * 0.0) * -2.25))))
        _ = v30
    case "hello":
        v29 = true
    default:
        var v31 bool = (0 == -761)
        _ = v31
    }
    switch genFn0(255) {
    case 128:
        v6 = ""
    case 32767:
        switch genFn0(32767) {
        case 447:
            v4 += 545
        case -554:
            var v32 float32 = 1e10
            _ = v32
        case 4294967295:
            for fi33 := 0; fi33 < 20; fi33++ {
                v3[(^(fi33))] = ""
                fmt.Println(v29, v6, v4)
                if (!v29) {
                    var v34 uint64 = 847
                    v4 -= 2147483647
                    var v35 float32 = (0.1 + float32((1e10 + float64(float32(-2.25)))))
                    _ = v34
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 12; fi36++ {
                        switch fi36 {
                        case -9223372036854775808:
                            var v37 uint8 = 42
                            _ = len(v3)
                            _ = v37
                        case -2147483648:
                            switch v6 {
                            case "\n\t":
                                if v29 {
                                    v38 := GenS1{F0: -65, F1: (v6 + ((v6 + v6) + v6))}
                                    v39 := GenS1{F0: 0, F1: v6}
                                    var v40 float64 = 0.1
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 uint = 2
                                    switch v6 {
                                    case "\n\t":
                                        var v42 uint64 = 1
                                        _ = v42
                                    case "":
                                        _ = len(v1)
                                    case "a":
                                        var v43 uint32 = 2147483647
                                        _ = v43
                                    default:
                                        v44 := GenS1{F0: -1, F1: "foo bar"}
                                        _ = v44
                                    }
                                    fmt.Println(fi36)
                                    _ = v41
                                }
                            default:
                                var v45 float32 = 0.1
                                _ = v45
                            }
                        case 0:
                            fmt.Println(v6, fi36)
                        default:
                            v3[(fi36 + (fi33 << ((65535) & 63)))] = "hello"
                            fmt.Println(fi36, v4, fi33)
                            v5 = v4
                        }
                        _ = fi36
                    }
                    switch "hello" {
                    case "\n\t":
                        if true {
                            v46 := v3[genFn0(2)]
                            v29 = v29
                            if (!((1 > 1) && v29)) {
                                v6 = v46
                                switch 32767 {
                                case -882:
                                    var v47 float64 = 0.0
                                    _ = v47
                                case 1:
                                    v6 = (v6 + "a")
                                }
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (true || true) {
                                    if len(v2) > 0 {
                                        v2[0] = 462
                                    }
                                    if true {
                                        v29 = (42 != 331)
                                        var v48 uint32 = 2147483648
                                        for fi49 := 0; fi49 < 20; fi49++ {
                                            _ = len(v2)
                                            _ = fi49
                                        }
                                        _ = v48
                                    }
                                }
                                if v29 {
                                    v46 = "\n\t"
                                } else {
                                    var v50 float64 = (1e10 + 0.0)
                                    if v29 {
                                        fmt.Println(v29, v50, v29)
                                        v6 = v46
                                        fmt.Println(v6)
                                    }
                                    _ = v50
                                }
                            }
                            _ = v46
                        }
                    case "foo bar":
                        for fi51 := 0; fi51 < 6; fi51++ {
                            var v52 bool = (v29 && (42 >= 1))
                            v52 = (0.1 >= -2.25)
                            _ = fi51
                            _ = v52
                        }
                    case "hello":
                        var v53 uint = 886
                        _ = v53
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    for fi54 := 0; fi54 < 12; fi54++ {
                        v55 := GenS2{F0: -128, F1: 101}
                        fmt.Println(fi54, v6, v6)
                        var v56 uint64 = 849
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                }
                _ = fi33
            }
        default:
            for fi57 := 0; fi57 < 11; fi57++ {
                v6 = "a"
                _ = fi57
            }
        }
    }
    _ = len(v3)
    if v29 {
        var v58 float64 = 3.14159
        v59 := GenS2{F0: -1, F1: 42}
        _ = v58
        _ = v59
    }
    v60 := GenS3{F0: v29, F1: -9223372036854775808, F2: 24}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
