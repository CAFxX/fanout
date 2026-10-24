package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int32; F1 int16 }
func genFn0(p0 int, p1 int) int16 {
    var v7 uint16 = 32767
    fmt.Println(p0, v7)
    var v8 string = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint64, p1 int64) uint32 {
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{(fi9 * 32767), (fi9 | 4294967295)}
        _ = fi9
        _ = v10
    }
    if (4294967295 == 4294967295) {
        fmt.Println(p0)
        p1 = -9223372036854775808
        if (354 <= 28) {
            fmt.Println(p1, p0)
            if false {
                var v11 int = int(p0)
                v12 := GenS1{F0: 68, F1: 0, F2: ""}
                _ = v11
                _ = v12
            }
            var v13 float64 = -2.25
            _ = v13
        } else {
            switch -604 {
            case 1:
                if true {
                    switch "a" {
                    case "":
                        v14 := []int{128, 32768}
                        _ = v14
                    case "a":
                        p0 = (p0 << ((uint(p0)) & 63))
                    }
                } else {
                    v15 := []int{}
                    p1 += 42
                    var v16 int64 = p1
                    _ = v15
                    _ = v16
                }
            }
            var v17 uint8 = 0
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 1, F1: 1, F2: ("foo bar" + "a")}
        _ = v18
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    var v3 int32 = -296
    v4 := make(map[string]bool)
    v5 := []int{0, 42, 237, -790}
    v6 := make(map[string]string)
    if true {
        fmt.Println(v3, v2)
        var v19 int8 = -1
        var v20 uint64 = 4294967295
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: "foo bar", F1: 1, F2: 696}
    var v22 float64 = 3.14159
    fmt.Println(v2, v3, v2)
    if (false || true) {
        _ = len(v1)
    }
    v23 := GenS2{F0: ((("\n\t" + "") + ((("a" + ("a" + ((("\n\t" + "foo bar") + ("hello" + "hello")) + "\n\t"))) + ((("" + "hello") + ("" + "foo bar")) + "foo bar")) + "hello")) + "hello"), F1: 42, F2: 1}
    var v24 int64 = 868
    if true {
        v2 -= v2
    } else {
        v3 = 0
    }
    var v25 float32 = (-2.25 - 0.1)
    if false {
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                switch (fi27 | 9223372036854775807) {
                case 9223372036854775807:
                    var v28 int64 = 9223372036854775807
                    _ = v28
                }
                var v29 int16 = 0
                v30 := GenS3{F0: 255}
                _ = fi27
                _ = v29
                _ = v30
            }
            _ = fi26
        }
        if (1 > uint(v2)) {
            var v31 uint64 = 18446744073709551615
            _ = v31
        } else {
            for fi32 := 0; fi32 < 8; fi32++ {
                var v33 uint64 = 0
                for fi34 := 0; fi34 < 20; fi34++ {
                    v35 := GenS4{F0: "a", F1: ("\n\t" + ("hello" + "hello"))}
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 15; fi36++ {
                    fmt.Println(fi36)
                    _ = fi36
                }
                _ = fi32
                _ = v33
            }
            v2 += (v2 | v2)
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            v6[("" + "\n\t")] = "hello"
            switch (-(fi37)) {
            case 593:
                var v38 int64 = -9223372036854775808
                _ = v38
            case 483:
                v39 := GenS1{F0: 128, F1: 4294967295, F2: ("a" + ("a" + ((("hello" + "\n\t") + "\n\t") + "\n\t")))}
                _ = v39
            default:
                v25 = (1.5 / (float32((((v22 / v22) * v22) / 1.0)) + (v25 * -2.25)))
            }
            var v40 int16 = int16(v2)
            _ = fi37
            _ = v40
        }
    } else {
        v2 = (v2 << ((255) & 63))
        v41 := GenS5{F0: 0, F1: 1}
        _ = v41
    }
    var v42 uint8 = 179
    for fi43 := 0; fi43 < 13; fi43++ {
        v6["\n\t"] = ""
        _ = fi43
    }
    v25 = (3.14159 / (-2.25 - 1e10))
    v3 = (v3 ^ v3)
    var v44 int32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
