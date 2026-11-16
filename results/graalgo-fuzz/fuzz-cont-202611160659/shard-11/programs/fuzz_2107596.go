package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0(p0 uint8, p1 int, p2 uint32) string {
    v6 := []int{468}
    p1 += (p1 * (-(p1)))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("" + ("foo bar" + ("\n\t" + "hello")))
}

func main() {
    v1 := []string{"", ("hello" + ((("hello" + "\n\t") + "foo bar") + ("" + "")))}
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 uint64 = 844
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = "foo bar"
        var v9 int32 = -759
        var v10 uint = 18446744073709551615
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 = append(v1, "a")
        if (74 > 1) {
            delete(v4, 773)
        } else {
            v5 -= 1
            switch fi11 {
            case -1:
                var v12 int16 = 1
                _ = v12
            case 60:
                if true {
                    switch ("" + (("\n\t" + "\n\t") + ("foo bar" + "hello"))) {
                    case "":
                        for fi13 := 0; fi13 < 13; fi13++ {
                            _ = len(v3)
                            v5 = (v5 * (^(v5)))
                            var v14 uint32 = 2147483648
                            _ = fi13
                            _ = v14
                        }
                    case "a":
                        var v15 int16 = 32767
                        _ = v15
                    case "foo bar":
                        var v16 float64 = ((1.5 + (-2.25 * -2.25)) + 3.14159)
                        _ = v16
                    }
                }
            default:
                v5 = 174
            }
        }
        _ = fi11
    }
    if true {
        v5 = 2
        switch 128 {
        case -684:
            if true {
                delete(v4, 590)
            } else {
                v17 := GenS1{F0: "hello", F1: 1, F2: 0}
                if (-48 <= 100) {
                    v2 = append(v2, "a")
                }
                if false {
                    v18 := GenS2{F0: 127}
                    fmt.Println(v5)
                    var v19 uint = 255
                    _ = v18
                    _ = v19
                }
                _ = v17
            }
        case 210:
            v5 += v5
        case 128:
            var v20 string = "hello"
            _ = v20
        }
        var v21 int = 163
        _ = v21
    } else {
        fmt.Println(v5)
    }
    v22 := GenS1{F0: "a", F1: 704, F2: 64}
    if len(v1) > 0 {
        v1[0] = "a"
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if false {
            var v24 int = fi23
            _ = v24
        } else {
            v25 := v3[2147483648]
            var v26 int16 = 174
            _ = v25
            _ = v26
        }
        if ((true || true) && false) {
            v5 -= 4294967295
        } else {
            _ = len(v4)
            fmt.Println(fi23, fi23)
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v5)
            default:
                switch "foo bar" {
                case "a":
                    var v27 string = "a"
                    _ = v27
                case "":
                    var v28 int32 = 0
                    _ = v28
                default:
                    v29 := GenS2{F0: 193}
                    _ = v29
                }
            }
        }
        _ = fi23
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        fmt.Println(v5, v5)
        if false {
            v31 := GenS3{F0: 157, F1: -1, F2: true}
            switch "hello" {
            case "a":
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            case "":
                fmt.Println(v5, fi30)
            case "foo bar":
                if ((fi30 * fi30) < fi30) {
                    var v32 bool = ((3.14159 - 3.14159) <= float64((1.5 + 3.14159)))
                    if (2147483647 != 42) {
                        var v33 uint64 = v5
                        v34 := GenS3{F0: uint32(v5), F1: 127, F2: v32}
                        _ = v33
                        _ = v34
                    }
                    var v35 int8 = -1
                    _ = v32
                    _ = v35
                }
            default:
                v5 = (v5 + 416)
            }
            _ = v31
        } else {
            v36 := GenS4{F0: (fi30 * (fi30 << ((4294967295) & 63))), F1: v5, F2: 2147483647}
            _ = v36
        }
        _ = fi30
    }
    v37 := GenS5{F0: 65535, F1: 42}
    _ = len(v4)
    v1 = append(v1, "hello")
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 15; fi39++ {
            v5 -= v5
            _ = fi39
        }
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
