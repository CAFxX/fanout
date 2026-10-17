package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint16) int8 {
    var v5 int16 = -957
    v6 := []int{}
    v5 = (v5 << ((172) & 63))
    for fi7 := 0; fi7 < 7; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, (fi7 ^ (fi7 | 32768)))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 int) int64 {
    if (2147483647 >= (-(p0))) {
        if (465 > 2147483648) {
            var v8 bool = (2 > 234)
            for fi9 := 0; fi9 < 5; fi9++ {
                if v8 {
                    fmt.Println(v8, fi9, fi9)
                    var v10 int32 = 42
                    _ = v10
                } else {
                    v11 := []int{}
                    _ = v11
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := []int{128, p0, (fi9 + (p0 << ((106) & 63))), (p0 >> ((895) & 63))}
                    _ = fi12
                    _ = v13
                }
                _ = fi9
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 700
                var v16 float32 = -2.25
                if (694 != 65535) {
                    var v17 int8 = 127
                    var v18 int32 = -1
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS1{F0: (("" + "") + ("foo bar" + "a")), F1: v8, F2: 42}
                    _ = v19
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v8
        }
        v20 := GenS2{F0: 32767}
        if false {
            switch -9223372036854775808 {
            case 32767:
                switch "\n\t" {
                case "":
                    var v21 int16 = 1
                    _ = v21
                case "\n\t":
                    switch (^(p0)) {
                    case -849:
                        v22 := GenS3{F0: 65535, F1: 1}
                        _ = v22
                    default:
                        fmt.Println(p0)
                    }
                case "hello":
                    p0 = (p0 & p0)
                }
            case -247:
                p0 -= 32767
            }
            var v23 int = p0
            _ = v23
        }
        _ = v20
    } else {
        fmt.Println(p0)
        var v24 uint64 = 964
        var v25 int8 = -10
        _ = v24
        _ = v25
    }
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 1e10
    var v3 uint64 = 275
    var v4 int16 = 32767
    v26 := GenS4{F0: ""}
    v27 := []int{}
    v28 := GenS5{F0: 1}
    v2 += (v2 + 0.1)
    fmt.Println(v3)
    var v29 uint = 531
    v1 = uint16(v29)
    v1 -= v1
    var v30 int16 = v4
    if (-908 == -997) {
        v30 += (v4 % (((v30 ^ (v30 * (-(v4))))) | 1))
        v31 := GenS2{F0: (v1 | (v1 << ((uint(v1)) & 63)))}
        var v32 int8 = 127
        _ = v31
        _ = v32
    } else {
        var v33 uint64 = (v3 >> ((v29) & 63))
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
