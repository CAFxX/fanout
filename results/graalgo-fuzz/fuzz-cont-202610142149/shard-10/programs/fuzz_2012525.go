package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint8, p1 uint, p2 int64) uint {
    fmt.Println(p0)
    var v5 int16 = -657
    var v6 uint = 0
    var v7 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 303
}

func main() {
    v1 := []int{-765, -2147483648}
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 string = v2
    v4 = v2
    v8 := GenS1{F0: 4294967295}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        var v9 uint32 = 1
        _ = v9
    } else {
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint = 255
                switch fi10 {
                case 128:
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: 18446744073709551615}
                        var v14 string = v4
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 2147483647:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 float64 = 0.0
                        if (1 <= 0) {
                            fmt.Println(v4)
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 9223372036854775807
                            }
                        }
                        v2 = (v2 + v4)
                        _ = fi15
                        _ = v16
                    }
                case 428:
                    v17 := GenS1{F0: 109}
                    _ = v17
                default:
                    v2 = "hello"
                }
                v18 := GenS1{F0: 2}
                _ = fi10
                _ = v11
                _ = v18
            }
            v19 := GenS1{F0: 0}
            _ = v19
        }
    }
    v20 := GenS1{F0: 42}
    var v21 int = 32768
    var v22 uint = 0
    switch ("foo bar" + ((v2 + "foo bar") + "\n\t")) {
    case "":
        if (false && (!true)) {
            var v23 int = v21
            var v24 uint8 = 127
            switch v23 {
            case -110:
                v25 := GenS2{F0: 32767, F1: 65535, F2: (^(v22))}
                _ = v25
            case 2147483647:
                _ = len(v1)
            case 1:
                v23 = (-(v21))
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v23
            _ = v24
        } else {
            fmt.Println(v4)
            var v26 uint64 = 1
            fmt.Println(v21, v26, v4)
            _ = v26
        }
    case "a":
        var v27 float32 = 1.5
        _ = v27
    }
    _ = len(v3)
    for fi28 := 0; fi28 < 10; fi28++ {
        v21 = fi28
        _ = fi28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
