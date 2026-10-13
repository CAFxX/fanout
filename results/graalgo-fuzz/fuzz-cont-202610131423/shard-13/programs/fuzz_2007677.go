package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 bool }
func genFn0() int64 {
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 9223372036854775807, F1: (("foo bar" + ("a" + ("" + (((("a" + "") + "\n\t") + "\n\t") + "foo bar")))) + "\n\t")}
    v5 += v5
    switch 569 {
    case 289:
        if ((v5 * 42) != 16) {
            v9 := GenS1{F0: -2147483648, F1: "hello"}
            v10 := GenS2{F0: -36}
            _ = v10
            _ = v9
        }
    case 32768:
        v11 := []int{127, 0, 793, 672}
        _ = v11
    case 0:
        v5 += (^(v5))
    }
    _ = v5
    _ = v8
    return int64(v5)
}

func main() {
    var v1 uint = 255
    var v2 int64 = -9223372036854775808
    v3 := []string{("\n\t" + "\n\t"), (("a" + "\n\t") + (("hello" + "") + ("hello" + "a"))), "a", ((("a" + "a") + "a") + "a")}
    var v4 uint = (v1 % ((1) | 1))
    v12 := GenS3{F0: 100, F1: (1000 >= 142)}
    v2 = v2
    for fi13 := 0; fi13 < 11; fi13++ {
        v2 -= v2
        if (-1 >= 755) {
            for fi14 := 0; fi14 < 19; fi14++ {
                if len(v3) > 0 {
                    v3[0] = (("hello" + ((("a" + "foo bar") + "hello") + "foo bar")) + "\n\t")
                }
                _ = fi14
            }
            v15 := GenS3{F0: 100, F1: (3.14159 != 3.14159)}
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v1 = v4
        v17 := GenS3{F0: -59, F1: true}
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 32767
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 float64 = -2.25
        fmt.Println(v18, v1, v1)
        var v21 uint16 = 65535
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int8 = 1
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v1, v18, v18)
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(fi23, v18)
            _ = len(v3)
            _ = fi24
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            v18 -= v18
            v4 -= 2
            _ = fi25
        }
        _ = fi23
    }
    var v26 uint32 = 887
    var v27 uint16 = v18
    v1 -= (v1 * 1)
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 float32 = 0.1
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 bool = ("hello" < ("hello" + "hello"))
        _ = fi30
        _ = v31
    }
    if len(v3) > 0 {
        v3[0] = ("" + ("a" + (("a" + "hello") + "\n\t")))
    }
    var v32 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
