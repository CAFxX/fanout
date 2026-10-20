package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -1
    var v3 uint64 = 6
    var v4 int8 = 0
    fmt.Println(v3)
    if (0 <= 1) {
        v3 += 2
        v5 := GenS1{F0: 1000}
        _ = v5
    }
    var v6 uint16 = 421
    v2 = (v2 | (v2 * -933))
    switch "hello" {
    case "\n\t":
        v2 = 602
    case "a":
        v3 = v3
    case "":
        var v7 int64 = -9223372036854775808
        _ = v7
    default:
        var v8 float32 = (-2.25 - 1e10)
        _ = v8
    }
    v9 := GenS1{F0: -1}
    fmt.Println(v6)
    for fi10 := 0; fi10 < 19; fi10++ {
        switch (^(fi10)) {
        case -9223372036854775808:
            switch "foo bar" {
            case "foo bar":
                v4 -= (-(v4))
            }
        case 4294967295:
            if (true && (0.1 == 3.14159)) {
                var v11 int = fi10
                var v12 int = v11
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := GenS1{F0: -693}
                    if false {
                        v12 -= v11
                        v15 := []int{v12, v11, (fi13 >> ((65535) & 63)), (fi10 << ((0) & 63))}
                        switch (v12 * (fi10 * (fi13 * 128))) {
                        case -377:
                            var v16 bool = true
                            var v17 bool = v16
                            _ = v16
                            _ = v17
                        case -885:
                            if true {
                                var v18 int8 = (v4 >> ((173) & 63))
                                var v19 float64 = 3.14159
                                _ = v18
                                _ = v19
                            }
                        case -9223372036854775808:
                            fmt.Println(fi10)
                        default:
                            var v20 bool = (true && (1 < -2147483648))
                            _ = v20
                        }
                        if false {
                            v15 = append(v15, (^(fi13)))
                            var v21 bool = false
                            v22 := v1[("hello" + "hello")]
                            _ = v21
                            _ = v22
                        }
                        v4 = v4
                        _ = v15
                    } else {
                        var v23 bool = false
                        var v24 uint32 = uint32(v6)
                        var v25 int8 = (v4 * v4)
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint = 2
                    _ = fi13
                    _ = v14
                    _ = v26
                }
                _ = v11
                _ = v12
            }
        case -908:
            var v27 int8 = (v4 >> ((18446744073709551615) & 63))
            _ = v27
        default:
            v4 = v4
        }
        v4 += (v4 >> ((1) & 63))
        var v28 bool = true
        _ = fi10
        _ = v28
    }
    v29 := GenS2{F0: 2}
    var v30 uint64 = v3
    v31 := []int{0, 1, 128}
    var v32 int16 = 0
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v6
    _ = v9
}
