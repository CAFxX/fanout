package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
func genFn0() int32 {
    v7 := GenS1{F0: 320}
    if false {
        var v8 bool = (-2147483648 > 1)
        switch (("foo bar" + "foo bar") + ("foo bar" + (("a" + "") + ""))) {
        case "a":
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 19; fi10++ {
                    if false {
                        fmt.Println(fi9)
                        fmt.Println(fi9, fi9, fi9)
                        v11 := []int{(fi9 | 1), (fi10 | 642), 9223372036854775807}
                        _ = v11
                    }
                    v12 := []int{fi9, (fi9 + (fi9 / (((fi9 % ((1) | 1))) | 1))), fi9, (^(fi10))}
                    var v13 int = (fi10 % ((fi9) | 1))
                    _ = fi10
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        }
        _ = v8
    } else {
        var v14 float32 = (3.14159 - 1.5)
        var v15 int = 42
        for fi16 := 0; fi16 < 4; fi16++ {
            if ((42 < 376) || true) {
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 int = v15
                    var v19 bool = (false || (45 <= 444))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := []int{32768, fi16, (fi16 + (fi16 - -460))}
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = v20
            } else {
                v15 = 0
                var v21 int64 = 0
                _ = v21
            }
            var v22 string = "\n\t"
            v22 = "\n\t"
            _ = fi16
            _ = v22
        }
        _ = v14
        _ = v15
    }
    var v23 uint32 = 416
    _ = v23
    _ = v7
    return -976
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 int32 = 2147483647
    v4 := make(map[int]int)
    var v5 string = "hello"
    var v6 uint = 255
    v5 = v5
    v6 = v6
    var v24 uint64 = 18446744073709551615
    v25 := GenS1{F0: 32767}
    var v26 int64 = -319
    switch "hello" {
    case "":
        v27 := []int{968, -702, 128}
        _ = v27
    case "foo bar":
        var v28 int64 = v26
        _ = v28
    case "\n\t":
        v29 := GenS2{F0: -9223372036854775808, F1: 42, F2: v3}
        _ = v29
    default:
        switch -9223372036854775808 {
        case 0:
            v26 -= v26
        case 405:
            v3 = (v3 / (((v3 & (-(v3)))) | 1))
        default:
            v30 := GenS1{F0: 65535}
            _ = v30
        }
    }
    v31 := GenS1{F0: 434}
    if (v3 != v3) {
        var v32 uint = (v6 & 42)
        _ = v32
    }
    if false {
        v26 -= -9223372036854775808
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
