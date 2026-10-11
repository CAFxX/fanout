package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8; F2 string }

func main() {
    var v1 uint8 = 243
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 6; fi4++ {
        delete(v3, 0)
        v5 := GenS1{F0: false, F1: ("a" + ""), F2: 127}
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 string = "foo bar"
            switch "" {
            case "":
                var v8 int8 = 42
                _ = v8
            }
            delete(v3, 65535)
            var v9 int16 = 941
            _ = fi6
            _ = v7
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    v10 := GenS1{F0: true, F1: "foo bar", F2: 127}
    v11 := []int{-9223372036854775808}
    delete(v2, 204)
    v1 += (^(v1))
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: true, F1: "", F2: v1}
        if (871 >= 762) {
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 -= v1
                _ = fi14
            }
            fmt.Println(v1, v1)
        }
        v1 = (v1 % ((102) | 1))
        _ = fi12
        _ = v13
    }
    v1 = v1
    v15 := GenS2{F0: 100}
    v16 := GenS2{F0: 42}
    var v17 uint8 = v1
    v18 := GenS1{F0: (v17 > (v1 | 42)), F1: (("a" + "\n\t") + (("\n\t" + "\n\t") + "hello")), F2: (v17 & 2)}
    v2[51] = true
    v17 -= (v17 << ((0) & 63))
    var v19 string = "hello"
    switch -9223372036854775808 {
    case -9223372036854775808:
        v20 := GenS3{F0: (v1 / ((6) | 1))}
        _ = v20
    }
    var v21 uint8 = v1
    var v22 int64 = 1
    v23 := GenS4{F0: 42, F1: 10, F2: v19}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
