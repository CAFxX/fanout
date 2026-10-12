package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
func genFn0() uint8 {
    v6 := []int{-2147483648}
    var v7 int32 = 285
    v8 := GenS1{F0: 128, F1: "hello"}
    if (-93 > 42) {
        fmt.Println(v7)
    } else {
        var v9 uint64 = 67
        v10 := GenS1{F0: 128, F1: ""}
        v11 := GenS1{F0: 149, F1: "hello"}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 231
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 82
    var v3 int8 = 100
    var v4 int = 600
    var v5 uint64 = v1
    v3 -= v3
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: 1, F1: ""}
        v5 = (v1 >> (((v2 & 2)) & 63))
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: ((("\n\t" + "\n\t") + "a") + "foo bar")}
    v1 += (v1 - 2)
    var v15 string = ("hello" + "\n\t")
    v4 = (v4 | v4)
    v2 = v2
    v4 += (v4 % (((v4 & v4)) | 1))
    var v16 string = "\n\t"
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := []int{}
        v19 := GenS1{F0: 238, F1: v15}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v1)
    if (!false) {
        fmt.Println(v16, v5, v2)
        var v20 float64 = (-2.25 + float64(1.5))
        if (int8(v2) > v3) {
            v1 += 42
            var v21 uint32 = 2
            _ = v21
        }
        _ = v20
    }
    v3 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
