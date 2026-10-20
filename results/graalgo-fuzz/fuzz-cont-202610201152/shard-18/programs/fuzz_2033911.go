package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() int32 {
    v5 := GenS1{F0: false, F1: 127}
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS2{F0: -54, F1: (fi6 ^ (fi6 - fi6)), F2: 18446744073709551615}
        if false {
            var v8 uint64 = 4294967295
            var v9 uint = 42
            _ = v8
            _ = v9
        }
        switch (^(fi6)) {
        case -433:
            v10 := []int{fi6, (fi6 << ((18446744073709551615) & 63)), fi6, -742}
            _ = v10
        case -2147483648:
            var v11 float32 = 1.5
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v12 := []int{0}
    var v13 float64 = -2.25
    _ = v12
    _ = v13
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint = 747
    var v2 string = "a"
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    fmt.Println(v1, v1)
    var v14 uint16 = 599
    v15 := GenS2{F0: int8(v14), F1: -262, F2: 18446744073709551615}
    if false {
        v2 = ((v2 + v2) + v2)
    } else {
        switch v2 {
        case "hello":
            v14 += v14
        case "\n\t":
            if true {
                v2 = v2
            }
        }
        _ = len(v3)
    }
    v16 := GenS1{F0: true, F1: 99}
    v2 = (("\n\t" + ((v2 + v2) + ("foo bar" + ("\n\t" + v2)))) + (("foo bar" + v2) + v2))
    var v17 uint32 = 65535
    if (!(1000 == -1)) {
        v18 := GenS1{F0: (false && (1 == 42)), F1: 0}
        _ = v18
    } else {
        v19 := []int{32768, 32768, 1, int(v14)}
        if (159 >= 2) {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
        var v20 float64 = -2.25
        _ = v19
        _ = v20
    }
    delete(v4, v17)
    switch "hello" {
    case "hello":
        v21 := GenS3{F0: uint64(v14), F1: 73}
        _ = v21
    case "foo bar":
        var v22 bool = (!((v17 >= v17) || true))
        _ = v22
    default:
        var v23 float64 = 0.0
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
