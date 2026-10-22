package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0() int32 {
    var v7 string = ("" + (("foo bar" + "\n\t") + ""))
    v7 = ""
    v7 = v7
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 bool) int16 {
    v8 := GenS1{F0: -1, F1: genFn0()}
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(fi9, p1, p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 692
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    var v3 float64 = (0.0 / ((float64(float32(-2.25)) / 1e10) + (-2.25 / float64((((float32(1.5) - 0.1) + float32((0.0 / (3.14159 + 0.0)))) / 1.0)))))
    var v4 uint32 = 2147483647
    var v5 uint = 2
    var v6 uint64 = 453
    fmt.Println(v4, v6)
    fmt.Println(v3, v5)
    var v10 int8 = 127
    v11 := GenS2{F0: 65535}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint64 = (v6 - (v6 >> ((v5) & 63)))
        for fi14 := 0; fi14 < 10; fi14++ {
            v3 = 1.5
            var v15 uint16 = 455
            _ = fi14
            _ = v15
        }
        v2 = genFn1(genFn0(), (42 > 1))
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v6)
    var v16 int16 = -1
    var v17 float64 = v3
    if false {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v6, v16)
            _ = fi18
        }
        v19 := []int{42, 32768}
        _ = v19
    } else {
        v5 += (v5 - v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
