package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 string; F1 int; F2 int32 }
func genFn0() uint {
    var v7 uint = 255
    v7 = (v7 | v7)
    var v8 int = 128
    _ = v7
    _ = v8
    return (v7 % ((v7) | 1))
}
func genFn1(p0 uint64, p1 int16) uint16 {
    var v9 int64 = 9223372036854775807
    var v10 int8 = 1
    v11 := []int{-466, int(p0), -1, 32768}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = -1
    v3 := make(map[string]bool)
    var v4 float32 = 3.14159
    v5 := make(map[string]bool)
    var v6 uint8 = 168
    var v12 int8 = 100
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: 32767, F1: (1e10 == 0.0)}
        v15 := GenS1{F0: 416, F1: (true || (1 != (v6 >> ((2) & 63))))}
        switch "" {
        case "":
            var v16 uint = 4294967295
            _ = v16
        case "\n\t":
            fmt.Println(v2, v12, v2)
        default:
            v3[("hello" + "foo bar")] = false
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 += (v2 ^ v2)
    v17 := GenS2{F0: ("foo bar" + "a"), F1: 727, F2: 42}
    var v18 uint64 = (v1 | v1)
    var v19 uint = 4294967295
    v20 := GenS2{F0: "a", F1: 2147483647, F2: -318}
    fmt.Println(v12)
    fmt.Println(v18)
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if ((^(v19)) <= v19) {
                if false {
                    v6 -= (^(v6))
                    v12 = v12
                }
                var v23 int32 = 0
                var v24 uint32 = 0
                _ = v23
                _ = v24
            } else {
                v25 := GenS1{F0: genFn1(v18, -53), F1: true}
                _ = v25
            }
            var v26 uint32 = 2
            switch "foo bar" {
            case "hello":
                v18 -= 255
            case "\n\t":
                fmt.Println(v2, fi22, v1)
                if true {
                    v26 += uint32(v6)
                    var v27 string = ("foo bar" + ("" + ""))
                    fmt.Println(v27, v27, v1)
                    _ = v27
                }
            case "a":
                v2 = 1
            default:
                var v28 int64 = int64(v6)
                _ = v28
            }
            _ = fi22
            _ = v26
        }
        _ = fi21
    }
    v29 := GenS1{F0: genFn1(v18, 948), F1: true}
    v30 := GenS2{F0: "foo bar", F1: 1, F2: 710}
    for fi31 := 0; fi31 < 15; fi31++ {
        v1 += 1
        fmt.Println(v1, v12, v19)
        for fi32 := 0; fi32 < 7; fi32++ {
            switch "\n\t" {
            case "hello":
                var v33 uint32 = 2
                _ = v33
            }
            v34 := GenS1{F0: 32767, F1: false}
            if (false || false) {
                switch "a" {
                case "\n\t":
                    v35 := []int{fi31}
                    _ = v35
                }
            } else {
                switch fi31 {
                case -1:
                    switch ("hello" + ("a" + "a")) {
                    case "\n\t":
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            var v36 uint32 = 2147483647
                            _ = v36
                        case "a":
                            var v37 float32 = float32(-2.25)
                            _ = v37
                        default:
                            v6 += v6
                        }
                    }
                default:
                    var v38 uint16 = uint16(v6)
                    _ = v38
                }
                v4 = 1.5
            }
            if false {
                var v39 int8 = (v12 | v12)
                _ = len(v5)
                v40 := GenS2{F0: "foo bar", F1: (fi32 >> ((genFn0()) & 63)), F2: 42}
                _ = v39
                _ = v40
            } else {
                v1 += v18
                if false {
                    v19 += v19
                    var v41 uint16 = 696
                    v42 := v3[("hello" + ("a" + "a"))]
                    v12 -= v12
                    _ = v41
                    _ = v42
                } else {
                    var v43 int16 = 1000
                    v44 := GenS2{F0: "\n\t", F1: (fi32 ^ (fi32 >> ((v19) & 63))), F2: -1}
                    _ = v43
                    _ = v44
                }
            }
            _ = fi32
            _ = v34
        }
        _ = fi31
    }
    if false {
        fmt.Println(v19)
        fmt.Println(v18)
        v45 := GenS1{F0: 32767, F1: (true || (false || false))}
        _ = v45
    } else {
        var v46 int8 = int8(v19)
        _ = v46
    }
    v47 := []int{1, 9}
    if len(v47) > 0 {
        v47[0] = 42
    }
    switch 127 {
    case -342:
        if len(v47) >= 2 {
            v47 = v47[1:len(v47)]
        }
    case -2147483648:
        var v48 float32 = ((float32(3.14159) / (v4 * 1e10)) - 0.1)
        _ = v48
    }
    v19 += genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
