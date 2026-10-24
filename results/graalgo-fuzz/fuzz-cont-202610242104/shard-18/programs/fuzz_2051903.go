package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }

func main() {
    var v1 float64 = ((1e10 / 3.14159) / 0.1)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := []int{-381, 32767}
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 14; fi6++ {
        switch (fi6 & (fi6 >> ((4294967295) & 63))) {
        case -1:
            var v7 int8 = 12
            _ = v7
        case -570:
            var v8 bool = ((!true) && true)
            _ = v8
        }
        var v9 int = fi6
        if (false || ((42 != 100) || false)) {
            v10 := GenS1{F0: ("hello" + (("hello" + "") + (("hello" + "foo bar") + ("hello" + "\n\t")))), F1: "a", F2: 42}
            var v11 int32 = 42
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int16 = 214
        var v14 int8 = 1
        if (65535 != 615) {
            if false {
                v1 = v1
                v1 = float64(1e10)
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    if (((!false) && false) || (true && (!(1e10 > 1.5)))) {
        if false {
            delete(v2, ("foo bar" + "\n\t"))
        }
        switch 128 {
        case 4294967295:
            v15 := GenS1{F0: "foo bar", F1: "foo bar", F2: 65535}
            _ = v15
        default:
            switch ("" + ("\n\t" + "foo bar")) {
            case "":
                fmt.Println(v1)
            default:
                v1 -= ((-2.25 - (0.0 - (v1 + v1))) - 1.5)
            }
        }
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 string = ""
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v5, v1)
    var v18 uint = 255
    v19 := GenS1{F0: "\n\t", F1: "a", F2: 811}
    fmt.Println(v1, v18, v18)
    if len(v4) > 0 {
        v4[0] = 351
    }
    fmt.Println(v18)
    v1 = (((v1 - -2.25) + v1) + 1e10)
    var v20 float32 = -2.25
    v4 = append(v4, -623)
    v18 += v18
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 := v2[("hello" + ("\n\t" + ""))]
    var v22 uint8 = 128
    if true {
        if true {
            fmt.Println(v18, v21)
            var v23 string = ((((("\n\t" + "foo bar") + "foo bar") + "foo bar") + "\n\t") + "foo bar")
            _ = v23
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v22 -= v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
