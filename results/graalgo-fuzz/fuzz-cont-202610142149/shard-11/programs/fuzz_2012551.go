package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    var v1 int64 = 0
    var v2 int64 = 1
    var v3 int16 = -1
    var v4 uint8 = 0
    var v5 int8 = 51
    var v6 string = "foo bar"
    var v7 uint8 = 2
    var v8 int64 = v2
    if true {
        v9 := GenS1{F0: 275, F1: ("a" + "hello")}
        _ = v9
    }
    v10 := []int{}
    v10 = append(v10, 127)
    v2 -= (v1 % (((v8 / ((v2) | 1))) | 1))
    var v11 int8 = -77
    var v12 uint16 = 32767
    var v13 uint8 = 255
    var v14 int64 = (v2 + v2)
    v10 = append(v10, 32768)
    switch v6 {
    case "\n\t":
        if (!(!(0 == 0))) {
            switch ("" + v6) {
            case "foo bar":
                v15 := GenS1{F0: 669, F1: "a"}
                _ = v15
            }
            switch 42 {
            case 1:
                for fi16 := 0; fi16 < 13; fi16++ {
                    fmt.Println(v2, v12, v4)
                    _ = fi16
                }
            case 0:
                v7 -= (v4 ^ v13)
            }
        } else {
            switch v6 {
            case "":
                for fi17 := 0; fi17 < 2; fi17++ {
                    v10 = append(v10, fi17)
                    _ = fi17
                }
            case "\n\t":
                var v18 uint64 = 18446744073709551615
                _ = v18
            }
            fmt.Println(v11, v11, v8)
            if false {
                v2 = v8
                switch ("" + v6) {
                case "foo bar":
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                case "\n\t":
                    if false {
                        if (!true) {
                            v19 := GenS1{F0: 421, F1: (v6 + v6)}
                            _ = v19
                        }
                    }
                default:
                    var v20 int32 = -2147483648
                    _ = v20
                }
            }
        }
    case "hello":
        var v21 uint32 = 65535
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 float64 = 0.0
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
