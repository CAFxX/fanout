package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
func genFn0(p0 uint) int16 {
    p0 -= 337
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int64 = -644
        _ = fi5
        _ = v6
    }
    _ = p0
    return -32768
}

func main() {
    var v1 float32 = (1.5 - -2.25)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int64 = -1
    for fi7 := 0; fi7 < 1; fi7++ {
        v4 = -362
        v4 = 42
        _ = fi7
    }
    var v8 uint8 = 255
    v9 := GenS1{F0: false, F1: -578}
    var v10 float64 = (-2.25 + 3.14159)
    for fi11 := 0; fi11 < 8; fi11++ {
        v4 += 678
        fmt.Println(v4, fi11)
        if false {
            v1 = (((v1 + ((((((v1 - v1) + 0.1) + v1) + (((0.0 + (v1 / 1.0)) / v1) - (v1 - 1.5))) * (v1 - (0.1 - 0.1))) - (0.1 + v1))) * (float32(v10) - v1)) * v1)
        }
        _ = fi11
    }
    var v12 uint8 = 0
    var v13 int32 = 227
    v14 := []int{2147483647, 32768}
    if (!((!("hello" == ("hello" + "foo bar"))) && (!((127 >= 1) && (v12 <= 131))))) {
        fmt.Println(v13, v12, v1)
        v14 = append(v14, 127)
        var v15 uint = 941
        _ = v15
    } else {
        switch "foo bar" {
        case "a":
            var v16 uint8 = (v8 + v12)
            _ = v16
        case "":
            v17 := GenS1{F0: (true && false), F1: -171}
            _ = v17
        case "foo bar":
            if true {
                if (false || true) {
                    switch 32768 {
                    case -575:
                        v13 += (v13 & v13)
                    case 32767:
                        for fi18 := 0; fi18 < 14; fi18++ {
                            if true {
                                switch ("\n\t" + "hello") {
                                case "a":
                                    var v19 float32 = (v1 / v1)
                                    _ = v19
                                case "\n\t":
                                    var v20 uint16 = 42
                                    _ = v20
                                case "hello":
                                    for fi21 := 0; fi21 < 12; fi21++ {
                                        v10 = 1e10
                                        _ = fi21
                                    }
                                default:
                                    var v22 int8 = 122
                                    _ = v22
                                }
                                v23 := GenS2{F0: -86, F1: 255, F2: 2147483648}
                                _ = v23
                            }
                            if (-285 < genFn0(390)) {
                                _ = len(v14)
                            } else {
                                var v24 uint16 = 713
                                _ = v24
                            }
                            _ = fi18
                        }
                    case -2147483648:
                        v8 = (^(v12))
                    }
                }
            } else {
                fmt.Println(v12, v1)
                v25 := GenS2{F0: 127, F1: 316, F2: 65535}
                _ = v25
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
