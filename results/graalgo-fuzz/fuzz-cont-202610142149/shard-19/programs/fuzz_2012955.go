package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32) uint16 {
    p0 = 2147483647
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            v8 := GenS1{F0: "foo bar", F1: 0, F2: fi7}
            fmt.Println(p0, fi7)
            _ = v8
        }
        if ((3.14159 / -2.25) < (((((0.0 - 0.0) + (((0.0 / ((float64(-2.25) / 1e10) / 1.0)) - 1.5) / -2.25)) * 1.5) + -2.25) * 1e10)) {
            var v9 uint32 = (p0 & p0)
            _ = v9
        } else {
            var v10 float64 = 1e10
            var v11 float64 = 3.14159
            if false {
                p0 = 2
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    fmt.Println(fi7)
                    var v13 float64 = ((v11 - v10) / v10)
                    if true {
                        var v14 uint16 = 196
                        v10 = ((v10 * 1e10) + v10)
                        _ = v14
                    } else {
                        var v15 int64 = 0
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                v16 := GenS1{F0: "a", F1: 42, F2: fi7}
                _ = v16
            }
            _ = v10
            _ = v11
        }
        p0 = 159
        _ = fi7
    }
    v17 := []int{4294967295}
    _ = p0
    _ = v17
    return 255
}
func genFn1(p0 int8, p1 int32, p2 int) int {
    p0 = 1
    switch "a" {
    case "\n\t":
        var v18 uint16 = 65535
        _ = v18
    default:
        var v19 int8 = 42
        _ = v19
    }
    if true {
        p2 = p2
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                if false {
                    p0 = p0
                    p0 -= 100
                    switch fi21 {
                    case 128:
                        switch "foo bar" {
                        case "foo bar":
                            switch 0 {
                            case 0:
                                if false {
                                    p0 += (-(p0))
                                } else {
                                    var v22 uint = 18446744073709551615
                                    _ = v22
                                }
                            case 128:
                                p2 += fi20
                            case -2147483648:
                                p0 += (p0 % ((p0) | 1))
                            default:
                                switch "" {
                                case "":
                                    p2 = (^(fi21))
                                case "hello":
                                    v23 := []int{}
                                    _ = v23
                                case "foo bar":
                                    v24 := GenS2{F0: 4294967295}
                                    _ = v24
                                }
                            }
                        }
                    case 263:
                        v25 := GenS3{F0: 32767}
                        _ = v25
                    case 2147483647:
                        v26 := []int{(p2 | 122), (^(p2)), 0, (-(fi20))}
                        _ = v26
                    }
                }
                var v27 uint32 = 4294967295
                v28 := []int{(-(fi20))}
                var v29 uint32 = v27
                p1 = p1
                _ = fi21
                _ = v27
                _ = v28
                _ = v29
            }
            fmt.Println(p1, p0, fi20)
            var v30 string = "a"
            _ = fi20
            _ = v30
        }
        var v31 uint = 65535
        _ = v31
    }
    v32 := []int{p2, p2, p2, -9223372036854775808}
    _ = len(v32)
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return (p2 + (^(p2)))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    var v4 int = 2147483647
    v5 := make(map[int]uint32)
    var v6 int = (v4 & 32768)
    for fi33 := 0; fi33 < 17; fi33++ {
        if true {
            var v34 int8 = -38
            var v35 string = ("a" + "a")
            v34 -= v34
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    v6 += v4
    v36 := []int{v4}
    switch "" {
    case "a":
        if (!false) {
            var v37 float32 = 0.1
            var v38 float64 = (-2.25 / 1.5)
            _ = v37
            _ = v38
        }
    default:
        v39 := GenS2{F0: 65535}
        _ = v39
    }
    var v40 string = ((("foo bar" + "") + "\n\t") + ((("foo bar" + "\n\t") + ("a" + "a")) + "a"))
    v4 = 1
    if ((!true) && true) {
        v40 = v40
        var v41 bool = (false && true)
        fmt.Println(v4, v41)
        _ = v41
    } else {
        v42 := GenS1{F0: ("foo bar" + (("a" + "\n\t") + (((("hello" + v40) + v40) + v40) + (("\n\t" + ("hello" + "hello")) + v40)))), F1: -94, F2: (v4 + genFn1(7, -237, (v4 << ((65535) & 63))))}
        v40 = v40
        _ = v42
    }
    var v43 float32 = (1e10 * 0.0)
    v2 = (v2 >> ((255) & 63))
    var v44 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
