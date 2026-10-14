package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 string, p1 int64) uint32 {
    v7 := GenS1{F0: 32767, F1: 127}
    v8 := GenS1{F0: 963, F1: 0}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint32) uint16 {
    switch ("foo bar" + "a") {
    case "foo bar":
        var v9 string = ((("foo bar" + "foo bar") + (("a" + ("hello" + ("\n\t" + "a"))) + "a")) + "a")
        _ = v9
    }
    p0 += 185
    _ = p0
    return 0
}

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    v4 := []int{0, 128, 1, 1}
    v5 := []int{127, -9223372036854775808}
    var v6 uint = 768
    var v10 uint8 = 243
    fmt.Println(v10, v6, v10)
    var v11 float32 = 1e10
    fmt.Println(v3, v11, v11)
    v3 = v3
    var v12 uint8 = v10
    fmt.Println(v3, v12, v12)
    switch v3 {
    case "foo bar":
        if ((v3 + v3) < ((v3 + "\n\t") + v3)) {
            var v13 uint32 = 114
            if true {
                if (false && false) {
                    switch -606 {
                    case -594:
                        v1 = append(v1, "a")
                    default:
                        fmt.Println(v3, v6, v12)
                    }
                }
            }
            _ = v13
        } else {
            var v14 bool = true
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if len(v5) > 0 {
            v5[0] = -202
        }
        _ = fi15
    }
    switch v3 {
    case "foo bar":
        switch (v3 + "foo bar") {
        case "a":
            if ((2147483648 <= 424) || true) {
                fmt.Println(v3, v3)
            } else {
                fmt.Println(v3)
                v10 = (v10 | (v10 + 232))
            }
        case "foo bar":
            v3 = "\n\t"
        default:
            v16 := GenS2{F0: 2, F1: 0}
            _ = v16
        }
    case "":
        switch -9223372036854775808 {
        case -1:
            if (2 < 2) {
                v3 = v3
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case 128:
            var v17 int = -114
            _ = v17
        default:
            v3 = (v3 + v3)
        }
    default:
        v18 := GenS3{F0: ((0 < 0) || true), F1: v3}
        _ = v18
    }
    v19 := GenS3{F0: false, F1: "hello"}
    for fi20 := 0; fi20 < 1; fi20++ {
        v2[v3] = fi20
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
