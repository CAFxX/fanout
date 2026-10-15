package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 int }

func main() {
    var v1 uint8 = 128
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 uint = 394
    v5 := make(map[uint32]string)
    var v6 uint = 255
    switch "foo bar" {
    case "a":
        if false {
            if v2 {
                switch 0 {
                case -450:
                    delete(v5, 2147483647)
                }
                fmt.Println(v4)
            }
            v7 := []int{0, 1}
            _ = v7
        } else {
            if true {
                fmt.Println(v4, v4, v2)
            }
        }
    case "foo bar":
        v8 := GenS1{F0: -1, F1: -1, F2: 1}
        _ = v8
    }
    v6 = (v6 & v4)
    v4 = v4
    for fi9 := 0; fi9 < 5; fi9++ {
        v2 = (!(((true && ((-(fi9)) != -9223372036854775808)) && true) || false))
        var v10 int16 = 1000
        fmt.Println(v2)
        _ = fi9
        _ = v10
    }
    v6 += 18446744073709551615
    for fi11 := 0; fi11 < 20; fi11++ {
        if v2 {
            var v12 string = ("hello" + "\n\t")
            switch (v12 + v12) {
            case "\n\t":
                v13 := []int{544}
                _ = v13
            case "a":
                v6 = v6
            default:
                switch v12 {
                case "hello":
                    if v2 {
                        v14 := GenS1{F0: 1, F1: fi11, F2: (fi11 ^ fi11)}
                        _ = v14
                    } else {
                        if v2 {
                            v6 = 65535
                            v4 = (v4 * (v6 << ((v4) & 63)))
                        }
                        v12 = ((v12 + (v12 + ("a" + "\n\t"))) + v12)
                        var v15 uint8 = 127
                        _ = v15
                    }
                default:
                    var v16 uint32 = 42
                    _ = v16
                }
            }
            _ = v12
        }
        _ = fi11
    }
    v4 = 314
    if (624 >= 18446744073709551615) {
        v2 = v2
        var v17 float32 = 1.5
        v17 = v17
        _ = v17
    }
    if v2 {
        v18 := GenS2{F0: -1}
        fmt.Println(v1, v4, v4)
        switch "a" {
        case "":
            fmt.Println(v1)
        default:
            if false {
                for fi19 := 0; fi19 < 10; fi19++ {
                    _ = len(v3)
                    _ = fi19
                }
            }
        }
        _ = v18
    } else {
        switch 128 {
        case -858:
            fmt.Println(v2)
        }
        switch "a" {
        case "a":
            v20 := GenS2{F0: 32767}
            _ = v20
        case "\n\t":
            v6 = (v4 + v4)
        case "foo bar":
            v21 := GenS3{F0: -1, F1: 2, F2: -1}
            _ = v21
        }
        var v22 uint = 42
        _ = v22
    }
    v2 = (!(32767 != 0))
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
