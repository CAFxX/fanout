package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 uint16) int32 {
    p0 += p0
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    _ = p0
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := []string{"hello", "foo bar", "\n\t", "hello"}
    v5 := []uint32{551}
    var v6 uint16 = 787
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9, fi9)
        v3["hello"] = true
        _ = fi9
    }
    v6 = v6
    if (true || (2 >= (v6 | v6))) {
        v6 = 255
        for fi10 := 0; fi10 < 17; fi10++ {
            switch "" {
            case "a":
                _ = len(v5)
            case "\n\t":
                fmt.Println(v6, v6)
            default:
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 float32 = 3.14159
                    v6 = (v6 + 1)
                    var v13 int16 = -1
                    v13 = v13
                    fmt.Println(v13)
                    var v14 uint16 = (v6 + v6)
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi10
        }
        v15 := GenS1{F0: 614}
        _ = v15
    }
    var v16 bool = true
    var v17 bool = (v16 && v16)
    v2[32767] = -257
    var v18 uint32 = 2147483648
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int8 = 0
        for fi21 := 0; fi21 < 11; fi21++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi21
        }
        if (508 == 629) {
            var v22 uint64 = 42
            _ = v22
        } else {
            v18 = v18
            if ((v20 - v20) > 0) {
                if true {
                    if v16 {
                        var v23 int32 = 2147483647
                        v24 := GenS2{F0: (v18 * (v18 - v18)), F1: 2147483648}
                        _ = v23
                        _ = v24
                    }
                    if (42 > 0) {
                        var v25 int8 = int8(v18)
                        var v26 int32 = -2147483648
                        _ = v25
                        _ = v26
                    }
                }
            }
            var v27 uint8 = uint8(v6)
            _ = v27
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v16, v16)
    switch 4294967295 {
    case 0:
        fmt.Println(v6)
    case 32768:
        var v28 uint = uint(v18)
        _ = v28
    case 1:
        v29 := GenS3{F0: 255, F1: "foo bar"}
        _ = v29
    default:
        v16 = v17
    }
    v18 = v18
    v30 := GenS1{F0: 960}
    v18 = (v18 ^ (v18 << ((1) & 63)))
    delete(v2, -1)
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 int16 = 32767
        v4 = append(v4, "\n\t")
        _ = fi31
        _ = v32
    }
    v6 = (v6 << ((65535) & 63))
    if (623 >= 65535) {
        v33 := GenS3{F0: 2, F1: ("a" + "")}
        v6 = 42
        _ = v33
    } else {
        var v34 float64 = 1e10
        v35 := GenS2{F0: (v18 | v18), F1: (v18 + v18)}
        var v36 uint8 = 126
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
