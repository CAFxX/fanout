package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 bool, p1 int64) bool {
    var v6 uint32 = 65535
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return false
}
func genFn1(p0 bool) int32 {
    v7 := []int{-1}
    v8 := GenS1{F0: (1 != 0)}
    if (!(1000 > 32767)) {
        var v9 int16 = 42
        _ = v9
    } else {
        switch -888 {
        case 9223372036854775807:
            p0 = (65535 < 580)
        case 127:
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int = fi10
                var v12 bool = true
                _ = fi10
                _ = v11
                _ = v12
            }
        case -243:
            p0 = (65535 > 1)
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if false {
            if ("a" > "") {
                p0 = p0
            } else {
                p0 = (p0 || (-1 != -231))
                var v14 string = ("" + "")
                _ = v14
            }
            p0 = p0
            if (!p0) {
                v15 := GenS2{F0: 606}
                fmt.Println(p0, p0)
                _ = v15
            }
        }
        v16 := GenS1{F0: false}
        v17 := GenS1{F0: true}
        _ = fi13
        _ = v16
        _ = v17
    }
    if false {
        var v18 uint32 = 2147483647
        for fi19 := 0; fi19 < 6; fi19++ {
            _ = len(v7)
            fmt.Println(v18, fi19, p0)
            for fi20 := 0; fi20 < 1; fi20++ {
                if len(v7) > 0 {
                    v7[0] = (fi20 << ((42) & 63))
                }
                p0 = true
                var v21 uint64 = 42
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        _ = v18
    } else {
        p0 = (1 < -1)
        var v22 uint8 = 255
        _ = v22
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 488
    var v2 float64 = 1.5
    var v3 uint32 = 0
    var v4 uint32 = v3
    v5 := make(map[uint32]int)
    fmt.Println(v1)
    switch -2147483648 {
    case 128:
        v2 = -2.25
    case -2147483648:
        var v23 float64 = 3.14159
        _ = v23
    }
    v2 -= (-2.25 + v2)
    fmt.Println(v1, v2)
    if true {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float64 = (v2 * (float64((0.1 / 1.0)) - 3.14159))
            fmt.Println(fi24, v2)
            if (1 > -605) {
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 bool = true
                    if (!true) {
                        v25 = v2
                    } else {
                        var v28 int8 = -68
                        var v29 int64 = 0
                        _ = v28
                        _ = v29
                    }
                    v25 += 0.0
                    _ = fi26
                    _ = v27
                }
                v4 -= (v3 - (v4 & 42))
                switch "\n\t" {
                case "\n\t":
                    var v30 float64 = 0.0
                    _ = v30
                case "foo bar":
                    v31 := GenS3{F0: (^(v3)), F1: v1, F2: 2}
                    _ = v31
                }
            } else {
                v32 := []int{fi24, (fi24 & fi24), fi24, (fi24 & fi24)}
                var v33 int64 = 9223372036854775807
                v34 := GenS4{F0: 42, F1: int32(v1)}
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi24
            _ = v25
        }
        var v35 int32 = 42
        _ = v35
    }
    var v36 uint16 = 255
    var v37 float32 = 1e10
    fmt.Println(v1)
    v38 := GenS2{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
