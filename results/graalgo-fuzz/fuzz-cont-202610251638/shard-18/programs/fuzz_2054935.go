package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }

func main() {
    var v1 int16 = -445
    var v2 int = 42
    var v3 int64 = -9223372036854775808
    var v4 int = 32767
    var v5 int = v2
    v6 := []string{}
    var v7 string = ("hello" + "\n\t")
    var v8 string = ("\n\t" + v7)
    switch (v4 >> ((0) & 63)) {
    case 127:
        _ = len(v6)
    }
    v9 := GenS1{F0: true, F1: v3, F2: v5}
    v3 -= 611
    fmt.Println(v3)
    v7 = v8
    var v10 bool = (42 < 671)
    _ = len(v6)
    v11 := GenS1{F0: v10, F1: v3, F2: (v4 % ((v4) | 1))}
    var v12 uint64 = 906
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v13 uint = 71
    var v14 float64 = (-2.25 - 0.1)
    var v15 int = (v4 * (v4 << ((0) & 63)))
    if (((1.5 + 0.1) > (3.14159 / float64(1.5))) && (v10 || (v12 >= v12))) {
        for fi16 := 0; fi16 < 13; fi16++ {
            switch v4 {
            case 127:
                var v17 float64 = v14
                _ = v17
            default:
                _ = len(v6)
            }
            _ = fi16
        }
        var v18 int64 = -1
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
