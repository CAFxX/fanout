package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint }
func genFn0(p0 int) uint16 {
    v4 := []int{(p0 * -2147483648)}
    v4 = append(v4, (^(p0)))
    var v5 int16 = 1
    v5 += v5
    _ = p0
    _ = v4
    _ = v5
    return 820
}
func genFn1(p0 int32, p1 int64, p2 uint32) bool {
    switch "" {
    case "hello":
        p0 = (p0 / ((-316) | 1))
    default:
        p2 = 42
    }
    v6 := []int{}
    if ((false && (!(747 == 902))) || (!true)) {
        for fi7 := 0; fi7 < 18; fi7++ {
            p1 -= 0
            var v8 uint = 65535
            _ = fi7
            _ = v8
        }
        switch -2147483648 {
        case -37:
            switch int(p1) {
            case 127:
                switch int(p0) {
                case -186:
                    if true {
                        if (0 >= 65535) {
                            var v9 uint8 = 127
                            for fi10 := 0; fi10 < 3; fi10++ {
                                _ = len(v6)
                                _ = fi10
                            }
                            p0 = (p0 - 666)
                            _ = v9
                        }
                        v11 := GenS1{F0: 128}
                        p1 = (p1 << ((18446744073709551615) & 63))
                        _ = v11
                    } else {
                        var v12 uint = 384
                        for fi13 := 0; fi13 < 10; fi13++ {
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = (fi13 / ((32767) | 1))
                                }
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    if ((v12 > (v12 ^ v12)) || true) {
                                        p2 = 65535
                                        v12 -= (v12 & v12)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    } else {
                                        switch fi13 {
                                        case 9223372036854775807:
                                            var v15 int32 = p0
                                            _ = v15
                                        case 127:
                                            p1 += (p1 << (((^(v12))) & 63))
                                        default:
                                            if (!true) {
                                                p2 = p2
                                            }
                                        }
                                    }
                                    var v16 uint64 = 18446744073709551615
                                    v17 := GenS2{F0: (p1 / ((214) | 1))}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi14
                                    _ = v16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    v19 := GenS3{F0: -33, F1: 1}
                                    var v20 float32 = 0.1
                                    fmt.Println(fi18)
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            }
                            _ = fi13
                        }
                        v12 -= 57
                        _ = v12
                    }
                }
            case 60:
                var v21 string = "\n\t"
                _ = v21
            case -1:
                var v22 bool = (false || false)
                _ = v22
            }
        case 42:
            fmt.Println(p0)
        case 0:
            v6 = append(v6, 32767)
        default:
            var v23 uint8 = 127
            _ = v23
        }
    } else {
        switch -592 {
        case 42:
            if len(v6) > 0 {
                v6[0] = 32767
            }
        case -478:
            for fi24 := 0; fi24 < 19; fi24++ {
                if len(v6) > 0 {
                    v6[0] = (fi24 / ((fi24) | 1))
                }
                fmt.Println(p0, p2, p2)
                var v25 int64 = p1
                _ = fi24
                _ = v25
            }
        }
        v26 := GenS4{F0: 0, F1: p2, F2: 119}
        fmt.Println(p2, p1)
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    var v1 uint8 = 108
    var v2 int64 = 467
    v3 := make(map[uint32]int)
    var v27 int32 = 2147483647
    var v28 int64 = 9223372036854775807
    if false {
        v28 -= 844
    }
    v29 := []int{int(v27)}
    v28 -= v28
    fmt.Println(v28, v1)
    v28 = (v2 >> ((294) & 63))
    if false {
        var v30 int8 = -1
        var v31 uint64 = 4294967295
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
