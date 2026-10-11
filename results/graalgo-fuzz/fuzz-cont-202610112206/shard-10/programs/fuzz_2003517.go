package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0() string {
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint = 889
    var v7 int64 = -768
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 uint32, p1 uint, p2 string) int16 {
    v8 := GenS1{F0: 344, F1: -627, F2: 255}
    switch p2 {
    case "\n\t":
        var v9 int64 = 946
        _ = v9
    }
    p2 = p2
    p0 -= (p0 ^ 478)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 374
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]int)
    v3 = 0
    v3 = v3
    var v10 int = 128
    v11 := GenS1{F0: 42, F1: 2147483647, F2: 2}
    v12 := GenS1{F0: -1, F1: 42, F2: 255}
    var v13 uint16 = 399
    switch v10 {
    case 42:
        v14 := v1[v10]
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            switch fi15 {
            case 492:
                switch "a" {
                case "\n\t":
                    var v17 uint = 1
                    var v18 uint16 = (v13 & v13)
                    var v19 int16 = genFn1(550, 42, ("hello" + ("foo bar" + (("\n\t" + "hello") + "hello"))))
                    if (true || true) {
                        v2[1] = ("a" + "hello")
                    }
                    v17 += 42
                    v17 -= (^(v17))
                    fmt.Println(fi15)
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
