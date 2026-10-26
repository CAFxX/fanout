package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 bool) uint {
    p1 = (-62 <= 4294967295)
    if (!((0 == 128) || (0 != 9223372036854775807))) {
        p0 = "foo bar"
        var v4 int64 = 0
        if ((p0 + "a") >= p0) {
            if true {
                v5 := GenS1{F0: 42, F1: 918}
                _ = v5
            } else {
                for fi6 := 0; fi6 < 15; fi6++ {
                    fmt.Println(fi6, fi6)
                    v7 := GenS2{F0: 42}
                    _ = fi6
                    _ = v7
                }
                v8 := GenS1{F0: 32767, F1: 0}
                v4 -= (v4 ^ 0)
                _ = v8
            }
            v9 := GenS1{F0: 394, F1: 1}
            _ = v9
        } else {
            p1 = (p1 || p1)
        }
        _ = v4
    }
    if (true && (p1 || p1)) {
        if p1 {
            switch 995 {
            case 127:
                var v10 int16 = -444
                _ = v10
            default:
                fmt.Println(p1)
            }
            var v11 int32 = 2147483647
            _ = v11
        }
        p1 = (1 >= 1)
        v12 := []int{127, 32767, 42, 181}
        _ = v12
    }
    _ = p0
    _ = p1
    return 362
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 42
    v3 := make(map[int]uint32)
    var v13 uint32 = 1
    for fi14 := 0; fi14 < 6; fi14++ {
        if (42 < 42) {
            v13 -= (v13 | 4294967295)
            v15 := v1[fi14]
            _ = v15
        } else {
            var v16 int = fi14
            for fi17 := 0; fi17 < 7; fi17++ {
                v13 = v13
                var v18 string = ("a" + ("" + ("" + ("foo bar" + "a"))))
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        v3[(fi14 & (fi14 ^ (fi14 - fi14)))] = v13
        v2 = v2
        _ = fi14
    }
    v1[9223372036854775807] = (2 <= 170)
    var v19 int64 = 42
    switch -817 {
    case 0:
        var v20 int16 = 0
        _ = v20
    case 32767:
        if false {
            var v21 uint16 = 837
            for fi22 := 0; fi22 < 17; fi22++ {
                switch ("" + (("a" + "hello") + "foo bar")) {
                case "hello":
                    v2 += (v2 + (v2 ^ v2))
                default:
                    var v23 int = (fi22 >> ((0) & 63))
                    _ = v23
                }
                switch "\n\t" {
                case "a":
                    var v24 uint64 = (^(v2))
                    _ = v24
                case "":
                    var v25 bool = ("\n\t" <= "hello")
                    _ = v25
                }
                _ = fi22
            }
            _ = v21
        } else {
            var v26 int32 = 2147483647
            _ = v26
        }
    default:
        fmt.Println(v13, v2, v13)
    }
    v27 := GenS3{F0: 100}
    var v28 float64 = 3.14159
    v29 := GenS3{F0: 1}
    v13 = (v13 & (^(v13)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
