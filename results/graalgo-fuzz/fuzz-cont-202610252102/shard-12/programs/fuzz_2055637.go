package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64) bool {
    v4 := GenS1{F0: -1, F1: 814}
    p0 -= (p0 << ((0) & 63))
    _ = p0
    _ = v4
    return (100 <= 100)
}
func genFn1(p0 int8, p1 int16, p2 string) uint {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p1, p0)
        var v6 uint8 = 121
        var v7 int32 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: p0, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 128
    var v3 float32 = 0.1
    var v9 string = "hello"
    var v10 float64 = 0.1
    v11 := GenS1{F0: v1, F1: 76}
    var v12 uint64 = 940
    switch 422 {
    case 4294967295:
        var v13 string = v9
        _ = v13
    default:
        v9 = (v9 + "\n\t")
    }
    v14 := []int{}
    switch ("\n\t" + (v9 + "foo bar")) {
    case "\n\t":
        v15 := GenS2{F0: 2147483647, F1: (^(v1))}
        _ = v15
    case "a":
        v10 -= 3.14159
    default:
        if true {
            for fi16 := 0; fi16 < 4; fi16++ {
                if (v1 != (v1 ^ (v1 + v1))) {
                    var v17 string = ("\n\t" + v9)
                    var v18 bool = true
                    _ = v17
                    _ = v18
                } else {
                    _ = len(v14)
                    fmt.Println(v9)
                }
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi16
            }
            switch (("" + v9) + "foo bar") {
            case "":
                for fi19 := 0; fi19 < 8; fi19++ {
                    v10 = -2.25
                    _ = fi19
                }
            }
            fmt.Println(v12, v2)
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 = 42
        _ = fi20
    }
    if true {
        v10 = -2.25
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint16 = 255
            _ = fi21
            _ = v22
        }
    } else {
        v2 = (^(v2))
    }
    var v23 int8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v9
}
