package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0(p0 int16) int16 {
    p0 = p0
    v4 := GenS1{F0: false}
    _ = p0
    _ = v4
    return (p0 + p0)
}
func genFn1() uint16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS2{F0: 255}
            v8 := GenS1{F0: false}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 bool = (false && false)
        _ = fi5
        _ = v9
    }
    if (false && true) {
        switch "" {
        case "hello":
            if false {
                var v10 string = "hello"
                v11 := GenS1{F0: false}
                _ = v10
                _ = v11
            }
        case "":
            v12 := []int{-817}
            _ = v12
        case "\n\t":
            if true {
                var v13 bool = (((32767 >= 0) || (-1 < 1000)) || (!true))
                var v14 int8 = -1
                _ = v13
                _ = v14
            }
        }
    }
    return 42
}

func main() {
    v1 := []int{-274, 32768, -9223372036854775808}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v15 int32 = 1
    if (false || true) {
        switch -983 {
        case -454:
            v15 = v15
        case 128:
            v16 := v3[4294967295]
            _ = v16
        }
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    switch "foo bar" {
    case "\n\t":
        switch 4294967295 {
        case -2147483648:
            v15 -= v15
        case -9223372036854775808:
            var v17 int8 = 100
            _ = v17
        }
    case "a":
        v18 := GenS1{F0: false}
        _ = v18
    default:
        var v19 uint8 = 184
        _ = v19
    }
    switch "" {
    case "\n\t":
        if (!((127 > 0) || true)) {
            if ((0 != 32767) || (!("foo bar" < (("\n\t" + "hello") + ("foo bar" + "\n\t"))))) {
                var v20 float32 = 1.5
                v15 = v15
                switch 1 {
                case 32767:
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 string = "\n\t"
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v20
            }
            for fi23 := 0; fi23 < 10; fi23++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v24 int8 = int8(v15)
                _ = fi23
                _ = v24
            }
        } else {
            v25 := GenS1{F0: false}
            fmt.Println(v15)
            _ = v25
        }
    }
    v26 := GenS3{F0: 0}
    v15 += (v15 + v15)
    if ((v15 - v15) < v15) {
        v15 = (v15 ^ (v15 + v15))
        fmt.Println(v15)
        v15 = v15
    }
    delete(v3, 42)
    var v27 int16 = -32768
    if true {
        switch 42 {
        case 888:
            if true {
                v27 = v27
            }
        case -453:
            switch "hello" {
            case "foo bar":
                var v28 string = ("hello" + "\n\t")
                _ = v28
            case "hello":
                v27 = (^(v27))
            default:
                v15 = (v15 * (^(v15)))
            }
        case 4294967295:
            v15 = -1
        default:
            if true {
                switch 32768 {
                case 32767:
                    v29 := GenS4{F0: 127}
                    _ = v29
                case 42:
                    v15 = v15
                case -2147483648:
                    var v30 uint8 = 1
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 uint8 = 164
                    v32 -= v32
                    _ = fi31
                    _ = v32
                }
                var v33 bool = false
                _ = v33
            }
        }
    }
    if true {
        switch "hello" {
        case "hello":
            v27 = (v27 | 1000)
        default:
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := GenS5{F0: (v27 / (((^(v27))) | 1)), F1: 168, F2: (v15 / (((v15 ^ (v15 * int32(v27)))) | 1))}
                switch "a" {
                case "a":
                    v36 := v3[42]
                    _ = v36
                case "foo bar":
                    v37 := GenS6{F0: 133, F1: 0, F2: 42}
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
        }
        fmt.Println(v27, v15)
    }
    v38 := GenS6{F0: 195, F1: 2, F2: 1}
    v39 := GenS4{F0: 127}
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS3{F0: 0}
        var v42 uint16 = 32767
        fmt.Println(v15, v27, v15)
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v39
    _ = v43
}
