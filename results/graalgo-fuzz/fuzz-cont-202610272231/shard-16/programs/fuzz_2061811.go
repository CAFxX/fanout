package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32) int16 {
    if (0.0 == float32((-2.25 * 3.14159))) {
        p0 += (p0 % ((42) | 1))
        p0 = (p0 << ((694) & 63))
        switch -2147483648 {
        case 4:
            var v6 float64 = ((-2.25 * (-2.25 / 0.1)) - float64(-2.25))
            _ = v6
        case 1:
            if (true && (0.1 <= -2.25)) {
                switch "" {
                case "":
                    var v7 int64 = 0
                    _ = v7
                case "foo bar":
                    p0 -= 2147483647
                default:
                    p0 = 0
                }
                v8 := []int{}
                _ = v8
            }
        }
    } else {
        switch 4294967295 {
        case 42:
            var v9 uint16 = 0
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = 42
            var v13 string = "\n\t"
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 uint = 255
    v14 = v14
    var v15 int8 = -99
    _ = p0
    _ = v14
    _ = v15
    return 711
}

func main() {
    v1 := []string{(("" + ((("foo bar" + "\n\t") + "") + "hello")) + ("foo bar" + (("" + "\n\t") + ("hello" + ("a" + "hello"))))), "\n\t", "hello"}
    var v2 float32 = (-2.25 - ((-2.25 * 1.5) * 0.0))
    v3 := make(map[int]string)
    var v4 uint32 = 2147483648
    v5 := []uint32{82, (v4 / ((v4) | 1))}
    if false {
        v5 = append(v5, v4)
        for fi16 := 0; fi16 < 19; fi16++ {
            v2 = v2
            if true {
                for fi17 := 0; fi17 < 16; fi17++ {
                    switch fi17 {
                    case -1:
                        var v18 uint8 = 128
                        _ = v18
                    default:
                        v4 += v4
                    }
                    v2 = v2
                    _ = fi17
                }
                v19 := GenS1{F0: 342, F1: -108}
                _ = v19
            } else {
                v2 = ((v2 + (float32(float64(v2)) * float32((3.14159 - float64(v2))))) + 1.5)
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 uint16 = 32767
                    var v22 int32 = 42
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            _ = fi16
        }
    } else {
        var v23 uint32 = v4
        _ = v23
    }
    if false {
        var v24 uint = 176
        _ = v24
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            v4 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi25
        }
    }
    var v26 uint = 65535
    var v27 int64 = 854
    if false {
        var v28 uint64 = 278
        fmt.Println(v26, v28)
        _ = v28
    } else {
        var v29 int16 = -1
        if (true || true) {
            var v30 int = 2147483647
            _ = v30
        }
        _ = v29
    }
    v31 := GenS2{F0: 117}
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 uint = (^(v26))
            if ((v27 ^ (v27 | v27)) < 0) {
                var v35 bool = true
                v36 := GenS3{F0: 4294967295}
                var v37 int64 = v27
                _ = v35
                _ = v36
                _ = v37
            }
            v34 = (v26 ^ (v34 ^ v34))
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        for fi39 := 0; fi39 < 17; fi39++ {
            v27 = v27
            switch "a" {
            case "\n\t":
                var v40 bool = true
                _ = v40
            }
            delete(v3, 0)
            _ = fi39
        }
        v41 := GenS2{F0: 255}
        v4 += (v4 - v4)
        _ = fi38
        _ = v41
    }
    var v42 uint32 = (v4 & 2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
