package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 int64) int16 {
    var v6 int8 = 98
    fmt.Println(p1, p1, p0)
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(p1, v6, v6)
        _ = fi7
    }
    var v8 float32 = (3.14159 + (1.5 / ((-2.25 - 3.14159) + -2.25)))
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 389
    var v2 string = (("a" + "foo bar") + "")
    var v3 uint16 = 42
    v4 := []string{""}
    var v5 float64 = float64(1e10)
    if ((v3 >= 873) && ((v2 <= (v2 + "hello")) && true)) {
        switch 1 {
        case 32767:
            v9 := GenS1{F0: 448}
            _ = v9
        case 52:
            v3 -= (v3 + (^(v3)))
        default:
            v4 = append(v4, v2)
        }
        v10 := GenS1{F0: v3}
        var v11 int32 = 382
        _ = v10
        _ = v11
    } else {
        var v12 int16 = -1
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = v12
    }
    v3 = 32767
    v13 := GenS1{F0: (v3 - (v3 | (v3 + v3)))}
    switch 9223372036854775807 {
    case 32767:
        switch (v2 + v2) {
        case "\n\t":
            v5 = (0.0 / 0.1)
        case "foo bar":
            v3 -= (v3 & v3)
        case "":
            var v14 uint64 = 2
            _ = v14
        default:
            v4 = append(v4, ((v2 + "foo bar") + ((v2 + "foo bar") + ("\n\t" + v2))))
        }
    case 0:
        fmt.Println(v2, v5, v3)
    case 42:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if len(v4) > 0 {
            v4[0] = v2
        }
    }
    v2 = v2
    if true {
        fmt.Println(v3, v3, v1)
        if (v3 < (v3 * v3)) {
            v4 = append(v4, (v2 + ((v2 + "\n\t") + (((("" + v2) + v2) + (v2 + v2)) + v2))))
            for fi15 := 0; fi15 < 5; fi15++ {
                v3 += (v3 << ((255) & 63))
                if (!(!(!(!(!false))))) {
                    var v16 float64 = ((1.5 - 1.5) / v5)
                    _ = v16
                }
                _ = fi15
            }
            if (32767 > 0) {
                var v17 uint32 = (v1 << ((2) & 63))
                _ = v17
            } else {
                v3 = v3
            }
        }
    } else {
        fmt.Println(v3, v3, v1)
        v3 += v3
        var v18 uint8 = 31
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int8 = 80
        v21 := GenS1{F0: v3}
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = len(v4)
    fmt.Println(v1, v1)
    if false {
        var v22 int = -91
        _ = v22
    }
    var v23 uint64 = 658
    v3 = (v3 << ((368) & 63))
    v24 := GenS2{F0: genFn0(v1, 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
