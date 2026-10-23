package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
func genFn0() uint64 {
    v7 := GenS1{F0: 170, F1: "\n\t", F2: 128}
    var v8 float64 = (float64(1.5) + -2.25)
    _ = v7
    _ = v8
    return 35
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 4294967295
    var v3 string = ""
    var v4 string = v3
    v5 := make(map[uint32]uint32)
    var v6 int16 = -1
    v4 = "\n\t"
    var v9 string = v3
    v6 -= v6
    fmt.Println(v4, v2)
    var v10 uint = 406
    switch "hello" {
    case "":
        switch (v2 << ((1) & 63)) {
        case 850:
            var v11 uint16 = 32767
            _ = v11
        case 0:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int = fi12
                _ = len(v5)
                _ = fi12
                _ = v13
            }
        case 9223372036854775807:
            if true {
                if (0 < 182) {
                    v14 := []int{0, v2, v2, (v2 + v2)}
                    _ = v14
                } else {
                    fmt.Println(v2)
                }
            }
        }
    case "hello":
        switch 32767 {
        case -1:
            var v15 int8 = 1
            _ = v15
        }
    case "a":
        v6 += (v6 << (((v10 * (v10 % ((v10) | 1)))) & 63))
    }
    var v16 string = v3
    v17 := GenS1{F0: 0, F1: ("" + (v16 + v4)), F2: 128}
    v3 = "hello"
    v4 = (v4 + v16)
    var v18 int32 = 0
    v19 := []int{51, v2, 128, v2}
    if len(v19) > 0 {
        v19[0] = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
