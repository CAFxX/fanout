package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []int{883, 4294967295, 9223372036854775807, -53}
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647, 42}
    var v4 int = 4294967295
    var v5 uint64 = 646
    switch ("hello" + "\n\t") {
    case "a":
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(v4)
            if false {
                if true {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        switch "foo bar" {
                        case "":
                            var v8 uint8 = 255
                            _ = v8
                        case "\n\t":
                            v9 := GenS1{F0: 4294967295}
                            _ = v9
                        default:
                            fmt.Println(fi7)
                        }
                        if false {
                            fmt.Println(v5)
                        }
                        _ = fi7
                    }
                    var v10 float32 = 3.14159
                    v10 -= (3.14159 + (v10 / v10))
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v4 = (v4 | fi11)
                        fmt.Println(v5, v5)
                        var v12 uint8 = 1
                        _ = fi11
                        _ = v12
                    }
                    delete(v2, 2)
                }
                v1 = append(v1, v4)
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(fi6, fi13, v4)
                _ = fi13
            }
            _ = fi6
        }
    }
    var v14 int64 = -656
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v3) > 0 {
            v3[0] = 15
        }
        var v16 bool = (2 != v5)
        _ = fi15
        _ = v16
    }
    var v17 int = (v4 % ((v4) | 1))
    delete(v2, 807)
    if true {
        if true {
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := GenS1{F0: 669}
                switch "foo bar" {
                case "hello":
                    v5 = (v5 | (v5 / ((v5) | 1)))
                case "foo bar":
                    v5 = (v5 << ((255) & 63))
                case "a":
                    fmt.Println(v5, v14, v14)
                }
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                v3 = append(v3, 2147483648)
                v4 = (-(fi20))
                _ = fi20
            }
        }
        var v21 uint16 = 0
        _ = v21
    }
    if true {
        var v22 bool = false
        _ = v22
    }
    var v23 uint64 = v5
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    fmt.Println(v5, v5, v23)
    var v24 uint16 = 0
    var v25 int64 = (-(v14))
    v17 += (v4 + -180)
    v1 = append(v1, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
