package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{-981}
    var v3 string = (((("\n\t" + "foo bar") + "\n\t") + "foo bar") + "a")
    var v4 uint = 255
    var v5 uint16 = 425
    switch v3 {
    case "\n\t":
        fmt.Println(v3, v5, v4)
    case "foo bar":
        for fi6 := 0; fi6 < 6; fi6++ {
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS1{F0: 1}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    case "":
        var v9 uint32 = 2147483648
        _ = v9
    default:
        fmt.Println(v3, v5)
    }
    v10 := GenS1{F0: 0}
    v11 := GenS1{F0: 1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -9223372036854775808 {
    case 376:
        var v12 string = (v3 + "foo bar")
        _ = v12
    default:
        if (!false) {
            var v13 float32 = 1.5
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 11; fi15++ {
                    if false {
                        var v16 uint = (^(v4))
                        v17 := GenS2{F0: fi14, F1: (!(v5 >= v5)), F2: v5}
                        v18 := v1[v3]
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 string = ("" + (("\n\t" + "") + v3))
                    _ = fi15
                    _ = v19
                }
                fmt.Println(fi14)
                v5 -= 255
                _ = fi14
            }
            var v20 string = "\n\t"
            _ = v20
        }
    }
    v5 += 2
    delete(v1, "a")
    var v21 float64 = ((1.5 * 1.5) * 1.5)
    if false {
        if (false || true) {
            v22 := GenS1{F0: 495}
            _ = v22
        } else {
            var v23 int32 = 2147483647
            v23 = v23
            var v24 float32 = 1e10
            _ = v23
            _ = v24
        }
    }
    v21 += v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
