package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 int8; F1 string; F2 uint }
func genFn0() int32 {
    var v5 uint32 = 778
    v6 := GenS1{F0: (v5 ^ 4294967295), F1: 2}
    v7 := []int{435, -489, 32768, -353}
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 int64, p2 int32) uint {
    p2 -= (p2 | (p0 ^ 42))
    var v8 uint16 = 2
    v9 := GenS1{F0: 65535, F1: 294}
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{718, 127, 0}
    v3 := make(map[string]string)
    var v4 int16 = 0
    var v10 int = -136
    switch "hello" {
    case "hello":
        if (!(true || true)) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v1 = v1
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 bool = (42 > 0)
                    if (318 == 0) {
                        var v14 float64 = 1e10
                        _ = v14
                    } else {
                        fmt.Println(v13)
                        fmt.Println(v10)
                    }
                    v15 := GenS1{F0: 386, F1: 18446744073709551615}
                    _ = fi12
                    _ = v13
                    _ = v15
                }
                fmt.Println(v1, v4, v10)
                _ = fi11
            }
            v16 := GenS2{F0: -128, F1: (v10 & v10)}
            var v17 uint64 = 90
            _ = v16
            _ = v17
        } else {
            v18 := GenS1{F0: 2, F1: 0}
            _ = len(v2)
            _ = v18
        }
    case "a":
        _ = len(v3)
    case "":
        var v19 int32 = 42
        _ = v19
    }
    v4 -= v4
    var v20 float64 = (-2.25 - 1.5)
    v4 += v4
    v10 -= 32767
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            delete(v3, "hello")
            _ = fi22
        }
        if true {
            v1 -= 0.1
            fmt.Println(fi21)
            v23 := GenS3{F0: (67 != 100), F1: 32767, F2: ("foo bar" + ("\n\t" + "\n\t"))}
            _ = v23
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v1 = v1
                var v25 uint8 = 114
                _ = fi24
                _ = v25
            }
        }
        _ = fi21
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 uint = 4294967295
            _ = fi27
            _ = v28
        }
        v29 := GenS4{F0: 0, F1: "", F2: genFn1(-622, -9223372036854775808, -2147483648)}
        _ = fi26
        _ = v29
    }
    switch (v10 * -9223372036854775808) {
    case 128:
        var v30 uint64 = 648
        _ = v30
    case -794:
        fmt.Println(v10)
    }
    switch "foo bar" {
    case "hello":
        v1 -= float32((v20 / -2.25))
    }
    if false {
        var v31 uint16 = 933
        _ = v31
    }
    var v32 uint16 = 246
    v32 += v32
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
}
