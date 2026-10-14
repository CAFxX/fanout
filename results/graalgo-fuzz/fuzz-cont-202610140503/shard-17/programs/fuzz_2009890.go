package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 int64) int {
    var v5 uint8 = 0
    var v6 uint8 = v5
    if (((((1.5 / -2.25) / ((1e10 * float64(0.0)) + (-2.25 + 0.0))) == (-2.25 / (float64((float32(0.1) / -2.25)) + 1e10))) && (42 > 174)) && (!true)) {
        var v7 bool = (42 != 32767)
        switch "foo bar" {
        case "a":
            v8 := []int{}
            _ = v8
        default:
            if (704 < 65535) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 uint64 = 4294967295
                    var v11 float64 = 0.1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        }
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := []uint32{885, 0, 4294967295, 339}
    var v2 int8 = 100
    var v3 uint32 = 65535
    v4 := make(map[string]uint32)
    var v12 uint64 = 44
    if true {
        if (495 == 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v13 int64 = -1
            v13 -= v13
            _ = v13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: 4294967295, F1: 4294967295, F2: fi14}
            fmt.Println(v2)
            v16 := GenS1{F0: 65535, F1: v12, F2: fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        v3 = v3
        for fi17 := 0; fi17 < 18; fi17++ {
            v2 -= 127
            var v18 int = genFn0(-308)
            v18 = fi17
            _ = fi17
            _ = v18
        }
        var v19 int = 922
        _ = v19
    }
    v2 = (v2 ^ (^(v2)))
    v4["foo bar"] = (^(v3))
    var v20 bool = false
    var v21 uint = 1
    v12 = 891
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v3, v2)
        if (v20 && ((0 >= 1) && true)) {
            var v23 string = "hello"
            var v24 string = v23
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    v25 := GenS2{F0: 1, F1: 88, F2: 0}
    v26 := GenS2{F0: 128, F1: 127, F2: (v12 - v12)}
    fmt.Println(v21, v2, v3)
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 string = "foo bar"
        v29 := GenS1{F0: v21, F1: (v12 + v12), F2: 9223372036854775807}
        for fi30 := 0; fi30 < 10; fi30++ {
            switch ((v28 + v28) + ("hello" + v28)) {
            case "":
                var v31 int16 = 0
                _ = v31
            default:
                if v20 {
                    var v32 uint8 = 2
                    var v33 int8 = (-(v2))
                    _ = v32
                    _ = v33
                }
            }
            _ = fi30
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        switch ("hello" + ("a" + (((("foo bar" + "\n\t") + ((("\n\t" + ("" + "")) + ("" + "hello")) + ("foo bar" + ""))) + "a") + ("\n\t" + "foo bar")))) {
        case "":
            fmt.Println(v12)
        case "hello":
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 int16 = 0
                _ = fi34
                _ = v35
            }
        default:
            var v36 int8 = (v2 * (v2 / ((v2) | 1)))
            _ = v36
        }
        _ = len(v4)
        if v20 {
            var v37 int16 = 1000
            _ = v37
        }
    } else {
        switch int(v2) {
        case 2147483647:
            fmt.Println(v20, v21)
        default:
            v3 -= v3
        }
        if true {
            fmt.Println(v21)
            var v38 uint64 = 4294967295
            _ = v38
        }
        if true {
            v39 := GenS3{F0: 4294967295, F1: 0}
            for fi40 := 0; fi40 < 18; fi40++ {
                for fi41 := 0; fi41 < 3; fi41++ {
                    var v42 int32 = -1
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
            fmt.Println(v21, v2)
            _ = v39
        } else {
            v43 := GenS2{F0: 255, F1: 42, F2: 18446744073709551615}
            fmt.Println(v12)
            fmt.Println(v3, v20, v12)
            _ = v43
        }
    }
    v12 -= (v12 | 1)
    fmt.Println(v12, v21, v2)
    fmt.Println(v12, v20, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
