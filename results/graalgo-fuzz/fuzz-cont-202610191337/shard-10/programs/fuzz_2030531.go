package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 int }
func genFn0() int16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "hello" {
        case "\n\t":
            var v7 int64 = -9223372036854775808
            _ = v7
        default:
            var v8 uint64 = 1
            _ = v8
        }
        _ = fi6
    }
    var v9 uint32 = 302
    switch 0 {
    case -9223372036854775808:
        if ((false && (!(518 != 9223372036854775807))) || (true && (1e10 >= 3.14159))) {
            var v10 int8 = 42
            _ = v10
        } else {
            if false {
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS1{F0: "hello"}
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 17; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        fmt.Println(v9)
                        var v15 int32 = 0
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := GenS1{F0: ("a" + "a")}
                _ = fi16
                _ = v17
            }
        }
    default:
        var v18 uint64 = 353
        _ = v18
    }
    _ = v9
    return 1
}
func genFn1(p0 int) int64 {
    p0 = p0
    for fi19 := 0; fi19 < 11; fi19++ {
        p0 += (^(fi19))
        if false {
            switch p0 {
            case 32768:
                p0 = (fi19 >> ((972) & 63))
            case 900:
                p0 += fi19
            }
            var v20 bool = (!(!true))
            _ = v20
        } else {
            var v21 float64 = (0.0 + 0.1)
            var v22 float32 = 0.0
            switch "" {
            case "":
                v23 := GenS1{F0: (((("a" + "a") + "") + "a") + ((("a" + "foo bar") + "") + ("" + "a")))}
                _ = v23
            case "a":
                if (396 >= 221) {
                    v24 := GenS1{F0: "a"}
                    _ = v24
                } else {
                    v25 := GenS1{F0: "a"}
                    p0 -= 4294967295
                    _ = v25
                }
            case "\n\t":
                var v26 uint = 42
                _ = v26
            }
            _ = v21
            _ = v22
        }
        v27 := []int{(fi19 * fi19), fi19, fi19}
        _ = fi19
        _ = v27
    }
    var v28 uint16 = 2
    var v29 float32 = 1.5
    _ = p0
    _ = v28
    _ = v29
    return -371
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    var v4 string = (("\n\t" + "a") + ("a" + ("hello" + "a")))
    v5 := make(map[string]uint32)
    v30 := GenS2{F0: 740, F1: 2147483647, F2: v2}
    switch "foo bar" {
    case "a":
        var v31 uint16 = 446
        _ = v31
    case "foo bar":
        v4 = v4
    default:
        v32 := []int{-9223372036854775808, -9223372036854775808, -1, 1}
        _ = v32
    }
    v4 = ""
    switch "hello" {
    case "hello":
        fmt.Println(v2, v4)
    case "":
        for fi33 := 0; fi33 < 17; fi33++ {
            v2 += (v2 & genFn0())
            _ = fi33
        }
    case "a":
        v34 := []int{-1, 42}
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        v36 := GenS3{F0: v2, F1: 42, F2: (fi35 & (fi35 + fi35))}
        fmt.Println(v4, fi35, v2)
        _ = fi35
        _ = v36
    }
    v37 := GenS1{F0: (v4 + v4)}
    var v38 float64 = float64(float32((1.5 / 0.1)))
    var v39 int8 = 0
    v40 := []int{}
    v41 := GenS1{F0: "a"}
    var v42 int64 = 650
    switch int(v42) {
    case 233:
        for fi43 := 0; fi43 < 1; fi43++ {
            fmt.Println(v4)
            fmt.Println(fi43, v4)
            _ = fi43
        }
    case 4294967295:
        v44 := v1[v4]
        _ = v44
    case 291:
        var v45 int64 = (-(v42))
        _ = v45
    default:
        v46 := GenS3{F0: (v2 * v2), F1: 127, F2: -1}
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
