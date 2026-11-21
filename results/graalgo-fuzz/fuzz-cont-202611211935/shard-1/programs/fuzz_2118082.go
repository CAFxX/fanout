package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    v3 := []uint32{}
    v4 := []string{"\n\t", "hello", (("" + ("hello" + ("foo bar" + ""))) + ("a" + ""))}
    _ = len(v3)
    v5 := GenS1{F0: -1}
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 = (v6 - (^(v6)))
        _ = fi7
    }
    var v8 int = 4
    switch "" {
    case "a":
        switch v8 {
        case 128:
            v9 := GenS2{F0: 0, F1: ""}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                v6 -= 2
                var v12 int16 = 104
                v12 += (v12 ^ v12)
                var v13 string = "foo bar"
                v13 = v13
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint64 = 4294967295
            _ = v14
        }
        var v15 uint16 = 255
        _ = fi10
        _ = v15
    }
    v8 = v8
    var v16 bool = (!(4294967295 >= 649))
    v6 -= v6
    v17 := v1[int(v6)]
    delete(v1, (v8 - (v8 * (v8 / ((32767) | 1)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
