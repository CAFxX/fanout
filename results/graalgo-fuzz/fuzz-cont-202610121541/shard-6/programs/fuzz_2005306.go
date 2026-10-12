package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 uint32 = v1
    v4 = 2147483647
    for fi5 := 0; fi5 < 13; fi5++ {
        if true {
            switch (("\n\t" + "foo bar") + "a") {
            case "foo bar":
                v1 = (^(v1))
            case "hello":
                var v6 bool = (true || true)
                _ = v6
            }
        } else {
            _ = len(v3)
            var v7 uint64 = 415
            _ = v7
        }
        _ = fi5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 uint16 = 1
    fmt.Println(v8, v1, v1)
    switch "" {
    case "foo bar":
        v9 := GenS1{F0: 822}
        _ = v9
    case "a":
        switch -1 {
        case -9223372036854775808:
            v2[0] = true
        case 9223372036854775807:
            v8 += (^(v8))
        }
    case "":
        for fi10 := 0; fi10 < 10; fi10++ {
            v8 += 32767
            if (326 != int32(v1)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch "hello" {
            case "\n\t":
                if true {
                    if (!(((0.1 + 1.5) + 1e10) != (0.1 * 3.14159))) {
                        switch (^(fi10)) {
                        case 96:
                            var v11 int64 = 9223372036854775807
                            _ = v11
                        case 411:
                            v12 := GenS2{F0: -531}
                            _ = v12
                        case -731:
                            var v13 string = ("hello" + "hello")
                            _ = v13
                        default:
                            for fi14 := 0; fi14 < 16; fi14++ {
                                var v15 int16 = -32768
                                if true {
                                    v16 := GenS2{F0: (v15 & v15)}
                                    v17 := GenS2{F0: v15}
                                    if true {
                                        var v18 uint64 = 444
                                        var v19 int = fi10
                                        _ = v18
                                        _ = v19
                                    }
                                    fmt.Println(v4, v15)
                                    _ = v16
                                    _ = v17
                                } else {
                                    v4 = uint32(v8)
                                    v20 := GenS3{F0: 42}
                                    switch ("hello" + "hello") {
                                    case "a":
                                        v8 = (v8 % (((^(v8))) | 1))
                                    default:
                                        v8 += v8
                                    }
                                    _ = v20
                                }
                                switch fi10 {
                                case 641:
                                    v4 -= 467
                                case 32768:
                                    fmt.Println(v15)
                                case 602:
                                    var v21 int = (fi10 + (fi14 + fi10))
                                    v21 = fi10
                                    v22 := GenS1{F0: 65535}
                                    _ = v21
                                    _ = v22
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                fmt.Println(v8, v1, v8)
                                var v23 float32 = (0.1 * ((1.5 + 0.0) + 0.1))
                                _ = fi14
                                _ = v15
                                _ = v23
                            }
                        }
                    } else {
                        switch (^(fi10)) {
                        case 4294967295:
                            _ = len(v2)
                        case 32768:
                            switch "hello" {
                            case "":
                                v2[17] = (true || (false && false))
                            case "hello":
                                v1 -= v1
                            case "foo bar":
                                var v24 uint32 = (v1 / ((v4) | 1))
                                _ = v24
                            }
                        }
                        v4 += uint32(v8)
                    }
                    for fi25 := 0; fi25 < 19; fi25++ {
                        fmt.Println(fi10)
                        var v26 string = "hello"
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            default:
                v4 += (v4 % ((v1) | 1))
            }
            _ = fi10
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 uint8 = 77
        for fi29 := 0; fi29 < 18; fi29++ {
            if false {
                if (1e10 >= 3.14159) {
                    v30 := GenS2{F0: 557}
                    var v31 float64 = ((3.14159 - -2.25) * float64(1.5))
                    v32 := GenS2{F0: 32767}
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(fi27, v1, v4)
                }
            } else {
                v4 -= (v4 & (v1 << ((42) & 63)))
                v33 := GenS4{F0: 0, F1: -9223372036854775808}
                _ = v33
            }
            v34 := GenS4{F0: 0, F1: -924}
            _ = fi29
            _ = v34
        }
        var v35 uint32 = 51
        _ = fi27
        _ = v28
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
