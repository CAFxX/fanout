package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint64 = 1
    v2 := []int{288, -430, 4294967295, -683}
    var v3 uint8 = 127
    var v4 int16 = -1
    v5 := []int{127, 4294967295}
    v6 := make(map[string]bool)
    var v7 uint = 65535
    var v8 bool = true
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println(v1)
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "foo bar" {
        case "\n\t":
            v10 := GenS1{F0: 65535}
            _ = v10
        }
        if ("\n\t" >= ("hello" + "foo bar")) {
            v11 := GenS1{F0: 2147483647}
            v12 := GenS1{F0: 1}
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: v1, F1: fi9, F2: 85}
            v14 := GenS1{F0: 82}
            for fi15 := 0; fi15 < 20; fi15++ {
                if (!true) {
                    switch "hello" {
                    case "foo bar":
                        var v16 int16 = v4
                        _ = v16
                    case "":
                        v6["foo bar"] = false
                    case "a":
                        fmt.Println(fi15, fi15)
                    }
                    v1 -= 0
                    switch (-(fi9)) {
                    case -1:
                        if (1 < 0) {
                            var v17 uint = 65535
                            v5 = append(v5, fi15)
                            _ = v17
                        }
                    default:
                        if v8 {
                            v7 = v7
                            var v18 uint64 = 2
                            var v19 uint = 30
                            _ = v18
                            _ = v19
                        }
                    }
                }
                v20 := GenS2{F0: v1, F1: 383, F2: 42}
                _ = fi15
                _ = v20
            }
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    v7 -= (v7 * 0)
    if ((v3 * (v3 & uint8(v7))) < uint8(v1)) {
        if true {
            _ = len(v6)
        }
        if ((^(v7)) >= (v7 << ((v7) & 63))) {
            var v21 float32 = ((0.0 / (0.1 + 0.1)) / 0.1)
            v1 -= 4294967295
            _ = v21
        }
        v4 = 42
    } else {
        v22 := GenS1{F0: 2}
        var v23 uint64 = v1
        _ = v22
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 17; fi24++ {
            if v8 {
                v8 = v8
            }
            v4 -= -728
            _ = fi24
        }
        v4 += (v4 - -32768)
        v25 := GenS2{F0: v1, F1: -530, F2: 255}
        _ = v25
    } else {
        v26 := GenS3{F0: (v3 / ((42) | 1))}
        _ = v26
    }
    v7 = 344
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
