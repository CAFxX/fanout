package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 bool; F2 int64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 int8, p1 int) uint {
    v7 := GenS1{F0: (((("a" + "\n\t") + "foo bar") + "hello") + (("hello" + "a") + (("" + "a") + "hello")))}
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{fi8, p1}
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi8)
            p0 = (^(p0))
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    switch 0 {
    case 0:
        v11 := GenS1{F0: "foo bar"}
        _ = v11
    case 32768:
        p1 = p1
    default:
        if true {
            if true {
                v12 := []int{int(p0), p1, p1}
                if (0 == 18446744073709551615) {
                    if (42 <= 193) {
                        v13 := GenS1{F0: ("\n\t" + (("a" + ("\n\t" + "\n\t")) + "foo bar"))}
                        fmt.Println(p1, p0)
                        p0 += p0
                        _ = v13
                    }
                    var v14 uint8 = 255
                    var v15 bool = (3.14159 == 1e10)
                    _ = v14
                    _ = v15
                } else {
                    _ = len(v12)
                    p0 -= p0
                    for fi16 := 0; fi16 < 4; fi16++ {
                        p1 += 32767
                        for fi17 := 0; fi17 < 15; fi17++ {
                            p0 += (p0 & 1)
                            _ = fi17
                        }
                        p1 = (fi16 + (fi16 / ((-130) | 1)))
                        _ = fi16
                    }
                }
                _ = v12
            } else {
                var v18 bool = (867 != 42)
                _ = v18
            }
            var v19 int = (p1 % (((^(p1))) | 1))
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint, p2 int8) int8 {
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        if false {
            v20 := GenS1{F0: "foo bar"}
            _ = v20
        }
    case "":
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int = (fi21 ^ (fi21 / ((0) | 1)))
            if (true || true) {
                switch "foo bar" {
                case "":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        p0 -= p0
                        _ = fi23
                    }
                case "foo bar":
                    fmt.Println(p1, v22, p2)
                case "hello":
                    var v24 uint32 = 42
                    _ = v24
                default:
                    p0 -= (^(p0))
                }
            }
            v25 := []int{127, v22}
            _ = fi21
            _ = v22
            _ = v25
        }
    default:
        v26 := GenS1{F0: "a"}
        _ = v26
    }
    p1 += (p1 - p1)
    v27 := GenS2{F0: 732, F1: (^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint = v1
    var v3 float32 = (0.0 - 0.1)
    v4 := []string{}
    var v5 uint16 = 393
    var v6 int32 = 680
    v28 := GenS3{F0: "\n\t", F1: false, F2: 9223372036854775807}
    fmt.Println(v5, v6)
    v1 += (^(v1))
    switch (("" + "\n\t") + ("\n\t" + "")) {
    case "hello":
        v29 := GenS2{F0: (v6 | 310), F1: (v1 / (((v2 - (v1 & genFn0(-1, 9223372036854775807)))) | 1))}
        _ = v29
    case "":
        v5 = (v5 << ((v2) & 63))
    }
    var v30 uint = (v1 + v1)
    fmt.Println(v6)
    v3 += (1.5 + float32(0.0))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v31 uint32 = 808
    if true {
        var v32 uint16 = (v5 << ((uint(v31)) & 63))
        var v33 int16 = 0
        _ = v32
        _ = v33
    }
    var v34 bool = (false || (0 != 0))
    fmt.Println(v6, v6, v1)
    v4 = append(v4, "foo bar")
    for fi35 := 0; fi35 < 4; fi35++ {
        fmt.Println(v31, fi35, v3)
        switch (fi35 >> (((v30 & (^(v1)))) & 63)) {
        case -703:
            v2 += 709
        case 9223372036854775807:
            var v36 uint64 = 2
            _ = v36
        default:
            if (!("hello" >= "foo bar")) {
                if (true && (!v34)) {
                    v37 := GenS4{F0: 9223372036854775807, F1: "foo bar"}
                    _ = v37
                }
            }
        }
        v38 := GenS5{F0: -1, F1: 1}
        _ = fi35
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
