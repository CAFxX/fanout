package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int16 = 1
    v2 := make(map[string]bool)
    var v3 int8 = -128
    v4 := []int{0, 4294967295, -2147483648, 85}
    if true {
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := GenS1{F0: true}
            v1 = v1
            for fi7 := 0; fi7 < 10; fi7++ {
                if (42 < 604) {
                    var v8 uint64 = 465
                    v9 := v2["foo bar"]
                    _ = v8
                    _ = v9
                }
                var v10 int8 = 4
                _ = fi7
                _ = v10
            }
            _ = fi5
            _ = v6
        }
        v3 -= v3
    }
    var v11 int = 2147483647
    v12 := GenS2{F0: 0}
    if (!true) {
        if ((0 > 0) && false) {
            for fi13 := 0; fi13 < 16; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    if (!(!true)) {
                        if true {
                            fmt.Println(fi13)
                            v15 := GenS3{F0: (v1 - v1)}
                            fmt.Println(v3)
                            _ = v15
                        }
                        var v16 int32 = 42
                        _ = v16
                    } else {
                        v17 := v2[("hello" + "a")]
                        delete(v2, "a")
                        var v18 uint32 = 2147483648
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS3{F0: v1}
                    var v20 int64 = -9223372036854775808
                    var v21 uint32 = 738
                    if (1 > 4294967295) {
                        v1 += 0
                    } else {
                        v11 = fi14
                    }
                    _ = fi14
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v1 += (v1 | -32768)
                if false {
                    var v22 float64 = 1e10
                    var v23 float32 = 1.5
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v1)
                    v1 = (^(v1))
                }
                _ = fi13
            }
        }
    } else {
        switch (^(v11)) {
        case 433:
            fmt.Println(v1, v3, v1)
        case -1:
            if len(v4) > 0 {
                v4[0] = v11
            }
        case 127:
            var v24 uint16 = 255
            _ = v24
        default:
            for fi25 := 0; fi25 < 16; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 uint16 = 1
                    _ = fi26
                    _ = v27
                }
                switch 128 {
                case 1:
                    fmt.Println(v11)
                }
                fmt.Println(v3)
                _ = fi25
            }
        }
    }
    switch "hello" {
    case "hello":
        v28 := GenS1{F0: (18446744073709551615 < 535)}
        _ = v28
    case "foo bar":
        v3 += 1
    }
    delete(v2, "\n\t")
    if false {
        fmt.Println(v3, v3, v1)
    }
    v11 -= 754
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
