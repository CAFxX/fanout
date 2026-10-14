package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int32) int32 {
    var v7 string = ("hello" + "hello")
    v7 = v7
    var v8 int8 = 0
    _ = p0
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint16) int16 {
    var v9 int32 = 0
    var v10 int32 = (v9 % (((v9 << ((42) & 63))) | 1))
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = v10
    _ = v9
    return -32768
}

func main() {
    var v1 uint32 = 1
    var v2 float64 = float64(3.14159)
    var v3 float64 = ((v2 * v2) / float64(float32(v2)))
    var v4 int64 = -617
    v5 := make(map[int]string)
    var v6 uint16 = 72
    v6 = 2
    v6 = (v6 << ((255) & 63))
    for fi11 := 0; fi11 < 1; fi11++ {
        if (65535 < 255) {
            if (234 != 128) {
                var v12 string = ("foo bar" + "\n\t")
                _ = v12
            }
            var v13 uint64 = 588
            var v14 float64 = v2
            _ = v13
            _ = v14
        } else {
            if true {
                fmt.Println(v4)
                for fi15 := 0; fi15 < 12; fi15++ {
                    if (false && (("" < "") || true)) {
                        v16 := GenS1{F0: ("" + ("a" + "a"))}
                        _ = v16
                    }
                    _ = fi15
                }
            } else {
                v2 -= (3.14159 - (v2 / 1e10))
                v4 = v4
            }
            v3 += float64(0.0)
            fmt.Println(v6, v3)
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if false {
            var v18 int8 = 0
            for fi19 := 0; fi19 < 7; fi19++ {
                v2 = 0.1
                _ = fi19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := v5[(-(fi17))]
                _ = fi20
                _ = v21
            }
            _ = v18
        }
        _ = fi17
    }
    if ((v4 & v4) > 9223372036854775807) {
        v1 = 2
    }
    fmt.Println(v2)
    v22 := []int{32767}
    fmt.Println(v4, v3)
    if len(v22) > 0 {
        v22[0] = 650
    }
    var v23 int16 = 32767
    switch 841 {
    case -1:
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 string = "a"
            delete(v5, int(v1))
            v25 = (v25 + "hello")
            _ = fi24
            _ = v25
        }
    case 906:
        var v26 int16 = v23
        _ = v26
    default:
        v6 = (v6 / (((v6 & 195)) | 1))
    }
    var v27 int64 = (^(v4))
    v23 += v23
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
