package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 int) bool {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint32 = 4294967295
        var v9 uint16 = 727
        var v10 bool = ((true && false) && false)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        switch (p0 ^ p0) {
        case 776:
            for fi12 := 0; fi12 < 6; fi12++ {
                p0 -= (fi12 % ((p0) | 1))
                v13 := []int{p0, (-(fi11)), (p0 ^ fi12), 32767}
                _ = fi12
                _ = v13
            }
        case -1:
            var v14 uint16 = 65535
            _ = v14
        default:
            switch (p0 % ((p0) | 1)) {
            case -9223372036854775808:
                v15 := GenS1{F0: 1, F1: 119, F2: 1}
                _ = v15
            case 2147483647:
                var v16 int16 = 802
                _ = v16
            }
        }
        p0 = 127
        _ = fi11
    }
    _ = p0
    return false
}
func genFn1(p0 uint16, p1 int64) uint {
    p1 += (p1 ^ p1)
    if true {
        p1 -= p1
    } else {
        switch 565 {
        case 32768:
            v17 := GenS1{F0: -1, F1: 582, F2: -15}
            _ = v17
        case -9223372036854775808:
            var v18 int = 128
            _ = v18
        case 4294967295:
            p1 = (p1 / (((p1 | (p1 | (p1 & (p1 % ((p1) | 1)))))) | 1))
        }
    }
    switch ("hello" + "foo bar") {
    case "":
        switch "" {
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                if false {
                    var v20 uint16 = 255
                    if true {
                        fmt.Println(p1, p0, p0)
                        v21 := GenS2{F0: 1}
                        _ = v21
                    }
                    fmt.Println(fi19, p1)
                    _ = v20
                }
                _ = fi19
            }
        case "foo bar":
            if (!true) {
                p0 = (p0 % ((527) | 1))
                switch 2147483647 {
                case -2147483648:
                    switch "foo bar" {
                    case "\n\t":
                        v22 := []int{713, 2147483647}
                        _ = v22
                    case "a":
                        v23 := GenS1{F0: 1, F1: 255, F2: -1}
                        _ = v23
                    }
                default:
                    v24 := []int{128, int(p0)}
                    _ = v24
                }
                v25 := GenS1{F0: 127, F1: 529, F2: 42}
                _ = v25
            }
        }
    case "\n\t":
        v26 := GenS1{F0: 127, F1: 214, F2: 0}
        _ = v26
    default:
        v27 := GenS3{F0: 100}
        _ = v27
    }
    var v28 int16 = -1
    v29 := []int{42, 128}
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return 0
}

func main() {
    var v1 bool = false
    var v2 bool = (3.14159 > 0.0)
    var v3 uint16 = 0
    var v4 int16 = 990
    v5 := make(map[uint32]uint32)
    v6 := []uint32{2, 2}
    var v30 uint8 = 1
    for fi31 := 0; fi31 < 20; fi31++ {
        v4 += 42
        v3 += v3
        var v32 uint8 = (^(v30))
        _ = fi31
        _ = v32
    }
    v33 := GenS4{F0: v4, F1: (42 <= -1)}
    v3 = (^(v3))
    v34 := GenS4{F0: 1, F1: v2}
    var v35 uint16 = v3
    for fi36 := 0; fi36 < 20; fi36++ {
        if (1.5 <= float32(3.14159)) {
            var v37 float32 = 3.14159
            v38 := GenS3{F0: 0}
            _ = v37
            _ = v38
        } else {
            if (v1 || v1) {
                fmt.Println(fi36)
            } else {
                v3 = 255
                v30 = v30
                if v1 {
                    v30 -= 1
                }
            }
            var v39 uint = 540
            v2 = (v1 && ((v4 << (((v39 - 255)) & 63)) != v4))
            _ = v39
        }
        v40 := v5[994]
        v35 = (v3 * v3)
        _ = fi36
        _ = v40
    }
    var v41 int8 = int8(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
