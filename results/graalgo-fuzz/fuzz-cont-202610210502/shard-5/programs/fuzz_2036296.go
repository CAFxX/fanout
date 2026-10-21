package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 string, p1 int16) int {
    v5 := GenS1{F0: p1}
    p0 = "foo bar"
    p0 = "foo bar"
    v6 := GenS2{F0: (p0 + "\n\t"), F1: 2, F2: 675}
    v7 := []int{712, -5, int(p1), 785}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint16 = 877
    var v2 int8 = 1
    v3 := make(map[string]uint32)
    var v4 int = int(v2)
    v8 := []int{}
    if false {
        var v9 bool = (false || true)
        fmt.Println(v9)
        v1 += v1
        _ = v9
    }
    v10 := GenS3{F0: 141, F1: 978}
    switch -789 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 18; fi11++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi11)
            case "\n\t":
                var v12 int32 = 2147483647
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 5; fi13++ {
                    fmt.Println(fi11, v4, fi13)
                    if (127 < uint8(v1)) {
                        var v14 uint8 = 1
                        v1 = 915
                        switch "a" {
                        case "":
                            v15 := GenS1{F0: -32768}
                            _ = v15
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v1)
                            }
                        default:
                            var v16 int32 = 940
                            _ = v16
                        }
                        _ = v14
                    }
                    switch v4 {
                    case 1:
                        v17 := GenS4{F0: 42}
                        _ = v17
                    default:
                        var v18 uint8 = 255
                        _ = v18
                    }
                    _ = fi13
                }
            }
            var v19 uint32 = 989
            var v20 int8 = int8(v1)
            _ = fi11
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v4, v4)
    v2 = 127
    v4 = v4
    if false {
        var v21 float64 = 0.1
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 string = ""
            v21 = 0.0
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        v1 -= 32767
        fmt.Println(v4, v2, v2)
        fmt.Println(v1, v4)
    }
    if (false || true) {
        v2 = v2
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 += v1
            _ = fi24
        }
    } else {
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
