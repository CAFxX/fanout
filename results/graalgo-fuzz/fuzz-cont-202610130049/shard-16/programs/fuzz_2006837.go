package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 int16) int64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch -1 {
        case 1:
            switch "" {
            case "\n\t":
                if false {
                    p1 -= (p0 - p0)
                    for fi7 := 0; fi7 < 10; fi7++ {
                        if (!false) {
                            v8 := GenS1{F0: false}
                            _ = v8
                        }
                        var v9 int64 = -170
                        if false {
                            var v10 int8 = 1
                            _ = v10
                        } else {
                            var v11 uint8 = 42
                            v12 := []int{fi7, fi6, 32768, (fi6 % ((fi6) | 1))}
                            _ = v11
                            _ = v12
                        }
                        _ = fi7
                        _ = v9
                    }
                }
            case "foo bar":
                if true {
                    var v13 int16 = (-(p1))
                    p1 = (p0 >> ((4294967295) & 63))
                    var v14 uint16 = 2
                    _ = v13
                    _ = v14
                }
            case "":
                var v15 uint8 = 0
                _ = v15
            }
        case 127:
            v16 := GenS1{F0: (true || true)}
            _ = v16
        case -2147483648:
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint = 2
                v19 := GenS2{F0: 1, F1: 2, F2: p0}
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            switch ("a" + "hello") {
            case "foo bar":
                var v20 uint32 = 2147483647
                _ = v20
            }
        }
        v21 := GenS3{F0: 2}
        fmt.Println(fi6, p0)
        _ = fi6
        _ = v21
    }
    var v22 uint32 = 955
    _ = p0
    _ = p1
    _ = v22
    return 1
}

func main() {
    var v1 string = "a"
    v2 := []int{42}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 198
    var v23 string = "hello"
    if len(v2) > 0 {
        v2[0] = -883
    }
    fmt.Println(v5, v23, v23)
    v1 = (((((((v1 + "a") + v1) + ("a" + v1)) + "foo bar") + "\n\t") + v1) + "\n\t")
    if (0 != 2147483648) {
        switch 4294967295 {
        case 4294967295:
            v23 = ("hello" + "")
        }
    }
    var v24 uint32 = 4294967295
    var v25 int64 = 654
    v26 := GenS2{F0: 299, F1: (v24 & (^(v24))), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
