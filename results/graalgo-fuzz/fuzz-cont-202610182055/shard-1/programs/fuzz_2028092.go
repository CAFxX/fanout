package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    v1 := []string{("" + "hello"), "a"}
    v2 := []string{"a", ("hello" + "hello"), "foo bar"}
    var v3 uint32 = 930
    var v4 bool = (false && (false && false))
    switch ("foo bar" + "hello") {
    case "":
        var v5 uint16 = 65535
        _ = v5
    default:
        if (2147483647 >= -92) {
            for fi6 := 0; fi6 < 11; fi6++ {
                v3 = (v3 + 676)
                var v7 uint16 = 0
                _ = len(v2)
                _ = fi6
                _ = v7
            }
            switch -328 {
            case -433:
                v8 := GenS1{F0: 377}
                _ = v8
            case 2147483647:
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int64 = 0
                    _ = fi9
                    _ = v10
                }
            case -1:
                v4 = (!(1 == 0))
            }
        } else {
            _ = len(v2)
        }
    }
    v11 := GenS1{F0: 0}
    var v12 uint16 = 344
    var v13 float32 = (-2.25 / 1.5)
    var v14 uint = 1
    v15 := GenS1{F0: v12}
    v12 = 255
    switch "a" {
    case "foo bar":
        v13 -= float32(0.0)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v12 += v12
    for fi16 := 0; fi16 < 19; fi16++ {
        v2 = append(v2, ("" + (("hello" + "a") + "a")))
        v17 := GenS2{F0: 529, F1: v14}
        for fi18 := 0; fi18 < 9; fi18++ {
            v14 = (^(v14))
            v12 = 32767
            var v19 bool = (127 >= 100)
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v13 -= v13
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
