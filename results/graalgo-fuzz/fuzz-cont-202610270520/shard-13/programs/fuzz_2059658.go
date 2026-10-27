package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    if true {
        if (544 == 106) {
            p0 -= (-(p0))
        }
        p0 += 32767
        v5 := []int{(-(p0))}
        _ = v5
    }
    _ = p0
    return 403
}
func genFn1(p0 int32, p1 uint8, p2 uint) uint {
    fmt.Println(p2)
    p0 += 42
    _ = p0
    _ = p1
    _ = p2
    return (p2 - (^(p2)))
}

func main() {
    v1 := []uint32{303, 0}
    v2 := make(map[uint32]int)
    var v3 string = (("foo bar" + "hello") + "")
    var v4 bool = true
    if (!v4) {
        v4 = false
        for fi6 := 0; fi6 < 17; fi6++ {
            v3 = ""
            fmt.Println(v3)
            _ = fi6
        }
    } else {
        v3 = v3
        if (-99 == 0) {
            v7 := GenS1{F0: -292}
            _ = v7
        } else {
            var v8 int16 = -1
            _ = v8
        }
    }
    if true {
        if v4 {
            if false {
                var v9 int32 = genFn0(-9223372036854775808)
                _ = len(v2)
                _ = len(v2)
                _ = v9
            }
        }
        var v10 int16 = -32768
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            v2[2147483648] = (^(fi11))
            v4 = v4
            var v12 uint32 = 2147483647
            _ = fi11
            _ = v12
        }
    }
    var v13 int32 = -278
    for fi14 := 0; fi14 < 14; fi14++ {
        if v4 {
            var v15 float64 = 1e10
            v13 = (v13 % ((int32(fi14)) | 1))
            var v16 int64 = 1
            _ = v15
            _ = v16
        } else {
            delete(v2, 2147483648)
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int = fi17
                v1 = append(v1, 42)
                fmt.Println(v13, fi14)
                fmt.Println(v4, v13)
                v3 = "a"
                v13 -= (v13 ^ v13)
                switch 127 {
                case -9223372036854775808:
                    var v19 uint32 = 2
                    _ = v19
                }
                var v20 uint = 42
                if true {
                    switch v3 {
                    case "\n\t":
                        _ = len(v2)
                    case "foo bar":
                        var v21 uint8 = 128
                        _ = v21
                    case "hello":
                        fmt.Println(fi17, fi17, v4)
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v18 -= (fi17 & fi14)
                    var v22 uint16 = 42
                    _ = v22
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v23 bool = ((((0.1 < 1e10) || v4) && true) && v4)
                v20 = (v20 | (v20 - v20))
                if (true || (!false)) {
                    var v24 uint16 = 1
                    _ = v24
                } else {
                    v23 = true
                    var v25 uint32 = 709
                    _ = v25
                }
                _ = fi17
                _ = v18
                _ = v20
                _ = v23
            }
        }
        v13 = (v13 >> ((genFn1((v13 | v13), 0, 4294967295)) & 63))
        v13 -= 42
        _ = fi14
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int8 = 127
        _ = fi26
        _ = v27
    }
    fmt.Println(v4)
    v28 := v2[65535]
    v29 := GenS2{F0: 4294967295, F1: 115}
    if false {
        fmt.Println(v4, v28, v13)
    }
    var v30 int = (v28 >> ((79) & 63))
    v28 = v30
    v31 := GenS3{F0: 0, F1: (v4 && v4), F2: 420}
    for fi32 := 0; fi32 < 7; fi32++ {
        v13 = (-(v13))
        var v33 string = v3
        v13 += v13
        _ = fi32
        _ = v33
    }
    v3 = ("hello" + (v3 + (("hello" + v3) + ("foo bar" + ((("hello" + v3) + ("hello" + v3)) + v3)))))
    v34 := GenS1{F0: (v30 * (v28 & 127))}
    var v35 int16 = 32767
    var v36 uint32 = 796
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
