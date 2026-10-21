package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint32; F1 string; F2 string }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int, p2 uint16) uint {
    v5 := GenS1{F0: 2, F1: 128}
    switch "a" {
    case "a":
        v6 := []int{p1}
        _ = v6
    default:
        fmt.Println(p1, p2, p0)
    }
    switch ("a" + ((("" + "\n\t") + "\n\t") + "a")) {
    case "hello":
        if ((("foo bar" + "\n\t") <= "a") && (true || ((p1 < p1) || (1 <= p0)))) {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v7 float32 = (1.5 + (((-2.25 / float32(3.14159)) + (((((0.0 + ((0.1 + (3.14159 + 1e10)) + float32(3.14159))) * float32(0.0)) + 0.0) / (1e10 - 3.14159)) + (-2.25 - 1.5))) + 0.0))
                _ = v7
            case "hello":
                v8 := GenS1{F0: 0, F1: (p1 ^ p1)}
                _ = v8
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(p2, p0)
                _ = fi9
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                switch p1 {
                case -598:
                    p1 -= (p1 % ((fi10) | 1))
                case 126:
                    var v11 int8 = 0
                    _ = v11
                default:
                    var v12 float64 = 3.14159
                    _ = v12
                }
                v13 := GenS1{F0: 406, F1: (^(fi10))}
                p1 = 127
                _ = fi10
                _ = v13
            }
        } else {
            if (3.14159 < 0.1) {
                if true {
                    var v14 int16 = -1
                    var v15 int16 = v14
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 string = (("" + ("\n\t" + "hello")) + "")
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                    var v19 float64 = float64(3.14159)
                    switch "a" {
                    case "":
                        v20 := []int{-2147483648}
                        _ = v20
                    case "\n\t":
                        for fi21 := 0; fi21 < 18; fi21++ {
                            switch "hello" {
                            case "foo bar":
                                p1 = (fi21 / (((p1 % ((fi21) | 1))) | 1))
                            case "a":
                                var v22 int = 4294967295
                                _ = v22
                            }
                            _ = fi21
                        }
                    case "foo bar":
                        var v23 bool = true
                        _ = v23
                    }
                    _ = v19
                }
                p1 = p1
                v24 := GenS2{F0: 1000, F1: true}
                _ = v24
            }
            var v25 int8 = -1
            _ = v25
        }
    default:
        var v26 uint64 = 901
        _ = v26
    }
    var v27 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v5
    return 65535
}
func genFn1(p0 uint8) uint32 {
    var v28 uint16 = 913
    var v29 int8 = -128
    var v30 uint64 = 2
    if false {
        v31 := []int{9223372036854775807, 0}
        v28 = v28
        _ = v31
    }
    _ = p0
    _ = v28
    _ = v29
    _ = v30
    return 2
}

func main() {
    var v1 uint16 = 1
    var v2 float64 = (0.1 * 1e10)
    var v3 int = 2147483647
    v4 := make(map[string]int)
    v3 = -842
    if (!false) {
        for fi32 := 0; fi32 < 10; fi32++ {
            fmt.Println(v3, v2)
            _ = fi32
        }
        v4[("hello" + "\n\t")] = (v3 & 0)
        v33 := []int{v3, (v3 & 32768)}
        _ = v33
    }
    if false {
        if (174 == uint8(v1)) {
            v3 = v3
        }
        v4["a"] = 127
    }
    var v34 int64 = 0
    switch "a" {
    case "\n\t":
        for fi35 := 0; fi35 < 13; fi35++ {
            v36 := GenS2{F0: -32768, F1: (0.0 > float32(v2))}
            _ = fi35
            _ = v36
        }
    case "":
        var v37 uint8 = 73
        _ = v37
    case "a":
        if true {
            v34 = v34
            var v38 int64 = 9223372036854775807
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            if true {
                var v40 string = "a"
                var v41 int8 = 100
                v42 := GenS3{F0: 2, F1: v40}
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi39
        }
    }
    if true {
        for fi43 := 0; fi43 < 3; fi43++ {
            if false {
                v44 := GenS3{F0: 255, F1: "a"}
                v3 += v3
                var v45 int64 = 0
                _ = v44
                _ = v45
            } else {
                for fi46 := 0; fi46 < 1; fi46++ {
                    v47 := v4[""]
                    v48 := GenS4{F0: 42, F1: "\n\t", F2: "a"}
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                fmt.Println(v3, v34)
                v2 = float64(float32(0.0))
            }
            _ = fi43
        }
        v4["a"] = 334
    }
    var v49 uint32 = 0
    var v50 bool = (!false)
    fmt.Println(v1, v50)
    v51 := []int{(v3 << ((411) & 63)), v3, v3}
    for fi52 := 0; fi52 < 6; fi52++ {
        v2 -= v2
        var v53 int8 = 1
        fmt.Println(v50)
        _ = fi52
        _ = v53
    }
    for fi54 := 0; fi54 < 17; fi54++ {
        var v55 int = int(v1)
        switch ("a" + "foo bar") {
        case "foo bar":
            for fi56 := 0; fi56 < 18; fi56++ {
                v57 := GenS1{F0: 0, F1: v3}
                v1 = 896
                v55 += 2147483647
                v3 += (fi56 * fi54)
                _ = fi56
                _ = v57
            }
        case "":
            for fi58 := 0; fi58 < 13; fi58++ {
                v55 = fi54
                _ = fi58
            }
        case "hello":
            v34 = v34
        }
        _ = fi54
        _ = v55
    }
    v59 := GenS5{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v59
}
