package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 uint8 = 101
    var v4 int = 9223372036854775807
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{(v4 >> ((4294967295) & 63)), (fi7 ^ 32768), (v4 - v4), (v4 / ((32768) | 1))}
        _ = fi7
        _ = v8
    }
    if ((v2 % ((v2) | 1)) < (^(v2))) {
        if (100 < -1) {
            v3 -= v3
            v9 := []int{v4, v4, v4}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint32 = 65535
                v3 = v3
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint64 = 4294967295
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: 42}
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        for fi16 := 0; fi16 < 10; fi16++ {
                            v17 := GenS1{F0: -108}
                            _ = fi16
                            _ = v17
                        }
                    case "foo bar":
                        v18 := GenS1{F0: -128}
                        _ = v18
                    }
                    var v19 bool = true
                    var v20 int = (^(v4))
                    _ = v19
                    _ = v20
                } else {
                    v21 := []int{}
                    _ = v21
                }
                v3 = (v3 + v3)
                _ = fi14
                _ = v15
            }
        }
        var v22 uint64 = 42
        _ = v22
    } else {
        var v23 bool = true
        switch ((("" + "a") + ("a" + ("" + "\n\t"))) + "hello") {
        case "foo bar":
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint8 = (v3 >> ((94) & 63))
                var v26 float64 = ((0.0 + 1e10) * (0.0 - 0.1))
                _ = fi24
                _ = v25
                _ = v26
            }
        case "a":
            v27 := GenS2{F0: (("\n\t" + ("hello" + "a")) + (("foo bar" + "a") + "hello")), F1: 2, F2: ("foo bar" + (("foo bar" + "foo bar") + (("\n\t" + "foo bar") + "hello")))}
            _ = v27
        default:
            var v28 int16 = 42
            _ = v28
        }
        v2 -= v2
        _ = v23
    }
    if true {
        if (!(65535 >= 32767)) {
            var v29 float32 = 3.14159
            v4 += (v4 / ((v4) | 1))
            _ = v29
        } else {
            v3 -= (v3 ^ v3)
        }
        if (true && true) {
            var v30 uint8 = v3
            v31 := []int{32768, 32767}
            if len(v31) > 0 {
                v31[0] = -1
            }
            _ = v30
            _ = v31
        }
        v32 := GenS2{F0: "", F1: 65535, F2: (("foo bar" + "foo bar") + ("foo bar" + "hello"))}
        _ = v32
    } else {
        v3 -= (v3 ^ v3)
        for fi33 := 0; fi33 < 8; fi33++ {
            v2 = (-(v2))
            var v34 uint32 = 98
            var v35 uint = 18446744073709551615
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v36 := []int{(v4 >> ((0) & 63)), -214, (v4 & v4), -669}
    v3 = v3
    fmt.Println(v4)
    delete(v5, (v4 - (v4 * (v4 - v4))))
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    v37 := GenS2{F0: ("hello" + "\n\t"), F1: 2147483648, F2: ("foo bar" + "hello")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
