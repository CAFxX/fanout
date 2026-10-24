package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 bool, p1 int32, p2 uint32) string {
    var v5 bool = p0
    v6 := []int{-366}
    switch "" {
    case "\n\t":
        var v7 uint64 = 0
        _ = v7
    case "":
        v5 = ((3.14159 / -2.25) >= 1.5)
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = 1
        }
    }
    var v8 float32 = (float32(1e10) / 1.0)
    var v9 float32 = (v8 * (-2.25 + 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 string) uint64 {
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{fi10}
        p0 = p0
        var v12 int16 = -148
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = ((p0 + ((((p0 + p0) + "") + (p0 + "hello")) + "foo bar")) + p0)
    switch 9223372036854775807 {
    case 738:
        v13 := GenS1{F0: p0, F1: 0, F2: 128}
        _ = v13
    default:
        v14 := []int{2147483647, 128, 1}
        _ = v14
    }
    _ = p0
    return 24
}

func main() {
    v1 := []int{219, 127}
    var v2 int64 = -9223372036854775808
    v3 := []int{}
    var v4 uint8 = 42
    v15 := GenS1{F0: "foo bar", F1: 0, F2: v4}
    fmt.Println(v2)
    switch 401 {
    case 456:
        var v16 float64 = float64((0.0 - (-2.25 - (0.1 / (-2.25 - 1.5)))))
        _ = v16
    default:
        v2 -= (-(v2))
    }
    var v17 bool = (((v2 | (v2 * (-(v2)))) > v2) && true)
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint8 = v4
        _ = fi18
        _ = v19
    }
    v4 = (v4 * 2)
    if len(v1) > 0 {
        v1[0] = 127
    }
    v2 = v2
    var v20 uint = 4294967295
    var v21 bool = v17
    for fi22 := 0; fi22 < 7; fi22++ {
        switch fi22 {
        case 4294967295:
            v20 = uint(v4)
        case 32768:
            if v17 {
                var v23 float32 = 1e10
                for fi24 := 0; fi24 < 16; fi24++ {
                    v20 += 307
                    _ = fi24
                }
                _ = v23
            }
        case 127:
            if len(v3) > 0 {
                v3[0] = fi22
            }
        default:
            var v25 bool = false
            _ = v25
        }
        var v26 int8 = 44
        switch fi22 {
        case 0:
            var v27 uint64 = genFn1(("a" + ((("hello" + "hello") + ("foo bar" + "\n\t")) + ("foo bar" + ("\n\t" + ("a" + ""))))))
            _ = v27
        case 999:
            v2 = v2
        default:
            v20 = (^(v20))
        }
        _ = fi22
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
