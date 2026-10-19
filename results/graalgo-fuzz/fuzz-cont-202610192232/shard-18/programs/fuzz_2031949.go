package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint }

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]uint32)
    var v3 uint = 255
    v3 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    if false {
        fmt.Println(v3)
        var v4 float32 = (1.5 + ((1.5 + 0.1) * (-2.25 + 0.1)))
        _ = v4
    } else {
        v3 += (v3 / (((v3 | (^(v3)))) | 1))
        for fi5 := 0; fi5 < 5; fi5++ {
            v3 = 0
            _ = fi5
        }
    }
    v6 := GenS1{F0: -1, F1: 255, F2: v3}
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 bool = (2 == 596)
            v11 := GenS1{F0: (fi8 << ((783) & 63)), F1: 2, F2: 1}
            switch (fi8 - fi9) {
            case 0:
                var v12 float64 = 3.14159
                _ = v12
            case 42:
                fmt.Println(v7, v10, v3)
            case 971:
                v13 := GenS1{F0: (^(fi9)), F1: 127, F2: 452}
                _ = v13
            }
            if v10 {
                if ((v10 && (65535 < 2147483648)) && (v10 && v10)) {
                    v14 := GenS1{F0: fi9, F1: 177, F2: (v3 + (^(v3)))}
                    v7 = ""
                    _ = v14
                }
            }
            if (1 >= 42) {
                var v15 int8 = 94
                fmt.Println(v10, fi8, fi9)
                _ = v15
            } else {
                fmt.Println(fi8, fi9, v7)
                delete(v2, ((((("\n\t" + ((v7 + v7) + v7)) + v7) + v7) + v7) + v7))
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(fi8)
        _ = len(v2)
        _ = fi8
    }
    v16 := GenS1{F0: 9223372036854775807, F1: 255, F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
