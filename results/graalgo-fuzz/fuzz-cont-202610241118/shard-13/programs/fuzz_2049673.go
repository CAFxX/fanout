package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() int8 {
    if false {
        var v6 uint16 = 101
        if (235 > 2) {
            v7 := GenS1{F0: 814}
            var v8 int = 42
            if (!(true || false)) {
                if ((0.0 + (-2.25 - (0.1 + 1e10))) == 0.0) {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v6, v8)
                    case "\n\t":
                        v6 += v6
                    }
                    v9 := GenS1{F0: -2147483648}
                    v10 := []int{(v8 - 128), (v8 / ((9223372036854775807) | 1))}
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{v8}
                    v12 := GenS2{F0: 577, F1: false}
                    var v13 int = v8
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := GenS3{F0: (!true), F1: (v6 + 743)}
                v15 := GenS1{F0: v8}
                _ = v14
                _ = v15
            }
            _ = v7
            _ = v8
        } else {
            if true {
                v6 = v6
                v16 := []int{-1, 0, 573}
                var v17 int32 = 753
                _ = v16
                _ = v17
            } else {
                v6 += v6
                v18 := GenS1{F0: -1}
                _ = v18
            }
            var v19 uint64 = 2
            fmt.Println(v6, v19)
            _ = v19
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint16 = 32767
                v22 += v6
                var v23 float64 = 3.14159
                _ = fi21
                _ = v22
                _ = v23
            }
            v6 -= 614
            _ = fi20
        }
        _ = v6
    } else {
        var v24 uint8 = 0
        v25 := []int{}
        v24 -= (v24 % (((v24 ^ v24)) | 1))
        _ = v24
        _ = v25
    }
    v26 := []int{}
    var v27 float64 = 0.1
    _ = v26
    _ = v27
    return 1
}
func genFn1(p0 string, p1 string) int32 {
    switch (((p0 + (("foo bar" + p1) + (p0 + p1))) + p1) + ("" + (p0 + p1))) {
    case "a":
        if false {
            fmt.Println(p0, p1)
        }
    case "\n\t":
        var v28 uint16 = 539
        _ = v28
    case "":
        p0 = "foo bar"
    }
    var v29 bool = false
    p0 = ((p0 + p0) + p1)
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS1{F0: (fi30 % ((fi30) | 1))}
        var v32 int32 = -468
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = v29
    return -2147483648
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 float64 = float64(float32((0.0 - 1.5)))
    var v4 int64 = -268
    v5 := make(map[int]bool)
    _ = len(v5)
    for fi33 := 0; fi33 < 2; fi33++ {
        v1 = append(v1, fi33)
        _ = fi33
    }
    fmt.Println(v3, v3)
    switch -1 {
    case 32767:
        for fi34 := 0; fi34 < 16; fi34++ {
            if ((1e10 + v3) != ((float64((-2.25 / 1.5)) * 3.14159) + 1e10)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v1)
                var v35 int16 = -32768
                _ = v35
            }
            if len(v1) > 0 {
                v1[0] = fi34
            }
            _ = fi34
        }
    case 42:
        var v36 uint8 = 42
        _ = v36
    }
    v3 -= v3
    v37 := GenS2{F0: -458, F1: ((42 > 764) || false)}
    var v38 uint32 = 2
    var v39 float32 = 0.1
    if false {
        _ = len(v1)
        v38 += 752
    } else {
        v40 := GenS2{F0: 1, F1: true}
        _ = v40
    }
    fmt.Println(v38, v3)
    v41 := GenS3{F0: false, F1: 341}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
