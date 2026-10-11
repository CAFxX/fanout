package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint64 = 813
    var v2 string = ("foo bar" + "hello")
    v3 := []string{v2, (v2 + "hello"), ((v2 + (("foo bar" + v2) + "")) + (("a" + v2) + "\n\t"))}
    var v4 bool = false
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 127, F1: (0 != -2147483648)}
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = v1
        var v9 uint64 = v1
        _ = fi8
        _ = v9
    }
    var v10 bool = false
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint8 = 106
        if (65535 > uint(v12)) {
            var v13 int16 = -89
            v2 = "a"
            _ = v13
        } else {
            fmt.Println(v4, v10, v4)
            v2 = "hello"
        }
        _ = len(v5)
        _ = fi11
        _ = v12
    }
    var v14 uint16 = 42
    fmt.Println(v4, v2, v1)
    var v15 uint8 = 138
    v15 += 100
    v4 = (!(!(false && (0 > 0))))
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v10)
        if ((fi16 * (fi16 << ((349) & 63))) < (fi16 & (fi16 << ((255) & 63)))) {
            var v17 int = (^(fi16))
            _ = v17
        } else {
            var v18 uint8 = v15
            switch (fi16 << ((1) & 63)) {
            case -2147483648:
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS2{F0: 42}
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                    _ = fi19
                    _ = v20
                }
            }
            _ = v18
        }
        v5[0] = v10
        _ = fi16
    }
    switch v2 {
    case "hello":
        var v21 int8 = 100
        _ = v21
    case "foo bar":
        if v4 {
            fmt.Println(v15, v2)
            var v22 uint = 0
            _ = v22
        }
    default:
        v7 += 1.5
    }
    var v23 uint32 = 185
    switch v2 {
    case "":
        var v24 uint16 = v14
        _ = v24
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v25 string = v2
    if v4 {
        if len(v3) > 0 {
            v3[0] = v2
        }
        v7 -= v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
