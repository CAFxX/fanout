package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string }
func genFn0() int32 {
    var v4 bool = true
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        fmt.Println(v4)
        _ = fi5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 0, F1: 817, F2: 936}
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 0, F1: 225, F2: 2}
        v4 = v4
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint64 = 131
            fmt.Println(v11, fi8, fi10)
            if v4 {
                fmt.Println(fi8, v11)
                if true {
                    v12 := []int{(^(fi8)), (fi10 / ((fi10) | 1))}
                    var v13 int64 = -9223372036854775808
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: 42, F1: 1, F2: 931}
                v11 += 1
                fmt.Println(v11, fi10)
                _ = v14
            } else {
                v15 := GenS1{F0: 1000, F1: 255, F2: 42}
                var v16 string = "foo bar"
                _ = v15
                _ = v16
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = v4
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 string = "\n\t"
    var v17 int8 = -73
    v17 -= v17
    v17 = v17
    v3 = (v3 + ((v3 + "foo bar") + "foo bar"))
    v18 := GenS2{F0: 18446744073709551615, F1: -9223372036854775808, F2: 1}
    v19 := GenS2{F0: 4294967295, F1: -302, F2: -1}
    v3 = v3
    var v20 int64 = 1
    v21 := GenS3{F0: v3, F1: 65535}
    var v22 uint = 309
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int64 = (v20 - (v20 >> ((v22) & 63)))
        var v25 int32 = int32(v20)
        _ = fi23
        _ = v24
        _ = v25
    }
    if false {
        v26 := GenS3{F0: (("a" + "") + ""), F1: 65535}
        fmt.Println(v3, v17)
        v27 := GenS4{F0: v3}
        _ = v26
        _ = v27
    }
    delete(v1, (v3 + v3))
    var v28 int8 = -33
    v17 -= (v17 ^ (v28 << (((v22 ^ 427)) & 63)))
    v17 -= v17
    var v29 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
}
