package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 string, p1 int) uint16 {
    v6 := GenS1{F0: 32767}
    v7 := []int{(p1 >> ((0) & 63)), (p1 % ((p1) | 1)), p1, p1}
    fmt.Println(p1, p0)
    v8 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = (v1 - (v1 / ((v1) | 1)))
    var v3 float64 = 1.5
    var v4 int16 = -32768
    var v5 float64 = (v3 / 1e10)
    v3 += 0.0
    v3 += float64(1e10)
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 -= (v5 * v5)
        v10 := []int{fi9, 32767, (fi9 - -1)}
        var v11 uint8 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    if (1 < genFn0((("a" + "foo bar") + ("hello" + "a")), 13)) {
        v12 := GenS1{F0: 255}
        _ = v12
    }
    var v13 float64 = -2.25
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: genFn0("\n\t", 9223372036854775807)}
            v16 := GenS1{F0: 32767}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v17 := []int{-899, 32767, 1, 9223372036854775807}
    v18 := GenS1{F0: uint16(v2)}
    var v19 uint16 = uint16(v1)
    if false {
        v20 := GenS1{F0: v19}
        _ = v20
    } else {
        var v21 uint32 = 2
        for fi22 := 0; fi22 < 19; fi22++ {
            v4 = 834
            for fi23 := 0; fi23 < 19; fi23++ {
                if false {
                    if false {
                        v21 -= v21
                        var v24 uint = 65535
                        v1 = v1
                        v1 += (v2 ^ v2)
                        v1 -= v2
                        _ = v24
                    } else {
                        if false {
                            v25 := GenS1{F0: (^(v19))}
                            v3 -= v5
                            _ = v25
                        } else {
                            var v26 float32 = float32(v5)
                            fmt.Println(fi23, fi23)
                            _ = v26
                        }
                        v27 := GenS1{F0: (v19 << ((731) & 63))}
                        _ = v27
                    }
                    switch 42 {
                    case 2147483647:
                        v28 := GenS1{F0: 42}
                        _ = v28
                    case -567:
                        fmt.Println(fi22, v19)
                    }
                } else {
                    var v29 int8 = 42
                    _ = v29
                }
                fmt.Println(v19)
                var v30 int8 = 42
                var v31 float32 = (0.0 * -2.25)
                _ = fi23
                _ = v30
                _ = v31
            }
            if (!true) {
                v5 += 1.5
                var v32 int8 = -93
                _ = v32
            } else {
                if false {
                    switch "" {
                    case "hello":
                        var v33 uint32 = 65535
                        _ = v33
                    case "":
                        fmt.Println(v5, fi22, v5)
                    case "a":
                        switch ("a" + "a") {
                        case "a":
                            var v34 float32 = float32((v3 / v5))
                            _ = v34
                        }
                    }
                    _ = len(v17)
                    for fi35 := 0; fi35 < 19; fi35++ {
                        var v36 int = (fi35 >> ((18446744073709551615) & 63))
                        var v37 float32 = 0.0
                        var v38 int = -2147483648
                        v39 := GenS1{F0: (v19 ^ 65535)}
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                fmt.Println(v4, v2)
                var v40 int16 = 648
                _ = v40
            }
            _ = fi22
        }
        v3 = (v13 / (v13 / float64((float32(v3) + 3.14159))))
        _ = v21
    }
    v17 = append(v17, 32767)
    fmt.Println(v4, v13)
    fmt.Println(v2, v1, v19)
    v41 := GenS2{F0: -37, F1: 259, F2: -253}
    if (4294967295 < (v1 + v1)) {
        if (false && ("foo bar" >= "foo bar")) {
            v42 := GenS3{F0: true}
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
