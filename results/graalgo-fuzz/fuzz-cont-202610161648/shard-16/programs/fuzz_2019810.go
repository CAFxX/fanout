package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0() uint32 {
    var v5 uint64 = 4294967295
    var v6 int = 127
    var v7 int = 4294967295
    var v8 float32 = 3.14159
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint16) int16 {
    v9 := []int{}
    p0 -= (^(p0))
    p0 = (p0 >> ((167) & 63))
    _ = p0
    _ = v9
    return 2
}

func main() {
    var v1 float64 = ((3.14159 - (1e10 - float64(3.14159))) - (0.0 * -2.25))
    var v2 uint = 0
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v1 = (1.5 / 3.14159)
    var v10 string = "a"
    v10 = v10
    switch "a" {
    case "a":
        v10 = v10
    case "hello":
        v10 = ("foo bar" + "")
    default:
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                if (-9223372036854775808 < 9223372036854775807) {
                    var v13 uint8 = 151
                    var v14 string = "foo bar"
                    v15 := []int{}
                    var v16 int32 = 2147483647
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    v1 += v1
                    fmt.Println(fi11)
                    if false {
                        v1 -= (v1 / (((float64((((float32(v1) * (0.1 - 3.14159)) - 0.0) * float32(1e10))) - (1.5 / 1.5)) + float64(0.0)) / ((v1 * (-2.25 - ((float64(0.1) / 0.1) + v1))) - float64(3.14159))))
                        v17 := GenS1{F0: 504, F1: 519, F2: 873}
                        var v18 string = "hello"
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint64 = 1
                        _ = v19
                    }
                    v2 += v2
                    var v20 int32 = 42
                    fmt.Println(v10, fi12, fi11)
                    _ = v20
                }
                if false {
                    v10 = v10
                } else {
                    var v21 int16 = genFn1(535)
                    v21 += v21
                    _ = v21
                }
                switch fi12 {
                case -307:
                    if (true || true) {
                        fmt.Println(v1)
                    }
                case -4:
                    var v22 uint16 = 2
                    _ = v22
                case -75:
                    var v23 float32 = float32((v1 * v1))
                    _ = v23
                }
                _ = fi12
            }
            var v24 int32 = 1
            if true {
                v25 := GenS2{F0: 2}
                fmt.Println(v10, fi11, v10)
                _ = v25
            }
            _ = fi11
            _ = v24
        }
    }
    v26 := []int{}
    var v27 float64 = 1.5
    fmt.Println(v27, v1)
    var v28 int64 = 328
    var v29 uint = v2
    var v30 int = int(v2)
    for fi31 := 0; fi31 < 7; fi31++ {
        switch v10 {
        case "hello":
            v32 := GenS1{F0: 2, F1: -165, F2: fi31}
            _ = v32
        case "\n\t":
            for fi33 := 0; fi33 < 16; fi33++ {
                if (1 > 0) {
                    v29 -= 941
                    delete(v3, (((v10 + "a") + (v10 + v10)) + ((v10 + ((("a" + (v10 + "hello")) + "a") + "hello")) + "\n\t")))
                }
                if true {
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                }
                _ = len(v26)
                _ = fi33
            }
        case "":
            v30 = v30
        }
        _ = fi31
    }
    switch (v30 ^ (v30 * (v30 >> (((v2 + (^(v29)))) & 63)))) {
    case 32767:
        if (103 < 32767) {
            v34 := GenS1{F0: 2, F1: -1, F2: v30}
            _ = v34
        } else {
            var v35 float64 = 1.5
            for fi36 := 0; fi36 < 17; fi36++ {
                v37 := GenS2{F0: 186}
                _ = fi36
                _ = v37
            }
            _ = v35
        }
    case 2147483647:
        switch 42 {
        case -1:
            var v38 uint16 = 2
            _ = v38
        case -9223372036854775808:
            v39 := GenS2{F0: 255}
            _ = v39
        default:
            v40 := GenS1{F0: 1, F1: 42, F2: v30}
            _ = v40
        }
    default:
        if false {
            v29 = v29
        } else {
            v1 = 1.5
        }
    }
    v41 := GenS3{F0: 2, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
}
