package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 uint = uint(v1)
    var v3 bool = false
    v4 := GenS1{F0: (("" + "a") + ("hello" + (((("foo bar" + "foo bar") + "hello") + (("a" + ("foo bar" + "")) + "hello")) + "hello"))), F1: 889}
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            switch "foo bar" {
            case "a":
                if (v3 && v3) {
                    v2 += 523
                    v2 -= uint(v1)
                } else {
                    v7 := GenS2{F0: 234, F1: fi6}
                    fmt.Println(v1, v2, v1)
                    var v8 uint8 = 42
                    _ = v7
                    _ = v8
                }
            default:
                if (0.1 == 0.1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            _ = fi6
        }
    }
    var v9 string = "foo bar"
    var v10 int = 4294967295
    fmt.Println(v9, v2, v3)
    v11 := GenS1{F0: v9, F1: -1}
    v12 := GenS3{F0: uint16(v1)}
    var v13 string = v9
    v1 -= 42
    if (!(1 > 225)) {
        var v14 int8 = 1
        _ = v14
    } else {
        if v3 {
            var v15 uint = (v2 | v2)
            for fi16 := 0; fi16 < 19; fi16++ {
                if len(v5) > 0 {
                    v5[0] = (-(fi16))
                }
                v1 -= v1
                v17 := GenS3{F0: 0}
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    var v18 bool = (!(!(1 <= 297)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
