package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = ((1.5 * 0.1) / ((0.0 + 3.14159) + 3.14159))
    var v3 uint = 1
    v4 := make(map[uint32]int)
    v5 := []int{}
    if false {
        v2 = (((1e10 / v2) / float32(1.5)) / v2)
    }
    v2 = float32(0.1)
    v2 -= (v2 + (v2 / (-2.25 + v2)))
    switch ("\n\t" + "") {
    case "foo bar":
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 uint64 = 4294967295
            switch "hello" {
            case "hello":
                if false {
                    v7 -= v7
                    var v8 bool = (1 != 1)
                    var v9 bool = (v8 && v8)
                    _ = v8
                    _ = v9
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if true {
                        switch (fi6 / (((^(fi6))) | 1)) {
                        case 4294967295:
                            var v10 float64 = 1e10
                            _ = v10
                        default:
                            v11 := GenS1{F0: 4294967295, F1: v3, F2: 32767}
                            _ = v11
                        }
                    }
                }
            case "a":
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int16 = -75
                    _ = fi12
                    _ = v13
                }
            case "":
                v14 := GenS1{F0: 367, F1: 200, F2: (fi6 >> (((v3 + (v3 << ((v3) & 63)))) & 63))}
                _ = v14
            }
            _ = fi6
            _ = v7
        }
    case "hello":
        if (127 < 1) {
            v15 := GenS1{F0: 662, F1: v3, F2: 2147483647}
            v3 += 65535
            _ = v15
        }
    case "a":
        v16 := v1[0]
        _ = v16
    }
    _ = len(v5)
    v2 = v2
    var v17 uint = 859
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
