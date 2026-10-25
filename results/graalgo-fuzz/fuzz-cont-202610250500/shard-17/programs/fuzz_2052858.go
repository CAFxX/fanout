package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 int8 = 59
    var v3 uint8 = 127
    var v4 float32 = 1e10
    v5 := []string{(v1 + v1), ((("\n\t" + (v1 + "foo bar")) + "") + v1)}
    var v6 int64 = 465
    var v7 int8 = v2
    if false {
        v8 := GenS1{F0: (uint64(v3) >= 42)}
        var v9 uint16 = 89
        if true {
            var v10 string = v1
            if true {
                v2 = (-(v7))
                switch -848 {
                case 32768:
                    var v11 int8 = 100
                    _ = v11
                }
                var v12 int8 = -69
                _ = v12
            } else {
                v3 = 1
                var v13 int8 = v7
                switch 460 {
                case 2147483647:
                    v14 := GenS1{F0: false}
                    _ = v14
                }
                _ = v13
            }
            v6 = v6
            _ = v10
        } else {
            v6 += v6
            switch ("\n\t" + ("\n\t" + v1)) {
            case "\n\t":
                v15 := GenS1{F0: false}
                _ = v15
            case "hello":
                fmt.Println(v1, v6, v7)
            case "a":
                for fi16 := 0; fi16 < 16; fi16++ {
                    fmt.Println(v9, v3)
                    var v17 int16 = 32767
                    _ = fi16
                    _ = v17
                }
            default:
                v7 -= v2
            }
            fmt.Println(v3, v7)
        }
        _ = v8
        _ = v9
    } else {
        var v18 string = "\n\t"
        _ = v18
    }
    var v19 bool = false
    for fi20 := 0; fi20 < 12; fi20++ {
        fmt.Println(v1, v4, v19)
        _ = fi20
    }
    switch 32767 {
    case 647:
        fmt.Println(v2, v19, v1)
    case -2147483648:
        v4 = (((float32(0.1) / 0.1) + v4) - 1.5)
    case 345:
        if false {
            v7 -= 42
            switch -2147483648 {
            case 42:
                if (v3 != v3) {
                    v4 += v4
                }
            case 2147483647:
                v5 = append(v5, v1)
            case -579:
                v6 = v6
            default:
                if len(v5) > 0 {
                    v5[0] = ("foo bar" + v1)
                }
            }
        }
    }
    var v21 uint8 = (v3 | 127)
    v3 = v3
    v1 = "hello"
    var v22 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
