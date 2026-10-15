package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 int16 = 1
    v2 += 0
    v2 = (v2 * (v2 & v2))
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: "hello", F1: 410, F2: 4294967295}
        if ("\n\t" > "") {
            v8 := GenS2{F0: 1}
            v9 := GenS2{F0: 42}
            _ = v8
            _ = v9
        } else {
            var v10 string = "\n\t"
            switch (v10 + v10) {
            case "\n\t":
                if true {
                    switch (fi6 >> ((2) & 63)) {
                    case -425:
                        var v11 uint32 = 840
                        _ = v11
                    }
                    delete(v3, 653)
                } else {
                    v12 := v1[v10]
                    v13 := []int{fi6}
                    _ = v12
                    _ = v13
                }
            }
            _ = v10
        }
        switch "hello" {
        case "a":
            v2 = v2
        case "hello":
            v2 -= (v2 % ((-1) | 1))
        default:
            var v14 string = "foo bar"
            _ = v14
        }
        _ = fi6
        _ = v7
    }
    delete(v3, 907)
    _ = len(v3)
    for fi15 := 0; fi15 < 14; fi15++ {
        delete(v1, ("foo bar" + "hello"))
        fmt.Println(v2)
        _ = fi15
    }
    v16 := GenS2{F0: -9223372036854775808}
    v17 := GenS2{F0: 0}
    v5 -= v5
    v18 := []int{878, 0, -176}
    for fi19 := 0; fi19 < 11; fi19++ {
        if len(v18) > 0 {
            v18[0] = (fi19 & -9223372036854775808)
        }
        v5 = -32768
        for fi20 := 0; fi20 < 10; fi20++ {
            v18 = append(v18, (fi19 * fi20))
            fmt.Println(v2, fi19, v5)
            _ = fi20
        }
        _ = fi19
    }
    var v21 uint16 = 32767
    var v22 int8 = 127
    v2 -= (v2 % (((^(v2))) | 1))
    v4[4294967295] = (!(v21 == v21))
    v23 := GenS2{F0: 888}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
