package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int32 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    var v8 int64 = 0
    v9 := []int{-552, 1, 0}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    switch ((("hello" + "a") + "foo bar") + ("" + ("" + "hello"))) {
    case "hello":
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v8 -= (^(v8))
        case "":
            if len(v9) > 0 {
                v9[0] = 4294967295
            }
        default:
            var v10 uint = 1
            _ = v10
        }
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 35
    v2 := make(map[uint32]bool)
    var v3 string = "hello"
    v4 := []int{32768}
    v5 := make(map[string]uint32)
    var v11 int32 = -2147483648
    v3 = v3
    var v12 int64 = -9223372036854775808
    var v13 string = (((v3 + ("\n\t" + v3)) + "") + v3)
    var v14 int = 32768
    switch int(v11) {
    case 1:
        switch v14 {
        case 128:
            var v15 uint = 225
            _ = v15
        default:
            var v16 int64 = 1
            _ = v16
        }
    case 32767:
        if (false || false) {
            if false {
                if true {
                    fmt.Println(v14, v14)
                    var v17 uint64 = 940
                    v17 = uint64(v1)
                    _ = v17
                }
                v18 := GenS1{F0: 42, F1: 0, F2: v12}
                _ = v18
            }
            v12 = 0
            if false {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int = (v14 - v14)
                    _ = fi19
                    _ = v20
                }
                v14 = v14
            }
        }
    case 127:
        switch (v13 + v13) {
        case "foo bar":
            var v21 uint64 = 0
            _ = v21
        case "hello":
            v22 := GenS1{F0: 127, F1: 136, F2: v12}
            _ = v22
        case "a":
            for fi23 := 0; fi23 < 13; fi23++ {
                v1 += (v1 * 0)
                v11 = v11
                _ = fi23
            }
        }
    }
    var v24 uint32 = 0
    if false {
        if (1.5 != float32(1e10)) {
            _ = len(v2)
        }
    } else {
        fmt.Println(v24, v1)
        v25 := GenS2{F0: 1}
        _ = v25
    }
    var v26 bool = true
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 uint32 = (v24 * (v24 | v24))
        for fi29 := 0; fi29 < 7; fi29++ {
            v11 -= v11
            fmt.Println(v11, v24)
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
