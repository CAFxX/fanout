package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 521
    v3 := []int{32767, -650, 2147483647}
    var v4 int32 = -1
    v5 := make(map[uint32]int)
    var v6 int8 = 0
    var v7 string = "foo bar"
    fmt.Println(v2, v2)
    var v8 float64 = 1.5
    fmt.Println(v2, v7, v2)
    v9 := GenS1{F0: v7, F1: 128}
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 1
        var v12 uint = 903
        switch v7 {
        case "\n\t":
            fmt.Println(v11, fi10, v8)
        case "foo bar":
            var v13 uint64 = v2
            _ = v13
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                v2 += v2
                var v15 uint16 = (v11 | (v11 - uint16(v12)))
                _ = fi14
                _ = v15
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if (true && ((!false) && (false && true))) {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch 9223372036854775807 {
            case 127:
                var v17 bool = false
                _ = v17
            case 0:
                delete(v1, 161)
            case -9223372036854775808:
                if (!(!(42 >= 42))) {
                    var v18 bool = true
                    _ = v18
                } else {
                    switch fi16 {
                    case 286:
                        for fi19 := 0; fi19 < 2; fi19++ {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            if true {
                                v4 = int32(fi19)
                                switch (^(fi16)) {
                                case 4294967295:
                                    v20 := v1[42]
                                    _ = v20
                                case 2147483647:
                                    _ = len(v3)
                                }
                            } else {
                                _ = len(v3)
                                if (true && true) {
                                    v6 = -77
                                    v21 := GenS2{F0: fi19}
                                    _ = v21
                                }
                            }
                            v7 = (v7 + v7)
                            _ = fi19
                        }
                    case 9223372036854775807:
                        v4 -= v4
                    case 971:
                        for fi22 := 0; fi22 < 18; fi22++ {
                            var v23 int64 = 9223372036854775807
                            _ = fi22
                            _ = v23
                        }
                    }
                    switch (fi16 / (((fi16 << ((65535) & 63))) | 1)) {
                    case -669:
                        if true {
                            for fi24 := 0; fi24 < 12; fi24++ {
                                if (25 > 317) {
                                    var v25 float32 = 1.5
                                    _ = v25
                                } else {
                                    fmt.Println(v8)
                                }
                                v2 = 18446744073709551615
                                v26 := GenS1{F0: ((v7 + ((v7 + v7) + ("hello" + (v7 + v7)))) + v7), F1: 2}
                                _ = fi24
                                _ = v26
                            }
                            v8 -= 0.0
                            v4 = v4
                        } else {
                            var v27 bool = false
                            v7 = "\n\t"
                            _ = v27
                        }
                    case -901:
                        v8 += 0.1
                    case 766:
                        v7 = v7
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            var v28 int = (fi16 >> ((382) & 63))
            _ = fi16
            _ = v28
        }
        if len(v3) > 0 {
            v3[0] = -193
        }
        v2 -= (^(v2))
    }
    fmt.Println(v7, v6, v8)
    v29 := GenS2{F0: -528}
    var v30 bool = true
    for fi31 := 0; fi31 < 8; fi31++ {
        v6 += (v6 / (((-(v6))) | 1))
        v7 = v7
        _ = fi31
    }
    var v32 uint32 = 31
    var v33 uint16 = 65535
    _ = len(v3)
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS2{F0: fi34}
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
