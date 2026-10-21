package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint, p1 string, p2 int8) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int = (fi8 % ((-467) | 1))
            v10 := GenS1{F0: 2, F1: 1, F2: int(p0)}
            p1 = (p1 + p1)
            _ = fi8
            _ = v10
            _ = v9
        }
        if false {
            p1 = p1
            var v11 int16 = 1000
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := []int{fi12, (fi7 - (fi12 ^ fi7)), (^(fi12))}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi7
    }
    v14 := GenS2{F0: uint16(p0), F1: 127, F2: 127}
    if (!(!(!(65535 <= 2)))) {
        if (796 != 261) {
            var v15 int8 = p2
            p1 = (((("" + ("" + ((p1 + p1) + p1))) + ("foo bar" + p1)) + (("hello" + p1) + p1)) + p1)
            fmt.Println(p0, p0, p0)
            _ = v15
        } else {
            var v16 uint32 = 42
            p0 = (p0 << ((p0) & 63))
            v17 := []int{-961, 128, 32768}
            _ = v16
            _ = v17
        }
        var v18 uint16 = 32767
        _ = v18
    } else {
        p1 = p1
        if (18446744073709551615 < 2) {
            var v19 uint8 = 127
            _ = v19
        } else {
            switch p1 {
            case "a":
                if true {
                    p2 = 100
                    v20 := GenS2{F0: 625, F1: 42, F2: -460}
                    _ = v20
                }
            case "hello":
                for fi21 := 0; fi21 < 13; fi21++ {
                    fmt.Println(p2, p2)
                    var v22 int8 = (p2 + 0)
                    _ = fi21
                    _ = v22
                }
            case "":
                fmt.Println(p1, p0, p0)
            }
            fmt.Println(p2, p2)
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 127
}

func main() {
    var v1 int64 = 1
    var v2 uint8 = 103
    var v3 uint16 = 0
    var v4 int8 = 1
    v5 := make(map[string]int)
    var v6 float32 = ((1.5 - 0.1) - float32(0.1))
    v1 += (v1 + (v1 & v1))
    var v23 int64 = (v1 - v1)
    if false {
        _ = len(v5)
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 float32 = v6
            switch (fi24 << ((uint(v2)) & 63)) {
            case -933:
                var v26 uint16 = v3
                _ = v26
            default:
                v27 := []int{fi24, -2147483648, -438}
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            v6 += v6
            _ = fi28
        }
    }
    var v29 int16 = -32768
    var v30 string = ("hello" + "")
    v30 = (((v30 + ((((v30 + (("hello" + v30) + "hello")) + v30) + v30) + "foo bar")) + "\n\t") + v30)
    v5[v30] = -284
    v31 := GenS3{F0: 4294967295, F1: 0}
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 uint32 = 0
        var v34 uint16 = (^(v3))
        var v35 int8 = v4
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
