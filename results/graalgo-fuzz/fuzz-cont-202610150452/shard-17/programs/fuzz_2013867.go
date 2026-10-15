package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 int64, p1 int64) int8 {
    fmt.Println(p1)
    v4 := []int{}
    _ = p0
    _ = p1
    _ = v4
    return -8
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 2
    var v3 float32 = float32(0.0)
    if ((!(true || (false || true))) && true) {
        if true {
            if ((true && true) && (false && false)) {
                if false {
                    v1 -= 1
                }
            } else {
                v2 = (v2 & 65535)
                for fi5 := 0; fi5 < 14; fi5++ {
                    fmt.Println(fi5, v3)
                    var v6 int16 = 32767
                    _ = fi5
                    _ = v6
                }
            }
            v1 = genFn0(583, 9223372036854775807)
        } else {
            var v7 float32 = (0.0 / 0.1)
            v3 -= v3
            _ = v7
        }
    }
    var v8 uint8 = 127
    switch (((((("foo bar" + "a") + "foo bar") + "\n\t") + ("foo bar" + "")) + ("" + "")) + "hello") {
    case "foo bar":
        fmt.Println(v1)
    default:
        var v9 int32 = -2147483648
        _ = v9
    }
    fmt.Println(v3)
    if false {
        switch "\n\t" {
        case "foo bar":
            v8 = v8
        }
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v2, v3, fi10)
                fmt.Println(v2)
                _ = fi11
            }
            var v12 int8 = (^(v1))
            _ = fi10
            _ = v12
        }
        var v13 uint16 = 65535
        v14 := GenS1{F0: "", F1: 2}
        _ = v13
        _ = v14
    }
    v15 := []int{0, 776, 612, -747}
    fmt.Println(v1, v8)
    v3 += float32(float64(((v3 * float32(1e10)) * v3)))
    fmt.Println(v2, v1)
    var v16 uint8 = v8
    for fi17 := 0; fi17 < 13; fi17++ {
        switch "hello" {
        case "foo bar":
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint32 = v2
                switch fi17 {
                case -902:
                    fmt.Println(v2, fi17, fi17)
                    var v20 int = (fi17 >> ((4294967295) & 63))
                    var v21 float64 = 1e10
                    _ = v20
                    _ = v21
                }
                v3 = v3
                _ = fi18
                _ = v19
            }
        case "":
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        default:
            v22 := GenS2{F0: 0, F1: 1}
            _ = v22
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            v15 = append(v15, -795)
            _ = len(v15)
            var v24 int64 = -9223372036854775808
            if len(v15) > 0 {
                v15[0] = 1
            }
            v15 = append(v15, (fi23 - fi17))
            _ = fi23
            _ = v24
        }
        _ = fi17
    }
    v1 = (^(v1))
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int8 = v1
        switch "foo bar" {
        case "":
            v16 -= v16
        }
        v27 := GenS2{F0: -84, F1: -683}
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS3{F0: -2147483648, F1: true}
        _ = fi28
        _ = v29
    }
    v15 = append(v15, -704)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
