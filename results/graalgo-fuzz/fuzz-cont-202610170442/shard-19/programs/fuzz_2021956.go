package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 667
    var v2 bool = true
    v3 := []int{42}
    v4 := make(map[int]string)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 32767, F1: 128, F2: (fi5 << ((425) & 63))}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = v2
        v2 = v8
        if ((!(!false)) || (42 < -73)) {
            var v9 string = ""
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int32 = -369
                v11 -= (v11 + v11)
                var v12 float32 = 0.1
                _ = fi10
                _ = v11
                _ = v12
            }
            fmt.Println(v9, v2)
            _ = v9
        } else {
            var v13 int = fi7
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    v2 = v2
    var v14 int32 = 42
    v14 = v14
    v3 = append(v3, 520)
    fmt.Println(v2, v2, v14)
    v15 := GenS1{F0: 825, F1: 128, F2: -127}
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v1, v2)
        var v17 bool = true
        var v18 uint8 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 int32 = (v14 - (v14 % ((v14) | 1)))
        fmt.Println(v20, fi19, v2)
        _ = len(v4)
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 2
    for fi22 := 0; fi22 < 5; fi22++ {
        if false {
            if (v2 || v2) {
                var v23 uint16 = (^(v21))
                fmt.Println(v2)
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS2{F0: 134}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
        }
        var v26 int32 = (v14 / ((v14) | 1))
        _ = fi22
        _ = v26
    }
    var v27 uint8 = 2
    if (("hello" + "foo bar") != ("" + "a")) {
        var v28 string = "a"
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
