package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 32767
    var v3 uint64 = 2
    v4 := make(map[uint32]bool)
    var v5 uint16 = 255
    v4[2] = true
    v6 := []int{823}
    v2 = v2
    var v7 int64 = -882
    v5 -= v5
    switch 4294967295 {
    case 197:
        v8 := GenS1{F0: 65535, F1: 1, F2: -1}
        _ = v8
    case -723:
        v2 = (v2 | v2)
    case 2147483647:
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint32 = 67
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v10, v2, v10)
                _ = fi11
            }
            fmt.Println(v3)
            _ = fi9
            _ = v10
        }
    default:
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 1, F1: 127, F2: 127}
            var v14 uint64 = 753
            var v15 uint32 = 73
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    v5 = (v5 * v5)
    if false {
        fmt.Println(v5, v3)
    } else {
        if (!true) {
            var v16 int16 = 1
            _ = v16
        } else {
            if false {
                if true {
                    v7 += (v7 >> ((20) & 63))
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(v3, v2, fi17)
                    _ = fi17
                }
                if ((v5 << ((188) & 63)) >= v5) {
                    var v18 int8 = 100
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        for fi20 := 0; fi20 < 12; fi20++ {
                            v7 -= v7
                            if len(v6) > 0 {
                                v6[0] = fi20
                            }
                            _ = fi20
                        }
                        v5 = v5
                        _ = fi19
                    }
                    v1[9223372036854775807] = (v7 == v7)
                }
            } else {
                if false {
                    v7 += 1
                    var v21 int32 = 2147483647
                    v5 -= 0
                    _ = v21
                } else {
                    v22 := GenS1{F0: 4294967295, F1: 1, F2: 0}
                    v23 := GenS1{F0: 255, F1: 255, F2: 42}
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 10; fi24++ {
                    v25 := GenS2{F0: v3}
                    for fi26 := 0; fi26 < 13; fi26++ {
                        if (-50 != 127) {
                            var v27 string = "foo bar"
                            switch (v27 + "hello") {
                            case "a":
                                switch "\n\t" {
                                case "":
                                    switch v27 {
                                    case "hello":
                                        var v28 uint32 = 4294967295
                                        _ = v28
                                    }
                                }
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi24 - fi26)
                                }
                            }
                            if false {
                                var v29 int32 = -381
                                v29 += (v29 / (((v29 / ((0) | 1))) | 1))
                                _ = v29
                            }
                            _ = v27
                        }
                        v30 := GenS3{F0: 0}
                        if len(v6) > 0 {
                            v6[0] = 32768
                        }
                        v4[4294967295] = ((!(!("foo bar" <= "foo bar"))) && true)
                        v7 = (v7 - -920)
                        v7 -= 525
                        v31 := GenS1{F0: uint(v3), F1: 127, F2: -13}
                        _ = fi26
                        _ = v30
                        _ = v31
                    }
                    _ = fi24
                    _ = v25
                }
            }
            v32 := GenS4{F0: -2147483648, F1: -1, F2: ("foo bar" + "\n\t")}
            v3 += 18446744073709551615
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
