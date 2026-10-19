package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
func genFn0(p0 uint8) uint {
    switch 440 {
    case -9223372036854775808:
        var v5 uint = 918
        _ = v5
    default:
        v6 := GenS1{F0: 32767, F1: 743, F2: 61}
        _ = v6
    }
    fmt.Println(p0)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = 10
        p0 = (p0 | p0)
        _ = fi7
        _ = v8
    }
    _ = p0
    return 4294967295
}
func genFn1() int {
    if false {
        var v9 uint16 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = (-(fi10))
        v12 := GenS1{F0: -152, F1: 4294967295, F2: 75}
        for fi13 := 0; fi13 < 1; fi13++ {
            switch ("foo bar" + "\n\t") {
            case "hello":
                var v14 uint32 = 4294967295
                _ = v14
            case "\n\t":
                fmt.Println(fi10)
            }
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    switch (("hello" + ("" + ("" + ""))) + "hello") {
    case "foo bar":
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := []int{626, (fi16 | 127), (fi15 ^ (^(fi16)))}
                    if len(v17) > 0 {
                        v17[0] = 127
                    }
                    var v18 int8 = 121
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
            var v19 bool = (!(42 > 18446744073709551615))
            _ = v19
        } else {
            var v20 uint = 4294967295
            _ = v20
        }
    case "":
        v21 := []int{42}
        _ = v21
    default:
        var v22 uint16 = 32767
        _ = v22
    }
    var v23 uint16 = 42
    _ = v23
    return -696
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 2
    v3 := make(map[int]string)
    var v4 bool = true
    var v24 float32 = (1e10 / 0.1)
    _ = len(v1)
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            fmt.Println(v24, fi27)
            fmt.Println(v25, fi26, v25)
            _ = fi27
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            v4 = (1 != 788)
            var v29 uint64 = 4294967295
            _ = fi28
            _ = v29
        }
        _ = fi26
    }
    v24 = 3.14159
    fmt.Println(v2, v24, v25)
    var v30 int16 = 423
    v31 := GenS1{F0: v30, F1: 182, F2: 42}
    if (!true) {
        v1[""] = (v25 * (v25 | v25))
    }
    var v32 int64 = 587
    v24 = v24
    v33 := []int{32768, -129, 127, -755}
    fmt.Println(v4)
    _ = len(v33)
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint32 = (v25 % ((2) | 1))
            _ = fi35
            _ = v36
        }
        var v37 int64 = -1
        _ = fi34
        _ = v37
    }
    v4 = v4
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 string = "\n\t"
        fmt.Println(fi38, v39, v30)
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v30 += (v30 >> ((18446744073709551615) & 63))
        v25 += 2
        fmt.Println(v30, v4, v2)
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
