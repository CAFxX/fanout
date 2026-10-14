package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint8 = 1
    v2 := []int{-9223372036854775808}
    var v3 string = ((("hello" + "hello") + ("foo bar" + "")) + ("\n\t" + "a"))
    var v4 int = -878
    v5 := make(map[int]string)
    fmt.Println(v4, v1, v4)
    fmt.Println(v4)
    var v6 int = (v4 & (v4 - (-(v4))))
    v6 -= -101
    fmt.Println(v1)
    switch v4 {
    case -1:
        switch ("hello" + (v3 + v3)) {
        case "":
            v7 := GenS1{F0: 29}
            _ = v7
        default:
            v8 := GenS1{F0: 0}
            _ = v8
        }
    default:
        switch (v6 << ((65535) & 63)) {
        case 0:
            for fi9 := 0; fi9 < 17; fi9++ {
                switch (^(fi9)) {
                case 2147483647:
                    switch (v4 & 32767) {
                    case 958:
                        for fi10 := 0; fi10 < 20; fi10++ {
                            var v11 uint8 = (v1 | v1)
                            var v12 uint8 = (v11 >> ((354) & 63))
                            var v13 uint16 = 193
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    case 4294967295:
                        if (!(780 <= 0)) {
                            for fi14 := 0; fi14 < 20; fi14++ {
                                if ((0.0 != 0.1) || true) {
                                    var v15 uint16 = 197
                                    var v16 bool = true
                                    _ = v15
                                    _ = v16
                                } else {
                                    v1 = 0
                                    switch (fi14 | 1) {
                                    case 494:
                                        switch ((v3 + "foo bar") + v3) {
                                        case "hello":
                                            switch "hello" {
                                            case "":
                                                var v17 bool = (0 != 18446744073709551615)
                                                _ = v17
                                            case "a":
                                                var v18 uint64 = 77
                                                _ = v18
                                            case "hello":
                                                var v19 int64 = -982
                                                _ = v19
                                            }
                                            fmt.Println(v4, v1, v4)
                                            v3 = v3
                                        case "a":
                                            switch fi9 {
                                            case 365:
                                                var v20 int64 = -370
                                                _ = v20
                                            case -437:
                                                v4 += v6
                                            case 501:
                                                var v21 float32 = float32((0.0 * (1.5 - 0.1)))
                                                _ = v21
                                            default:
                                                switch v6 {
                                                case 245:
                                                    var v22 uint64 = 519
                                                    _ = v22
                                                case -1:
                                                    _ = len(v2)
                                                default:
                                                    var v23 bool = (true && ((!(false || (-470 >= 0))) || ((!false) && (true && (0.1 >= -2.25)))))
                                                    _ = v23
                                                }
                                            }
                                        }
                                    case 4294967295:
                                        var v24 uint16 = 1
                                        _ = v24
                                    case 131:
                                        var v25 int64 = 38
                                        _ = v25
                                    }
                                    v3 = v3
                                    if (v3 <= ((v3 + (((v3 + v3) + v3) + v3)) + ("foo bar" + "foo bar"))) {
                                        fmt.Println(v1, v1, v1)
                                        v4 = 9223372036854775807
                                    }
                                }
                                _ = fi14
                            }
                            v2 = append(v2, v4)
                        } else {
                            fmt.Println(v1)
                            v3 = ((v3 + (v3 + "a")) + "")
                            switch v3 {
                            case "a":
                                v3 = "\n\t"
                            case "":
                                fmt.Println(v4, v1)
                            case "foo bar":
                                v26 := GenS1{F0: 100}
                                _ = v26
                            }
                        }
                    case -1:
                        _ = len(v2)
                    }
                case -9223372036854775808:
                    v4 -= 128
                default:
                    for fi27 := 0; fi27 < 15; fi27++ {
                        if false {
                            var v28 int16 = 32767
                            if false {
                                fmt.Println(v28, fi27, fi9)
                                var v29 float64 = 1e10
                                var v30 uint = 646
                                _ = v29
                                _ = v30
                            }
                            v31 := v5[(v4 - 0)]
                            _ = v28
                            _ = v31
                        } else {
                            v5[9223372036854775807] = "\n\t"
                        }
                        fmt.Println(fi27, fi9, fi9)
                        v3 = v3
                        _ = fi27
                    }
                }
                v4 += (v4 % ((v4) | 1))
                _ = fi9
            }
        default:
            var v32 int8 = 127
            _ = v32
        }
    }
    v33 := GenS1{F0: 80}
    for fi34 := 0; fi34 < 11; fi34++ {
        switch v3 {
        case "foo bar":
            var v35 float32 = (1e10 + 1.5)
            _ = v35
        default:
            for fi36 := 0; fi36 < 18; fi36++ {
                if false {
                    switch v3 {
                    case "":
                        var v37 uint64 = 42
                        _ = v37
                    case "a":
                        v38 := GenS1{F0: 0}
                        _ = v38
                    }
                    var v39 int8 = 57
                    _ = v39
                }
                v6 -= (fi36 << ((428) & 63))
                delete(v5, v6)
                v2 = append(v2, 1)
                if true {
                    v40 := GenS2{F0: false, F1: -636}
                    v4 += v4
                    _ = v40
                }
                _ = fi36
            }
        }
        _ = fi34
    }
    fmt.Println(v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
