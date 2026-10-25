package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int32) int32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint64 = 80
        var v8 uint16 = 32767
        fmt.Println(v8, v7, p0)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -19 {
    case 750:
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 4294967295, F1: true, F2: 395}
            var v11 int16 = -1
            _ = fi9
            _ = v10
            _ = v11
        }
    case 377:
        if true {
            var v12 int = 0
            _ = v12
        }
    case 2147483647:
        var v13 uint64 = 4294967295
        _ = v13
    default:
        v14 := []int{-2147483648}
        _ = v14
    }
    var v15 uint16 = 536
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 780, F1: 0}
        _ = fi16
        _ = v17
    }
    var v18 string = "a"
    _ = p0
    _ = v15
    _ = v18
    return int32(v15)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 42
    v3 := []int{1}
    v4 := []string{"", (("\n\t" + ("foo bar" + ("" + ""))) + "foo bar")}
    v5 := make(map[string]int)
    var v19 int = 42
    switch -2147483648 {
    case 824:
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int64 = 0
            var v22 uint = 255
            for fi23 := 0; fi23 < 2; fi23++ {
                v2 -= (^(v2))
                v1[(v19 << ((0) & 63))] = "hello"
                _ = fi23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        _ = len(v5)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        switch 32768 {
        case 32767:
            v2 += v2
        case 2147483647:
            v2 += v2
        case -265:
            var v25 uint8 = 128
            _ = v25
        default:
            v19 = v19
        }
        switch (("hello" + (("\n\t" + "") + ("\n\t" + "foo bar"))) + ("hello" + ((("\n\t" + (("a" + ("a" + "")) + "a")) + "\n\t") + "a"))) {
        case "":
            var v26 string = ("" + "hello")
            _ = v26
        case "\n\t":
            var v27 string = ""
            _ = v27
        case "foo bar":
            fmt.Println(v19, v19, fi24)
        default:
            v5[(((("a" + "a") + ("a" + (("a" + "\n\t") + (("foo bar" + "hello") + ("a" + "a"))))) + "\n\t") + (("hello" + "\n\t") + ""))] = fi24
        }
        v19 += 783
        _ = fi24
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int = 32768
        _ = fi28
        _ = v29
    }
    v4 = append(v4, "")
    var v30 int = v19
    fmt.Println(v19)
    v2 -= v2
    fmt.Println(v30)
    for fi31 := 0; fi31 < 5; fi31++ {
        v2 = v2
        fmt.Println(fi31, v30, v30)
        _ = fi31
    }
    v19 -= v30
    v2 = v2
    v32 := GenS2{F0: 0, F1: 65535}
    var v33 uint32 = 1
    var v34 int16 = 32767
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 float64 = 1.5
        _ = fi35
        _ = v36
    }
    v33 = (v33 | 42)
    if (2 <= v33) {
        v5[("\n\t" + ("\n\t" + "a"))] = v19
        v34 = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
