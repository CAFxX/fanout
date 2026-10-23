package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0() uint64 {
    var v6 string = "a"
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 780, F1: 98}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: 4294967295}
        _ = fi9
        _ = v10
    }
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = -12
    var v3 float64 = 3.14159
    var v4 uint16 = v1
    v5 := make(map[string]string)
    v11 := v5["\n\t"]
    v1 = (v1 + (v1 >> ((65535) & 63)))
    switch v11 {
    case "":
        fmt.Println(v2)
    default:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int = fi12
            v5[(("a" + (v11 + "")) + "hello")] = "\n\t"
            v11 = "a"
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float32 = (0.0 * 1e10)
        _ = fi14
        _ = v15
    }
    v16 := []int{-2147483648}
    fmt.Println(v4, v1)
    v17 := v5[v11]
    var v18 uint64 = 0
    for fi19 := 0; fi19 < 15; fi19++ {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            if ((fi20 | (-(fi20))) == (fi20 | fi19)) {
                var v21 uint64 = (^(v18))
                if false {
                    var v22 int16 = -32768
                    _ = v22
                }
                v16 = append(v16, (-(fi19)))
                _ = v21
            } else {
                if true {
                    v23 := GenS2{F0: uint64(v1)}
                    _ = v23
                } else {
                    switch (fi19 >> ((uint(v18)) & 63)) {
                    case 2147483647:
                        var v24 int8 = 100
                        _ = v24
                    case 779:
                        if (65535 <= 65535) {
                            var v25 float32 = 0.0
                            _ = v25
                        }
                    case -2147483648:
                        var v26 float64 = (v3 + (3.14159 / v3))
                        _ = v26
                    }
                    v1 = 65535
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                switch -1 {
                case -419:
                    _ = len(v16)
                default:
                    var v27 uint = 0
                    _ = v27
                }
            }
            _ = fi20
        }
        switch ((("\n\t" + (v17 + v17)) + v17) + v17) {
        case "a":
            fmt.Println(fi19, v11)
        case "":
            if false {
                v28 := GenS2{F0: v18}
                _ = v28
            }
        case "\n\t":
            v4 += (v1 % (((v1 >> ((62) & 63))) | 1))
        }
        _ = fi19
    }
    var v29 int = 127
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
