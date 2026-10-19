package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 int = -1
    var v3 uint16 = 0
    var v4 bool = false
    var v5 uint16 = v3
    var v6 int32 = 189
    if ("hello" >= "a") {
        v4 = (42 >= 100)
    }
    fmt.Println(v4, v5)
    fmt.Println(v4)
    var v7 int8 = -79
    if v4 {
        if v4 {
            switch v2 {
            case 1:
                var v8 bool = ((1 >= 690) && (-1 <= int64(v5)))
                _ = v8
            default:
                v9 := GenS1{F0: -1}
                _ = v9
            }
        } else {
            v10 := []int{}
            _ = v10
        }
    }
    v7 -= v7
    var v11 int16 = 32767
    fmt.Println(v1)
    var v12 int = int(v6)
    v13 := []int{v2, (v12 ^ v2)}
    fmt.Println(v5, v6, v12)
    if (0 != 351) {
        fmt.Println(v3, v2)
        if (v12 == 9223372036854775807) {
            if (-332 > (v11 | v11)) {
                fmt.Println(v5)
                var v14 int64 = -9223372036854775808
                _ = v14
            }
            v4 = (v4 && (867 < 1))
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            if v4 {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v12, v4, v11)
                case "hello":
                    v16 := GenS2{F0: 2}
                    _ = v16
                case "foo bar":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        switch "a" {
                        case "foo bar":
                            var v18 int8 = (v7 | v7)
                            _ = v18
                        }
                        _ = fi17
                    }
                default:
                    var v19 string = "\n\t"
                    _ = v19
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    v13 = append(v13, v12)
                    _ = fi20
                }
                v13 = append(v13, v2)
            } else {
                fmt.Println(v7, v6, v5)
            }
            var v21 bool = ((1 == 42) && false)
            _ = fi15
            _ = v21
        }
    } else {
        var v22 int64 = 1
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
