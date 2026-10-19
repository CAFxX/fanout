package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0() string {
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            if (!(42 <= 32767)) {
                var v6 string = ("hello" + "foo bar")
                v6 = "hello"
                v6 = v6
                _ = v6
            }
            var v7 uint16 = 42
            _ = fi5
            _ = v7
        }
    }
    v8 := []int{243}
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -961
    var v3 uint8 = 119
    v4 := make(map[string]int)
    var v9 float32 = (0.1 - 1e10)
    v10 := []int{-1, 57, 656, 9223372036854775807}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch -9223372036854775808 {
    case 1:
        switch 42 {
        case 128:
            if (883 >= 18446744073709551615) {
                v4["hello"] = 127
            }
        default:
            v9 = (v9 / 1e10)
        }
    case 128:
        fmt.Println(v2)
    case 32767:
        var v11 uint16 = 550
        _ = v11
    }
    v12 := GenS1{F0: v3, F1: (-(v2)), F2: v2}
    if (!((^(v3)) > v3)) {
        if true {
            _ = len(v10)
            for fi13 := 0; fi13 < 17; fi13++ {
                v3 -= 237
                v14 := GenS1{F0: 23, F1: (v2 % (((v2 % ((-54) | 1))) | 1)), F2: 42}
                fmt.Println(v2)
                _ = fi13
                _ = v14
            }
            var v15 uint8 = v3
            _ = v15
        }
    }
    if false {
        v2 -= v2
        if true {
            if true {
                if (false || true) {
                    v4[(("a" + ("foo bar" + ("hello" + "a"))) + "foo bar")] = 486
                    switch "a" {
                    case "hello":
                        fmt.Println(v3, v2)
                    default:
                        v10 = append(v10, 42)
                    }
                }
            }
        }
        if false {
            fmt.Println(v9, v2)
            if false {
                var v16 uint64 = 0
                var v17 int32 = v2
                if true {
                    v17 = -1
                    var v18 int8 = 100
                    _ = v18
                } else {
                    if (!true) {
                        fmt.Println(v9)
                        v16 += v16
                    } else {
                        if false {
                            v19 := GenS2{F0: (v3 << ((0) & 63))}
                            switch "\n\t" {
                            case "foo bar":
                                var v20 int8 = 0
                                _ = v20
                            }
                            _ = v19
                        }
                        var v21 uint64 = 0
                        _ = v21
                    }
                }
                _ = v16
                _ = v17
            }
            v2 = (-(v2))
        } else {
            var v22 bool = false
            _ = v22
        }
    } else {
        fmt.Println(v2, v3)
        if true {
            if (!true) {
                v23 := GenS2{F0: (v3 ^ 1)}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        fmt.Println(v9, v9)
                        v3 = v3
                        v26 := GenS2{F0: v3}
                        _ = fi25
                        _ = v26
                    }
                    delete(v4, "hello")
                    _ = fi24
                }
                _ = len(v4)
            }
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS2{F0: (v3 + (^(v3)))}
        _ = len(v4)
        _ = fi27
        _ = v28
    }
    switch ((("a" + "foo bar") + (("" + ("hello" + "\n\t")) + "a")) + ("foo bar" + "a")) {
    case "hello":
        v2 -= 682
    }
    var v29 uint8 = 128
    for fi30 := 0; fi30 < 9; fi30++ {
        v2 += -2147483648
        _ = fi30
    }
    v29 += (v29 << ((1) & 63))
    var v31 uint32 = 0
    v31 = v31
    for fi32 := 0; fi32 < 18; fi32++ {
        if false {
            var v33 int32 = v2
            switch "hello" {
            case "\n\t":
                _ = len(v10)
            default:
                if (1 == -262) {
                    var v34 uint32 = 2147483647
                    _ = v34
                } else {
                    fmt.Println(v9, fi32)
                    v35 := GenS2{F0: (^(v3))}
                    for fi36 := 0; fi36 < 19; fi36++ {
                        v33 -= (-(v33))
                        _ = fi36
                    }
                    _ = v35
                }
            }
            _ = v33
        }
        _ = fi32
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v38 := GenS1{F0: (v29 >> ((65535) & 63)), F1: 1, F2: -1}
        v9 = v9
        _ = len(v4)
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v9
}
