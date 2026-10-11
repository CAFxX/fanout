package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int = 561
    v2 := make(map[string]uint32)
    var v3 int64 = -71
    var v4 int16 = -673
    var v5 uint64 = 0
    switch 37 {
    case 712:
        v6 := GenS1{F0: 255, F1: (v5 + v5)}
        _ = v6
    case 32768:
        var v7 uint = 533
        _ = v7
    case 4294967295:
        var v8 float64 = 3.14159
        _ = v8
    }
    var v9 uint8 = 255
    v10 := GenS1{F0: 65535, F1: 0}
    var v11 bool = false
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 bool = v11
        if ((v1 >> ((65535) & 63)) > v1) {
            if v11 {
                switch -2147483648 {
                case -2147483648:
                    v4 = (-(v4))
                default:
                    switch (v1 & fi12) {
                    case 1:
                        switch "\n\t" {
                        case "":
                            var v14 int8 = 100
                            _ = v14
                        case "foo bar":
                            v5 += 18446744073709551615
                        case "\n\t":
                            fmt.Println(v4, v9)
                        }
                    }
                }
                v3 = (v3 | v3)
                v15 := GenS1{F0: 42, F1: v5}
                _ = v15
            } else {
                v16 := GenS2{F0: 6}
                _ = v16
            }
            if (9223372036854775807 < v3) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v5 = v5
                    if v13 {
                        var v18 string = "foo bar"
                        var v19 int32 = int32(v4)
                        _ = v18
                        _ = v19
                    } else {
                        v3 = v3
                        var v20 int8 = 127
                        _ = v20
                    }
                    _ = fi17
                }
                v21 := []int{fi12}
                _ = v21
            } else {
                var v22 bool = ((v5 - v5) < v5)
                var v23 uint = 616
                v24 := []int{(fi12 - fi12), fi12}
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 bool = ((v1 + fi12) != v1)
            _ = v25
        } else {
            var v26 int = (v1 * (^(v1)))
            v4 -= (v4 | v4)
            _ = v26
        }
        _ = fi12
        _ = v13
    }
    v27 := []int{v1}
    switch ("hello" + "\n\t") {
    case "hello":
        v28 := GenS3{F0: (v4 % (((v4 | int16(v9))) | 1)), F1: 65535}
        _ = v28
    case "foo bar":
        var v29 uint64 = (^(v5))
        _ = v29
    default:
        var v30 int = -113
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int = (fi31 + fi31)
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
