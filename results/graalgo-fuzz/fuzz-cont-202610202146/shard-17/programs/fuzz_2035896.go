package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 bool, p2 int8) int16 {
    if (-2.25 < 1e10) {
        v5 := []int{-9223372036854775808}
        var v6 uint16 = 1
        _ = v5
        _ = v6
    }
    switch 422 {
    case 0:
        var v7 string = ("" + ("a" + "a"))
        _ = v7
    case 755:
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 65535}
            _ = fi8
            _ = v9
        }
    default:
        var v10 int32 = -1
        _ = v10
    }
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint8, p1 bool, p2 uint16) int64 {
    v11 := []int{}
    fmt.Println(p0, p1)
    var v12 int = 569
    if (42 > 1) {
        var v13 uint16 = 65535
        v12 += v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1000
}

func main() {
    var v1 int = 127
    v2 := make(map[int]string)
    v3 := []int{32768, 32767, v1}
    var v4 string = "a"
    v2[376] = "\n\t"
    if true {
        switch v1 {
        case 42:
            v14 := v2[(v1 & 0)]
            _ = v14
        case -1:
            switch v1 {
            case 411:
                v3 = append(v3, 2147483647)
            case 32768:
                var v15 int8 = -128
                _ = v15
            }
        case -2147483648:
            v16 := GenS1{F0: 438}
            _ = v16
        default:
            v1 = -201
        }
        switch (v4 + ((v4 + v4) + v4)) {
        case "foo bar":
            if (!(-529 != -2147483648)) {
                var v17 string = v4
                _ = v17
            }
        }
    } else {
        if (!((true && true) || true)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v18 int64 = 9223372036854775807
            if (v1 == v1) {
                var v19 uint32 = 42
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int32 = 1
                    v21 = (v21 | 222)
                    _ = fi20
                    _ = v21
                }
                var v22 uint8 = 38
                _ = v19
                _ = v22
            }
            _ = v18
        }
        var v23 uint64 = 31
        v2[v1] = "foo bar"
        _ = v23
    }
    fmt.Println(v1, v1)
    fmt.Println(v4)
    v2[v1] = v4
    v24 := GenS1{F0: 741}
    v25 := GenS2{F0: true}
    v4 = "\n\t"
    _ = len(v3)
    var v26 int = v1
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 float64 = float64(-2.25)
        var v30 int8 = 42
        v26 -= (v26 & v26)
        _ = fi28
        _ = v29
        _ = v30
    }
    delete(v2, v26)
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
