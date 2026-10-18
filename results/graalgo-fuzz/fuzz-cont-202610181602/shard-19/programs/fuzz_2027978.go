package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    var v6 int = 127
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = -280
        _ = fi7
        _ = v8
    }
    v6 = v6
    var v9 string = "hello"
    _ = v6
    _ = v9
    return 2
}

func main() {
    var v1 int = -1
    var v2 int32 = 0
    var v3 int16 = 32767
    var v4 uint32 = 65535
    v5 := []int{(v1 * (^(v1)))}
    v2 += (v2 ^ (v2 & 885))
    v5 = append(v5, v1)
    v10 := GenS1{F0: true, F1: 9223372036854775807, F2: 913}
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi11
    }
    v12 := GenS2{F0: uint(v4), F1: 822}
    var v13 uint = 840
    v13 += 4294967295
    if false {
        if true {
            v14 := GenS3{F0: v2}
            _ = v14
        }
        v4 += (^(v4))
        var v15 float64 = (1e10 - 3.14159)
        _ = v15
    } else {
        switch (^(v1)) {
        case 42:
            switch "" {
            case "\n\t":
                v16 := GenS2{F0: v13, F1: 2}
                _ = v16
            default:
                if false {
                    fmt.Println(v13, v3)
                }
            }
        case -409:
            fmt.Println(v1, v4)
        case 1:
            v4 += (v4 - 0)
        }
    }
    v4 += 2
    var v17 int16 = (v3 / (((v3 >> (((^(v13))) & 63))) | 1))
    switch (-(v1)) {
    case -9223372036854775808:
        var v18 uint8 = 128
        _ = v18
    case 690:
        var v19 float32 = 0.0
        _ = v19
    case 127:
        if false {
            if len(v5) > 0 {
                v5[0] = v1
            }
            var v20 int8 = -128
            for fi21 := 0; fi21 < 5; fi21++ {
                if ("hello" >= "foo bar") {
                    v20 = (v20 - int8(v2))
                }
                _ = fi21
            }
            _ = v20
        }
    }
    var v22 int32 = (^(v2))
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 = v17
            _ = fi24
        }
        var v25 int16 = (v17 + (v3 >> (((v13 % ((4294967295) | 1))) & 63)))
        _ = fi23
        _ = v25
    }
    var v26 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
