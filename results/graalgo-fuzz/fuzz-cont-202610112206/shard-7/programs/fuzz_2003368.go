package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 string; F2 int }
func genFn0() uint8 {
    v5 := GenS1{F0: 617, F1: 65535}
    var v6 uint = 255
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int) uint32 {
    p0 = p0
    v7 := []int{-1}
    fmt.Println(p0)
    var v8 bool = true
    var v9 uint8 = 128
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 int = 32768
    var v3 uint64 = 4294967295
    var v4 float64 = 3.14159
    v2 -= 232
    switch ((("a" + "") + "a") + "a") {
    case "a":
        if false {
            v10 := []int{v2, -2147483648, (^(v2)), (v2 >> ((18446744073709551615) & 63))}
            v11 := GenS1{F0: 557, F1: 2}
            _ = v10
            _ = v11
        }
    case "":
        v1 -= v1
    case "\n\t":
        v3 = 678
    default:
        v12 := []int{v2, 32767, (v2 % (((v2 ^ (v2 * v2))) | 1)), -2147483648}
        _ = v12
    }
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "foo bar":
        for fi13 := 0; fi13 < 18; fi13++ {
            v2 += (v2 >> ((339) & 63))
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS2{F0: 1, F1: 2147483648, F2: true}
                if (false || ((v2 | 4294967295) > fi14)) {
                    var v16 int = (fi13 + v2)
                    var v17 uint = 42
                    _ = v16
                    _ = v17
                } else {
                    if ((-2.25 - v4) < (v4 + (v4 - v4))) {
                        v18 := GenS2{F0: 873, F1: 901, F2: (-1 > 852)}
                        fmt.Println(fi14)
                        switch "" {
                        case "hello":
                            v19 := GenS3{F0: fi14, F1: (fi14 + (fi14 & fi13))}
                            _ = v19
                        case "a":
                            fmt.Println(v3, v1)
                        case "\n\t":
                            switch "a" {
                            case "":
                                v3 = (v3 + 172)
                            }
                        }
                        _ = v18
                    } else {
                        if false {
                            switch "" {
                            case "foo bar":
                                v4 = ((v4 / v4) * v4)
                            }
                        }
                    }
                }
                fmt.Println(v3, fi14)
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    v20 := GenS4{F0: -32768, F1: ("\n\t" + ("hello" + "\n\t")), F2: (^(v2))}
    if true {
        v3 = 2
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int8 = 42
            var v23 uint8 = genFn0()
            v4 = v4
            _ = fi21
            _ = v22
            _ = v23
        }
        if true {
            var v24 float64 = float64(1.5)
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS3{F0: -2147483648, F1: (-(v2))}
                switch fi25 {
                case 9223372036854775807:
                    v27 := GenS4{F0: 1000, F1: "a", F2: (fi25 % ((fi25) | 1))}
                    _ = v27
                case 128:
                    for fi28 := 0; fi28 < 15; fi28++ {
                        v4 = float64(((3.14159 * 3.14159) - (1.5 / float32(1.5))))
                        _ = fi28
                    }
                case -2147483648:
                    v2 = (fi25 + v2)
                default:
                    var v29 float32 = 1e10
                    _ = v29
                }
                v30 := []int{fi25, int(v3)}
                _ = fi25
                _ = v26
                _ = v30
            }
            _ = v24
        }
    } else {
        var v31 string = "foo bar"
        _ = v31
    }
    v32 := []int{(v2 & -1), 0, (^(v2)), v2}
    if (false || false) {
        v2 += v2
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
    }
    v3 += 824
    v3 += v3
    v3 = 0
    if len(v32) > 0 {
        v32[0] = -859
    }
    switch "hello" {
    case "a":
        _ = len(v32)
    }
    v2 = -9223372036854775808
    switch v2 {
    case 9223372036854775807:
        for fi33 := 0; fi33 < 20; fi33++ {
            if (false || false) {
                fmt.Println(v3, v4)
                var v34 int32 = -1
                for fi35 := 0; fi35 < 20; fi35++ {
                    v3 -= (v3 + v3)
                    v36 := GenS2{F0: 15, F1: 0, F2: (true || true)}
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                if (false && true) {
                    if true {
                        v2 += fi33
                        var v37 float32 = 3.14159
                        var v38 float32 = 3.14159
                        _ = v37
                        _ = v38
                    }
                } else {
                    fmt.Println(fi33, fi33)
                    v4 = v4
                }
                v39 := GenS3{F0: v2, F1: (-(v2))}
                _ = v39
            }
            _ = fi33
        }
    case 0:
        v1 = (^(v1))
    case -290:
        var v40 int8 = 42
        _ = v40
    default:
        for fi41 := 0; fi41 < 2; fi41++ {
            _ = len(v32)
            var v42 uint32 = 4294967295
            _ = fi41
            _ = v42
        }
    }
    v3 = 42
    var v43 uint64 = v3
    v44 := GenS4{F0: 169, F1: ("foo bar" + "\n\t"), F2: v2}
    var v45 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
