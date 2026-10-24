package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 639
    var v3 string = ("foo bar" + "")
    var v4 bool = (0 <= 952)
    if v4 {
        v3 = v3
        var v5 int8 = 42
        _ = v5
    }
    var v6 string = v3
    if v4 {
        v7 := []int{int(v2), 4294967295, 2147483647, 128}
        fmt.Println(v6, v4, v3)
        _ = v7
    } else {
        v1 = (v1 ^ 2)
        switch -272 {
        case 9223372036854775807:
            v8 := []int{2147483647, -978, -505}
            _ = v8
        case -9223372036854775808:
            v9 := []int{-1, 4294967295}
            _ = v9
        }
    }
    var v10 int64 = -247
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            switch (fi11 + (^(fi12))) {
            case 422:
                v13 := GenS1{F0: v6}
                _ = v13
            case 9223372036854775807:
                fmt.Println(v6)
            }
            _ = fi12
        }
        v14 := GenS1{F0: (("a" + v3) + v6)}
        fmt.Println(v10, v3)
        _ = fi11
        _ = v14
    }
    switch v6 {
    case "hello":
        v2 = v2
    default:
        var v15 uint16 = 805
        _ = v15
    }
    switch ("a" + v3) {
    case "foo bar":
        for fi16 := 0; fi16 < 8; fi16++ {
            if (v4 || (807 >= 2)) {
                v17 := []int{fi16, fi16, fi16}
                _ = v17
            } else {
                var v18 float32 = 1e10
                if v4 {
                    v10 = v10
                }
                _ = v18
            }
            v19 := GenS2{F0: v1, F1: int(v10)}
            _ = fi16
            _ = v19
        }
    case "a":
        v20 := []int{int(v2), 128}
        _ = v20
    }
    v21 := []int{32767, 42, -570, 9223372036854775807}
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
