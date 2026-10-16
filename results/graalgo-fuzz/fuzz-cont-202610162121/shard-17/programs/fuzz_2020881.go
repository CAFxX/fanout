package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 uint8, p1 string) uint8 {
    p0 += 127
    p0 -= p0
    p1 = (p1 + (p1 + "hello"))
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 = 128
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{-9223372036854775808, (fi6 & (fi6 + (-(fi6)))), (fi6 - 755), fi6}
            p0 = p0
            switch (p1 + "\n\t") {
            case "a":
                v8 = append(v8, (^(fi6)))
            case "\n\t":
                v9 := GenS1{F0: fi7, F1: 65535}
                _ = v9
            default:
                p1 = p1
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return (p0 << ((65535) & 63))
}

func main() {
    var v1 uint16 = 0
    v2 := []int{-110, -2147483648}
    var v3 uint8 = 255
    v4 := make(map[int]string)
    var v5 int16 = 32767
    var v10 int64 = -696
    v11 := GenS2{F0: (v3 << ((4294967295) & 63)), F1: v1}
    fmt.Println(v3)
    var v12 int8 = 1
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            v3 += v3
            v14 := GenS2{F0: v3, F1: (v1 % ((v1) | 1))}
            v4[0] = ("" + ((("\n\t" + "hello") + "\n\t") + (("hello" + "a") + ("" + (("hello" + ("foo bar" + (("hello" + "\n\t") + "foo bar"))) + "")))))
            _ = v14
        }
        v10 = v10
        _ = fi13
    }
    v3 = v3
    var v15 int = 679
    if (0 < 65535) {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = ""
            _ = fi16
            _ = v17
        }
        var v18 bool = false
        v15 += (-(v15))
        _ = v18
    }
    var v19 uint = 616
    fmt.Println(v1)
    v1 -= v1
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS3{F0: 0, F1: "\n\t"}
        for fi22 := 0; fi22 < 4; fi22++ {
            _ = len(v4)
            var v23 int = (fi20 - 4294967295)
            _ = fi22
            _ = v23
        }
        fmt.Println(v19, fi20, v1)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
