package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 int64, p1 bool) bool {
    fmt.Println(p1)
    var v4 int = 0
    v5 := GenS1{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (0.1 * (-2.25 + 0.0))
    v3 := make(map[int]bool)
    v6 := v3[-57]
    var v7 int64 = -9223372036854775808
    if (32767 < 32767) {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int64 = v7
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            if (fi10 > fi10) {
                v11 := v1[(fi10 * (fi10 << ((65535) & 63)))]
                fmt.Println(v7)
                _ = v11
            }
            v1[fi10] = "hello"
            _ = fi10
        }
        v7 = (v7 - (^(v7)))
    }
    v2 -= v2
    for fi12 := 0; fi12 < 2; fi12++ {
        v2 = ((3.14159 + float32((1.5 - -2.25))) + 1e10)
        v2 += -2.25
        _ = len(v1)
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            v2 += v2
            if true {
                switch fi13 {
                case 0:
                    var v15 string = ""
                    _ = v15
                case 127:
                    fmt.Println(v6, fi14, fi14)
                    v2 = v2
                case 323:
                    switch (fi14 & fi14) {
                    case 128:
                        v16 := GenS2{F0: 1000, F1: 42, F2: 608}
                        _ = v16
                    case -658:
                        v17 := GenS3{F0: -249, F1: v7}
                        _ = v17
                    default:
                        var v18 uint = 2
                        _ = v18
                    }
                default:
                    var v19 uint64 = 500
                    _ = v19
                }
                _ = len(v3)
            }
            v2 = 1e10
            _ = fi14
        }
        switch "" {
        case "a":
            switch ("hello" + "a") {
            case "":
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := v1[32768]
                    var v22 int8 = 45
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                v6 = (true && (((^(v7)) == (^(v7))) && (255 < 2)))
            }
        case "foo bar":
            fmt.Println(v2)
        }
        _ = fi13
    }
    fmt.Println(v6)
    var v23 int16 = 42
    v24 := []int{718, -172}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v6
    _ = v7
}
