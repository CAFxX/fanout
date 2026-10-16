package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }

func main() {
    var v1 float64 = ((1.5 * 0.1) / (3.14159 * float64(3.14159)))
    v2 := []uint32{4294967295, 42, 1, 0}
    var v3 int64 = 42
    var v4 uint8 = 255
    switch 32767 {
    case 850:
        for fi5 := 0; fi5 < 2; fi5++ {
            v3 = v3
            v3 -= 125
            _ = fi5
        }
    }
    var v6 bool = false
    switch 42 {
    case 2147483647:
        for fi7 := 0; fi7 < 3; fi7++ {
            if (42 == 4294967295) {
                v8 := GenS1{F0: fi7}
                _ = v8
            } else {
                v3 = v3
                v3 -= (v3 >> ((0) & 63))
                fmt.Println(v1, v4, v4)
            }
            _ = fi7
        }
    case 4294967295:
        if (!(!((float32(v1) * (0.0 + 3.14159)) <= (0.1 * float32(v1))))) {
            if v6 {
                switch 352 {
                case 127:
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v4 = (v4 & v4)
                        _ = fi9
                    }
                }
                var v10 uint32 = 65535
                _ = v10
            } else {
                _ = len(v2)
                for fi11 := 0; fi11 < 9; fi11++ {
                    v2 = append(v2, 0)
                    fmt.Println(v6, v4)
                    switch ("" + ((("\n\t" + ((((("hello" + ("" + ("" + ""))) + "a") + "") + "") + "foo bar")) + "") + "")) {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        for fi12 := 0; fi12 < 7; fi12++ {
                            var v13 uint = 2
                            _ = fi12
                            _ = v13
                        }
                    }
                    _ = fi11
                }
            }
        }
    default:
        if (uint(v4) == 2) {
            v1 = v1
        }
    }
    v14 := GenS1{F0: 626}
    v6 = ((0.0 <= (((3.14159 / 1.0) * v1) - 1e10)) || ((!false) && ("" >= "foo bar")))
    fmt.Println(v1, v6, v4)
    var v15 uint = 2
    v16 := GenS2{F0: -995}
    var v17 int16 = -406
    fmt.Println(v3)
    if v6 {
        var v18 string = ""
        v19 := GenS3{F0: 0, F1: 0, F2: 42}
        var v20 uint = (^(v15))
        _ = v18
        _ = v19
        _ = v20
    }
    switch 127 {
    case 32767:
        v21 := GenS2{F0: v17}
        _ = v21
    default:
        v6 = (-885 <= -1)
    }
    v15 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
