package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{""}
    v3 := []int{1}
    var v4 uint8 = 138
    var v5 uint8 = (v4 * v4)
    v6 := GenS1{F0: 450, F1: 1}
    v4 = (v5 * 209)
    v4 = v4
    if false {
        if (((0 >= 42) && (!true)) || true) {
            v5 = 255
        } else {
            var v7 string = ""
            _ = v7
        }
        var v8 uint16 = 786
        v8 = (v8 << ((636) & 63))
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v4 -= v5
        fmt.Println(v5, fi9)
        _ = fi9
    }
    v4 = v4
    fmt.Println(v4, v4)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = -546
        if (4294967295 == 0) {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int16 = 880
                if true {
                    v5 += (v5 + v4)
                    if false {
                        var v14 uint = 42
                        v15 := GenS2{F0: 0, F1: (fi10 < int(v5)), F2: 524}
                        if len(v2) > 0 {
                            v2[0] = ("hello" + "hello")
                        }
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi12
                _ = v13
            }
            fmt.Println(v5, v5)
            v16 := GenS2{F0: 0, F1: (!true), F2: 2147483647}
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    var v17 int16 = -707
    switch 42 {
    case 770:
        fmt.Println(v5, v4, v17)
    default:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    delete(v1, 775)
    fmt.Println(v17, v17, v4)
    var v18 string = "foo bar"
    if len(v3) > 0 {
        v3[0] = -905
    }
    v5 = 0
    if (!true) {
        var v19 float64 = 3.14159
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(v5, fi20)
            for fi21 := 0; fi21 < 5; fi21++ {
                switch (("foo bar" + "\n\t") + "a") {
                case "a":
                    v4 += 128
                case "":
                    v22 := v1[(-(fi20))]
                    _ = v22
                default:
                    _ = len(v3)
                }
                v4 = (^(v5))
                if len(v3) > 0 {
                    v3[0] = 839
                }
                _ = fi21
            }
            _ = fi20
        }
        v3 = append(v3, 4294967295)
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
