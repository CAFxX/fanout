package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 uint16; F2 int32 }
type GenS4 struct { F0 bool }
func genFn0(p0 int16, p1 int16, p2 uint64) uint16 {
    switch 0 {
    case -413:
        v7 := []int{32768, -9223372036854775808}
        _ = v7
    }
    v8 := []int{138, -473}
    for fi9 := 0; fi9 < 20; fi9++ {
        if ((false || true) || false) {
            p2 = p2
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int8 = -1
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 float64 = 0.1
            _ = fi12
            _ = v13
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            p2 = p2
        }
        _ = len(v8)
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 2
    v3 := []string{}
    v4 := []int{-1, -1}
    var v5 int32 = -1
    v6 := make(map[int]string)
    fmt.Println(v1)
    v14 := GenS1{F0: 1}
    var v15 int32 = (^(v5))
    if true {
        var v16 int64 = int64(v15)
        delete(v6, -395)
        v17 := GenS2{F0: "", F1: (v16 % (((-(v16))) | 1))}
        _ = v16
        _ = v17
    }
    if (100 >= 1) {
        v2 -= 159
        v2 = 255
    }
    if false {
        delete(v6, 565)
        var v18 float32 = 1.5
        _ = len(v6)
        _ = v18
    }
    if ("hello" <= ("" + ("\n\t" + "a"))) {
        var v19 int32 = v5
        var v20 int8 = -1
        _ = v19
        _ = v20
    } else {
        v21 := GenS3{F0: 32767, F1: 42, F2: (^(v15))}
        fmt.Println(v2)
        _ = v21
    }
    v22 := GenS4{F0: ("\n\t" >= (("\n\t" + ("\n\t" + ("a" + "hello"))) + ("" + ("" + ("hello" + ((("foo bar" + "hello") + "foo bar") + ("foo bar" + ("foo bar" + ("a" + "")))))))))}
    v5 -= -619
    var v23 string = "hello"
    v24 := v6[-880]
    var v25 float64 = (v1 * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
