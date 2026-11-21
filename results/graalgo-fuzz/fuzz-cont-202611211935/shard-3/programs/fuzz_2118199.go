package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int8 = -128
    var v2 int16 = 1
    var v3 uint = 783
    v4 := make(map[string]int)
    var v5 int = 1
    v6 := make(map[uint32]bool)
    if ((v2 / (((v2 | (^(v2)))) | 1)) == v2) {
        v7 := GenS1{F0: ""}
        fmt.Println(v2, v5, v5)
        _ = v7
    }
    switch (^(v5)) {
    case -86:
        var v8 float32 = -2.25
        _ = v8
    case 32767:
        v9 := GenS1{F0: ""}
        _ = v9
    case -2147483648:
        v5 = (^(v5))
    }
    fmt.Println(v2, v2, v3)
    fmt.Println(v1)
    v10 := GenS2{F0: (v1 * (-(v1))), F1: (v1 >> ((v3) & 63))}
    var v11 int = 596
    v12 := GenS2{F0: 69, F1: (^(v1))}
    for fi13 := 0; fi13 < 6; fi13++ {
        if (v3 > (^(v3))) {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{fi13, v11, 9223372036854775807, (fi14 >> (((v3 % ((18446744073709551615) | 1))) & 63))}
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: -978}
            _ = v16
        }
        switch v11 {
        case -358:
            v5 = v11
        }
        var v17 int16 = v2
        _ = fi13
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
