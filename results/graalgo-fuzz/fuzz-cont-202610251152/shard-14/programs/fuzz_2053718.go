package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0() uint64 {
    var v4 uint32 = 777
    switch -1 {
    case -671:
        var v5 uint16 = 42
        _ = v5
    case 9223372036854775807:
        var v6 float32 = float32(0.0)
        _ = v6
    case 32767:
        for fi7 := 0; fi7 < 20; fi7++ {
            v4 -= 0
            if true {
                var v8 string = ("a" + "a")
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 uint64 = 0
                    _ = fi9
                    _ = v10
                }
                var v11 int64 = -9223372036854775808
                _ = v11
                _ = v8
            }
            _ = fi7
        }
    default:
        switch 127 {
        case 472:
            var v12 int8 = 53
            _ = v12
        case 42:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: (v4 ^ (v4 - 65535)), F1: 18446744073709551615, F2: 4294967295}
                _ = fi13
                _ = v14
            }
        case 128:
            if false {
                v15 := GenS1{F0: (v4 + 0), F1: 93, F2: 0}
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: v4, F1: 1, F2: v4}
                    switch ("hello" + ((("a" + "") + ("hello" + "a")) + ("" + "hello"))) {
                    case "foo bar":
                        v18 := GenS1{F0: (v4 - 759), F1: 2, F2: v4}
                        _ = v18
                    case "hello":
                        v4 -= 2
                    case "":
                        v19 := GenS1{F0: (^(v4)), F1: 0, F2: v4}
                        _ = v19
                    default:
                        v20 := GenS1{F0: (v4 + v4), F1: 542, F2: v4}
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
                v21 := []int{-783, -1}
                _ = v15
                _ = v21
            } else {
                var v22 int = -2147483648
                var v23 int = v22
                _ = v22
                _ = v23
            }
        }
    }
    var v24 int32 = -681
    v4 = (v4 | v4)
    _ = v24
    _ = v4
    return 655
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 float32 = (((-2.25 / 3.14159) + float32(float64(0.1))) / -2.25)
    delete(v1, "hello")
    delete(v2, 355)
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint8 = 42
        v3 = (((3.14159 * (0.1 - (1e10 * float32(((1.5 - 1e10) - (3.14159 * (0.0 - 1e10))))))) - 0.1) + v3)
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS2{F0: int64(fi27)}
        v3 += -2.25
        _ = fi27
        _ = v28
    }
    v29 := GenS3{F0: -1}
    v3 -= v3
    _ = len(v1)
    v3 -= v3
    v30 := []int{32767, 42, 995, 32767}
    v3 += v3
    v31 := v1[("hello" + ("\n\t" + "\n\t"))]
    v32 := v1[("foo bar" + "foo bar")]
    v33 := GenS3{F0: 916}
    v32 -= (^(v31))
    var v34 uint = 1
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            if false {
                if true {
                    var v37 uint8 = 157
                    _ = v37
                }
            }
            var v38 int64 = 1
            _ = fi36
            _ = v38
        }
        _ = fi35
    }
    if (42 >= -29) {
        v39 := v2[42]
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
