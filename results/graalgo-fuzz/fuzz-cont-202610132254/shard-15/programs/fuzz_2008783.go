package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 uint, p1 int16, p2 uint8) int8 {
    var v7 float32 = 0.0
    if (-2.25 < (1.5 / 3.14159)) {
        fmt.Println(v7)
    } else {
        if (77 >= 1) {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{fi8, (fi8 / (((fi8 % ((fi8) | 1))) | 1)), (fi8 & (fi8 % ((14) | 1)))}
                _ = fi8
                _ = v9
            }
        }
    }
    if false {
        if true {
            fmt.Println(p2, p2)
            v7 += 0.0
        } else {
            p0 -= 4294967295
            var v10 uint16 = 65535
            var v11 float64 = 0.1
            _ = v10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            switch (^(fi12)) {
            case -413:
                if ("\n\t" > "\n\t") {
                    p0 = p0
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v14 := []int{fi12, -2147483648, (fi12 + -9223372036854775808), int(p2)}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        v7 += float32(0.1)
                        _ = fi13
                        _ = v14
                    }
                    var v15 int = (fi12 << ((p0) & 63))
                    _ = v15
                } else {
                    var v16 bool = (true && true)
                    _ = v16
                }
            case 4294967295:
                if false {
                    switch 4294967295 {
                    case -200:
                        fmt.Println(p0)
                    }
                    var v17 int8 = 127
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        switch (^(fi12)) {
                        case 709:
                            fmt.Println(fi18)
                        case 4294967295:
                            v19 := []int{fi12, (-(fi12)), fi18, (fi18 >> ((65535) & 63))}
                            _ = v19
                        }
                        v7 = 1.5
                        p1 += p1
                        _ = fi18
                    }
                }
            case 2147483647:
                p2 -= (p2 + p2)
            default:
                fmt.Println(p0)
            }
            _ = fi12
        }
        if ((false && ((p2 + 255) < 127)) || (!(("hello" + "hello") != "foo bar"))) {
            var v20 float32 = 1e10
            _ = v20
        }
        if false {
            v21 := []int{-9223372036854775808, 128, 32768, 526}
            _ = v21
        }
    }
    v22 := GenS1{F0: -1, F1: (p0 + p0), F2: 0}
    var v23 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v7
    return 1
}

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 699
    var v3 int8 = 0
    var v4 uint8 = 43
    v5 := make(map[string]bool)
    var v6 uint32 = 0
    var v24 int = v1
    var v25 int8 = (v3 & (^(v3)))
    for fi26 := 0; fi26 < 2; fi26++ {
        v25 += v3
        _ = fi26
    }
    v27 := GenS1{F0: (v25 - genFn0(24, -107, v4)), F1: 255, F2: 2147483647}
    v25 += (v3 / ((-1) | 1))
    v24 -= 2147483647
    switch (v1 + (v1 * (v24 * v24))) {
    case 117:
        v28 := []int{(v24 ^ (v1 >> ((768) & 63)))}
        _ = v28
    case 1:
        v29 := v5["hello"]
        _ = v29
    }
    v24 += v24
    for fi30 := 0; fi30 < 7; fi30++ {
        if (true && false) {
            v31 := GenS2{F0: 2147483647}
            _ = v31
        } else {
            v32 := GenS1{F0: (-(v25)), F1: 664, F2: -1}
            switch "hello" {
            case "foo bar":
                fmt.Println(v24)
            case "a":
                switch "foo bar" {
                case "hello":
                    var v33 int16 = -32768
                    _ = v33
                case "\n\t":
                    v34 := GenS3{F0: -678, F1: -185, F2: 777}
                    _ = v34
                }
            }
            var v35 int8 = genFn0(4294967295, 42, 128)
            _ = v32
            _ = v35
        }
        v36 := v5[("\n\t" + ("foo bar" + "foo bar"))]
        v37 := []int{fi30, fi30, v1, int(v2)}
        _ = fi30
        _ = v36
        _ = v37
    }
    delete(v5, "")
    v38 := []int{v24, v1, (v24 - (v24 >> ((0) & 63))), v24}
    v1 += v1
    for fi39 := 0; fi39 < 4; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            delete(v5, (("a" + "a") + "foo bar"))
            fmt.Println(v1, fi39, v24)
            if len(v38) > 0 {
                v38[0] = fi40
            }
            var v41 string = "hello"
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v1)
                v42 := GenS2{F0: 42}
                _ = v42
            default:
                var v43 int16 = 358
                _ = v43
            }
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
