package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []int{-2147483648, 9223372036854775807}
    v2 := make(map[int]bool)
    v3 := []string{"hello", (("hello" + "foo bar") + ("foo bar" + ("\n\t" + "foo bar")))}
    v4 := []int{-807, -2147483648, 32767, -247}
    v5 := make(map[string]uint32)
    var v6 float64 = 3.14159
    v7 := v2[-190]
    v4 = append(v4, -587)
    var v8 float32 = ((float32(1.5) + 0.1) * 1.5)
    v8 -= (1.5 * (v8 / (1e10 * v8)))
    v6 = 1.5
    v6 -= 0.0
    v9 := GenS1{F0: 42, F1: 31}
    switch 370 {
    case 0:
        for fi10 := 0; fi10 < 11; fi10++ {
            v7 = v7
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 32767, F1: 42}
                _ = fi11
                _ = v12
            }
            v7 = (32767 != 1)
            _ = fi10
        }
    case 1:
        var v13 int16 = -277
        _ = v13
    case 42:
        var v14 int8 = -128
        _ = v14
    default:
        var v15 int16 = 66
        _ = v15
    }
    v6 += (-2.25 - 0.1)
    if v7 {
        v16 := GenS2{F0: 42}
        fmt.Println(v8)
        if (!((641 > 9223372036854775807) || (42 <= -889))) {
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 13; fi18++ {
                    _ = len(v4)
                    _ = fi18
                }
                v7 = (!v7)
                switch fi17 {
                case 37:
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v7 = ((v8 - v8) < (3.14159 * 0.1))
                        _ = fi19
                    }
                case -1:
                    switch 32767 {
                    case 32767:
                        v8 = 3.14159
                    }
                case -924:
                    var v20 uint = 65535
                    _ = v20
                default:
                    var v21 float64 = ((v6 * v6) + (v6 * (v6 / 1.0)))
                    _ = v21
                }
                _ = fi17
            }
        } else {
            var v22 uint8 = 127
            for fi23 := 0; fi23 < 17; fi23++ {
                for fi24 := 0; fi24 < 16; fi24++ {
                    v8 -= 0.0
                    if (!(!v7)) {
                        v6 -= (float64((1e10 - float32(((float64(0.1) / (float64(v8) - v6)) + v6)))) * v6)
                        var v25 int16 = 1
                        _ = v25
                    } else {
                        if false {
                            delete(v2, (fi24 - (fi23 ^ (-(fi24)))))
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = (fi24 ^ 128)
                    }
                    _ = fi24
                }
                if v7 {
                    if (0.0 < (v6 + v6)) {
                        v6 += ((v6 / 1.5) * v6)
                        v22 = 128
                        v7 = (105 <= 127)
                    }
                    if v7 {
                        var v26 string = "hello"
                        fmt.Println(v6)
                        v6 = v6
                        _ = v26
                    }
                } else {
                    switch (-(fi23)) {
                    case 9223372036854775807:
                        if v7 {
                            var v27 float32 = float32(v6)
                            _ = v27
                        }
                    case 128:
                        v28 := GenS1{F0: 0, F1: (v22 - v22)}
                        _ = v28
                    case 0:
                        v8 -= ((v8 * float32(float64(1.5))) * ((v8 * float32(v6)) + 0.0))
                    }
                    switch 1 {
                    case -132:
                        var v29 string = (("a" + ("\n\t" + ("hello" + "foo bar"))) + ("\n\t" + ("foo bar" + "hello")))
                        _ = v29
                    default:
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 int32 = 609
                            var v32 int16 = 949
                            v32 = 0
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                    v8 = v8
                }
                _ = fi23
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
        }
        _ = v16
    }
    v6 -= v6
    fmt.Println(v6, v6, v6)
    var v33 int = 307
    var v34 int8 = -1
    if v7 {
        var v35 int = (-(v33))
        v36 := GenS2{F0: v34}
        _ = v35
        _ = v36
    } else {
        v8 += float32(1e10)
    }
    switch (v33 + (v33 & 32767)) {
    case 32767:
        if true {
            fmt.Println(v8)
            delete(v2, (v33 + v33))
        }
    case -1:
        v34 = v34
    case 32768:
        v33 = v33
    default:
        v37 := GenS3{F0: 2147483648, F1: -1}
        _ = v37
    }
    if (true || true) {
        var v38 int32 = 2147483647
        v39 := GenS3{F0: 160, F1: -9223372036854775808}
        var v40 int16 = 0
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 int16 = 451
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
