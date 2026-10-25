package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 bool = false
    var v2 int32 = 0
    v3 := []int{0, 42}
    var v4 int = 367
    if (!(!v1)) {
        var v5 string = ("hello" + ("" + ("hello" + "\n\t")))
        _ = len(v3)
        v6 := GenS1{F0: v1}
        _ = v5
        _ = v6
    } else {
        fmt.Println(v2, v2)
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        switch ("a" + "foo bar") {
        case "hello":
            var v8 uint32 = 2147483648
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: true}
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(fi7)
            _ = fi10
        }
        _ = fi7
    }
    _ = len(v3)
    v11 := GenS2{F0: 2, F1: 65535}
    switch (v4 + v4) {
    case -969:
        v1 = v1
    case -2147483648:
        var v12 uint8 = 0
        _ = v12
    case 118:
        v13 := GenS3{F0: v1, F1: -9223372036854775808, F2: (^(v4))}
        _ = v13
    default:
        if v1 {
            switch "foo bar" {
            case "\n\t":
                for fi14 := 0; fi14 < 13; fi14++ {
                    v3 = append(v3, v4)
                    v15 := GenS2{F0: 2147483647, F1: 42}
                    var v16 string = ""
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "foo bar":
                v17 := GenS3{F0: false, F1: 898, F2: (-(v4))}
                _ = v17
            case "":
                v4 -= v4
            }
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            if false {
                switch (fi19 * fi18) {
                case 0:
                    var v20 int16 = -770
                    _ = v20
                default:
                    switch (fi19 & 42) {
                    case 42:
                        _ = len(v3)
                    }
                }
                var v21 uint16 = 65535
                _ = v21
            }
            var v22 int16 = 1
            v23 := GenS4{F0: 42, F1: 32767}
            _ = fi19
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(v2, v4, fi24)
        var v25 uint = 65535
        _ = fi24
        _ = v25
    }
    var v26 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
