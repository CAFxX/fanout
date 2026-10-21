package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 float32 = float32(((1e10 * 3.14159) + 3.14159))
    var v2 int8 = -128
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint8 = 194
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint16 = 575
            var v8 uint64 = 0
            v9 := GenS1{F0: 1, F1: 100}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 float32 = ((-2.25 * float32(float64((0.1 / v1)))) - (v1 - float32(v3)))
        _ = fi4
        _ = v10
        _ = v5
    }
    var v11 float64 = -2.25
    fmt.Println(v3, v11, v3)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{fi12, fi12, (fi12 % ((fi12) | 1))}
        v14 := GenS2{F0: 4294967295}
        v15 := GenS3{F0: "\n\t", F1: (-(v2))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 uint32 = 42
    var v17 bool = ((746 < 1) && (true && false))
    var v18 string = "foo bar"
    for fi19 := 0; fi19 < 3; fi19++ {
        v18 = "\n\t"
        switch (v18 + v18) {
        case "foo bar":
            fmt.Println(v3, v11)
        case "\n\t":
            v20 := GenS2{F0: 2}
            _ = v20
        case "hello":
            if v17 {
                v21 := []int{(fi19 - fi19), (fi19 ^ fi19), 42}
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 string = ((v18 + v18) + ("hello" + "foo bar"))
                    var v24 uint8 = 127
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            }
        default:
            if (!true) {
                v25 := GenS1{F0: 42, F1: -128}
                for fi26 := 0; fi26 < 20; fi26++ {
                    fmt.Println(v17, v16)
                    if false {
                        v16 = 282
                        fmt.Println(v3, fi26, v17)
                        switch (fi26 ^ (fi19 >> ((0) & 63))) {
                        case -567:
                            var v27 float32 = v1
                            var v28 bool = v17
                            if (fi19 == fi26) {
                                if false {
                                    v29 := []int{-776, fi26}
                                    v30 := GenS1{F0: 0, F1: (v2 % ((v2) | 1))}
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := []int{fi19, fi26, fi26, 0}
                                    _ = v31
                                }
                                v32 := []int{fi26}
                                _ = v32
                            } else {
                                v33 := []int{(fi19 * fi19), -816, (fi19 - 128), 128}
                                v34 := GenS3{F0: v18, F1: v2}
                                _ = v33
                                _ = v34
                            }
                            var v35 int32 = 2147483647
                            _ = v27
                            _ = v28
                            _ = v35
                        case -2147483648:
                            var v36 int64 = 1
                            if (65535 <= 42) {
                                v37 := []int{fi26, (fi26 | 32767), (fi19 + fi26), fi26}
                                _ = v37
                            }
                            v38 := []int{fi26, 4294967295, 32768}
                            switch ("hello" + v18) {
                            case "":
                                switch v18 {
                                case "\n\t":
                                    v16 -= 4294967295
                                }
                            case "\n\t":
                                v39 := GenS4{F0: (!(v18 > v18))}
                                _ = v39
                            }
                            var v40 uint = 2
                            _ = v36
                            _ = v38
                            _ = v40
                        }
                    } else {
                        fmt.Println(v11, v18)
                    }
                    _ = fi26
                }
                _ = v25
            }
        }
        _ = fi19
    }
    v2 = -121
    var v41 int32 = 0
    var v42 uint64 = 0
    v43 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
