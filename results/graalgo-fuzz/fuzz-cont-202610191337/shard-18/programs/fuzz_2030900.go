package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 uint64) int32 {
    p0 = p0
    p0 -= p0
    var v4 float64 = 3.14159
    switch 2147483647 {
    case 2147483647:
        if false {
            var v5 uint = 65535
            fmt.Println(v4, p0, v5)
            _ = v5
        } else {
            v6 := []int{42, -336}
            _ = v6
        }
    }
    _ = p0
    _ = v4
    return -366
}
func genFn1() uint16 {
    var v7 int64 = 0
    if true {
        v8 := GenS1{F0: -128, F1: 908}
        fmt.Println(v7)
        _ = v8
    }
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := []uint32{770, 342, 2147483647, 65535}
    var v9 bool = (100 > 7)
    var v10 bool = (1 == 0)
    v9 = false
    for fi11 := 0; fi11 < 10; fi11++ {
        if (2 >= 2) {
            if true {
                for fi12 := 0; fi12 < 20; fi12++ {
                    switch ("foo bar" + ("a" + "\n\t")) {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "a":
                        var v13 uint16 = 65535
                        _ = v13
                    default:
                        v14 := GenS1{F0: 100, F1: 42}
                        _ = v14
                    }
                    fmt.Println(fi12, v10)
                    v10 = v10
                    _ = fi12
                }
                delete(v2, fi11)
            } else {
                var v15 string = ("\n\t" + "")
                _ = v15
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                v10 = v9
                switch (fi16 ^ fi16) {
                case 128:
                    switch fi11 {
                    case 690:
                        v9 = (481 <= 1)
                    case 42:
                        var v17 int16 = 42
                        _ = v17
                    default:
                        fmt.Println(v10, v10, v10)
                    }
                default:
                    v10 = v10
                }
                v18 := GenS2{F0: 799, F1: (2 < 334), F2: 173}
                if false {
                    switch (fi11 | (^(fi11))) {
                    case 32767:
                        fmt.Println(fi11, fi16)
                        v10 = ((-2.25 < ((1.5 * 0.0) + float32(0.1))) && v9)
                        v10 = v9
                        if true {
                            var v19 int64 = 397
                            _ = v19
                        }
                    }
                    fmt.Println(fi16, fi16)
                } else {
                    var v20 float32 = 1e10
                    _ = v20
                }
                _ = fi16
                _ = v18
            }
            if (fi11 > fi11) {
                var v21 float32 = 0.1
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    fmt.Println(v10, v10)
                    _ = fi22
                }
                fmt.Println(v9)
                if (255 >= 42) {
                    var v23 int = fi11
                    _ = v23
                }
            }
        }
        _ = fi11
    }
    v9 = v9
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int32 = 740
        if v10 {
            v10 = true
            var v26 uint32 = 473
            fmt.Println(v26, v26)
            _ = v26
        } else {
            v3 = append(v3, 2147483647)
            v9 = v9
            for fi27 := 0; fi27 < 12; fi27++ {
                v9 = (!((1 >= 1) || ((v25 << ((407) & 63)) <= 1)))
                var v28 uint8 = 2
                switch "" {
                case "":
                    v28 -= 95
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 127
                    }
                case "foo bar":
                    if v9 {
                        var v29 int = 32767
                        _ = v29
                    } else {
                        var v30 string = "\n\t"
                        _ = v30
                    }
                }
                _ = fi27
                _ = v28
            }
        }
        fmt.Println(fi24, v25)
        _ = fi24
        _ = v25
    }
    if (0 == 65535) {
        var v31 int16 = 396
        v10 = true
        v32 := GenS2{F0: 1, F1: true, F2: 977}
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 17; fi33++ {
            v9 = v10
            delete(v1, "foo bar")
            v9 = v9
            _ = fi33
        }
    }
    v10 = v9
    v34 := v1[("\n\t" + "foo bar")]
    fmt.Println(v9, v10)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v9
}
