package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int; F1 uint32; F2 string }
func genFn0(p0 int8, p1 int16, p2 int) uint64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = (p0 >> ((2) & 63))
            var v9 uint32 = 1
            fmt.Println(v8, p2, p1)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    p1 -= (p1 % ((p1) | 1))
    var v10 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint8, p1 int32, p2 int) int {
    switch (("" + "\n\t") + ("hello" + "")) {
    case "hello":
        if true {
            v11 := []int{p2, -1}
            v12 := GenS1{F0: 4294967295}
            _ = v11
            _ = v12
        } else {
            if true {
                v13 := []int{p2, (p2 ^ (p2 % ((p2) | 1))), p2, 335}
                var v14 uint = 1
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 int16 = -698
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            } else {
                fmt.Println(p0)
                p0 = p0
            }
        }
    case "a":
        p1 = p1
    }
    fmt.Println(p2, p1)
    var v17 bool = true
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS2{F0: 255}
        for fi20 := 0; fi20 < 18; fi20++ {
            if false {
                v17 = v17
            }
            switch 9223372036854775807 {
            case 32768:
                v21 := []int{p2, p2, -161}
                _ = v21
            }
            p1 += (p1 | (^(p1)))
            _ = fi20
        }
        v22 := GenS3{F0: 658, F1: 2147483647}
        _ = fi18
        _ = v19
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p2
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1000
    var v3 float64 = 0.0
    v4 := make(map[int]bool)
    v5 := []uint32{2147483648, 0}
    v4[-190] = false
    fmt.Println(v2, v2)
    var v23 uint64 = 1
    fmt.Println(v2, v2, v2)
    switch -196 {
    case 32768:
        v3 += 1.5
    }
    if (true || true) {
        var v24 int32 = -2147483648
        v25 := GenS4{F0: 0, F1: (v23 % (((v23 ^ v23)) | 1))}
        v2 = int16(v23)
        _ = v24
        _ = v25
    }
    fmt.Println(v2)
    for fi26 := 0; fi26 < 15; fi26++ {
        switch ("\n\t" + "hello") {
        case "hello":
            if (82 > 42) {
                v3 += (((((0.1 + v3) * ((1.5 * -2.25) + v3)) * -2.25) / v3) + (0.0 * v3))
                var v27 int = fi26
                _ = v27
            }
        case "\n\t":
            _ = len(v5)
        case "a":
            v28 := GenS5{F0: -9223372036854775808, F1: 4294967295, F2: "a"}
            _ = v28
        }
        var v29 string = ""
        _ = fi26
        _ = v29
    }
    v3 -= v3
    if true {
        var v30 int8 = 70
        fmt.Println(v2)
        _ = v30
    } else {
        var v31 float64 = 1e10
        _ = len(v5)
        _ = v31
    }
    var v32 float32 = 1.5
    if (1e10 != ((v3 / (v3 / v3)) - float64(v32))) {
        delete(v1, "hello")
        v32 -= float32((((1e10 + ((v3 / v3) - v3)) + v3) - 0.0))
    } else {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 uint64 = 18446744073709551615
            _ = fi33
            _ = v34
        }
        if true {
            var v35 uint = 255
            var v36 int64 = -1
            _ = v35
            _ = v36
        } else {
            for fi37 := 0; fi37 < 2; fi37++ {
                var v38 float64 = 0.1
                _ = fi37
                _ = v38
            }
        }
        if false {
            v39 := v1["a"]
            fmt.Println(v3, v2)
            _ = v39
        } else {
            for fi40 := 0; fi40 < 7; fi40++ {
                var v41 uint32 = 2
                _ = fi40
                _ = v41
            }
        }
    }
    var v42 float64 = (1.5 / -2.25)
    switch "hello" {
    case "hello":
        v2 = (^(v2))
    case "a":
        var v43 uint = 0
        _ = v43
    case "foo bar":
        for fi44 := 0; fi44 < 15; fi44++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi44
        }
    }
    switch -20 {
    case 440:
        v32 = 3.14159
    default:
        v32 = (v32 / -2.25)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v5
}
