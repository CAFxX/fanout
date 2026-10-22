package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 10
    v3 := []uint32{}
    var v4 int8 = v2
    v5 := make(map[int]bool)
    v6 := make(map[int]bool)
    var v7 uint16 = 0
    var v8 int16 = 42
    var v9 string = ("hello" + "a")
    if true {
        fmt.Println(v9)
        v9 = v9
        var v10 int = 2147483647
        _ = v10
    }
    if (0.1 == ((0.1 - 1e10) * ((1e10 + (0.1 / ((1e10 - 0.0) * ((-2.25 - 1e10) * (1e10 / 1.0))))) * ((0.0 - 0.1) * float32(float64(1.5)))))) {
        for fi11 := 0; fi11 < 15; fi11++ {
            switch (fi11 + fi11) {
            case 2147483647:
                v8 = -1
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint64 = 654
                var v14 int32 = -2147483648
                if true {
                    var v15 int64 = -9223372036854775808
                    switch fi12 {
                    case 160:
                        fmt.Println(fi12, v4)
                    case 861:
                        var v16 uint32 = 2147483648
                        _ = v16
                    }
                    switch (-(fi12)) {
                    case 128:
                        var v17 int = 32768
                        _ = v17
                    }
                    _ = v15
                }
                fmt.Println(v8, v14, v1)
                v13 -= v13
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        var v18 uint64 = 1
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint64 = v18
            fmt.Println(fi19, v2, v8)
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v21 int64 = 9223372036854775807
    v22 := GenS1{F0: 300, F1: 32767}
    var v23 uint = (^(v1))
    v24 := GenS2{F0: 455, F1: 42}
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(v7, v23, v9)
        fmt.Println(v2)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
