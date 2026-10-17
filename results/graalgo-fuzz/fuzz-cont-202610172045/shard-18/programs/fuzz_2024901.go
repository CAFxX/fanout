package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 uint, p1 uint8, p2 int16) bool {
    fmt.Println(p1)
    switch 2147483647 {
    case 42:
        var v5 int32 = 1
        _ = v5
    case 928:
        p1 = 42
    }
    var v6 uint = (p0 >> (((p0 % ((2) | 1))) & 63))
    var v7 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!(false || (p2 < -972)))
}

func main() {
    var v1 uint32 = 2
    var v2 uint = 385
    v3 := []string{}
    var v4 string = "\n\t"
    v8 := GenS1{F0: 1, F1: -9223372036854775808, F2: -389}
    v9 := GenS1{F0: -824, F1: -9223372036854775808, F2: -740}
    var v10 uint = 656
    var v11 int32 = -707
    switch -9223372036854775808 {
    case 32768:
        if ((-(v11)) < v11) {
            _ = len(v3)
            for fi12 := 0; fi12 < 13; fi12++ {
                v1 = uint32(v10)
                if ((fi12 % ((32767) | 1)) > fi12) {
                    v3 = append(v3, v4)
                    if false {
                        switch (fi12 >> ((1) & 63)) {
                        case -492:
                            _ = len(v3)
                        case -905:
                            for fi13 := 0; fi13 < 4; fi13++ {
                                var v14 int8 = 0
                                _ = fi13
                                _ = v14
                            }
                        case -2147483648:
                            v15 := GenS1{F0: 1, F1: (fi12 / (((-(fi12))) | 1)), F2: (^(fi12))}
                            _ = v15
                        }
                    }
                } else {
                    if true {
                        if (false && true) {
                            if ((true && false) && (true || ((false && true) && true))) {
                                v16 := GenS1{F0: -9223372036854775808, F1: fi12, F2: (fi12 << (((v2 + 1)) & 63))}
                                var v17 string = ("\n\t" + v4)
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(v11)
                            v1 += 42
                        } else {
                            var v18 float64 = float64(-2.25)
                            _ = v18
                        }
                        fmt.Println(fi12, v2)
                        for fi19 := 0; fi19 < 20; fi19++ {
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                            v20 := GenS2{F0: (^(v1)), F1: fi19}
                            if (true || false) {
                                fmt.Println(v1)
                                v21 := GenS1{F0: 0, F1: 128, F2: fi12}
                                v11 += -2147483648
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                    }
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(v11)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi22
                    }
                    switch fi12 {
                    case -2147483648:
                        v10 = v2
                    case 304:
                        var v23 int = fi12
                        _ = v23
                    default:
                        v2 = (v10 >> (((v2 ^ (v10 >> ((v10) & 63)))) & 63))
                    }
                }
                _ = fi12
            }
        }
    }
    v2 = v2
    switch 0 {
    case -1:
        v24 := GenS1{F0: 42, F1: -9223372036854775808, F2: -2147483648}
        _ = v24
    case 42:
        var v25 float32 = float32(0.0)
        _ = v25
    }
    var v26 bool = (!(1.5 != ((0.0 - float64(1.5)) - 0.1)))
    v27 := GenS3{F0: -308}
    for fi28 := 0; fi28 < 12; fi28++ {
        v26 = v26
        fmt.Println(v1, v1, v26)
        v29 := GenS2{F0: v1, F1: -532}
        _ = fi28
        _ = v29
    }
    switch 128 {
    case -851:
        if true {
            if v26 {
                v26 = true
                fmt.Println(v2)
            }
            v2 = (v2 << ((v2) & 63))
        } else {
            for fi30 := 0; fi30 < 1; fi30++ {
                switch (fi30 * int(v1)) {
                case 42:
                    var v31 float32 = 1e10
                    _ = v31
                case -86:
                    fmt.Println(v2, v26, v1)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi30
            }
        }
    case 4294967295:
        v26 = (4294967295 >= (v10 * (v2 >> (((^(v10))) & 63))))
    case 42:
        if len(v3) > 0 {
            v3[0] = (v4 + v4)
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v32 := GenS4{F0: (!((!v26) && true)), F1: 294, F2: -73}
    v10 += (v10 | v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
