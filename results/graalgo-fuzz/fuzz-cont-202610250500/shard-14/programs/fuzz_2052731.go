package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int = 128
    var v2 uint16 = 42
    var v3 float64 = -2.25
    v4 := []string{("" + "\n\t"), "\n\t"}
    v5 := []string{("" + ("foo bar" + ("\n\t" + "foo bar"))), "a", "a", ("" + "hello")}
    var v6 string = "hello"
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint8 = 255
            switch v6 {
            case "foo bar":
                v2 = 690
            case "":
                _ = len(v4)
            default:
                switch fi8 {
                case 128:
                    v5 = append(v5, (v6 + "hello"))
                }
            }
            _ = fi8
            _ = v9
        }
        v3 = 1.5
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = fi7
    }
    v2 = (v2 << ((uint(v2)) & 63))
    v1 = (v1 & v1)
    var v10 string = v6
    if false {
        var v11 float32 = 1e10
        v12 := GenS1{F0: 255}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v11
        _ = v12
    } else {
        v2 -= 724
        v1 = v1
    }
    if ((3.14159 / (0.0 / (-2.25 - (1.5 + float32(v3))))) < 0.1) {
        if len(v5) > 0 {
            v5[0] = v10
        }
    } else {
        switch (v1 + -2147483648) {
        case 42:
            var v13 bool = (!(false && true))
            _ = v13
        case 273:
            v6 = "foo bar"
        }
        fmt.Println(v10, v2, v1)
        var v14 uint32 = 2
        _ = v14
    }
    v6 = (v10 + v10)
    var v15 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
