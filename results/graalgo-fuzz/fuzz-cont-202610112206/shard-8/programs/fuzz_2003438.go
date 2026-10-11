package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint8 }

func main() {
    v1 := []int{-402, 9223372036854775807, 2147483647, 127}
    v2 := []string{}
    var v3 uint64 = 0
    var v4 int8 = 42
    var v5 uint8 = 2
    var v6 int16 = 32767
    v7 := GenS1{F0: (v5 >> ((255) & 63))}
    switch (("" + "\n\t") + ("\n\t" + ("" + "hello"))) {
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS2{F0: -965, F1: (v4 << ((18446744073709551615) & 63))}
            var v10 uint64 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        v2 = append(v2, ("" + "a"))
    }
    if false {
        if (-9223372036854775808 != -585) {
            _ = len(v1)
            v3 -= (^(v3))
        } else {
            fmt.Println(v5, v4)
            v5 -= (v5 | v5)
        }
    }
    if ((65535 >= 222) && false) {
        switch 32768 {
        case 42:
            v2 = append(v2, ("\n\t" + ("" + "hello")))
        }
        var v11 uint32 = 2
        _ = v11
    } else {
        var v12 bool = true
        var v13 float32 = 3.14159
        v12 = v12
        _ = v12
        _ = v13
    }
    var v14 uint8 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 uint32 = 2147483648
    if false {
        var v16 int = -1
        var v17 uint = 18446744073709551615
        _ = v16
        _ = v17
    }
    if false {
        var v18 bool = false
        _ = v18
    }
    v19 := GenS3{F0: 1, F1: 2, F2: v14}
    if (2147483647 > 1) {
        var v20 uint16 = 32767
        _ = v20
    } else {
        var v21 uint16 = 42
        for fi22 := 0; fi22 < 7; fi22++ {
            v14 += (^(v14))
            _ = fi22
        }
        _ = v21
    }
    v4 += (v4 * (-(v4)))
    var v23 int8 = 1
    switch 312 {
    case -2147483648:
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 float32 = ((0.1 / float32(0.1)) - -2.25)
                if false {
                    switch "a" {
                    case "foo bar":
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v27 := GenS3{F0: 220, F1: 255, F2: (v5 >> ((255) & 63))}
                            var v28 int32 = 672
                            v28 = (v28 - -573)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                    var v29 float64 = -2.25
                    v30 := GenS3{F0: 4294967295, F1: 2, F2: v5}
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v3)
                    var v31 int8 = v4
                    v32 := GenS3{F0: 42, F1: 255, F2: v14}
                    _ = v31
                    _ = v32
                }
                _ = fi24
                _ = v25
            }
            v6 = (v6 ^ (-(v6)))
        } else {
            v23 = v23
            var v33 uint = 765
            var v34 uint64 = (v3 - (v3 * v3))
            _ = v33
            _ = v34
        }
    case 128:
        v35 := GenS3{F0: 712, F1: 255, F2: (v14 | v14)}
        _ = v35
    case 0:
        v36 := GenS1{F0: v5}
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
