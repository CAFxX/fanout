package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{(("hello" + ("" + "foo bar")) + ""), "hello", "foo bar", "\n\t"}
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 uint16 = 1
    var v5 uint16 = 42
    fmt.Println(v5)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch fi6 {
            case 66:
                var v8 int = fi6
                _ = v8
            case 42:
                v4 = (v4 >> ((255) & 63))
            }
            _ = fi7
        }
        var v9 float64 = (1.5 / -2.25)
        _ = fi6
        _ = v9
    }
    var v10 uint8 = 28
    v10 -= 162
    for fi11 := 0; fi11 < 10; fi11++ {
        _ = len(v3)
        for fi12 := 0; fi12 < 16; fi12++ {
            if (-456 <= -1) {
                v5 = uint16(v10)
                v4 += v4
                v5 += 876
            }
            _ = fi12
        }
        var v13 uint = 334
        _ = fi11
        _ = v13
    }
    var v14 string = "foo bar"
    v14 = v14
    v15 := GenS1{F0: 100, F1: v14, F2: "hello"}
    fmt.Println(v4, v10)
    fmt.Println(v4, v5, v10)
    for fi16 := 0; fi16 < 9; fi16++ {
        v10 = 124
        if (787 != 42) {
            var v17 int64 = -1
            _ = v17
        } else {
            var v18 int32 = 714
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS1{F0: -1, F1: ("foo bar" + "foo bar"), F2: v14}
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = fi16
    }
    var v21 uint8 = (v10 + 44)
    if false {
        v4 = (v4 ^ v4)
        switch 127 {
        case 42:
            v22 := GenS2{F0: 127}
            _ = v22
        case 4294967295:
            fmt.Println(v21, v10)
        case -1:
            if true {
                v23 := GenS2{F0: 42}
                v14 = ("a" + v14)
                fmt.Println(v21, v10, v4)
                _ = v23
            } else {
                v10 -= 0
                for fi24 := 0; fi24 < 3; fi24++ {
                    delete(v3, v14)
                    v21 = (v10 | (v10 | v10))
                    for fi25 := 0; fi25 < 7; fi25++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v1)
                        _ = fi25
                    }
                    _ = fi24
                }
            }
        }
    }
    fmt.Println(v14, v10, v10)
    v26 := v3[v14]
    switch -9223372036854775808 {
    case 502:
        v26 = "foo bar"
    case 0:
        fmt.Println(v21, v21, v10)
    case 1:
        v26 = (((v26 + v26) + (v14 + v14)) + "")
    }
    v4 = (v5 << ((1) & 63))
    v27 := v2[9223372036854775807]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
