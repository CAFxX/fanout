package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 int, p2 uint8) int {
    var v4 uint16 = 815
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{}
        if len(v6) > 0 {
            v6[0] = -371
        }
        v7 := GenS1{F0: 767, F1: 4294967295, F2: 2147483647}
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 += p1
    v8 := []int{574, -93}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return p1
}
func genFn1(p0 int16) uint {
    v9 := []int{128, 9223372036854775807, 310}
    v10 := GenS2{F0: 647, F1: 2147483648}
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{977, -1}
    var v3 bool = (true && ((0.1 <= -2.25) || (false && (-1 != 144))))
    var v11 uint64 = 0
    v3 = false
    v12 := GenS3{F0: 32767}
    fmt.Println(v11, v3)
    var v13 string = ""
    var v14 int8 = 100
    switch ((v13 + ("" + ("\n\t" + v13))) + "") {
    case "\n\t":
        for fi15 := 0; fi15 < 19; fi15++ {
            v14 += (v14 + int8(fi15))
            switch v13 {
            case "\n\t":
                if v3 {
                    v3 = false
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint8 = 255
                        switch "foo bar" {
                        case "a":
                            var v18 int32 = -174
                            _ = v18
                        case "foo bar":
                            _ = len(v2)
                        default:
                            v19 := GenS2{F0: (fi16 | -9223372036854775808), F1: 667}
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                    var v20 int16 = 1
                    _ = v20
                }
            case "hello":
                var v21 int = fi15
                _ = v21
            }
            _ = fi15
        }
    case "a":
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 uint64 = v11
            _ = fi22
            _ = v23
        }
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
