package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 uint, p1 int) string {
    p0 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 794
        v9 := GenS1{F0: 0, F1: 0, F2: "hello"}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += 1
    v10 := GenS2{F0: -74}
    v11 := GenS1{F0: -1, F1: 460, F2: "a"}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (("" + "") + "")
}
func genFn1(p0 int16, p1 uint16) int {
    if false {
        fmt.Println(p1, p1)
        p0 += (p0 >> ((255) & 63))
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        switch (-(fi12)) {
        case 9223372036854775807:
            p1 -= p1
        case 4294967295:
            fmt.Println(p1, p1)
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -128
    var v3 string = "\n\t"
    v4 := []string{"foo bar", ("a" + v3), v3}
    var v5 int16 = 912
    var v6 int8 = 100
    v2 -= (v2 + 0)
    var v13 int = 32767
    var v14 uint8 = 127
    v4 = append(v4, v3)
    switch (("foo bar" + (v3 + v3)) + (v3 + v3)) {
    case "\n\t":
        if (2 == 32767) {
            v15 := GenS1{F0: -2147483648, F1: 236, F2: (v3 + (v3 + (v3 + ("foo bar" + "a"))))}
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS3{F0: int8(fi16), F1: v5}
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            v13 = 1
        }
    case "a":
        switch (v3 + v3) {
        case "\n\t":
            var v18 int8 = -111
            _ = v18
        case "foo bar":
            v14 = 137
        default:
            if (((v3 + (((v3 + v3) + v3) + (((v3 + v3) + v3) + (v3 + ("\n\t" + ""))))) <= v3) && (255 == 255)) {
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        _ = len(v1)
                        v3 = "a"
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 int32 = 897
                        _ = fi21
                        _ = v22
                    }
                    _ = fi19
                }
            }
        }
    case "hello":
        if (952 <= 42) {
            v6 = int8(v5)
            if len(v4) > 0 {
                v4[0] = (v3 + "a")
            }
        }
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            fmt.Println(v13, fi23, v2)
            fmt.Println(v6)
            fmt.Println(v13, v6, v2)
            _ = fi23
        }
    }
    var v24 uint32 = 2147483648
    var v25 string = v3
    v3 = v3
    v26 := GenS4{F0: 904, F1: 42}
    v4 = append(v4, (((("" + "hello") + v3) + "a") + (v25 + v3)))
    for fi27 := 0; fi27 < 4; fi27++ {
        v6 -= 100
        _ = fi27
    }
    v5 = int16(v24)
    delete(v1, (v24 % ((v24) | 1)))
    var v28 string = ("a" + (v25 + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
