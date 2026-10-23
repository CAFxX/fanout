package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int32 = -1
    var v2 string = (("" + "a") + "foo bar")
    var v3 uint32 = 42
    v4 := []int{43, 9223372036854775807, 0}
    v5 := []uint32{v3, v3, (^(v3))}
    v1 -= 982
    var v6 float32 = 1.5
    if (v2 > (v2 + "a")) {
        v5 = append(v5, v3)
        var v7 int64 = -899
        _ = v7
    }
    switch 163 {
    case 4294967295:
        fmt.Println(v1, v6, v3)
    case 2147483647:
        var v8 uint32 = 658
        _ = v8
    case 42:
        v9 := GenS1{F0: 143, F1: -56, F2: v2}
        _ = v9
    }
    var v10 int8 = 1
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: (fi11 % ((849) | 1)), F1: v10, F2: v2}
        if (!(fi11 < (fi11 + int(v1)))) {
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v6, v10)
                v14 := GenS2{F0: 458}
                var v15 int = (fi11 >> ((671) & 63))
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        switch 2147483647 {
        case 2147483647:
            var v16 uint64 = 18446744073709551615
            _ = v16
        case 4294967295:
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint8 = 255
                var v19 string = v2
                if (!(false || false)) {
                    v20 := GenS3{F0: v18, F1: v18}
                    var v21 string = v19
                    _ = v20
                    _ = v21
                } else {
                    switch 213 {
                    case 0:
                        var v22 uint8 = (v18 / ((v18) | 1))
                        _ = v22
                    case 295:
                        var v23 int8 = -128
                        _ = v23
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi11
        _ = v12
    }
    var v24 uint = 924
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            fmt.Println(v10)
            fmt.Println(v24, v6, v24)
            _ = fi26
        }
        var v27 uint64 = 735
        if false {
            v3 = v3
        } else {
            for fi28 := 0; fi28 < 3; fi28++ {
                if true {
                    var v29 int64 = 601
                    switch 42 {
                    case 128:
                        v30 := GenS4{F0: 19}
                        _ = v30
                    default:
                        var v31 uint = 1
                        _ = v31
                    }
                    _ = v29
                }
                if len(v5) > 0 {
                    v5[0] = 1
                }
                fmt.Println(v6, fi28, v1)
                v24 += uint(v3)
                _ = fi28
            }
        }
        _ = fi25
        _ = v27
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 string = ("hello" + (("" + v2) + v2))
        _ = fi32
        _ = v33
    }
    if len(v5) > 0 {
        v5[0] = v3
    }
    v6 += 3.14159
    v10 -= -128
    v24 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
