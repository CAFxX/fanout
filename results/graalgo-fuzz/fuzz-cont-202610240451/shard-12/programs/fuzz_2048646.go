package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint, p1 int32) int64 {
    v4 := []int{84, 32767}
    var v5 string = ("\n\t" + "")
    fmt.Println(v5)
    var v6 int16 = 1000
    v5 = v5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 83
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -489
    v3 := make(map[uint32]int)
    v7 := GenS1{F0: v2}
    if ("" <= "hello") {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: (-(v2))}
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := GenS2{F0: (v2 + -1), F1: v2}
                _ = fi10
                _ = v11
            }
            delete(v1, fi8)
            _ = fi8
            _ = v9
        }
    }
    v12 := GenS3{F0: ("foo bar" + ("a" + "hello")), F1: -490}
    v13 := v1[-304]
    var v14 uint8 = 127
    v13 = false
    v14 -= v14
    for fi15 := 0; fi15 < 11; fi15++ {
        v14 -= v14
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v2, fi16, v14)
            _ = fi16
        }
        var v17 float64 = -2.25
        _ = fi15
        _ = v17
    }
    v18 := []int{32768, -1, 9223372036854775807}
    switch -2147483648 {
    case 5:
        v19 := GenS3{F0: (("\n\t" + (("" + "hello") + "a")) + "a"), F1: 1}
        _ = v19
    case 545:
        if v13 {
            delete(v3, 4294967295)
            v14 -= 0
        } else {
            if (!(65535 != 0)) {
                switch 32767 {
                case 983:
                    v14 = 31
                }
                var v20 int16 = 32767
                _ = v20
            } else {
                switch 2147483647 {
                case 782:
                    v2 = v2
                }
                var v21 int64 = 1
                _ = v21
            }
            v13 = (373 == 65535)
        }
    case 127:
        if len(v18) > 0 {
            v18[0] = 644
        }
    }
    switch ("hello" + "hello") {
    case "\n\t":
        v22 := GenS4{F0: 0, F1: 1}
        _ = v22
    case "a":
        v14 -= 128
    case "hello":
        v23 := v1[-2147483648]
        _ = v23
    }
    var v24 uint32 = 0
    for fi25 := 0; fi25 < 14; fi25++ {
        v18 = append(v18, 9223372036854775807)
        v26 := GenS5{F0: 1}
        var v27 int16 = 1000
        _ = fi25
        _ = v26
        _ = v27
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v2)
    default:
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := v1[30]
            var v30 bool = v29
            var v31 bool = (100 > 73)
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v7
}
