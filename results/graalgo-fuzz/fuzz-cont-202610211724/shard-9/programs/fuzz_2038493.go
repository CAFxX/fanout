package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 4294967295
    var v2 float64 = float64(0.1)
    var v3 int = -1
    var v4 string = "foo bar"
    v5 := GenS1{F0: (v3 - v3)}
    fmt.Println(v3, v2, v2)
    fmt.Println(v2, v1)
    var v6 int = (v3 & v3)
    v7 := []int{v3, (v3 >> ((980) & 63)), (v3 / ((42) | 1)), (v6 >> ((2) & 63))}
    var v8 uint32 = v1
    v3 -= v6
    if false {
        var v9 bool = (int64(v1) != -1)
        switch (v4 + (((v4 + "foo bar") + v4) + v4)) {
        case "":
            if v9 {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v8 -= (^(v8))
                    v9 = false
                    _ = fi10
                }
                var v11 uint64 = 66
                _ = v11
            }
        case "\n\t":
            v4 = (("hello" + "hello") + "\n\t")
        default:
            switch v6 {
            case -484:
                switch (((("hello" + (v4 + (v4 + (("a" + (v4 + ("a" + v4))) + (v4 + v4))))) + (v4 + (v4 + "foo bar"))) + "hello") + v4) {
                case "foo bar":
                    var v12 uint8 = 127
                    _ = v12
                case "hello":
                    v2 = 1e10
                case "a":
                    var v13 int16 = -765
                    _ = v13
                default:
                    var v14 float64 = 3.14159
                    _ = v14
                }
            default:
                for fi15 := 0; fi15 < 11; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if len(v7) > 0 {
                            v7[0] = (fi15 - (fi16 & (v6 / ((fi16) | 1))))
                        }
                        _ = fi16
                    }
                    _ = fi15
                }
            }
        }
        _ = v9
    }
    var v17 float64 = (0.1 * float64(3.14159))
    var v18 float32 = (((0.1 / 0.1) - (3.14159 / 1e10)) / -2.25)
    v6 = (v3 >> ((65535) & 63))
    v7 = append(v7, (v6 ^ v3))
    v2 -= float64(v18)
    switch v3 {
    case -2147483648:
        if false {
            var v19 float64 = v17
            v3 -= v3
            v18 = v18
            _ = v19
        } else {
            var v20 uint8 = 255
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int64 = -593
            _ = fi22
            _ = v23
        }
        v24 := GenS1{F0: 373}
        var v25 uint = 110
        _ = fi21
        _ = v24
        _ = v25
    }
    if len(v7) > 0 {
        v7[0] = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
