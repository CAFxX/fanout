package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
func genFn0() int32 {
    var v6 uint16 = 32767
    var v7 int = 9223372036854775807
    fmt.Println(v6, v7)
    fmt.Println(v6, v7)
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint32 = 0
    var v2 string = ("\n\t" + ((("\n\t" + "\n\t") + ("" + ("" + ("hello" + "")))) + ("" + "")))
    var v3 uint16 = 255
    var v4 uint16 = 0
    v5 := make(map[uint32]int)
    v8 := GenS1{F0: 707, F1: -614, F2: int(v1)}
    v3 -= 566
    v4 += (v3 ^ (v4 / ((0) | 1)))
    var v9 uint64 = uint64(v1)
    if false {
        switch -2147483648 {
        case 9223372036854775807:
            switch ("a" + v2) {
            case "\n\t":
                switch v2 {
                case "a":
                    var v10 uint64 = v9
                    _ = v10
                }
            case "a":
                var v11 bool = (!(1000 != 577))
                _ = v11
            case "foo bar":
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS1{F0: 0, F1: (fi12 / ((-1) | 1)), F2: (^(fi12))}
                    _ = fi12
                    _ = v13
                }
            }
        default:
            v3 -= v4
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v2)
        v4 -= 0
        var v15 int = 9223372036854775807
        _ = fi14
        _ = v15
    }
    var v16 uint = 1
    v4 = (v4 & v4)
    var v17 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
