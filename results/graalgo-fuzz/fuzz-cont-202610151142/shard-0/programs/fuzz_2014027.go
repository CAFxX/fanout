package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 string) int64 {
    var v6 uint8 = 42
    if true {
        var v7 int8 = 0
        _ = v7
    } else {
        switch p0 {
        case "foo bar":
            if true {
                p0 = p0
            } else {
                var v8 int32 = -603
                _ = v8
            }
        default:
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := []int{fi9, (fi9 * (fi9 & (fi9 >> ((0) & 63)))), fi9}
                _ = fi9
                _ = v10
            }
        }
        p0 = (("" + ("hello" + ((p0 + (p0 + ("" + (p0 + ("foo bar" + ""))))) + p0))) + (p0 + ""))
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            v12 := []int{fi11, fi11}
            _ = v12
        } else {
            var v13 uint8 = v6
            fmt.Println(fi11, v6)
            for fi14 := 0; fi14 < 18; fi14++ {
                v15 := []int{(fi14 | (fi14 / ((-1) | 1))), int(v6), 0, (-(fi11))}
                switch 261 {
                case -640:
                    var v16 uint = 972
                    _ = v16
                default:
                    fmt.Println(p0, fi11)
                }
                var v17 float32 = 3.14159
                v17 = (0.0 - 1.5)
                var v18 uint32 = 2
                v18 = v18
                v19 := GenS1{F0: 730}
                _ = fi14
                _ = v15
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v13
        }
        _ = fi11
    }
    v20 := []int{}
    _ = p0
    _ = v20
    _ = v6
    return 0
}
func genFn1() int {
    var v21 uint16 = 65535
    if (true && (v21 <= v21)) {
        if (42 >= 42) {
            for fi22 := 0; fi22 < 2; fi22++ {
                switch 128 {
                case 0:
                    v21 = v21
                case -66:
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v21 = (v21 >> ((0) & 63))
                        fmt.Println(fi23, fi23, fi23)
                        v24 := []int{(fi23 * fi23)}
                        _ = fi23
                        _ = v24
                    }
                }
                v21 += (v21 + (^(v21)))
                v21 = 2
                _ = fi22
            }
            switch (("foo bar" + "hello") + ("a" + ("foo bar" + ("a" + ("a" + "\n\t"))))) {
            case "a":
                v25 := GenS1{F0: 255}
                _ = v25
            case "\n\t":
                v26 := []int{-2147483648, -1}
                _ = v26
            }
        }
        v27 := []int{}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 11; fi28++ {
            v21 += 973
            _ = fi28
        }
        v29 := []int{32768, 1}
        if (2147483647 <= 165) {
            v21 = v21
            var v30 uint32 = 547
            var v31 uint = 255
            _ = v30
            _ = v31
        }
        _ = v29
    }
    v21 -= v21
    _ = v21
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    v5 := []int{4294967295, 9223372036854775807}
    var v32 int32 = -543
    switch 128 {
    case 127:
        v33 := GenS2{F0: 255, F1: 0}
        _ = v33
    case -113:
        v34 := v4[8]
        _ = v34
    case 32767:
        v5 = append(v5, 42)
    default:
        for fi35 := 0; fi35 < 20; fi35++ {
            for fi36 := 0; fi36 < 11; fi36++ {
                fmt.Println(v32, fi36)
                v32 -= (v32 | 0)
                _ = fi36
            }
            _ = fi35
        }
    }
    v32 -= v32
    var v37 string = "a"
    v1[524] = genFn1()
    var v38 int16 = 0
    fmt.Println(v32)
    _ = len(v4)
    switch genFn1() {
    case -1:
        for fi39 := 0; fi39 < 16; fi39++ {
            if len(v5) > 0 {
                v5[0] = (fi39 - (fi39 >> ((255) & 63)))
            }
            v37 = ((v37 + ("foo bar" + v37)) + (v37 + v37))
            fmt.Println(fi39, fi39)
            _ = fi39
        }
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        if true {
            for fi41 := 0; fi41 < 5; fi41++ {
                if (false && (!(929 > 2147483647))) {
                    var v42 uint32 = 1
                    _ = v42
                }
                if (4294967295 != 110) {
                    var v43 float32 = (0.0 - ((0.0 - (float32((0.1 - 0.1)) - (((0.0 - (-2.25 + float32(3.14159))) + (((3.14159 - (1.5 - 0.1)) * 1e10) + 0.1)) / 1e10))) / (1e10 - 0.0)))
                    var v44 int = fi40
                    if true {
                        v43 += v43
                    } else {
                        var v45 int = 52
                        _ = v45
                    }
                    _ = v43
                    _ = v44
                }
                v37 = v37
                v5 = append(v5, genFn1())
                _ = fi41
            }
        }
        _ = fi40
    }
    v38 -= v38
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
