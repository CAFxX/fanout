package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 string = ("hello" + "hello")
    var v3 bool = (42 >= 127)
    var v4 int = -309
    var v5 uint32 = 2147483648
    v6 := []int{(v4 - v4), (v4 << ((65535) & 63)), 1}
    fmt.Println(v5, v2, v4)
    var v7 uint8 = 42
    if ((v3 && true) && v3) {
        v7 += 2
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            delete(v1, v2)
            if (!v3) {
                v9 := GenS1{F0: (v4 * fi8)}
                v10 := GenS1{F0: v4}
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi8)
                if len(v6) > 0 {
                    v6[0] = v4
                }
                _ = fi11
            }
            _ = fi8
        }
    }
    var v12 int64 = -474
    if (uint64(v7) == 42) {
        if false {
            v13 := GenS1{F0: 128}
            switch 127 {
            case -1:
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS2{F0: 940, F1: 184, F2: 1}
                    _ = fi14
                    _ = v15
                }
            case -993:
                var v16 int8 = -59
                _ = v16
            default:
                var v17 string = v2
                _ = v17
            }
            _ = v13
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint32 = 2147483648
            v12 = v12
            _ = fi18
            _ = v19
        }
        if false {
            v1["hello"] = true
            v20 := GenS2{F0: 65535, F1: (v7 / ((v7) | 1)), F2: 2147483647}
            _ = v20
        }
    } else {
        if (-1 == 2147483647) {
            var v21 int8 = 1
            fmt.Println(v7, v7)
            _ = v21
        } else {
            if (255 >= 0) {
                v22 := GenS3{F0: 210}
                _ = v22
            }
            var v23 uint8 = (^(v7))
            if len(v6) > 0 {
                v6[0] = int(v5)
            }
            _ = v23
        }
        fmt.Println(v2, v12)
    }
    v7 = (v7 * 255)
    if len(v6) > 0 {
        v6[0] = v4
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: (-(v4))}
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
