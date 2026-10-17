package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint) uint32 {
    var v5 float64 = 0.0
    v5 = float64(0.0)
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            v7 := []int{(fi6 << (((^(p1))) & 63)), (fi6 << (((p1 >> (((p1 / ((1) | 1))) & 63))) & 63))}
            var v8 string = ("foo bar" + "foo bar")
            _ = v7
            _ = v8
        } else {
            switch "" {
            case "\n\t":
                fmt.Println(fi6, p1)
            case "a":
                fmt.Println(p0, p0, p0)
            default:
                var v9 uint64 = 42
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int64 = -9223372036854775808
                if true {
                    var v12 int = -1
                    v12 -= (fi6 * 2147483647)
                    _ = v12
                } else {
                    var v13 float32 = -2.25
                    if true {
                        if false {
                            var v14 int64 = (v11 >> ((42) & 63))
                            _ = v14
                        } else {
                            var v15 string = ((("\n\t" + "a") + ("" + "a")) + ("\n\t" + "foo bar"))
                            p1 += 539
                            _ = v15
                        }
                    } else {
                        if false {
                            v16 := GenS1{F0: "\n\t", F1: 32767}
                            fmt.Println(fi6)
                            _ = v16
                        } else {
                            var v17 bool = ((!((v11 > 1) && false)) || (549 < 42))
                            fmt.Println(v5)
                            _ = v17
                        }
                        v18 := []int{fi10}
                        _ = v18
                    }
                    var v19 int64 = 42
                    _ = v13
                    _ = v19
                }
                switch (fi10 ^ (fi6 % ((fi6) | 1))) {
                case 553:
                    fmt.Println(fi6)
                case 815:
                    fmt.Println(v11, fi10)
                case 2147483647:
                    var v20 int8 = 57
                    _ = v20
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi6
    }
    v5 -= float64(-2.25)
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 uint8 = 1
    v4 := make(map[uint32]string)
    v3 = v3
    if (!false) {
        fmt.Println(v3, v3)
        v21 := GenS2{F0: 0, F1: 100, F2: 914}
        var v22 uint32 = 685
        _ = v21
        _ = v22
    }
    if true {
        v3 = 42
        switch "a" {
        case "\n\t":
            if ((678 < 32767) || (v3 != (v3 ^ (v3 | v3)))) {
                v3 = (v3 >> ((255) & 63))
            }
        default:
            var v23 int = 127
            _ = v23
        }
    } else {
        v24 := []int{0, 685, 1}
        _ = v24
    }
    if true {
        v25 := []int{9223372036854775807, 42, 4294967295, -9223372036854775808}
        v26 := GenS3{F0: 523, F1: v3}
        var v27 string = ((("" + ("\n\t" + ((("" + "") + "foo bar") + "hello"))) + "hello") + ("" + "a"))
        _ = v25
        _ = v26
        _ = v27
    }
    switch "hello" {
    case "foo bar":
        if true {
            v28 := GenS2{F0: 584, F1: 0, F2: -1}
            v3 = v3
            _ = v28
        }
    default:
        delete(v1, -569)
    }
    v29 := v4[2]
    switch v29 {
    case "hello":
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 uint16 = 65535
            fmt.Println(v3, v29, v29)
            switch "foo bar" {
            case "a":
                switch v29 {
                case "hello":
                    var v32 string = (v29 + "hello")
                    _ = v32
                case "a":
                    v33 := GenS2{F0: 18446744073709551615, F1: -1, F2: 112}
                    _ = v33
                default:
                    v34 := []int{}
                    _ = v34
                }
            default:
                for fi35 := 0; fi35 < 17; fi35++ {
                    v36 := GenS2{F0: 42, F1: -128, F2: 3}
                    var v37 string = "a"
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            _ = fi30
            _ = v31
        }
    case "\n\t":
        fmt.Println(v29)
    case "":
        var v38 int32 = -1
        _ = v38
    default:
        v39 := []int{1, 4294967295}
        _ = v39
    }
    v40 := []int{}
    var v41 uint8 = v3
    v42 := v1[2147483647]
    var v43 int32 = -2147483648
    var v44 bool = (!v42)
    fmt.Println(v41)
    v45 := GenS4{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
