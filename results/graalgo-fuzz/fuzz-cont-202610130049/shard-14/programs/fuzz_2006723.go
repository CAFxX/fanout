package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 809
    var v3 int32 = 2147483647
    switch 440 {
    case 128:
        v4 := GenS1{F0: ((("foo bar" + ("" + "")) + "\n\t") + (("a" + "foo bar") + "a"))}
        _ = v4
    }
    v5 := GenS1{F0: "hello"}
    v3 -= (v3 << ((553) & 63))
    v6 := []int{2147483647}
    v6 = append(v6, 127)
    var v7 uint8 = 255
    switch ("" + "\n\t") {
    case "":
        v2 -= v2
    case "a":
        var v8 uint64 = (v2 & v2)
        _ = v8
    case "\n\t":
        var v9 bool = false
        _ = v9
    default:
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(v2, v3)
                v2 = (v2 & v2)
                v3 = -2147483648
                _ = fi10
            }
            var v11 int8 = -1
            fmt.Println(v2)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                delete(v1, 4294967295)
                var v13 uint16 = 379
                for fi14 := 0; fi14 < 10; fi14++ {
                    v7 = v7
                    _ = fi14
                }
                _ = fi12
                _ = v13
            }
        }
    }
    v1[776] = 42
    fmt.Println(v2)
    switch "" {
    case "\n\t":
        if (("a" + "a") >= "\n\t") {
            var v15 int = 42
            _ = v15
        }
    case "foo bar":
        if true {
            v3 = -416
        } else {
            v7 -= (v7 | v7)
            switch "foo bar" {
            case "foo bar":
                var v16 int64 = 1
                _ = v16
            case "hello":
                v1[2] = 2147483647
            case "a":
                v17 := GenS2{F0: 42, F1: (^(v7)), F2: 42}
                _ = v17
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v2 = v2
                    fmt.Println(v7, v3, v3)
                    _ = fi19
                }
                if false {
                    v2 += (^(v2))
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 uint8 = v7
                        var v22 int32 = -1
                        v21 = 1
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                }
                v23 := v1[2147483648]
                _ = fi18
                _ = v23
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
