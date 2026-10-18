package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "foo bar"}
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    var v5 uint8 = 101
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := v4[("a" + ("a" + (("\n\t" + ((("hello" + "\n\t") + "foo bar") + "hello")) + "hello")))]
            _ = fi6
            _ = v7
        }
    }
    fmt.Println(v5, v5)
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = 3.14159
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v3, v5)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v11 := v4[("hello" + "hello")]
    v2 = append(v2, (v11 + v11))
    if false {
        var v12 uint32 = 2147483648
        for fi13 := 0; fi13 < 14; fi13++ {
            v2 = append(v2, ("hello" + ""))
            v3 += v3
            var v14 uint32 = v12
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 bool = false
                v5 = 1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi16
                _ = v17
            }
            v18 := v4[("a" + "a")]
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint16 = 49
                v5 += 2
                _ = fi19
                _ = v20
            }
            _ = fi15
            _ = v18
        }
        _ = v12
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = 2147483648
    if ((float64(((1.5 * ((((0.1 * (0.1 - 1e10)) / float32((1.5 - -2.25))) + -2.25) * 3.14159)) + -2.25)) > 3.14159) && (!true)) {
        for fi21 := 0; fi21 < 18; fi21++ {
            v3 -= v3
            _ = fi21
        }
        _ = len(v1)
    }
    if true {
        var v22 bool = (true && (!false))
        v23 := GenS1{F0: (!v22), F1: -55, F2: -875}
        v3 = 2
        _ = v22
        _ = v23
    }
    if true {
        if (2147483647 > (v3 << ((42) & 63))) {
            if (255 < 247) {
                v2 = append(v2, v11)
                if true {
                    switch v11 {
                    case "\n\t":
                        delete(v4, "foo bar")
                    }
                    var v24 int64 = 297
                    _ = v24
                } else {
                    var v25 int32 = 576
                    _ = v25
                }
            }
            v11 = v11
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    v5 = 1
    fmt.Println(v5, v3, v3)
    v2 = append(v2, "")
    v26 := v4[(v11 + "a")]
    v27 := GenS1{F0: false, F1: 672, F2: -681}
    v28 := GenS2{F0: v26, F1: -128, F2: 0}
    var v29 float64 = (((3.14159 * 3.14159) * 3.14159) + ((3.14159 - -2.25) - (3.14159 - 0.1)))
    fmt.Println(v5, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
