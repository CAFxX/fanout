package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 string; F2 uint }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int16; F1 int8; F2 string }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    fmt.Println(p0)
    var v7 int32 = 767
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 int) int {
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 0
        v10 := GenS1{F0: 1, F1: (v9 / ((65) | 1))}
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: fi8, F1: 671, F2: 0}
            p0 -= (fi8 + 128)
            var v13 int = genFn0(2)
            v14 := GenS1{F0: 32767, F1: 100}
            v13 -= genFn0(18446744073709551615)
            v15 := GenS3{F0: (!(true || (2147483648 != 992)))}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v16 float32 = 0.1
    p0 = (p0 / ((p0) | 1))
    switch ("foo bar" + "") {
    case "foo bar":
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := []int{p0, 0}
            for fi19 := 0; fi19 < 3; fi19++ {
                v16 -= 0.0
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    case "hello":
        v16 -= v16
    default:
        v20 := []int{(^(p0)), genFn0(42)}
        _ = v20
    }
    v21 := []int{(p0 >> ((622) & 63)), (p0 - (p0 + -1)), (p0 << ((42) & 63))}
    _ = p0
    _ = v16
    _ = v21
    return (p0 / ((p0) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 42
    v3 := make(map[string]int)
    var v4 int16 = 32767
    var v5 uint16 = 42
    var v6 uint64 = 367
    v22 := GenS1{F0: v4, F1: -94}
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v4)
        fmt.Println(v6)
        v5 += 42
        _ = fi23
    }
    v2 = v2
    var v24 bool = (2147483647 <= 865)
    if (121 < genFn0((v6 | 665))) {
        var v25 uint64 = (v6 | 2)
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v6, v6)
            v27 := GenS4{F0: 0, F1: -81, F2: (v25 * (v25 << ((2) & 63)))}
            _ = fi26
            _ = v27
        }
        v28 := GenS5{F0: ("\n\t" + "a")}
        _ = v25
        _ = v28
    }
    if (!v24) {
        var v29 int16 = 1000
        _ = v29
    } else {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 uint = 4294967295
            var v32 float64 = (1.5 + ((3.14159 * 1.5) + (float64(((1.5 - (1.5 * 1.5)) - 1.5)) / 3.14159)))
            _ = fi30
            _ = v31
            _ = v32
        }
        v33 := []int{127, 128, 2147483647, 0}
        fmt.Println(v5)
        _ = v33
    }
    var v34 uint64 = (^(v6))
    var v35 uint64 = v34
    switch "a" {
    case "a":
        switch genFn0((v34 | v34)) {
        case 2147483647:
            v36 := GenS6{F0: v4, F1: (("foo bar" + "hello") + "a"), F2: 1}
            _ = v36
        case 127:
            for fi37 := 0; fi37 < 17; fi37++ {
                switch fi37 {
                case 32768:
                    fmt.Println(fi37, v35, v34)
                case -2147483648:
                    delete(v3, (("hello" + (("" + "foo bar") + (("" + "\n\t") + "a"))) + "\n\t"))
                case -957:
                    var v38 uint32 = uint32(v6)
                    _ = v38
                }
                _ = fi37
            }
        case 42:
            v35 += 18446744073709551615
        }
    case "foo bar":
        if (!(false || v24)) {
            if (((1.5 * 0.0) - 1.5) != (0.1 * (1.5 + 0.0))) {
                v5 = (^(v5))
                var v39 int = 134
                v40 := GenS7{F0: 42, F1: (v2 - v2)}
                _ = v39
                _ = v40
            } else {
                v24 = false
                v35 = (v35 - 380)
            }
            if true {
                switch "" {
                case "a":
                    var v41 uint8 = 1
                    _ = v41
                case "foo bar":
                    v42 := GenS7{F0: 2, F1: v2}
                    _ = v42
                default:
                    if true {
                        v6 += v6
                        var v43 uint64 = v34
                        _ = v43
                    } else {
                        v44 := GenS8{F0: 1, F1: int32(v4)}
                        v45 := []int{int(v2)}
                        _ = v44
                        _ = v45
                    }
                }
            } else {
                if (!(42 < 42)) {
                    switch "" {
                    case "a":
                        switch "a" {
                        case "hello":
                            if true {
                                v46 := GenS5{F0: ("hello" + "hello")}
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    v48 := GenS4{F0: 2147483648, F1: -88, F2: v6}
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        if (!(v5 < v5)) {
                                            v4 = v4
                                            v50 := GenS9{F0: 0, F1: -62, F2: ("a" + ("foo bar" + ((("hello" + "hello") + "\n\t") + "hello")))}
                                            _ = v50
                                        }
                                        _ = fi49
                                    }
                                    var v51 int = 2147483647
                                    _ = fi47
                                    _ = v48
                                    _ = v51
                                }
                                _ = v46
                            }
                        case "a":
                            v2 = (v2 - (v2 - -620))
                        }
                    case "hello":
                        v52 := GenS7{F0: 261, F1: v2}
                        _ = v52
                    default:
                        for fi53 := 0; fi53 < 15; fi53++ {
                            v54 := GenS2{F0: 4294967295, F1: v35, F2: (v4 & (v4 * (v4 % (((v4 >> ((uint(v6)) & 63))) | 1))))}
                            _ = fi53
                            _ = v54
                        }
                    }
                } else {
                    v55 := GenS9{F0: 42, F1: 1, F2: ("hello" + "foo bar")}
                    _ = v55
                }
            }
        } else {
            var v56 bool = (((3.14159 > (3.14159 * 1e10)) || (!(18446744073709551615 == 1))) || (!(762 == 267)))
            var v57 float64 = 0.0
            _ = v56
            _ = v57
        }
    case "":
        v58 := GenS10{F0: v2}
        _ = v58
    }
    v59 := []int{32768, -373, genFn1(genFn1(32767)), -9223372036854775808}
    var v60 float32 = (float32(1.5) * ((0.0 + (0.1 / 0.1)) - 0.1))
    var v61 int = 32767
    var v62 uint32 = 4294967295
    v4 -= v4
    v4 = (v4 ^ v4)
    fmt.Println(v60, v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
