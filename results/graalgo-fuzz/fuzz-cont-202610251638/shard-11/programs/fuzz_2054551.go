package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 int32 }

func main() {
    var v1 bool = true
    var v2 int32 = 0
    var v3 string = (((("foo bar" + (("" + ("" + "a")) + "")) + "a") + ("\n\t" + "")) + "a")
    var v4 uint16 = 65535
    var v5 int32 = int32(v4)
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        v1 = v1
        if (true && (4294967295 >= v6)) {
            v1 = v1
        }
        _ = fi7
    }
    if true {
        v8 := GenS1{F0: 42, F1: 42, F2: v6}
        v3 = "a"
        v9 := []int{203, -540, 4294967295}
        _ = v8
        _ = v9
    }
    v10 := []int{32768, -9223372036854775808, -155}
    var v11 uint = 700
    v2 += v5
    v12 := GenS1{F0: 0, F1: 2, F2: (v6 - 18446744073709551615)}
    var v13 uint64 = 679
    for fi14 := 0; fi14 < 8; fi14++ {
        if ("hello" < v3) {
            switch fi14 {
            case 2147483647:
                switch ((v3 + v3) + v3) {
                case "\n\t":
                    v15 := GenS1{F0: 4294967295, F1: 255, F2: 3}
                    _ = v15
                case "hello":
                    v16 := GenS1{F0: 2, F1: 169, F2: v13}
                    _ = v16
                case "foo bar":
                    v3 = v3
                }
            default:
                var v17 uint64 = (^(v6))
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS1{F0: 4294967295, F1: 128, F2: v13}
                switch fi18 {
                case 1:
                    v4 -= (v4 | v4)
                default:
                    var v20 int64 = -1
                    _ = v20
                }
                switch (fi14 * fi14) {
                case 931:
                    v5 = (v5 << (((v11 * v11)) & 63))
                }
                if true {
                    fmt.Println(v6)
                    if (!v1) {
                        v21 := GenS2{F0: (v6 ^ v13)}
                        var v22 uint32 = 350
                        if (v3 == v3) {
                            var v23 int64 = -9223372036854775808
                            var v24 int16 = -115
                            _ = v23
                            _ = v24
                        } else {
                            var v25 uint32 = (v22 ^ (^(v22)))
                            _ = v25
                        }
                        v2 = v2
                        var v26 uint64 = 18446744073709551615
                        _ = v21
                        _ = v22
                        _ = v26
                    } else {
                        v13 = 312
                        fmt.Println(v5)
                    }
                } else {
                    v5 = -2147483648
                }
                _ = fi18
                _ = v19
            }
        } else {
            var v27 uint = 42
            _ = v27
        }
        _ = fi14
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v6 += v13
        var v29 uint8 = 43
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: (^(v11)), F1: ((v3 + v3) + ("foo bar" + v3)), F2: -232}
    fmt.Println(v1, v3)
    for fi31 := 0; fi31 < 14; fi31++ {
        fmt.Println(v5)
        v13 += (v6 / ((v6) | 1))
        var v32 uint8 = 37
        _ = fi31
        _ = v32
    }
    v11 -= (v11 / ((uint(v4)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
