package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64) int64 {
    p0 = (p0 * (^(p0)))
    fmt.Println(p0)
    var v5 int64 = p0
    var v6 uint64 = 1
    _ = p0
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 float32 = float32(float64(0.1))
    var v2 uint = 0
    v3 := []string{"a", ("foo bar" + ("" + "")), "foo bar", "hello"}
    v4 := []uint32{2, 1, 1}
    for fi7 := 0; fi7 < 5; fi7++ {
        v2 = (v2 | (v2 >> ((v2) & 63)))
        _ = fi7
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println(v1)
    var v8 uint8 = 175
    switch 42 {
    case -110:
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                var v10 int64 = 0
                v10 = (v10 << ((101) & 63))
                _ = len(v3)
                _ = v10
            }
            _ = fi9
        }
    case 42:
        v2 = v2
    case 1:
        v1 = 1.5
    default:
        _ = len(v3)
    }
    v11 := GenS1{F0: 67, F1: 94}
    var v12 bool = (-522 <= -1)
    v2 += 2
    v2 = v2
    v13 := GenS1{F0: -77, F1: -1}
    if false {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 uint = v2
            v12 = v12
            v8 = (v8 << (((v2 * (v2 % (((^(v15))) | 1)))) & 63))
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS2{F0: 288}
            switch fi16 {
            case 32768:
                v3 = append(v3, "foo bar")
            case 128:
                var v18 uint = v2
                _ = v18
            }
            v19 := GenS2{F0: 65535}
            _ = fi16
            _ = v17
            _ = v19
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if true {
            if true {
                v12 = (!v12)
                var v20 float64 = (1.5 - 0.0)
                _ = len(v3)
                _ = v20
            } else {
                var v21 uint8 = 2
                _ = v21
            }
        } else {
            v3 = append(v3, "hello")
            v2 = v2
            switch 236 {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 18; fi22++ {
                    v12 = ((32767 == 1) && ((42 == 64) && (-128 < 80)))
                    var v23 int32 = -1
                    var v24 int16 = -32768
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case -418:
                v12 = false
            case -2147483648:
                var v25 int32 = -2147483648
                _ = v25
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
