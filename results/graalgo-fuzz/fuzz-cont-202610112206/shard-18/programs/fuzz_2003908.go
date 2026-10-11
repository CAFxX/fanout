package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    var v4 float32 = 0.0
    v2[0] = (^(v1))
    if true {
        v3 += (^(v3))
    } else {
        var v5 uint64 = v3
        if false {
            if (255 >= 154) {
                v3 -= (v5 >> ((255) & 63))
                v6 := []int{-683, -1, 127, 42}
                var v7 int = -1
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: true}
                var v9 string = "a"
                _ = v8
                _ = v9
            }
        } else {
            if false {
                delete(v2, -9223372036854775808)
            } else {
                if true {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 uint32 = (v1 % (((v1 | (v1 % ((0) | 1)))) | 1))
                            fmt.Println(fi11, v5)
                            delete(v2, (fi10 << ((0) & 63)))
                            var v13 float32 = -2.25
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        _ = fi10
                    }
                    var v14 uint8 = 127
                    v14 = v14
                    _ = v14
                } else {
                    v15 := []int{0, 0}
                    _ = v15
                }
                v16 := []int{42}
                _ = v16
            }
            v4 = v4
            v4 = 0.0
        }
        _ = v5
    }
    var v17 int8 = 127
    switch 657 {
    case 42:
        for fi18 := 0; fi18 < 5; fi18++ {
            v4 -= ((v4 - v4) - (v4 + float32(1.5)))
            _ = fi18
        }
    case 128:
        switch -2147483648 {
        case 127:
            if true {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 int64 = -1
                    var v21 uint32 = v1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 int16 = 1000
                var v23 string = ("\n\t" + "a")
                _ = v22
                _ = v23
            }
        case 2147483647:
            fmt.Println(v1)
        default:
            if false {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v17, v17, v17)
                }
                var v24 string = "\n\t"
                _ = v24
            }
        }
    default:
        if (1 >= 18446744073709551615) {
            if (!false) {
                switch "\n\t" {
                case "\n\t":
                    v25 := GenS1{F0: true}
                    _ = v25
                case "a":
                    var v26 int = 32767
                    _ = v26
                default:
                    var v27 int16 = 32767
                    _ = v27
                }
            } else {
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 uint = 65535
                    _ = fi28
                    _ = v29
                }
                v1 -= (^(v1))
            }
            delete(v2, 174)
        } else {
            v30 := []int{-781}
            switch -2147483648 {
            case 128:
                v17 = v17
            case 127:
                v4 -= v4
            }
            _ = len(v2)
            _ = v30
        }
    }
    var v31 int8 = -128
    switch "foo bar" {
    case "a":
        v32 := GenS1{F0: ((1 >= 258) && ((((566 == 42) && false) || false) && (!false)))}
        _ = v32
    default:
        v31 = v17
    }
    v33 := []int{128}
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 uint32 = 42
        if (false || true) {
            if false {
                v1 = 2147483648
                v1 += v35
                var v36 int16 = -32768
                _ = v36
            } else {
                v33 = append(v33, -20)
                v17 = (-(v17))
                for fi37 := 0; fi37 < 17; fi37++ {
                    v1 -= v35
                    fmt.Println(fi34, v17, v31)
                    _ = fi37
                }
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                var v39 string = ("" + ("\n\t" + ("" + "a")))
                _ = fi38
                _ = v39
            }
        }
        _ = fi34
        _ = v35
    }
    var v40 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v40
}
