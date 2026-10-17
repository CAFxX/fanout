package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 32768
    var v3 uint64 = 715
    var v4 int32 = 620
    var v5 int32 = (v4 | -320)
    fmt.Println(v4, v4)
    v3 -= 671
    v2 = 2147483647
    v4 = (v5 % ((-1) | 1))
    switch "" {
    case "":
        for fi6 := 0; fi6 < 2; fi6++ {
            v2 += (v2 - (fi6 - fi6))
            _ = fi6
        }
    default:
        delete(v1, v2)
    }
    v7 := GenS1{F0: v2}
    v5 = v5
    v8 := []int{(v2 << ((4294967295) & 63))}
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v2 += fi9
            _ = fi10
        }
        v11 := GenS2{F0: 127, F1: 42, F2: 589}
        _ = fi9
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int32 = (^(v4))
            switch "a" {
            case "\n\t":
                for fi14 := 0; fi14 < 8; fi14++ {
                    v1[fi12] = "\n\t"
                    if len(v8) > 0 {
                        v8[0] = 128
                    }
                    v1[1] = ""
                    var v15 float32 = -2.25
                    var v16 string = ((("foo bar" + "") + "foo bar") + ("" + ("a" + "")))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "":
                var v17 int16 = -900
                _ = v17
            default:
                switch ("a" + "hello") {
                case "foo bar":
                    v2 = (v2 * fi12)
                case "hello":
                    v18 := GenS2{F0: 127, F1: 0, F2: 65535}
                    _ = v18
                case "a":
                    v3 = 886
                default:
                    fmt.Println(v4, v4)
                }
            }
            _ = fi12
            _ = v13
        }
        if true {
            var v19 uint64 = (^(v3))
            _ = v19
        }
    } else {
        if true {
            var v20 uint16 = 32767
            _ = v20
        } else {
            var v21 int64 = 455
            for fi22 := 0; fi22 < 7; fi22++ {
                switch 1 {
                case -717:
                    v21 = v21
                case 857:
                    var v23 uint8 = 2
                    _ = v23
                case 9223372036854775807:
                    fmt.Println(v2, v21)
                }
                _ = fi22
            }
            _ = v21
        }
        if (-2.25 == (3.14159 / (1e10 + 0.1))) {
            var v24 int8 = -128
            var v25 uint = 42
            _ = v24
            _ = v25
        } else {
            fmt.Println(v4, v5, v2)
            if true {
                var v26 uint64 = v3
                for fi27 := 0; fi27 < 10; fi27++ {
                    v26 = (v26 ^ v3)
                    v2 = int(v26)
                    _ = fi27
                }
                _ = v26
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint8 = 127
                v1[0] = "\n\t"
                _ = fi28
                _ = v29
            }
        }
    }
    delete(v1, int(v5))
    _ = len(v1)
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 uint64 = (v3 >> ((18446744073709551615) & 63))
        delete(v1, 42)
        _ = len(v1)
        _ = fi30
        _ = v31
    }
    v4 = (v5 * v5)
    for fi32 := 0; fi32 < 1; fi32++ {
        v2 = fi32
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
