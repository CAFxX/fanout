package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() uint16 {
    v7 := GenS1{F0: -9223372036854775808}
    v8 := []int{0}
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int64, p1 uint8) int32 {
    p1 -= p1
    if (!(true || false)) {
        if (!(true || (42 > 18446744073709551615))) {
            var v9 bool = true
            switch -464 {
            case 127:
                p0 += p0
            default:
                var v10 float32 = ((((1e10 / 1.0) / 1.5) * (1e10 * float32(1.5))) - (1e10 + (1e10 * 0.1)))
                _ = v10
            }
            _ = v9
        } else {
            var v11 uint16 = 255
            var v12 string = "foo bar"
            p1 += 72
            _ = v11
            _ = v12
        }
        fmt.Println(p0, p0)
        fmt.Println(p1, p0)
    } else {
        var v13 int32 = 42
        v14 := GenS1{F0: -2147483648}
        fmt.Println(v13, p1, p0)
        _ = v13
        _ = v14
    }
    var v15 float32 = 0.1
    _ = p0
    _ = p1
    _ = v15
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    var v3 int32 = -1
    var v4 float64 = 3.14159
    var v5 float64 = -2.25
    var v6 int16 = 32767
    var v16 uint = 381
    v17 := GenS1{F0: 32768}
    fmt.Println(v3, v5, v3)
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            if true {
                v20 := GenS2{F0: (^(v16)), F1: 4294967295, F2: (v6 ^ (v6 ^ v6))}
                switch (("a" + "foo bar") + "") {
                case "\n\t":
                    v3 = (v3 >> ((1) & 63))
                default:
                    fmt.Println(v16, v16)
                }
                _ = v20
            }
            v4 += (3.14159 - (((v5 - v4) * (v4 * ((v5 - (float64(float32((((v4 - 0.0) / v4) - v4))) - ((((v4 * 1e10) * (v4 - (((v4 / v4) - v4) + 0.1))) - ((v4 / 1e10) - 3.14159)) - 0.1))) / (v5 / ((v5 / (-2.25 / v5)) + (1.5 / (0.1 - (v4 - v5)))))))) + (v5 / 1.0)))
            v5 = ((v4 - ((1e10 / v5) * (v5 / (v4 / (1.5 - float64(3.14159)))))) - float64(0.1))
            _ = fi19
        }
        _ = fi18
    }
    fmt.Println(v3)
    _ = len(v2)
    var v21 int64 = -937
    for fi22 := 0; fi22 < 1; fi22++ {
        v3 = v3
        var v23 bool = (v21 >= v21)
        _ = fi22
        _ = v23
    }
    if true {
        fmt.Println(v6, v4)
        fmt.Println(v4)
    }
    v6 = v6
    var v24 int8 = int8(v21)
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS3{F0: 967}
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
