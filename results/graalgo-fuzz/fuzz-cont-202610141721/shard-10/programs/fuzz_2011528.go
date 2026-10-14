package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 1
    var v6 float32 = 0.1
    var v7 int32 = 368
    v6 = v6
    var v8 uint32 = 2
    var v9 uint8 = 42
    v5 -= (v5 + (v8 ^ (v8 >> ((544) & 63))))
    v9 = v9
    v10 := GenS1{F0: v3, F1: "hello"}
    v2[42] = "foo bar"
    var v11 uint16 = 2
    for fi12 := 0; fi12 < 16; fi12++ {
        if (("" == "\n\t") && v3) {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 bool = (18446744073709551615 <= 0)
                _ = fi13
                _ = v14
            }
        } else {
            v15 := GenS1{F0: (-2.25 < 1.5), F1: ("hello" + "hello")}
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := GenS2{F0: int8(v7), F1: v3}
                fmt.Println(fi12, v7)
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        var v18 int = 479
        _ = fi12
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            v20 := []int{(fi19 | -105), (fi19 * -9223372036854775808), (fi19 & fi19)}
            _ = v20
        } else {
            v7 = (-(v7))
            var v21 int8 = -63
            v22 := GenS2{F0: v21, F1: v3}
            _ = v21
            _ = v22
        }
        v23 := []int{fi19, (fi19 & 42)}
        v8 = (v8 | 489)
        _ = fi19
        _ = v23
    }
    v24 := []int{}
    if ((-1 != -128) && false) {
        v25 := GenS2{F0: -128, F1: v3}
        if v3 {
            var v26 float64 = 3.14159
            v27 := v4[v8]
            v27 = 2147483647
            _ = v26
            _ = v27
        }
        _ = v25
    }
    if (2147483647 == (v7 | v7)) {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint32 = 2
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint = 42
                var v32 int64 = 9223372036854775807
                _ = fi30
                _ = v31
                _ = v32
            }
            fmt.Println(v5)
            _ = fi28
            _ = v29
        }
        if (("hello" + ((("" + "a") + ("foo bar" + "\n\t")) + "\n\t")) < ("\n\t" + "foo bar")) {
            if ((("foo bar" + "a") + ("foo bar" + "\n\t")) >= "hello") {
                v33 := GenS1{F0: v3, F1: ("hello" + ("foo bar" + "hello"))}
                var v34 int8 = 30
                switch "hello" {
                case "":
                    if true {
                        v34 = (v34 >> ((18446744073709551615) & 63))
                    } else {
                        v34 -= 42
                        v35 := GenS1{F0: (1 == 33), F1: "\n\t"}
                        _ = v35
                    }
                case "hello":
                    if false {
                        v7 += v7
                        v36 := GenS2{F0: v34, F1: v3}
                        var v37 uint = 2
                        _ = v36
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            v11 -= v11
                            _ = fi38
                        }
                        if v3 {
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        }
                        if len(v24) > 0 {
                            v24[0] = 2147483647
                        }
                    }
                case "\n\t":
                    fmt.Println(v34)
                default:
                    if true {
                        v6 -= v6
                    }
                }
                _ = v33
                _ = v34
            }
            v1[9223372036854775807] = 2147483647
            v39 := GenS2{F0: 0, F1: v3}
            _ = v39
        }
    }
    v8 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
