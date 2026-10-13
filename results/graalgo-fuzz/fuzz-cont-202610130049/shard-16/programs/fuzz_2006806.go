package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    var v2 uint32 = 1
    v3 := make(map[uint32]bool)
    var v4 int32 = v1
    var v5 string = (("hello" + "foo bar") + "a")
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 127 {
        case 42:
            v6 := []int{0}
            _ = v6
        case 128:
            if false {
                for fi7 := 0; fi7 < 18; fi7++ {
                    v2 -= v2
                    if false {
                        v2 -= 2147483647
                    }
                    _ = fi7
                }
                v1 = (v1 >> ((2) & 63))
            } else {
                v4 += 849
            }
        case 32767:
            for fi8 := 0; fi8 < 10; fi8++ {
                fmt.Println(v2, v5, v1)
                v5 = v5
                v9 := GenS1{F0: false, F1: -1}
                _ = fi8
                _ = v9
            }
        default:
            v5 = v5
        }
    case 28:
        v5 = v5
    case 428:
        fmt.Println(v4, v5, v4)
    }
    switch -9223372036854775808 {
    case 4294967295:
        fmt.Println(v5, v2)
    default:
        v10 := []int{473, 533}
        _ = v10
    }
    v4 += v1
    if (false && false) {
        for fi11 := 0; fi11 < 20; fi11++ {
            v3[(v2 % (((v2 + 4294967295)) | 1))] = true
            _ = fi11
        }
        if (-429 < int(v2)) {
            if (0 >= 123) {
                v12 := []int{355, -585}
                v1 = 2147483647
                switch "\n\t" {
                case "foo bar":
                    var v13 float64 = float64((0.1 + ((0.0 + 0.0) + 0.0)))
                    _ = v13
                case "\n\t":
                    var v14 uint = 42
                    _ = v14
                case "hello":
                    switch 2147483647 {
                    case 127:
                        if false {
                            v4 += v4
                        }
                    case -748:
                        v2 = 2147483648
                    case 32767:
                        v2 -= 42
                    }
                }
                _ = v12
            }
        }
    }
    fmt.Println(v2, v4)
    v1 -= (v4 % (((v4 ^ v1)) | 1))
    v15 := GenS1{F0: false, F1: 0}
    delete(v3, v2)
    v16 := GenS1{F0: (true && (!(1e10 >= 1e10))), F1: 1}
    var v17 int16 = 155
    var v18 uint32 = (v2 | (^(v2)))
    v19 := GenS1{F0: false, F1: -60}
    fmt.Println(v17, v17, v18)
    v4 -= v1
    switch 127 {
    case -9223372036854775808:
        v20 := v3[v18]
        _ = v20
    default:
        var v21 uint = 255
        _ = v21
    }
    if (!(v4 < 1)) {
        v4 -= (v4 >> ((65535) & 63))
        v2 += (v2 * (v2 >> ((18446744073709551615) & 63)))
        if ((v5 + v5) < "") {
            v18 += (v18 + 2)
        } else {
            _ = len(v3)
            v22 := GenS2{F0: -64}
            _ = v22
        }
    }
    delete(v3, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
