package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 uint64 }

func main() {
    var v1 int64 = 42
    var v2 uint = 927
    var v3 uint32 = 2
    var v4 uint = (^(v2))
    var v5 uint8 = uint8(v4)
    v6 := make(map[string]string)
    v4 -= 0
    switch "hello" {
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            switch (^(fi7)) {
            case 42:
                var v8 uint = v2
                _ = v8
            case -556:
                v9 := GenS1{F0: 4294967295}
                _ = v9
            case -25:
                var v10 int8 = -1
                _ = v10
            default:
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := []int{int(v2), (fi7 | fi7), fi7}
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi7
        }
    case "foo bar":
        switch 440 {
        case 32767:
            v13 := GenS1{F0: 220}
            _ = v13
        case -803:
            var v14 int64 = v1
            _ = v14
        default:
            v15 := GenS1{F0: 42}
            _ = v15
        }
    }
    fmt.Println(v5)
    v4 = 571
    v16 := []int{127, -749}
    _ = len(v6)
    var v17 int8 = 42
    _ = len(v6)
    v16 = append(v16, -2147483648)
    var v18 uint16 = 0
    for fi19 := 0; fi19 < 9; fi19++ {
        v3 = 65535
        switch ("foo bar" + "\n\t") {
        case "a":
            for fi20 := 0; fi20 < 7; fi20++ {
                switch 32767 {
                case 144:
                    var v21 string = "foo bar"
                    _ = v21
                }
                _ = fi20
            }
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS1{F0: 811}
            var v24 uint64 = 42
            v25 := GenS1{F0: (v24 << ((2) & 63))}
            v26 := GenS2{F0: v1, F1: 32768, F2: 960}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
