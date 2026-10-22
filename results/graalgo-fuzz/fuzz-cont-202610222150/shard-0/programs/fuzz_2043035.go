package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = 1
    var v2 string = "hello"
    v3 := []uint32{2, 0}
    var v4 int16 = -1
    var v5 int = 472
    var v6 bool = true
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: ("hello" + "\n\t"), F1: 4294967295}
            _ = fi8
            _ = v9
        }
        switch fi7 {
        case 9223372036854775807:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 20:
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS2{F0: 0}
                _ = fi10
                _ = v11
            }
        }
        _ = fi7
    }
    v2 = "\n\t"
    v2 = v2
    fmt.Println(v6, v4)
    if ((2147483647 < 346) && (1 >= v1)) {
        switch (v2 + v2) {
        case "":
            var v12 uint64 = 1
            _ = v12
        default:
            v6 = false
        }
        if v6 {
            if (((1e10 / (0.1 - 1e10)) <= -2.25) || v6) {
                v13 := GenS2{F0: -9223372036854775808}
                v5 += 2147483647
                for fi14 := 0; fi14 < 16; fi14++ {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        var v16 uint64 = 991
                        v17 := GenS3{F0: (v2 <= ("foo bar" + ("\n\t" + "hello")))}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                _ = v13
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint = 255
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (!(3.14159 >= 1.5)) {
                        v21 := GenS4{F0: 111}
                        _ = v21
                    } else {
                        v4 += int16(v1)
                    }
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
            if ("" < ((v2 + v2) + v2)) {
                switch "" {
                case "":
                    var v22 int32 = 0
                    _ = v22
                case "a":
                    if (v5 <= -403) {
                        _ = len(v3)
                    } else {
                        v23 := GenS3{F0: v6}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v5 -= 127
                        _ = v23
                    }
                default:
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 string = "hello"
                        v5 += (fi24 & (^(v5)))
                        if (v1 > (v1 - (v1 ^ v1))) {
                            v2 = ("a" + v2)
                            switch -936 {
                            case 17:
                                var v26 uint64 = 4294967295
                                _ = v26
                            default:
                                var v27 uint16 = 255
                                _ = v27
                            }
                            v5 = (^(fi24))
                        } else {
                            var v28 string = "a"
                            _ = v28
                        }
                        _ = fi24
                        _ = v25
                    }
                }
            }
        }
        v5 += 32767
    } else {
        for fi29 := 0; fi29 < 15; fi29++ {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = fi29
        }
    }
    var v30 int64 = 805
    fmt.Println(v6, v4, v1)
    fmt.Println(v30, v30, v1)
    v1 += v1
    v4 -= (v4 | (v4 * 32767))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
