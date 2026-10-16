package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 int16 }
func genFn0(p0 int16) uint16 {
    v7 := GenS1{F0: (p0 / (((p0 >> ((65535) & 63))) | 1)), F1: "\n\t", F2: -128}
    var v8 float64 = 3.14159
    switch 67 {
    case 0:
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := []int{}
                    v12 := GenS2{F0: -1, F1: 9223372036854775807, F2: (p0 & 42)}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v8 += (v8 / 0.1)
                var v13 uint = 65535
                _ = v13
            } else {
                var v14 uint32 = 42
                if true {
                    fmt.Println(v8)
                    var v15 float32 = 1.5
                    v8 += v8
                    _ = v15
                } else {
                    var v16 uint16 = 1
                    v17 := []int{fi9, (-(fi9)), fi9}
                    var v18 uint32 = v14
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
            v19 := []int{(fi9 + fi9), fi9, (fi9 ^ (^(fi9))), (fi9 + (fi9 / ((2147483647) | 1)))}
            if true {
                var v20 float64 = 1e10
                _ = v20
            } else {
                var v21 uint8 = 255
                if (uint(v21) > 65535) {
                    v22 := GenS2{F0: -333, F1: -1, F2: p0}
                    for fi23 := 0; fi23 < 16; fi23++ {
                        switch 32768 {
                        case 32768:
                            v21 = v21
                        }
                        _ = fi23
                    }
                    _ = v22
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = v21
            }
            _ = fi9
            _ = v19
        }
    }
    v8 = v8
    for fi24 := 0; fi24 < 20; fi24++ {
        switch ((("a" + "foo bar") + "a") + ("foo bar" + (("" + "hello") + "foo bar"))) {
        case "\n\t":
            var v25 bool = (545 > 255)
            _ = v25
        }
        v8 = (1e10 + (v8 - v8))
        _ = fi24
    }
    _ = p0
    _ = v7
    _ = v8
    return 692
}

func main() {
    var v1 uint8 = 121
    var v2 string = (("hello" + (("a" + ((("hello" + "foo bar") + ("hello" + "foo bar")) + "hello")) + ("\n\t" + ""))) + "a")
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = uint64(v1)
    v5 := []string{v2, "", v2, v2}
    v6 := make(map[int]bool)
    var v26 int16 = -837
    v27 := GenS1{F0: v26, F1: v2, F2: 0}
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(fi28)
        var v29 uint16 = 1
        _ = fi28
        _ = v29
    }
    v2 = (v2 + ("\n\t" + (v2 + v2)))
    var v30 float64 = 1e10
    switch "hello" {
    case "foo bar":
        fmt.Println(v2, v26, v4)
    default:
        v6[-126] = true
    }
    v4 += (^(v3))
    var v31 float32 = float32(v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
