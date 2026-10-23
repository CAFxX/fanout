package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }

func main() {
    var v1 int64 = 189
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 float32 = (1.5 - 0.1)
    v5 := make(map[uint32]int)
    var v6 uint16 = 2
    if false {
        if true {
            v1 = (v1 % ((v1) | 1))
        }
        switch 0 {
        case -2147483648:
            v7 := GenS1{F0: (v3 == v3), F1: "a", F2: ("a" + "\n\t")}
            _ = v7
        case 956:
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(v3, v6)
                _ = fi8
            }
        default:
            v9 := []int{-2147483648}
            _ = v9
        }
        switch 61 {
        case 127:
            var v10 int64 = v1
            _ = v10
        case -9223372036854775808:
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 float32 = 0.1
                for fi13 := 0; fi13 < 4; fi13++ {
                    v1 = v1
                    var v14 int8 = 117
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        case -373:
            fmt.Println(v4)
        default:
            v6 = (v6 * (v6 >> ((131) & 63)))
        }
    }
    var v15 float32 = (3.14159 / 1.0)
    v15 -= v15
    var v16 uint = uint(v3)
    v15 = v4
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v6, v15, fi17)
        _ = fi17
    }
    var v18 float32 = v4
    if false {
        v19 := []int{}
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := []int{int(v16), fi20}
            if (-421 <= v1) {
                var v22 float32 = (v18 + (v15 / ((v18 + float32(((0.0 - float64(1e10)) - 0.1))) / v15)))
                if false {
                    switch (^(fi20)) {
                    case 4294967295:
                        var v23 uint32 = 65535
                        _ = v23
                    default:
                        v24 := GenS1{F0: true, F1: "hello", F2: "a"}
                        _ = v24
                    }
                    var v25 uint64 = 18446744073709551615
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 string = ""
                        fmt.Println(fi26, v16)
                        v6 = v6
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 float64 = (float64(v22) + 0.0)
                    _ = v28
                }
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        switch "a" {
        case "hello":
            switch -654 {
            case -1:
                var v29 bool = false
                _ = v29
            case -9223372036854775808:
                if false {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v2[("\n\t" + "")] = "\n\t"
                        _ = fi30
                    }
                    delete(v2, "hello")
                    v31 := []int{}
                    _ = v31
                } else {
                    switch 2147483647 {
                    case 0:
                        v32 := []int{888, 2147483647, -963}
                        _ = v32
                    default:
                        var v33 float64 = (1.5 + float64(v15))
                        _ = v33
                    }
                    switch -2147483648 {
                    case 609:
                        fmt.Println(v4)
                    case 42:
                        v3 = v3
                    case -2147483648:
                        var v34 int = 32767
                        _ = v34
                    default:
                        if (float32(0.0) != 0.1) {
                            v16 += (v16 / (((^(v16))) | 1))
                            for fi35 := 0; fi35 < 16; fi35++ {
                                if (false || (2 < 4294967295)) {
                                    switch (^(fi35)) {
                                    case 128:
                                        for fi36 := 0; fi36 < 5; fi36++ {
                                            v6 += (v6 * 834)
                                            if (1000 < -32768) {
                                                v6 -= (v6 * (^(v6)))
                                            }
                                            v2[""] = "\n\t"
                                            _ = fi36
                                        }
                                    case -1:
                                        var v37 int32 = -863
                                        _ = v37
                                    default:
                                        switch ("a" + (("hello" + ("\n\t" + "foo bar")) + "foo bar")) {
                                        case "\n\t":
                                            for fi38 := 0; fi38 < 18; fi38++ {
                                                var v39 uint16 = (v6 + 0)
                                                _ = fi38
                                                _ = v39
                                            }
                                        case "a":
                                            v3 = (^(v3))
                                        case "hello":
                                            var v40 int32 = 1
                                            _ = v40
                                        }
                                    }
                                    var v41 int = (fi35 * fi35)
                                    v4 = v18
                                    _ = v41
                                }
                                _ = fi35
                            }
                        }
                    }
                    var v42 int32 = -2147483648
                    _ = v42
                }
            default:
                v3 = 309
            }
        case "":
            v3 -= v3
        }
        v4 = v15
    }
    v43 := []int{-2147483648, 32767, 181}
    v18 -= float32(1.5)
    v18 = ((v4 / 0.1) * float32(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
