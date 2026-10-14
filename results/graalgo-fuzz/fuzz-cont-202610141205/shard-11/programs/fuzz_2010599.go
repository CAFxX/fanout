package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint64 = 1
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    var v5 uint64 = 4294967295
    switch -1 {
    case -2147483648:
        fmt.Println(v2)
    case -9223372036854775808:
        var v6 string = "a"
        _ = v6
    default:
        var v7 int16 = 1
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: v2}
    v3[-870] = "\n\t"
    if true {
        v9 := v3[4294967295]
        var v10 string = v9
        var v11 uint64 = v5
        _ = v10
        _ = v11
        _ = v9
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint64 = (v1 ^ v5)
            var v14 float64 = float64(1.5)
            _ = len(v3)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if (false || (v1 <= 4294967295)) {
        switch ("a" + "foo bar") {
        case "\n\t":
            delete(v4, "")
        }
        v2 = v2
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint16 = 32767
            if true {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int16 = 0
                    switch (("" + ("hello" + ("hello" + "hello"))) + "hello") {
                    case "a":
                        v18 = v18
                    case "":
                        v19 := GenS2{F0: 42, F1: (v16 < v16), F2: v2}
                        _ = v19
                    case "hello":
                        v20 := GenS2{F0: 4294967295, F1: (!true), F2: (v2 >> ((18) & 63))}
                        _ = v20
                    default:
                        if (-1 != 1) {
                            v21 := GenS3{F0: 496, F1: 42}
                            _ = v21
                        } else {
                            v5 += (v1 + 0)
                            v1 = 0
                            var v22 uint8 = 42
                            _ = v22
                        }
                    }
                    v23 := GenS2{F0: 1, F1: ((fi15 + -1) >= fi17), F2: -1}
                    v24 := GenS1{F0: 0, F1: v2}
                    _ = fi17
                    _ = v18
                    _ = v23
                    _ = v24
                }
            }
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v5)
        if true {
            var v25 bool = false
            _ = v25
        }
        fmt.Println(v1, v2)
    }
    v2 = (v2 ^ 42)
    fmt.Println(v1)
    v26 := GenS2{F0: 255, F1: (255 != 128), F2: v2}
    var v27 uint64 = (v1 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
