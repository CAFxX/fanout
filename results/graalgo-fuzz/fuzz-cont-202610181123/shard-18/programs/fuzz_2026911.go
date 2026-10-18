package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (0 == 0)
    var v3 int8 = 42
    var v4 float64 = 3.14159
    var v5 uint16 = 1
    for fi6 := 0; fi6 < 14; fi6++ {
        switch -70 {
        case 9223372036854775807:
            for fi7 := 0; fi7 < 13; fi7++ {
                v2 = ("\n\t" <= "foo bar")
                _ = fi7
            }
        default:
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := []int{-818, fi6, fi8, (fi6 % (((fi6 + fi6)) | 1))}
                var v10 int16 = 32767
                switch "foo bar" {
                case "a":
                    v10 -= v10
                case "hello":
                    switch 81 {
                    case 128:
                        v1[(fi8 + (fi8 >> ((255) & 63)))] = (fi8 * (fi6 / ((fi6) | 1)))
                    case 32767:
                        fmt.Println(v2)
                    default:
                        v4 = v4
                    }
                default:
                    if v2 {
                        switch ("hello" + ("a" + (("foo bar" + "a") + ""))) {
                        case "\n\t":
                            v3 = v3
                        case "a":
                            v5 -= (^(v5))
                        default:
                            v10 += v10
                        }
                        v1[fi6] = (fi6 | -2147483648)
                    } else {
                        v11 := v1[(^(fi8))]
                        fmt.Println(v3, v10, v4)
                        if v2 {
                            v10 -= (v10 % (((v10 << ((862) & 63))) | 1))
                        }
                        _ = v11
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
    }
    if (v5 <= (v5 >> ((uint(v5)) & 63))) {
        v12 := GenS1{F0: 1000}
        _ = v12
    }
    fmt.Println(v2, v4)
    switch 517 {
    case 4294967295:
        var v13 int = 32768
        _ = v13
    case -2147483648:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 42
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v15, fi16)
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    default:
        v3 = 1
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS2{F0: fi18, F1: (v3 >> ((0) & 63)), F2: 42}
            _ = fi18
            _ = v19
        }
        delete(v1, fi17)
        _ = fi17
    }
    fmt.Println(v5, v5)
    if (642 > -9223372036854775808) {
        switch 4294967295 {
        case 0:
            v5 = (^(v5))
        default:
            v2 = ((!false) || v2)
        }
        if (!false) {
            v2 = (uint(v5) < 0)
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v3, v2, v4)
                _ = len(v1)
                _ = fi20
            }
        }
    }
    v4 += ((1.5 / 1.5) / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
