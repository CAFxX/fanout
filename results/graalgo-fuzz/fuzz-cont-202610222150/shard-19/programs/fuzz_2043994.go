package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 int64) string {
    v6 := GenS1{F0: "a"}
    p0 += p0
    _ = p0
    _ = v6
    return ""
}
func genFn1(p0 uint32) int {
    var v7 int16 = 42
    if false {
        switch "" {
        case "foo bar":
            for fi8 := 0; fi8 < 6; fi8++ {
                if true {
                    if (2147483647 <= (fi8 % ((fi8) | 1))) {
                        v9 := GenS2{F0: 2, F1: fi8}
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v7 += (v7 * v7)
                            _ = fi10
                        }
                        v11 := GenS1{F0: "\n\t"}
                        _ = v11
                        _ = v9
                    }
                    fmt.Println(fi8, p0, fi8)
                    switch 127 {
                    case 4294967295:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            switch "a" {
                            case "foo bar":
                                var v13 uint8 = 2
                                _ = v13
                            default:
                                if false {
                                    fmt.Println(fi8, v7, p0)
                                    v7 += (-(v7))
                                    v14 := GenS1{F0: "\n\t"}
                                    _ = v14
                                }
                            }
                            _ = fi12
                        }
                    default:
                        v15 := []int{(fi8 ^ fi8), -2147483648}
                        _ = v15
                    }
                }
                var v16 uint16 = 65535
                v17 := GenS3{F0: fi8, F1: -1}
                _ = fi8
                _ = v16
                _ = v17
            }
        case "":
            fmt.Println(p0, v7)
        case "a":
            v18 := []int{32768, 32767}
            _ = v18
        }
    }
    _ = p0
    _ = v7
    return int(p0)
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 285
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    var v19 int8 = 1
    var v20 uint16 = 1
    v1 += (v1 - 1)
    fmt.Println(v20)
    if false {
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS1{F0: ("\n\t" + "\n\t")}
                var v24 uint64 = 739
                if ((v1 >> ((uint(v24)) & 63)) <= v1) {
                    if ("a" != "") {
                        v25 := GenS2{F0: 235, F1: (fi22 + genFn1(65535))}
                        var v26 float32 = (3.14159 + (1.5 / (1e10 / 1.5)))
                        v20 = (^(v20))
                        var v27 int = (fi22 & genFn1(42))
                        var v28 float64 = 0.1
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            v20 += v20
            fmt.Println(v1)
            _ = fi21
        }
        switch "foo bar" {
        case "foo bar":
            v29 := GenS3{F0: -957, F1: -1}
            _ = v29
        case "hello":
            v30 := []int{-1, 4294967295}
            _ = v30
        case "a":
            v31 := v3["a"]
            _ = v31
        }
        v32 := []int{32767}
        _ = v32
    }
    fmt.Println(v1)
    var v33 uint64 = 2
    var v34 uint32 = 2147483647
    v1 -= v1
    var v35 float32 = ((((1.5 - 3.14159) / float32(0.1)) + 0.1) + (((1.5 + ((float32((-2.25 + float64((0.0 * 0.1)))) - 0.1) / 1.0)) * float32(3.14159)) / ((-2.25 * 1e10) / ((-2.25 - 1e10) + (0.1 - float32(1.5))))))
    v34 += v34
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 int32 = int32(v34)
        _ = fi36
        _ = v37
    }
    var v38 float64 = 0.0
    if false {
        v34 -= (v34 >> ((2) & 63))
        v1 += (v1 ^ (v1 / ((-407) | 1)))
        for fi39 := 0; fi39 < 6; fi39++ {
            v19 += (v19 ^ -128)
            _ = fi39
        }
    }
    var v40 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
