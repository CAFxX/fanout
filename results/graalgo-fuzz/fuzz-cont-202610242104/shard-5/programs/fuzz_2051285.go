package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16, p1 bool, p2 uint64) int {
    var v4 float64 = 0.1
    fmt.Println(v4)
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 807}
        _ = fi5
        _ = v6
    }
    if (!p1) {
        p0 += (p0 | (p0 - p0))
        v4 = (-2.25 - (1e10 * ((v4 - ((v4 * 1.5) - 0.1)) - v4)))
        v7 := []int{32767, 32768}
        _ = v7
    } else {
        var v8 uint = 705
        _ = v8
    }
    v9 := []int{128, 2147483647, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return 32768
}
func genFn1(p0 uint32, p1 int) string {
    p1 += 2147483647
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(p1, fi10)
        p1 += (fi10 - fi10)
        p0 += 615
        _ = fi10
    }
    _ = p0
    _ = p1
    return (("\n\t" + "hello") + "\n\t")
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"foo bar", ("\n\t" + (("foo bar" + ("hello" + ((("\n\t" + "foo bar") + "a") + ""))) + (("a" + ((("foo bar" + "hello") + "hello") + ("foo bar" + "\n\t"))) + "a"))), "hello", "foo bar"}
    v3 := []uint32{348, 1, 65535}
    var v11 uint32 = 2
    v11 += v11
    var v12 uint16 = 32767
    v13 := GenS2{F0: 128}
    v11 -= v11
    if false {
        v14 := GenS2{F0: 128}
        delete(v1, (v11 - v11))
        if false {
            var v15 float64 = float64(3.14159)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v16 int64 = 9223372036854775807
            _ = v15
            _ = v16
        }
        _ = v14
    }
    switch "foo bar" {
    case "hello":
        var v17 uint64 = 4294967295
        _ = v17
    case "\n\t":
        v11 = (v11 / (((v11 - 2147483647)) | 1))
    case "a":
        v11 = v11
    }
    v18 := v1[v11]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 := GenS3{F0: -1, F1: (0 <= 0)}
    v20 := GenS1{F0: 850}
    v18 = "foo bar"
    v21 := GenS2{F0: int(v11)}
    if len(v3) > 0 {
        v3[0] = (^(v11))
    }
    if (2 >= 84) {
        fmt.Println(v11, v12)
    } else {
        v22 := GenS4{F0: 18446744073709551615}
        fmt.Println(v18, v18, v12)
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
