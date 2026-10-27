package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0() uint64 {
    var v5 uint8 = 127
    v5 = v5
    var v6 int64 = -18
    switch (("hello" + "foo bar") + "\n\t") {
    case "":
        var v7 int64 = 801
        _ = v7
    }
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{440, 42, 2147483648, 0}
    var v3 uint32 = 65535
    v4 := []int{}
    var v8 int16 = -799
    switch ("foo bar" + "\n\t") {
    case "a":
        if (!(false && true)) {
            switch "foo bar" {
            case "hello":
                if (float32(((1e10 - 1e10) - -2.25)) >= (0.1 - 1.5)) {
                    _ = len(v4)
                } else {
                    var v9 int32 = -197
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 int16 = v8
                        switch (fi10 & fi10) {
                        case 0:
                            for fi12 := 0; fi12 < 7; fi12++ {
                                var v13 int64 = v1
                                fmt.Println(v8, v11)
                                _ = fi12
                                _ = v13
                            }
                        case 32768:
                            var v14 uint64 = 372
                            _ = v14
                        case 128:
                            if false {
                                var v15 int8 = 89
                                v15 = 1
                                var v16 uint64 = 2
                                _ = v15
                                _ = v16
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
            case "foo bar":
                fmt.Println(v3, v1, v3)
            default:
                fmt.Println(v3, v8, v8)
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "":
        v4 = append(v4, 933)
    case "foo bar":
        _ = len(v4)
    default:
        var v17 float32 = -2.25
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS1{F0: 42}
        v3 = (v3 | v3)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        switch "" {
        case "":
            v21 := GenS1{F0: 168}
            _ = v21
        case "a":
            v1 += v1
        case "\n\t":
            v22 := GenS2{F0: false}
            _ = v22
        default:
            var v23 int = (fi20 + (fi20 | fi20))
            _ = v23
        }
        switch fi20 {
        case 127:
            var v24 uint64 = 0
            _ = v24
        default:
            if (false && true) {
                switch "foo bar" {
                case "a":
                    fmt.Println(v3, v3, v1)
                }
                var v25 float32 = -2.25
                _ = v25
            } else {
                switch "" {
                case "\n\t":
                    v1 = v1
                case "hello":
                    if (0.1 <= (0.1 * 0.1)) {
                        v1 -= (v1 >> ((4294967295) & 63))
                    }
                case "":
                    v1 -= -1
                }
            }
        }
        _ = fi20
    }
    v26 := GenS1{F0: 1}
    var v27 int32 = 2147483647
    var v28 bool = ((!(true && (1 > genFn0()))) && (1 < -118))
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            _ = len(v2)
            v1 -= (v1 ^ int64(fi30))
            v8 = v8
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
