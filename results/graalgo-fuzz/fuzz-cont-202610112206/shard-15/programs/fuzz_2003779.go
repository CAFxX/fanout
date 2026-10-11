package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 bool = true
    var v2 uint64 = 42
    var v3 uint32 = 42
    v4 := []int{32767, 32767, 948}
    var v5 int16 = 534
    v6 := GenS1{F0: 128, F1: ((1.5 - 1e10) >= 0.1), F2: (true || v1)}
    var v7 int16 = (v5 / ((v5) | 1))
    var v8 float32 = -2.25
    v9 := GenS1{F0: 458, F1: v1, F2: (v2 > (v2 | 4294967295))}
    v5 -= v5
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int32 = 2147483647
        for fi12 := 0; fi12 < 1; fi12++ {
            if ((v2 ^ v2) < v2) {
                v8 = float32((-2.25 * 1e10))
            }
            if (2 == 119) {
                v13 := GenS1{F0: (fi10 >> ((0) & 63)), F1: (v1 || v1), F2: v1}
                if len(v4) > 0 {
                    v4[0] = int(v3)
                }
                _ = len(v4)
                _ = v13
            } else {
                v14 := GenS1{F0: (^(fi12)), F1: ((0 == 613) || false), F2: v1}
                if (9223372036854775807 <= (fi10 - fi10)) {
                    v15 := GenS1{F0: fi12, F1: v1, F2: (false || (!((v2 ^ 624) == v2)))}
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v16 uint8 = 0
                        _ = v16
                    case 4294967295:
                        v1 = false
                    case 32767:
                        switch "" {
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    var v17 int = -717
                    _ = v15
                    _ = v17
                }
                var v18 float32 = (1e10 + -2.25)
                _ = v14
                _ = v18
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v19 uint8 = 2
    switch "a" {
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: (fi20 ^ (fi20 / ((-2147483648) | 1))), F1: (v3 == 4294967295), F2: (!("\n\t" == ((("" + (((((("" + "") + "") + ((("\n\t" + "a") + "foo bar") + "\n\t")) + ("" + "hello")) + (("" + "\n\t") + "hello")) + ("a" + "\n\t"))) + (("" + "hello") + "\n\t")) + "")))}
            _ = fi20
            _ = v21
        }
    case "hello":
        fmt.Println(v7, v2, v7)
    case "\n\t":
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                fmt.Println(v8)
                _ = fi23
            }
            _ = fi22
        }
    }
    v24 := GenS2{F0: v19}
    _ = len(v4)
    v2 -= v2
    v25 := GenS3{F0: (v5 << ((0) & 63)), F1: 677, F2: 107}
    v8 += 0.0
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
