package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 int; F2 int16 }
func genFn0() string {
    var v6 uint32 = 2147483647
    v6 -= 4294967295
    v7 := []int{128, -241, 128}
    _ = v6
    _ = v7
    return "a"
}

func main() {
    var v1 int = -39
    v2 := make(map[string]bool)
    v3 := []string{"\n\t", ("" + "a"), "foo bar", ("a" + "a")}
    v4 := []uint32{42, 314, 42}
    var v5 float32 = -2.25
    v8 := v2[("\n\t" + "hello")]
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v1 ^ 174) {
    case -217:
        if (1 <= 42) {
            var v9 int32 = 1
            _ = v9
        }
    default:
        _ = len(v2)
    }
    v1 = (-(v1))
    _ = len(v2)
    fmt.Println(v5)
    for fi10 := 0; fi10 < 5; fi10++ {
        if v8 {
            v5 -= ((1.5 + v5) + float32((0.1 - 0.1)))
        }
        v5 -= v5
        _ = fi10
    }
    v8 = v8
    v11 := GenS1{F0: v1, F1: 29, F2: 621}
    _ = len(v2)
    v12 := GenS2{F0: v1, F1: (v1 - (v1 << ((42) & 63))), F2: 96}
    v13 := GenS2{F0: 32768, F1: (v1 >> ((4294967295) & 63)), F2: -198}
    v14 := GenS2{F0: v1, F1: 0, F2: -76}
    if v8 {
        if v8 {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        } else {
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    switch fi16 {
                    case 32767:
                        v5 -= float32((3.14159 - 3.14159))
                    case 127:
                        fmt.Println(v8, fi15)
                    default:
                        v8 = v8
                    }
                    _ = fi16
                }
                switch (v1 - (-(fi15))) {
                case -486:
                    fmt.Println(v5, fi15, v8)
                case -131:
                    var v17 string = (("a" + (("foo bar" + "\n\t") + ("" + ("hello" + "")))) + "")
                    _ = v17
                case 4294967295:
                    v3 = append(v3, ("foo bar" + "foo bar"))
                }
                _ = fi15
            }
            var v18 bool = v8
            _ = v18
        }
        v8 = v8
    } else {
        v8 = (2 >= 0)
        v19 := GenS2{F0: v1, F1: v1, F2: -32768}
        var v20 int16 = 1000
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
