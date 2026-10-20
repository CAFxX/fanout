package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 int8 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2, 580}
    var v3 uint16 = 0
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 16; fi5++ {
        if true {
            var v6 uint32 = 927
            var v7 float64 = 1e10
            _ = v6
            _ = v7
        }
        if false {
            v8 := v4[2147483647]
            _ = v8
        } else {
            v3 += (v3 % (((^(v3))) | 1))
            _ = len(v4)
            var v9 uint = 42
            _ = v9
        }
        var v10 float64 = (0.0 / 0.1)
        _ = fi5
        _ = v10
    }
    switch "" {
    case "":
        switch "foo bar" {
        case "":
            var v11 string = "a"
            _ = v11
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v12 uint64 = 4294967295
            _ = v12
        }
    case "foo bar":
        var v13 int16 = -1
        _ = v13
    }
    v1 += (v1 - v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 := GenS1{F0: true}
    var v15 bool = true
    switch ("hello" + "\n\t") {
    case "a":
        switch "hello" {
        case "foo bar":
            v3 = v3
        case "a":
            if v15 {
                if (v15 || (v15 || (32767 >= 705))) {
                    v3 += 0
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v2 = append(v2, 2147483648)
                        var v17 string = "foo bar"
                        v17 = "a"
                        _ = fi16
                        _ = v17
                    }
                    if v15 {
                        v15 = (int32(v3) >= -1)
                        v15 = (-1 == 2147483647)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        for fi18 := 0; fi18 < 18; fi18++ {
                            if v15 {
                                v2 = append(v2, 2147483647)
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    fmt.Println(v3, fi18)
                                    var v20 uint64 = 334
                                    var v21 int16 = -32768
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                fmt.Println(v3, v3, v15)
                            } else {
                                v15 = (!false)
                                fmt.Println(v1, fi18, v15)
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    var v23 int = 127
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v15 = (!v15)
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            var v24 uint16 = v3
                            _ = fi18
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 string = "a"
                            var v27 float64 = (-2.25 * 3.14159)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS2{F0: "hello", F1: v15, F2: -108}
                        _ = v28
                    }
                } else {
                    v3 = v3
                }
            }
        }
    case "hello":
        if v15 {
            v3 = v3
            v15 = v15
            v3 -= v3
        } else {
            fmt.Println(v15, v3, v1)
            var v29 int32 = 2147483647
            v30 := GenS2{F0: (("foo bar" + "a") + "a"), F1: ((18446744073709551615 == 0) && (v15 && (0 < 32767))), F2: 42}
            _ = v29
            _ = v30
        }
    }
    v31 := GenS1{F0: (((0 >= 32767) && (true && (1 == 2))) && v15)}
    if ((v3 ^ v3) >= 0) {
        if false {
            v1 += 0.1
            var v32 int16 = -1
            if v15 {
                v4[42] = "a"
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v15 = true
            }
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
