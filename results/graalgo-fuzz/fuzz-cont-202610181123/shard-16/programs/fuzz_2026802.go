package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 uint = 2
    var v2 uint8 = 0
    var v3 int16 = 0
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 2; fi5++ {
        if (922 == 0) {
            switch "\n\t" {
            case "":
                switch (fi5 << (((v1 & v1)) & 63)) {
                case 1:
                    var v6 int16 = v3
                    _ = v6
                default:
                    v7 := GenS1{F0: 2147483647}
                    _ = v7
                }
            case "\n\t":
                for fi8 := 0; fi8 < 15; fi8++ {
                    v3 = (v3 ^ (v3 ^ v3))
                    var v9 uint64 = 4294967295
                    var v10 bool = ("hello" < v4)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "hello":
                for fi11 := 0; fi11 < 13; fi11++ {
                    switch fi5 {
                    case 128:
                        v2 += v2
                    }
                    v1 += v1
                    _ = fi11
                }
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(v1)
                v13 := GenS1{F0: 165}
                if (3.14159 >= 1e10) {
                    fmt.Println(v1, v4)
                    fmt.Println(fi12, fi12, v2)
                    fmt.Println(fi5, fi12, fi5)
                }
                _ = fi12
                _ = v13
            }
        } else {
            v14 := GenS1{F0: -2147483648}
            if (int8(v1) < 42) {
                var v15 uint16 = 255
                v16 := GenS2{F0: int8(v1), F1: v4}
                _ = v15
                _ = v16
            }
            _ = v14
        }
        _ = fi5
    }
    v17 := []int{-2147483648, 128, 4294967295, -1}
    var v18 string = v4
    var v19 uint64 = 598
    var v20 uint = (^(v1))
    v19 = (v19 / (((^(v19))) | 1))
    var v21 int32 = -2147483648
    v22 := GenS1{F0: -602}
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
