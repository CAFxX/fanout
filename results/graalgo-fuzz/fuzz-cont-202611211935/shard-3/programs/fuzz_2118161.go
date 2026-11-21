package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 int64 = 722
    v4 := []int{-1}
    v2 = (v2 / (v2 + 0.1))
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int32 = -920
            _ = fi6
            _ = v7
        }
        v4 = append(v4, (fi5 + fi5))
        _ = fi5
    }
    if len(v4) > 0 {
        v4[0] = -195
    }
    v3 += v3
    v2 = (1.5 / 1e10)
    v8 := GenS1{F0: 18446744073709551615, F1: (("" + ("\n\t" + "foo bar")) + "hello")}
    v2 -= (3.14159 / -2.25)
    fmt.Println(v2)
    var v9 int64 = (v3 | 0)
    fmt.Println(v9)
    v9 = (-(v3))
    var v10 float32 = 1e10
    fmt.Println(v10, v9, v10)
    switch "\n\t" {
    case "foo bar":
        v3 = (v3 & v3)
    case "hello":
        v2 = v10
    }
    if (v9 <= (v9 & (v3 * (^(v9))))) {
        for fi11 := 0; fi11 < 1; fi11++ {
            if (((((v9 - v3) > 105) || false) || true) && ((3.14159 == (0.0 * 1e10)) || false)) {
                v10 -= float32(0.1)
            } else {
                var v12 uint64 = 18446744073709551615
                if false {
                    if len(v4) > 0 {
                        v4[0] = -258
                    }
                    if ((v10 - 1.5) >= v2) {
                        v3 -= v3
                        if (v10 == float32((float64((((0.1 * v10) + float32(1.5)) * float32(1e10))) - 3.14159))) {
                            var v13 bool = false
                            _ = v13
                        } else {
                            if false {
                                var v14 uint = 799
                                _ = v14
                            }
                            fmt.Println(v10, v3)
                        }
                    }
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    if len(v4) > 0 {
                        v4[0] = fi15
                    }
                    var v16 uint8 = 232
                    _ = fi15
                    _ = v16
                }
                _ = v12
            }
            var v17 int8 = 0
            fmt.Println(v17, v9)
            _ = fi11
            _ = v17
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v18 := GenS1{F0: 2, F1: (("foo bar" + "foo bar") + ("\n\t" + "foo bar"))}
        _ = v18
    }
    v3 = v9
    if (1 > 1) {
        if (((0.1 - 0.0) * (1e10 - -2.25)) <= ((3.14159 - 1e10) - (1e10 / -2.25))) {
            var v19 int32 = 249
            v20 := v1["hello"]
            _ = v19
            _ = v20
        }
        v2 = v2
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS2{F0: 4294967295}
            v23 := GenS2{F0: 1}
            var v24 float32 = 0.1
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 uint = 18446744073709551615
            var v27 int = -2147483648
            v28 := GenS2{F0: 877}
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
