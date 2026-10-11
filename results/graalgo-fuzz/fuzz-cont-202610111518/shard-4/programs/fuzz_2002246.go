package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8) int {
    var v4 uint8 = (p0 & 128)
    v4 = (^(p0))
    if true {
        for fi5 := 0; fi5 < 9; fi5++ {
            p0 = p0
            for fi6 := 0; fi6 < 11; fi6++ {
                switch (^(fi6)) {
                case -9223372036854775808:
                    p0 += 0
                case 557:
                    v7 := []int{}
                    _ = v7
                case -2147483648:
                    var v8 float64 = -2.25
                    _ = v8
                default:
                    v4 = (p0 * 0)
                }
                if (255 != uint(v4)) {
                    var v9 uint64 = 0
                    _ = v9
                } else {
                    v10 := GenS1{F0: true}
                    _ = v10
                }
                v11 := GenS2{F0: 2}
                _ = fi6
                _ = v11
            }
            fmt.Println(fi5, fi5)
            _ = fi5
        }
    }
    v12 := []int{127, -9223372036854775808, 608}
    _ = p0
    _ = v12
    _ = v4
    return 2147483647
}

func main() {
    v1 := []int{32768, 908, -2147483648}
    v2 := make(map[string]bool)
    var v3 int8 = 42
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v13 bool = false
    var v14 uint8 = 42
    v15 := GenS2{F0: 42}
    if ((v14 ^ 2) > (v14 % (((v14 ^ (v14 - (^(v14))))) | 1))) {
        var v16 int64 = -9223372036854775808
        _ = v16
    }
    fmt.Println(v13, v13)
    var v17 float32 = (1.5 + 1e10)
    for fi18 := 0; fi18 < 20; fi18++ {
        if ((v3 - (v3 >> ((0) & 63))) >= (v3 / (((v3 - v3)) | 1))) {
            v19 := GenS2{F0: 2}
            delete(v2, "\n\t")
            var v20 string = ("hello" + "hello")
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println(v13)
    v17 = (float32(3.14159) * v17)
    var v21 int8 = (v3 << ((966) & 63))
    var v22 uint64 = 1
    v23 := GenS1{F0: (!("foo bar" != "foo bar"))}
    var v24 float64 = (1.5 - -2.25)
    var v25 uint64 = (v22 << ((18446744073709551615) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
