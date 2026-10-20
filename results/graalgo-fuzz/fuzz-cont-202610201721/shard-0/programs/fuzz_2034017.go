package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    var v1 int32 = 758
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "foo bar"}
    v4 := make(map[uint32]string)
    var v5 uint16 = 2
    v6 := []int{}
    v5 = (v5 / ((0) | 1))
    for fi7 := 0; fi7 < 10; fi7++ {
        if (964 > fi7) {
            var v8 uint8 = 52
            var v9 int32 = v1
            _ = v8
            _ = v9
        } else {
            var v10 string = ("\n\t" + "hello")
            fmt.Println(v10, v10, v5)
            v10 = v10
            _ = v10
        }
        _ = fi7
    }
    if ((v5 > (v5 - 1)) && ((1 > 65535) || true)) {
        var v11 int64 = 1
        v12 := GenS1{F0: ("foo bar" + "foo bar"), F1: 1, F2: 759}
        _ = v11
        _ = v12
    }
    var v13 int8 = int8(v5)
    fmt.Println(v13)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v14 int16 = 32767
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
