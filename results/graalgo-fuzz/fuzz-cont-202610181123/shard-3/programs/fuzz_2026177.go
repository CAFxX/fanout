package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int; F2 string }
func genFn0() int32 {
    var v4 uint16 = 2
    if true {
        var v5 string = "foo bar"
        var v6 int32 = -392
        if true {
            var v7 float64 = (1.5 / 3.14159)
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        var v8 int16 = 42
        v9 := []int{710, -9223372036854775808, -957, -2147483648}
        for fi10 := 0; fi10 < 19; fi10++ {
            v9 = append(v9, fi10)
            for fi11 := 0; fi11 < 9; fi11++ {
                v9 = append(v9, fi11)
                _ = fi11
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    v12 := GenS1{F0: 18}
    if ((false || (!false)) && (180 >= 0)) {
        if (127 != 230) {
            var v13 uint8 = 127
            _ = v13
        } else {
            var v14 int8 = 42
            v14 -= v14
            if true {
                v15 := GenS2{F0: 2147483647, F1: -944, F2: "foo bar"}
                _ = v15
            }
            _ = v14
        }
    }
    if (0 > 42) {
        v16 := []int{-562, 42, 42}
        var v17 int = 32767
        var v18 int32 = 0
        _ = v16
        _ = v17
        _ = v18
    } else {
        v19 := GenS1{F0: 100}
        fmt.Println(v4)
        _ = v19
    }
    _ = v12
    _ = v4
    return 0
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 32767
    var v3 bool = (1.5 == 0.0)
    var v20 int8 = 42
    var v21 uint = 42
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{fi22, (fi22 * fi22), (fi22 - fi22), (fi22 % ((fi22) | 1))}
        v1 = v1
        var v24 string = ("" + "foo bar")
        _ = fi22
        _ = v23
        _ = v24
    }
    v21 = v21
    var v25 int32 = -134
    fmt.Println(v2)
    var v26 int16 = v1
    var v27 float64 = -2.25
    v25 = -138
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            v1 += v1
            _ = fi29
        }
        _ = fi28
    }
    switch (("" + (("hello" + "foo bar") + "hello")) + "") {
    case "hello":
        if (!v3) {
            for fi30 := 0; fi30 < 11; fi30++ {
                v31 := []int{fi30, (fi30 ^ fi30)}
                v32 := GenS1{F0: v20}
                fmt.Println(v27, v3)
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    case "foo bar":
        v33 := GenS2{F0: v25, F1: 0, F2: ("hello" + "a")}
        _ = v33
    }
    v25 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
