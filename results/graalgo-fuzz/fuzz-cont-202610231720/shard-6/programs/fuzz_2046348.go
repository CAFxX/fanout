package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int; F2 int64 }
func genFn0() uint32 {
    var v6 int64 = -1
    v7 := []int{}
    _ = v6
    _ = v7
    return 216
}
func genFn1(p0 int8, p1 bool, p2 int) int8 {
    v8 := []int{p2}
    p1 = (2147483648 >= 0)
    if true {
        switch p2 {
        case -341:
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 int64 = 9223372036854775807
                v11 := GenS1{F0: 205, F1: (p0 + p0)}
                v10 -= 31
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            _ = len(v8)
        }
    }
    if (!p1) {
        for fi12 := 0; fi12 < 10; fi12++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi12
        }
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = (3.14159 - 3.14159)
    var v3 float32 = (0.1 * 0.1)
    var v4 uint32 = 1
    var v5 uint = 42
    v13 := GenS2{F0: genFn0(), F1: 1000}
    v14 := GenS3{F0: "", F1: 1}
    fmt.Println(v1, v4)
    v15 := GenS4{F0: (("" + "hello") + ("hello" + "foo bar"))}
    for fi16 := 0; fi16 < 16; fi16++ {
        if (!false) {
            var v17 int8 = int8(fi16)
            v4 = genFn0()
            _ = v17
        } else {
            var v18 uint64 = v1
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int32 = 0
                _ = fi19
                _ = v20
            }
            var v21 int8 = 1
            _ = v18
            _ = v21
        }
        v22 := GenS3{F0: "", F1: 32767}
        _ = fi16
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "hello" {
        case "hello":
            if ((-69 >= 1000) && true) {
                v3 -= 1.5
                v1 = 699
                var v24 int64 = 9223372036854775807
                _ = v24
            } else {
                switch 487 {
                case 128:
                    v25 := []int{fi23, fi23}
                    _ = v25
                case -1:
                    switch "hello" {
                    case "":
                        v26 := []int{}
                        _ = v26
                    case "\n\t":
                        var v27 uint32 = v4
                        _ = v27
                    case "hello":
                        v4 += (^(v4))
                    default:
                        switch fi23 {
                        case 270:
                            var v28 float32 = 3.14159
                            _ = v28
                        }
                    }
                case 9223372036854775807:
                    v29 := []int{(fi23 * (fi23 << (((v5 & (v5 * 42))) & 63)))}
                    _ = v29
                }
                v30 := []int{(fi23 / (((-(fi23))) | 1)), (fi23 << (((v5 / (((v5 & (v5 >> ((65535) & 63)))) | 1))) & 63)), 4294967295, (fi23 >> ((718) & 63))}
                for fi31 := 0; fi31 < 11; fi31++ {
                    fmt.Println(fi31, v2)
                    _ = fi31
                }
                _ = v30
            }
        case "\n\t":
            fmt.Println(v1, v3, v5)
        }
        _ = fi23
    }
    if (1 > (v4 / ((genFn0()) | 1))) {
        v4 += genFn0()
    } else {
        v32 := []int{}
        switch (("" + "hello") + "foo bar") {
        case "":
            _ = len(v32)
        case "hello":
            var v33 uint8 = 2
            _ = v33
        case "foo bar":
            fmt.Println(v3, v5)
        }
        _ = v32
    }
    v34 := GenS3{F0: "a", F1: -1}
    if (!true) {
        var v35 uint = (v5 % ((v5) | 1))
        if true {
            v3 = ((v3 - 0.0) * 1e10)
            for fi36 := 0; fi36 < 2; fi36++ {
                for fi37 := 0; fi37 < 1; fi37++ {
                    var v38 int16 = 1
                    _ = fi37
                    _ = v38
                }
                var v39 bool = (1 != 923)
                var v40 int32 = -70
                _ = fi36
                _ = v39
                _ = v40
            }
        }
        v41 := GenS5{F0: (v1 | (v1 | (v1 / ((v1) | 1))))}
        _ = v35
        _ = v41
    }
    fmt.Println(v5, v5, v2)
    v4 = genFn0()
    v42 := []int{139, 0}
    if len(v42) > 0 {
        v42[0] = -627
    }
    v42 = append(v42, -2147483648)
    v2 -= 0.0
    v1 = v1
    if (v1 == (v1 & (v1 - v1))) {
        v43 := GenS6{F0: ("" + "a"), F1: 2147483647, F2: 1}
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
        v4 += genFn0()
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
}
