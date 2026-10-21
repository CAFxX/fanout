package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int {
    switch -795 {
    case 0:
        if (0 >= 18446744073709551615) {
            var v4 bool = false
            v5 := []int{2147483647}
            var v6 uint32 = 1
            _ = v4
            _ = v5
            _ = v6
        }
    case 831:
        var v7 string = "foo bar"
        _ = v7
    case -803:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 float64 = ((0.0 * 0.0) - -2.25)
            _ = fi8
            _ = v9
        }
    default:
        if (false && false) {
            var v10 int64 = 1
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    if ("foo bar" == "") {
                        var v13 int16 = 1000
                        var v14 uint16 = 42
                        v15 := []int{(fi12 + fi12), fi12, -190, 659}
                        fmt.Println(fi12, fi11)
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                }
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(fi11, fi16, fi16)
                        default:
                            v17 := []int{(-(fi16)), fi11, (fi11 | fi16), (fi11 / (((^(fi11))) | 1))}
                            _ = v17
                        }
                        if (2147483648 == 2) {
                            v10 = (v10 >> ((222) & 63))
                            var v18 uint32 = 898
                            _ = v18
                        }
                        _ = fi16
                    }
                    fmt.Println(fi11, v10)
                }
                _ = fi11
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v10 += (v10 >> ((65535) & 63))
                _ = fi19
            }
            _ = v10
        }
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(fi20)
        var v21 int16 = 1000
        var v22 uint32 = 65535
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 2
    _ = v23
    return 2147483647
}

func main() {
    v1 := []uint32{551}
    var v2 uint64 = 2
    v3 := make(map[uint32]uint32)
    var v24 int16 = 683
    v24 -= (v24 + (-(v24)))
    if false {
        var v25 int8 = -128
        v25 = 42
        _ = v25
    } else {
        if (1 < -859) {
            v2 -= v2
        }
        if true {
            v24 += -1
        } else {
            _ = len(v3)
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = 649
            }
        }
        var v26 uint16 = 240
        _ = v26
    }
    v24 = 32767
    v27 := GenS1{F0: 28}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v2 | (v2 << ((42) & 63)))
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        v28 := GenS1{F0: 255}
        _ = v28
    default:
        fmt.Println(v24)
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
