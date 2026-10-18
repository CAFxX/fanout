package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint16) int16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        switch -959 {
        case 4294967295:
            v7 := GenS1{F0: p0}
            _ = v7
        case -931:
            v8 := GenS2{F0: false, F1: uint8(p0), F2: 604}
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
    }
    v9 := GenS3{F0: 107}
    fmt.Println(p0)
    v10 := []int{-2147483648, 2147483647}
    _ = p0
    _ = v10
    _ = v9
    return 1000
}
func genFn1(p0 string) int8 {
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint16 = 0
        var v13 uint16 = v12
        v12 = 255
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: 0}
    _ = p0
    _ = v14
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 int = 0
    var v3 string = ""
    var v4 int8 = 0
    v5 := make(map[int]string)
    switch v3 {
    case "foo bar":
        fmt.Println(v1)
    }
    if ((0 >= -1) && false) {
        if true {
            if true {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v4 += -128
                    fmt.Println(v3)
                    fmt.Println(v4, fi15, v4)
                    _ = fi15
                }
            } else {
                v4 -= (v4 >> ((255) & 63))
            }
            v1 = (v1 ^ (v1 ^ -9223372036854775808))
        }
        var v16 bool = true
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := []int{fi17, -320, v2}
        if true {
            v19 := GenS4{F0: genFn0(46)}
            v20 := GenS2{F0: false, F1: 42, F2: 0}
            _ = v19
            _ = v20
        } else {
            fmt.Println(v4, v2, v3)
            var v21 int32 = 0
            _ = v21
        }
        v22 := GenS4{F0: 42}
        _ = fi17
        _ = v18
        _ = v22
    }
    v1 += v1
    var v23 string = v3
    var v24 float32 = (0.1 * ((((1.5 - 1e10) * 1e10) - -2.25) / 3.14159))
    var v25 uint32 = 2147483647
    if true {
        v3 = ("hello" + v23)
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS3{F0: 4294967295}
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                switch v2 {
                case -244:
                    var v30 float32 = 1e10
                    _ = v30
                case 738:
                    v31 := GenS2{F0: (v24 != ((1e10 + v24) - 0.1)), F1: 127, F2: v25}
                    _ = v31
                case 4294967295:
                    v32 := GenS3{F0: 2}
                    _ = v32
                }
                _ = fi29
            }
            switch (v2 & 9223372036854775807) {
            case -647:
                fmt.Println(v25, v2)
            case -86:
                fmt.Println(v25, v1, v1)
            default:
                for fi33 := 0; fi33 < 10; fi33++ {
                    v34 := GenS5{F0: v25, F1: 727}
                    _ = fi33
                    _ = v34
                }
            }
            var v35 uint64 = 304
            _ = fi28
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
