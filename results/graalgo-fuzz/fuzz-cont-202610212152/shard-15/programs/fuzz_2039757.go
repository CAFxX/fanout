package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    v3 := []string{"\n\t", "a", "", "hello"}
    var v4 uint32 = 2
    v4 += (v4 - v4)
    switch -243 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 7; fi5++ {
            switch (((("" + "hello") + "a") + "a") + (("hello" + "") + ("hello" + ("a" + ("hello" + "hello"))))) {
            case "\n\t":
                v1 = (v1 ^ 2147483647)
            }
            _ = fi5
        }
    case 893:
        v6 := GenS1{F0: 543, F1: "", F2: 1}
        _ = v6
    default:
        var v7 uint8 = 127
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int = (fi8 / ((-58) | 1))
        if ((v4 << ((850) & 63)) != v4) {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint16 = 212
                _ = fi10
                _ = v11
            }
        }
        switch (fi8 >> ((529) & 63)) {
        case 264:
            var v12 float64 = (-2.25 - 1.5)
            _ = v12
        case 32768:
            v13 := GenS1{F0: 255, F1: ("a" + "a"), F2: -47}
            _ = v13
        default:
            var v14 int64 = -1
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int32 = -1
        _ = fi15
        _ = v16
    }
    var v17 bool = true
    v2[65535] = true
    v1 -= 0
    v1 += v1
    v4 = v4
    var v18 uint64 = 1
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: 483, F1: 32767, F2: 151}
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 += 42
            fmt.Println(v18, fi21)
            if false {
                v3 = append(v3, (("hello" + (("\n\t" + "foo bar") + "foo bar")) + ""))
                switch (fi19 | 1) {
                case 0:
                    var v22 uint32 = (v4 - v4)
                    if true {
                        var v23 int32 = (v1 * v1)
                        switch (fi21 - fi21) {
                        case 9223372036854775807:
                            var v24 int32 = (v1 & 1)
                            _ = v24
                        default:
                            v4 = 935
                        }
                        v23 -= (v23 + (v23 * v23))
                        _ = v23
                    }
                    v22 = v4
                    _ = v22
                }
            }
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
