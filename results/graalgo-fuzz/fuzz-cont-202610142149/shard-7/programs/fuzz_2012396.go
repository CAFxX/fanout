package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int; F2 uint32 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint64, p2 int) uint16 {
    p0 = p0
    if true {
        v5 := GenS1{F0: 255, F1: 42}
        var v6 float64 = 1.5
        v6 -= v6
        _ = v5
        _ = v6
    } else {
        p0 = (-(p0))
    }
    switch -1 {
    case -1:
        switch ("hello" + "") {
        case "hello":
            v7 := GenS1{F0: 559, F1: 42}
            _ = v7
        case "foo bar":
            var v8 int8 = 42
            _ = v8
        case "\n\t":
            var v9 int8 = -119
            _ = v9
        }
    default:
        v10 := GenS1{F0: 4294967295, F1: (^(p1))}
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            fmt.Println(p2, p2, p2)
            v12 := []int{fi11, (p2 % (((^(fi11))) | 1)), p2, p2}
            _ = v12
        } else {
            fmt.Println(fi11, p1)
            var v13 float64 = ((1.5 + (1.5 + (-2.25 * -2.25))) + -2.25)
            _ = v13
        }
        v14 := GenS2{F0: p2}
        var v15 float32 = (-2.25 + (1e10 - 1e10))
        _ = fi11
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 int16, p1 uint32) uint8 {
    v16 := []int{}
    v17 := GenS3{F0: 1, F1: -1}
    var v18 uint64 = 0
    var v19 int = 128
    switch "foo bar" {
    case "\n\t":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        switch "hello" {
        case "foo bar":
            v18 = v18
        }
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    v2 := []uint32{}
    var v3 int = 0
    var v4 float64 = (1e10 / float64(0.1))
    var v20 int64 = 42
    v21 := GenS4{F0: genFn0(0, 4294967295, v3), F1: 713, F2: 2147483648}
    var v22 uint = 255
    for fi23 := 0; fi23 < 18; fi23++ {
        v2 = append(v2, 4294967295)
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 uint8 = 128
            if (42 >= 18446744073709551615) {
                v3 = (-(fi24))
            } else {
                v26 := GenS5{F0: (v25 << (((v22 / (((v22 | (v22 - v22))) | 1))) & 63))}
                v25 = (v25 ^ v25)
                fmt.Println(v20, v20)
                fmt.Println(fi23)
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        v4 -= 1e10
        _ = fi23
    }
    v27 := GenS6{F0: 4294967295, F1: v3, F2: 689}
    var v28 string = "hello"
    if (v22 != v22) {
        for fi29 := 0; fi29 < 3; fi29++ {
            v4 += (0.0 * (v4 * float64((1.5 - 0.0))))
            var v30 int16 = 0
            _ = fi29
            _ = v30
        }
        v2 = append(v2, 2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v31 uint16 = 32767
        _ = v31
    }
    var v32 uint = (v22 ^ (v22 | 2))
    fmt.Println(v28, v22)
    v22 += (v22 | v32)
    if true {
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 int16 = -729
            _ = fi33
            _ = v34
        }
        switch (v3 * (v3 & (v3 >> (((v22 >> (((^(v22))) & 63))) & 63)))) {
        case 42:
            var v35 string = (("a" + ("" + v28)) + "a")
            _ = v35
        case 0:
            v4 += 1.5
        default:
            var v36 int16 = -32768
            _ = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            if (genFn1(1, 42) <= 117) {
                v38 := GenS7{F0: 4294967295}
                for fi39 := 0; fi39 < 19; fi39++ {
                    fmt.Println(v32, v28, v22)
                    v40 := GenS6{F0: 4294967295, F1: 127, F2: 2147483647}
                    var v41 uint = (^(v32))
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = v38
            }
            if true {
                v42 := GenS6{F0: 1, F1: fi37, F2: 2147483647}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v42
            } else {
                if len(v2) > 0 {
                    v2[0] = 784
                }
                v28 = (("foo bar" + ("\n\t" + v28)) + v28)
                var v43 uint8 = 42
                _ = v43
            }
            _ = fi37
        }
    } else {
        var v44 uint64 = 42
        switch v3 {
        case 4294967295:
            v4 -= (-2.25 / (v4 / v4))
        case -473:
            var v45 int16 = -426
            _ = v45
        case 668:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v44
    }
    fmt.Println(v3)
    var v46 int32 = -955
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v47 int = 42
    v22 = 18446744073709551615
    v20 -= 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v46
    _ = v47
}
