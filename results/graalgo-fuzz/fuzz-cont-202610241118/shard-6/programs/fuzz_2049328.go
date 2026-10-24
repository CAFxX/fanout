package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int32; F2 int8 }
type GenS5 struct { F0 int8; F1 uint }
func genFn0(p0 uint, p1 uint8, p2 uint) uint16 {
    switch 42 {
    case -9223372036854775808:
        p2 = p0
    case -2147483648:
        var v4 uint64 = 339
        _ = v4
    case -767:
        var v5 uint32 = 1
        _ = v5
    }
    v6 := []int{32768, 1}
    var v7 float64 = 0.1
    var v8 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = -751
    var v2 uint32 = 19
    var v3 uint16 = 2
    fmt.Println(v2, v2)
    var v9 uint32 = (v2 / ((v2) | 1))
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int16 = v1
            if (-1 <= 606) {
                fmt.Println(fi11, fi11)
                var v13 int16 = v1
                v2 += v2
                v3 = v3
                _ = v13
            } else {
                v14 := GenS1{F0: "foo bar"}
                v2 -= v2
                _ = v14
            }
            switch "foo bar" {
            case "a":
                v9 = v2
            case "\n\t":
                v2 = 0
            }
            _ = fi11
            _ = v12
        }
        switch "\n\t" {
        case "a":
            v15 := GenS1{F0: "hello"}
            _ = v15
        case "foo bar":
            var v16 float32 = 3.14159
            _ = v16
        case "hello":
            switch -2147483648 {
            case 4294967295:
                switch (fi10 + fi10) {
                case 0:
                    var v17 uint16 = v3
                    _ = v17
                default:
                    if false {
                        v18 := GenS1{F0: ("a" + "\n\t")}
                        _ = v18
                    }
                }
            case 9223372036854775807:
                v1 -= 764
            }
        default:
            if false {
                switch "" {
                case "hello":
                    v19 := []int{fi10, fi10, -1, 127}
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v21 := []int{42, (fi10 / ((fi10) | 1)), (^(fi20)), fi20}
                        fmt.Println(fi20)
                        var v22 int8 = -64
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case "":
                    v23 := GenS1{F0: "a"}
                    _ = v23
                default:
                    v9 -= v2
                }
                v24 := GenS2{F0: -288}
                _ = v24
            } else {
                fmt.Println(v2, v9, v1)
                fmt.Println(v3, fi10)
            }
        }
        if ((!false) && (v2 <= (v9 << ((4294967295) & 63)))) {
            var v25 uint8 = 92
            v26 := GenS3{F0: "hello"}
            v27 := GenS1{F0: ""}
            _ = v25
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 20; fi28++ {
                fmt.Println(v9, v3, v2)
                var v29 int32 = 1
                _ = fi28
                _ = v29
            }
            var v30 uint8 = 102
            _ = v30
        }
        _ = fi10
    }
    v31 := []int{}
    if true {
        v3 += 2
        if (1 == 42) {
            fmt.Println(v9)
            var v32 uint64 = 4294967295
            _ = v32
        }
    } else {
        v33 := GenS4{F0: v3, F1: 1, F2: 100}
        v34 := GenS1{F0: ("" + (("\n\t" + "\n\t") + "foo bar"))}
        _ = v33
        _ = v34
    }
    v31 = append(v31, 552)
    v35 := GenS5{F0: 33, F1: 65535}
    var v36 int = 1
    v3 -= (v3 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v9
}
