package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }

func main() {
    v1 := []string{}
    var v2 bool = true
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 uint16 = 65535
    var v6 int = 135
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if (((v5 & (v5 << ((121) & 63))) == (v5 - (v5 << ((2) & 63)))) || (!(!false))) {
        v7 := GenS1{F0: "\n\t", F1: 1, F2: 213}
        if (100 == 42) {
            v8 := GenS1{F0: "\n\t", F1: 100, F2: 18446744073709551615}
            _ = v8
        }
        if (true && false) {
            v5 += v5
            for fi9 := 0; fi9 < 7; fi9++ {
                v2 = (-191 != 0)
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch 32768 {
                    case -504:
                        v5 = v5
                    }
                    v2 = false
                    _ = fi10
                }
                _ = len(v1)
                _ = fi9
            }
            v6 = v6
        }
        _ = v7
    }
    fmt.Println(v6, v6, v6)
    v2 = v2
    v2 = v2
    var v11 int16 = -32768
    v2 = v2
    var v12 uint32 = 2147483647
    v13 := GenS1{F0: "foo bar", F1: -1, F2: 42}
    v12 += (v12 | v12)
    v14 := GenS1{F0: "\n\t", F1: 127, F2: 244}
    switch v6 {
    case 127:
        var v15 int = -9223372036854775808
        _ = v15
    case 32767:
        v2 = (!(2147483647 <= 32768))
    }
    v16 := GenS1{F0: "", F1: 1, F2: 4294967295}
    var v17 string = "foo bar"
    v18 := GenS2{F0: (!v2), F1: 952, F2: (-(v11))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
