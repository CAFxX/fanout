package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int64 = 1
    var v2 string = ""
    v3 := []uint32{298, 42}
    var v4 int64 = 467
    var v5 int64 = 1
    var v6 int64 = v5
    v2 = "hello"
    _ = len(v3)
    if (false || (!(((-98 < 100) || true) && ((false && false) && (!(false || (1e10 != (1e10 / (((-2.25 * 3.14159) + (1.5 - float32(0.1))) - 1e10))))))))) {
        var v7 uint8 = 0
        _ = v7
    }
    v2 = ("\n\t" + v2)
    v1 = (v5 << ((65535) & 63))
    switch v2 {
    case "foo bar":
        var v8 uint = 582
        _ = v8
    case "":
        if (!(!true)) {
            if len(v3) > 0 {
                v3[0] = 65535
            }
            switch v2 {
            case "\n\t":
                if (1 > 255) {
                    fmt.Println(v1, v5)
                    fmt.Println(v5, v1)
                    v6 += v5
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case "a":
                switch 128 {
                case -1:
                    fmt.Println(v6, v5)
                }
            }
        }
    }
    var v9 int32 = -721
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint16 = 2
        _ = fi10
        _ = v11
    }
    switch 885 {
    case 0:
        fmt.Println(v6, v2)
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            v3 = append(v3, 2)
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int32 = v9
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v5, v6)
        _ = fi15
    }
    var v16 float64 = 0.0
    var v17 int64 = v6
    switch v2 {
    case "foo bar":
        v18 := GenS1{F0: false}
        _ = v18
    default:
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS1{F0: false}
            fmt.Println(v16, v2, fi19)
            _ = fi19
            _ = v20
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            if (((("foo bar" + v2) + ((v2 + v2) + ("" + "foo bar"))) >= ("a" + (v2 + v2))) && true) {
                v23 := GenS2{F0: v17}
                var v24 int32 = -1
                _ = v23
                _ = v24
            } else {
                fmt.Println(v1, fi21)
                var v25 uint = 1
                v9 = 0
                _ = v25
            }
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
