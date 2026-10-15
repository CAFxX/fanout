package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
func genFn0(p0 string, p1 int64, p2 int) int {
    if false {
        p0 = p0
        fmt.Println(p2, p2)
        var v4 uint64 = 188
        _ = v4
    }
    switch p2 {
    case -877:
        fmt.Println(p1, p2, p1)
    default:
        v5 := []int{p2, 127}
        _ = v5
    }
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() int {
    if ((825 < 1) || false) {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint32 = 0
            v8 := GenS1{F0: 4294967295, F1: 42, F2: ("a" + (("a" + "") + ("hello" + (("a" + "a") + (("\n\t" + ("\n\t" + "a")) + "hello")))))}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint32 = 65535
        _ = v9
    } else {
        var v10 int64 = 9223372036854775807
        var v11 int8 = -1
        if false {
            if ((687 > -2147483648) || true) {
                v11 = 100
                v12 := GenS1{F0: 18446744073709551615, F1: 808, F2: ("hello" + "")}
                _ = v12
            } else {
                v10 += -892
            }
            v13 := GenS1{F0: 4294967295, F1: 555, F2: "hello"}
            _ = v13
        }
        _ = v10
        _ = v11
    }
    if (0 <= 2147483647) {
        var v14 uint32 = 2147483647
        for fi15 := 0; fi15 < 2; fi15++ {
            v14 += v14
            v16 := []int{}
            _ = fi15
            _ = v16
        }
        var v17 uint = 2
        _ = v14
        _ = v17
    } else {
        v18 := []int{894, 32767, -356, 128}
        var v19 int64 = 9223372036854775807
        _ = v18
        _ = v19
    }
    var v20 uint = 4294967295
    var v21 uint = 2
    _ = v20
    _ = v21
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    switch "foo bar" {
    case "foo bar":
        var v22 uint64 = 2
        _ = v22
    }
    v3[0] = ((2147483647 < 690) || (!(false || true)))
    v1 -= -1
    fmt.Println(v1)
    v1 = (v1 | (-(v1)))
    fmt.Println(v1)
    var v23 int16 = -32768
    delete(v3, genFn1())
    v24 := GenS1{F0: 1, F1: 0, F2: "hello"}
    var v25 int8 = -128
    var v26 int8 = v25
    if true {
        var v27 uint16 = 65535
        fmt.Println(v26)
        _ = v27
    }
    var v28 int16 = v23
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 float32 = 1e10
        _ = fi29
        _ = v30
    }
    var v31 string = ("foo bar" + (("" + ("hello" + "\n\t")) + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v31
}
