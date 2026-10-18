package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int16, p1 uint) int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        p1 -= p1
        for fi7 := 0; fi7 < 7; fi7++ {
            if false {
                var v8 float64 = (float64(0.1) - 3.14159)
                fmt.Println(fi6)
                v9 := GenS1{F0: -240, F1: 440, F2: (("" + "hello") + (("foo bar" + "hello") + "\n\t"))}
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    fmt.Println(p1, p1)
    p1 = 18446744073709551615
    if (!((!(!(false || true))) && (uint16(p1) <= 65535))) {
        fmt.Println(p1)
    }
    v10 := GenS2{F0: 361}
    _ = p0
    _ = p1
    _ = v10
    return -128
}

func main() {
    var v1 float64 = (1.5 + 0.1)
    var v2 uint8 = 107
    var v3 int16 = 42
    var v4 int = 974
    var v5 int64 = 42
    switch (v4 / ((v4) | 1)) {
    case 792:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = 259
            var v13 int16 = v3
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    if ((!false) || (!(!false))) {
        var v14 float32 = float32(3.14159)
        _ = v14
    } else {
        var v15 int = -893
        var v16 uint64 = 2
        _ = v15
        _ = v16
    }
    v17 := []int{v4, 4294967295}
    var v18 int = 938
    v4 = 32768
    var v19 uint = 2
    v3 = 1000
    v17 = append(v17, -122)
    v18 -= (-(v18))
    v19 -= (v19 * v19)
    var v20 float32 = 0.0
    for fi21 := 0; fi21 < 12; fi21++ {
        switch -535 {
        case 42:
            var v22 int = 32767
            _ = v22
        }
        var v23 int = v18
        v24 := GenS3{F0: true}
        _ = fi21
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v17 = append(v17, (v18 * v4))
        if len(v17) > 0 {
            v17[0] = fi25
        }
        _ = fi25
    }
    v17 = append(v17, v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
