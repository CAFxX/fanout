package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
func genFn0() uint {
    var v5 int = 9223372036854775807
    if (true && (-677 < 902)) {
        v6 := []int{127, 2147483647}
        v5 = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v5 -= (v5 / ((fi7) | 1))
            fmt.Println(fi7)
            _ = fi7
        }
    }
    _ = v5
    return 2
}
func genFn1(p0 int, p1 int16) uint {
    p0 += (p0 % (((p0 ^ -330)) | 1))
    v8 := []int{2147483647, 346}
    var v9 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int8 = -28
    v2 := make(map[uint32]bool)
    var v3 uint16 = 906
    var v4 uint8 = 42
    v10 := GenS1{F0: -1, F1: 42, F2: "hello"}
    var v11 int64 = 9223372036854775807
    fmt.Println(v1, v3, v11)
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: v11, F1: 1, F2: ""}
            v1 = (v1 << ((18446744073709551615) & 63))
            var v15 float64 = 1e10
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint16 = v3
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v1 = v1
        _ = fi17
    }
    switch "foo bar" {
    case "foo bar":
        if (("a" != "hello") || false) {
            _ = len(v2)
            v18 := GenS1{F0: 296, F1: 54, F2: ("\n\t" + (((("\n\t" + "a") + "\n\t") + ("a" + "a")) + ("foo bar" + ("a" + "foo bar"))))}
            _ = v18
        }
    case "":
        if false {
            if true {
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch "foo bar" {
                    case "":
                        var v20 uint8 = (v4 & 255)
                        _ = v20
                    case "foo bar":
                        v21 := GenS1{F0: (v11 * v11), F1: 4294967295, F2: "\n\t"}
                        _ = v21
                    }
                    _ = fi19
                }
            } else {
                v22 := v2[261]
                _ = v22
            }
            fmt.Println(v11, v3, v4)
        }
    case "hello":
        v23 := []int{-2147483648, -924, 32768}
        _ = v23
    default:
        v24 := v2[559]
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v1 = (v1 / ((int8(v3)) | 1))
        _ = fi25
    }
    delete(v2, 951)
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
