package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint; F2 int32 }
func genFn0(p0 string) uint16 {
    p0 = (p0 + (p0 + (((p0 + p0) + ("a" + ((((p0 + ("\n\t" + p0)) + p0) + "\n\t") + (p0 + "")))) + p0)))
    fmt.Println(p0)
    v4 := []int{}
    var v5 int32 = -1
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]int)
    var v3 int = 4294967295
    v1 -= (-2.25 / 0.1)
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            if true {
                v7 := []int{}
                _ = v7
            }
            var v8 int16 = -32768
            for fi9 := 0; fi9 < 19; fi9++ {
                if false {
                    v3 -= v3
                }
                var v10 string = "foo bar"
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        delete(v2, (fi6 >> ((2) & 63)))
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                v1 = 1.5
                v1 = 0.1
                fmt.Println(fi11, fi11)
                v1 = ((v1 - v1) - 0.1)
            } else {
                v12 := GenS1{F0: genFn0("")}
                if (("hello" + "hello") >= "a") {
                    if true {
                        switch "hello" {
                        case "":
                            var v13 int = fi6
                            _ = v13
                        case "\n\t":
                            fmt.Println(v1, fi11, v3)
                        }
                        var v14 uint = 18446744073709551615
                        v15 := GenS1{F0: 520}
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint = 2
                        v17 := GenS1{F0: 829}
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS1{F0: 65535}
                    _ = v18
                }
                var v19 uint8 = 2
                _ = v12
                _ = v19
            }
            _ = fi11
        }
        _ = fi6
    }
    v20 := GenS2{F0: -604, F1: 476, F2: -946}
    v21 := v2[32767]
    v22 := v2[(v21 | (v21 << ((192) & 63)))]
    var v23 float64 = (v1 - v1)
    if false {
        v23 += 0.0
        var v24 uint8 = 128
        _ = v24
    }
    if (0.0 <= v1) {
        var v25 uint16 = 738
        v3 = 4294967295
        v26 := GenS1{F0: (v25 * v25)}
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                if (-2.25 != (1e10 / -2.25)) {
                    v1 -= -2.25
                } else {
                    switch (fi27 ^ 128) {
                    case 42:
                        v29 := GenS2{F0: 0, F1: 18446744073709551615, F2: -646}
                        _ = v29
                    case 452:
                        v30 := GenS2{F0: -1, F1: 18446744073709551615, F2: -2147483648}
                        _ = v30
                    }
                    var v31 uint = 302
                    v2[(v3 % (((v22 / ((-33) | 1))) | 1))] = (v3 - (-(fi27)))
                    _ = v31
                }
                if (127 == 0) {
                    delete(v2, (^(v22)))
                    if false {
                        switch (fi28 & (v22 ^ (v3 * 2147483647))) {
                        case -170:
                            fmt.Println(fi27, fi28, fi28)
                        case -634:
                            fmt.Println(fi27, v22, v1)
                        case 1:
                            v32 := []int{v3, (fi28 / ((v3) | 1)), -866, fi28}
                            _ = v32
                        }
                        fmt.Println(v3)
                        v22 = (v21 & -1)
                    }
                    switch "a" {
                    case "foo bar":
                        switch v21 {
                        case -365:
                            var v33 float64 = (v23 / v1)
                            _ = v33
                        case 42:
                            var v34 uint64 = 1
                            _ = v34
                        default:
                            var v35 int32 = -2147483648
                            _ = v35
                        }
                    case "":
                        v3 -= v22
                        v36 := GenS1{F0: 0}
                        switch "" {
                        case "foo bar":
                            var v37 int = -555
                            _ = v37
                        }
                        _ = v36
                    case "\n\t":
                        fmt.Println(fi27, v1)
                    }
                }
                _ = fi28
            }
            fmt.Println(v22, v3, v23)
            _ = fi27
        }
        v23 += float64(1e10)
    }
    fmt.Println(v21)
    fmt.Println(v22)
    fmt.Println(v1, v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
