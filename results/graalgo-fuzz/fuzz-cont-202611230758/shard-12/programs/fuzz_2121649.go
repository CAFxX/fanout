package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    v6 := GenS1{F0: 689}
    var v7 string = "a"
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            v9 := GenS1{F0: 65535}
            if true {
                if false {
                    v7 = "\n\t"
                } else {
                    switch (fi8 * fi8) {
                    case 32768:
                        v10 := []int{fi8}
                        _ = v10
                    case 32767:
                        var v11 bool = (((v7 + "hello") >= v7) && false)
                        _ = v11
                    case 128:
                        v12 := GenS2{F0: 255, F1: -496, F2: -513}
                        _ = v12
                    default:
                        switch (fi8 << ((1) & 63)) {
                        case 553:
                            v13 := GenS2{F0: 128, F1: 0, F2: 0}
                            _ = v13
                        default:
                            v14 := []int{(fi8 - -9223372036854775808)}
                            _ = v14
                        }
                    }
                    v15 := []int{fi8, fi8}
                    _ = v15
                }
                var v16 int16 = -1
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 int16 = 1000
                    var v19 int64 = -453
                    switch v7 {
                    case "a":
                        v19 = (v19 ^ 1)
                    case "\n\t":
                        var v20 string = (v7 + "foo bar")
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                var v21 uint16 = 554
                _ = v21
            }
            v7 = "foo bar"
            _ = v9
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 1000
}
func genFn1() string {
    var v22 uint8 = 2
    var v23 uint8 = (v22 ^ v22)
    _ = v22
    _ = v23
    return ((("\n\t" + "") + "foo bar") + ("foo bar" + "hello"))
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 127
    var v3 string = "a"
    var v4 int8 = 42
    var v5 uint = 42
    fmt.Println(v4)
    v24 := GenS1{F0: v5}
    v25 := v1[v3]
    if false {
        v1[v3] = (^(v2))
        switch (v25 ^ v2) {
        case 4294967295:
            v5 -= 65535
        }
        v26 := GenS3{F0: 42, F1: -629}
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(fi27, v5, v5)
        v28 := []int{}
        if false {
            if ((-(v25)) >= int(v5)) {
                var v29 int16 = 1
                _ = v29
            }
            v25 += (v25 >> ((v5) & 63))
            v5 += 42
        }
        _ = fi27
        _ = v28
    }
    v25 += v25
    v30 := []int{(v25 % ((v25) | 1)), 32767, 1, 32767}
    if (875 >= 0) {
        if false {
            var v31 int32 = -1
            v32 := GenS1{F0: v5}
            v1[(v3 + v3)] = 1
            _ = v31
            _ = v32
        } else {
            switch v25 {
            case -862:
                v4 = v4
            case -338:
                var v33 uint8 = 255
                _ = v33
            case 128:
                v4 = 100
            default:
                _ = len(v30)
            }
        }
    }
    var v34 uint8 = 139
    if (!true) {
        v2 -= -859
        if ((0.1 * ((float32(3.14159) + -2.25) + 3.14159)) < 1e10) {
            v35 := GenS4{F0: 1}
            if (65535 != 362) {
                var v36 uint8 = 2
                fmt.Println(v34, v2)
                _ = v36
            }
            _ = v35
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 uint64 = 4294967295
        v30 = append(v30, v25)
        var v39 uint8 = (v34 & v34)
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int16 = int16(v34)
    v5 = v5
    v2 -= (v25 + (-(v25)))
    v34 = v34
    var v41 int32 = -1
    var v42 int16 = (v40 + (v40 * 1000))
    v4 -= (v4 ^ (^(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
