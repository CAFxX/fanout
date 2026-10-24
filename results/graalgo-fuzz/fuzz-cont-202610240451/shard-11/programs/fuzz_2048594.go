package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "a"
    var v3 uint64 = 4294967295
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]int)
    var v6 uint32 = 445
    v3 -= 4294967295
    v3 += v3
    var v7 int16 = 32767
    if (((-128 == -54) || (662 != -1)) || ((!true) || true)) {
        var v8 uint32 = 347
        var v9 uint = 65535
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint32 = v6
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v2)
    var v12 int32 = -2147483648
    switch "\n\t" {
    case "\n\t":
        v2 = "foo bar"
    default:
        var v13 int16 = v7
        _ = v13
    }
    if (255 < 2) {
        var v14 uint64 = (v3 & (v3 - 4294967295))
        var v15 float32 = float32(-2.25)
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 bool = false
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    } else {
        var v18 string = v2
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: "a", F1: 193}
            for fi21 := 0; fi21 < 6; fi21++ {
                if ((^(fi21)) <= fi21) {
                    fmt.Println(v2, v2)
                } else {
                    var v22 int = int(v3)
                    _ = v22
                }
                switch v2 {
                case "\n\t":
                    switch ("a" + ("foo bar" + "foo bar")) {
                    case "":
                        v23 := []int{(fi21 << ((0) & 63)), 32767, int(v6), (fi21 >> ((2) & 63))}
                        _ = v23
                    }
                case "":
                    var v24 int16 = v7
                    _ = v24
                case "a":
                    v5[fi19] = fi19
                }
                _ = fi21
            }
            v4 = v4
            _ = fi19
            _ = v20
        }
        v18 = v18
        _ = v18
    }
    v25 := []int{}
    var v26 int8 = -1
    v27 := GenS1{F0: "hello", F1: 128}
    var v28 uint = 775
    switch 1 {
    case 2147483647:
        var v29 int = 0
        _ = v29
    case -242:
        for fi30 := 0; fi30 < 4; fi30++ {
            v2 = ""
            _ = fi30
        }
    case -2147483648:
        _ = len(v5)
    }
    var v31 int8 = (v26 % ((v26) | 1))
    if (!false) {
        v32 := GenS1{F0: (v2 + v2), F1: uint8(v3)}
        _ = v32
    }
    v3 += (v3 ^ 18446744073709551615)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
