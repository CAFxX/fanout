package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int64 {
    switch "\n\t" {
    case "a":
        v6 := GenS1{F0: false}
        _ = v6
    case "\n\t":
        var v7 uint32 = 65535
        _ = v7
    case "foo bar":
        v8 := []int{}
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if (-27 == 38) {
            v10 := []int{fi9, fi9, (fi9 << ((805) & 63)), (^(fi9))}
            _ = v10
        }
        _ = fi9
    }
    v11 := GenS1{F0: true}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint32 = 2147483647
        _ = fi12
        _ = v13
    }
    _ = v11
    return 42
}

func main() {
    var v1 float32 = ((0.0 / -2.25) / float32((0.1 / 1.5)))
    var v2 int = 128
    var v3 float64 = 0.1
    var v4 bool = false
    v5 := []uint32{}
    var v14 int64 = 9223372036854775807
    var v15 int8 = -128
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint = 564
        if false {
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS1{F0: true}
                v20 := GenS1{F0: (1 >= 128)}
                _ = fi18
                _ = v19
                _ = v20
            }
            switch -9223372036854775808 {
            case 2147483647:
                var v21 float64 = v3
                _ = v21
            }
        } else {
            var v22 uint16 = 42
            _ = v22
        }
        var v23 float64 = v3
        _ = fi16
        _ = v17
        _ = v23
    }
    v24 := GenS1{F0: v4}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v15, v14)
    if v4 {
        switch "\n\t" {
        case "\n\t":
            v2 = (^(v2))
        case "":
            var v25 int = v2
            _ = v25
        }
    }
    var v26 uint8 = 128
    v14 = 9223372036854775807
    for fi27 := 0; fi27 < 6; fi27++ {
        v3 -= ((v3 * ((1.5 * 0.1) - (0.1 * float64(-2.25)))) + v3)
        if v4 {
            v3 -= (-2.25 + 0.0)
            fmt.Println(v26)
        } else {
            var v28 float64 = 1.5
            _ = v28
        }
        v5 = append(v5, 2147483648)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
