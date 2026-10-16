package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint16 }

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 65535
    var v3 float64 = ((0.0 * 0.1) - 1.5)
    v4 := GenS1{F0: -9223372036854775808, F1: 0}
    if (v3 > v3) {
        if ("" == "\n\t") {
            for fi5 := 0; fi5 < 5; fi5++ {
                v6 := GenS1{F0: (fi5 / ((-897) | 1)), F1: -1}
                var v7 uint16 = 255
                v2 = 0
                _ = fi5
                _ = v6
                _ = v7
            }
            v8 := GenS1{F0: 398, F1: 42}
            v2 = (v2 % (((v2 - (^(v2)))) | 1))
            _ = v8
        } else {
            v3 += 1.5
            v2 = 42
            v1 -= 3.14159
        }
        if (!((!true) || ((^(v2)) > v2))) {
            v9 := GenS2{F0: 1, F1: (-32768 >= -32768), F2: ""}
            v10 := GenS3{F0: 959}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: 2147483647, F1: true, F2: "a"}
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint = 0
                var v14 int = 2147483647
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 int8 = 0
                    v3 = (v3 * (-2.25 * v3))
                    if (v16 == (v16 + (v16 ^ (-(v16))))) {
                        v2 += v2
                        fmt.Println(v13)
                        var v17 uint32 = 2
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v18 int = 55
            _ = v11
            _ = v18
        }
    }
    var v19 int16 = 582
    var v20 uint16 = 2
    v21 := []int{32768, 9223372036854775807, -918, -2147483648}
    for fi22 := 0; fi22 < 19; fi22++ {
        if ((4294967295 > 18446744073709551615) && true) {
            if true {
                v20 = (^(v20))
                fmt.Println(v20, v19)
                if len(v21) > 0 {
                    v21[0] = 32768
                }
            }
            v23 := GenS2{F0: 128, F1: false, F2: "\n\t"}
            _ = v23
        }
        _ = fi22
    }
    var v24 bool = (true || false)
    fmt.Println(v24, v24, v3)
    if false {
        var v25 uint64 = 714
        var v26 uint8 = 229
        v27 := GenS4{F0: 42, F1: v20}
        _ = v25
        _ = v26
        _ = v27
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
