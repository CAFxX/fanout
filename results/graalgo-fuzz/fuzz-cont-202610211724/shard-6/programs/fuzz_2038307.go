package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0() bool {
    var v5 uint64 = 456
    switch ("\n\t" + "a") {
    case "":
        var v6 uint16 = 563
        _ = v6
    case "\n\t":
        v7 := []int{85}
        _ = v7
    case "foo bar":
        switch 1 {
        case -9223372036854775808:
            var v8 string = (("\n\t" + "foo bar") + ("hello" + "hello"))
            _ = v8
        case 128:
            if (!(867 < 1)) {
                if true {
                    var v9 int32 = -1
                    var v10 int64 = 9223372036854775807
                    fmt.Println(v5, v9)
                    _ = v10
                    _ = v9
                }
                var v11 float64 = 3.14159
                _ = v11
            }
        case 32767:
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int8 = -1
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        }
    default:
        fmt.Println(v5)
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v5)
            v17 := []int{fi15, int(v5)}
            v18 := GenS1{F0: fi15}
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: (fi15 | 4294967295)}
        _ = fi15
        _ = v19
    }
    v20 := []int{42, 1}
    v5 -= (v5 / ((42) | 1))
    _ = v20
    _ = v5
    return (687 > 2147483647)
}
func genFn1(p0 int32) int8 {
    for fi21 := 0; fi21 < 19; fi21++ {
        p0 += (p0 * (p0 << ((1) & 63)))
        _ = fi21
    }
    var v22 float32 = 1.5
    var v23 int16 = -141
    _ = p0
    _ = v22
    _ = v23
    return 127
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    var v4 float32 = (-2.25 * 0.1)
    v1 -= v1
    v24 := GenS1{F0: 2147483647}
    v4 += (float32(0.1) - 1e10)
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := []int{(^(fi25)), (fi25 & fi25), fi25, fi25}
        for fi27 := 0; fi27 < 8; fi27++ {
            if (0 != 42) {
                fmt.Println(v3, fi27)
                v3 = v3
                switch int(v1) {
                case -2147483648:
                    var v28 uint16 = 42
                    switch "foo bar" {
                    case "foo bar":
                        v4 = (v4 / (float32(-2.25) * float32((1e10 - (0.1 / float64(float32(1e10)))))))
                    case "a":
                        var v29 bool = true
                        _ = v29
                    }
                    var v30 int32 = -2147483648
                    _ = v28
                    _ = v30
                case 105:
                    v1 += (^(v1))
                default:
                    v3 = -9223372036854775808
                }
                v31 := GenS2{F0: ("foo bar" + (("a" + "") + (("" + ("" + "foo bar")) + "foo bar")))}
                v2[fi25] = (^(v1))
                _ = v31
            }
            _ = fi27
        }
        var v32 int32 = 42
        _ = fi25
        _ = v26
        _ = v32
    }
    if (128 != 0) {
        v3 += v3
    }
    var v33 uint64 = 4294967295
    switch 9223372036854775807 {
    case 6:
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := GenS2{F0: ("foo bar" + "a")}
            v36 := []int{}
            var v37 int64 = -551
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
    case -144:
        v3 = 192
    default:
        v33 -= 1
    }
    var v38 int32 = -2147483648
    if true {
        var v39 uint = 18446744073709551615
        _ = v39
    }
    v40 := []int{-940}
    var v41 int64 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v40
    _ = v41
}
