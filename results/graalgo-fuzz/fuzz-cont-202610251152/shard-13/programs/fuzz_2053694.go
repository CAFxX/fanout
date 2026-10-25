package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    if (false && true) {
        var v4 uint32 = 978
        _ = v4
    } else {
        if (("" + "hello") != "") {
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := GenS1{F0: 32767}
                switch ("a" + ((("" + "a") + "") + "hello")) {
                case "a":
                    for fi7 := 0; fi7 < 3; fi7++ {
                        switch 127 {
                        case -1:
                            delete(v2, 2147483647)
                        case 32768:
                            fmt.Println(v1, fi7, fi5)
                        case -2147483648:
                            var v8 int16 = v1
                            _ = v8
                        default:
                            var v9 int8 = -1
                            _ = v9
                        }
                        _ = fi7
                    }
                default:
                    var v10 bool = ((((0.0 - (1.5 + 3.14159)) * (((-2.25 * (-2.25 - -2.25)) / 1e10) - 0.1)) / 3.14159) != -2.25)
                    _ = v10
                }
                if false {
                    fmt.Println(v1, fi5)
                } else {
                    var v11 int8 = -38
                    v12 := GenS1{F0: 32767}
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v14 := GenS2{F0: (^(v1)), F1: 0}
                        var v15 int64 = 42
                        v16 := GenS1{F0: 255}
                        v17 := v2[4294967295]
                        fmt.Println(v15, fi13)
                        v17 = v17
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v11
                    _ = v12
                }
                _ = fi5
                _ = v6
            }
            if true {
                v1 = v1
            } else {
                v18 := GenS3{F0: -1}
                fmt.Println(v1)
                _ = v18
            }
        } else {
            var v19 int32 = 2147483647
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 uint8 = 196
                v21 -= 1
                fmt.Println(v1, v19)
                _ = fi20
                _ = v21
            }
            v22 := GenS2{F0: v1, F1: v19}
            _ = v19
            _ = v22
        }
    }
    fmt.Println(v1)
    var v23 uint16 = 135
    var v24 uint8 = 90
    var v25 int32 = 42
    for fi26 := 0; fi26 < 17; fi26++ {
        v1 = v1
        var v27 int32 = (v25 << ((42) & 63))
        v28 := []int{(^(fi26)), fi26}
        _ = fi26
        _ = v27
        _ = v28
    }
    delete(v3, 2147483647)
    fmt.Println(v25)
    var v29 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
