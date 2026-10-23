package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8; F2 int }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 float64 = (-2.25 + (3.14159 / -2.25))
    var v4 bool = (v1 <= 696)
    var v5 float32 = float32((v3 * (v3 * (((-2.25 + 0.1) * 0.0) + v3))))
    var v6 uint64 = 18446744073709551615
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 11; fi8++ {
        v3 -= -2.25
        fmt.Println(fi8, fi8)
        for fi9 := 0; fi9 < 19; fi9++ {
            v4 = v4
            v10 := []int{(-(fi9))}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    delete(v2, "foo bar")
    switch "hello" {
    case "foo bar":
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint32 = 4294967295
                if (-1 >= 2147483647) {
                    var v14 uint32 = (v13 & v13)
                    v14 = (v14 - (v14 % ((4294967295) | 1)))
                    _ = v14
                }
                if ((v1 + 322) != uint(v6)) {
                    if v4 {
                        v15 := GenS1{F0: 0, F1: v1}
                        v16 := GenS2{F0: v13}
                        if (v3 <= (float64(v7) + ((((v3 / 3.14159) - v3) - (-2.25 - (v3 * 1e10))) * float64(v7)))) {
                            var v17 float32 = float32((v3 / 1.0))
                            v3 += (v3 / (float64((v5 * ((float32(v3) / float32((float64((((0.0 - ((v5 + float32((v3 - v3))) * 3.14159)) + v7) + 0.1)) * (v3 + ((1e10 / v3) + -2.25))))) / (float32((1.5 + float64((v5 + (0.1 - 3.14159))))) - v5)))) - v3))
                            _ = v17
                        } else {
                            v7 = v7
                            v2[("" + "a")] = false
                            v18 := []int{fi11}
                            _ = v18
                        }
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    case "hello":
        for fi19 := 0; fi19 < 19; fi19++ {
            v7 = (0.1 - (v5 * ((float32(v3) - (3.14159 - v7)) / v7)))
            var v20 uint16 = 1
            fmt.Println(v6)
            _ = fi19
            _ = v20
        }
    case "a":
        var v21 uint16 = 42
        _ = v21
    default:
        fmt.Println(v6)
    }
    v22 := []int{128}
    fmt.Println(v1, v4)
    switch ("a" + "\n\t") {
    case "":
        if ("foo bar" < ("" + ("foo bar" + ""))) {
            v4 = true
            v23 := GenS3{F0: 0, F1: -45, F2: 128}
            _ = v23
        } else {
            v24 := v2[""]
            switch ("hello" + "\n\t") {
            case "hello":
                v5 += 3.14159
            case "foo bar":
                v6 = (v6 + v6)
            default:
                var v25 uint = v1
                _ = v25
            }
            if false {
                v26 := GenS3{F0: (v1 + v1), F1: -1, F2: 32767}
                var v27 uint = v1
                _ = v26
                _ = v27
            } else {
                v28 := GenS4{F0: 27, F1: 2147483648}
                _ = v28
            }
            _ = v24
        }
    default:
        var v29 int = 0
        _ = v29
    }
    var v30 int32 = 191
    v7 = float32((v3 - v3))
    v3 -= 1e10
    switch "\n\t" {
    case "":
        for fi31 := 0; fi31 < 15; fi31++ {
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 float64 = v3
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
