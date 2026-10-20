package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint16, p1 bool) uint16 {
    var v7 uint16 = (^(p0))
    fmt.Println(p0, v7, p1)
    var v8 float32 = ((((float32((3.14159 / float64(0.1))) + 1.5) / float32(((((3.14159 / 1.5) + 3.14159) + (1.5 - 3.14159)) / ((float64((1e10 / ((3.14159 * 3.14159) * 1.5))) + 0.0) + 0.1)))) / (0.1 + 1.5)) / 1e10)
    for fi9 := 0; fi9 < 15; fi9++ {
        if p1 {
            switch "a" {
            case "a":
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := []int{}
                    var v12 uint16 = (v7 | v7)
                    if (v8 <= -2.25) {
                        fmt.Println(p1, v7, v8)
                        var v13 uint32 = 65535
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "":
                var v14 uint32 = 65535
                _ = v14
            default:
                var v15 int64 = 896
                _ = v15
            }
            v16 := []int{}
            if p1 {
                v17 := GenS1{F0: 42}
                var v18 string = "foo bar"
                fmt.Println(v18)
                _ = v17
                _ = v18
            }
            _ = v16
        } else {
            var v19 string = ("foo bar" + ("a" + (("a" + "") + "foo bar")))
            v7 -= (p0 * (v7 / ((v7) | 1)))
            _ = v19
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 int8, p1 int64, p2 int32) uint8 {
    if (42 == 22) {
        switch 42 {
        case 4294967295:
            v20 := GenS1{F0: 0}
            _ = v20
        case 1:
            var v21 int8 = (p0 / ((p0) | 1))
            _ = v21
        case -581:
            p2 = p2
        default:
            var v22 int = 9223372036854775807
            _ = v22
        }
    } else {
        v23 := GenS1{F0: 42}
        p2 -= (-(p2))
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 float64 = 0.0
        v26 := []int{}
        var v27 uint = 848
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.0
    var v3 int = 128
    var v4 string = ""
    v5 := make(map[uint32]string)
    var v6 int = v3
    if true {
        _ = len(v1)
    } else {
        var v28 uint = 255
        fmt.Println(v3)
        v2 += v2
        _ = v28
    }
    v4 = (v4 + v4)
    v29 := []int{(^(v3))}
    v2 -= (((v2 - v2) + v2) / float32(((0.1 / 1.5) - 0.1)))
    var v30 uint = 42
    v4 = v4
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS1{F0: 0}
        var v33 int8 = 108
        switch -2147483648 {
        case 0:
            switch (v6 + (fi31 >> ((4294967295) & 63))) {
            case 4294967295:
                v34 := GenS2{F0: 0, F1: 4294967295}
                _ = v34
            case 42:
                var v35 float32 = (-2.25 - 1e10)
                _ = v35
            case 32768:
                fmt.Println(v4, v3, fi31)
            }
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    delete(v5, 42)
    for fi36 := 0; fi36 < 15; fi36++ {
        _ = len(v29)
        var v37 bool = true
        if (1 >= 100) {
            var v38 int16 = 0
            if ((-1 != 9223372036854775807) || true) {
                v29 = append(v29, (^(fi36)))
            } else {
                if (42 > 0) {
                    v5[705] = v4
                    v29 = append(v29, fi36)
                    fmt.Println(v3, v4, v30)
                }
            }
            var v39 string = (v4 + "foo bar")
            _ = v38
            _ = v39
        } else {
            var v40 int32 = -1
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v2 = (((float32((1e10 * 0.0)) - 3.14159) / -2.25) * 1.5)
        _ = fi41
    }
    v42 := GenS2{F0: 645, F1: 829}
    v43 := GenS1{F0: 668}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
