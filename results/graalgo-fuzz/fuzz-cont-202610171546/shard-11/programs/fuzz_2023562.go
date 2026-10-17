package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 679
    var v5 int16 = -32768
    v6 := []int{709}
    switch (("hello" + "hello") + ("\n\t" + (("a" + "a") + ""))) {
    case "foo bar":
        fmt.Println(v5)
    default:
        v1 = 1e10
    }
    var v7 int16 = (v5 >> ((42) & 63))
    var v8 int32 = 1
    var v9 int32 = 1
    var v10 int = 5
    if (((1.5 * 0.1) - -2.25) < 1e10) {
        fmt.Println(v8)
    } else {
        _ = len(v3)
        fmt.Println(v5, v10)
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                v10 = (v10 | (fi11 ^ v10))
                v5 = (-(v7))
                delete(v3, 42)
                if true {
                    if (1e10 <= (v1 + 0.1)) {
                        var v13 int8 = 0
                        var v14 uint32 = 0
                        v1 = 1e10
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(fi12)
                    if (!true) {
                        switch fi12 {
                        case 950:
                            fmt.Println(v10, v5)
                        default:
                            v15 := GenS1{F0: -1, F1: 37}
                            _ = v15
                        }
                        var v16 int16 = (^(v5))
                        var v17 string = ("foo bar" + "a")
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v16
                        _ = v17
                    } else {
                        v5 -= (v7 | (v7 + (v5 * -1)))
                        if (fi11 < (v10 << ((255) & 63))) {
                            var v18 int32 = v8
                            _ = v18
                        } else {
                            fmt.Println(v8, fi11, fi11)
                        }
                        v4 = v4
                        delete(v2, "")
                    }
                    v19 := GenS2{F0: 96, F1: 65535}
                    var v20 uint8 = 127
                    fmt.Println(v7, v7)
                    _ = v19
                    _ = v20
                }
                _ = fi12
            }
            var v21 float64 = (v1 / 0.1)
            _ = fi11
            _ = v21
        }
    }
    var v22 int32 = 473
    v23 := GenS3{F0: (v9 >> ((65535) & 63)), F1: 99}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
