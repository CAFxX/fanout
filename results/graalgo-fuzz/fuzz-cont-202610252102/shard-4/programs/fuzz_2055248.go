package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 string) int64 {
    if (true || false) {
        var v7 int32 = 2147483647
        if true {
            if true {
                switch p0 {
                case "":
                    if true {
                        v8 := []int{}
                        v8 = append(v8, 1)
                        var v9 float64 = (1.5 + -2.25)
                        _ = v8
                        _ = v9
                    }
                case "\n\t":
                    var v10 uint16 = 236
                    _ = v10
                }
                switch p0 {
                case "foo bar":
                    p0 = "\n\t"
                case "":
                    v7 = v7
                }
            } else {
                switch 0 {
                case 148:
                    v11 := []int{869, 1, -572, -1}
                    _ = v11
                case 2147483647:
                    v7 -= (v7 / (((v7 - v7)) | 1))
                case -832:
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                fmt.Println(v7, v7)
            }
        } else {
            v13 := []int{}
            v7 = -2147483648
            _ = v13
        }
        if (4294967295 != 18446744073709551615) {
            v14 := GenS1{F0: 42}
            var v15 uint32 = 2147483647
            var v16 uint = 795
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v7
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 string = ("" + ("foo bar" + p0))
            v19 := []int{}
            var v20 float64 = 0.1
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        fmt.Println(p0)
        switch -238 {
        case -2147483648:
            v21 := []int{42, -2147483648, -279, 32767}
            _ = v21
        default:
            var v22 uint64 = 0
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        switch p0 {
        case "\n\t":
            p0 = ((p0 + ("" + "a")) + (p0 + ""))
        case "a":
            v24 := GenS1{F0: 4294967295}
            _ = v24
        case "hello":
            p0 = "a"
        default:
            for fi25 := 0; fi25 < 13; fi25++ {
                p0 = p0
                _ = fi25
            }
        }
        p0 = (p0 + (p0 + "hello"))
        _ = fi23
    }
    if false {
        fmt.Println(p0)
        p0 = ((p0 + p0) + p0)
    }
    if false {
        v26 := []int{-1}
        var v27 int64 = -1
        var v28 string = p0
        _ = v26
        _ = v27
        _ = v28
    } else {
        switch "hello" {
        case "a":
            var v29 bool = true
            _ = v29
        }
        p0 = ""
        p0 = p0
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1e10
    var v3 uint64 = 42
    v4 := make(map[uint32]uint32)
    v5 := []string{"foo bar"}
    var v6 bool = false
    v30 := GenS1{F0: 2}
    v2 = 1e10
    v31 := GenS1{F0: 944}
    v2 -= (((3.14159 / (v2 + ((1e10 / 0.1) / (v2 * (v2 / float32(float64(v2))))))) / float32(-2.25)) / -2.25)
    if true {
        var v32 uint16 = 0
        var v33 bool = v6
        var v34 float64 = 0.1
        _ = v32
        _ = v33
        _ = v34
    } else {
        var v35 float32 = 1.5
        _ = v35
    }
    v2 = v2
    var v36 uint32 = 2
    v36 += 262
    var v37 float32 = (v2 + v2)
    v36 = v36
    v38 := GenS1{F0: (v3 * (v3 / ((v3) | 1)))}
    _ = len(v5)
    var v39 int8 = 32
    var v40 string = "hello"
    for fi41 := 0; fi41 < 17; fi41++ {
        v2 -= v37
        _ = fi41
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := GenS1{F0: 82}
            _ = fi43
            _ = v44
        }
        var v45 int16 = 1000
        _ = fi42
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
