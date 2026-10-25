package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint) int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{}
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 string = "a"
            v7 = v7
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    if (510 != 1) {
        p0 += (p0 % ((143) | 1))
        for fi8 := 0; fi8 < 19; fi8++ {
            if (1 < 258) {
                v9 := GenS1{F0: -32768, F1: 770, F2: (fi8 + 4294967295)}
                _ = v9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int16 = 0
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 9; fi12++ {
                p0 = (^(p0))
                p0 += 761
                fmt.Println(p0)
                _ = fi12
            }
            _ = fi8
        }
    } else {
        p0 += 2
        switch 1 {
        case 302:
            v13 := []int{42, 127, 81, 0}
            _ = v13
        }
    }
    _ = p0
    return 353
}

func main() {
    var v1 int64 = -824
    var v2 int8 = 0
    var v3 uint64 = 2
    var v14 int16 = -1
    v15 := []int{128, -970}
    for fi16 := 0; fi16 < 20; fi16++ {
        v15 = append(v15, (fi16 << ((4294967295) & 63)))
        _ = fi16
    }
    v17 := GenS2{F0: 128}
    switch (("" + ("a" + "a")) + "\n\t") {
    case "":
        if len(v15) > 0 {
            v15[0] = 177
        }
    case "hello":
        fmt.Println(v2)
    case "foo bar":
        var v18 int8 = 127
        _ = v18
    default:
        v3 = 952
    }
    fmt.Println(v3, v1)
    var v19 uint8 = 0
    v2 -= 100
    var v20 uint64 = v3
    v20 -= (^(v3))
    var v21 float64 = 0.0
    var v22 float32 = 3.14159
    fmt.Println(v1)
    fmt.Println(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
