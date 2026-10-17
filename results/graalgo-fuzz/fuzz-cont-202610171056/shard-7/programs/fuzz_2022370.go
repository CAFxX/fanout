package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 string }

func main() {
    v1 := []string{(("hello" + "hello") + ("hello" + "foo bar")), "hello", "foo bar"}
    var v2 uint = 65535
    var v3 int64 = -785
    v4 := []int{9223372036854775807, 4294967295, 32768, 127}
    v5 := make(map[int]bool)
    var v6 int8 = 1
    fmt.Println(v3, v2, v3)
    if (4294967295 != 42) {
        var v7 uint32 = 1
        _ = v7
    } else {
        v8 := v5[9223372036854775807]
        _ = v8
    }
    var v9 string = "hello"
    var v10 int32 = 0
    fmt.Println(v2, v3)
    if (!false) {
        _ = len(v1)
        switch 9223372036854775807 {
        case 9223372036854775807:
            v10 = 2147483647
        case -9223372036854775808:
            v11 := GenS1{F0: false}
            _ = v11
        case 128:
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 uint = v2
                _ = fi12
                _ = v13
            }
        }
        v14 := v5[-892]
        _ = v14
    } else {
        v10 = (v10 << (((^(v2))) & 63))
        fmt.Println(v3)
        v2 = v2
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        if true {
            var v16 uint32 = 2
            switch (v9 + "hello") {
            case "hello":
                if len(v4) > 0 {
                    v4[0] = fi15
                }
            case "":
                _ = len(v1)
            }
            if (((0.1 / -2.25) >= 0.0) && (26 <= 499)) {
                if (((v9 + v9) + v9) >= "\n\t") {
                    var v17 uint16 = 42
                    v1 = append(v1, ((((v9 + "\n\t") + v9) + v9) + ""))
                    _ = v17
                } else {
                    fmt.Println(v16)
                    var v18 uint64 = 1
                    _ = v18
                }
                if ((^(v2)) <= v2) {
                    v16 = v16
                    var v19 int8 = 127
                    _ = v19
                }
                var v20 int8 = 112
                _ = v20
            }
            _ = v16
        } else {
            switch ("foo bar" + ("a" + "hello")) {
            case "":
                fmt.Println(fi15)
            case "a":
                for fi21 := 0; fi21 < 20; fi21++ {
                    v4 = append(v4, (-(fi15)))
                    v22 := GenS1{F0: false}
                    _ = fi21
                    _ = v22
                }
            default:
                v23 := GenS2{F0: (^(v2)), F1: -32768, F2: 1}
                _ = v23
            }
            v24 := GenS1{F0: ((-(v6)) == int8(v2))}
            _ = v24
        }
        var v25 float64 = ((0.1 * 3.14159) / 1.0)
        v3 += v3
        _ = fi15
        _ = v25
    }
    if false {
        v26 := GenS3{F0: (v9 + v9)}
        _ = v26
    } else {
        fmt.Println(v10, v2, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
