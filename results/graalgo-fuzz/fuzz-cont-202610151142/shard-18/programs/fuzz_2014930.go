package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0(p0 uint) int32 {
    p0 -= (^(p0))
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1() int8 {
    var v7 uint16 = 32767
    fmt.Println(v7)
    fmt.Println(v7)
    _ = v7
    return -128
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 999
    var v3 uint = 1
    v4 := make(map[uint32]int)
    v5 := []uint32{2147483648, 348, 0, 121}
    var v6 uint8 = 127
    v1 = v1
    if len(v5) > 0 {
        v5[0] = uint32(v2)
    }
    switch 4294967295 {
    case -2147483648:
        if (v3 > v3) {
            if (genFn1() == 0) {
                v6 += 42
                v1 = v1
            }
            v8 := GenS1{F0: 1, F1: ""}
            _ = v8
        } else {
            var v9 bool = ((v2 != (v2 * (v2 | (v2 << ((v3) & 63))))) && ((2147483647 <= 2) && (!false)))
            _ = v9
        }
    default:
        if true {
            for fi10 := 0; fi10 < 2; fi10++ {
                if true {
                    v3 = (^(v3))
                } else {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v6 -= v6
                        var v12 uint16 = (^(v2))
                        var v13 int16 = 579
                        fmt.Println(v13, v1)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 1, F1: (("foo bar" + "") + v1)}
                    var v15 uint = v3
                    _ = v14
                    _ = v15
                }
                _ = fi10
            }
        }
    }
    if (v1 >= ("hello" + ("foo bar" + v1))) {
        fmt.Println(v1, v2, v2)
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    }
    v16 := GenS1{F0: -554, F1: "a"}
    if len(v5) > 0 {
        v5[0] = 743
    }
    switch 4294967295 {
    case -1:
        if len(v5) > 0 {
            v5[0] = uint32(v2)
        }
    }
    var v17 uint = 255
    fmt.Println(v6, v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 string = "hello"
        v20 := GenS1{F0: 0, F1: ""}
        if (false || (!false)) {
            for fi21 := 0; fi21 < 17; fi21++ {
                v19 = ("" + ("" + ("foo bar" + "a")))
                switch v1 {
                case "foo bar":
                    v2 += v2
                }
                _ = fi21
            }
            var v22 int8 = 0
            _ = v22
        } else {
            v3 -= v17
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
