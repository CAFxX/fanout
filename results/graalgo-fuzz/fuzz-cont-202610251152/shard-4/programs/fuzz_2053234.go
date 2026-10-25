package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 uint = 42
    v5 := make(map[string]uint32)
    v6 := []uint32{42, 791}
    switch 9223372036854775807 {
    case -499:
        var v7 uint = (v4 - 0)
        _ = v7
    default:
        var v8 int32 = -2147483648
        _ = v8
    }
    v3 = v3
    v1 -= ((v1 - v1) * (v1 - (v1 + float64(3.14159))))
    if ((!(true && v3)) && true) {
        v5[("foo bar" + "")] = 2147483648
        if (!v3) {
            var v9 int8 = 127
            v10 := GenS1{F0: 0, F1: "a"}
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            if v3 {
                for fi12 := 0; fi12 < 19; fi12++ {
                    if v3 {
                        var v13 bool = v3
                        v14 := GenS1{F0: 72, F1: (("" + "a") + "\n\t")}
                        var v15 uint = 1
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        v6 = append(v6, 42)
                        v1 += v1
                        v1 = 0.0
                    }
                    fmt.Println(v4, v4)
                    var v16 uint = v4
                    _ = fi12
                    _ = v16
                }
                switch (("hello" + ("hello" + "")) + "") {
                case "\n\t":
                    var v17 float32 = 1e10
                    _ = v17
                default:
                    var v18 float32 = (0.1 + 3.14159)
                    _ = v18
                }
            } else {
                v4 = (v4 << ((0) & 63))
                v4 = (^(v4))
            }
            _ = fi11
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint32 = 42
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 2, F1: "a"}
        var v23 int8 = 0
        _ = fi19
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            v1 = ((v1 * 0.1) * float64(0.0))
            var v26 uint8 = 25
            _ = fi25
            _ = v26
        }
        fmt.Println(fi24, v3)
        _ = fi24
    }
    v27 := GenS2{F0: -255, F1: 0}
    v1 = 3.14159
    v3 = true
    var v28 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
