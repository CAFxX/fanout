package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint32, p1 uint) int64 {
    v7 := GenS1{F0: 128, F1: 9223372036854775807}
    var v8 uint64 = 18446744073709551615
    if (false && (-2147483648 >= 1)) {
        if true {
            v9 := []int{}
            if len(v9) > 0 {
                v9[0] = 42
            }
            v10 := GenS1{F0: 204, F1: 0}
            _ = v10
            _ = v9
        }
    }
    v11 := []int{-1, 32767, 0}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            p0 -= p0
            v14 := GenS1{F0: 127, F1: 0}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return -374
}
func genFn1(p0 int8) int64 {
    if (4294967295 >= -813) {
        var v15 uint = 255
        switch "hello" {
        case "foo bar":
            v16 := []int{4294967295, -129}
            _ = v16
        case "hello":
            fmt.Println(v15, p0)
        }
        _ = v15
    } else {
        v17 := GenS2{F0: 2}
        var v18 uint8 = 148
        _ = v17
        _ = v18
    }
    fmt.Println(p0)
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            fmt.Println(p0, fi20)
            _ = fi20
        }
        v21 := []int{fi19, fi19, fi19}
        _ = fi19
        _ = v21
    }
    switch ("hello" + ("" + "foo bar")) {
    case "foo bar":
        p0 = (p0 - 1)
    case "a":
        fmt.Println(p0)
    default:
        p0 = (p0 + 127)
    }
    _ = p0
    return -1
}

func main() {
    var v1 int16 = 489
    var v2 float64 = 1.5
    var v3 int64 = 909
    var v4 uint8 = 87
    var v5 int32 = int32(v1)
    var v6 uint8 = (v4 * (v4 + (v4 ^ 162)))
    for fi22 := 0; fi22 < 15; fi22++ {
        v2 += 0.0
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 int64 = (v3 % ((v3) | 1))
            switch fi23 {
            case 4294967295:
                v5 = (v5 << ((65535) & 63))
            case -1:
                switch 1 {
                case 128:
                    v6 += 127
                case -1:
                    v24 = 1
                case 4294967295:
                    v24 = genFn1(1)
                }
            case 930:
                var v25 string = ("hello" + "hello")
                _ = v25
            default:
                var v26 string = ("a" + "foo bar")
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v27 := []int{32768, 1, 9223372036854775807, 660}
    var v28 float64 = float64(1e10)
    var v29 int = 128
    v1 = (v1 | (v1 % ((-1) | 1)))
    var v30 int64 = (v3 >> ((280) & 63))
    var v31 int = v29
    v32 := GenS3{F0: -1, F1: v6}
    fmt.Println(v6, v28, v31)
    v27 = append(v27, 4294967295)
    var v33 uint = 1
    for fi34 := 0; fi34 < 11; fi34++ {
        switch "hello" {
        case "a":
            v29 = fi34
        case "hello":
            fmt.Println(fi34, v30)
        default:
            if true {
                var v35 uint = (v33 >> ((849) & 63))
                var v36 string = ("\n\t" + "a")
                _ = v35
                _ = v36
            }
        }
        _ = fi34
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v30 = genFn1(25)
        var v38 uint8 = 127
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
