package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 bool, p1 int) uint32 {
    v5 := []int{(p1 - (p1 - p1)), p1}
    v5 = append(v5, 4294967295)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint8 = 42
            fmt.Println(p0, p1)
            p1 += fi7
            _ = fi7
            _ = v8
        }
        fmt.Println(p0, p0, fi6)
        _ = fi6
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if true {
            var v10 uint8 = 20
            var v11 uint32 = 970
            var v12 bool = (!((!p0) || ((1 <= 1) || (p0 || ((42 < 0) && p0)))))
            _ = v10
            _ = v11
            _ = v12
        }
        fmt.Println(fi9)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := []string{"hello", ("foo bar" + "\n\t")}
    v2 := []string{("a" + (("foo bar" + ("hello" + "hello")) + ("a" + "a"))), "foo bar", "foo bar"}
    v3 := make(map[uint32]int)
    v4 := []string{"\n\t", ((("foo bar" + ("foo bar" + (("hello" + "a") + ("foo bar" + "foo bar")))) + "\n\t") + ("a" + "")), "", "foo bar"}
    var v13 string = "\n\t"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 int64 = -273
    delete(v3, 233)
    v14 = -9223372036854775808
    v15 := GenS1{F0: 2, F1: -1, F2: -9223372036854775808}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32768 {
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 2147483647:
        v14 = -9223372036854775808
    }
    if len(v1) > 0 {
        v1[0] = v13
    }
    fmt.Println(v14)
    fmt.Println(v13)
    var v16 string = ("" + v13)
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: 0, F1: 480, F2: fi17}
        switch fi17 {
        case 9223372036854775807:
            var v19 bool = (!(18446744073709551615 >= 42))
            _ = v19
        }
        v16 = "a"
        _ = fi17
        _ = v18
    }
    if true {
        fmt.Println(v16)
    } else {
        var v20 float64 = 3.14159
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v20)
            v22 := GenS1{F0: 4294967295, F1: 13, F2: fi21}
            _ = fi21
            _ = v22
        }
        if true {
            var v23 float32 = 1e10
            v20 = 1e10
            fmt.Println(v13, v16, v13)
            _ = v23
        }
        _ = v20
    }
    var v24 bool = (100 <= -1)
    v14 = v14
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
