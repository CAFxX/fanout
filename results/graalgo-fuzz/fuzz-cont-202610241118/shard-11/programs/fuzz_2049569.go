package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64(0.0)
    v4 := make(map[uint32]bool)
    var v5 float64 = (v3 - v3)
    delete(v4, 114)
    var v6 bool = (1 != 2147483648)
    v5 = 1.5
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v5 = v5
            fmt.Println(fi8)
            _ = fi8
        }
        if ((!v6) && v6) {
            var v9 float64 = (1.5 + 3.14159)
            _ = v9
        } else {
            switch -2147483648 {
            case 9223372036854775807:
                v10 := []int{fi7, (fi7 / ((fi7) | 1)), (fi7 << ((2) & 63))}
                _ = v10
            case -774:
                if (!false) {
                    v3 = v3
                    v11 := GenS1{F0: 0, F1: -78}
                    _ = v11
                }
            case 32768:
                var v12 float64 = v5
                _ = v12
            default:
                switch (fi7 | fi7) {
                case 356:
                    var v13 string = "a"
                    _ = v13
                case -848:
                    v6 = false
                case -9223372036854775808:
                    v3 = v3
                default:
                    var v14 string = ("\n\t" + (("foo bar" + ("a" + "foo bar")) + "foo bar"))
                    _ = v14
                }
            }
        }
        v5 = v5
        _ = fi7
    }
    switch "a" {
    case "":
        v6 = (true || (592 < 0))
    default:
        v6 = (!v6)
    }
    fmt.Println(v5, v3)
    var v15 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
