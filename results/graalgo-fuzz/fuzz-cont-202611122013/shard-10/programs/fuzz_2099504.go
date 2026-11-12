package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-292}
    v3 := make(map[uint32]bool)
    var v4 uint = 837
    var v5 uint8 = 109
    var v6 bool = (true || true)
    var v7 uint8 = (^(v5))
    v8 := GenS1{F0: 65535, F1: 0, F2: v6}
    var v9 uint64 = 1
    _ = len(v3)
    for fi10 := 0; fi10 < 3; fi10++ {
        delete(v1, 127)
        fmt.Println(v9, v9)
        _ = fi10
    }
    v6 = (true && false)
    v4 = (v4 >> ((255) & 63))
    v11 := GenS1{F0: 42, F1: (v9 % (((v9 / (((v9 + (v9 + (v9 - (v9 & v9))))) | 1))) | 1)), F2: (v6 || (!((v5 % ((v7) | 1)) > v5)))}
    v12 := v1[0]
    v7 = v7
    if ((!v6) && v6) {
        v13 := GenS1{F0: 763, F1: (v9 & (^(v9))), F2: (!v6)}
        if (!false) {
            v14 := GenS1{F0: 32767, F1: (^(v9)), F2: false}
            if (!(!true)) {
                v15 := GenS1{F0: 0, F1: (^(v9)), F2: (0.0 == 0.1)}
                _ = v15
            } else {
                v6 = v6
                v5 -= (v5 << ((4294967295) & 63))
            }
            _ = v14
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v16 := GenS1{F0: 255, F1: 4294967295, F2: ("a" > "a")}
            for fi17 := 0; fi17 < 17; fi17++ {
                for fi18 := 0; fi18 < 1; fi18++ {
                    switch "hello" {
                    case "a":
                        v19 := GenS1{F0: 42, F1: 0, F2: ((false || v6) && v6)}
                        _ = v19
                    default:
                        if true {
                            v4 -= v4
                            v20 := v3[4294967295]
                            var v21 int16 = 0
                            v22 := GenS1{F0: 2, F1: v9, F2: (float64(-2.25) < 0.0)}
                            v5 += v7
                            v6 = v6
                            if (1 >= 2147483647) {
                                var v23 int64 = 42
                                var v24 float64 = (0.0 / 1.5)
                                _ = v23
                                _ = v24
                            } else {
                                v5 += (v7 << ((255) & 63))
                            }
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v25 := GenS1{F0: 2, F1: v9, F2: false}
                            var v26 int32 = 2147483647
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi18
                }
                if (42 == 32767) {
                    switch "a" {
                    case "\n\t":
                        v4 += v4
                    default:
                        v9 = v9
                    }
                    if v6 {
                        v3[uint32(v5)] = (v6 && false)
                        _ = len(v1)
                    } else {
                        var v27 int16 = 536
                        if v6 {
                            v28 := GenS2{F0: -261}
                            _ = v28
                        } else {
                            switch 128 {
                            case 39:
                                var v29 float64 = 0.1
                                _ = v29
                            default:
                                var v30 uint16 = 255
                                _ = v30
                            }
                        }
                        _ = v27
                    }
                    var v31 uint8 = v5
                    _ = v31
                } else {
                    fmt.Println(v7)
                    v1[65535] = (fi17 - v12)
                    var v32 string = "foo bar"
                    _ = v32
                }
                for fi33 := 0; fi33 < 19; fi33++ {
                    v34 := GenS1{F0: 32767, F1: (v9 & (v9 + (v9 + 18446744073709551615))), F2: false}
                    fmt.Println(v6)
                    _ = fi33
                    _ = v34
                }
                _ = fi17
            }
            _ = v16
        }
        _ = v13
    } else {
        v9 -= (v9 << ((18446744073709551615) & 63))
        _ = len(v3)
        v2 = append(v2, v12)
    }
    v35 := GenS2{F0: 811}
    if v6 {
        v5 += (v7 >> ((v4) & 63))
        v36 := v1[444]
        for fi37 := 0; fi37 < 19; fi37++ {
            fmt.Println(v4)
            for fi38 := 0; fi38 < 11; fi38++ {
                v2 = append(v2, -498)
                fmt.Println(v4, v36)
                v39 := GenS1{F0: 607, F1: uint64(v7), F2: (!v6)}
                _ = fi38
                _ = v39
            }
            _ = len(v2)
            _ = fi37
        }
        _ = v36
    }
    switch ("" + ("foo bar" + "")) {
    case "a":
        if len(v2) > 0 {
            v2[0] = (v12 + v12)
        }
    case "":
        var v40 uint = v4
        _ = v40
    case "hello":
        var v41 int32 = -1
        _ = v41
    default:
        var v42 int64 = 1
        _ = v42
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v43 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
