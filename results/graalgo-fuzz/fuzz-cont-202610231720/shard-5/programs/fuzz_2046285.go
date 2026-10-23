package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint8 = 0
    var v2 int16 = 42
    v3 := make(map[uint32]string)
    var v4 string = "a"
    var v5 uint8 = v1
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: (fi7 % ((fi7) | 1)), F1: 127}
        _ = fi7
        _ = v8
    }
    var v9 int64 = 0
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            fmt.Println(v9, v4, v1)
            var v11 bool = true
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            switch int(v9) {
            case -284:
                var v13 uint8 = v5
                _ = v13
            case -2147483648:
                var v14 uint8 = 199
                _ = v14
            }
            _ = fi12
        }
        v1 -= v5
        _ = fi10
    }
    if false {
        var v15 int = 91
        v4 = v4
        _ = v15
    }
    v16 := []int{697}
    v2 += v2
    v9 = v9
    var v17 int = -2147483648
    v17 = (v17 * 2147483647)
    if true {
        v18 := GenS2{F0: 12}
        v1 = 127
        var v19 int = v17
        _ = v18
        _ = v19
    } else {
        if ((v4 >= "\n\t") || false) {
            switch v17 {
            case 9223372036854775807:
                delete(v3, 4294967295)
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                for fi21 := 0; fi21 < 8; fi21++ {
                    switch 32768 {
                    case 32767:
                        fmt.Println(fi21)
                        v5 = (v5 * v5)
                    case 0:
                        fmt.Println(v2, v5)
                    }
                    v9 = v9
                    _ = fi21
                }
                _ = fi20
            }
        }
    }
    v5 = (^(v5))
    v4 = v4
    switch ("a" + (v4 + v4)) {
    case "":
        v16 = append(v16, v17)
    }
    switch 4294967295 {
    case -410:
        v2 -= v2
    default:
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 int16 = v2
            v5 += (^(v5))
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
