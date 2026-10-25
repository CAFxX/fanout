package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int64; F1 int16 }
func genFn0(p0 int64) uint32 {
    p0 += (p0 | (p0 ^ (p0 + (p0 | p0))))
    v6 := []int{32768, 447, 0}
    for fi7 := 0; fi7 < 4; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v8 bool = (!(66 < -1))
        v9 := GenS1{F0: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = append(v6, 1)
    _ = p0
    _ = v6
    return 590
}

func main() {
    v1 := []string{("" + "foo bar"), (("hello" + "\n\t") + "hello"), "a"}
    var v2 uint64 = 1
    v3 := []int{2147483647, 9223372036854775807}
    var v4 int32 = -955
    v5 := make(map[uint32]int)
    v10 := GenS2{F0: 255, F1: -1}
    v2 -= v2
    v11 := GenS3{F0: -1}
    v4 += v4
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v2 = v2
            v4 -= int32(fi12)
            if true {
                var v14 uint64 = 2
                var v15 uint32 = 42
                if true {
                    v16 := GenS4{F0: uint(v2), F1: ("\n\t" + "\n\t")}
                    v14 -= (v14 + (v2 * (v14 / (((v2 >> ((374) & 63))) | 1))))
                    v2 -= (^(v14))
                    _ = v16
                } else {
                    fmt.Println(v14, v15, v4)
                }
                _ = v14
                _ = v15
            }
            _ = fi13
        }
        _ = fi12
    }
    switch ("foo bar" + "a") {
    case "hello":
        for fi17 := 0; fi17 < 16; fi17++ {
            switch "\n\t" {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = ""
                }
            case "":
                v2 -= (^(v2))
            case "foo bar":
                fmt.Println(fi17, fi17)
            }
            v18 := GenS5{F0: 164, F1: 1}
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int8 = 1
                var v21 string = ""
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
    }
    switch "hello" {
    case "a":
        v22 := GenS3{F0: 0}
        _ = v22
    case "\n\t":
        v23 := GenS4{F0: 1, F1: "a"}
        _ = v23
    }
    switch 32768 {
    case -459:
        var v24 int8 = 122
        _ = v24
    case 319:
        v25 := GenS6{F0: -691, F1: 0}
        _ = v25
    case -29:
        switch "foo bar" {
        case "a":
            var v26 int64 = -526
            _ = v26
        default:
            for fi27 := 0; fi27 < 12; fi27++ {
                v28 := GenS7{F0: -1, F1: 32767}
                v29 := GenS4{F0: 18446744073709551615, F1: (((("hello" + "foo bar") + "foo bar") + "a") + ("foo bar" + ("hello" + ("foo bar" + ""))))}
                var v30 uint16 = 0
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
    default:
        var v31 uint8 = 255
        _ = v31
    }
    var v32 uint8 = 2
    delete(v5, 2147483648)
    v33 := v5[4294967295]
    v2 = v2
    v4 = (v4 & int32(v32))
    v33 -= v33
    v32 += v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
