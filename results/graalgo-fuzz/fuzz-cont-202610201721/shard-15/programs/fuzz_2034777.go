package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 string; F2 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint8) uint32 {
    var v7 int8 = 1
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch "a" {
            case "":
                var v10 int16 = -32768
                _ = v10
            case "hello":
                var v11 float64 = 1e10
                _ = v11
            default:
                switch ("foo bar" + "a") {
                case "hello":
                    v12 := []int{}
                    _ = v12
                }
            }
            v13 := []int{int(p0), (fi9 >> (((p0 - p0)) & 63)), fi8}
            v14 := GenS1{F0: int(p0)}
            if ((0.0 / float64(-2.25)) != (((0.0 + -2.25) / 1.0) + 1.5)) {
                if len(v13) > 0 {
                    v13[0] = fi8
                }
                var v15 string = (((("a" + ("a" + "\n\t")) + "") + ("a" + "a")) + "a")
                var v16 bool = true
                v17 := GenS1{F0: 9223372036854775807}
                _ = v15
                _ = v16
                _ = v17
            }
            p1 -= (p1 ^ 128)
            v18 := GenS2{F0: p0, F1: (p0 - 255)}
            _ = fi9
            _ = v13
            _ = v14
            _ = v18
        }
        if true {
            fmt.Println(fi8)
            for fi19 := 0; fi19 < 14; fi19++ {
                v7 = v7
                v20 := []int{fi8, int(p0)}
                _ = fi19
                _ = v20
            }
        } else {
            if (((0.1 / 1.0) >= 1e10) || (float64((1.5 - -2.25)) >= (0.1 - float64(float32(float64(0.1)))))) {
                switch (fi8 ^ (fi8 >> (((^(p0))) & 63))) {
                case 311:
                    var v21 string = "\n\t"
                    _ = v21
                case 171:
                    var v22 string = "foo bar"
                    _ = v22
                case 4294967295:
                    fmt.Println(p1, p0)
                default:
                    p0 -= p0
                }
            } else {
                p0 = (^(p0))
                p0 += p0
                var v23 int64 = 42
                _ = v23
            }
            if false {
                v24 := []int{(fi8 ^ (fi8 >> ((p0) & 63)))}
                _ = v24
            }
            v25 := GenS3{F0: (p1 - 47)}
            _ = v25
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 343
}
func genFn1(p0 bool, p1 int8) uint32 {
    var v26 float64 = (3.14159 / 0.1)
    v26 = v26
    switch 9223372036854775807 {
    case 128:
        v27 := []int{128, 42, -304}
        _ = v27
    }
    var v28 uint = 255
    if (128 > 183) {
        if (2 < genFn0(v28, 0)) {
            var v29 string = "hello"
            if (128 <= 128) {
                if ((42 < v28) || true) {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v31 := []int{-801, -9223372036854775808}
                        _ = fi30
                        _ = v31
                    }
                    v28 = (v28 % (((v28 + (^(v28)))) | 1))
                    v32 := GenS4{F0: -9223372036854775808, F1: "foo bar", F2: 0}
                    _ = v32
                } else {
                    v33 := GenS5{F0: 42}
                    _ = v33
                }
            }
            _ = v29
        }
        fmt.Println(v26, v26, p0)
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v28
    return 65535
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]uint32)
    var v3 int = 0
    v4 := make(map[string]uint32)
    v5 := []uint32{909}
    var v6 int = -2147483648
    var v34 uint8 = 128
    var v35 float64 = -2.25
    var v36 int = 9223372036854775807
    switch v6 {
    case 4294967295:
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 bool = true
            v39 := GenS1{F0: (fi37 - v36)}
            v40 := GenS1{F0: 931}
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    default:
        var v41 int64 = 1
        _ = v41
    }
    v42 := GenS3{F0: 127}
    if ((1 > 42) && false) {
        switch (("a" + "\n\t") + ("hello" + ((("hello" + (("hello" + "\n\t") + "")) + "hello") + "foo bar"))) {
        case "hello":
            v43 := GenS4{F0: 0, F1: ("\n\t" + "foo bar"), F2: 126}
            _ = v43
        case "\n\t":
            v36 -= (v6 & 32768)
        }
        var v44 int32 = (-(v1))
        v6 = v6
        _ = v44
    } else {
        var v45 int64 = 42
        for fi46 := 0; fi46 < 3; fi46++ {
            _ = len(v2)
            _ = fi46
        }
        _ = v45
    }
    var v47 float64 = (v35 / 1.0)
    if (255 > 4294967295) {
        var v48 uint16 = 42
        _ = v48
    }
    v47 = 3.14159
    v36 = v36
    v49 := GenS5{F0: 42}
    if true {
        v6 -= 127
        switch v3 {
        case 221:
            var v50 int = v36
            _ = v50
        case 42:
            fmt.Println(v1, v3, v6)
        case -2147483648:
            var v51 string = ("foo bar" + "a")
            _ = v51
        }
        fmt.Println(v1)
    }
    var v52 uint64 = 961
    v53 := GenS2{F0: 0, F1: 767}
    fmt.Println(v36)
    v35 += 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v47
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
