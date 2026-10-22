package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int, p2 int64) string {
    if true {
        p0 = (^(p0))
    } else {
        if (461 > 0) {
            fmt.Println(p2, p2, p1)
            if true {
                var v4 uint32 = 4294967295
                var v5 uint16 = 65535
                p0 = (p0 + 855)
                _ = v4
                _ = v5
            } else {
                v6 := []int{(p1 >> ((4294967295) & 63))}
                p1 = p1
                _ = v6
            }
        }
    }
    p1 = (p1 + (p1 - 32767))
    v7 := GenS1{F0: p0}
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("a" + "")
}
func genFn1() uint {
    switch -842 {
    case 128:
        var v10 float32 = 1e10
        _ = v10
    case 9223372036854775807:
        v11 := []int{-168, 127, 9223372036854775807}
        _ = v11
    case 206:
        var v12 uint = 4294967295
        _ = v12
    }
    var v13 uint16 = 32767
    var v14 int64 = 379
    var v15 int64 = -631
    fmt.Println(v15, v13)
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := []int{9223372036854775807, 537}
    var v2 int16 = 1000
    var v3 float64 = (float64(3.14159) - 3.14159)
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 = v3
        _ = fi16
    }
    switch "\n\t" {
    case "\n\t":
        var v17 string = ("foo bar" + "a")
        _ = v17
    case "":
        if false {
            var v18 int = 1
            switch (^(v18)) {
            case 717:
                if true {
                    if len(v1) > 0 {
                        v1[0] = v18
                    }
                } else {
                    v18 -= 109
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v18
        }
    }
    var v19 int8 = -1
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v20 uint64 = 636
    v19 -= -23
    v20 = (^(v20))
    if (811 < 508) {
        var v21 int16 = 424
        v19 -= v19
        _ = v21
    } else {
        v22 := GenS1{F0: 42}
        _ = v22
    }
    var v23 uint8 = 182
    v19 = v19
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint8 = v23
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi24
        _ = v25
    }
    var v26 int16 = (v2 >> ((319) & 63))
    v2 = v2
    if true {
        fmt.Println(v23, v2)
        var v27 uint32 = 165
        if len(v1) > 0 {
            v1[0] = -924
        }
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        if false {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 int8 = v19
                v20 = v20
                _ = fi29
                _ = v30
            }
        }
        v3 = v3
        _ = fi28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v3
}
