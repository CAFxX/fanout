package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := []string{("" + "foo bar"), "a", "", ("" + "")}
    v5 := make(map[uint32]uint32)
    var v6 bool = true
    v1 -= (v1 >> ((4294967295) & 63))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = 0
        var v9 float32 = float32((-2.25 + (1e10 - 3.14159)))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = "\n\t"
    if true {
        switch 749 {
        case 4294967295:
            v11 := GenS1{F0: v1, F1: 204}
            _ = v11
        case 0:
            v6 = v6
        case 9223372036854775807:
            v1 -= v1
        }
    }
    if (!(!true)) {
        fmt.Println(v6, v10, v10)
        for fi12 := 0; fi12 < 10; fi12++ {
            if (v6 || true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    fmt.Println(v10, fi13, fi13)
                    if v6 {
                        var v14 int16 = -1
                        v2[fi12] = fi12
                        var v15 uint32 = 65535
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS2{F0: 42}
                    _ = fi13
                    _ = v16
                }
            }
            if (!false) {
                v17 := GenS2{F0: 0}
                v1 = 0
                _ = v17
            } else {
                switch v10 {
                case "foo bar":
                    if v6 {
                        switch v10 {
                        case "":
                            var v18 int8 = 42
                            _ = v18
                        case "a":
                            if v6 {
                                switch v10 {
                                case "":
                                    v19 := GenS3{F0: ((v10 + ("a" + "\n\t")) + ((v10 + (v10 + v10)) + "foo bar"))}
                                    _ = v19
                                default:
                                    v10 = (v10 + ("hello" + ("" + ("\n\t" + (v10 + v10)))))
                                }
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    var v21 uint = 18446744073709551615
                                    _ = fi20
                                    _ = v21
                                }
                                var v22 uint64 = 1
                                _ = v22
                            } else {
                                fmt.Println(v10, v6, v10)
                                v23 := GenS2{F0: 1}
                                _ = v23
                            }
                        case "foo bar":
                            if v6 {
                                var v24 uint16 = v1
                                _ = v24
                            }
                        }
                    } else {
                        delete(v3, ((v10 + (v10 + (v10 + v10))) + v10))
                        v25 := GenS4{F0: 0, F1: 1, F2: 2}
                        _ = v25
                    }
                case "hello":
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v27 := GenS4{F0: 42, F1: 0, F2: 0}
                        var v28 bool = (!(100 <= 80))
                        v1 = (^(v1))
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case "a":
                    var v29 int32 = 84
                    _ = v29
                default:
                    var v30 float32 = (0.1 + -2.25)
                    _ = v30
                }
            }
            if true {
                for fi31 := 0; fi31 < 15; fi31++ {
                    v1 -= (v1 / ((32767) | 1))
                    var v32 int8 = 127
                    _ = fi31
                    _ = v32
                }
                var v33 float32 = 0.1
                if v6 {
                    fmt.Println(v33)
                }
                _ = v33
            } else {
                var v34 bool = v6
                v6 = (!false)
                v35 := GenS5{F0: 32767}
                _ = v34
                _ = v35
            }
            _ = fi12
        }
        v36 := GenS4{F0: 42, F1: 641, F2: 2}
        _ = v36
    }
    if len(v4) > 0 {
        v4[0] = (v10 + "\n\t")
    }
    switch "a" {
    case "a":
        var v37 bool = (35 < 128)
        _ = v37
    }
    if v6 {
        var v38 float32 = (-2.25 * 3.14159)
        _ = v38
    } else {
        switch (((v10 + "hello") + v10) + v10) {
        case "hello":
            v39 := GenS4{F0: -863, F1: 42, F2: 2}
            _ = v39
        case "a":
            v40 := GenS6{F0: int64(v1)}
            _ = v40
        case "\n\t":
            var v41 float64 = 0.0
            _ = v41
        default:
            var v42 string = ("hello" + (v10 + v10))
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v44 uint8 = 127
        _ = fi43
        _ = v44
    }
    v45 := GenS3{F0: "\n\t"}
    v46 := GenS7{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
