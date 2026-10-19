package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int32; F2 uint64 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 uint32 = 135
            var v7 float32 = 0.1
            _ = fi5
            _ = v6
            _ = v7
        }
        if true {
            var v8 int16 = 32767
            _ = v8
        } else {
            var v9 float32 = (0.0 * 0.1)
            _ = v9
        }
        v10 := GenS1{F0: 9223372036854775807}
        _ = fi4
        _ = v10
    }
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS2{F0: "hello", F1: -329, F2: 0}
        switch (-(fi12)) {
        case 903:
            if (false && (255 > 206)) {
                v11 -= (v11 / v11)
                var v14 string = ("\n\t" + ((("foo bar" + "hello") + "") + "hello"))
                _ = v14
            }
        }
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        switch (fi15 - (fi15 ^ (fi15 % (((fi15 >> ((640) & 63))) | 1)))) {
        case 369:
            v16 := GenS2{F0: "a", F1: 42, F2: 534}
            _ = v16
        case 349:
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int16 = 42
                v11 = float32((0.0 / 1e10))
                var v19 uint64 = 86
                _ = fi17
                _ = v18
                _ = v19
            }
        case -1:
            v11 = v11
        default:
            switch "\n\t" {
            case "a":
                var v20 uint = 18446744073709551615
                _ = v20
            case "\n\t":
                if true {
                    var v21 float32 = v11
                    _ = v21
                }
            case "":
                var v22 uint64 = 1
                _ = v22
            }
        }
        fmt.Println(v11, v11)
        _ = fi15
    }
    _ = v11
    return 88
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -9223372036854775808
    v3 := []string{""}
    v2 += -969
    if false {
        fmt.Println(v1)
    } else {
        v1 -= (^(v1))
    }
    v1 += 4294967295
    fmt.Println(v1)
    for fi23 := 0; fi23 < 16; fi23++ {
        if false {
            var v24 uint32 = 0
            _ = v24
        }
        _ = fi23
    }
    var v25 uint32 = 2147483648
    switch -855 {
    case 32767:
        var v26 bool = (0 > int8(v25))
        _ = v26
    }
    if false {
        v1 = 4294967295
        v25 = (v25 | (v25 + genFn0()))
        var v27 uint = 2
        _ = v27
    } else {
        v25 += 257
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 int8 = 42
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if (65535 == uint16(v25)) {
            v2 = (v2 ^ -9223372036854775808)
            fmt.Println(v1)
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 int8 = 42
                _ = fi31
                _ = v32
            }
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            fmt.Println(v25, v1, v1)
            var v34 int16 = int16(v25)
            _ = fi33
            _ = v34
        }
        var v35 uint8 = 29
        _ = fi30
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
