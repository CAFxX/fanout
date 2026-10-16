package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 bool, p2 int8) int32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        p0 = p0
        for fi5 := 0; fi5 < 18; fi5++ {
            if (866 <= 324) {
                switch (fi4 >> ((18446744073709551615) & 63)) {
                case 128:
                    if false {
                        p1 = true
                    }
                case -1:
                    var v6 int64 = -9223372036854775808
                    _ = v6
                case 1:
                    p2 = p2
                default:
                    switch "\n\t" {
                    case "\n\t":
                        if (!p1) {
                            var v7 bool = (2 >= 2)
                            if v7 {
                                p0 = (p0 * (p0 ^ (^(p0))))
                                switch (("\n\t" + ("a" + "foo bar")) + "") {
                                case "hello":
                                    var v8 int32 = -1
                                    _ = v8
                                case "\n\t":
                                    var v9 uint = 0
                                    _ = v9
                                case "foo bar":
                                    var v10 int32 = 2147483647
                                    _ = v10
                                }
                                if p1 {
                                    v11 := []int{fi4}
                                    _ = v11
                                } else {
                                    switch fi5 {
                                    case 1:
                                        if false {
                                            fmt.Println(p1)
                                            p2 = -1
                                            p0 -= 42
                                            v7 = ("hello" < "hello")
                                            v12 := GenS1{F0: 40}
                                            if true {
                                                p0 += p2
                                                var v13 int64 = 9223372036854775807
                                                _ = v13
                                            }
                                            switch fi4 {
                                            case -2147483648:
                                                v14 := []int{-378, 32767}
                                                _ = v14
                                            default:
                                                if p1 {
                                                    p1 = (v7 || v7)
                                                    v15 := GenS2{F0: 190, F1: 1}
                                                    _ = v15
                                                }
                                            }
                                            _ = v12
                                        }
                                    case -9223372036854775808:
                                        p2 -= (p2 % (((^(p2))) | 1))
                                    default:
                                        p1 = v7
                                    }
                                }
                            }
                            switch "hello" {
                            case "\n\t":
                                v16 := []int{}
                                v16 = append(v16, (fi4 * -335))
                                var v17 uint32 = 65535
                                _ = v16
                                _ = v17
                            }
                            _ = v7
                        } else {
                            if (!p1) {
                                p1 = true
                            } else {
                                p0 = p0
                            }
                            p0 = (p0 % ((p0) | 1))
                            v18 := []int{-930, -602}
                            _ = v18
                        }
                    case "hello":
                        p1 = (48 < 10)
                    case "a":
                        v19 := GenS1{F0: 42}
                        _ = v19
                    }
                }
            }
            _ = fi5
        }
        if p1 {
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int8 = 1
                v22 := GenS1{F0: 91}
                _ = fi20
                _ = v21
                _ = v22
            }
            if (42 != 80) {
                v23 := GenS2{F0: 175, F1: 2147483647}
                v24 := []int{fi4, fi4, (-(fi4))}
                _ = v23
                _ = v24
            }
            var v25 int32 = 952
            _ = v25
        }
        _ = fi4
    }
    var v26 float32 = (1.5 + -2.25)
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 int16 = -1
        fmt.Println(v26, v28)
        _ = fi27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return -1
}
func genFn1(p0 int64, p1 int32, p2 int8) uint16 {
    v29 := GenS1{F0: 1}
    fmt.Println(p1)
    if (42 != 126) {
        if (1.5 > -2.25) {
            fmt.Println(p1, p0, p0)
            var v30 uint = 735
            _ = v30
        }
        p2 -= 91
        v31 := GenS2{F0: 2, F1: p1}
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    v3 := []string{("\n\t" + v2), v2, v2}
    fmt.Println(v2)
    var v32 uint32 = 2147483647
    v32 += v32
    var v33 int16 = -1
    v33 = (-(v33))
    v2 = (v2 + ((v2 + ("" + "foo bar")) + ""))
    if (0.1 != -2.25) {
        v33 = v33
        v34 := GenS3{F0: 226}
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 uint = 255
        if true {
            v36 -= (v36 >> ((2) & 63))
            v32 = v32
            var v37 uint32 = (v32 ^ (v32 * (v32 % (((v32 * v32)) | 1))))
            _ = v37
        } else {
            fmt.Println(v2, v36)
        }
        fmt.Println(v32, fi35, v2)
        _ = fi35
        _ = v36
    }
    var v38 int = 127
    if ((1e10 + 1.5) > -2.25) {
        v39 := GenS4{F0: 0}
        var v40 uint32 = v32
        _ = v39
        _ = v40
    }
    var v41 uint = 65535
    v2 = (((v2 + "") + v2) + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v41
}
