package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint }
func genFn0(p0 string) string {
    switch (p0 + "\n\t") {
    case "hello":
        var v7 uint64 = 53
        _ = v7
    }
    p0 = p0
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: 65535}
    v8 = (v8 & v8)
    _ = p0
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int32 = -397
    v2 := make(map[int]uint32)
    var v3 uint64 = 849
    var v4 int64 = 26
    var v5 float64 = -2.25
    var v6 int16 = 1000
    fmt.Println(v6, v1, v1)
    var v10 float32 = -2.25
    v11 := GenS2{F0: ("a" + ("hello" + "a"))}
    v1 += 0
    fmt.Println(v5)
    v12 := []int{305}
    v3 = v3
    v10 = float32(v5)
    var v13 int8 = -128
    if true {
        if true {
            v3 = v3
            switch 285 {
            case -1:
                fmt.Println(v5, v4)
            case 2147483647:
                var v14 float32 = v10
                _ = v14
            }
        } else {
            var v15 uint16 = 0
            _ = v15
        }
        switch 2147483647 {
        case 496:
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 uint = 117
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        case -2147483648:
            switch int(v6) {
            case 9223372036854775807:
                switch "foo bar" {
                case "a":
                    if len(v12) > 0 {
                        v12[0] = 127
                    }
                }
            case -1:
                fmt.Println(v6)
            case -9223372036854775808:
                v6 = (v6 % ((v6) | 1))
            default:
                if true {
                    v19 := GenS3{F0: v6, F1: 4294967295, F2: 65535}
                    fmt.Println(v5, v10, v5)
                    _ = v19
                }
            }
        }
    } else {
        fmt.Println(v10)
    }
    var v20 int32 = v1
    fmt.Println(v20)
    var v21 uint = 962
    fmt.Println(v10)
    v12 = append(v12, 127)
    for fi22 := 0; fi22 < 10; fi22++ {
        v6 -= (v6 | 0)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        if (v10 < (0.0 + v10)) {
            v23 := GenS4{F0: 18446744073709551615, F1: fi22, F2: 1}
            _ = v23
        } else {
            fmt.Println(v1, v1)
            var v24 int16 = (v6 + (-(v6)))
            fmt.Println(v13, v21, v1)
            _ = v24
        }
        _ = fi22
    }
    var v25 uint16 = 871
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
