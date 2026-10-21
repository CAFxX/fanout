package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0(p0 uint64, p1 string, p2 uint) string {
    var v4 uint16 = 2
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            p2 -= 0
            _ = fi5
        }
        v6 := GenS1{F0: 32767}
        for fi7 := 0; fi7 < 20; fi7++ {
            if (true || true) {
                var v8 uint16 = (v4 * 0)
                v8 += (v4 + v8)
                _ = v8
            }
            v4 = (v4 >> ((p2) & 63))
            switch -2147483648 {
            case 42:
                p1 = "hello"
            case -521:
                for fi9 := 0; fi9 < 14; fi9++ {
                    v4 -= v4
                    switch 127 {
                    case -1:
                        var v10 string = ""
                        _ = v10
                    case -9223372036854775808:
                        var v11 int32 = -2147483648
                        _ = v11
                    default:
                        p1 = "a"
                    }
                    var v12 uint8 = 128
                    _ = fi9
                    _ = v12
                }
            case -307:
                v13 := GenS1{F0: (fi7 / ((fi7) | 1))}
                _ = v13
            default:
                if (1 < 0) {
                    if true {
                        v14 := []int{(fi7 >> (((p2 * p2)) & 63)), (fi7 + (fi7 * (fi7 ^ (fi7 - fi7))))}
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 string = p1
                            _ = fi15
                            _ = v16
                        }
                        p2 -= p2
                        _ = v14
                    } else {
                        p1 = "hello"
                        for fi17 := 0; fi17 < 14; fi17++ {
                            var v18 int32 = 0
                            _ = fi17
                            _ = v18
                        }
                    }
                    p2 += uint(v4)
                }
            }
            _ = fi7
        }
        _ = v6
    }
    fmt.Println(p0, p1, p0)
    v19 := []int{42}
    fmt.Println(p0, p0, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v4
    return "\n\t"
}
func genFn1(p0 int16, p1 int, p2 int16) uint8 {
    v20 := GenS2{F0: -2147483648, F1: 0}
    v21 := []int{}
    switch "a" {
    case "\n\t":
        v21 = append(v21, p1)
    case "foo bar":
        var v22 uint = 2
        _ = v22
    }
    if (0 <= 100) {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS1{F0: (^(p1))}
            fmt.Println(fi23, p1, p0)
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 int32 = -850
                p0 = 32767
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        var v27 bool = true
        p1 -= (^(p1))
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 2
}

func main() {
    var v1 int16 = 1000
    var v2 int32 = -1
    var v3 uint = 18446744073709551615
    var v28 float32 = 0.1
    v28 = ((v28 - v28) * v28)
    v29 := GenS3{F0: 81, F1: "foo bar", F2: (v1 * -52)}
    var v30 int8 = -128
    v28 = (3.14159 - (((1.5 + float32(-2.25)) / 1e10) * (v28 * 0.1)))
    var v31 uint32 = 2147483647
    fmt.Println(v28, v30)
    if true {
        fmt.Println(v2)
        var v32 uint64 = 4294967295
        v33 := []int{42}
        _ = v32
        _ = v33
    } else {
        v34 := []int{127, 775, 804, 32768}
        _ = v34
    }
    var v35 int = 42
    switch v35 {
    case 0:
        v36 := []int{0, (v35 * v35), (v35 * v35)}
        _ = v36
    case 1:
        var v37 uint8 = uint8(v3)
        _ = v37
    case -335:
        v38 := GenS4{F0: (!(v1 < v1)), F1: 2, F2: 255}
        _ = v38
    default:
        var v39 uint8 = 68
        _ = v39
    }
    fmt.Println(v3)
    var v40 bool = ((-2.25 > (3.14159 / 1.0)) && false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v40
}
