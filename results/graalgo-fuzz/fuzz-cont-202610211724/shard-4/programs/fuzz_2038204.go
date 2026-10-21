package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool, p1 uint32) int32 {
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: (p1 ^ (p1 >> ((0) & 63))), F1: uint32(v4)}
        v7 := []int{(fi5 % ((0) | 1)), fi5, 0}
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 = (p1 * (p1 - (^(p1))))
    for fi8 := 0; fi8 < 20; fi8++ {
        v4 = 32767
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = (-2.25 * (0.0 * -2.25))
    var v3 int32 = -2147483648
    fmt.Println(v2)
    v2 = v2
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int = fi10
        v12 := []int{fi10, -9223372036854775808, -2147483648, v11}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: uint32(v1), F1: 0}
    var v14 int8 = 12
    v15 := []int{1, 0, -1}
    var v16 int64 = 0
    v1 = v1
    var v17 float64 = (0.0 - (0.0 + (0.1 / 1.5)))
    v17 -= v17
    v14 = v14
    v16 -= v16
    if false {
        for fi18 := 0; fi18 < 14; fi18++ {
            _ = len(v15)
            _ = fi18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 uint8 = 42
            _ = fi19
            _ = v20
        }
        var v21 uint64 = (v1 | (v1 + (v1 & (^(v1)))))
        _ = v21
    } else {
        var v22 float32 = v2
        var v23 float32 = 0.1
        _ = v22
        _ = v23
    }
    var v24 int = -823
    if false {
        v25 := GenS2{F0: -230}
        v1 = 331
        if true {
            for fi26 := 0; fi26 < 15; fi26++ {
                if (true || ("a" == "foo bar")) {
                    switch int(v9) {
                    case 1:
                        v24 = fi26
                    }
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := GenS1{F0: 2, F1: 1}
                    v15 = append(v15, (fi27 ^ fi26))
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            if (false && ((true && (!(!false))) || false)) {
                var v29 int16 = 1
                for fi30 := 0; fi30 < 2; fi30++ {
                    v16 -= v16
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v29 -= -271
                        fmt.Println(v1, v17)
                        v32 := GenS1{F0: 0, F1: 4294967295}
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
                var v33 bool = false
                _ = v29
                _ = v33
            }
            fmt.Println(v24, v9)
        } else {
            if false {
                var v34 uint16 = 32767
                if len(v15) > 0 {
                    v15[0] = v24
                }
                _ = v34
            } else {
                switch v24 {
                case 0:
                    fmt.Println(v24, v14, v2)
                case -2147483648:
                    v1 -= (v1 - v1)
                case 42:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v17 = 0.1
                        v36 := GenS2{F0: 1}
                        var v37 int32 = -2147483648
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                default:
                    for fi38 := 0; fi38 < 2; fi38++ {
                        var v39 string = "hello"
                        var v40 float32 = 3.14159
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
                var v41 int64 = (v16 << ((18446744073709551615) & 63))
                _ = v41
            }
            switch ((("foo bar" + ("a" + "hello")) + "foo bar") + "foo bar") {
            case "foo bar":
                switch v24 {
                case 117:
                    fmt.Println(v9, v3)
                case -262:
                    v14 -= v14
                default:
                    var v42 uint32 = 4294967295
                    _ = v42
                }
            }
            v9 -= (v9 % ((v9) | 1))
        }
        _ = v25
    } else {
        v24 = v24
        var v43 int8 = (v14 - (v14 / (((-(v14))) | 1)))
        var v44 float64 = v17
        _ = v43
        _ = v44
    }
    var v45 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v45
    _ = v9
}
