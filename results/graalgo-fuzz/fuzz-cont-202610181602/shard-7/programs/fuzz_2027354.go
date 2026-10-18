package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    var v4 uint16 = 0
    v5 := make(map[string]bool)
    var v6 int16 = 42
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v4)
        v8 := []int{}
        v3 += (^(v3))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: (v3 + v3)}
    var v10 uint64 = 18446744073709551615
    if true {
        v11 := GenS1{F0: 42, F1: 0}
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 17; fi14++ {
        switch (^(fi14)) {
        case 2147483647:
            fmt.Println(fi14, v10, fi14)
        case -825:
            v3 += v3
        case 9223372036854775807:
            var v15 uint32 = 1
            _ = v15
        default:
            var v16 string = (("\n\t" + (("foo bar" + ("hello" + ((((("" + "a") + (("hello" + "a") + "hello")) + ((("\n\t" + "foo bar") + "\n\t") + "a")) + "hello") + "hello"))) + "foo bar")) + "hello")
            _ = v16
        }
        _ = fi14
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            _ = len(v5)
        case "\n\t":
            var v17 uint16 = v4
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: uint(v4), F1: 337}
                v13 = append(v13, 42)
                var v20 float32 = 0.1
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: 255, F1: (v10 | v10)}
    if (true || true) {
        v22 := GenS1{F0: 255, F1: v3}
        _ = v22
    } else {
        v10 = 1
    }
    switch ("foo bar" + ("" + "a")) {
    case "\n\t":
        v1[4294967295] = 2
    case "a":
        var v23 float32 = (0.0 / 1e10)
        _ = v23
    }
    if true {
        var v24 int8 = 0
        v25 := GenS2{F0: v24, F1: (v4 % (((v4 | v4)) | 1))}
        fmt.Println(v6, v10)
        _ = v24
        _ = v25
    } else {
        fmt.Println(v10)
        var v26 string = "a"
        v4 = 42
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 int8 = 100
            _ = fi28
            _ = v29
        }
        var v30 int8 = int8(fi27)
        v3 = 0
        _ = fi27
        _ = v30
    }
    var v31 uint16 = 939
    if true {
        _ = len(v13)
        for fi32 := 0; fi32 < 3; fi32++ {
            if (fi32 == fi32) {
                v5["a"] = false
                _ = len(v13)
                v5["a"] = false
            }
            _ = fi32
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 uint = 0
            v35 := GenS1{F0: 42, F1: v3}
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        v6 -= -682
    }
    switch "\n\t" {
    case "hello":
        v36 := GenS1{F0: 4294967295, F1: 18446744073709551615}
        _ = v36
    case "":
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := v5[(("a" + "\n\t") + "a")]
            v39 := GenS2{F0: 42, F1: 449}
            v40 := GenS1{F0: 42, F1: v3}
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
