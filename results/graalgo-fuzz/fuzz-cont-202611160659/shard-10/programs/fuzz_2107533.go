package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []int{42, -2147483648, -1}
    var v2 uint16 = 562
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    if (4294967295 <= 4294967295) {
        if true {
            v5 := GenS1{F0: 2}
            var v6 int = -9223372036854775808
            _ = len(v1)
            _ = v5
            _ = v6
        }
        var v7 uint8 = 128
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = -53
            switch "foo bar" {
            case "\n\t":
                for fi10 := 0; fi10 < 19; fi10++ {
                    v2 = 834
                    var v11 uint8 = v7
                    var v12 int16 = 42
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                var v13 int = (fi8 + 9223372036854775807)
                _ = v13
            }
            _ = len(v3)
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        v14 := v3[834]
        if false {
            delete(v4, 897)
            var v15 int16 = 706
            _ = v15
        } else {
            fmt.Println(v14, v2)
        }
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        _ = v14
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(fi17, fi17)
                _ = fi17
            }
        }
        _ = fi16
    }
    if true {
        if (18446744073709551615 != 1) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 = (v2 | 449)
        }
    }
    fmt.Println(v2)
    var v18 int32 = -245
    v18 = v18
    if (18446744073709551615 <= 2) {
        var v19 float64 = ((((1e10 + float64((float32(0.0) - 3.14159))) / ((0.1 + 0.0) + 1.5)) - 1e10) * 3.14159)
        _ = v19
    }
    if (0 != -107) {
        v20 := GenS1{F0: 2}
        _ = v20
    } else {
        var v21 uint8 = 0
        v21 = (v21 + v21)
        v2 = (v2 & v2)
        _ = v21
    }
    delete(v4, 175)
    v22 := GenS1{F0: 18446744073709551615}
    var v23 int64 = -1
    if len(v1) > 0 {
        v1[0] = 691
    }
    if (v2 <= v2) {
        fmt.Println(v2, v23, v2)
        for fi24 := 0; fi24 < 14; fi24++ {
            v18 = v18
            for fi25 := 0; fi25 < 8; fi25++ {
                v2 = v2
                v2 = v2
                switch "hello" {
                case "a":
                    v26 := GenS1{F0: 536}
                    _ = v26
                case "":
                    delete(v4, 42)
                case "foo bar":
                    fmt.Println(fi24, v2, v23)
                    var v27 uint32 = 4294967295
                    var v28 string = "foo bar"
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            if false {
                v2 = v2
                var v29 float64 = 1.5
                _ = v29
            } else {
                v23 += 814
                if true {
                    var v30 int8 = 42
                    v23 = v23
                    _ = v30
                }
            }
            _ = fi24
        }
        var v31 uint64 = 18446744073709551615
        _ = v31
    } else {
        v18 = -307
        if (true || (4294967295 >= 1)) {
            v32 := v3[4294967295]
            fmt.Println(v23, v2, v23)
            v2 += 32767
            _ = v32
        }
    }
    fmt.Println(v2, v23, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
