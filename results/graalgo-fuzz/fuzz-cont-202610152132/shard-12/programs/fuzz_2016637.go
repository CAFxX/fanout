package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 uint32, p1 uint8) int64 {
    var v6 uint = 18446744073709551615
    if true {
        var v7 int8 = 99
        fmt.Println(p1)
        _ = v7
    }
    v6 -= v6
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = v6
    return int64(p1)
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 int64 = 1
    v4 := make(map[string]uint32)
    var v5 string = ("\n\t" + ((("hello" + "foo bar") + ((("hello" + "") + "a") + "")) + "a"))
    if (1e10 <= (3.14159 + (3.14159 / 0.1))) {
        if false {
            fmt.Println(v3, v3)
            var v8 int32 = -1
            _ = v8
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 int = (-(fi9))
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 2, F1: 750, F2: 2147483647}
            _ = v11
        }
        v5 = "\n\t"
    }
    v5 = ""
    if false {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (v5 + (v5 + v5)) {
            case "\n\t":
                _ = len(v4)
            case "foo bar":
                if (1 == 255) {
                    var v13 float32 = (0.1 * 1e10)
                    _ = len(v1)
                    v14 := GenS1{F0: 128, F1: 42, F2: 43}
                    _ = v13
                    _ = v14
                }
            default:
                v4[((v5 + v5) + "\n\t")] = 0
            }
            _ = fi12
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int32 = 1
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            fmt.Println(v3, v3)
            _ = fi17
        }
    } else {
        fmt.Println(v5, v5)
    }
    fmt.Println(v5)
    _ = len(v2)
    fmt.Println(v3, v5)
    v3 += v3
    v3 += genFn0(996, 87)
    for fi18 := 0; fi18 < 15; fi18++ {
        if true {
            _ = len(v2)
            switch ("\n\t" + "foo bar") {
            case "a":
                v5 = "\n\t"
            default:
                switch (v5 + (v5 + v5)) {
                case "hello":
                    var v19 uint32 = 495
                    _ = v19
                case "a":
                    var v20 bool = true
                    _ = v20
                case "":
                    _ = len(v2)
                default:
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 float64 = 1e10
                        fmt.Println(v5)
                        v23 := GenS2{F0: -690, F1: 2147483648}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
            }
            v3 = genFn0(1, 194)
        } else {
            var v24 uint64 = 1
            var v25 int16 = -658
            var v26 int64 = (v3 - v3)
            _ = v24
            _ = v25
            _ = v26
        }
        switch "a" {
        case "hello":
            v27 := GenS2{F0: (^(fi18)), F1: 2147483647}
            _ = v27
        case "a":
            v5 = (("\n\t" + ("" + ("hello" + "\n\t"))) + "hello")
        default:
            v4[""] = 0
        }
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
