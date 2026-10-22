package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 int) int8 {
    if (p1 < p1) {
        if ((!true) || true) {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint32 = 318
                v9 := []int{}
                switch "hello" {
                case "hello":
                    p1 -= (fi7 << ((0) & 63))
                case "a":
                    switch "a" {
                    case "foo bar":
                        v9 = append(v9, p1)
                    case "a":
                        var v10 float32 = float32((0.1 + 0.0))
                        _ = v10
                    }
                case "":
                    var v11 float64 = (-2.25 / 1e10)
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            if (!false) {
                p1 -= p1
                p1 += (p1 << ((42) & 63))
            }
        }
        v12 := GenS1{F0: "foo bar"}
        _ = v12
    } else {
        v13 := GenS1{F0: ("foo bar" + "foo bar")}
        p1 -= p1
        _ = v13
    }
    p0 = 0
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 65535
    var v3 bool = (-1 >= 231)
    var v4 int = 456
    var v5 string = v1
    var v6 uint = 380
    fmt.Println(v5)
    switch "a" {
    case "foo bar":
        for fi14 := 0; fi14 < 14; fi14++ {
            v5 = ((v5 + "foo bar") + (v5 + v1))
            v15 := []int{fi14, (v4 ^ -2147483648)}
            _ = fi14
            _ = v15
        }
    case "":
        v16 := []int{}
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 3; fi17++ {
            v4 = (fi17 << (((^(v6))) & 63))
            _ = fi17
        }
    default:
        var v18 int = v4
        _ = v18
    }
    v19 := []int{(^(v4)), (v4 & (-(v4))), (-(v4))}
    var v20 int8 = 100
    v19 = append(v19, v4)
    v21 := GenS1{F0: "hello"}
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 int16 = 42
        _ = fi22
        _ = v23
    }
    v4 = (v4 & (v4 / ((-658) | 1)))
    var v24 float64 = float64(((1.5 + 1.5) - 0.1))
    var v25 int32 = 0
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS1{F0: "foo bar"}
        var v28 int16 = 698
        fmt.Println(v28, fi26, v3)
        _ = fi26
        _ = v27
        _ = v28
    }
    if (v3 && (v5 == (v5 + (v1 + (v5 + "hello"))))) {
        var v29 int = (v4 % (((v4 | v4)) | 1))
        _ = v29
    } else {
        fmt.Println(v24, v2)
        if len(v19) > 0 {
            v19[0] = 0
        }
        var v30 uint8 = 0
        _ = v30
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    switch int(v20) {
    case 0:
        var v31 int32 = v25
        _ = v31
    case 32768:
        var v32 bool = (!v3)
        _ = v32
    default:
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
