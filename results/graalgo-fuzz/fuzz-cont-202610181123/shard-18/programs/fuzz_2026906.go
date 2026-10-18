package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8, p1 string, p2 uint64) uint32 {
    var v5 uint8 = (p0 * 136)
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, (fi6 >> ((18446744073709551615) & 63)), (fi6 | fi6), fi6}
        var v8 int64 = 0
        if true {
            switch int(p2) {
            case 1:
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(p0, p0, p2)
                    v10 := GenS1{F0: (v8 % ((-1) | 1))}
                    fmt.Println(fi6)
                    fmt.Println(v5, p0, p2)
                    _ = fi9
                    _ = v10
                }
            case 32768:
                if (127 <= 103) {
                    fmt.Println(p0, v8, v8)
                    fmt.Println(fi6)
                }
            }
            var v11 float32 = float32(((float64(1e10) / ((-2.25 + 3.14159) / (1e10 * ((0.1 / -2.25) / ((0.1 - (0.0 * 3.14159)) - 1.5))))) * 3.14159))
            var v12 bool = (!true)
            _ = v11
            _ = v12
        } else {
            var v13 uint64 = (p2 ^ 264)
            fmt.Println(fi6, v8, v5)
            if (255 <= 2) {
                p2 = p2
            }
            _ = v13
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 333
}

func main() {
    var v1 int8 = -6
    var v2 uint16 = 32767
    var v3 uint16 = v2
    v4 := make(map[int]int)
    v2 = (v3 >> ((65535) & 63))
    v1 = (-(v1))
    if (v3 <= v3) {
        switch "hello" {
        case "foo bar":
            v14 := GenS2{F0: 100}
            _ = v14
        case "a":
            v3 = (v3 ^ v2)
        default:
            v2 = (^(v3))
        }
    }
    var v15 uint8 = uint8(v2)
    if (-748 > 2147483647) {
        if true {
            var v16 int16 = 282
            var v17 int64 = -727
            _ = v16
            _ = v17
        }
    } else {
        v18 := []int{42, 1}
        _ = v18
    }
    if false {
        if true {
            if true {
                var v19 uint8 = v15
                v20 := GenS3{F0: ""}
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 bool = true
                v1 = (v1 + v1)
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := []int{fi21}
                    _ = fi23
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
        } else {
            v4[128] = -82
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            v2 = (v3 & 773)
            switch (fi25 - fi25) {
            case 127:
                var v26 float64 = ((((((0.0 - (float64(0.1) + -2.25)) * ((((3.14159 + (1.5 / (-2.25 / 1.0))) - 0.1) - 3.14159) - 0.0)) - 1.5) / (0.1 - -2.25)) / 1.5) + 3.14159)
                _ = v26
            case -693:
                var v27 uint = 93
                _ = v27
            case 533:
                v28 := GenS2{F0: (v1 ^ v1)}
                _ = v28
            default:
                if true {
                    v29 := GenS4{F0: (v15 ^ (v15 << ((160) & 63))), F1: 42}
                    switch "a" {
                    case "hello":
                        v4[(fi25 % (((fi25 ^ -9223372036854775808)) | 1))] = fi25
                    case "a":
                        if true {
                            v1 = (v1 * v1)
                            v15 = (v15 / ((v15) | 1))
                            for fi30 := 0; fi30 < 19; fi30++ {
                                v2 -= (^(v3))
                                v31 := GenS1{F0: 0}
                                var v32 uint = 18446744073709551615
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        } else {
                            var v33 bool = (true && true)
                            _ = v33
                        }
                    case "\n\t":
                        fmt.Println(fi25)
                    }
                    _ = v29
                }
            }
            var v34 uint32 = 0
            _ = fi25
            _ = v34
        }
    }
    var v35 uint16 = v3
    switch "hello" {
    case "foo bar":
        _ = len(v4)
    case "":
        v2 -= (^(v3))
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        if (!(1 != 4294967295)) {
            switch (-(fi36)) {
            case -542:
                switch "\n\t" {
                case "hello":
                    v1 -= (v1 & (v1 | v1))
                case "a":
                    v3 = 918
                case "\n\t":
                    v37 := GenS2{F0: (v1 * v1)}
                    _ = v37
                }
            default:
                delete(v4, (fi36 << ((367) & 63)))
            }
            fmt.Println(v2)
            if (true && false) {
                v38 := GenS4{F0: (v15 << ((686) & 63)), F1: (^(v15))}
                _ = v38
            }
        } else {
            v1 = 1
        }
        _ = fi36
    }
    v15 = v15
    v2 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
