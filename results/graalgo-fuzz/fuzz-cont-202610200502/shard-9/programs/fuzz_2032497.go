package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint8 = 172
    var v2 uint = 794
    var v3 float64 = ((-2.25 + (1.5 + (float64((-2.25 / float32(((0.1 + 0.1) / 3.14159)))) * 0.1))) * 0.0)
    if true {
        var v4 uint32 = 657
        _ = v4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        if false {
            if true {
                for fi6 := 0; fi6 < 5; fi6++ {
                    var v7 string = "a"
                    v2 = (v2 << (((v2 >> ((v2) & 63))) & 63))
                    _ = fi6
                    _ = v7
                }
                v2 = (v2 ^ 18446744073709551615)
            } else {
                v8 := GenS1{F0: 255}
                v2 += v2
                var v9 int = fi5
                _ = v8
                _ = v9
            }
            v3 += v3
            fmt.Println(v3)
        }
        var v10 uint32 = 513
        v11 := GenS1{F0: uint16(v2)}
        _ = fi5
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 255}
    if (127 <= -117) {
        v13 := []int{1, -70, -887}
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                var v15 uint32 = 2
                var v16 uint32 = v15
                _ = v15
                _ = v16
            } else {
                v1 += (v1 - v1)
            }
            _ = fi14
        }
        _ = v13
    }
    var v17 float32 = 1.5
    v18 := []int{4294967295, -1, 32768}
    v19 := GenS1{F0: 0}
    if len(v18) > 0 {
        v18[0] = 32768
    }
    var v20 int = -9223372036854775808
    v17 = ((v17 * v17) * v17)
    v2 = v2
    switch v20 {
    case 42:
        v1 = v1
    case 128:
        v18 = append(v18, (v20 - v20))
    default:
        v3 += v3
    }
    switch "hello" {
    case "foo bar":
        switch int(v1) {
        case 878:
            if false {
                var v21 uint32 = 4294967295
                _ = v21
            }
        case -1:
            var v22 float64 = v3
            _ = v22
        case 448:
            if true {
                var v23 int16 = -32768
                _ = v23
            }
        }
    default:
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 16; fi25++ {
                v3 -= (v3 + float64(v17))
                if len(v18) > 0 {
                    v18[0] = 4294967295
                }
                if len(v18) > 0 {
                    v18[0] = -9223372036854775808
                }
                _ = fi25
            }
            var v26 uint32 = 275
            if true {
                v18 = append(v18, (fi24 - v20))
                switch (fi24 ^ 507) {
                case 32767:
                    var v27 int = fi24
                    _ = v27
                }
                var v28 uint32 = (v26 % (((v26 | (v26 - v26))) | 1))
                _ = v28
            }
            _ = fi24
            _ = v26
        }
    }
    _ = len(v18)
    var v29 int = v20
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
}
