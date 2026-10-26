package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0() int {
    var v7 uint8 = 112
    v8 := []int{}
    v7 = 224
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 2
    var v3 uint = 255
    v4 := []int{1}
    v5 := make(map[uint32]int)
    var v6 uint64 = 629
    v3 -= v3
    switch "foo bar" {
    case "foo bar":
        if ((^(v6)) == v6) {
            v9 := v5[846]
            _ = v9
        }
    }
    v6 -= (v6 % ((uint64(v3)) | 1))
    if true {
        var v10 int32 = 681
        var v11 uint32 = 2147483648
        delete(v5, (v11 | (v11 ^ 48)))
        _ = v10
        _ = v11
    } else {
        if true {
            v12 := GenS1{F0: (-1 == -128)}
            v1 += 0
            _ = v12
        } else {
            fmt.Println(v3)
        }
        if true {
            var v13 int = 660
            _ = v13
        }
    }
    if len(v4) > 0 {
        v4[0] = 433
    }
    delete(v5, 968)
    if (-563 == 0) {
        if true {
            var v14 uint8 = 0
            v15 := v5[1]
            v16 := GenS2{F0: (v6 - (v6 * (v6 % (((^(v6))) | 1))))}
            _ = v14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float32 = -2.25
            _ = fi17
            _ = v18
        }
        var v19 int32 = (^(v1))
        _ = v19
    }
    var v20 uint16 = 799
    switch (("\n\t" + "a") + ("\n\t" + "")) {
    case "hello":
        var v21 int64 = -1
        _ = v21
    case "":
        fmt.Println(v3, v1)
    default:
        v2 = v2
    }
    v3 += 730
    _ = len(v5)
    fmt.Println(v1, v2, v6)
    if true {
        v6 = uint64(v2)
        if (!(0 >= 2147483647)) {
            switch "" {
            case "foo bar":
                v20 = 255
            }
        }
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int32 = (v1 << ((v3) & 63))
            var v24 int = genFn0()
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint32 = 2
        var v26 int = -2147483648
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (false || false)}
    if false {
        switch "foo bar" {
        case "":
            v2 -= 42
        case "foo bar":
            var v28 int64 = 9223372036854775807
            _ = v28
        }
        v1 += v1
        fmt.Println(v6, v1, v6)
    }
    v3 -= v3
    var v29 uint16 = (v20 | (v20 % (((v20 | v20)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
