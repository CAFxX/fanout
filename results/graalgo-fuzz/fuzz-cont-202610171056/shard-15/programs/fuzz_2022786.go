package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-988}
    var v3 float32 = 3.14159
    var v4 int = 2147483647
    v5 := GenS1{F0: -9223372036854775808}
    var v6 int64 = int64(v4)
    var v7 int = -1
    v8 := GenS1{F0: (v4 % (((v4 / ((42) | 1))) | 1))}
    v9 := GenS2{F0: -1, F1: 255, F2: -2147483648}
    v10 := GenS3{F0: (v7 | 1), F1: "a"}
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v3, v6, fi12)
        _ = fi12
    }
    v6 += -847
    if (1 < 2) {
        _ = len(v2)
        var v13 uint8 = 42
        v14 := GenS4{F0: 876}
        _ = v13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: v7}
            for fi17 := 0; fi17 < 15; fi17++ {
                if (394 != 32767) {
                    v6 -= (v6 / ((v6) | 1))
                    var v18 uint64 = 739
                    _ = v18
                } else {
                    var v19 uint8 = 210
                    v20 := GenS3{F0: fi17, F1: ("\n\t" + "a")}
                    var v21 bool = (!(((!((-1 > 114) || (false && false))) && ((v6 << ((18446744073709551615) & 63)) != (^(v6)))) || false))
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi17
            }
            var v22 string = ((("hello" + "a") + "") + "a")
            _ = fi15
            _ = v16
            _ = v22
        }
        if (!true) {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 int = (v4 * 127)
                if (!(false || false)) {
                    var v25 string = ""
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v26 bool = false
                    _ = v25
                    _ = v26
                } else {
                    var v27 int32 = -2147483648
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
        }
        v4 -= v4
    } else {
        var v28 float32 = (0.0 + 1.5)
        var v29 float32 = float32(v11)
        v30 := GenS2{F0: 1, F1: 0, F2: 0}
        _ = v28
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
