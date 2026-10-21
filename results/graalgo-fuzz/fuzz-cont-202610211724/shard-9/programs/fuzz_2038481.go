package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 int64) uint16 {
    p0 = (p0 - 1)
    if false {
        v7 := GenS1{F0: 42, F1: (^(p1))}
        fmt.Println(p1, p1)
        if true {
            v8 := []int{32767, 32767}
            var v9 uint64 = 0
            var v10 uint32 = 65535
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                if (!true) {
                    p1 = p0
                    v12 := []int{(fi11 * 1), (fi11 | fi11), fi11, fi11}
                    _ = v12
                } else {
                    p1 = (p0 | 0)
                    v13 := GenS1{F0: 255, F1: p0}
                    fmt.Println(p0, p1, p1)
                    _ = v13
                }
                var v14 uint = 255
                _ = fi11
                _ = v14
            }
            v15 := GenS1{F0: 255, F1: 42}
            _ = v15
        }
        _ = v7
    }
    if true {
        fmt.Println(p0, p1)
        var v16 uint32 = 4294967295
        for fi17 := 0; fi17 < 3; fi17++ {
            switch "a" {
            case "a":
                if (true && false) {
                    var v18 float32 = (0.0 * -2.25)
                    _ = v18
                }
            case "hello":
                var v19 int8 = 0
                _ = v19
            }
            _ = fi17
        }
        _ = v16
    } else {
        if true {
            if true {
                for fi20 := 0; fi20 < 3; fi20++ {
                    p1 += p0
                    _ = fi20
                }
                if false {
                    if true {
                        p0 += 0
                        if false {
                            v21 := []int{2147483647}
                            var v22 bool = (true && false)
                            p1 += p0
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int32 = -1
                            var v24 string = "foo bar"
                            _ = v23
                            _ = v24
                        }
                    }
                    v25 := []int{-697, 461}
                    p1 -= (p1 * 0)
                    _ = v25
                } else {
                    var v26 int = -1
                    _ = v26
                }
            }
            v27 := GenS1{F0: 65535, F1: p0}
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint16 = 498
                _ = fi28
                _ = v29
            }
            _ = v27
        } else {
            p1 = 731
            fmt.Println(p0)
            var v30 int16 = 0
            _ = v30
        }
    }
    var v31 int64 = (p0 % (((p0 ^ (^(p1)))) | 1))
    _ = p0
    _ = p1
    _ = v31
    return 32767
}

func main() {
    var v1 int64 = 42
    var v2 uint = 0
    var v3 float64 = 1e10
    var v4 uint64 = 628
    v5 := []int{}
    var v6 float64 = v3
    fmt.Println(v3, v3)
    var v32 bool = (true || true)
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS1{F0: 65535, F1: (v1 - v1)}
        for fi35 := 0; fi35 < 5; fi35++ {
            v5 = append(v5, fi35)
            var v36 uint = 255
            switch (fi35 * 9223372036854775807) {
            case -1:
                fmt.Println(v36)
            default:
                v37 := GenS2{F0: 255, F1: genFn0(v1, (v1 >> ((1) & 63))), F2: (v4 << ((v36) & 63))}
                _ = v37
            }
            var v38 uint16 = 32767
            _ = fi35
            _ = v36
            _ = v38
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v1, v6)
    var v39 string = (("\n\t" + "foo bar") + (("\n\t" + "hello") + ("\n\t" + ("foo bar" + ""))))
    switch 441 {
    case -104:
        var v40 int32 = -1
        _ = v40
    default:
        fmt.Println(v6)
    }
    v41 := GenS3{F0: 42, F1: 0, F2: 2}
    var v42 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
