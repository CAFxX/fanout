package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int; F2 int16 }

func main() {
    var v1 bool = true
    var v2 int32 = 887
    var v3 uint8 = 125
    var v4 string = ""
    var v5 uint64 = 409
    var v6 uint64 = v5
    var v7 int = 1
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v4, v4)
        _ = fi8
    }
    fmt.Println(v3, v1)
    var v9 float64 = (0.0 - (1.5 / ((-2.25 + -2.25) / 1.5)))
    v6 = (v5 - 0)
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{v7, fi10, int(v5)}
        _ = fi10
        _ = v11
    }
    v3 -= v3
    if (989 >= 764) {
        v6 -= v6
        v6 -= (v6 * (v5 ^ v5))
        if v1 {
            v9 = v9
            v4 = ((v4 + ("a" + v4)) + (v4 + v4))
            v6 = 2
        }
    }
    v7 -= -399
    switch "\n\t" {
    case "a":
        switch v4 {
        case "":
            v12 := GenS1{F0: v4, F1: 545}
            _ = v12
        }
    case "hello":
        var v13 bool = false
        _ = v13
    case "":
        if v1 {
            if (1000 < -958) {
                v14 := GenS2{F0: (v2 % (((v2 - (^(v2)))) | 1))}
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 bool = ((v7 == 1) && v1)
                    v17 := GenS3{F0: 42, F1: (v7 & -691), F2: 1}
                    v7 += (-(fi15))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: v2}
                _ = v14
                _ = v18
            } else {
                if (!v1) {
                    v19 := GenS1{F0: "foo bar", F1: 42}
                    _ = v19
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
