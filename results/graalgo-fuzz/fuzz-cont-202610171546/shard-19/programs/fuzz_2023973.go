package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 0
    v3 := []string{"a"}
    var v4 bool = true
    var v5 int8 = 127
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    switch 32768 {
    case 873:
        var v6 int8 = 42
        _ = v6
    case 9223372036854775807:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch "" {
        case "hello":
            switch 61 {
            case 9223372036854775807:
                var v7 string = ""
                _ = v7
            case 127:
                _ = len(v3)
            default:
                if true {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        if ((!v4) && v4) {
                            for fi9 := 0; fi9 < 4; fi9++ {
                                var v10 int8 = (v5 + v5)
                                var v11 uint16 = 255
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            if v4 {
                                if (911 == uint64(v2)) {
                                    v5 += (-(v5))
                                    if false {
                                        fmt.Println(v2, v4)
                                    }
                                }
                                for fi12 := 0; fi12 < 16; fi12++ {
                                    fmt.Println(v1, fi12, fi8)
                                    if true {
                                        var v13 int8 = v5
                                        var v14 int8 = 0
                                        if (2 < 409) {
                                            v2 -= v2
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = "a"
                                            }
                                            var v15 int8 = (^(v14))
                                            v2 -= 2
                                            _ = v15
                                        }
                                        _ = v13
                                        _ = v14
                                    }
                                    _ = fi12
                                }
                            }
                            var v16 uint32 = (^(v1))
                            _ = v16
                        }
                        _ = fi8
                    }
                    v1 = (v1 | (^(v1)))
                    v5 += v5
                } else {
                    if (-1 >= 0) {
                        switch -1 {
                        case 127:
                            var v17 int16 = 1
                            _ = v17
                        case 42:
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + ("hello" + "foo bar"))
                            }
                        }
                    }
                }
            }
        case "":
            _ = len(v3)
        }
    }
    if v4 {
        v18 := GenS1{F0: ((-2.25 >= 1e10) && v4), F1: ("hello" + ""), F2: (v5 - v5)}
        var v19 uint = 65535
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 uint32 = v1
                _ = fi21
                _ = v22
            }
            v1 = v1
            _ = fi20
        }
        _ = v18
        _ = v19
    }
    v3 = append(v3, "a")
    var v23 int16 = 1
    for fi24 := 0; fi24 < 14; fi24++ {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: v4, F1: "foo bar", F2: v5}
            v5 = (v5 >> ((42) & 63))
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    _ = len(v3)
    var v27 int = -2147483648
    var v28 bool = (false || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
