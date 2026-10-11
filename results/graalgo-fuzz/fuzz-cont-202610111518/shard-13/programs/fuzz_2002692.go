package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 uint32) uint16 {
    v5 := []int{-2147483648, 477}
    if false {
        var v6 uint = 2
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 2
    var v3 float32 = 1.5
    var v4 int = -614
    fmt.Println(v4, v4)
    v7 := []int{-1, v4}
    if true {
        v3 -= float32(1.5)
        _ = len(v7)
        switch (v4 / ((v4) | 1)) {
        case -612:
            if true {
                var v8 uint16 = v1
                _ = v8
            }
        case 2147483647:
            v4 = (v4 | v4)
        case 0:
            switch (^(v4)) {
            case -9223372036854775808:
                v7 = append(v7, v4)
            }
        }
    } else {
        if true {
            v2 -= v2
        }
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int = 127
        _ = fi9
        _ = v10
    }
    var v11 float64 = ((3.14159 + ((1.5 + 1e10) + (0.1 + -2.25))) * (1.5 * ((3.14159 / (3.14159 - 0.0)) * -2.25)))
    var v12 float32 = -2.25
    fmt.Println(v2, v1)
    var v13 float64 = v11
    fmt.Println(v11, v11, v13)
    if ((v3 - (v12 - (v12 * float32((v13 - ((v11 * 1e10) * (v11 + float64(v3)))))))) > v3) {
        v4 = 0
    }
    v14 := GenS1{F0: 1, F1: 128, F2: 42}
    if (((1.5 / 1e10) * v3) == 0.0) {
        v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
        _ = len(v7)
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                v13 = -2.25
                var v17 uint32 = (v2 & v2)
                v7 = append(v7, (v4 >> ((255) & 63)))
                _ = fi16
                _ = v17
            }
            var v18 string = (("" + "foo bar") + ("hello" + ("\n\t" + ("foo bar" + "\n\t"))))
            _ = fi15
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
