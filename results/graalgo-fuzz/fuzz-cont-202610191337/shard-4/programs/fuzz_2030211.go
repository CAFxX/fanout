package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 int8, p2 int32) string {
    switch ("a" + (("\n\t" + "\n\t") + "")) {
    case "":
        p2 -= (^(p2))
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{32768, fi5, 42}
        _ = fi5
        _ = v6
    }
    var v7 bool = (32767 <= 1)
    var v8 float32 = float32(((float64(0.1) / ((-2.25 + 0.1) * 3.14159)) / 1e10))
    var v9 int64 = int64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (("\n\t" + ("foo bar" + (("" + "hello") + ("\n\t" + ("\n\t" + "a"))))) + ("a" + ""))
}
func genFn1(p0 string, p1 uint64) uint64 {
    p0 = p0
    v10 := []int{9223372036854775807, 4294967295}
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    var v1 bool = (2147483647 > 2147483647)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 42
    var v4 string = ""
    switch -1 {
    case 469:
        var v11 int32 = v3
        _ = v11
    case -2147483648:
        var v12 uint8 = 2
        _ = v12
    case 854:
        fmt.Println(v1, v4, v1)
    }
    fmt.Println(v1)
    var v13 float32 = 1e10
    if true {
        var v14 bool = v1
        var v15 uint16 = 42
        if (255 >= 574) {
            for fi16 := 0; fi16 < 6; fi16++ {
                v2[2] = 83
                switch fi16 {
                case -9223372036854775808:
                    switch v4 {
                    case "\n\t":
                        switch (((v4 + "\n\t") + v4) + "a") {
                        case "\n\t":
                            var v17 uint8 = 215
                            _ = v17
                        }
                    case "":
                        var v18 bool = (v1 || true)
                        _ = v18
                    case "a":
                        v15 = 632
                    }
                default:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if (!(v4 != ((v4 + (v4 + "\n\t")) + v4))) {
                            v3 -= -286
                            var v20 float64 = -2.25
                            var v21 int64 = 42
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                }
                fmt.Println(v4)
                _ = fi16
            }
        } else {
            fmt.Println(v15, v14, v3)
            var v22 float32 = (((v13 * 3.14159) * (v13 / v13)) / v13)
            v3 += v3
            _ = v22
        }
        _ = v14
        _ = v15
    } else {
        v23 := GenS1{F0: -71, F1: v4, F2: 0}
        if false {
            delete(v2, 166)
            switch ((v4 + v4) + ("hello" + (("\n\t" + v4) + ("foo bar" + "foo bar")))) {
            case "":
                v13 += (1.5 / (v13 * v13))
            case "hello":
                var v24 int32 = (v3 + v3)
                _ = v24
            case "\n\t":
                v25 := GenS2{F0: -128}
                _ = v25
            }
            var v26 uint8 = 1
            _ = v26
        }
        if true {
            fmt.Println(v1, v4)
            v27 := []int{4294967295, 2147483647}
            _ = len(v27)
            _ = v27
        } else {
            v28 := GenS1{F0: 47, F1: ("foo bar" + "a"), F2: 1}
            _ = v28
        }
        _ = v23
    }
    delete(v2, 42)
    v29 := GenS1{F0: -1, F1: v4, F2: 65535}
    var v30 string = v4
    var v31 uint16 = 283
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
