package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    v4 := []uint32{2, (v3 + (v3 << ((648) & 63)))}
    v2[-1] = true
    v3 = (v3 << ((65535) & 63))
    switch 42 {
    case -324:
        v3 -= (v3 / ((v3) | 1))
    case 128:
        switch "\n\t" {
        case "":
            fmt.Println(v3)
        case "\n\t":
            fmt.Println(v3)
        default:
            var v5 uint32 = (^(v3))
            _ = v5
        }
    default:
        v1[9223372036854775807] = (127 < -128)
    }
    v6 := GenS1{F0: -1, F1: true, F2: 1}
    var v7 int16 = -732
    switch int(v7) {
    case -122:
        switch ("hello" + "") {
        case "a":
            v3 = (v3 * (v3 << ((0) & 63)))
        case "\n\t":
            for fi8 := 0; fi8 < 1; fi8++ {
                if (((false || (false || (false || false))) || false) || false) {
                    var v9 int8 = 0
                    v9 = 1
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v3 = (v3 | 42)
                        fmt.Println(v7, v3, v7)
                        _ = fi10
                    }
                    var v11 bool = true
                    _ = v11
                }
                v12 := GenS1{F0: 127, F1: false, F2: fi8}
                v13 := GenS1{F0: 87, F1: (true || (3.14159 == 3.14159)), F2: fi8}
                _ = fi8
                _ = v12
                _ = v13
            }
        case "foo bar":
            if (671 >= 4294967295) {
                var v14 string = ""
                var v15 float32 = float32(-2.25)
                _ = v14
                _ = v15
            }
        default:
            v4 = append(v4, (v3 << ((612) & 63)))
        }
    case 227:
        v7 = (v7 / ((v7) | 1))
    case 32768:
        if true {
            _ = len(v1)
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
    }
    switch 42 {
    case -2147483648:
        delete(v1, -384)
    case 127:
        var v16 uint32 = 0
        _ = v16
    }
    v17 := GenS2{F0: (1 <= (v7 + v7))}
    var v18 uint64 = 24
    fmt.Println(v3, v3, v3)
    var v19 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
