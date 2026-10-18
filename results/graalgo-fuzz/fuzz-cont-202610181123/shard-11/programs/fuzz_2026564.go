package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0() uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{(fi4 / (((fi4 ^ fi4)) | 1)), fi4}
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 bool = (255 > 198)
            v5 = append(v5, (fi6 | (-(fi4))))
            v8 := GenS1{F0: (fi6 & fi6)}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int = 0
        _ = fi4
        _ = v5
        _ = v9
    }
    v10 := []int{128, 1, 0, 1}
    v11 := GenS2{F0: 0, F1: ((!true) && true), F2: true}
    _ = v10
    _ = v11
    return 65535
}
func genFn1(p0 uint16) int16 {
    if (162 == 42) {
        switch (("\n\t" + "\n\t") + "hello") {
        case "\n\t":
            v12 := GenS1{F0: -517}
            _ = v12
        }
    } else {
        p0 += p0
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint8 = 127
            switch -1 {
            case 4294967295:
                p0 += (p0 | 170)
            case -298:
                switch "a" {
                case "\n\t":
                    if (0 > (p0 * (p0 ^ 0))) {
                        var v15 uint64 = 0
                        fmt.Println(p0, p0, p0)
                        if ((fi13 >> ((621) & 63)) < (fi13 * (fi13 & (fi13 - fi13)))) {
                            fmt.Println(v15, v15, fi13)
                            v15 = (v15 - v15)
                            var v16 int32 = 42
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 17; fi17++ {
                                v18 := []int{fi13, (fi13 & -248), (^(fi13)), (fi17 + (fi13 & 32767))}
                                if (-1 <= 0) {
                                    var v19 string = "\n\t"
                                    var v20 int16 = 32767
                                    _ = v19
                                    _ = v20
                                } else {
                                    v21 := GenS3{F0: 42, F1: p0}
                                    _ = v21
                                }
                                fmt.Println(v15, p0)
                                v15 += v15
                                _ = fi17
                                _ = v18
                            }
                            v15 -= v15
                        }
                        _ = v15
                    }
                case "foo bar":
                    fmt.Println(p0, p0)
                case "":
                    v14 -= (v14 - 128)
                }
            default:
                if true {
                    v22 := []int{fi13}
                    var v23 int32 = 48
                    _ = v22
                    _ = v23
                } else {
                    v24 := []int{(fi13 | (-(fi13))), fi13}
                    _ = v24
                }
            }
            _ = fi13
            _ = v14
        }
    }
    switch "hello" {
    case "a":
        v25 := []int{32768, 9223372036854775807}
        _ = v25
    case "hello":
        p0 += (^(p0))
    case "foo bar":
        switch 477 {
        case 291:
            var v26 int8 = 64
            _ = v26
        case 32768:
            v27 := []int{}
            _ = v27
        case -1:
            v28 := []int{}
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 float32 = 0.1
            v31 := GenS1{F0: fi29}
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v32 string = "foo bar"
    var v33 int64 = 42
    if true {
        var v34 uint8 = 42
        fmt.Println(v33, v32)
        _ = v34
    } else {
        v35 := GenS2{F0: 9223372036854775807, F1: false, F2: (0.1 >= 1.5)}
        v36 := GenS4{F0: 2, F1: 840, F2: 0}
        if true {
            if ((!(953 <= 399)) || (("hello" < ("foo bar" + (v32 + v32))) && true)) {
                v33 -= v33
            }
            v37 := []int{32768, 32768, 2147483647}
            _ = v37
        }
        _ = v35
        _ = v36
    }
    v38 := []int{-163, -2147483648, 127, 127}
    var v39 int = -938
    fmt.Println(v39)
    if true {
        fmt.Println(v33, v32, v39)
        v40 := GenS5{F0: 42, F1: false, F2: 2}
        if len(v38) >= 2 {
            v38 = v38[1:len(v38)]
        }
        _ = v40
    }
    v33 += (^(v33))
    switch (v39 * -1) {
    case 42:
        var v41 string = ((v32 + (v32 + (v32 + "\n\t"))) + v32)
        _ = v41
    case -2147483648:
        v42 := GenS1{F0: 223}
        _ = v42
    default:
        if true {
            v32 = "foo bar"
            for fi43 := 0; fi43 < 4; fi43++ {
                v33 += (-(v33))
                var v44 int8 = 42
                if false {
                    delete(v3, 965)
                } else {
                    var v45 int64 = 1
                    v38 = append(v38, v39)
                    _ = v45
                }
                _ = fi43
                _ = v44
            }
        }
    }
    v1["foo bar"] = v39
    v46 := v2[v39]
    v33 = (^(v33))
    var v47 uint8 = 128
    fmt.Println(v46, v39, v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v46
    _ = v47
}
