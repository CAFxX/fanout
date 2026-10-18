package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0() uint16 {
    if (!false) {
        var v6 float64 = 1e10
        _ = v6
    }
    v7 := GenS1{F0: 0}
    switch 128 {
    case 502:
        v8 := []int{4294967295, 421, 473, 9223372036854775807}
        _ = v8
    case 2147483647:
        var v9 int8 = -128
        _ = v9
    default:
        if false {
            var v10 int16 = 31
            _ = v10
        }
    }
    _ = v7
    return 473
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{0, 2}
    var v3 string = "hello"
    var v4 float32 = 1.5
    var v5 int32 = 133
    v3 = v3
    v4 = (0.0 + v4)
    v4 += v4
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v5 -= (v5 << ((1) & 63))
            v13 := GenS1{F0: 1000}
            _ = fi12
            _ = v13
        }
        v3 = (((v3 + ("hello" + (v3 + (("foo bar" + (v3 + "")) + "hello")))) + "foo bar") + "")
        if true {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 bool = (true || (42 == 1))
                if v15 {
                    if len(v2) > 0 {
                        v2[0] = 850
                    }
                } else {
                    var v16 int16 = 0
                    v17 := GenS1{F0: v16}
                    var v18 uint = 0
                    _ = v16
                    _ = v17
                    _ = v18
                }
                v19 := GenS2{F0: 518, F1: 65535, F2: 0}
                _ = fi14
                _ = v15
                _ = v19
            }
        }
        _ = fi11
    }
    var v20 int = 128
    v21 := GenS2{F0: -958, F1: 176, F2: 0}
    v1[v3] = 127
    if (!true) {
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS3{F0: 4294967295, F1: (((v3 + (v3 + (v3 + v3))) + ("\n\t" + (((((("\n\t" + v3) + ("foo bar" + "foo bar")) + (("foo bar" + v3) + "foo bar")) + (v3 + "hello")) + v3) + ("a" + "a")))) + v3)}
            fmt.Println(fi22, v4)
            v4 = ((v4 + float32((1e10 * 0.1))) / float32(((float64(1.5) - 0.0) + 3.14159)))
            _ = fi22
            _ = v23
        }
    } else {
        if (!false) {
            var v24 uint8 = 218
            _ = v24
        }
        v5 = v5
        switch v20 {
        case 1:
            delete(v1, v3)
        case 9223372036854775807:
            if (false && true) {
                v3 = v3
                for fi25 := 0; fi25 < 20; fi25++ {
                    v20 += (^(fi25))
                    v26 := v1[(("foo bar" + "hello") + ((v3 + v3) + "foo bar"))]
                    if ((430 < 4294967295) && (!true)) {
                        _ = len(v1)
                        v4 = (-2.25 - (3.14159 * float32(0.0)))
                    }
                    _ = fi25
                    _ = v26
                }
            }
        }
    }
    delete(v1, ((v3 + v3) + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
