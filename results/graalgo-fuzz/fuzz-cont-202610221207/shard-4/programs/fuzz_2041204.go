package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint, p2 int32) uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            p0 = (p0 ^ (p0 ^ p0))
        }
        if false {
            v8 := []int{}
            _ = v8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 42}
    var v10 bool = true
    if v10 {
        var v11 float32 = 1e10
        if v10 {
            fmt.Println(v11, p0, p2)
            var v12 uint64 = 299
            v13 := GenS1{F0: 2}
            _ = v12
            _ = v13
        } else {
            switch ("" + ("hello" + "hello")) {
            case "a":
                p1 -= (p1 * p1)
            }
        }
        _ = v11
    } else {
        var v14 int = -367
        v15 := []int{947, v14, (-(v14)), 9223372036854775807}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 174
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint = 889
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v6 uint32 = 2147483647
    fmt.Println(v3)
    _ = len(v4)
    v3 -= v3
    v2[("foo bar" + "hello")] = ""
    fmt.Println(v6, v6)
    v3 = (v3 >> (((v3 + 255)) & 63))
    v16 := GenS1{F0: 42}
    v17 := []int{1, 1}
    _ = len(v17)
    switch -9223372036854775808 {
    case -295:
        switch 42 {
        case -282:
            var v18 float64 = 1e10
            _ = v18
        case 127:
            v19 := GenS1{F0: uint8(v6)}
            _ = v19
        case -1:
            if ((("hello" + "a") < "foo bar") && false) {
                v20 := v5[(v6 >> ((627) & 63))]
                switch "a" {
                case "a":
                    v6 = (v6 << ((4294967295) & 63))
                case "hello":
                    fmt.Println(v3, v3, v20)
                default:
                    if true {
                        switch "a" {
                        case "a":
                            v20 += v20
                        case "foo bar":
                            var v21 uint = (v3 - 4294967295)
                            _ = v21
                        }
                        fmt.Println(v20, v3, v6)
                        v6 -= v6
                    } else {
                        if (!true) {
                            switch ("foo bar" + (("\n\t" + ("hello" + "hello")) + "foo bar")) {
                            case "hello":
                                var v22 uint8 = 255
                                _ = v22
                            }
                            v23 := GenS1{F0: 0}
                            v24 := GenS1{F0: 187}
                            _ = v23
                            _ = v24
                        } else {
                            switch v20 {
                            case 4294967295:
                                v3 -= v3
                            default:
                                var v25 int16 = 32767
                                _ = v25
                            }
                        }
                        v26 := GenS1{F0: 1}
                        _ = v26
                    }
                }
                var v27 uint64 = 698
                _ = v20
                _ = v27
            }
        default:
            if (-2147483648 >= 0) {
                if true {
                    fmt.Println(v3)
                    var v28 int16 = 1
                    _ = v28
                } else {
                    var v29 string = ("\n\t" + "")
                    v3 = (v3 - v3)
                    fmt.Println(v29, v29)
                    _ = v29
                }
                v3 = (v3 % (((v3 >> (((v3 - (^(v3)))) & 63))) | 1))
                v6 = v6
            }
        }
    default:
        var v30 int = 42
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
