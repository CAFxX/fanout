package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    v4 := []int{-939, 42, 2147483647, 42}
    switch "hello" {
    case "":
        fmt.Println(v3)
    }
    v5 := v2[v3]
    if false {
        v3 = (v3 ^ (v3 | v3))
        if true {
            v6 := GenS1{F0: 1}
            for fi7 := 0; fi7 < 15; fi7++ {
                v5 = v5
                _ = fi7
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                if (138 >= 0) {
                    delete(v2, (v3 - v3))
                    if true {
                        if (!true) {
                            v3 += v3
                        }
                        var v9 uint64 = 2
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 7; fi10++ {
                        fmt.Println(v3, v5, fi10)
                        if false {
                            if false {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                _ = len(v2)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v11 uint64 = 2
                            if true {
                                var v12 string = ("hello" + v5)
                                _ = v12
                            } else {
                                switch v5 {
                                case "":
                                    delete(v2, v3)
                                    v4 = append(v4, (-(fi8)))
                                    var v13 int16 = 0
                                    var v14 uint64 = v11
                                    if false {
                                        switch "a" {
                                        case "\n\t":
                                            switch (v5 + (v5 + "hello")) {
                                            case "":
                                                fmt.Println(v14, v11)
                                            case "a":
                                                v3 += 2147483647
                                            default:
                                                var v15 float32 = -2.25
                                                _ = v15
                                            }
                                        }
                                        v14 = 2
                                        v3 = 4294967295
                                    }
                                    _ = v13
                                    _ = v14
                                case "\n\t":
                                    var v16 int64 = int64(v3)
                                    _ = v16
                                default:
                                    v5 = v5
                                }
                                v4 = append(v4, (fi10 >> ((255) & 63)))
                                if true {
                                    var v17 string = v5
                                    _ = v17
                                }
                                if false {
                                    v18 := GenS1{F0: 2}
                                    v19 := v2[(v3 ^ v3)]
                                    if (v19 == v19) {
                                        v20 := GenS1{F0: 128}
                                        if ((!(!true)) && ("\n\t" != (((("\n\t" + "hello") + "\n\t") + "hello") + v5))) {
                                            fmt.Println(v3, v5, v11)
                                        } else {
                                            if false {
                                                var v21 uint = 65535
                                                v1 = v1
                                                _ = v21
                                            } else {
                                                _ = len(v2)
                                                var v22 uint8 = 2
                                                v4 = append(v4, (fi8 | 32768))
                                                _ = v22
                                            }
                                            v19 = "\n\t"
                                            var v23 string = (v5 + (v19 + "a"))
                                            _ = v23
                                        }
                                        _ = v20
                                    }
                                    _ = v18
                                    _ = v19
                                }
                                v24 := GenS1{F0: 1}
                                _ = v24
                            }
                            _ = v11
                        } else {
                            switch v5 {
                            case "\n\t":
                                v5 = "foo bar"
                            }
                        }
                        _ = fi10
                    }
                }
                var v25 uint16 = 851
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := v2[(v3 >> ((959) & 63))]
                    _ = fi26
                    _ = v27
                }
                _ = fi8
                _ = v25
            }
            _ = v6
        }
        var v28 uint32 = 2147483647
        _ = v28
    }
    v5 = v5
    var v29 int64 = 9223372036854775807
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 15; fi31++ {
        v3 = (v3 & (v3 - (v3 ^ 2)))
        var v32 float32 = ((1e10 - (v1 / (((float32(1.5) / (v1 / (v1 - (float32(1e10) * v1)))) / (0.0 - v1)) - (((-2.25 - (v1 + 0.0)) + v1) - (v1 - 1.5))))) * (((v1 - v1) + v1) + 0.0))
        _ = fi31
        _ = v32
    }
    if (!((!true) && true)) {
        v3 -= (v3 << ((18446744073709551615) & 63))
        v29 -= (v29 >> ((42) & 63))
        v33 := GenS1{F0: 127}
        _ = v33
    }
    v34 := GenS2{F0: int(v29), F1: -899}
    v4 = append(v4, -297)
    var v35 uint8 = 255
    v36 := GenS2{F0: 127, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
