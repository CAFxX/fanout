package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16, p1 int, p2 int) int64 {
    if true {
        for fi4 := 0; fi4 < 3; fi4++ {
            var v5 float64 = (1.5 - -2.25)
            _ = fi4
            _ = v5
        }
        switch "a" {
        case "a":
            v6 := GenS1{F0: 1, F1: true}
            _ = v6
        case "":
            v7 := GenS2{F0: 431}
            _ = v7
        case "\n\t":
            var v8 uint8 = 108
            _ = v8
        }
        p1 = p1
    } else {
        p1 += p2
        if true {
            if true {
                var v9 uint64 = 42
                fmt.Println(p2, p1, p1)
                if (!false) {
                    var v10 uint64 = 0
                    fmt.Println(p1, p2)
                    var v11 uint32 = 0
                    _ = v10
                    _ = v11
                } else {
                    if (0 == 0) {
                        v12 := []int{p1, (p2 * 541), p2, 953}
                        _ = v12
                    }
                }
                _ = v9
            } else {
                v13 := GenS1{F0: 702, F1: (18446744073709551615 < 0)}
                if false {
                    if (false && true) {
                        var v14 int = p2
                        var v15 float32 = 3.14159
                        if (((v15 / float32((0.0 - -2.25))) - ((v15 + 1.5) + float32(3.14159))) > v15) {
                            p2 += (p1 ^ (^(v14)))
                        }
                        _ = v14
                        _ = v15
                    } else {
                        v16 := GenS3{F0: "hello", F1: -357}
                        var v17 int32 = 2147483647
                        for fi18 := 0; fi18 < 12; fi18++ {
                            v19 := GenS4{F0: 128, F1: -51}
                            var v20 float64 = 0.1
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v16
                        _ = v17
                    }
                    if (0 > -9223372036854775808) {
                        for fi21 := 0; fi21 < 16; fi21++ {
                            fmt.Println(fi21, p0, p0)
                            _ = fi21
                        }
                        var v22 string = "foo bar"
                        _ = v22
                    }
                    var v23 uint64 = 828
                    _ = v23
                }
                v24 := []int{-835, (p2 + (^(p2))), (p1 / (((p2 ^ 2147483647)) | 1)), -652}
                _ = v13
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint8 = 1
            fmt.Println(v26)
            v27 := GenS5{F0: -629}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        if false {
            if (-655 != -281) {
                v29 := GenS4{F0: 128, F1: -111}
                _ = v29
            }
        } else {
            if true {
                var v30 float32 = 0.1
                _ = v30
            }
            var v31 int32 = 1
            _ = v31
        }
        _ = fi28
    }
    v32 := []int{}
    p2 += (p2 >> ((255) & 63))
    var v33 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v33
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{65535, 65535, 464, 928}
    var v3 uint64 = 18446744073709551615
    v3 = v3
    var v34 int16 = -1
    v34 = v34
    switch -12 {
    case 32768:
        if true {
            var v35 string = ""
            switch 9223372036854775807 {
            case 0:
                fmt.Println(v34, v35)
            case 805:
                var v36 int16 = 42
                _ = v36
            case -9223372036854775808:
                if true {
                    v34 -= (v34 * v34)
                }
            default:
                fmt.Println(v34)
            }
            _ = v35
        }
    case 0:
        fmt.Println(v34, v3)
    }
    fmt.Println(v3, v3)
    v34 = -1
    if (!false) {
        v34 += v34
    } else {
        var v37 float64 = (1.5 - 3.14159)
        _ = v37
    }
    var v38 string = ("hello" + ((("foo bar" + "\n\t") + ("" + "a")) + ("a" + "foo bar")))
    v39 := GenS4{F0: 18, F1: -1}
    fmt.Println(v38, v34, v3)
    fmt.Println(v3, v38, v38)
    v38 = "foo bar"
    for fi40 := 0; fi40 < 9; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            delete(v1, -336)
            var v42 uint32 = 0
            v38 = v38
            _ = fi41
            _ = v42
        }
        v1[fi40] = 1
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
}
