package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0() int8 {
    var v7 int64 = 0
    if true {
        v7 -= (v7 ^ 117)
        if true {
            v8 := []int{-777}
            v7 -= v7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: (^(v7)), F1: 0, F2: 4294967295}
        var v10 float64 = 0.1
        v11 := []int{42, 42}
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint16 = 32767
    if (316 > 42) {
        v12 = v12
    } else {
        if true {
            var v13 int8 = 36
            v12 = v12
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    fmt.Println(v12)
    _ = v12
    _ = v7
    return -1
}

func main() {
    var v1 float32 = 0.1
    var v2 uint = 511
    var v3 uint32 = 1
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 int32 = 956
    v3 = (v3 << ((2) & 63))
    switch 42 {
    case 128:
        v5["foo bar"] = (65535 >= 2)
    case 32767:
        var v15 bool = ((false && ((("\n\t" + ("foo bar" + (((("" + ((("foo bar" + (("hello" + "foo bar") + ("hello" + ("hello" + (("" + "\n\t") + ("\n\t" + "foo bar")))))) + "a") + ("" + "hello"))) + ("a" + ("foo bar" + "a"))) + "hello") + (("hello" + "foo bar") + "")))) == "a") && true)) || false)
        _ = v15
    case 127:
        if ((v6 << (((v2 * v2)) & 63)) <= -1) {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 string = (("hello" + ("foo bar" + (("foo bar" + (("\n\t" + "\n\t") + "foo bar")) + "\n\t"))) + "foo bar")
                _ = fi16
                _ = v17
            }
            if false {
                v6 += (v6 ^ v6)
            } else {
                fmt.Println(v6, v6)
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint32 = 149
                _ = fi18
                _ = v19
            }
        }
    default:
        v2 = 255
    }
    var v20 int8 = 1
    var v21 string = "a"
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint32 = v3
        _ = fi22
        _ = v23
    }
    v24 := []int{903, 895, -9223372036854775808}
    v25 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 662}
    var v26 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
