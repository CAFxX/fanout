package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint8 = 128
    v2 := []uint32{2147483647, 2147483647}
    var v3 int8 = 127
    var v4 uint = 1
    v4 = 255
    var v5 uint = v4
    if false {
        if ((false || true) && (0.1 <= 0.0)) {
            v2 = append(v2, 1)
            var v6 int64 = -9223372036854775808
            _ = v6
        }
        v3 = (v3 & (v3 ^ v3))
    }
    var v7 int64 = -406
    fmt.Println(v3)
    if true {
        var v8 int8 = v3
        var v9 int = -614
        for fi10 := 0; fi10 < 12; fi10++ {
            if (32767 <= -2147483648) {
                if ((-(v7)) >= (v7 ^ (v7 | (v7 & 42)))) {
                    v9 = int(v5)
                    for fi11 := 0; fi11 < 19; fi11++ {
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            if true {
                                if false {
                                    v8 = v8
                                }
                                fmt.Println(v5)
                            } else {
                                v12 := GenS1{F0: int16(v7)}
                                var v13 uint16 = 42
                                fmt.Println(v5)
                                _ = v12
                                _ = v13
                            }
                        case "foo bar":
                            v14 := GenS2{F0: 781}
                            _ = v14
                        case "a":
                            var v15 int16 = 383
                            _ = v15
                        }
                        _ = fi11
                    }
                    var v16 uint8 = 128
                    _ = v16
                } else {
                    v9 = -1
                }
                var v17 int64 = v7
                _ = v17
            }
            var v18 uint64 = 546
            var v19 float32 = float32((float64(3.14159) + (((3.14159 / 1.5) - (1.5 / 1.0)) + 0.1)))
            _ = fi10
            _ = v18
            _ = v19
        }
        _ = v8
        _ = v9
    }
    var v20 uint8 = v1
    if ((1e10 - float32((0.1 + 1e10))) < float32((1.5 + float64((1e10 + float32(3.14159)))))) {
        v21 := GenS2{F0: 27}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v5 = (v4 ^ (v4 % (((^(v5))) | 1)))
                v5 = (^(v5))
                _ = fi23
            }
            _ = fi22
        }
    }
    var v24 uint = (v4 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
