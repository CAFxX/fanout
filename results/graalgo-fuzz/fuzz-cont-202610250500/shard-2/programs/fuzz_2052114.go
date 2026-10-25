package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 int64, p2 int) uint16 {
    if false {
        fmt.Println(p0, p1, p1)
    } else {
        var v6 uint = 352
        v7 := GenS1{F0: "hello"}
        _ = v6
        _ = v7
    }
    p0 -= 2
    var v8 int32 = -1
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{-876}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if (227 >= 36) {
            fmt.Println(p0, p1)
        }
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 0
}
func genFn1() int32 {
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch (fi12 << ((81) & 63)) {
            case 42:
                switch "foo bar" {
                case "a":
                    v14 := []int{}
                    _ = v14
                case "\n\t":
                    fmt.Println(fi13, fi12)
                default:
                    v15 := GenS1{F0: "\n\t"}
                    _ = v15
                }
            case 4294967295:
                switch (^(fi13)) {
                case -1:
                    if (true && false) {
                        var v16 int32 = 42
                        v17 := GenS1{F0: "hello"}
                        switch fi13 {
                        case -9223372036854775808:
                            v18 := []int{32768}
                            _ = v18
                        case 237:
                            v19 := GenS2{F0: 4294967295, F1: 80}
                            _ = v19
                        default:
                            if false {
                                switch (fi13 << ((255) & 63)) {
                                case 42:
                                    v16 = v16
                                case -2147483648:
                                    v20 := GenS3{F0: -1}
                                    _ = v20
                                default:
                                    var v21 string = "hello"
                                    _ = v21
                                }
                                fmt.Println(v16)
                                v22 := GenS1{F0: ("a" + "foo bar")}
                                _ = v22
                            }
                        }
                        _ = v16
                        _ = v17
                    } else {
                        if (232 <= -752) {
                            var v23 float32 = (3.14159 * (1e10 + (1e10 + ((0.0 + 0.1) - 1e10))))
                            var v24 uint64 = 18446744073709551615
                            if true {
                                v24 -= v24
                                if true {
                                    v23 = v23
                                    v25 := GenS1{F0: "foo bar"}
                                    _ = v25
                                } else {
                                    v23 += v23
                                    v26 := []int{-9223372036854775808, (fi13 + fi12), fi12}
                                    var v27 uint16 = 32767
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                if (-1 <= -1) {
                                    fmt.Println(v24, v24)
                                    v28 := []int{fi13, (fi13 % ((32768) | 1))}
                                    _ = v28
                                }
                                var v29 uint8 = 127
                                _ = v29
                            }
                            _ = v23
                            _ = v24
                        } else {
                            var v30 string = ""
                            v31 := GenS3{F0: -32768}
                            v32 := GenS4{F0: ((v30 + (v30 + "")) + v30)}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                case 32768:
                    var v33 uint16 = 1
                    fmt.Println(v33, fi12, fi12)
                    _ = v33
                }
            }
            _ = fi13
        }
        var v34 float64 = 3.14159
        _ = fi12
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        v36 := GenS5{F0: 65535}
        _ = fi35
        _ = v36
    }
    return -2147483648
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    delete(v2, 127)
    fmt.Println(v1)
    var v37 int8 = -118
    fmt.Println(v1)
    fmt.Println(v1, v37)
    v37 = v37
    fmt.Println(v37, v37)
    if ((!(float64(-2.25) == 1.5)) && true) {
        v38 := GenS6{F0: (42 >= 0), F1: int64(v37), F2: 769}
        _ = v38
    } else {
        v1 = v1
    }
    if (false || false) {
        var v39 int = -9223372036854775808
        _ = v39
    } else {
        v37 -= v37
        var v40 uint8 = 1
        _ = v40
    }
    v41 := []int{174, 32768}
    var v42 uint32 = 776
    var v43 uint8 = 0
    v37 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
