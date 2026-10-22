package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := []string{""}
    for fi4 := 0; fi4 < 5; fi4++ {
        v3 = append(v3, "\n\t")
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 bool = ((227 == 0) && false)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            if true {
                fmt.Println(fi7)
            }
            var v8 float32 = (((1e10 - 1.5) + 1.5) - 1e10)
            var v9 int8 = 127
            _ = v8
            _ = v9
        } else {
            switch "hello" {
            case "hello":
                if (-2.25 != float64(((((0.1 + 1.5) * 0.1) / (0.0 + 1e10)) * -2.25))) {
                    var v10 string = ""
                    _ = v10
                }
            default:
                fmt.Println(fi7)
            }
        }
        var v11 uint = 998
        var v12 string = ("a" + "hello")
        _ = fi7
        _ = v11
        _ = v12
    }
    var v13 bool = true
    v13 = (!(v13 || false))
    for fi14 := 0; fi14 < 7; fi14++ {
        if (-1 == -1) {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(fi14, v13)
            default:
                var v15 uint = 349
                _ = v15
            }
            v16 := GenS1{F0: 977, F1: 128, F2: (fi14 % (((-(fi14))) | 1))}
            fmt.Println(fi14, fi14)
            _ = v16
        }
        var v17 int64 = 9223372036854775807
        fmt.Println(v17)
        _ = fi14
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                v19 := GenS1{F0: 42, F1: 42, F2: fi18}
                var v20 int64 = -1
                v13 = (4294967295 <= 2147483648)
                _ = v19
                _ = v20
            } else {
                v13 = false
                for fi21 := 0; fi21 < 7; fi21++ {
                    v13 = (!(("a" + "hello") <= ((("\n\t" + "foo bar") + "a") + "a")))
                    v13 = v13
                    var v22 uint = 100
                    _ = fi21
                    _ = v22
                }
                var v23 uint64 = 0
                _ = v23
            }
            v24 := GenS1{F0: 1, F1: 144, F2: 9223372036854775807}
            _ = fi18
            _ = v24
        }
        var v25 uint64 = 397
        _ = v25
    }
    var v26 uint16 = 1
    fmt.Println(v26, v26)
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int16 = 0
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
}
