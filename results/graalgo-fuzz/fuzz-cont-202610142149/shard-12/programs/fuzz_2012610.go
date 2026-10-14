package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint32 = 65535
    var v2 int32 = 0
    v3 := make(map[int]int)
    v4 := []int{127, -1, 4294967295}
    var v5 int = 42
    fmt.Println(v1, v5, v2)
    var v6 uint32 = v1
    fmt.Println(v5, v1)
    v1 = v6
    if true {
        v7 := GenS1{F0: ((46 != 18446744073709551615) || true), F1: 0, F2: 97}
        _ = v7
    }
    v1 += v6
    if true {
        v8 := GenS1{F0: ((v1 >> ((4294967295) & 63)) >= v6), F1: v5, F2: -320}
        for fi9 := 0; fi9 < 18; fi9++ {
            if (true || true) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch "hello" {
                    case "foo bar":
                        v11 := GenS2{F0: int(v1), F1: 32767, F2: 18446744073709551615}
                        _ = v11
                    }
                    switch 127 {
                    case 127:
                        v6 += 2147483647
                    }
                    _ = fi10
                }
                v6 = (v1 - v1)
                v2 = 0
            }
            _ = fi9
        }
        _ = v8
    } else {
        v3[v5] = v5
        v1 += (v1 & v6)
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint = 498
            switch "\n\t" {
            case "\n\t":
                v1 += (v1 ^ (^(v1)))
            case "":
                switch v5 {
                case 128:
                    fmt.Println(v13, v2, v6)
                case -972:
                    fmt.Println(v13, fi12)
                default:
                    v14 := GenS2{F0: (fi12 >> ((v13) & 63)), F1: 164, F2: uint64(v1)}
                    _ = v14
                }
            case "hello":
                fmt.Println(v5, v5, v2)
            }
            v15 := GenS3{F0: -1}
            _ = fi12
            _ = v13
            _ = v15
        }
    }
    var v16 int64 = int64(v2)
    switch ("" + "foo bar") {
    case "a":
        if (!(false && true)) {
            _ = len(v3)
            v5 = (v5 / ((365) | 1))
            fmt.Println(v2, v5)
        }
    case "foo bar":
        var v17 uint64 = 815
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v6)
        _ = len(v4)
        _ = fi18
    }
    v2 -= v2
    delete(v3, v5)
    if false {
        v16 = (v16 + v16)
    }
    var v19 uint8 = 54
    if false {
        v1 = v1
        v5 = (v5 % ((127) | 1))
    } else {
        v5 = v5
        fmt.Println(v6, v16, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
