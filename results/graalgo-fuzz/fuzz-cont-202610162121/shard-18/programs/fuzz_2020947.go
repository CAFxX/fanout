package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
func genFn0(p0 uint, p1 int32) uint {
    switch -2147483648 {
    case 42:
        var v4 uint64 = 42
        _ = v4
    default:
        switch "a" {
        case "hello":
            v5 := GenS1{F0: (1 > 127), F1: 2, F2: ((19 != -1) || (false || true))}
            _ = v5
        case "\n\t":
            p0 += p0
        default:
            var v6 string = ("\n\t" + "\n\t")
            _ = v6
        }
    }
    var v7 int64 = 382
    _ = p0
    _ = p1
    _ = v7
    return (p0 + p0)
}

func main() {
    v1 := []int{9223372036854775807, 2147483647, -9223372036854775808, 42}
    v2 := []string{}
    var v3 uint16 = 0
    if true {
        v3 = (v3 & (v3 >> ((683) & 63)))
        v3 = 42
        var v8 string = "hello"
        _ = v8
    }
    v3 = (^(v3))
    switch 949 {
    case 9223372036854775807:
        var v9 uint16 = 42
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 42, F2: (!(42 > -2147483648))}
    v3 -= 32767
    switch 930 {
    case 32768:
        v3 -= (v3 >> ((42) & 63))
    case 2147483647:
        if false {
            v3 += v3
            var v11 uint16 = (^(v3))
            var v12 float32 = float32(0.0)
            _ = v11
            _ = v12
        } else {
            v3 = (v3 / (((v3 - 0)) | 1))
        }
    default:
        if true {
            if false {
                var v13 float32 = -2.25
                _ = v13
            } else {
                var v14 uint8 = 0
                var v15 int16 = 1
                _ = v14
                _ = v15
            }
            if true {
                _ = len(v1)
            } else {
                v3 = (v3 ^ (v3 | v3))
                fmt.Println(v3)
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS1{F0: (("hello" + "\n\t") == ("foo bar" + "foo bar")), F1: 259, F2: true}
                    _ = fi16
                    _ = v17
                }
            }
            v3 = v3
        } else {
            v1 = append(v1, 4294967295)
        }
    }
    v18 := GenS2{F0: -829, F1: "", F2: 1}
    v3 += 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
}
