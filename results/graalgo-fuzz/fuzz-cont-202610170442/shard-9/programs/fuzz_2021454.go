package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 353
    v3 := []string{((("hello" + "hello") + "a") + "foo bar"), "\n\t", "foo bar"}
    if (0 > 0) {
        fmt.Println(v1)
        var v4 bool = (v1 == (v1 & (^(v1))))
        _ = v4
    }
    switch 32768 {
    case -87:
        var v5 uint16 = (v1 | (v1 / ((2) | 1)))
        _ = v5
    case 2147483647:
        switch ("foo bar" + ("a" + "\n\t")) {
        case "\n\t":
            v1 = (^(v1))
        case "a":
            if false {
                if ((1 <= 0) || (-350 != 32767)) {
                    v1 = (v1 + 65535)
                    v6 := GenS1{F0: 531, F1: false, F2: 42}
                    _ = v6
                }
                v1 += (v1 / ((32767) | 1))
            }
        case "foo bar":
            if false {
                switch "a" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                case "":
                    v7 := GenS2{F0: (false || ((float64(1.5) / 1.0) != 1e10)), F1: v1, F2: 42}
                    _ = v7
                case "foo bar":
                    for fi8 := 0; fi8 < 9; fi8++ {
                        fmt.Println(fi8)
                        _ = fi8
                    }
                }
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int64 = v2
                    var v11 uint32 = 868
                    fmt.Println(v10, v2)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        default:
            if len(v3) > 0 {
                v3[0] = (("" + "a") + "hello")
            }
        }
    }
    var v12 string = ""
    v3 = append(v3, ((("" + v12) + v12) + ("hello" + "hello")))
    if (float32(3.14159) <= (3.14159 - ((0.1 * (1.5 * (1e10 / 1.5))) * (((float32((3.14159 / 3.14159)) / 1.0) * 3.14159) + 1e10)))) {
        v13 := GenS2{F0: (!(1 < 187)), F1: (v1 ^ (v1 ^ 32767)), F2: 128}
        fmt.Println(v12)
        _ = v13
    }
    var v14 uint64 = 513
    switch 1 {
    case -77:
        v15 := GenS1{F0: 18446744073709551615, F1: (-1 < -1), F2: uint8(v1)}
        _ = v15
    default:
        v1 = (v1 * v1)
    }
    v2 = int64(v14)
    fmt.Println(v2, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
