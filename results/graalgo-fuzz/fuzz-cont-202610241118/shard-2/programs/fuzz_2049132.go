package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 uint16, p1 uint8) uint8 {
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 10; fi6++ {
        if ("" <= "") {
            fmt.Println(v5)
            p0 = (p0 >> ((1) & 63))
        }
        _ = fi6
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v5
    return (p1 + (p1 & (p1 ^ p1)))
}

func main() {
    v1 := []string{""}
    var v2 bool = false
    var v3 int64 = -333
    v4 := make(map[string]bool)
    var v7 int = 4294967295
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v7 = v7
            v10 := GenS1{F0: "foo bar"}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 float32 = 0.1
            fmt.Println(v7, fi11, v2)
            if false {
                v2 = (127 <= 56)
                v13 := v4["\n\t"]
                var v14 uint32 = 4294967295
                v15 := GenS1{F0: "hello"}
                _ = v13
                _ = v14
                _ = v15
            } else {
                if (!(true && v2)) {
                    switch fi11 {
                    case 4294967295:
                        if ("" >= ("a" + ("hello" + "\n\t"))) {
                            _ = len(v4)
                            v3 -= (v3 & v3)
                        } else {
                            fmt.Println(v12)
                        }
                    case -1:
                        v12 -= v12
                    default:
                        delete(v4, "")
                    }
                    var v16 int16 = 1
                    _ = v16
                }
            }
            _ = fi11
            _ = v12
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: 42, F1: "hello", F2: ("foo bar" + "a")}
            var v19 int = fi17
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi8
    }
    var v20 uint8 = 38
    var v21 bool = v2
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += (v3 << ((2) & 63))
    v20 += (v20 >> ((4294967295) & 63))
    if (678 == 2) {
        if true {
            v22 := GenS1{F0: "foo bar"}
            _ = v22
        } else {
            v23 := GenS2{F0: 869, F1: ("foo bar" + "\n\t"), F2: "\n\t"}
            _ = v23
        }
        if (4294967295 == 584) {
            if ((float64((3.14159 * (float32(0.1) / (3.14159 / 1.5)))) >= 1.5) || (v21 && (1 > 617))) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    fmt.Println(v2, v7)
                    fmt.Println(v2, v2, v21)
                    fmt.Println(v2)
                    _ = fi24
                }
            } else {
                v20 = v20
            }
            v25 := GenS3{F0: -32768, F1: 1000}
            switch 1 {
            case 42:
                var v26 string = (("hello" + "hello") + "foo bar")
                _ = v26
            case 128:
                var v27 uint64 = 42
                _ = v27
            }
            _ = v25
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v28 := GenS2{F0: 1, F1: "a", F2: ""}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v7
}
