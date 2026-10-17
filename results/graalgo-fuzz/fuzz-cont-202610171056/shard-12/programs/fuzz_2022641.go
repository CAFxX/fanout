package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 string }
func genFn0(p0 uint8, p1 int32, p2 uint16) uint32 {
    var v5 bool = ((true && (-9223372036854775808 <= 9223372036854775807)) && (!(-25 > -213)))
    v6 := GenS1{F0: 42, F1: (349 <= 18446744073709551615)}
    v7 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    var v1 int8 = 107
    v2 := make(map[string]int)
    var v3 uint8 = 1
    var v4 float64 = 1.5
    v8 := GenS1{F0: 0, F1: (!(false && (false || (!true))))}
    switch 656 {
    case 0:
        fmt.Println(v1, v1)
    case 2147483647:
        if (!((v1 ^ v1) != (v1 << ((18446744073709551615) & 63)))) {
            for fi9 := 0; fi9 < 17; fi9++ {
                v1 -= v1
                if true {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        if false {
                            delete(v2, "foo bar")
                        } else {
                            v11 := []int{-1, fi10, fi10}
                            _ = v11
                        }
                        var v12 string = "hello"
                        fmt.Println(fi10)
                        fmt.Println(fi10, v3)
                        _ = fi10
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v4 = (v4 - v4)
                        v3 += v3
                        _ = fi13
                    }
                    var v14 int8 = (-(v1))
                    _ = v14
                }
                v15 := GenS2{F0: 4294967295, F1: 42}
                _ = fi9
                _ = v15
            }
        } else {
            switch 127 {
            case 42:
                var v16 int16 = -298
                _ = v16
            case -846:
                v17 := GenS3{F0: 1, F1: 0, F2: ("\n\t" + "hello")}
                _ = v17
            default:
                v3 = 128
            }
        }
    case -403:
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int8 = v1
            _ = fi18
            _ = v19
        }
    default:
        if false {
            fmt.Println(v4, v3, v3)
        } else {
            fmt.Println(v1, v4)
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint = 255
                v22 := []int{2147483647, fi20, -9223372036854775808, 42}
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    switch 9223372036854775807 {
    case 32768:
        v23 := GenS1{F0: 2147483647, F1: false}
        _ = v23
    case 128:
        if true {
            if ((float32((1.5 * v4)) == (0.1 - 0.1)) && false) {
                v24 := GenS3{F0: 65535, F1: -1, F2: ("hello" + "foo bar")}
                var v25 uint8 = 127
                fmt.Println(v4, v1, v25)
                _ = v24
                _ = v25
            }
            v3 = (v3 >> ((201) & 63))
            var v26 float32 = 1e10
            _ = v26
        }
    default:
        var v27 int8 = v1
        _ = v27
    }
    v1 = (v1 + (v1 >> ((30) & 63)))
    switch 617 {
    case 127:
        var v28 int64 = -1
        _ = v28
    case -701:
        var v29 int = -2147483648
        _ = v29
    case 444:
        v4 = 1.5
    }
    delete(v2, "a")
    if (!true) {
        fmt.Println(v4)
        fmt.Println(v1, v4, v1)
        fmt.Println(v4, v1, v4)
    }
    v1 = (v1 * v1)
    var v30 float64 = v4
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
