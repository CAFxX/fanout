package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0(p0 string, p1 int, p2 uint) int8 {
    switch (p1 & p1) {
    case 40:
        var v7 float32 = 1.5
        _ = v7
    case -10:
        p1 = p1
    }
    var v8 uint32 = 2
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}
func genFn1(p0 uint32) int {
    p0 += 42
    fmt.Println(p0)
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(fi9, fi9, fi10)
            v11 := []int{(-(fi10))}
            _ = fi10
            _ = v11
        }
        v12 := []int{9223372036854775807, (fi9 + 127), -606, fi9}
        _ = fi9
        _ = v12
    }
    var v13 int8 = genFn0("hello", 32767, 42)
    p0 += 0
    _ = p0
    _ = v13
    return 729
}

func main() {
    v1 := []string{("" + "foo bar"), "a", (("a" + "hello") + "a")}
    v2 := make(map[int]int)
    var v3 uint8 = 42
    v4 := make(map[uint32]bool)
    var v5 string = "foo bar"
    v6 := make(map[string]bool)
    fmt.Println(v3)
    v5 = "a"
    switch -9223372036854775808 {
    case 128:
        v14 := GenS1{F0: 42, F1: -9223372036854775808}
        _ = v14
    }
    v15 := GenS2{F0: 2147483647, F1: genFn0(v5, 195, 671), F2: 255}
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            var v17 int64 = 1
            var v18 bool = (18446744073709551615 == 255)
            _ = v17
            _ = v18
        } else {
            switch ("foo bar" + "hello") {
            case "foo bar":
                var v19 int32 = int32(v3)
                _ = v19
            case "\n\t":
                if (true || ("" != "\n\t")) {
                    if false {
                        if false {
                            v20 := GenS3{F0: 2}
                            _ = v20
                        }
                    }
                    v21 := GenS4{F0: v5, F1: v5, F2: fi16}
                    v3 += (v3 >> ((365) & 63))
                    _ = v21
                } else {
                    var v22 int8 = -128
                    if (!false) {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            var v24 int8 = (v22 / (((v22 | (v22 | (v22 + (v22 / (((v22 ^ 100)) | 1)))))) | 1))
                            v25 := GenS2{F0: 2147483647, F1: 0, F2: 275}
                            v1 = append(v1, "")
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    v5 = v5
                    _ = v22
                }
            case "a":
                v5 = ""
            }
        }
        v26 := GenS1{F0: -1, F1: -1}
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS1{F0: 112, F1: 42}
            v29 := GenS5{F0: -9223372036854775808, F1: -424}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi16
        _ = v26
    }
    v3 += v3
    v5 = v5
    fmt.Println(v5, v5)
    v5 = "foo bar"
    if (!((2 == 4294967295) || (32768 < 604))) {
        var v30 float64 = (float64((1.5 - (3.14159 / 1e10))) * ((-2.25 / 1.0) / (3.14159 / 1.0)))
        var v31 uint16 = 0
        _ = v30
        _ = v31
    }
    var v32 float32 = 1.5
    var v33 int16 = 1000
    v32 = v32
    v34 := GenS4{F0: (v5 + ("" + (v5 + "a"))), F1: v5, F2: -1}
    if len(v1) > 0 {
        v1[0] = v5
    }
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
