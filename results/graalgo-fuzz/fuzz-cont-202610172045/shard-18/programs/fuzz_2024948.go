package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) bool {
    fmt.Println(p0)
    var v5 int64 = -513
    _ = p0
    _ = v5
    return (1 != 127)
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = v1
    var v3 uint8 = v2
    v4 := make(map[uint32]string)
    if true {
        if (42 == 16) {
            v6 := v4[315]
            _ = v6
        } else {
            var v7 uint16 = 2
            v8 := GenS1{F0: (!true), F1: -1}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            v2 = (v3 + v2)
            switch "a" {
            case "foo bar":
                if true {
                    v3 -= (v1 / ((206) | 1))
                }
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint32 = 1
                v1 -= (v3 ^ v3)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        v3 = 128
    } else {
        var v12 float32 = 3.14159
        _ = v12
    }
    fmt.Println(v1)
    if false {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint16 = 0
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            if (414 == 603) {
                v1 = v1
                var v16 uint16 = 255
                _ = v16
            } else {
                var v17 bool = false
                v18 := GenS1{F0: (-283 <= 1), F1: 919}
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int16 = 0
                v21 := GenS1{F0: true, F1: -975}
                var v22 int8 = int8(v20)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            fmt.Println(fi15)
            _ = fi15
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int8 = -38
        var v25 float64 = 1e10
        v1 -= 128
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := []int{int(v3), -317, 54, 1}
    for fi27 := 0; fi27 < 20; fi27++ {
        switch 4294967295 {
        case 32767:
            v28 := GenS1{F0: (0 == 665), F1: 1}
            _ = v28
        }
        _ = fi27
    }
    v29 := GenS2{F0: 634}
    var v30 uint8 = 64
    v31 := v4[0]
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
