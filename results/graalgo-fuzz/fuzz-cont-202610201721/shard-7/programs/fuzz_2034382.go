package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 1
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 uint8 = 42
    v6 := []int{}
    if true {
        if true {
            var v7 int32 = 1
            v8 := GenS1{F0: -32768, F1: 42}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float32 = 0.1
            v2 += 42
            _ = fi9
            _ = v10
        }
    }
    _ = len(v4)
    _ = len(v1)
    if (!true) {
        switch (("" + ("foo bar" + ((("hello" + "hello") + "a") + "\n\t"))) + "\n\t") {
        case "hello":
            v5 -= v5
        }
        var v11 uint16 = 268
        if (4294967295 < 42) {
            v2 += v2
            v12 := GenS1{F0: 32767, F1: 0}
            _ = v12
        }
        _ = v11
    }
    v13 := GenS1{F0: -1, F1: 254}
    v5 = (v5 ^ 255)
    fmt.Println(v2, v5)
    v5 += 127
    if (868 <= 42) {
        delete(v1, 1)
        var v14 uint16 = 2
        _ = v14
    } else {
        var v15 int = (-(v2))
        _ = v15
    }
    if false {
        var v16 float64 = 0.1
        v5 = (v5 | v5)
        _ = v16
    } else {
        _ = len(v3)
        var v17 uint16 = 2
        _ = v17
    }
    if (1.5 <= float32(1e10)) {
        v18 := v3[2147483647]
        v5 = 0
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 bool = false
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v21 bool = (100 > -1)
    fmt.Println(v2, v5)
    v22 := GenS1{F0: 180, F1: 2}
    if (true && true) {
        v23 := GenS1{F0: 1, F1: 53}
        _ = v23
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v24 := GenS1{F0: -41, F1: 0}
    v2 -= 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
