package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0() int16 {
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint32 = 2147483648
            var v9 uint16 = 541
            v8 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := []int{32767, 32767, 127, 1}
        var v11 uint32 = 65535
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        switch ("a" + "foo bar") {
        case "":
            var v13 float64 = 1.5
            _ = v13
        case "hello":
            var v14 uint = 65535
            _ = v14
        case "foo bar":
            var v15 int32 = 292
            _ = v15
        }
        var v16 float64 = (0.1 + 3.14159)
        _ = fi12
        _ = v16
    }
    return -581
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := []string{"", (((("foo bar" + "\n\t") + "hello") + "foo bar") + "a"), "a", ("\n\t" + "hello")}
    var v4 int8 = 126
    var v5 uint64 = 1
    var v6 bool = false
    v17 := GenS1{F0: 0, F1: 742}
    v18 := GenS1{F0: 2147483647, F1: 519}
    if v6 {
        var v19 uint16 = 2
        v19 -= (v19 << ((1) & 63))
        _ = v19
    }
    v4 = (^(v4))
    var v20 bool = v6
    v4 += v4
    var v21 int32 = 42
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint = 38
        _ = fi22
        _ = v23
    }
    var v24 float32 = 3.14159
    if v6 {
        for fi25 := 0; fi25 < 13; fi25++ {
            fmt.Println(v20, v20)
            var v26 float32 = v24
            _ = fi25
            _ = v26
        }
    } else {
        var v27 float32 = 0.1
        fmt.Println(v21, v21)
        _ = v27
    }
    v4 += -19
    var v28 uint8 = 1
    v29 := GenS1{F0: (^(v21)), F1: 0}
    switch 32768 {
    case 783:
        var v30 uint8 = (v28 >> ((392) & 63))
        _ = v30
    }
    v31 := v2[4294967295]
    var v32 int32 = int32(v4)
    for fi33 := 0; fi33 < 3; fi33++ {
        v4 -= v4
        var v34 int8 = v4
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
