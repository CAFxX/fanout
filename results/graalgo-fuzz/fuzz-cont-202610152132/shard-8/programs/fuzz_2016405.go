package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint32 = 2
    var v2 int16 = 971
    var v3 uint16 = 745
    v4 := make(map[int]string)
    var v5 string = "foo bar"
    var v6 uint16 = v3
    switch 127 {
    case 200:
        v7 := GenS1{F0: (^(v1))}
        _ = v7
    case 1:
        var v8 int8 = 1
        _ = v8
    default:
        v4[-385] = "hello"
    }
    v3 = 112
    var v9 uint32 = v1
    fmt.Println(v2)
    if true {
        if true {
            v2 = 32767
            switch 127 {
            case 32768:
                var v10 int16 = (v2 ^ (v2 << ((1) & 63)))
                _ = v10
            case 9223372036854775807:
                if false {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        fmt.Println(v3, v2, v5)
                        fmt.Println(v9, v5)
                        _ = fi11
                    }
                } else {
                    fmt.Println(v6, v5, v3)
                    v5 = v5
                }
            default:
                var v12 int32 = int32(v1)
                _ = v12
            }
            var v13 uint32 = v9
            _ = v13
        } else {
            fmt.Println(v1)
            if (0 != 0) {
                v14 := []int{754, -123, 42}
                _ = v14
            } else {
                v3 = 65535
            }
        }
        v5 = v5
        var v15 uint8 = 127
        _ = v15
    }
    var v16 int = -1
    switch (((v5 + v5) + (v5 + (v5 + v5))) + v5) {
    case "a":
        v5 = v5
    case "":
        v4[(v16 % (((v16 - 0)) | 1))] = "foo bar"
    case "foo bar":
        v17 := []int{(v16 ^ v16), 2147483647, -861}
        _ = v17
    }
    if true {
        v18 := GenS1{F0: v1}
        for fi19 := 0; fi19 < 4; fi19++ {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 uint = 42
                fmt.Println(v1, v5)
                fmt.Println(v2, v3, v6)
                fmt.Println(v2)
                _ = fi20
                _ = v21
            }
            var v22 int32 = -1
            v6 = (^(v3))
            _ = fi19
            _ = v22
        }
        _ = v18
    } else {
        v9 = 848
    }
    delete(v4, (v16 - int(v2)))
    v9 = 1
    var v23 uint = 20
    v9 -= v9
    if (!true) {
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v23, fi24)
            var v25 uint32 = (v1 / (((^(v1))) | 1))
            _ = fi24
            _ = v25
        }
        v16 -= 32768
        v5 = v5
    } else {
        if true {
            v26 := v4[v16]
            fmt.Println(v16, v3)
            v27 := []int{(v16 - v16), 451}
            _ = v26
            _ = v27
        } else {
            var v28 int = v16
            _ = v28
        }
        var v29 int = (v16 / (((v16 * (v16 >> ((v23) & 63)))) | 1))
        switch (v16 >> ((uint(v1)) & 63)) {
        case 128:
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS2{F0: (v3 ^ (v3 & 42)), F1: 42}
                _ = fi30
                _ = v31
            }
        case 9223372036854775807:
            v5 = v5
        case 32768:
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 uint8 = 127
                _ = fi32
                _ = v33
            }
        default:
            var v34 float32 = (3.14159 * 1e10)
            _ = v34
        }
        _ = v29
    }
    if ((42 != 1) || false) {
        if (!true) {
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int32 = 794
                _ = fi35
                _ = v36
            }
            var v37 float32 = -2.25
            _ = v37
        } else {
            var v38 uint64 = 1
            delete(v4, 457)
            _ = v38
        }
    }
    if false {
        var v39 float64 = (1e10 / 1.0)
        var v40 uint8 = 217
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
