package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint; F2 uint }
func genFn0(p0 uint) int32 {
    fmt.Println(p0)
    p0 -= (p0 & 757)
    fmt.Println(p0)
    _ = p0
    return 886
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int64 = 570
    if (65535 > 42) {
        var v4 float32 = (0.1 / -2.25)
        _ = v4
    } else {
        v5 := GenS1{F0: 32768}
        v6 := GenS1{F0: 4294967295}
        _ = v5
        _ = v6
    }
    var v7 float64 = ((((float64(0.0) / 1.0) + float64(0.0)) - -2.25) - (1e10 + 0.1))
    fmt.Println(v7, v3)
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float64 = ((v8 * 1e10) * v8)
        fmt.Println(fi9, fi9)
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{fi9}
            if (!(true && true)) {
                var v13 int32 = -36
                v14 := GenS1{F0: 2147483647}
                if true {
                    v12 = append(v12, -219)
                    v15 := GenS1{F0: (fi9 - fi9)}
                    var v16 int8 = 0
                    _ = v15
                    _ = v16
                }
                _ = v13
                _ = v14
            } else {
                var v17 float32 = 1e10
                _ = v17
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v7, v3)
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int = (fi18 << ((0) & 63))
        _ = fi18
        _ = v19
    }
    delete(v1, (("foo bar" + "foo bar") + ("hello" + "foo bar")))
    for fi20 := 0; fi20 < 16; fi20++ {
        v8 = (((v8 * float64(-2.25)) - v8) / v7)
        switch "hello" {
        case "":
            v21 := []int{fi20, fi20, fi20, (fi20 & (fi20 % ((32768) | 1)))}
            _ = v21
        case "hello":
            var v22 int = fi20
            _ = v22
        default:
            var v23 bool = false
            _ = v23
        }
        switch ("a" + "foo bar") {
        case "hello":
            var v24 bool = (true || (float64((0.1 * 1.5)) != (v7 / (float64(float32(v8)) * -2.25))))
            _ = v24
        default:
            fmt.Println(v7, v7)
        }
        _ = fi20
    }
    v3 -= 0
    v3 = v3
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS2{F0: false, F1: 42, F2: 617}
        v27 := GenS3{F0: 255}
        v28 := []int{127}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 bool = ("a" == "")
    v3 += -451
    if (v29 || (0.0 == 0.1)) {
        fmt.Println(v3, v7, v3)
        if v29 {
            switch ("" + ("a" + "\n\t")) {
            case "":
                v1["hello"] = "hello"
            default:
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := GenS3{F0: 0}
                    _ = fi30
                    _ = v31
                }
            }
            v32 := GenS1{F0: 9223372036854775807}
            v33 := GenS2{F0: (!(100 != -1)), F1: 503, F2: -951}
            _ = v32
            _ = v33
        }
    }
    v34 := GenS4{F0: -363, F1: 0, F2: 757}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v7
    _ = v8
}
