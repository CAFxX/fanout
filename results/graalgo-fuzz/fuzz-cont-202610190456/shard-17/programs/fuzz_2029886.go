package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 int8, p1 string) uint8 {
    v7 := []int{128, 1, 127, 32768}
    if len(v7) > 0 {
        v7[0] = -168
    }
    fmt.Println(p1, p1)
    switch 2147483647 {
    case -2147483648:
        var v8 string = p1
        _ = v8
    case 9223372036854775807:
        var v9 float64 = 3.14159
        _ = v9
    }
    if (true || true) {
        p0 = p0
        p1 = p1
        if len(v7) > 0 {
            v7[0] = 161
        }
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                p0 -= (p0 * -126)
                switch fi10 {
                case 127:
                    var v12 int32 = 0
                    _ = v12
                case 32768:
                    switch (fi11 * (fi11 | (fi10 ^ fi10))) {
                    case 4294967295:
                        var v13 int64 = 735
                        _ = v13
                    case 32767:
                        v14 := GenS1{F0: 2147483648, F1: -1}
                        _ = v14
                    case -9223372036854775808:
                        p0 = 104
                    }
                case 4294967295:
                    var v15 int8 = p0
                    _ = v15
                }
                _ = fi11
            }
            if false {
                for fi16 := 0; fi16 < 8; fi16++ {
                    if (481 != 4294967295) {
                        if len(v7) > 0 {
                            v7[0] = fi16
                        }
                        var v17 int16 = -365
                        p0 = p0
                        _ = v17
                    } else {
                        if false {
                            v18 := GenS1{F0: 2147483648, F1: 42}
                            _ = v18
                        }
                        if (255 <= 0) {
                            p1 = "a"
                        }
                        var v19 bool = true
                        p0 = (p0 >> ((4294967295) & 63))
                        var v20 uint64 = 1
                        _ = v19
                        _ = v20
                    }
                    _ = fi16
                }
                v21 := GenS1{F0: 2, F1: 880}
                _ = v21
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := []string{(("a" + "hello") + "\n\t")}
    var v2 uint32 = 65535
    var v3 uint64 = 504
    var v4 uint16 = 0
    var v5 string = ""
    var v6 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= (v4 - (^(v4)))
    var v22 uint = 18446744073709551615
    var v23 string = v5
    v1 = append(v1, ("a" + "foo bar"))
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 float64 = 3.14159
        v25 += v25
        if true {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS1{F0: v2, F1: 9223372036854775807}
                v22 += (^(v22))
                if (!false) {
                    fmt.Println(v23)
                    v1 = append(v1, "a")
                }
                _ = fi26
                _ = v27
            }
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v6, v2, v23)
    _ = len(v1)
    _ = len(v1)
    switch -1 {
    case 93:
        _ = len(v1)
    }
    _ = len(v1)
    _ = len(v1)
    v28 := GenS1{F0: v2, F1: 0}
    var v29 uint16 = 291
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
