package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 int8 = -88
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    var v4 int = -2147483648
    v5 := GenS1{F0: (v2 + ""), F1: 356}
    v6 := []int{}
    switch 4294967295 {
    case 960:
        var v7 uint = 42
        _ = v7
    default:
        fmt.Println(v4, v4)
    }
    if false {
        switch v2 {
        case "a":
            _ = len(v3)
        case "":
            for fi8 := 0; fi8 < 3; fi8++ {
                if (!true) {
                    v2 = ((v2 + v2) + v2)
                    var v9 uint = 2
                    _ = v9
                }
                _ = fi8
            }
        case "foo bar":
            var v10 int16 = 0
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v12 uint32 = 2147483647
            for fi13 := 0; fi13 < 9; fi13++ {
                if ((("a" + ("foo bar" + v2)) + ((v2 + v2) + (v2 + ((v2 + v2) + v2)))) > "hello") {
                    var v14 int = 9223372036854775807
                    _ = v14
                }
                if (v12 < (^(v12))) {
                    if true {
                        var v15 float64 = (1.5 / 3.14159)
                        var v16 uint64 = 42
                        var v17 int16 = int16(v16)
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS1{F0: (v2 + "foo bar"), F1: 895}
                    _ = v18
                }
                if (!(!(0 < 0))) {
                    v19 := v3[(v4 >> ((255) & 63))]
                    var v20 int64 = -1
                    fmt.Println(fi13, v12)
                    v21 := GenS1{F0: "hello", F1: -32768}
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        var v22 string = "a"
        _ = v22
    } else {
        fmt.Println(v2, v4)
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS1{F0: (v2 + "foo bar"), F1: -1}
            _ = fi23
            _ = v24
        }
    }
    var v25 bool = false
    fmt.Println(v25, v1, v1)
    v25 = (((189 < 2) || v25) || v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
