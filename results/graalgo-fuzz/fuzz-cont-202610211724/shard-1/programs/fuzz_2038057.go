package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -933
    var v3 uint64 = 939
    var v4 int64 = 1
    var v5 float32 = v1
    var v6 string = "hello"
    var v7 uint8 = 128
    v3 = v3
    switch (("" + "") + "") {
    case "foo bar":
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(v7)
                if true {
                    fmt.Println(fi8, v2, v5)
                    var v10 string = v6
                    v6 = v10
                    fmt.Println(v1)
                    v11 := GenS1{F0: "\n\t", F1: (false || true)}
                    _ = v10
                    _ = v11
                } else {
                    var v12 int32 = 1
                    var v13 int = fi8
                    v12 += v12
                    v14 := GenS1{F0: v6, F1: false}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
            _ = fi8
        }
    case "a":
        if (!(!((v5 / (v5 * v5)) >= v1))) {
            var v15 uint16 = 375
            _ = v15
        } else {
            v2 = v2
            v16 := []int{-1}
            _ = v16
        }
    }
    var v17 uint8 = uint8(v3)
    for fi18 := 0; fi18 < 19; fi18++ {
        v1 += (((1.5 - v1) * v5) - ((((v1 + (v1 + v5)) * ((v1 - 1.5) - (v1 - v5))) - -2.25) / ((v5 * 3.14159) - (v5 - (1e10 / ((v5 - v1) / float32(0.1)))))))
        var v19 int16 = (-(v2))
        var v20 int32 = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    v5 += (v1 - v5)
    v21 := GenS1{F0: v6, F1: true}
    v22 := GenS1{F0: v6, F1: (255 != 1)}
    var v23 int8 = 1
    fmt.Println(v3, v4, v23)
    var v24 int32 = 42
    v25 := GenS1{F0: v6, F1: (-2147483648 == 128)}
    fmt.Println(v17, v3, v23)
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS1{F0: v6, F1: (true || (fi26 <= (fi27 & -717)))}
            fmt.Println(v2)
            var v29 bool = ("\n\t" >= ("hello" + v6))
            _ = fi27
            _ = v28
            _ = v29
        }
        v1 += v1
        _ = fi26
    }
    var v30 uint64 = (v3 % ((v3) | 1))
    fmt.Println(v17, v3)
    v31 := GenS1{F0: v6, F1: (false && ((v24 >> ((255) & 63)) >= v24))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
