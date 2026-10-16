package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 string) uint64 {
    v6 := GenS1{F0: 1}
    v7 := []int{-983}
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []string{"foo bar", "a", "\n\t"}
    var v4 uint64 = 820
    v5 := []int{-9223372036854775808, 107}
    v8 := GenS1{F0: 1}
    switch "hello" {
    case "foo bar":
        for fi9 := 0; fi9 < 10; fi9++ {
            v2[(fi9 >> ((538) & 63))] = 2147483647
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(fi9)
                v4 -= (v4 << ((513) & 63))
                v11 := GenS2{F0: 42, F1: 32767, F2: (("hello" + "") + ("\n\t" + "foo bar"))}
                _ = fi10
                _ = v11
            }
            if true {
                v12 := GenS2{F0: 1, F1: 2, F2: "foo bar"}
                v13 := GenS3{F0: 2147483648}
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS1{F0: -77}
                    var v16 string = "hello"
                    fmt.Println(v4)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v12
                _ = v13
            } else {
                var v17 string = (("hello" + "foo bar") + "hello")
                switch "a" {
                case "":
                    switch ("\n\t" + (v17 + v17)) {
                    case "foo bar":
                        v18 := GenS1{F0: 9223372036854775807}
                        _ = v18
                    case "hello":
                        var v19 uint8 = 255
                        _ = v19
                    case "a":
                        fmt.Println(v17)
                    default:
                        v20 := GenS4{F0: 131, F1: genFn0(((((v17 + v17) + (v17 + "foo bar")) + v17) + (("foo bar" + ("\n\t" + v17)) + v17)))}
                        _ = v20
                    }
                case "\n\t":
                    if (true && (1 <= 0)) {
                        fmt.Println(v17)
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    }
                case "foo bar":
                    var v22 bool = true
                    _ = v22
                }
                _ = v17
            }
            _ = fi9
        }
    case "\n\t":
        delete(v1, 2147483647)
    default:
        _ = len(v3)
    }
    v4 -= (v4 / ((genFn0(("foo bar" + ("foo bar" + ("" + "hello"))))) | 1))
    v23 := GenS5{F0: (!(-293 != 9223372036854775807)), F1: 1, F2: 255}
    var v24 uint64 = v4
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS5{F0: false, F1: 0, F2: 30}
        v27 := GenS2{F0: 0, F1: 786, F2: ("a" + "a")}
        _ = fi25
        _ = v26
        _ = v27
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    _ = len(v3)
    switch 127 {
    case -2147483648:
        var v28 bool = false
        _ = v28
    default:
        fmt.Println(v4, v24)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
