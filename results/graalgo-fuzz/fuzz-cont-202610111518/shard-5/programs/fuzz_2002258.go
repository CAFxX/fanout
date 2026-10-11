package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0() uint8 {
    v7 := GenS1{F0: 1}
    v8 := GenS1{F0: -1}
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int16) uint32 {
    p0 = p0
    v9 := GenS2{F0: -487}
    var v10 uint32 = 4294967295
    _ = p0
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 int16 = 1
    v2 := []string{("foo bar" + "\n\t")}
    var v3 bool = (!(221 >= 603))
    v4 := []int{0}
    var v5 uint64 = 2
    var v6 int16 = (v1 + v1)
    var v11 float32 = 3.14159
    v6 += 42
    v5 += (v5 << ((1) & 63))
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v2) > 0 {
            v2[0] = ((("foo bar" + "\n\t") + "") + "")
        }
        var v13 int16 = v1
        if ((9 < fi12) && v3) {
            var v14 int64 = -374
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if false {
        var v15 bool = v3
        var v16 bool = (!v3)
        _ = v15
        _ = v16
    }
    if false {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(fi17, v1)
                _ = fi18
            }
            _ = fi17
        }
        switch "hello" {
        case "hello":
            var v19 float32 = ((((((v11 * 0.1) / (-2.25 + 1.5)) * 0.0) * v11) - ((3.14159 - ((3.14159 / v11) / (0.0 + v11))) * (float32(1.5) + v11))) + v11)
            _ = v19
        case "\n\t":
            var v20 float32 = float32(3.14159)
            _ = v20
        }
        var v21 uint8 = 1
        _ = v21
    } else {
        fmt.Println(v11, v3)
    }
    fmt.Println(v3)
    v5 = (v5 + (v5 >> ((644) & 63)))
    var v22 string = "foo bar"
    var v23 int64 = 1
    switch v22 {
    case "":
        fmt.Println(v22)
    case "hello":
        v11 -= -2.25
    }
    v2 = append(v2, (("\n\t" + v22) + v22))
    var v24 int8 = -106
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 float32 = v11
        v26 += v11
        _ = fi25
        _ = v26
    }
    var v27 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
