package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 string }
func genFn0(p0 uint64, p1 uint64) int8 {
    p1 = p0
    p1 -= p1
    _ = p0
    _ = p1
    return 25
}

func main() {
    var v1 uint8 = 149
    v2 := make(map[int]uint32)
    v3 := []uint32{124, 2147483648}
    v4 := []string{"\n\t", (("foo bar" + ((("\n\t" + ("foo bar" + "hello")) + "foo bar") + "\n\t")) + "")}
    var v5 string = ("foo bar" + "hello")
    if len(v4) > 0 {
        v4[0] = v5
    }
    var v6 int16 = -1
    var v7 uint64 = 4294967295
    var v8 uint64 = v7
    v8 -= 985
    v9 := GenS1{F0: v6, F1: -9223372036854775808}
    v10 := v2[-2147483648]
    _ = len(v4)
    fmt.Println(v8, v7, v6)
    if false {
        fmt.Println(v8)
        v6 += v6
    } else {
        if true {
            var v11 int64 = -9223372036854775808
            _ = v11
        }
        v12 := GenS1{F0: v6, F1: 457}
        _ = v12
    }
    _ = len(v3)
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint8 = (^(v1))
        fmt.Println(v6)
        v15 := GenS1{F0: (v6 * (v6 * -32768)), F1: 1}
        _ = fi13
        _ = v14
        _ = v15
    }
    v10 = 4294967295
    _ = len(v2)
    if true {
        switch -1 {
        case 127:
            for fi16 := 0; fi16 < 12; fi16++ {
                fmt.Println(v5)
                var v17 int64 = 9223372036854775807
                _ = fi16
                _ = v17
            }
        case 586:
            if false {
                var v18 uint16 = 0
                switch v5 {
                case "a":
                    fmt.Println(v1)
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = (((v5 + "hello") + v5) + v5)
                    }
                }
                _ = v18
            } else {
                v2[0] = (v10 | (v10 ^ 878))
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS2{F0: true, F1: (fi19 * fi19), F2: (fi19 + (^(fi19)))}
                    _ = fi19
                    _ = v20
                }
            }
        case 2147483647:
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 float32 = (float32((3.14159 * float64(3.14159))) * 1.5)
                _ = fi21
                _ = v22
            }
        }
        v7 = (v7 % ((v7) | 1))
        switch (v5 + (v5 + v5)) {
        case "\n\t":
            v23 := GenS3{F0: v5}
            _ = v23
        default:
            v24 := GenS3{F0: "a"}
            _ = v24
        }
    }
    fmt.Println(v5, v8, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
