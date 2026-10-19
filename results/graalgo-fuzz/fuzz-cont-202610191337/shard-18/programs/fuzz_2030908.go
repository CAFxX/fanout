package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 int; F2 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 int8 = 42
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int32 = -1
        v6 := []int{}
        v3 -= (-(v3))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: int64(v1), F1: ("a" + "\n\t"), F2: 65535}
        var v9 int8 = v3
        if true {
            fmt.Println(fi7, fi7)
            v10 := GenS1{F0: -148, F1: "", F2: 65535}
            var v11 uint16 = 74
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v1, v3, v1)
        switch "foo bar" {
        case "hello":
            var v13 uint16 = 117
            _ = v13
        case "":
            if (1000 > 0) {
                v14 := v2[fi12]
                v15 := GenS1{F0: 536, F1: "", F2: 4294967295}
                _ = v14
                _ = v15
            }
        case "a":
            fmt.Println(v1)
        }
        _ = fi12
    }
    v16 := GenS2{F0: 9223372036854775807, F1: 1, F2: v3}
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := []int{}
        for fi19 := 0; fi19 < 13; fi19++ {
            if false {
                var v20 string = "\n\t"
                var v21 uint = 18446744073709551615
                _ = v20
                _ = v21
            } else {
                v3 = v3
            }
            v1 -= (^(v1))
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    var v22 int16 = -1
    v3 -= (v3 & v3)
    v3 = 1
    v23 := GenS2{F0: -9223372036854775808, F1: 32767, F2: v3}
    if (!(true && (!(uint(v1) > 2)))) {
        var v24 int16 = v22
        _ = v24
    }
    _ = len(v2)
    var v25 int64 = 943
    var v26 uint16 = 213
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
}
