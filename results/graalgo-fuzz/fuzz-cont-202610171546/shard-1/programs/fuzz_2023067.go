package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 uint8, p1 int64) string {
    fmt.Println(p1)
    var v6 int = 32767
    _ = p0
    _ = p1
    _ = v6
    return ("\n\t" + "")
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 128
    var v4 uint8 = v3
    var v5 int32 = 42
    delete(v1, (("foo bar" + "\n\t") + ("hello" + "hello")))
    v2 -= 42
    v3 = v4
    if (1 > 493) {
        for fi7 := 0; fi7 < 8; fi7++ {
            if false {
                var v8 int16 = 76
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := []int{286}
                    v8 -= v8
                    fmt.Println(v3, v4, v3)
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
            v11 := []int{fi7, fi7, fi7, (fi7 / ((2147483647) | 1))}
            _ = fi7
            _ = v11
        }
        v12 := GenS1{F0: 1, F1: 145, F2: -9223372036854775808}
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v2 -= 1
            v5 = 0
            _ = fi14
        }
        var v15 uint = 844
        v16 := GenS1{F0: -13, F1: (v4 / (((^(v4))) | 1)), F2: fi13}
        _ = fi13
        _ = v15
        _ = v16
    }
    var v17 int8 = 100
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 13; fi18++ {
            v3 += (v3 << ((65535) & 63))
            _ = fi18
        }
    case "hello":
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: (v17 & 100), F1: 1, F2: int(v17)}
                if (!(v5 > (v5 - (v5 >> ((255) & 63))))) {
                    fmt.Println(v5, v3)
                    fmt.Println(fi19)
                }
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: v17, F1: v3, F2: 0}
            if false {
                v4 = (v3 << ((566) & 63))
            }
            _ = fi19
            _ = v22
        }
    default:
        switch 2147483647 {
        case 159:
            v17 = -1
        case 582:
            switch -1 {
            case 62:
                var v23 float64 = 1.5
                _ = v23
            default:
                v17 = (v17 ^ (^(v17)))
            }
        case 9223372036854775807:
            fmt.Println(v3)
        default:
            v24 := GenS2{F0: 127, F1: "a", F2: true}
            _ = v24
        }
    }
    fmt.Println(v5, v2, v5)
    if (18446744073709551615 == 4294967295) {
        v1[("foo bar" + "")] = -2147483648
        fmt.Println(v4, v4)
    }
    fmt.Println(v2, v17, v3)
    if (4294967295 != 2147483647) {
        delete(v1, "\n\t")
        v25 := GenS2{F0: 42, F1: ("a" + "hello"), F2: (!false)}
        for fi26 := 0; fi26 < 2; fi26++ {
            v2 -= v2
            v17 = v17
            var v27 uint64 = 820
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        var v28 float32 = (3.14159 - (-2.25 / 0.1))
        _ = v28
    }
    var v29 int8 = (v17 & v17)
    v2 = v2
    v29 -= (v17 % ((v17) | 1))
    var v30 float32 = 0.1
    if (("a" + "hello") <= "\n\t") {
        v2 = v2
        if (2147483648 == 0) {
            v3 = v3
            v31 := GenS3{F0: ((("hello" + ("foo bar" + "hello")) + ("hello" + "")) + "hello"), F1: ("foo bar" + ("foo bar" + ""))}
            v2 -= (v2 - (v2 | v2))
            _ = v31
        }
    }
    if false {
        var v32 int64 = v2
        if false {
            v5 -= (v5 - v5)
            v2 = (v2 / ((v2) | 1))
            v17 = v17
        }
        var v33 uint64 = 4294967295
        _ = v32
        _ = v33
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
