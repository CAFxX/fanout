package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8) uint8 {
    v5 := []int{-2147483648, 2147483647, 32768, -685}
    switch 0 {
    case -2147483648:
        switch 42 {
        case -2147483648:
            _ = len(v5)
        }
    case 32767:
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint32 = 134
            _ = fi6
            _ = v7
        }
    default:
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    _ = p0
    _ = v5
    return 57
}
func genFn1(p0 uint8, p1 uint32) uint {
    v8 := []int{922, 0, 2147483647, -1}
    if false {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: true, F1: (0 < 100), F2: 1}
            switch (fi9 % ((-2147483648) | 1)) {
            case -9223372036854775808:
                if (42 <= 1) {
                    p1 = (^(p1))
                } else {
                    _ = len(v8)
                    var v11 uint32 = p1
                    _ = v11
                }
            case 1:
                fmt.Println(p0, p0)
            }
            fmt.Println(p1, p1, fi9)
            _ = fi9
            _ = v10
        }
    }
    p0 = (p0 * (p0 - (p0 * 255)))
    p0 -= 42
    _ = p0
    _ = p1
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    var v4 int = -9223372036854775808
    v12 := GenS1{F0: false, F1: (true && true), F2: 0}
    fmt.Println(v4)
    v3 += (v3 ^ (v3 * v3))
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v4, v4, v4)
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint = 20
            if true {
                v15 -= (v15 << ((42) & 63))
            }
            var v16 int = -955
            _ = fi14
            _ = v15
            _ = v16
        }
        v4 = -9223372036854775808
        _ = fi13
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int = 4294967295
        fmt.Println(fi17, v18, v4)
        switch (fi17 - (fi17 >> ((255) & 63))) {
        case 0:
            for fi19 := 0; fi19 < 1; fi19++ {
                v4 = v18
                fmt.Println(v4, v3)
                delete(v1, 42)
                _ = fi19
            }
        case 128:
            var v20 float32 = float32(0.1)
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    var v21 int = (v4 - v4)
    switch "hello" {
    case "a":
        switch "foo bar" {
        case "foo bar":
            var v22 bool = (true && ((!true) || ((42 != 127) || true)))
            _ = v22
        case "":
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS2{F0: ("\n\t" + "hello"), F1: 2147483647}
                var v25 float32 = (1e10 + -2.25)
                fmt.Println(v25, v21)
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            v26 := GenS3{F0: 256}
            _ = v26
        }
    case "\n\t":
        fmt.Println(v21, v3)
    default:
        v21 = (v21 * (v4 % ((v21) | 1)))
    }
    delete(v1, 2147483647)
    var v27 int64 = 9223372036854775807
    if true {
        v4 -= v21
    } else {
        v1[65535] = v21
        v28 := []int{v4, (v4 >> ((108) & 63)), (-(v21))}
        _ = v28
    }
    v4 += v4
    v29 := GenS3{F0: 0}
    var v30 int16 = v3
    if (249 > 18446744073709551615) {
        var v31 int = v4
        v21 = v31
        _ = v31
    } else {
        var v32 uint = 504
        v21 -= 272
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := []int{v4, v4, (-(fi33)), (-(fi33))}
            v1[2] = 32767
            var v35 uint32 = 2
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    var v36 int = (v21 | v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
