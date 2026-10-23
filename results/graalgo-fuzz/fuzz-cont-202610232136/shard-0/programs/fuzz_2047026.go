package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32) string {
    var v6 int8 = -1
    switch "\n\t" {
    case "hello":
        v6 = (v6 - v6)
    default:
        fmt.Println(v6, v6)
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: "a"}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = 42
        _ = fi9
        _ = v10
    }
    var v11 int64 = 42
    _ = p0
    _ = v11
    _ = v6
    return ("" + (("" + "\n\t") + ("a" + ("" + ("" + "hello")))))
}

func main() {
    v1 := []string{"", (("foo bar" + "\n\t") + (("\n\t" + (("\n\t" + "hello") + ("hello" + "a"))) + "\n\t")), "hello"}
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := make(map[string]string)
    var v5 int32 = 42
    v5 = v5
    for fi12 := 0; fi12 < 15; fi12++ {
        v2[2] = ""
        switch (fi12 | (fi12 % ((fi12) | 1))) {
        case 333:
            if false {
                v5 = (v5 << ((14) & 63))
            }
        }
        _ = fi12
    }
    var v13 int32 = (v5 << ((255) & 63))
    var v14 uint32 = 2147483648
    if (3.14159 != 0.1) {
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS2{F0: -9223372036854775808, F1: (v14 & (v14 ^ 822))}
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v13, v14, v13)
        v5 = v5
        var v17 uint64 = 1
        _ = v17
    }
    v5 -= 0
    switch 479 {
    case 520:
        switch -399 {
        case 0:
            switch 871 {
            case -1:
                for fi18 := 0; fi18 < 3; fi18++ {
                    v13 = 56
                    fmt.Println(v5, v13)
                    _ = fi18
                }
            }
        case 185:
            for fi19 := 0; fi19 < 8; fi19++ {
                switch "foo bar" {
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    var v20 float64 = 0.1
                    _ = v20
                }
                _ = fi19
            }
        default:
            switch "foo bar" {
            case "a":
                var v21 int = -852
                _ = v21
            case "\n\t":
                var v22 uint32 = 1
                _ = v22
            }
        }
    }
    v14 += 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
