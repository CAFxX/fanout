package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0(p0 string, p1 int) uint32 {
    v5 := []int{9223372036854775807, p1, (p1 ^ 128)}
    v6 := GenS1{F0: "", F1: 0}
    var v7 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 float64 = 1.5
    var v2 int16 = 187
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    if (840 != (v2 >> ((625) & 63))) {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: ("\n\t" + "a"), F1: 0}
            v4["foo bar"] = (4294967295 < 968)
            _ = fi8
            _ = v9
        }
        v10 := []int{242, 1, 32768, -9223372036854775808}
        v2 += (v2 * v2)
        _ = v10
    } else {
        fmt.Println(v1, v1)
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS2{F0: (v2 - 1000)}
            if true {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v1 -= v1
                    var v14 float64 = (1.5 * v1)
                    fmt.Println(fi13, v1)
                    _ = fi13
                    _ = v14
                }
                v15 := []int{-2147483648, 1, 2147483647}
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    }
    switch "\n\t" {
    case "hello":
        v16 := GenS2{F0: -1}
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 2; fi17++ {
            switch "a" {
            case "foo bar":
                var v18 int = -573
                _ = v18
            case "a":
                v2 -= 1000
            case "\n\t":
                v4["foo bar"] = ((0.1 <= -2.25) || (false && (255 != 65535)))
            default:
                v19 := []int{fi17, -495}
                _ = v19
            }
            _ = fi17
        }
    case "foo bar":
        v2 += v2
    default:
        fmt.Println(v2, v1)
    }
    v2 = (v2 - -1)
    v20 := GenS1{F0: "foo bar", F1: 4294967295}
    if true {
        if false {
            fmt.Println(v1, v1)
            fmt.Println(v1, v2)
        }
    }
    v21 := GenS3{F0: v2, F1: (2 > 839), F2: 42}
    delete(v3, ((("hello" + "") + "\n\t") + "hello"))
    if false {
        v22 := v4["hello"]
        _ = v22
    }
    v2 -= (v2 - (v2 ^ v2))
    v23 := []int{42, 9223372036854775807}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
