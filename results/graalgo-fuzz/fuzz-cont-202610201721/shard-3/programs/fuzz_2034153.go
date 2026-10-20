package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 string = v1
    var v4 string = v3
    if false {
        switch "foo bar" {
        case "a":
            for fi5 := 0; fi5 < 1; fi5++ {
                switch fi5 {
                case -202:
                    v6 := GenS1{F0: 1, F1: (fi5 * fi5), F2: 1}
                    _ = v6
                case 4294967295:
                    v3 = v3
                case -9223372036854775808:
                    v7 := GenS2{F0: 1}
                    _ = v7
                default:
                    v4 = v4
                }
                _ = fi5
            }
        }
        var v8 bool = true
        fmt.Println(v3, v8)
        _ = v8
    } else {
        var v9 int = 32767
        var v10 uint32 = 522
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v3 = (v3 + ("a" + "foo bar"))
            v13 := GenS3{F0: 255}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch 2147483647 {
    case 42:
        if true {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int = 127
                var v16 int16 = 1000
                var v17 int16 = (v16 & -210)
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v4 = ((v3 + "\n\t") + v1)
        }
    default:
        var v18 int = 0
        _ = v18
    }
    v19 := GenS4{F0: 42}
    fmt.Println(v1, v3)
    v20 := GenS4{F0: 18446744073709551615}
    if false {
        v3 = ("hello" + ((v1 + "hello") + "hello"))
    }
    v3 = "foo bar"
    for fi21 := 0; fi21 < 1; fi21++ {
        fmt.Println(v1)
        v22 := GenS2{F0: 806}
        var v23 int16 = 42
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        var v24 uint16 = 65535
        _ = v24
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
