package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
func genFn0(p0 int32) string {
    p0 += p0
    v5 := []int{2147483647, 819, 1, -1}
    _ = p0
    _ = v5
    return ("hello" + ((("hello" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t")) + "foo bar"))
}

func main() {
    var v1 bool = (!false)
    var v2 int64 = -135
    v3 := []uint32{}
    v4 := []string{("\n\t" + ""), ""}
    v1 = v1
    if ((0.1 + 0.1) != ((1e10 / float64(float32((-2.25 / (1.5 * float64(((((-2.25 * (float32(0.0) + 3.14159)) - 1e10) / 1.0) / 1.5))))))) - ((1e10 * (1e10 - (3.14159 + ((((((3.14159 / 0.1) / 3.14159) * ((1e10 - float64(0.1)) / ((1.5 - 1e10) * 1.5))) + -2.25) * 3.14159) / (3.14159 - (float64((0.0 + 0.0)) + ((-2.25 * 0.0) / float64(float32((1e10 / (0.0 + 3.14159))))))))))) / (-2.25 / 3.14159)))) {
        fmt.Println(v1, v1)
        v2 = -1
    } else {
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v6 = (v6 % ((v6) | 1))
                var v9 int = 128
                v6 -= 4294967295
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            v2 += v2
            fmt.Println(v1, v1)
            _ = fi10
        }
        _ = v6
    }
    var v11 string = ""
    switch 2147483647 {
    case -538:
        var v12 int16 = -32768
        _ = v12
    default:
        var v13 int32 = 42
        _ = v13
    }
    v14 := GenS1{F0: 32767, F1: 634}
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS1{F0: -261, F1: 9223372036854775807}
        _ = fi15
        _ = v16
    }
    var v17 int32 = 1
    fmt.Println(v17, v11)
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint = 65535
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v20
        }
        fmt.Println(v11, fi18, v17)
        _ = fi18
    }
    if true {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        if ((-(v2)) > 0) {
            for fi21 := 0; fi21 < 12; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 uint32 = 1
                    _ = fi22
                    _ = v23
                }
                var v24 float32 = (3.14159 / -2.25)
                _ = fi21
                _ = v24
            }
        } else {
            v25 := GenS2{F0: v1}
            _ = v25
        }
    } else {
        v1 = (582 != 422)
        var v26 int16 = 265
        _ = v26
    }
    v1 = false
    if (0.0 >= -2.25) {
        var v27 float64 = -2.25
        v28 := GenS2{F0: false}
        for fi29 := 0; fi29 < 17; fi29++ {
            _ = len(v4)
            if len(v4) > 0 {
                v4[0] = v11
            }
            _ = fi29
        }
        _ = v27
        _ = v28
    } else {
        fmt.Println(v11, v11)
        v30 := GenS2{F0: (255 <= 127)}
        _ = v30
    }
    var v31 string = v11
    if (!true) {
        if ((555 > 0) && (!true)) {
            v32 := GenS3{F0: 1, F1: 32767, F2: 538}
            _ = v32
        }
    }
    var v33 int32 = 2147483647
    v34 := GenS1{F0: -99, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
}
