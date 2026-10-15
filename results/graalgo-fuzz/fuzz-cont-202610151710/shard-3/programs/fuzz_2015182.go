package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 string }

func main() {
    v1 := []int{128, -2147483648, 1}
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    v4 := []string{("a" + "\n\t")}
    v5 := make(map[int]string)
    if (18446744073709551615 < 42) {
        v3 = 171
        if (-9223372036854775808 != 9223372036854775807) {
            if (801 <= v3) {
                var v6 int32 = v3
                for fi7 := 0; fi7 < 9; fi7++ {
                    if len(v1) > 0 {
                        v1[0] = fi7
                    }
                    var v8 uint = 65535
                    v9 := GenS1{F0: 199, F1: (v3 % ((42) | 1))}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = v6
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v3 = (v3 * (^(v3)))
                    switch "" {
                    case "\n\t":
                        v3 += v3
                    }
                    var v11 uint32 = 2
                    _ = fi10
                    _ = v11
                }
            }
            v4 = append(v4, "")
        } else {
            v3 -= -2147483648
            fmt.Println(v3)
        }
    }
    var v12 uint64 = 421
    if len(v4) > 0 {
        v4[0] = ((("" + "hello") + ("foo bar" + "hello")) + "")
    }
    var v13 int = -628
    v1 = append(v1, -1)
    v14 := GenS2{F0: false, F1: 400}
    for fi15 := 0; fi15 < 6; fi15++ {
        v13 = fi15
        _ = fi15
    }
    if false {
        v12 += 4294967295
    }
    fmt.Println(v13, v13)
    v13 = (v13 - v13)
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            fmt.Println(v3, v3, v3)
            for fi17 := 0; fi17 < 18; fi17++ {
                v13 -= (fi17 % ((fi17) | 1))
                fmt.Println(v13, fi17)
                delete(v2, (fi16 % (((fi17 - (fi17 / ((fi17) | 1)))) | 1)))
                _ = fi17
            }
        } else {
            v18 := GenS3{F0: "", F1: 42, F2: ""}
            if true {
                var v19 int16 = 32767
                _ = v19
            }
            var v20 float64 = (-2.25 - 1.5)
            _ = v18
            _ = v20
        }
        v13 = -971
        var v21 int8 = int8(fi16)
        _ = fi16
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
