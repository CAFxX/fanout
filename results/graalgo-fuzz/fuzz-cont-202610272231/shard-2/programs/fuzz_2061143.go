package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int16 = 722
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int = 127
    v5 := GenS1{F0: 32768}
    switch -2147483648 {
    case 2147483647:
        if true {
            fmt.Println(v4, v1)
            if (v4 == v4) {
                fmt.Println(v1)
            }
        }
    case 191:
        fmt.Println(v1, v4)
    default:
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 bool = (true || (!(true && (9223372036854775807 >= 1))))
            v8 := v3["hello"]
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v4 = v4
    switch (v4 ^ v4) {
    case 1:
        if (245 < -9223372036854775808) {
            fmt.Println(v1)
        } else {
            if (2 >= 91) {
                var v9 int = v4
                _ = v9
            }
        }
    case -2147483648:
        v4 += (v4 << ((42) & 63))
    }
    switch v4 {
    case 128:
        v10 := []int{}
        _ = v10
    case -9223372036854775808:
        v1 -= v1
    case 32767:
        v4 += v4
    default:
        v1 -= (^(v1))
    }
    if false {
        v1 = (v1 - v1)
        v11 := v3[(("" + "\n\t") + "hello")]
        for fi12 := 0; fi12 < 16; fi12++ {
            v3["a"] = -476
            _ = fi12
        }
        _ = v11
    }
    fmt.Println(v4, v4)
    if (0 <= 298) {
        var v13 bool = (397 > int64(v1))
        var v14 int64 = 470
        var v15 int64 = -328
        _ = v13
        _ = v14
        _ = v15
    } else {
        v1 = (v1 + v1)
        switch (v4 >> ((2) & 63)) {
        case -2147483648:
            v16 := []int{v4, v4, (v4 & (^(v4))), (v4 | (v4 / ((v4) | 1)))}
            _ = v16
        case -146:
            for fi17 := 0; fi17 < 13; fi17++ {
                switch ("" + "foo bar") {
                case "hello":
                    v4 -= (v4 & 128)
                default:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 int = fi17
                        _ = fi18
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 20; fi20++ {
                    fmt.Println(fi17, v1)
                    v4 += fi20
                    _ = fi20
                }
                var v21 int8 = 100
                _ = fi17
                _ = v21
            }
        case -818:
            var v22 int64 = 1
            _ = v22
        default:
            switch (v4 & 127) {
            case -548:
                v23 := GenS1{F0: 1}
                _ = v23
            case 0:
                v24 := GenS1{F0: v4}
                _ = v24
            case 9223372036854775807:
                delete(v3, "a")
            }
        }
        var v25 uint = 18446744073709551615
        _ = v25
    }
    var v26 uint32 = 2147483647
    fmt.Println(v26, v26, v4)
    fmt.Println(v4)
    for fi27 := 0; fi27 < 13; fi27++ {
        switch (v4 - v4) {
        case 58:
            v4 = 2147483647
        case 2147483647:
            v28 := GenS1{F0: (v4 >> ((343) & 63))}
            _ = v28
        }
        var v29 float32 = float32((1e10 / (((1e10 - 0.1) - -2.25) - (0.1 + (-2.25 * 0.0)))))
        _ = fi27
        _ = v29
    }
    switch ("\n\t" + "a") {
    case "":
        v30 := GenS1{F0: 42}
        _ = v30
    }
    v1 += v1
    v31 := []int{0, (^(v4)), 670}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
