package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 string }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0() int32 {
    if true {
        v4 := []int{9223372036854775807, -512, 32767, 127}
        _ = v4
    }
    var v5 uint32 = 461
    fmt.Println(v5)
    _ = v5
    return -488
}

func main() {
    var v1 int8 = 42
    var v2 int8 = -77
    var v3 uint32 = 146
    v6 := GenS1{F0: 0, F1: 4294967295, F2: "hello"}
    v7 := []int{-570}
    var v8 uint64 = 0
    switch 2147483647 {
    case 32767:
        var v9 bool = (-9223372036854775808 < 1)
        _ = v9
    case 4294967295:
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(v1, v2, fi10)
            v8 = v8
            v2 += v2
            _ = fi10
        }
    }
    switch ("\n\t" + ("a" + "a")) {
    case "a":
        v1 = v2
    default:
        v2 += (v2 | (v1 - v2))
    }
    var v11 int16 = 32767
    v1 = (^(v2))
    switch 489 {
    case 2147483647:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int = fi12
            if (v13 <= fi12) {
                v14 := GenS1{F0: 1, F1: 42, F2: (("hello" + "foo bar") + (("" + "\n\t") + "a"))}
                var v15 int64 = 204
                _ = v14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    }
    var v16 int32 = 2147483647
    v7 = append(v7, 9223372036854775807)
    var v17 uint64 = (v8 % (((v8 & v8)) | 1))
    for fi18 := 0; fi18 < 13; fi18++ {
        if false {
            var v19 uint8 = 131
            var v20 int16 = v11
            _ = v19
            _ = v20
        } else {
            v11 += v11
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 int64 = 584
                switch ("\n\t" + "") {
                case "foo bar":
                    v8 -= v8
                    if (128 > fi21) {
                        var v23 int = fi18
                        v1 += v1
                        v22 += v22
                        _ = v23
                    }
                    v24 := GenS2{F0: v11, F1: 42}
                    _ = v24
                default:
                    var v25 uint8 = 105
                    _ = v25
                }
                v2 = v1
                _ = fi21
                _ = v22
            }
        }
        fmt.Println(fi18)
        v26 := GenS2{F0: v11, F1: v1}
        _ = fi18
        _ = v26
    }
    v7 = append(v7, -2147483648)
    v8 = (v8 % ((v17) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
