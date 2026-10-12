package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    var v4 uint64 = 4294967295
    var v5 string = "foo bar"
    v5 = (v5 + "foo bar")
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = v5
        if (!false) {
            if (false && false) {
                fmt.Println(fi6)
                v5 = v5
                var v7 int = fi6
                _ = v7
            }
            switch fi6 {
            case -336:
                for fi8 := 0; fi8 < 14; fi8++ {
                    v9 := GenS1{F0: (fi8 & fi8), F1: 429}
                    fmt.Println(fi6, v5)
                    _ = fi8
                    _ = v9
                }
            case 42:
                for fi10 := 0; fi10 < 2; fi10++ {
                    v4 -= v4
                    switch ("foo bar" + v5) {
                    case "foo bar":
                        var v11 float64 = 1.5
                        _ = v11
                    case "\n\t":
                        var v12 uint16 = 32767
                        _ = v12
                    }
                    fmt.Println(v3)
                    var v13 int64 = -9223372036854775808
                    v14 := []int{(fi10 + fi6)}
                    v5 = v5
                    fmt.Println(v4, v4)
                    _ = fi10
                    _ = v13
                    _ = v14
                }
            case 32767:
                v3 = (v3 >> ((uint(v4)) & 63))
            }
            var v15 int8 = 42
            _ = v15
        }
        _ = fi6
    }
    if true {
        fmt.Println(v5)
        if false {
            v16 := GenS2{F0: (v3 - (v3 & (v3 + (v3 % ((v3) | 1)))))}
            if ((0 > 32767) && true) {
                fmt.Println(v4, v3)
            }
            switch "a" {
            case "foo bar":
                var v17 float32 = 1e10
                _ = v17
            case "\n\t":
                v5 = (v5 + "hello")
            }
            _ = v16
        }
        var v18 int32 = -1
        _ = v18
    }
    _ = len(v2)
    v19 := GenS2{F0: uint(v4)}
    v5 = v5
    var v20 int8 = -59
    var v21 uint64 = (v4 + v4)
    var v22 int32 = 1
    var v23 bool = ((v5 + v5) != "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
