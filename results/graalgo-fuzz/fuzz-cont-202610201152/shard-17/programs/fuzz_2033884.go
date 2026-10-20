package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{"hello", "", "hello", ""}
    var v3 uint32 = 537
    var v4 int16 = 1000
    v5 := []string{((("hello" + "foo bar") + "a") + ("a" + "foo bar")), "foo bar"}
    var v6 string = "a"
    v7 := GenS1{F0: 127}
    v8 := GenS1{F0: 255}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v9 uint16 = 1
    v9 = (v9 | (^(v9)))
    if false {
        v10 := GenS2{F0: (238 > int(v3))}
        for fi11 := 0; fi11 < 18; fi11++ {
            switch ((v6 + (v6 + v6)) + "") {
            case "":
                v6 = "a"
            case "foo bar":
                var v12 uint8 = 0
                _ = v12
            default:
                var v13 int64 = -1
                _ = v13
            }
            v4 = v4
            v4 = v4
            _ = fi11
        }
        if false {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int = fi14
                v9 = v9
                _ = fi14
                _ = v15
            }
            v16 := GenS1{F0: 128}
            _ = v16
        }
        _ = v10
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint8 = 166
            var v19 uint16 = v9
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v20 := GenS3{F0: (v3 >> ((v1) & 63)), F1: -2147483648}
    if len(v2) > 0 {
        v2[0] = v6
    }
    var v21 int8 = int8(v4)
    var v22 int16 = 32767
    fmt.Println(v3, v4)
    switch 2147483647 {
    case 2147483647:
        var v23 float32 = (0.1 * -2.25)
        _ = v23
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = ("" + "a")
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v6 = v6
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
