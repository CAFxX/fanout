package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
func genFn0(p0 int64) string {
    if false {
        var v7 float32 = 3.14159
        _ = v7
    }
    v8 := []int{32768, 42, 2147483647}
    fmt.Println(p0)
    if false {
        for fi9 := 0; fi9 < 12; fi9++ {
            p0 += p0
            var v10 float64 = (0.0 + 3.14159)
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "a" {
            case "a":
                p0 = p0
            case "\n\t":
                if true {
                    p0 -= (p0 ^ (p0 * p0))
                    var v12 uint8 = 2
                    _ = v12
                }
            }
            var v13 uint32 = 0
            p0 = (p0 >> ((0) & 63))
            _ = fi11
            _ = v13
        }
        var v14 float32 = 1e10
        _ = v14
    }
    _ = p0
    _ = v8
    return ("a" + ("\n\t" + "\n\t"))
}

func main() {
    var v1 int64 = 942
    v2 := []uint32{2147483648, 2, 0, 42}
    var v3 float64 = (-2.25 / (1e10 + 3.14159))
    var v4 float32 = (1e10 / (0.1 + (-2.25 - float32(0.0))))
    var v5 float32 = v4
    v6 := make(map[int]uint32)
    var v15 int8 = 68
    fmt.Println(v5)
    fmt.Println(v3, v5, v15)
    var v16 int32 = -30
    for fi17 := 0; fi17 < 8; fi17++ {
        if (!(!false)) {
            var v18 int8 = 0
            var v19 uint = 920
            for fi20 := 0; fi20 < 10; fi20++ {
                switch "\n\t" {
                case "a":
                    v4 = v5
                case "\n\t":
                    v4 = (v5 + v4)
                }
                v5 = 0.0
                var v21 int32 = (v16 | (v16 - -2147483648))
                _ = fi20
                _ = v21
            }
            _ = v18
            _ = v19
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS1{F0: 0, F1: 1, F2: v15}
                _ = fi22
                _ = v23
            }
        }
        _ = fi17
    }
    v3 -= v3
    var v24 float64 = (v3 - v3)
    var v25 uint = 1
    if false {
        delete(v6, 42)
        if (255 != 127) {
            v15 = int8(v1)
            fmt.Println(v24)
            v3 += 1.5
        } else {
            v26 := GenS1{F0: int8(v25), F1: v15, F2: -128}
            switch "hello" {
            case "":
                if ("" == ("hello" + ("a" + "hello"))) {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        if true {
                            v4 = v4
                            v2 = append(v2, 65535)
                        }
                        var v28 float64 = v3
                        if (v4 == v4) {
                            for fi29 := 0; fi29 < 7; fi29++ {
                                var v30 float32 = v5
                                _ = fi29
                                _ = v30
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v15 = 115
                }
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if false {
                    var v31 bool = (true && true)
                    _ = v31
                } else {
                    v5 = (v4 * v5)
                    v32 := GenS2{F0: 69, F1: 32767, F2: 0}
                    var v33 uint8 = 0
                    _ = v32
                    _ = v33
                }
            }
            _ = v26
        }
        if (2 == 65535) {
            if (false && (42 != (v15 % ((42) | 1)))) {
                v3 += float64(v5)
                var v34 uint64 = 1
                _ = v34
            } else {
                if (false && ("foo bar" == "a")) {
                    for fi35 := 0; fi35 < 5; fi35++ {
                        v36 := GenS3{F0: (-(fi35)), F1: 1000, F2: -873}
                        fmt.Println(v3, v4, v24)
                        _ = fi35
                        _ = v36
                    }
                } else {
                    switch "a" {
                    case "a":
                        v37 := GenS1{F0: (v15 / ((v15) | 1)), F1: -56, F2: (v15 % ((v15) | 1))}
                        _ = v37
                    case "\n\t":
                        for fi38 := 0; fi38 < 16; fi38++ {
                            var v39 float64 = (v3 * 1.5)
                            var v40 uint8 = 255
                            var v41 int16 = 32767
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    default:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v15, v16, v15)
                        }
                    }
                    var v42 uint8 = 1
                    _ = v42
                }
                var v43 uint32 = 924
                _ = v43
            }
            v3 = v3
        }
    }
    var v44 float64 = v3
    fmt.Println(v3, v15)
    if false {
        var v45 int32 = (v16 | (v16 * (v16 + 720)))
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
