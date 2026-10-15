package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint8) uint8 {
    p0 = (p0 & (p0 % ((p0) | 1)))
    v7 := GenS1{F0: "foo bar", F1: 0}
    var v8 float32 = float32((0.1 + -2.25))
    var v9 int64 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (p0 % ((175) | 1))
}
func genFn1(p0 int, p1 string) int {
    var v10 uint32 = 0
    switch p0 {
    case -2147483648:
        switch ((p1 + "foo bar") + "foo bar") {
        case "foo bar":
            var v11 uint32 = 2
            _ = v11
        default:
            v12 := []int{}
            _ = v12
        }
    default:
        if true {
            if ((^(p0)) > (p0 >> ((240) & 63))) {
                p0 -= 32767
            }
        }
    }
    fmt.Println(v10, p1, p0)
    v13 := GenS1{F0: p1, F1: v10}
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    return p0
}

func main() {
    var v1 float32 = 1e10
    var v2 string = "a"
    var v3 int16 = 1
    v4 := make(map[uint32]int)
    var v5 uint64 = 116
    var v6 int16 = (v3 % ((42) | 1))
    fmt.Println(v3, v5, v2)
    var v14 float32 = ((1e10 * v1) / v1)
    v15 := GenS1{F0: v2, F1: 441}
    v14 = float32(1e10)
    switch 4294967295 {
    case -702:
        var v16 float32 = -2.25
        _ = v16
    case -914:
        fmt.Println(v5, v1)
    case 1:
        fmt.Println(v1, v1, v6)
    default:
        switch 42 {
        case -1:
            if (255 <= 146) {
                if (42 > int64(v3)) {
                    var v17 int32 = 730
                    v6 = v3
                    v18 := []int{127, 32768}
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v2, v14)
                    if (!(!(-1 >= 42))) {
                        var v19 uint64 = 4294967295
                        var v20 string = v2
                        _ = v19
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 7; fi21++ {
                            switch "" {
                            case "":
                                var v22 uint = 901
                                _ = v22
                            case "a":
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    var v24 uint = 4294967295
                                    _ = fi23
                                    _ = v24
                                }
                            case "hello":
                                fmt.Println(v2)
                            default:
                                switch genFn1((^(fi21)), v2) {
                                case 128:
                                    var v25 int = (fi21 ^ 128)
                                    _ = v25
                                case -885:
                                    v26 := []int{-676, (^(fi21)), (^(fi21))}
                                    _ = v26
                                }
                            }
                            v27 := []int{(fi21 & (fi21 >> ((2) & 63))), (fi21 >> ((46) & 63))}
                            _ = fi21
                            _ = v27
                        }
                        var v28 uint = 2
                        _ = v28
                    }
                }
            }
        case 127:
            v29 := []int{32768, -1}
            _ = v29
        }
    }
    switch "" {
    case "hello":
        var v30 uint32 = 2
        _ = v30
    case "\n\t":
        delete(v4, 692)
    case "a":
        if (!true) {
            var v31 float64 = (-2.25 * (0.0 * 0.1))
            var v32 int16 = v3
            var v33 uint32 = 674
            _ = v31
            _ = v32
            _ = v33
        } else {
            var v34 uint = 65535
            var v35 int64 = 9223372036854775807
            v36 := []int{1}
            _ = v34
            _ = v35
            _ = v36
        }
    default:
        v5 -= v5
    }
    switch "\n\t" {
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        v37 := GenS2{F0: 2, F1: 42}
        _ = v37
    }
    var v38 string = "hello"
    var v39 uint32 = 1
    switch -1 {
    case 646:
        v5 = (^(v5))
    case -1:
        if true {
            fmt.Println(v2, v3)
            for fi40 := 0; fi40 < 17; fi40++ {
                v41 := []int{genFn1((fi40 + (^(fi40))), v38), genFn1(fi40, (v2 + (v2 + "a"))), genFn1(fi40, "a")}
                var v42 uint16 = 0
                _ = fi40
                _ = v41
                _ = v42
            }
        }
    case 2147483647:
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 uint32 = v39
            v45 := []int{(-(fi43))}
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    v46 := GenS2{F0: 32767, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
