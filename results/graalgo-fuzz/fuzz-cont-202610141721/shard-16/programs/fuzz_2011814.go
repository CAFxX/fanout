package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = 0
        _ = fi7
        _ = v8
    }
    var v9 float32 = 1.5
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -53
    v3 := []int{-9223372036854775808, -2147483648}
    v4 := make(map[string]int)
    var v5 int8 = v2
    v6 := make(map[int]int)
    var v10 uint32 = 954
    if (3.14159 > (1e10 - float64((float32(-2.25) * 1e10)))) {
        fmt.Println(v5, v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v10, v2)
            _ = fi11
        }
        _ = len(v4)
    }
    var v12 int8 = (v5 >> ((724) & 63))
    var v13 uint = 480
    v14 := GenS1{F0: "a", F1: -2147483648}
    var v15 uint16 = 518
    var v16 int8 = (v2 << ((v13) & 63))
    v10 += (v10 - v10)
    switch ((("a" + "\n\t") + ("hello" + "\n\t")) + ("\n\t" + "foo bar")) {
    case "a":
        switch ("hello" + "\n\t") {
        case "hello":
            var v17 string = ""
            _ = v17
        case "foo bar":
            v18 := GenS1{F0: "", F1: 2147483647}
            _ = v18
        default:
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v16)
                _ = fi19
            }
        }
    }
    fmt.Println(v13, v10)
    var v20 float64 = 1e10
    if (0 != 2) {
        switch 4294967295 {
        case -1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -62:
            var v21 uint64 = uint64(v15)
            _ = v21
        }
    } else {
        v16 -= 100
    }
    v16 += -24
    v1["\n\t"] = (uint64(v10) > 446)
    fmt.Println(v16, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
