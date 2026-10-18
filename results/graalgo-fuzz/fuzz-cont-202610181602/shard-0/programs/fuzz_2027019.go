package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = ""
    v2 := []int{-1, 127}
    v3 := make(map[uint32]uint32)
    var v4 int32 = -1
    v1 = "foo bar"
    fmt.Println(v4)
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v1)
        v4 += 1
        switch 2147483647 {
        case 127:
            v4 += (v4 & -2147483648)
        case 366:
            var v6 int32 = 42
            _ = v6
        case -512:
            delete(v3, 65535)
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = (1e10 - 1e10)
        var v9 float64 = 1.5
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 62}
    var v11 bool = true
    v12 := GenS1{F0: 100}
    switch v1 {
    case "":
        for fi13 := 0; fi13 < 1; fi13++ {
            switch (v1 + "\n\t") {
            case "hello":
                if v11 {
                    delete(v3, 994)
                    v14 := GenS1{F0: 3}
                    if (42 == 42) {
                        fmt.Println(v1, fi13)
                        v4 += (-(v4))
                    } else {
                        var v15 int8 = 1
                        _ = v15
                    }
                    _ = v14
                }
            case "\n\t":
                var v16 string = v1
                _ = v16
            default:
                if ((v11 && true) || v11) {
                    _ = len(v2)
                }
            }
            _ = fi13
        }
    case "hello":
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int32 = v4
            var v19 uint8 = 127
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v1, v1, v11)
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
