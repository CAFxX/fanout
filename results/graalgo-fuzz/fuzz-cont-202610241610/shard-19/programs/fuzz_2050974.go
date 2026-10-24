package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    var v5 uint32 = 5
    var v6 int = 1
    if (false && true) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 bool = true
            _ = fi7
            _ = v8
        }
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    v3 := make(map[int]string)
    var v4 uint16 = v2
    switch 1 {
    case 0:
        var v9 int32 = 0
        _ = v9
    case -358:
        _ = len(v1)
    case -693:
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint = 65535
            if false {
                for fi12 := 0; fi12 < 12; fi12++ {
                    v2 = (^(v2))
                    v3[fi10] = "foo bar"
                    fmt.Println(v11, v11)
                    _ = fi12
                }
                v4 = v4
                v4 -= v4
            }
            v4 -= v2
            _ = fi10
            _ = v11
        }
    }
    delete(v1, 4294967295)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint8 = 127
        _ = fi13
        _ = v14
    }
    if false {
        var v15 uint8 = 122
        for fi16 := 0; fi16 < 12; fi16++ {
            v15 = 0
            _ = fi16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v4 -= v4
            _ = fi17
        }
        _ = v15
    } else {
        fmt.Println(v4)
    }
    v18 := GenS1{F0: 0, F1: 1, F2: 0}
    var v19 int = 777
    if true {
        fmt.Println(v4, v2, v4)
        var v20 int8 = -73
        v21 := GenS2{F0: v19}
        _ = v20
        _ = v21
    }
    v22 := []int{0, (v19 ^ 42), v19}
    var v23 int32 = -2147483648
    for fi24 := 0; fi24 < 11; fi24++ {
        v23 += v23
        if false {
            v25 := GenS1{F0: 9223372036854775807, F1: 2, F2: 65535}
            _ = v25
        }
        _ = fi24
    }
    v26 := GenS2{F0: (-(v19))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
