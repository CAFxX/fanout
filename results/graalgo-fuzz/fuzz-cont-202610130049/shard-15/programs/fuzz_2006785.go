package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int, p1 uint8, p2 int) int {
    p2 += p2
    var v5 string = (("foo bar" + (("foo bar" + "hello") + "foo bar")) + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 uint64, p1 uint32, p2 uint) int {
    p2 += 590
    v6 := GenS1{F0: -287, F1: (false && (p0 <= (p0 | (p0 & p0))))}
    var v7 int64 = -9223372036854775808
    var v8 bool = (p2 <= (p2 & p2))
    v8 = (!(-128 <= -128))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    var v3 string = ("foo bar" + "a")
    var v4 uint16 = 255
    v4 = v4
    if ("\n\t" >= ((v3 + "a") + "foo bar")) {
        switch v3 {
        case "foo bar":
            v2 = v2
        case "hello":
            v4 = 692
        case "\n\t":
            if (29 <= -35) {
                var v9 int16 = 868
                _ = v9
            } else {
                v4 = v4
                v10 := GenS2{F0: v4, F1: 491}
                v11 := v1[-181]
                _ = v10
                _ = v11
            }
        }
        v3 = v3
        switch -1 {
        case 2147483647:
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 uint = 65535
                _ = fi12
                _ = v13
            }
        }
    } else {
        fmt.Println(v2)
        v2 -= 412
        for fi14 := 0; fi14 < 20; fi14++ {
            v3 = ("hello" + v3)
            v15 := GenS3{F0: -710, F1: 9223372036854775807, F2: v4}
            _ = fi14
            _ = v15
        }
    }
    switch v3 {
    case "hello":
        v1[4294967295] = v3
    default:
        v16 := GenS4{F0: 2, F1: 128, F2: (^(v2))}
        _ = v16
    }
    v4 -= (v4 / (((v4 / ((65535) | 1))) | 1))
    v17 := []int{2147483647, genFn0(genFn0(412, 2, genFn0(9223372036854775807, 255, 2147483647)), 163, -2147483648)}
    v2 = uint32(v4)
    _ = len(v17)
    v3 = ((("\n\t" + v3) + v3) + ((v3 + "foo bar") + v3))
    var v18 int = genFn1(42, 433, 0)
    var v19 int = v18
    v3 = "a"
    if false {
        v20 := GenS2{F0: 986, F1: v18}
        v3 = ("\n\t" + "foo bar")
        delete(v1, (v19 | 128))
        _ = v20
    } else {
        v4 = (v4 * v4)
        v3 = v3
    }
    fmt.Println(v3, v3, v4)
    var v21 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
