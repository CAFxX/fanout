package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 bool, p2 uint8) int8 {
    if p1 {
        var v7 int32 = 2147483647
        v8 := []int{-1, -1}
        _ = v7
        _ = v8
    } else {
        v9 := []int{128}
        var v10 int8 = 36
        _ = v10
        _ = v9
    }
    var v11 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -41
}
func genFn1(p0 uint32, p1 int16) uint16 {
    switch -1 {
    case 9223372036854775807:
        v12 := GenS1{F0: 127, F1: (^(p1))}
        _ = v12
    }
    if (2147483647 < 2147483647) {
        p1 += (p1 << ((255) & 63))
        v13 := GenS1{F0: genFn0(-388, (2147483647 == p0), 0), F1: p1}
        var v14 uint = 255
        _ = v13
        _ = v14
    }
    v15 := []int{}
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint16 = 2
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v15
    return 255
}

func main() {
    var v1 uint16 = 848
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v6 uint32 = 130
    delete(v4, "a")
    var v18 int32 = 42
    v19 := GenS1{F0: 127, F1: 32767}
    _ = len(v4)
    v3 = (v3 << ((2) & 63))
    var v20 bool = false
    v1 += uint16(v6)
    var v21 int32 = v3
    switch "foo bar" {
    case "foo bar":
        v18 = -940
    case "\n\t":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 int8 = genFn0(42, (v20 || (("\n\t" < (("a" + "\n\t") + ("a" + (("hello" + "\n\t") + "a")))) || (v1 == (v1 | v1)))), 255)
            _ = fi22
            _ = v23
        }
    }
    switch ("" + "a") {
    case "\n\t":
        switch "hello" {
        case "foo bar":
            if (((130 != 0) || ((v3 & -1) != int32(v6))) || true) {
                if (v20 && (v6 <= (^(v6)))) {
                    if ("a" > "foo bar") {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                fmt.Println(v20)
                                var v26 uint = uint(v1)
                                _ = fi25
                                _ = v26
                            }
                            v27 := GenS1{F0: 1, F1: -1}
                            var v28 float32 = (0.0 * 0.0)
                            _ = fi24
                            _ = v27
                            _ = v28
                        }
                        switch -2147483648 {
                        case -9223372036854775808:
                            v18 = (v3 | (v21 % (((v18 << ((42) & 63))) | 1)))
                        default:
                            v29 := GenS2{F0: "foo bar", F1: 32767, F2: 0}
                            _ = v29
                        }
                        var v30 uint = 255
                        _ = v30
                    } else {
                        var v31 uint = 927
                        _ = v31
                    }
                    fmt.Println(v6, v20, v6)
                    v32 := GenS1{F0: -1, F1: 1000}
                    _ = v32
                }
                var v33 int64 = 1
                _ = v33
            }
        }
    case "foo bar":
        v1 += (v1 - v1)
    }
    v34 := []int{24, 2147483647, -2147483648}
    _ = len(v4)
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int32 = (v3 & v21)
            _ = fi36
            _ = v37
        }
        v1 = genFn1(408, 1000)
        var v38 uint64 = uint64(v1)
        _ = fi35
        _ = v38
    }
    v39 := GenS3{F0: 2147483647, F1: -1}
    var v40 bool = (127 < -1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
