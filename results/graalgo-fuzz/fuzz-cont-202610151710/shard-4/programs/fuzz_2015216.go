package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = 180
    var v2 string = "foo bar"
    var v3 int16 = (v1 >> ((952) & 63))
    v4 := []string{"\n\t", ("a" + (((v2 + v2) + v2) + v2)), v2}
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 += -1
        _ = fi6
    }
    v2 = v2
    var v7 int64 = 9223372036854775807
    switch v2 {
    case "foo bar":
        v7 += (v7 % ((v7) | 1))
    case "a":
        switch v2 {
        case "\n\t":
            for fi8 := 0; fi8 < 18; fi8++ {
                if ((v2 + ((v2 + "\n\t") + v2)) > ("\n\t" + "")) {
                    if len(v4) > 0 {
                        v4[0] = (v2 + v2)
                    }
                }
                _ = fi8
            }
        case "":
            var v9 uint32 = v5
            _ = v9
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(fi10)
                _ = fi10
            }
        }
    case "\n\t":
        v7 -= 42
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v1 -= v1
        _ = fi11
    }
    v2 = ("\n\t" + ("hello" + v2))
    var v12 float32 = 3.14159
    var v13 string = v2
    switch (("\n\t" + "hello") + "a") {
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            v3 = v3
            v13 = v13
            _ = fi14
        }
    case "a":
        var v15 float64 = 1e10
        _ = v15
    default:
        v5 += v5
    }
    if false {
        v4 = append(v4, (v2 + ("" + v2)))
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v3)
            if true {
                _ = len(v4)
            } else {
                v3 += v1
                v2 = ((v2 + v2) + v2)
            }
            if (v7 > v7) {
                if false {
                    if (904 == -1) {
                        switch 2147483647 {
                        case 2147483647:
                            fmt.Println(v12)
                        case 4294967295:
                            switch 86 {
                            case -353:
                                for fi17 := 0; fi17 < 17; fi17++ {
                                    var v18 float32 = v12
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if true {
                                        var v19 uint16 = 678
                                        v3 -= v1
                                        _ = v19
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                            default:
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    v21 := GenS1{F0: true, F1: 42, F2: 1}
                                    switch fi16 {
                                    case 536:
                                        var v22 float32 = float32(1.5)
                                        _ = v22
                                    case -2147483648:
                                        v13 = "hello"
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        case 32767:
                            var v23 int8 = 0
                            _ = v23
                        }
                    } else {
                        v12 -= (v12 + 3.14159)
                        fmt.Println(v3, v13, v7)
                    }
                    fmt.Println(v5)
                    var v24 int = 32767
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v5 += (v5 * v5)
                        v26 := GenS2{F0: 65535}
                        var v27 uint64 = 4294967295
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 int = fi16
                _ = v28
            }
            _ = fi16
        }
        v12 = 0.0
    } else {
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 int = fi29
            var v31 bool = (1 < 0)
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            v1 += v3
            v2 = "a"
            _ = fi33
        }
        fmt.Println(v13, fi32)
        for fi34 := 0; fi34 < 6; fi34++ {
            if false {
                var v35 float64 = 0.0
                _ = v35
            }
            v36 := GenS1{F0: (false && (!((!true) && true))), F1: 255, F2: fi32}
            _ = fi34
            _ = v36
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
