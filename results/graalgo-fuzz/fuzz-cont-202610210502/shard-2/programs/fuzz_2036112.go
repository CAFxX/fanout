package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 int; F2 uint }

func main() {
    var v1 uint8 = 131
    var v2 bool = (((9223372036854775807 >= -173) && (!(!false))) || (32767 != 34))
    v3 := []string{"hello", ("" + "\n\t"), ("hello" + ""), "\n\t"}
    switch int(v1) {
    case 2147483647:
        if ((1 == 2147483647) || ((269 == -903) || (-159 != -32768))) {
            fmt.Println(v2, v1)
            v2 = true
            v1 -= v1
        }
    case -2147483648:
        if (-32768 <= -227) {
            fmt.Println(v1)
            for fi4 := 0; fi4 < 1; fi4++ {
                for fi5 := 0; fi5 < 10; fi5++ {
                    v6 := GenS1{F0: -1, F1: 0}
                    _ = fi5
                    _ = v6
                }
                if len(v3) > 0 {
                    v3[0] = ("a" + "\n\t")
                }
                _ = fi4
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 uint16 = 255
                    _ = fi8
                    _ = v9
                }
                if (1 >= 1) {
                    v10 := GenS2{F0: 127, F1: 2147483647, F2: 65535}
                    _ = v10
                } else {
                    fmt.Println(v2, fi7, fi7)
                }
                _ = fi7
            }
        } else {
            var v11 uint8 = v1
            _ = v11
        }
    case -156:
        fmt.Println(v1)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 string = ("\n\t" + ("" + "foo bar"))
        switch (fi12 & (fi12 >> ((18446744073709551615) & 63))) {
        case -9223372036854775808:
            var v14 bool = (!(v2 || v2))
            _ = v14
        case -1:
            v3 = append(v3, "foo bar")
        default:
            var v15 uint16 = 1
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 = v1
    v1 += 42
    var v16 uint8 = v1
    var v17 float32 = 1.5
    v3 = append(v3, "foo bar")
    v16 = (v1 + 127)
    var v18 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
