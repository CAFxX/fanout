package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 uint }
func genFn0(p0 string, p1 uint32, p2 uint32) int8 {
    switch 127 {
    case 127:
        if (true && (2147483647 >= -1)) {
            v6 := GenS1{F0: (42 >= -174), F1: 127, F2: 253}
            p2 = (p1 * (p1 % ((p1) | 1)))
            _ = v6
        }
    }
    v7 := GenS1{F0: false, F1: 0, F2: 33}
    var v8 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 57
}

func main() {
    v1 := []uint32{65535}
    v2 := []uint32{35, 2, 2147483648}
    var v3 uint16 = 0
    var v4 int64 = -196
    v5 := []uint32{0}
    if len(v1) > 0 {
        v1[0] = uint32(v3)
    }
    var v9 int = -633
    v9 -= 1
    v10 := GenS1{F0: ((-2.25 >= 0.0) && false), F1: 100, F2: 1}
    var v11 float64 = -2.25
    var v12 uint16 = 65535
    if true {
        v13 := GenS2{F0: -128}
        switch v9 {
        case 128:
            fmt.Println(v11)
        case 2147483647:
            var v14 string = "\n\t"
            _ = v14
        case -9223372036854775808:
            switch (v9 ^ v9) {
            case -2147483648:
                for fi15 := 0; fi15 < 19; fi15++ {
                    if true {
                        v3 = v12
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                    } else {
                        fmt.Println(fi15)
                        switch ("hello" + "\n\t") {
                        case "a":
                            var v16 string = ("\n\t" + "hello")
                            _ = v16
                        default:
                            v12 += v12
                        }
                    }
                    _ = fi15
                }
            case 1:
                fmt.Println(v12, v9, v11)
            default:
                var v17 uint32 = 1
                _ = v17
            }
        }
        _ = v13
    }
    v9 -= v9
    if false {
        v18 := GenS1{F0: (4294967295 == 178), F1: 127, F2: 209}
        var v19 string = "\n\t"
        _ = v18
        _ = v19
    }
    var v20 int = v9
    switch ("hello" + "hello") {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v21 int16 = 42
    for fi22 := 0; fi22 < 17; fi22++ {
        v4 = -425
        _ = fi22
    }
    v23 := GenS3{F0: 2147483647, F1: "foo bar", F2: 26}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
