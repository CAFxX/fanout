package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    v7 := GenS1{F0: "a"}
    v8 := []int{9223372036854775807}
    var v9 uint = 4294967295
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int32, p1 bool) uint16 {
    v10 := GenS2{F0: (1 > 0)}
    if p1 {
        fmt.Println(p1, p1)
        v11 := GenS2{F0: p1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    return 255
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 42
    var v3 string = "foo bar"
    v4 := []string{(("foo bar" + "foo bar") + v3), v3, (((v1 + v3) + "a") + v3)}
    v5 := make(map[int]string)
    v6 := []int{-2147483648, 493, 128, -71}
    v5[-796] = v3
    var v12 int64 = 952
    fmt.Println(v3, v1)
    v13 := GenS2{F0: (255 <= 1)}
    v1 = (("hello" + ("" + "foo bar")) + (v1 + (v1 + (v1 + v3))))
    v12 += -761
    _ = len(v4)
    if true {
        v2 += v2
        switch -441 {
        case -2147483648:
            v14 := GenS1{F0: ""}
            _ = v14
        case -265:
            v2 += (v2 - 2)
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v2 = v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: (-32768 >= -32768)}
            v12 = -78
            switch (("a" + ("a" + "a")) + v3) {
            case "\n\t":
                fmt.Println(v1, fi15)
            default:
                switch (fi16 ^ fi16) {
                case -1:
                    switch fi15 {
                    case 42:
                        v2 += 214
                    case -754:
                        if (2147483647 == 42) {
                            switch fi15 {
                            case -1:
                                var v18 int16 = -32768
                                _ = v18
                            case 0:
                                v19 := GenS1{F0: (v3 + "foo bar")}
                                _ = v19
                            }
                            fmt.Println(v2, fi15, v12)
                        }
                    default:
                        v3 = v3
                        v20 := GenS1{F0: (v1 + v3)}
                        _ = v20
                    }
                default:
                    fmt.Println(fi16, v3)
                }
            }
            switch fi15 {
            case -406:
                delete(v5, -349)
            case 2147483647:
                if len(v4) > 0 {
                    v4[0] = (v3 + v3)
                }
            case 32768:
                v21 := GenS3{F0: 1000}
                _ = v21
            }
            switch fi15 {
            case 61:
                var v22 uint8 = 64
                _ = v22
            default:
                v1 = (v3 + v3)
            }
            _ = len(v6)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    delete(v5, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
