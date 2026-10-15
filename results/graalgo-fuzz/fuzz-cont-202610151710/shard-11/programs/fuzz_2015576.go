package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]uint32)
    v3 := []uint32{(v1 % (((v1 | (^(v1)))) | 1)), (v1 >> ((65535) & 63)), 4294967295}
    if false {
        var v4 uint16 = 65535
        _ = v4
    } else {
        v3 = append(v3, (v1 / (((v1 & (^(v1)))) | 1)))
        var v5 uint32 = (v1 >> ((1) & 63))
        for fi6 := 0; fi6 < 9; fi6++ {
            if (!(((v5 % ((v1) | 1)) != 2147483647) && (!((971 != 0) && false)))) {
                fmt.Println(v1)
                var v7 string = "\n\t"
                _ = v7
            }
            if false {
                v1 = (v5 - (v1 % (((v5 ^ 4294967295)) | 1)))
            }
            _ = fi6
        }
        _ = v5
    }
    fmt.Println(v1)
    if (!false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1)
    }
    v8 := GenS1{F0: "\n\t"}
    var v9 bool = (((!((-165 <= -11) || ((18446744073709551615 == 18446744073709551615) && (false || true)))) || true) || (518 == 127))
    switch 32767 {
    case 566:
        if (59 <= 0) {
            fmt.Println(v1, v1)
            v1 = (v1 * v1)
        } else {
            v9 = (18446744073709551615 > 362)
            fmt.Println(v1)
            fmt.Println(v9, v1)
        }
    case 2147483647:
        if v9 {
            var v10 bool = v9
            v2[-134] = 193
            _ = v10
        }
    }
    switch "" {
    case "foo bar":
        var v11 int = 2147483647
        _ = v11
    }
    v9 = true
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 string = "\n\t"
        _ = fi12
        _ = v13
    }
    if (125 == 127) {
        v9 = v9
    } else {
        var v14 int16 = 0
        var v15 string = "a"
        if v9 {
            if (!false) {
                switch 127 {
                case -864:
                    v16 := GenS1{F0: v15}
                    _ = v16
                case -782:
                    var v17 bool = (-128 != -67)
                    _ = v17
                case 719:
                    var v18 uint8 = 42
                    _ = v18
                default:
                    fmt.Println(v14, v1, v14)
                }
            }
        } else {
            var v19 uint32 = (v1 - (v1 >> ((369) & 63)))
            _ = v19
        }
        _ = v14
        _ = v15
    }
    fmt.Println(v9)
    v9 = ((!(v9 && (-1 > 0))) || (42 != 1))
    var v20 int64 = 1
    _ = len(v3)
    var v21 int16 = 32767
    v1 = (v1 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
