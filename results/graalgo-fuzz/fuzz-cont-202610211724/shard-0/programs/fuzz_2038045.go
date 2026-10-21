package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    var v3 int = 32767
    var v4 bool = (false || true)
    var v5 bool = v4
    v6 := make(map[int]int)
    fmt.Println(v4)
    var v7 int32 = -392
    v8 := GenS1{F0: 0, F1: 657, F2: 115}
    switch (("" + "") + "\n\t") {
    case "\n\t":
        v3 -= v3
    case "foo bar":
        var v9 uint8 = 108
        _ = v9
    case "":
        v5 = (!v4)
    default:
        var v10 float32 = (float32(1.5) - (-2.25 + -2.25))
        _ = v10
    }
    v11 := GenS2{F0: 1}
    v12 := GenS3{F0: -2147483648}
    var v13 int = v3
    if v4 {
        if true {
            for fi14 := 0; fi14 < 13; fi14++ {
                switch 1 {
                case -1:
                    v3 -= 2147483647
                case 42:
                    var v15 uint = 42
                    _ = v15
                case -921:
                    if v4 {
                        var v16 float64 = float64(1e10)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v3 = (fi14 << ((42) & 63))
                            v1 = v7
                            fmt.Println(v4)
                            _ = fi17
                        }
                    }
                default:
                    var v18 float64 = 1e10
                    _ = v18
                }
                _ = fi14
            }
            v7 -= 518
        }
        v19 := GenS2{F0: 4294967295}
        _ = v19
    } else {
        v20 := GenS3{F0: v3}
        _ = v20
    }
    var v21 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
