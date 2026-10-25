package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint }

func main() {
    var v1 int16 = -581
    var v2 uint16 = 65535
    var v3 float32 = (((-2.25 - float32(0.0)) - float32((-2.25 + (0.1 - 0.0)))) / (float32((float64(float32(((-2.25 - 1e10) / 3.14159))) / 1.0)) + -2.25))
    var v4 string = ("\n\t" + "\n\t")
    var v5 int = 0
    var v6 string = "foo bar"
    var v7 float64 = (0.1 + ((1.5 + (1.5 + 1e10)) * float64(0.0)))
    v4 = ("\n\t" + v6)
    switch v6 {
    case "foo bar":
        switch v5 {
        case 42:
            v7 -= v7
        case -1:
            var v8 uint = 343
            _ = v8
        default:
            v9 := GenS1{F0: v4, F1: 1, F2: 2}
            _ = v9
        }
    }
    v10 := GenS1{F0: ("a" + "hello"), F1: 100, F2: 712}
    var v11 uint64 = 34
    var v12 uint16 = (^(v2))
    v13 := []int{32768}
    if true {
        v12 = (v2 + v12)
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        var v14 int8 = 127
        _ = v14
    }
    fmt.Println(v6)
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            if false {
                _ = len(v13)
            }
            switch (^(v5)) {
            case 324:
                if len(v13) > 0 {
                    v13[0] = (v5 - 0)
                }
            case -1:
                switch (^(v5)) {
                case -149:
                    if len(v13) > 0 {
                        v13[0] = (fi15 ^ v5)
                    }
                case -450:
                    v1 = v1
                case 127:
                    v5 = (-(fi15))
                }
            default:
                v4 = (v4 + ((("hello" + "") + "") + v4))
            }
            _ = fi16
        }
        _ = fi15
    }
    var v17 int64 = 42
    for fi18 := 0; fi18 < 10; fi18++ {
        v13 = append(v13, fi18)
        v19 := GenS2{F0: 346, F1: 34}
        _ = fi18
        _ = v19
    }
    if ((false && (!true)) && (false && (v5 == v5))) {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
