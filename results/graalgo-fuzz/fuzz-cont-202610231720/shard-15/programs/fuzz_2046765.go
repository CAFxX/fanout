package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8, p1 string, p2 int16) int32 {
    p2 = (p2 - 0)
    var v7 uint8 = 141
    if true {
        switch -93 {
        case 32767:
            p2 = (p2 * p2)
        }
    }
    v8 := GenS1{F0: p1, F1: (2 > 2), F2: 65535}
    switch (p1 + "foo bar") {
    case "foo bar":
        v9 := GenS2{F0: 255, F1: int32(v7)}
        _ = v9
    case "\n\t":
        v10 := GenS3{F0: 42, F1: (20 >= 18446744073709551615), F2: (!false)}
        _ = v10
    default:
        var v11 float64 = -2.25
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -210
}

func main() {
    v1 := []uint32{}
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[uint32]bool)
    var v4 int8 = 0
    var v5 float64 = 0.1
    v6 := make(map[uint32]uint32)
    if false {
        var v12 int8 = (v4 % ((1) | 1))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int32 = 258
            _ = fi13
            _ = v14
        }
        if true {
            for fi15 := 0; fi15 < 11; fi15++ {
                if false {
                    var v16 float32 = (float32(v5) / 3.14159)
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if (true || (3.14159 > (v16 * 3.14159))) {
                            fmt.Println(v16, v5)
                        }
                        _ = fi17
                    }
                    _ = v16
                }
                var v18 int8 = 100
                _ = fi15
                _ = v18
            }
            var v19 string = v2
            v20 := GenS4{F0: 0}
            _ = v19
            _ = v20
        } else {
            var v21 float32 = (1e10 - ((0.0 * 1e10) * (3.14159 * 0.0)))
            v2 = (("hello" + (v2 + v2)) + "\n\t")
            _ = v21
        }
    }
    v22 := GenS3{F0: 42, F1: (((2 != 247) && (true && (false || false))) || ((v4 >> ((620) & 63)) < v4)), F2: true}
    switch 156 {
    case 762:
        var v23 uint32 = 4294967295
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = 305
    }
    v4 = v4
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            if len(v1) > 0 {
                v1[0] = 233
            }
            v26 := GenS4{F0: 918}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v27 int = 32767
    if true {
        v4 += v4
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(fi28)
        _ = fi28
    }
    if ((v4 - v4) <= 1) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v2 = (v2 + v2)
        v27 += v27
        switch (v27 * v27) {
        case -1:
            switch v27 {
            case -2147483648:
                v2 = v2
            case -603:
                v5 += (v5 - -2.25)
            }
        case 376:
            fmt.Println(v2, v4)
        }
    }
    v27 = (^(v27))
    v5 += v5
    var v29 string = ("\n\t" + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
