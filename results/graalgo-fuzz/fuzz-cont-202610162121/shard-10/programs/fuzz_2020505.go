package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint32; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32 }
func genFn0() uint {
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: "a"}
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: ("" + ("hello" + ("\n\t" + "")))}
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: 42, F1: (!("" <= "hello")), F2: 0}
        _ = fi8
        _ = v12
        _ = v9
    }
    fmt.Println(v7)
    _ = v7
    return 353
}
func genFn1(p0 int32, p1 int8, p2 int8) uint8 {
    var v13 uint32 = 32
    v14 := []int{431, 128, 9223372036854775807, 4294967295}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 128
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 int32 = -2147483648
    v4 := make(map[int]string)
    var v5 uint = 928
    var v6 int32 = v3
    if false {
        v5 -= genFn0()
        v6 = v3
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 4; fi16++ {
                switch ("a" + "hello") {
                case "a":
                    v5 = (v5 + 2)
                case "hello":
                    _ = len(v2)
                default:
                    v17 := GenS3{F0: (fi15 | -1)}
                    _ = v17
                }
                v5 += 255
                v6 = v6
                var v18 uint8 = 0
                _ = fi16
                _ = v18
            }
            _ = fi15
        }
    } else {
        if (false || true) {
            switch ("hello" + ("hello" + (("a" + ("a" + "\n\t")) + "\n\t"))) {
            case "foo bar":
                var v19 uint = v5
                _ = v19
            case "a":
                v3 = (v6 + (v6 - v3))
            default:
                var v20 uint8 = 255
                _ = v20
            }
            v6 += v6
            switch 287 {
            case 42:
                v21 := GenS1{F0: "\n\t"}
                _ = v21
            default:
                v6 = v3
            }
        }
    }
    v6 -= 0
    switch "" {
    case "foo bar":
        v22 := GenS4{F0: 42}
        _ = v22
    case "a":
        v23 := GenS5{F0: ("foo bar" + "hello"), F1: 558, F2: 42}
        _ = v23
    case "":
        v24 := GenS4{F0: 1}
        _ = v24
    default:
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int = (fi25 << ((1) & 63))
            _ = fi25
            _ = v26
        }
    }
    v6 = (v6 << (((v5 << ((847) & 63))) & 63))
    var v27 uint8 = genFn1(v3, -1, 42)
    var v28 uint32 = 196
    var v29 uint16 = 42
    v5 = uint(v28)
    v30 := GenS6{F0: int16(v29)}
    v31 := GenS7{F0: (v3 >> (((v5 << (((v5 / (((v5 % ((65535) | 1))) | 1))) & 63))) & 63))}
    var v32 uint64 = 352
    var v33 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
