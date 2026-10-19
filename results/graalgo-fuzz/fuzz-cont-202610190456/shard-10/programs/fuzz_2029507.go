package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int) uint16 {
    v7 := GenS1{F0: ("" + "a"), F1: 2, F2: ("a" + "a")}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    switch 127 {
    case 2147483647:
        fmt.Println(p0)
    case 1:
        if true {
            var v10 uint8 = 255
            v10 += v10
            switch -903 {
            case 32768:
                v10 = v10
            case 42:
                p0 = (p0 >> ((0) & 63))
            default:
                p0 -= p0
            }
            _ = v10
        }
    case 32767:
        switch ("foo bar" + "foo bar") {
        case "a":
            v11 := GenS2{F0: -977, F1: 998, F2: 42}
            _ = v11
        case "hello":
            v12 := []int{}
            _ = v12
        }
    default:
        p0 += (p0 | (-(p0)))
    }
    v13 := GenS3{F0: (-(p0)), F1: false, F2: "hello"}
    _ = p0
    _ = v13
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 bool = false
    v5 := make(map[int]uint32)
    var v6 int = 42
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 = (0.1 != (1.5 / 1.5))
        if (166 < 114) {
            if v4 {
                _ = len(v3)
            } else {
                v4 = v4
            }
            if (1 == 42) {
                delete(v3, 1)
                var v15 uint32 = 65535
                _ = v15
            }
        } else {
            v4 = (18446744073709551615 < 42)
        }
        _ = fi14
    }
    var v16 uint32 = 2147483647
    v6 = v6
    if v4 {
        v4 = ((((-2.25 / 0.1) - (0.0 / 3.14159)) / 0.1) >= -2.25)
        var v17 int = v6
        if (408 <= 18446744073709551615) {
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS1{F0: ("\n\t" + (("foo bar" + (("\n\t" + "") + ("\n\t" + ("foo bar" + "\n\t")))) + "a")), F1: 0, F2: ("a" + ("a" + "hello"))}
                _ = fi18
                _ = v19
            }
            v20 := []int{v6, 9223372036854775807, v17, v17}
            v6 += v6
            _ = v20
        }
        _ = v17
    }
    fmt.Println(v16, v6, v6)
    v21 := GenS1{F0: "\n\t", F1: 65535, F2: "\n\t"}
    v22 := GenS4{F0: 255}
    v23 := GenS1{F0: "\n\t", F1: 0, F2: "foo bar"}
    v4 = v4
    var v24 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
