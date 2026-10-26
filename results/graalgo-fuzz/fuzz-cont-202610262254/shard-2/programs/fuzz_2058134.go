package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int64) uint {
    var v6 uint32 = 843
    v7 := GenS1{F0: p1, F1: 42}
    if true {
        var v8 string = p0
        v8 = ((v8 + v8) + (v8 + (v8 + (p0 + "foo bar"))))
        _ = v8
    } else {
        if (!(false && true)) {
            v6 += (v6 & v6)
        } else {
            p0 = "\n\t"
        }
        p0 = p0
        v9 := []int{-213, -2147483648}
        _ = v9
    }
    switch 42 {
    case -1:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: (-(p1)), F1: 51}
            fmt.Println(p1, v6)
            _ = fi10
            _ = v11
        }
    case 127:
        var v12 uint8 = 1
        _ = v12
    case -416:
        v6 = v6
    default:
        v13 := []int{-1}
        _ = v13
    }
    switch -1 {
    case 128:
        v6 = (v6 >> ((4294967295) & 63))
    case 127:
        switch 916 {
        case 743:
            v14 := []int{}
            _ = v14
        }
    default:
        var v15 int16 = 0
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 bool) uint16 {
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 bool = p1
        v18 := []int{}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 933}
    var v20 int = 1
    v20 = (-(v20))
    v21 := []int{v20, v20}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 1
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v5 int32 = -489
    v5 = -1
    var v22 int16 = -557
    if (1.5 >= 0.1) {
        v23 := []int{9223372036854775807}
        _ = v23
    }
    var v24 int8 = (v2 - (v2 & (v2 * (v2 - 42))))
    v25 := []int{2147483647}
    var v26 float64 = 0.0
    if ((v22 << ((genFn0((("\n\t" + "") + ("foo bar" + (("hello" + (("hello" + "\n\t") + (("hello" + ("foo bar" + "\n\t")) + (("" + ("a" + "\n\t")) + ("hello" + "\n\t"))))) + ""))), 1)) & 63)) == (v22 | v22)) {
        v26 += (v26 - 1e10)
    }
    if (42 <= 178) {
        switch 1 {
        case 9223372036854775807:
            v24 += (v2 / (((v2 >> ((0) & 63))) | 1))
        case -1:
            var v27 int64 = 42
            _ = v27
        }
        v26 += (v26 / float64((1e10 / 1.5)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
