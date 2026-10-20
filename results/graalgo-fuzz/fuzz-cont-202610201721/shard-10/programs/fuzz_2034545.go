package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
func genFn0(p0 uint, p1 uint8) uint8 {
    p0 = (p0 << ((p0) & 63))
    var v4 uint16 = 42
    if true {
        v5 := []int{-621, -590, 1}
        _ = v5
    } else {
        var v6 uint64 = uint64(p1)
        var v7 int32 = 2147483647
        _ = v6
        _ = v7
    }
    switch -1 {
    case 882:
        v8 := GenS1{F0: (v4 & (v4 | v4)), F1: ("foo bar" + ("hello" + "foo bar")), F2: true}
        _ = v8
    case 32768:
        v9 := []int{int(v4), -954, -290}
        _ = v9
    case -9223372036854775808:
        v10 := []int{160, 2147483647, -223, 9223372036854775807}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 * (p1 & (p1 - p1)))
}

func main() {
    v1 := []string{"hello", (((("hello" + ((("a" + "\n\t") + "hello") + "")) + "") + "\n\t") + "hello"), "\n\t"}
    v2 := []uint32{0}
    v3 := make(map[string]bool)
    v3[("foo bar" + "a")] = (true && (("a" + "a") == "hello"))
    var v11 int64 = -9223372036854775808
    var v12 int32 = 1
    v12 -= 2147483647
    v11 = v11
    v13 := GenS1{F0: 32767, F1: "a", F2: ((!false) && ((1.5 - (1e10 + (0.0 + (((0.0 / -2.25) * 3.14159) * 1.5)))) >= 3.14159))}
    switch "hello" {
    case "foo bar":
        var v14 int32 = 0
        _ = v14
    }
    v12 = (v12 + 1)
    if false {
        v11 = (v11 / (((v11 ^ 42)) | 1))
        v15 := v3["foo bar"]
        if (v11 < (v11 * v11)) {
            v12 -= v12
        } else {
            v12 = v12
            var v16 uint64 = 301
            _ = v16
        }
        _ = v15
    } else {
        switch ("foo bar" + "\n\t") {
        case "":
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 32767, F1: "foo bar", F2: (161 == genFn0(579, 7))}
                _ = fi17
                _ = v18
            }
        default:
            v19 := GenS1{F0: 1, F1: "foo bar", F2: true}
            _ = v19
        }
    }
    var v20 float32 = -2.25
    var v21 int16 = 0
    v22 := GenS1{F0: 63, F1: "foo bar", F2: false}
    v21 += (^(v21))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
