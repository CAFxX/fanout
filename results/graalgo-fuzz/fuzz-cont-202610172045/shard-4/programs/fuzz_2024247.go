package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 1
    v3 := []int{-9223372036854775808, int(v1), 53}
    var v4 uint8 = 128
    v3 = append(v3, -707)
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(v4)
        _ = fi5
    }
    var v6 uint = 65535
    switch "\n\t" {
    case "":
        if false {
            switch 127 {
            case -256:
                v2 += v2
            case 32768:
                v2 = v2
            case -1:
                _ = len(v3)
            default:
                v6 = 18446744073709551615
            }
            switch -712 {
            case 128:
                fmt.Println(v1)
            case 880:
                var v7 int32 = 0
                _ = v7
            case -1:
                var v8 uint8 = (v4 / ((148) | 1))
                _ = v8
            default:
                var v9 int16 = v1
                _ = v9
            }
            var v10 int64 = -9223372036854775808
            _ = v10
        } else {
            var v11 int16 = v1
            var v12 uint16 = 0
            var v13 uint64 = 392
            _ = v11
            _ = v12
            _ = v13
        }
    }
    if (!true) {
        v14 := GenS1{F0: (^(v2))}
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int = 127
            v6 -= (^(v6))
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        v1 -= v1
        switch 32768 {
        case 32767:
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: v2}
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi17
                _ = v18
            }
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 32768:
            switch "" {
            case "hello":
                switch "a" {
                case "foo bar":
                    fmt.Println(v1, v6)
                case "a":
                    v19 := GenS2{F0: -1, F1: false, F2: v6}
                    _ = v19
                default:
                    if ((-693 <= -526) && false) {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v1 -= int16(v4)
                            v3 = append(v3, (fi20 - fi20))
                            if len(v3) > 0 {
                                v3[0] = (^(fi20))
                            }
                            _ = fi20
                        }
                        fmt.Println(v2, v6, v2)
                    } else {
                        for fi21 := 0; fi21 < 8; fi21++ {
                            v22 := GenS1{F0: (^(v2))}
                            var v23 uint32 = v2
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        fmt.Println(v1, v4)
                    }
                }
            default:
                v6 += (v6 >> ((v6) & 63))
            }
        }
        var v24 uint64 = 18446744073709551615
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v25 int16 = (-(v1))
    v2 -= (v2 % ((v2) | 1))
    if ((-(v25)) != (v1 + v25)) {
        if false {
            var v26 int64 = 0
            _ = v26
        }
        var v27 bool = (1 <= 4294967295)
        _ = v27
    } else {
        v6 += (v6 ^ (v6 - (v6 | 2)))
    }
    v6 = v6
    var v28 float64 = float64(float32((1.5 - (float64(-2.25) * 1e10))))
    var v29 float64 = 3.14159
    switch (("a" + "hello") + "a") {
    case "":
        if true {
            var v30 uint8 = (v4 << ((v6) & 63))
            for fi31 := 0; fi31 < 10; fi31++ {
                v3 = append(v3, 42)
                _ = fi31
            }
            _ = len(v3)
            _ = v30
        } else {
            var v32 int16 = v1
            var v33 string = ""
            _ = v32
            _ = v33
        }
    default:
        fmt.Println(v1, v25, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v6
}
