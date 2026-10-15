package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }

func main() {
    var v1 uint8 = 127
    v2 := []string{("\n\t" + "")}
    var v3 uint8 = v1
    var v4 int32 = 1
    var v5 uint16 = 765
    v2 = append(v2, "foo bar")
    for fi6 := 0; fi6 < 18; fi6++ {
        if false {
            if false {
                if true {
                    var v7 int8 = 42
                    _ = v7
                }
                for fi8 := 0; fi8 < 17; fi8++ {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    _ = fi8
                }
            }
        } else {
            v9 := GenS1{F0: "\n\t"}
            _ = v9
        }
        v10 := GenS2{F0: -28}
        _ = fi6
        _ = v10
    }
    v4 += v4
    switch 32767 {
    case 9223372036854775807:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    if ((^(v4)) == -2147483648) {
        v5 -= (v5 * v5)
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int = 857
            v5 += (v5 | v5)
            _ = fi11
            _ = v12
        }
        var v13 uint = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        _ = len(v2)
        v2 = append(v2, "")
        var v15 bool = (255 > 42)
        _ = fi14
        _ = v15
    }
    if (!(int8(v4) >= -1)) {
        var v16 int = 32768
        var v17 int32 = (v4 | v4)
        _ = v16
        _ = v17
    } else {
        v1 -= v1
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: 2147483647, F1: 42, F2: false}
            v20 := GenS2{F0: int8(v5)}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
