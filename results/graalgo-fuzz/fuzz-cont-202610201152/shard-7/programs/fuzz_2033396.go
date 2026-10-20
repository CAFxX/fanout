package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 int, p2 bool) uint16 {
    v5 := GenS1{F0: 0, F1: 641, F2: 42}
    var v6 int8 = -56
    var v7 uint = uint(p0)
    var v8 uint16 = 353
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (^(v8))
}
func genFn1(p0 int64, p1 bool, p2 string) int64 {
    p2 = p2
    var v9 int32 = 2147483647
    v10 := GenS2{F0: 0}
    var v11 bool = (((p2 + (p2 + p2)) + p2) != ((((p2 + "") + p2) + ("foo bar" + p2)) + (p2 + p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = ""
    var v2 int16 = 1
    var v3 int8 = -1
    var v4 int16 = -1
    if false {
        var v12 uint8 = 255
        for fi13 := 0; fi13 < 10; fi13++ {
            v4 += v4
            var v14 uint32 = 61
            var v15 int64 = -872
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        var v16 int16 = 569
        for fi17 := 0; fi17 < 17; fi17++ {
            v1 = v1
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 uint16 = 65535
                _ = fi18
                _ = v19
            }
            var v20 int16 = (v16 ^ (v16 * v4))
            _ = fi17
            _ = v20
        }
        _ = v16
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        switch v1 {
        case "foo bar":
            for fi22 := 0; fi22 < 6; fi22++ {
                switch v1 {
                case "hello":
                    var v23 float64 = (float64(float32((1e10 * 1.5))) - (0.1 + 0.0))
                    _ = v23
                case "foo bar":
                    v1 = ""
                }
                var v24 uint16 = 0
                _ = fi22
                _ = v24
            }
        case "hello":
            switch fi21 {
            case 2147483647:
                if (22 > 255) {
                    var v25 uint16 = genFn0(83, (fi21 + fi21), (1.5 != 3.14159))
                    var v26 uint = 42
                    _ = v25
                    _ = v26
                } else {
                    switch (("\n\t" + v1) + (v1 + "")) {
                    case "foo bar":
                        fmt.Println(v3, fi21, v3)
                    }
                    switch (fi21 | (-(fi21))) {
                    case 784:
                        var v27 bool = true
                        _ = v27
                    case 128:
                        switch v1 {
                        case "a":
                            v28 := GenS3{F0: 1, F1: fi21, F2: v3}
                            _ = v28
                        case "foo bar":
                            if true {
                                v29 := GenS1{F0: (v3 ^ (-(v3))), F1: 0, F2: 2147483647}
                                _ = v29
                            } else {
                                v30 := []int{(fi21 % (((fi21 % (((fi21 % ((fi21) | 1))) | 1))) | 1)), -954}
                                if ((-(fi21)) >= int(v2)) {
                                    v1 = v1
                                }
                                var v31 int = fi21
                                _ = v30
                                _ = v31
                            }
                        case "hello":
                            v2 = (v2 - v2)
                        default:
                            for fi32 := 0; fi32 < 4; fi32++ {
                                fmt.Println(fi32)
                                _ = fi32
                            }
                        }
                    default:
                        if ("foo bar" >= "a") {
                            v4 += 748
                            v33 := GenS1{F0: (v3 >> ((427) & 63)), F1: 2147483648, F2: 0}
                            if true {
                                var v34 string = "a"
                                _ = v34
                            } else {
                                for fi35 := 0; fi35 < 5; fi35++ {
                                    var v36 uint32 = 1
                                    v4 = (v2 | (v4 % ((v2) | 1)))
                                    var v37 int16 = 119
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                v4 -= (-(v2))
                            }
                            _ = v33
                        }
                    }
                    var v38 int16 = (v4 - (-(v2)))
                    _ = v38
                }
            }
        }
        v39 := []int{(fi21 % ((fi21) | 1)), (fi21 * fi21)}
        _ = fi21
        _ = v39
    }
    var v40 int64 = 1
    v3 = (v3 - -90)
    var v41 bool = ((524 == 1) || (!false))
    var v42 uint64 = 128
    v43 := []int{-346, -665}
    v43 = append(v43, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
