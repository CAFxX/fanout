package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = 1
    var v2 int8 = 1
    v3 := []int{-323, 643, 2147483647}
    var v4 float32 = -2.25
    var v5 uint32 = 2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v5 -= 795
            _ = fi7
        }
        v1 = v1
        var v8 int = int(v2)
        _ = fi6
        _ = v8
    }
    switch ("a" + "") {
    case "foo bar":
        v1 += -1
    case "a":
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v4)
            if true {
                if (!(!false)) {
                    fmt.Println(fi9, v5)
                    var v10 int16 = -167
                    v5 -= (v5 + v5)
                    _ = v10
                } else {
                    var v11 bool = true
                    for fi12 := 0; fi12 < 3; fi12++ {
                        fmt.Println(v1)
                        var v13 int = int(v5)
                        _ = fi12
                        _ = v13
                    }
                    if len(v3) > 0 {
                        v3[0] = (fi9 - int(v2))
                    }
                    _ = v11
                }
                if (!true) {
                    if false {
                        if false {
                            v14 := GenS1{F0: 2, F1: 127, F2: 0}
                            switch "foo bar" {
                            case "foo bar":
                                var v15 int64 = 9223372036854775807
                                _ = v15
                            case "a":
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 32768
                                    }
                                }
                            }
                            _ = v14
                        }
                    }
                    v5 -= v5
                }
                v5 += 2
            } else {
                v5 = 1
                v16 := GenS1{F0: 42, F1: 255, F2: -32768}
                var v17 int = (fi9 ^ fi9)
                _ = v16
                _ = v17
            }
            _ = fi9
        }
    case "hello":
        v4 += -2.25
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    case "a":
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: 42, F1: 1, F2: int16(v2)}
            _ = fi18
            _ = v19
        }
    }
    if true {
        var v20 uint32 = (v5 + v5)
        var v21 uint = 42
        v22 := GenS2{F0: (v2 | (v2 * 1))}
        _ = v20
        _ = v21
        _ = v22
    } else {
        var v23 bool = true
        _ = v23
    }
    v2 = (v2 << ((4294967295) & 63))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
