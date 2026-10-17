package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 string, p1 int32, p2 int64) uint64 {
    v7 := []int{128, 42, 0, -9223372036854775808}
    fmt.Println(p2)
    p0 = "\n\t"
    switch "foo bar" {
    case "a":
        var v8 int64 = (-(p2))
        _ = v8
    case "foo bar":
        if true {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: 1, F1: 0}
                _ = fi9
                _ = v10
            }
            p0 = ("foo bar" + p0)
        }
    case "":
        var v11 int16 = 1
        _ = v11
    default:
        if true {
            if true {
                p1 = 276
            }
        } else {
            var v12 bool = (p0 != p0)
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = 1e10
    var v2 float32 = 1e10
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    var v6 bool = (!true)
    if (!true) {
        switch "\n\t" {
        case "foo bar":
            v2 = (float32(3.14159) + (float32(float64((float32(3.14159) + v2))) * float32((-2.25 - 1e10))))
        case "":
            var v13 uint8 = 255
            _ = v13
        case "a":
            if ((1e10 * 0.1) != (1.5 / 3.14159)) {
                var v14 uint64 = genFn0((((("foo bar" + ("hello" + "\n\t")) + ("a" + ("hello" + "foo bar"))) + "a") + "hello"), 2147483647, 9223372036854775807)
                _ = v14
            }
        }
        if true {
            for fi15 := 0; fi15 < 1; fi15++ {
                v2 += 0.0
                _ = fi15
            }
            fmt.Println(v2)
            v16 := GenS1{F0: 0, F1: 1}
            _ = v16
        }
    } else {
        v1 += float32(float64((1.5 - (v1 + float32(1.5)))))
    }
    var v17 uint = 4294967295
    switch 4294967295 {
    case -2147483648:
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 uint16 = 1
            _ = fi18
            _ = v19
        }
    case 0:
        v20 := []int{624, 127, -918, -1}
        _ = v20
    case 32767:
        v21 := []int{32767, 32767}
        _ = v21
    default:
        var v22 uint64 = 65
        _ = v22
    }
    v2 -= 3.14159
    if v6 {
        v2 = -2.25
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS1{F0: 128, F1: 637}
            _ = fi23
            _ = v24
        }
    }
    var v25 uint16 = 32767
    v6 = (9223372036854775807 != 986)
    if (!((^(v25)) == (v25 ^ (^(v25))))) {
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                v25 = v25
                if (false || v6) {
                    fmt.Println(v2)
                    v3[("hello" + "a")] = fi27
                }
                v28 := GenS2{F0: v17, F1: 4294967295, F2: v6}
                _ = fi27
                _ = v28
            }
            v6 = (42 <= -941)
            var v29 uint16 = (^(v25))
            _ = fi26
            _ = v29
        }
    } else {
        switch 32768 {
        case 9223372036854775807:
            v30 := GenS3{F0: 0, F1: 127}
            _ = v30
        case 674:
            if (!v6) {
                var v31 string = ""
                v2 -= v2
                _ = v31
            } else {
                v32 := []int{}
                for fi33 := 0; fi33 < 14; fi33++ {
                    if (!(-2.25 >= float64((float32(1.5) - (1e10 / v1))))) {
                        if v6 {
                            fmt.Println(v25, fi33, v2)
                        }
                    }
                    v34 := GenS3{F0: 1, F1: 4}
                    v35 := GenS2{F0: 4294967295, F1: 42, F2: v6}
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 bool = true
                    _ = fi36
                    _ = v37
                }
                _ = v32
            }
        default:
            var v38 int = int(v17)
            _ = v38
        }
        v39 := GenS4{F0: v17, F1: 1, F2: v6}
        _ = v39
    }
    fmt.Println(v17, v2, v25)
    v40 := []int{2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
