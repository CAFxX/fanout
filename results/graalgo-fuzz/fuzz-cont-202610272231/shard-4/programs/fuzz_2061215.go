package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 string, p1 uint64) int8 {
    var v5 float64 = 0.1
    switch p0 {
    case "":
        var v6 int8 = -1
        _ = v6
    case "foo bar":
        if false {
            for fi7 := 0; fi7 < 8; fi7++ {
                if true {
                    var v8 int16 = -32768
                    v5 += (v5 / (3.14159 / v5))
                    fmt.Println(p1, p1)
                    _ = v8
                }
                _ = fi7
            }
            v9 := []int{}
            v5 = (v5 - 0.0)
            _ = v9
        }
    default:
        fmt.Println(v5, p1, v5)
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    return 100
}
func genFn1(p0 int8, p1 string, p2 uint64) int32 {
    switch 9223372036854775807 {
    case 128:
        v10 := []int{-1}
        _ = v10
    }
    switch -1 {
    case -963:
        var v11 uint32 = 239
        _ = v11
    case 41:
        if true {
            v12 := []int{310, -906, 32768, 489}
            v13 := GenS1{F0: 127}
            if false {
                var v14 int8 = p0
                _ = v14
            } else {
                p2 = 2
                p0 -= (p0 >> ((65535) & 63))
            }
            _ = v12
            _ = v13
        } else {
            var v15 int8 = p0
            v16 := GenS1{F0: 127}
            v17 := GenS1{F0: 255}
            _ = v15
            _ = v16
            _ = v17
        }
    case -71:
        var v18 float64 = float64((3.14159 - float32(float64(float32(1.5)))))
        _ = v18
    default:
        var v19 uint32 = 65535
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar", ("hello" + "foo bar"), ("" + "\n\t")}
    var v4 uint32 = 1
    v4 = (^(v4))
    if true {
        if (134 == 1) {
            v4 += (v4 * (v4 ^ 766))
        } else {
            var v20 int = 9223372036854775807
            var v21 uint64 = 1
            _ = len(v3)
            _ = v20
            _ = v21
        }
        switch "hello" {
        case "foo bar":
            if true {
                var v22 int32 = -2147483648
                if true {
                    _ = len(v1)
                    var v23 int64 = 0
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v23
                }
                fmt.Println(v22, v4)
                _ = v22
            } else {
                switch "a" {
                case "":
                    if (84 <= 84) {
                        var v24 bool = false
                        fmt.Println(v4)
                        _ = v24
                    }
                default:
                    var v25 uint32 = (v4 | v4)
                    _ = v25
                }
            }
        case "a":
            v3 = append(v3, "foo bar")
        }
        var v26 float64 = 0.1
        _ = v26
    }
    v27 := GenS1{F0: 0}
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS2{F0: 1}
        v30 := GenS3{F0: 165, F1: "hello"}
        _ = fi28
        _ = v29
        _ = v30
    }
    v4 = (v4 * v4)
    switch "foo bar" {
    case "\n\t":
        v4 = (v4 >> ((255) & 63))
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(fi31, fi31)
        if true {
            var v32 bool = false
            v4 = (^(v4))
            _ = v32
        } else {
            v4 = 2147483647
        }
        _ = fi31
    }
    var v33 int16 = 1000
    v34 := GenS4{F0: 214}
    var v35 uint64 = 42
    var v36 bool = (false && ((v4 <= v4) || (true || true)))
    var v37 float64 = ((-2.25 / 1.0) * float64(1.5))
    v38 := GenS3{F0: 0, F1: "foo bar"}
    v39 := GenS2{F0: (v33 * 211)}
    fmt.Println(v35)
    for fi40 := 0; fi40 < 19; fi40++ {
        if v36 {
            if v36 {
                v35 = v35
                var v41 int32 = 1
                var v42 uint16 = 32767
                _ = v41
                _ = v42
            } else {
                var v43 int64 = int64(v4)
                fmt.Println(v4, v35, v37)
                switch fi40 {
                case 1:
                    v44 := GenS4{F0: 255}
                    _ = v44
                case -100:
                    if (!(v36 || (625 > 1))) {
                        v45 := GenS5{F0: "foo bar", F1: ("\n\t" + "hello")}
                        _ = v45
                    }
                }
                _ = v43
            }
            fmt.Println(v35, v35)
        } else {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            for fi46 := 0; fi46 < 17; fi46++ {
                v37 = (1.5 * 1e10)
                var v47 uint16 = 42
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                _ = fi46
                _ = v47
            }
        }
        _ = fi40
    }
    var v48 uint8 = 13
    var v49 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
}
