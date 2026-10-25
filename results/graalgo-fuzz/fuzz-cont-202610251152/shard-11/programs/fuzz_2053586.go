package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 bool) int16 {
    v7 := []int{42, -1}
    p0 = p0
    fmt.Println(p1, p0)
    var v8 uint = 755
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(fi10, fi10)
            p0 += p0
            v8 = (v8 | v8)
            if (-612 != -1) {
                v8 = 0
            }
            v11 := GenS1{F0: 13, F1: v8, F2: 977}
            if len(v7) > 0 {
                v7[0] = fi10
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(p1, p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}

func main() {
    v1 := []int{745, 128, 42, -1}
    var v2 uint32 = 580
    v3 := make(map[string]bool)
    var v4 uint = 4294967295
    v5 := []uint32{v2, (v2 | (v2 >> ((v4) & 63)))}
    var v6 float64 = ((0.0 - (((3.14159 - 1.5) * 0.1) * 3.14159)) + 0.0)
    for fi12 := 0; fi12 < 3; fi12++ {
        if false {
            var v13 bool = false
            if false {
                var v14 int32 = 42
                var v15 uint8 = 0
                v6 = v6
                _ = v14
                _ = v15
            } else {
                var v16 float64 = float64(((0.0 / 0.1) * 3.14159))
                _ = v16
            }
            switch (("a" + "\n\t") + "foo bar") {
            case "":
                fmt.Println(v6, v13)
            }
            _ = v13
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 string = "foo bar"
        var v19 uint32 = v2
        var v20 int16 = 32767
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    if true {
        v21 := GenS1{F0: -1, F1: v4, F2: 1}
        v22 := GenS2{F0: true}
        _ = v21
        _ = v22
    }
    fmt.Println(v2, v4, v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS2{F0: false}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v6 -= float64((3.14159 * (3.14159 + ((-2.25 / -2.25) + -2.25))))
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
