package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int) int {
    switch p0 {
    case 32767:
        var v7 uint16 = 190
        _ = v7
    case -776:
        var v8 uint = 1
        _ = v8
    default:
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    v10 := []int{p0, (p0 >> ((255) & 63))}
    fmt.Println(p0)
    var v11 float32 = 1e10
    v10 = append(v10, 1)
    _ = p0
    _ = v10
    _ = v11
    return (p0 % (((p0 / ((0) | 1))) | 1))
}
func genFn1(p0 string, p1 uint8) bool {
    for fi12 := 0; fi12 < 2; fi12++ {
        switch (fi12 - genFn0((fi12 % (((fi12 + (fi12 >> ((300) & 63)))) | 1)))) {
        case 2147483647:
            for fi13 := 0; fi13 < 12; fi13++ {
                p1 = 0
                v14 := GenS1{F0: (fi12 + fi13), F1: (p0 + (p0 + (((("hello" + p0) + (p0 + p0)) + ("a" + (("foo bar" + "hello") + p0))) + p0))), F2: -9223372036854775808}
                var v15 string = (p0 + p0)
                fmt.Println(p1, fi12, p0)
                _ = fi13
                _ = v14
                _ = v15
            }
        case -2147483648:
            for fi16 := 0; fi16 < 11; fi16++ {
                p1 -= (p1 & p1)
                p1 += p1
                _ = fi16
            }
        case 32767:
            var v17 bool = (0.0 > (0.1 - (0.0 * (1e10 - (((0.1 * 0.0) - 1.5) + 0.0)))))
            _ = v17
        default:
            if (-150 <= 680) {
                var v18 uint16 = 65535
                _ = v18
            }
        }
        switch genFn0(genFn0((^(fi12)))) {
        case -9223372036854775808:
            fmt.Println(fi12)
        }
        _ = fi12
    }
    p1 = (^(p1))
    fmt.Println(p0, p1)
    switch -650 {
    case 656:
        var v19 uint = 770
        _ = v19
    case 32767:
        var v20 bool = (4294967295 > genFn0(355))
        _ = v20
    }
    var v21 bool = true
    _ = p0
    _ = p1
    _ = v21
    return (2 < 2)
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 2147483647
    var v3 uint64 = 249
    var v4 uint8 = 113
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    switch 2147483647 {
    case 1:
        delete(v5, -2147483648)
    case 0:
        v2 -= 0
    case -1:
        switch "" {
        case "a":
            if (2 == v4) {
                var v22 uint = 4294967295
                v3 += 1
                _ = v22
            }
        case "":
            var v23 bool = false
            _ = v23
        default:
            var v24 int8 = -1
            _ = v24
        }
    }
    var v25 int64 = 42
    var v26 uint16 = 684
    for fi27 := 0; fi27 < 14; fi27++ {
        v25 = v25
        _ = fi27
    }
    var v28 int = int(v1)
    for fi29 := 0; fi29 < 11; fi29++ {
        v3 -= (v3 << ((4294967295) & 63))
        _ = fi29
    }
    _ = len(v6)
    v30 := []int{v28}
    var v31 int32 = (v2 * v2)
    var v32 int8 = -38
    v33 := GenS2{F0: (v32 ^ (^(v32)))}
    switch "foo bar" {
    case "foo bar":
        var v34 int8 = 48
        _ = v34
    case "\n\t":
        fmt.Println(v32, v32, v32)
    case "":
        for fi35 := 0; fi35 < 9; fi35++ {
            if false {
                for fi36 := 0; fi36 < 16; fi36++ {
                    if (!(!((v4 - (^(v4))) > (v4 / ((uint8(v3)) | 1))))) {
                        v37 := GenS1{F0: v28, F1: "\n\t", F2: v25}
                        v3 = v3
                        v38 := GenS2{F0: (v32 & 0)}
                        _ = v37
                        _ = v38
                    }
                    var v39 uint8 = 1
                    var v40 int16 = -526
                    _ = fi36
                    _ = v39
                    _ = v40
                }
                v41 := GenS3{F0: v1, F1: "a"}
                _ = v41
            }
            switch "hello" {
            case "":
                v42 := GenS4{F0: (v3 >> ((255) & 63)), F1: 0}
                _ = v42
            }
            switch (v28 - genFn0(genFn0((-(fi35))))) {
            case 32767:
                var v43 uint = 262
                _ = v43
            default:
                var v44 bool = (false || false)
                _ = v44
            }
            _ = fi35
        }
    default:
        var v45 uint64 = 4294967295
        _ = v45
    }
    var v46 int32 = 530
    var v47 int16 = 32767
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
