package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0() uint {
    var v5 uint64 = 18446744073709551615
    v5 = v5
    _ = v5
    return 635
}
func genFn1(p0 uint8) uint8 {
    p0 = 255
    v6 := []int{32767, int(p0), 32768}
    p0 = (p0 | (p0 - (p0 % (((p0 % (((p0 % (((^(p0))) | 1))) | 1))) | 1))))
    var v7 int = -249
    _ = p0
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a", ("a" + (("\n\t" + "") + "a"))}
    v2 := make(map[uint32]int)
    v3 := []int{}
    v4 := []int{-681}
    switch ("a" + "a") {
    case "a":
        v8 := GenS1{F0: 0}
        _ = v8
    case "hello":
        var v9 float64 = 3.14159
        _ = v9
    default:
        var v10 float32 = 3.14159
        _ = v10
    }
    v2[2147483647] = -2147483648
    var v11 float64 = float64(0.0)
    v11 -= v11
    var v12 string = ("a" + "foo bar")
    if (361 != 311) {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint8 = 1
            _ = len(v2)
            if (false && (102 == 0)) {
                v15 := GenS2{F0: "foo bar"}
                if false {
                    fmt.Println(v11, v11)
                } else {
                    var v16 float64 = v11
                    _ = v16
                }
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    } else {
        var v17 int8 = 101
        v12 = ((v12 + "foo bar") + "\n\t")
        var v18 uint = 952
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "foo bar":
        var v19 int = -9223372036854775808
        _ = v19
    case "\n\t":
        switch -1 {
        case 2147483647:
            v20 := v2[42]
            _ = v20
        case 1:
            v12 = (v12 + v12)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
