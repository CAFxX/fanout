package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = 650
    var v2 float64 = 3.14159
    var v3 float64 = (1e10 + (1e10 - float64(((1e10 - 0.0) - (0.1 + 1e10)))))
    var v4 int32 = -601
    v5 := []string{"\n\t", ("\n\t" + ("" + "a")), ("foo bar" + "a")}
    v6 := make(map[int]string)
    if false {
        v2 += (v3 + v2)
        v3 = 3.14159
    } else {
        v1 = v1
        v2 += 1.5
    }
    v7 := GenS1{F0: (false || false)}
    if len(v5) > 0 {
        v5[0] = ("" + "\n\t")
    }
    var v8 int16 = 1000
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int16 = (v1 + 1)
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: (true || (("" + ("hello" + "hello")) == (("" + "\n\t") + "hello")))}
            if true {
                switch "\n\t" {
                case "a":
                    var v13 float32 = 0.1
                    _ = v13
                case "foo bar":
                    v14 := v6[0]
                    _ = v14
                default:
                    var v15 float64 = (v2 + (3.14159 + float64((0.1 - float32(v3)))))
                    _ = v15
                }
            } else {
                v3 = (((v2 - (0.1 * (v2 + (0.1 / v2)))) + 1e10) - (v3 / (((1e10 * (v2 * ((0.0 * (float64(float32(1e10)) / v3)) / 0.1))) + -2.25) - v3)))
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v16 uint64 = 94
    var v17 uint32 = 826
    v8 = (v8 - v1)
    var v18 int8 = 0
    var v19 uint16 = 223
    var v20 int64 = 1
    var v21 uint16 = 254
    var v22 int16 = (v1 << ((65535) & 63))
    v17 += 533
    var v23 int64 = 91
    var v24 uint64 = 1
    var v25 uint16 = (v19 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
