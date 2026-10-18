package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 0
    var v2 int = 434
    v3 := []string{"a"}
    var v4 string = "foo bar"
    if len(v3) > 0 {
        v3[0] = ("a" + ("" + ((v4 + "") + v4)))
    }
    if len(v3) > 0 {
        v3[0] = v4
    }
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int64 = 1
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: 127}
    switch "foo bar" {
    case "a":
        v2 = v2
    case "hello":
        v2 = -9223372036854775808
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int8 = -11
                fmt.Println(fi9, v11)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    fmt.Println(v2)
    switch v2 {
    case 48:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        fmt.Println(v4, v1)
    }
    var v12 bool = false
    v13 := GenS2{F0: -1}
    v12 = false
    v12 = (179 == 255)
    fmt.Println(v2, v1)
    var v14 float64 = (0.1 + 0.0)
    if v12 {
        fmt.Println(v12)
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int8 = 0
                _ = fi15
                _ = v16
            }
            var v17 float64 = (v14 - v14)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 uint = 0
                _ = fi18
                _ = v19
            }
        }
    }
    if true {
        fmt.Println(v12, v2)
        v2 = v2
    }
    v12 = (4294967295 == 18446744073709551615)
    var v20 int64 = -270
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
