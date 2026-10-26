package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 4294967295
    var v3 uint = 65535
    v2 -= 42
    v4 := GenS1{F0: (v3 & 2)}
    v3 = 4294967295
    v3 += (^(v3))
    v3 += v3
    if (!((v2 % ((32767) | 1)) == 128)) {
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                v2 = (fi5 << (((v3 >> ((v3) & 63))) & 63))
                if true {
                    var v7 float64 = 1e10
                    switch (fi5 << ((v3) & 63)) {
                    case 302:
                        v8 := []int{v2, (fi6 / ((v2) | 1)), fi5}
                        _ = v8
                    case -388:
                        fmt.Println(fi6)
                    case -9223372036854775808:
                        var v9 int32 = 1
                        _ = v9
                    }
                    var v10 float64 = v7
                    _ = v10
                    _ = v7
                }
                v3 = v3
                _ = fi6
            }
            _ = fi5
        }
    } else {
        var v11 uint64 = 820
        v12 := GenS1{F0: v3}
        _ = v11
        _ = v12
    }
    var v13 int16 = 1
    if false {
        switch -312 {
        case 127:
            v14 := GenS2{F0: -128, F1: 255}
            _ = v14
        case -1:
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int = (v2 >> ((v3) & 63))
                _ = fi15
                _ = v16
            }
        }
        v17 := []int{v2, v2, (^(v2)), (v2 >> ((0) & 63))}
        switch 2147483647 {
        case 318:
            var v18 uint32 = 4294967295
            _ = v18
        }
        _ = v17
    }
    var v19 int64 = 1
    switch "\n\t" {
    case "a":
        fmt.Println(v19)
    }
    fmt.Println(v19)
    v20 := GenS3{F0: 128}
    var v21 uint8 = 42
    v13 += v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
