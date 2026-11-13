package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    var v1 int64 = 0
    var v2 string = "foo bar"
    var v3 int64 = (v1 & v1)
    v4 := GenS1{F0: (v2 + v2), F1: v3}
    var v5 int8 = 28
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6, v1)
        v5 -= v5
        v7 := GenS1{F0: (v2 + v2), F1: (v3 >> ((18446744073709551615) & 63))}
        _ = fi6
        _ = v7
    }
    v8 := []int{32768, 822, 661}
    v9 := GenS1{F0: v2, F1: v3}
    for fi10 := 0; fi10 < 16; fi10++ {
        v1 += (v3 % (((^(v1))) | 1))
        var v11 float32 = 0.1
        _ = fi10
        _ = v11
    }
    v1 -= 1
    if ((v2 + ((v2 + (v2 + v2)) + "hello")) > "a") {
        if (float32(0.0) >= -2.25) {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint32 = 399
                var v14 bool = (true || false)
                if v14 {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    if len(v8) > 0 {
                        v8[0] = 2147483647
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            if (9223372036854775807 == 4294967295) {
                _ = len(v8)
                v5 += v5
                if (true && (false && (1 < 2147483647))) {
                    v1 -= v1
                    switch v2 {
                    case "a":
                        v15 := GenS1{F0: (v2 + v2), F1: (^(v1))}
                        _ = v15
                    case "hello":
                        if (!(true && (127 > 739))) {
                            var v16 float32 = 1.5
                            v16 = v16
                            _ = v16
                        }
                    case "\n\t":
                        var v17 uint8 = 148
                        _ = v17
                    }
                } else {
                    v2 = v2
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS1{F0: ("\n\t" + v2), F1: v1}
                        var v20 bool = (1 != 128)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 int8 = -1
                    _ = v21
                }
            }
            switch 454 {
            case 32768:
                v22 := GenS1{F0: v2, F1: (v3 + (^(v3)))}
                _ = v22
            case 42:
                v3 -= int64(v5)
            }
        }
    }
    v8 = append(v8, -475)
    v3 = -1
    var v23 uint32 = 65535
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
