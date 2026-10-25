package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []uint32{}
    var v2 int = 25
    var v3 string = "hello"
    var v4 int = v2
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int8 = 127
            v2 += fi6
            _ = fi6
            _ = v7
        }
        switch (v2 - v2) {
        case 1:
            if false {
                switch "\n\t" {
                case "":
                    var v8 uint16 = 1
                    _ = v8
                default:
                    v2 = -978
                }
                v4 -= (v4 >> ((4294967295) & 63))
            }
        case 42:
            v2 -= 32768
        case 64:
            v1 = append(v1, 65535)
        default:
            var v9 uint32 = 65535
            _ = v9
        }
        v2 += (fi5 - (v4 - 128))
        _ = fi5
    }
    if true {
        switch v3 {
        case "\n\t":
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint64 = 908
                    if false {
                        if false {
                            v3 = (((v3 + (v3 + v3)) + (v3 + "a")) + "foo bar")
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v11 -= 418
                                var v13 int16 = 0
                                fmt.Println(v13)
                                _ = fi12
                                _ = v13
                            }
                            v14 := GenS1{F0: "\n\t"}
                            _ = v14
                        }
                        fmt.Println(fi10)
                    } else {
                        if (((1e10 * (1.5 * 1.5)) - 1.5) > 1e10) {
                            var v15 uint8 = uint8(v11)
                            var v16 uint64 = 757
                            if true {
                                if (false || true) {
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        if true {
                                            var v18 int = (fi10 % ((fi17) | 1))
                                            _ = v18
                                        } else {
                                            v19 := GenS1{F0: v3}
                                            _ = v19
                                        }
                                        var v20 int64 = 1
                                        _ = fi17
                                        _ = v20
                                    }
                                    if (false || true) {
                                        var v21 int16 = 1
                                        v22 := GenS2{F0: (v11 + v11)}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        for fi23 := 0; fi23 < 4; fi23++ {
                                            var v24 uint32 = 946
                                            _ = fi23
                                            _ = v24
                                        }
                                        v4 = v4
                                        v3 = v3
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                }
                                var v25 uint64 = (v16 * v11)
                                _ = v25
                            } else {
                                v15 += v15
                                var v26 int16 = 0
                                switch v3 {
                                case "\n\t":
                                    v27 := GenS1{F0: v3}
                                    _ = v27
                                case "":
                                    v28 := GenS2{F0: (^(v11))}
                                    _ = v28
                                case "hello":
                                    v29 := GenS3{F0: ((((("foo bar" + v3) + "a") + v3) + (v3 + ("hello" + ((v3 + v3) + (((v3 + "a") + v3) + v3))))) > v3), F1: -9223372036854775808}
                                    _ = v29
                                default:
                                    v26 += (-(v26))
                                }
                                _ = v26
                            }
                            _ = v15
                            _ = v16
                        }
                        switch 1 {
                        case 1:
                            var v30 uint32 = 0
                            _ = v30
                        case 42:
                            v2 += (fi10 - (v4 * v4))
                        case 32767:
                            if false {
                                v11 = v11
                                var v31 uint = 42
                                _ = v31
                            }
                        default:
                            switch fi10 {
                            case 128:
                                v2 = fi10
                            case 0:
                                var v32 uint16 = 959
                                _ = v32
                            default:
                                v11 -= v11
                            }
                        }
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 uint16 = 1
                            var v35 float32 = 1e10
                            switch ("\n\t" + (v3 + "hello")) {
                            case "a":
                                v2 += 1
                            default:
                                v36 := GenS2{F0: v11}
                                _ = v36
                            }
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = fi10
                    _ = v11
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    switch (fi37 % ((1) | 1)) {
                    case 2147483647:
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                        }
                    case -2147483648:
                        v2 += v4
                    case 581:
                        v3 = v3
                    default:
                        v3 = v3
                    }
                    _ = fi37
                }
            }
        }
    }
    fmt.Println(v3, v3, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v38 float64 = 1e10
    v38 += (v38 + 1.5)
    for fi39 := 0; fi39 < 3; fi39++ {
        v40 := GenS4{F0: 0, F1: 1, F2: 2147483648}
        v41 := GenS5{F0: 1}
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint16 = 325
    v42 = (v42 / ((255) | 1))
    fmt.Println(v38)
    switch v4 {
    case -1:
        var v43 uint16 = v42
        _ = v43
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
}
