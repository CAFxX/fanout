package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    v1 := []uint32{4294967295, 2147483647, 65535}
    var v2 uint32 = 2147483647
    var v3 uint16 = 65535
    if ("foo bar" < "\n\t") {
        if false {
            var v4 uint16 = v3
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v4
        } else {
            for fi5 := 0; fi5 < 19; fi5++ {
                var v6 int16 = 42
                v6 = (^(v6))
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: 128, F1: 128, F2: 42}
            _ = v7
        }
        v3 += (v3 - 657)
    } else {
        var v8 int64 = 42
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                var v11 uint64 = 68
                v11 = v11
                _ = v11
            } else {
                v3 += v3
                v1 = append(v1, (v2 ^ v2))
            }
            _ = fi10
        }
        _ = fi9
    }
    switch ("hello" + "") {
    case "":
        if len(v1) > 0 {
            v1[0] = 65535
        }
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 2, F1: (fi12 * (fi12 >> ((255) & 63))), F2: 42}
        fmt.Println(v3, v3, fi12)
        _ = fi12
        _ = v13
    }
    v2 = v2
    v3 -= (v3 + v3)
    if false {
        switch "hello" {
        case "":
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    if true {
                        var v16 uint16 = (v3 * uint16(v2))
                        fmt.Println(v16, v16, v16)
                        var v17 uint32 = 42
                        _ = v16
                        _ = v17
                    }
                    _ = fi15
                }
                _ = fi14
            }
        case "hello":
            var v18 string = "foo bar"
            _ = v18
        case "\n\t":
            v19 := GenS1{F0: 1, F1: 9223372036854775807, F2: -97}
            _ = v19
        }
    }
    var v20 string = "hello"
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(v2)
            _ = fi22
        }
        v3 = uint16(v2)
        _ = fi21
    }
    switch (v20 + (v20 + v20)) {
    case "\n\t":
        var v23 int = -1
        _ = v23
    case "hello":
        v20 = v20
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS1{F0: 42, F1: int(v3), F2: -1}
        v2 += v2
        v26 := GenS1{F0: 128, F1: fi24, F2: 20}
        _ = fi24
        _ = v25
        _ = v26
    }
    switch 167 {
    case -9223372036854775808:
        switch 42 {
        case 516:
            for fi27 := 0; fi27 < 13; fi27++ {
                v20 = v20
                fmt.Println(v20, v3)
                _ = fi27
            }
        case -721:
            for fi28 := 0; fi28 < 18; fi28++ {
                if false {
                    if true {
                        var v29 float32 = (0.1 / 1.0)
                        fmt.Println(v3)
                        _ = v29
                    } else {
                        v2 = v2
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS2{F0: (v3 & (v3 << ((1) & 63))), F1: "\n\t"}
                            _ = fi30
                            _ = v31
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 int = (fi32 << ((339) & 63))
                        _ = fi32
                        _ = v33
                    }
                }
                v20 = (v20 + ("a" + (("\n\t" + "\n\t") + v20)))
                _ = fi28
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
