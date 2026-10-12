package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int32) bool {
    var v5 uint64 = 479
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 = (((v6 + (v6 + "a")) + "foo bar") + "")
        var v8 uint8 = 26
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int8 = -36
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 int8 = 1
    _ = p0
    _ = v12
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = (v1 | (v1 / ((v1) | 1)))
    v13 := []int{(v1 << ((142) & 63)), v1, v1, v1}
    var v14 int = -916
    v1 = 4294967295
    if ((v14 % ((v14) | 1)) <= (v14 * v1)) {
        var v15 uint32 = 586
        for fi16 := 0; fi16 < 19; fi16++ {
            v15 = v15
            fmt.Println(v15, v15, v1)
            for fi17 := 0; fi17 < 8; fi17++ {
                v1 += v1
                fmt.Println(v14)
                _ = fi17
            }
            _ = fi16
        }
        if false {
            v18 := GenS1{F0: ("hello" + ("" + ""))}
            _ = v18
        }
        _ = v15
    }
    var v19 bool = false
    v20 := GenS1{F0: ""}
    v14 = 4294967295
    if v19 {
        for fi21 := 0; fi21 < 18; fi21++ {
            v19 = (18446744073709551615 > 549)
            v19 = (-127 == 42)
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi21
        }
        fmt.Println(v19, v14, v14)
    }
    var v22 int8 = 42
    v14 -= v1
    switch (^(v1)) {
    case 123:
        v14 -= -2147483648
    case -834:
        v1 += v1
    }
    if v19 {
        var v23 int = (v1 % ((v1) | 1))
        _ = v23
    }
    var v24 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
