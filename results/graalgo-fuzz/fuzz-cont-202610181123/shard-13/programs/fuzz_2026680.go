package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0() int {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 42, F1: -128, F2: ((("hello" + "a") + ((("a" + "\n\t") + "a") + "hello")) + (("" + "hello") + ("foo bar" + "")))}
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10)
        v11 := []int{4294967295}
        var v12 int16 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = 206
    _ = v13
    return (v13 * (v13 ^ 32767))
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 42, 681}
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 uint8 = 42
    var v5 int = 9223372036854775807
    v6 := make(map[uint32]bool)
    var v14 uint32 = 0
    var v15 uint64 = 302
    v4 = (v4 ^ 255)
    fmt.Println(v15, v4, v14)
    v5 -= (v5 + (^(v5)))
    if false {
        v16 := GenS1{F0: 2, F1: -1, F2: ""}
        _ = v16
    } else {
        v15 -= (^(v15))
    }
    _ = len(v1)
    for fi17 := 0; fi17 < 13; fi17++ {
        if ((!(false || ((v14 != (v14 + 686)) && true))) || ((true || true) && false)) {
            var v18 int32 = -91
            v19 := GenS1{F0: 248, F1: 100, F2: ("hello" + "")}
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: (^(v15)), F1: false, F2: 1}
        v4 += v4
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v5, v5)
        v14 = 2147483647
        _ = fi21
    }
    var v22 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
