package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint }
func genFn0(p0 uint, p1 uint64) uint8 {
    var v4 uint32 = 120
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := []int{(fi5 ^ 32767), (-(fi5)), 4294967295, fi5}
        fmt.Println(fi5)
        for fi7 := 0; fi7 < 13; fi7++ {
            p0 += (p0 >> ((uint(v4)) & 63))
            switch (("" + "") + "") {
            case "\n\t":
                fmt.Println(p0)
            }
            if false {
                switch "foo bar" {
                case "\n\t":
                    if ((float32(0.0) / 1e10) >= 0.1) {
                        p0 -= 42
                        fmt.Println(p1)
                        v8 := GenS1{F0: 42}
                        _ = v8
                    }
                case "a":
                    p0 = (p0 - 642)
                }
                if false {
                    if ((false || true) || true) {
                        v4 = v4
                        v9 := GenS2{F0: 1}
                        _ = v9
                    } else {
                        var v10 float32 = 0.0
                        _ = v10
                    }
                    p0 = p0
                }
                var v11 int = fi5
                _ = v11
            } else {
                var v12 uint = 255
                v12 = (p0 >> (((v12 * 42)) & 63))
                _ = v12
            }
            var v13 uint32 = 0
            _ = fi7
            _ = v13
        }
        _ = fi5
        _ = v6
    }
    var v14 int = -1
    if (false || false) {
        var v15 int16 = 42
        v16 := GenS2{F0: 18446744073709551615}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS2{F0: 0}
        for fi19 := 0; fi19 < 4; fi19++ {
            p1 = p1
            var v20 int16 = -289
            v20 = (v20 & (v20 | -1))
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v4
    return 1
}

func main() {
    v1 := []uint32{4294967295}
    v2 := []int{-9223372036854775808}
    v3 := make(map[int]int)
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint32 = 2
        _ = fi21
        _ = v22
    }
    var v23 uint = 1
    v23 = (v23 >> ((v23) & 63))
    var v24 float64 = 1.5
    v25 := GenS3{F0: -868}
    v26 := GenS1{F0: -428}
    v27 := v3[42]
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint8 = 1
        _ = fi28
        _ = v29
    }
    v24 += float64((0.0 - (3.14159 * -2.25)))
    fmt.Println(v23, v24, v24)
    for fi30 := 0; fi30 < 10; fi30++ {
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v31 uint32 = 4294967295
            _ = v31
        case "foo bar":
            switch (v27 % ((2147483647) | 1)) {
            case 355:
                var v32 uint8 = genFn0(v23, 4294967295)
                _ = v32
            case 32767:
                switch ("\n\t" + ("foo bar" + "\n\t")) {
                case "hello":
                    fmt.Println(v24)
                case "foo bar":
                    v24 = 3.14159
                }
            case -9223372036854775808:
                var v33 int = int(v23)
                _ = v33
            default:
                for fi34 := 0; fi34 < 16; fi34++ {
                    fmt.Println(fi34)
                    _ = fi34
                }
            }
        }
        _ = fi30
    }
    var v35 uint = 177
    fmt.Println(v35, v23)
    v36 := GenS4{F0: 1, F1: 1}
    var v37 int16 = 32767
    v38 := GenS5{F0: 337, F1: 1, F2: v35}
    v27 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
