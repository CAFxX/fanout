package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
func genFn0(p0 string, p1 uint64, p2 bool) int64 {
    v7 := []int{175, 1, -2147483648, 4294967295}
    var v8 int32 = 183
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 0
    var v3 int32 = 2147483647
    var v4 uint8 = 128
    var v5 uint16 = 42
    v6 := []string{"hello", (("\n\t" + "hello") + "foo bar")}
    var v9 int8 = -1
    var v10 uint64 = v2
    var v11 uint64 = 1
    if (42 != int16(v3)) {
        var v12 int = 42
        _ = v12
    }
    v1[128] = 128
    switch "" {
    case "":
        var v13 uint8 = (v4 / (((v4 / (((v4 * (^(v4)))) | 1))) | 1))
        _ = v13
    case "hello":
        delete(v1, 42)
    case "a":
        var v14 uint8 = (v4 >> ((2) & 63))
        _ = v14
    default:
        var v15 uint64 = v10
        _ = v15
    }
    v9 = -128
    var v16 string = (((("\n\t" + "\n\t") + ("hello" + ("a" + "a"))) + "a") + "hello")
    var v17 int32 = -295
    switch v16 {
    case "foo bar":
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint32 = 0
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS1{F0: "foo bar", F1: ("" + (v16 + "hello")), F2: v5}
                v22 := GenS1{F0: v16, F1: (("a" + (v16 + ((("\n\t" + v16) + (v16 + (((((("hello" + v16) + v16) + v16) + v16) + v16) + v16))) + (((v16 + v16) + v16) + ("" + v16))))) + (v16 + v16)), F2: 1}
                if true {
                    v23 := GenS1{F0: v16, F1: ("" + (v16 + v16)), F2: 42}
                    var v24 int8 = v9
                    _ = v23
                    _ = v24
                } else {
                    if true {
                        if true {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    fmt.Println(v11)
                                } else {
                                    var v25 int16 = 681
                                    v9 += -88
                                    _ = v25
                                }
                            default:
                                _ = len(v1)
                            }
                            v1[fi20] = (fi18 << ((2) & 63))
                        } else {
                            var v26 int32 = (^(v3))
                            v27 := GenS2{F0: fi18, F1: v26, F2: v4}
                            _ = v26
                            _ = v27
                        }
                        var v28 string = v16
                        _ = v28
                    } else {
                        var v29 uint32 = (^(v19))
                        _ = v29
                    }
                }
                v10 += (^(v2))
                _ = fi20
                _ = v21
                _ = v22
            }
            switch (("hello" + ("foo bar" + "\n\t")) + "") {
            case "hello":
                var v30 uint = 188
                _ = v30
            default:
                _ = len(v1)
            }
            _ = fi18
            _ = v19
        }
    case "":
        for fi31 := 0; fi31 < 13; fi31++ {
            if ((true && (((v9 - int8(v2)) > v9) && true)) && false) {
                v2 -= (v11 * (v2 / ((v11) | 1)))
                delete(v1, fi31)
                for fi32 := 0; fi32 < 5; fi32++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi32
                }
            } else {
                v33 := v1[fi31]
                var v34 uint16 = v5
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    var v35 uint8 = (v4 / ((127) | 1))
    v4 -= 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
