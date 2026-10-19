package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 string, p1 int32) uint32 {
    p1 = p1
    p0 = ((("\n\t" + "hello") + (p0 + "\n\t")) + (p0 + (("foo bar" + (p0 + (p0 + "hello"))) + p0)))
    switch "foo bar" {
    case "hello":
        var v6 int16 = 1
        _ = v6
    default:
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(p1)
            p1 -= p1
            switch (fi7 * fi7) {
            case -776:
                v8 := []int{}
                _ = v8
            }
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1() uint {
    var v9 bool = false
    v10 := GenS1{F0: v9, F1: 2}
    fmt.Println(v9)
    _ = v10
    _ = v9
    return 659
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 1e10
    var v3 string = "hello"
    var v4 float32 = (float32(float64(v2)) + v2)
    var v5 uint64 = 448
    for fi11 := 0; fi11 < 16; fi11++ {
        v3 = "a"
        if false {
            v12 := GenS2{F0: 0, F1: v3, F2: v5}
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: (!true), F1: genFn1()}
                if false {
                    switch "hello" {
                    case "hello":
                        var v15 int8 = 0
                        _ = v15
                    case "foo bar":
                        switch ((((v3 + "\n\t") + "hello") + v3) + "a") {
                        case "foo bar":
                            v16 := GenS3{F0: (v5 + v5), F1: (^(v5))}
                            _ = v16
                        }
                    case "":
                        if (true || (1 < 791)) {
                            v3 = v3
                        } else {
                            var v17 int = fi11
                            v4 = v2
                            _ = v17
                        }
                    default:
                        if (!true) {
                            _ = len(v1)
                            if ((v3 < (v3 + v3)) && false) {
                                fmt.Println(v2)
                            }
                            v18 := []int{127, (fi11 + -837), fi13, (fi13 << ((0) & 63))}
                            var v19 float32 = (0.1 - 0.1)
                            v1[938] = v3
                            v4 = ((v4 + (v19 * (v19 - v2))) + (1e10 + (((float32(((0.1 - 0.1) + 3.14159)) + v4) / (0.1 / 3.14159)) - (((v19 - (v4 + 3.14159)) + 1e10) + v19))))
                            fmt.Println(v3, v5, v19)
                            _ = v18
                            _ = v19
                        } else {
                            switch fi13 {
                            case -314:
                                v3 = v3
                            case 914:
                                v5 += v5
                            case 32768:
                                v5 += 2
                            default:
                                _ = len(v1)
                            }
                            delete(v1, 957)
                            v20 := []int{(fi11 / ((fi11) | 1)), fi13, 127, (fi13 >> ((704) & 63))}
                            _ = v20
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            switch (v3 + "") {
            case "":
                switch v3 {
                case "\n\t":
                    v2 += float32((1.5 / 1e10))
                case "foo bar":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 uint64 = (^(v5))
                        _ = fi21
                        _ = v22
                    }
                }
            case "\n\t":
                var v23 uint64 = v5
                _ = v23
            case "hello":
                v4 = 3.14159
            default:
                if (!true) {
                    fmt.Println(v2, fi11)
                }
            }
            _ = v12
        }
        v24 := GenS1{F0: (396 <= 42), F1: 0}
        _ = fi11
        _ = v24
    }
    v4 = 1.5
    v2 = float32(0.1)
    v2 = (0.1 - v2)
    fmt.Println(v4, v2)
    var v25 string = "a"
    if (!((0 <= 0) && true)) {
        v4 -= v4
    }
    switch 690 {
    case 32768:
        var v26 uint64 = v5
        _ = v26
    case 32767:
        var v27 int32 = -2147483648
        _ = v27
    case 46:
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := []int{(-(fi29)), fi29}
                var v31 uint8 = 2
                fmt.Println(v25, fi29)
                v5 += 159
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 int32 = 1
            var v33 int16 = 267
            _ = fi28
            _ = v32
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
