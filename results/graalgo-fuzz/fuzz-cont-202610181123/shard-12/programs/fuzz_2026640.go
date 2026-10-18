package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 string, p1 int32, p2 uint8) uint32 {
    p0 = ("a" + "hello")
    fmt.Println(p1, p1)
    if (165 == (^(p2))) {
        fmt.Println(p1, p1)
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = p0
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}
func genFn1(p0 bool) int64 {
    for fi8 := 0; fi8 < 4; fi8++ {
        if (-1 != -32768) {
            var v9 uint = 1
            switch -2147483648 {
            case -731:
                switch -606 {
                case 128:
                    v9 = (v9 + v9)
                case 32768:
                    v9 = v9
                default:
                    var v10 uint8 = 1
                    _ = v10
                }
            case 32768:
                v11 := []int{}
                _ = v11
            default:
                if ((p0 || p0) || (!((p0 || false) && true))) {
                    v12 := GenS1{F0: 100, F1: 2147483648}
                    var v13 int32 = 1
                    _ = v12
                    _ = v13
                } else {
                    v9 = v9
                    v14 := GenS1{F0: 10, F1: 2147483648}
                    _ = v14
                }
            }
            _ = v9
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            p0 = (100 < 8)
            _ = fi15
        }
        switch -2147483648 {
        case 522:
            fmt.Println(p0, p0)
        case -533:
            for fi16 := 0; fi16 < 20; fi16++ {
                switch 9223372036854775807 {
                case -141:
                    v17 := []int{fi8, fi16}
                    _ = v17
                case -1:
                    if (4294967295 < 169) {
                        p0 = p0
                        var v18 uint64 = 357
                        fmt.Println(p0, v18)
                        _ = v18
                    }
                case 2147483647:
                    switch -9223372036854775808 {
                    case -932:
                        if (2 >= 18446744073709551615) {
                            var v19 int64 = -1
                            _ = v19
                        } else {
                            switch "foo bar" {
                            case "a":
                                var v20 bool = (0 < 42)
                                _ = v20
                            default:
                                p0 = (true && p0)
                            }
                            fmt.Println(fi16, fi8)
                        }
                    }
                default:
                    var v21 int32 = 2147483647
                    _ = v21
                }
                _ = fi16
            }
        default:
            p0 = (p0 || true)
        }
        _ = fi8
    }
    switch (("hello" + "a") + "foo bar") {
    case "a":
        var v22 int16 = 32767
        _ = v22
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    var v4 float32 = (((-2.25 + v3) - 0.1) * v3)
    var v5 uint8 = 128
    var v6 string = ""
    v5 = (v5 ^ v5)
    v23 := GenS1{F0: -1, F1: 65535}
    var v24 float32 = ((float32(1e10) - (1e10 - -2.25)) - v4)
    v25 := GenS1{F0: -1, F1: genFn0("hello", 846, v5)}
    _ = len(v2)
    v26 := []int{}
    if (0 <= 2147483647) {
        v3 = (v3 / v24)
        for fi27 := 0; fi27 < 1; fi27++ {
            v6 = v6
            _ = fi27
        }
    }
    v2[-938] = 118
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
