package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]uint32)
    var v3 int = 262
    var v4 uint8 = 127
    if ((!false) && false) {
        v4 -= v4
    } else {
        for fi5 := 0; fi5 < 18; fi5++ {
            if (-88 != 0) {
                var v6 uint8 = (v4 >> ((739) & 63))
                _ = v6
            }
            v3 += (v3 * (fi5 << ((2) & 63)))
            _ = fi5
        }
    }
    if true {
        if true {
            v7 := GenS1{F0: 1}
            var v8 uint64 = 882
            fmt.Println(v3, v4, v8)
            _ = v7
            _ = v8
        }
        fmt.Println(v4, v3, v3)
    }
    v9 := GenS2{F0: -235}
    var v10 int = (^(v3))
    var v11 int16 = (v1 & v1)
    v3 = (v10 & 2147483647)
    fmt.Println(v3, v10)
    var v12 uint32 = 2147483647
    var v13 string = "\n\t"
    if false {
        var v14 int32 = 1
        var v15 int32 = v14
        switch 9223372036854775807 {
        case 1:
            v12 -= 287
        case -908:
            for fi16 := 0; fi16 < 18; fi16++ {
                if (!((false && true) || false)) {
                    v14 = 474
                    v15 = 1
                }
                switch 291 {
                case -2147483648:
                    var v17 float64 = 3.14159
                    _ = v17
                default:
                    v18 := []int{127}
                    _ = v18
                }
                _ = fi16
            }
        case 4294967295:
            if true {
                fmt.Println(v10, v13)
                v1 = int16(v4)
                var v19 int64 = 9223372036854775807
                _ = v19
            }
        default:
            var v20 int64 = -825
            _ = v20
        }
        _ = v14
        _ = v15
    } else {
        v21 := GenS2{F0: 0}
        var v22 int16 = v1
        switch 127 {
        case 2147483647:
            var v23 int8 = 127
            _ = v23
        }
        _ = v21
        _ = v22
    }
    var v24 string = v13
    for fi25 := 0; fi25 < 17; fi25++ {
        v24 = (("a" + "foo bar") + (v24 + (v24 + (v13 + v13))))
        v4 = (v4 + 255)
        if false {
            if false {
                v4 = 121
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    if false {
                        v4 -= 128
                        v13 = ("a" + (v24 + v24))
                        v27 := GenS1{F0: 260}
                        var v28 uint32 = 441
                        _ = v27
                        _ = v28
                    }
                    var v29 int16 = (v11 - v11)
                    if true {
                        var v30 float64 = 1.5
                        if (1 <= 42) {
                            v13 = v24
                            fmt.Println(v29)
                        } else {
                            v31 := GenS1{F0: 42}
                            fmt.Println(v30, v29)
                            _ = v31
                        }
                        _ = v30
                    } else {
                        if true {
                            v2[fi25] = (v12 >> ((1) & 63))
                            v32 := []int{v10, 0}
                            fmt.Println(v29, v12)
                            v33 := v2[(fi26 << ((2) & 63))]
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS1{F0: 42}
                        if true {
                            var v35 float64 = ((1.5 + 1e10) + ((0.1 + 1.5) - 1e10))
                            v1 = v1
                            v36 := GenS2{F0: -1}
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                    v2[(fi26 & (v3 >> ((1) & 63)))] = (v12 * v12)
                    _ = fi26
                    _ = v29
                }
            }
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 int = 939
                _ = fi37
                _ = v38
            }
            fmt.Println(v3, v3)
        } else {
            v3 += v10
            if (!(true && false)) {
                _ = len(v2)
                v3 = (fi25 << ((0) & 63))
            }
            var v39 string = (v13 + "foo bar")
            _ = v39
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
