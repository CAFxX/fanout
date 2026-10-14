package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 42
    var v3 uint64 = v2
    v4 := make(map[uint32]int)
    var v5 float64 = ((v1 * float64(0.1)) + -2.25)
    v6 := GenS1{F0: 128}
    v5 = v1
    v7 := []int{-884, -259}
    v7 = append(v7, 4294967295)
    var v8 uint32 = 8
    var v9 float32 = 3.14159
    if (-32768 == -32768) {
        switch -9223372036854775808 {
        case 680:
            var v10 uint64 = (^(v3))
            _ = v10
        case 9223372036854775807:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            if (((("\n\t" + "") >= "a") && (true && false)) && false) {
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 uint32 = 9
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    if true {
                        var v15 int64 = 1
                        _ = v15
                    }
                    if (false || true) {
                        switch (("foo bar" + "a") + ("hello" + ("" + "foo bar"))) {
                        case "":
                            v4[(^(v8))] = 0
                        default:
                            var v16 float32 = (v9 * (float32(v1) * float32(3.14159)))
                            _ = v16
                        }
                    }
                    v17 := GenS2{F0: ("" + ("a" + ((("\n\t" + "hello") + "foo bar") + ""))), F1: (v8 % ((v8) | 1))}
                    _ = fi14
                    _ = v17
                }
            }
            _ = fi11
        }
    }
    if false {
        fmt.Println(v2)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS3{F0: 2, F1: ""}
        var v20 uint = 871
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint64 = v3
    v22 := GenS3{F0: 239, F1: ("" + "\n\t")}
    var v23 float64 = (v5 - v5)
    _ = len(v4)
    v2 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
