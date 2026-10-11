package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []uint32{79, 188, 2147483647}
    var v2 int = 4294967295
    v3 := make(map[int]bool)
    fmt.Println(v2)
    v2 = -1
    if true {
        var v4 uint16 = 164
        _ = v4
    } else {
        if false {
            var v5 uint16 = 613
            fmt.Println(v2)
            _ = v5
        }
        var v6 uint64 = 60
        _ = v6
    }
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            fmt.Println(fi8, fi8)
            if true {
                fmt.Println(fi8)
                if (true || false) {
                    var v9 int16 = -644
                    _ = v9
                } else {
                    var v10 uint16 = 65535
                    var v11 float32 = (float32((((((-2.25 + (float64(1e10) * 1e10)) + (0.1 / float64(1.5))) - 0.1) * (1.5 - 1.5)) - 0.1)) / 1e10)
                    _ = v10
                    _ = v11
                }
                switch (("hello" + "hello") + (("a" + "") + "\n\t")) {
                case "foo bar":
                    switch v2 {
                    case 128:
                        v12 := GenS1{F0: v2}
                        _ = v12
                    case -487:
                        switch "a" {
                        case "\n\t":
                            if true {
                                v3[(v2 / ((v2) | 1))] = (!true)
                                v2 += fi8
                            } else {
                                var v13 int = fi8
                                _ = v13
                            }
                        default:
                            switch (v2 & (-(v2))) {
                            case 32768:
                                if (false && (0.1 != 3.14159)) {
                                    v14 := v3[fi8]
                                    v14 = v14
                                    _ = len(v3)
                                    _ = v14
                                }
                            case 4294967295:
                                if true {
                                    var v15 int64 = 291
                                    _ = v15
                                }
                            default:
                                if false {
                                    switch "" {
                                    case "":
                                        for fi16 := 0; fi16 < 4; fi16++ {
                                            v17 := GenS1{F0: 1}
                                            var v18 float32 = 0.1
                                            delete(v3, -763)
                                            _ = fi16
                                            _ = v17
                                            _ = v18
                                        }
                                    case "a":
                                        for fi19 := 0; fi19 < 10; fi19++ {
                                            var v20 int64 = 9223372036854775807
                                            _ = fi19
                                            _ = v20
                                        }
                                    default:
                                        var v21 float64 = float64(0.1)
                                        _ = v21
                                    }
                                    v2 += v2
                                } else {
                                    if false {
                                        for fi22 := 0; fi22 < 20; fi22++ {
                                            v2 = fi8
                                            _ = fi22
                                        }
                                        var v23 int = fi8
                                        v2 = -9223372036854775808
                                        _ = v23
                                    }
                                    if (2 <= 42) {
                                        if true {
                                            _ = len(v1)
                                        }
                                        v2 += v2
                                    } else {
                                        v2 = (v2 >> ((0) & 63))
                                        v3[(v2 / ((v2) | 1))] = ((19 == 1) || (0 >= 1))
                                        v24 := GenS1{F0: 0}
                                        _ = v24
                                    }
                                }
                            }
                        }
                    }
                case "a":
                    v2 += (fi8 + 32768)
                default:
                    _ = len(v3)
                }
            }
            v25 := GenS2{F0: 979}
            _ = v25
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            if (!false) {
                var v27 bool = (0 == 949)
                var v28 int64 = -1
                _ = v27
                _ = v28
            } else {
                switch (fi8 >> ((696) & 63)) {
                case -1:
                    v2 = v2
                case 0:
                    v2 -= (v2 * fi26)
                default:
                    fmt.Println(v2, fi8)
                }
                _ = len(v3)
            }
            _ = fi26
        }
        _ = fi8
    }
    if (("foo bar" + "foo bar") == "") {
        switch 42 {
        case 127:
            switch "foo bar" {
            case "\n\t":
                v29 := GenS1{F0: (v2 - (v2 / ((v2) | 1)))}
                _ = v29
            default:
                var v30 uint8 = 14
                _ = v30
            }
        case 0:
            var v31 uint32 = 42
            _ = v31
        default:
            var v32 uint8 = 224
            _ = v32
        }
        v33 := GenS3{F0: 32767}
        _ = v33
    }
    v2 = v2
    var v34 string = "\n\t"
    v35 := GenS4{F0: 405}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v7
}
