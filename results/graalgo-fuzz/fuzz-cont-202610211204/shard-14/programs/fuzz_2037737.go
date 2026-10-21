package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }

func main() {
    var v1 uint16 = 198
    var v2 float64 = 1.5
    var v3 int64 = 1
    var v4 int = 127
    fmt.Println(v4, v1, v3)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: ("hello" + ""), F1: (v4 * (fi5 % (((v4 ^ fi5)) | 1))), F2: (fi5 | v4)}
        fmt.Println(v2)
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    v4 += 127
    v4 += 9223372036854775807
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if ("a" > ("foo bar" + "a")) {
        var v8 int16 = 1
        v4 += v4
        _ = v8
    }
    v1 = (v1 + (v1 - 32767))
    if false {
        var v9 string = ("" + "foo bar")
        if false {
            switch v9 {
            case "a":
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch v9 {
                    case "\n\t":
                        for fi11 := 0; fi11 < 3; fi11++ {
                            var v12 uint8 = 54
                            fmt.Println(fi10)
                            _ = fi11
                            _ = v12
                        }
                    default:
                        var v13 uint16 = 154
                        _ = v13
                    }
                    _ = fi10
                }
            default:
                var v14 int16 = 32767
                _ = v14
            }
            fmt.Println(v1)
            v1 = (v1 / ((v1) | 1))
        } else {
            var v15 uint16 = 829
            v9 = v9
            if len(v7) > 0 {
                v7[0] = v4
            }
            _ = v15
        }
        _ = v9
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
