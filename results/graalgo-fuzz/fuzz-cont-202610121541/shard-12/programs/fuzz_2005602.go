package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }

func main() {
    var v1 int = -315
    v2 := make(map[uint32]string)
    var v3 uint16 = 65535
    var v4 float32 = 3.14159
    switch "\n\t" {
    case "a":
        v5 := []int{v1}
        _ = v5
    case "":
        var v6 string = "a"
        _ = v6
    case "\n\t":
        v7 := GenS1{F0: v3, F1: 0, F2: -176}
        _ = v7
    default:
        v8 := GenS1{F0: v3, F1: 25, F2: -167}
        _ = v8
    }
    fmt.Println(v3, v3, v4)
    var v9 string = ("a" + "a")
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = 526
        switch 9223372036854775807 {
        case 917:
            v4 += ((v4 - (3.14159 + (v4 + ((float32(0.0) / float32(0.1)) / (v4 + v4))))) + 3.14159)
        case -149:
            for fi12 := 0; fi12 < 2; fi12++ {
                if true {
                    var v13 uint8 = 1
                    _ = v13
                }
                v14 := v2[4294967295]
                if (!((v1 + (fi10 ^ (v1 / (((fi12 | fi12)) | 1)))) >= v1)) {
                    var v15 int64 = -1
                    v11 = v11
                    _ = v15
                }
                _ = fi12
                _ = v14
            }
        case -489:
            switch fi10 {
            case -255:
                v16 := GenS1{F0: (^(v3)), F1: 42, F2: 0}
                _ = v16
            case 42:
                var v17 int = fi10
                _ = v17
            default:
                v9 = v9
            }
        }
        _ = fi10
        _ = v11
    }
    v9 = (((v9 + v9) + v9) + "a")
    switch (v9 + v9) {
    case "a":
        switch (v9 + v9) {
        case "foo bar":
            if (1 <= -9223372036854775808) {
                var v18 int8 = -1
                _ = v18
            } else {
                fmt.Println(v3, v4, v3)
            }
        case "\n\t":
            var v19 int16 = 1000
            _ = v19
        }
    case "":
        v3 = 2
    default:
        fmt.Println(v4, v4)
    }
    if ((785 == 4294967295) && true) {
        v20 := GenS1{F0: v3, F1: 0, F2: 0}
        _ = v20
    } else {
        var v21 string = v9
        v22 := []int{v1}
        _ = v21
        _ = v22
    }
    fmt.Println(v9, v4, v4)
    v23 := GenS1{F0: 810, F1: 1, F2: 901}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
