package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{-2147483648, -550}
        switch ("\n\t" + "") {
        case "":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "a":
            var v8 int = 42
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{42, -11, -308, 0}
    v10 := GenS1{F0: 2}
    _ = len(v9)
    _ = p0
    _ = v10
    _ = v9
    return -2147483648
}
func genFn1(p0 uint32) uint8 {
    v11 := []int{0}
    if (!((false && (1 >= 403)) || true)) {
        p0 += p0
        if (("foo bar" < ("foo bar" + "hello")) || false) {
            for fi12 := 0; fi12 < 6; fi12++ {
                p0 = p0
                var v13 uint8 = 228
                _ = fi12
                _ = v13
            }
        }
    }
    _ = p0
    _ = v11
    return 128
}

func main() {
    var v1 int64 = -356
    var v2 uint16 = 31
    var v3 uint32 = 2147483648
    var v4 int16 = 1000
    var v5 uint32 = v3
    v14 := []int{-1, 307}
    v15 := GenS2{F0: 11}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 string = ""
        _ = fi16
        _ = v17
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v5 = 4294967295
    for fi18 := 0; fi18 < 9; fi18++ {
        v2 = (^(v2))
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int64 = 1
            var v21 int8 = int8(v2)
            v22 := GenS2{F0: (v21 % (((v21 % (((v21 / ((-82) | 1))) | 1))) | 1))}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        switch "a" {
        case "hello":
            fmt.Println(fi18, v4, v4)
        }
        _ = fi18
    }
    v5 = (^(v5))
    v23 := GenS1{F0: 538}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
