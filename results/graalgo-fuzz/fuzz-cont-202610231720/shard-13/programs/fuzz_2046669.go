package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
func genFn0() int {
    var v6 string = ("foo bar" + ("hello" + "a"))
    fmt.Println(v6)
    _ = v6
    return 127
}

func main() {
    var v1 float32 = (float32((1.5 / 1e10)) * 1e10)
    v2 := []int{}
    var v3 uint64 = 4294967295
    v4 := make(map[int]int)
    var v5 int32 = -1
    _ = len(v2)
    v5 -= v5
    if len(v2) > 0 {
        v2[0] = 127
    }
    if true {
        var v7 int16 = 32767
        switch (("" + "\n\t") + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v8 string = "hello"
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 12; fi9++ {
                if (v1 <= v1) {
                    fmt.Println(v3, v7)
                    v7 -= 1000
                }
                _ = fi9
            }
        }
        _ = v7
    }
    v1 += 1e10
    v3 += v3
    var v10 uint64 = (v3 | (v3 | (v3 & (^(v3)))))
    v11 := GenS1{F0: "a", F1: true, F2: 0}
    for fi12 := 0; fi12 < 8; fi12++ {
        if (true || true) {
            v5 = v5
        }
        _ = fi12
    }
    var v13 uint64 = v3
    if false {
        v14 := GenS1{F0: "hello", F1: false, F2: 4294967295}
        var v15 int64 = -1
        _ = v14
        _ = v15
    } else {
        fmt.Println(v3, v3, v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
