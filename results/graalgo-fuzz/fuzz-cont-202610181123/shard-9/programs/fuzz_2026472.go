package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 int; F1 string; F2 int16 }
func genFn0() uint8 {
    var v7 string = "a"
    if (27 > 255) {
        v7 = (v7 + v7)
        v7 = (v7 + ((v7 + v7) + v7))
        v8 := []int{2147483647, -1, 4294967295, 42}
        _ = v8
    }
    v9 := GenS1{F0: 1}
    v7 = v7
    for fi10 := 0; fi10 < 19; fi10++ {
        switch v7 {
        case "\n\t":
            v11 := GenS1{F0: -70}
            _ = v11
        case "foo bar":
            v7 = ("hello" + v7)
        case "":
            v7 = "hello"
        }
        fmt.Println(v7, v7)
        v7 = (v7 + v7)
        _ = fi10
    }
    _ = v7
    _ = v9
    return 2
}
func genFn1(p0 uint8, p1 int32, p2 string) int {
    if (false && false) {
        var v12 float64 = 1e10
        var v13 int = 127
        if (127 <= 0) {
            var v14 uint = 1
            _ = v14
        }
        _ = v12
        _ = v13
    } else {
        p1 += p1
        if (!("foo bar" < "a")) {
            var v15 uint32 = 2
            p2 = p2
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 20; fi17++ {
                    if false {
                        fmt.Println(fi17, p0)
                        v18 := GenS2{F0: 115}
                        switch fi17 {
                        case 42:
                            fmt.Println(p0, p2)
                        case 4294967295:
                            fmt.Println(p1, fi16, p1)
                        case 0:
                            v19 := []int{32767, 1, fi17, int(p0)}
                            _ = v19
                        }
                        _ = v18
                    } else {
                        p1 += 0
                    }
                    v20 := []int{int(v15)}
                    _ = fi17
                    _ = v20
                }
                var v21 int32 = 1
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int64 = -62
                    _ = fi22
                    _ = v23
                }
                _ = fi16
                _ = v21
            }
            _ = v15
        } else {
            v24 := []int{}
            switch p2 {
            case "hello":
                var v25 float32 = (3.14159 / ((float32(1e10) / 0.1) * 1e10))
                _ = v25
            case "a":
                var v26 float64 = (0.1 * 1.5)
                _ = v26
            }
            var v27 bool = (42 >= 127)
            _ = v24
            _ = v27
        }
        if (0 != 1000) {
            v28 := []int{462, 42, -9223372036854775808}
            var v29 int64 = 42
            _ = v28
            _ = v29
        }
    }
    v30 := GenS3{F0: -10}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 235
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 2
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]string)
    fmt.Println(v1)
    _ = len(v6)
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 uint = 2
        switch fi31 {
        case -9223372036854775808:
            v32 = v32
        case 0:
            fmt.Println(v1, v32)
        case -1:
            delete(v6, "a")
        default:
            switch ("hello" + "foo bar") {
            case "\n\t":
                v33 := GenS4{F0: (v32 % (((v32 ^ (v32 * (v32 * (v32 << (((v32 * (v32 << ((0) & 63)))) & 63)))))) | 1)), F1: (^(v1))}
                _ = v33
            case "":
                var v34 uint16 = v2
                _ = v34
            }
        }
        fmt.Println(v1)
        _ = fi31
        _ = v32
    }
    var v35 uint32 = 2147483648
    v36 := []int{32768, -1, genFn1(128, 42, "\n\t"), 32768}
    v37 := GenS4{F0: uint(v35), F1: v1}
    if (true || (!(v5 <= v5))) {
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 float32 = 1.5
            _ = fi38
            _ = v39
        }
        switch "a" {
        case "hello":
            if false {
                if (true || v3) {
                    switch -957 {
                    case -898:
                        var v40 int8 = -128
                        _ = v40
                    default:
                        v41 := GenS2{F0: (v2 >> ((265) & 63))}
                        _ = v41
                    }
                    v35 = v35
                }
            } else {
                v42 := v6["hello"]
                v5 += 42
                if v3 {
                    switch -809 {
                    case -2147483648:
                        if ((1 <= 65535) && v3) {
                            fmt.Println(v5, v5, v3)
                            v43 := GenS5{F0: ((^(v2)) != (v2 * (v2 * 261))), F1: 4294967295, F2: v1}
                            fmt.Println(v3)
                            _ = v43
                        } else {
                            switch "a" {
                            case "hello":
                                var v44 uint8 = 251
                                _ = v44
                            }
                            v5 = (^(v5))
                        }
                    case 2147483647:
                        v35 = (v35 % (((v35 / (((v35 % ((v35) | 1))) | 1))) | 1))
                    }
                } else {
                    delete(v6, (v42 + (v42 + "")))
                }
                _ = v42
            }
        case "foo bar":
            switch ("foo bar" + "") {
            case "":
                v45 := GenS6{F0: -2147483648, F1: 1, F2: (!((-1 < -65) && v3))}
                _ = v45
            case "hello":
                delete(v6, "foo bar")
            }
        }
    } else {
        v46 := GenS7{F0: 32767, F1: ("" + ""), F2: (-(v1))}
        for fi47 := 0; fi47 < 1; fi47++ {
            fmt.Println(v3)
            delete(v6, "\n\t")
            switch "" {
            case "":
                v1 -= (v1 ^ (v1 ^ int16(fi47)))
            default:
                fmt.Println(v35)
            }
            _ = fi47
        }
        _ = v46
    }
    fmt.Println(v5, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
