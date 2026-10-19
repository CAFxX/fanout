package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0() int {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 42, F1: 18446744073709551615}
        var v9 int16 = -32768
        var v10 uint = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 string = ""
    if true {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{918, 42, -750, -1}
            v11 = ""
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 44
        _ = v14
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := []int{-9223372036854775808, (fi16 % ((9223372036854775807) | 1)), -2147483648}
                var v18 int8 = 1
                v17 = append(v17, fi15)
                _ = fi16
                _ = v17
                _ = v18
            }
            fmt.Println(v11, v11)
            _ = fi15
        }
        v19 := GenS1{F0: 42, F1: 36}
        _ = v19
    }
    var v20 int16 = 931
    if false {
        v20 = (v20 / ((1000) | 1))
        v11 = v11
    } else {
        v20 = v20
        v21 := []int{}
        _ = v21
    }
    _ = v11
    _ = v20
    return 4294967295
}

func main() {
    v1 := []string{("foo bar" + "a")}
    v2 := []int{1, -1, 767}
    var v3 uint16 = 42
    var v4 uint8 = 128
    var v5 int32 = 42
    v6 := make(map[uint32]string)
    v3 = (v3 / (((v3 + v3)) | 1))
    var v22 uint32 = 1
    var v23 uint8 = v4
    fmt.Println(v22, v3, v23)
    if (v22 == v22) {
        v24 := GenS1{F0: (^(v22)), F1: 206}
        _ = v24
    }
    var v25 uint64 = 201
    v4 = (^(v23))
    fmt.Println(v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
