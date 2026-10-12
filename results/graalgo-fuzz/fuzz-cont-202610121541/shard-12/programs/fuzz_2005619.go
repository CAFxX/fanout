package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int32 }
func genFn0() uint {
    v4 := []int{}
    var v5 int32 = -1
    v5 -= v5
    var v6 string = "foo bar"
    v5 -= (v5 % ((v5) | 1))
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int64, p1 int64, p2 int) uint8 {
    p0 = (p0 % ((p1) | 1))
    var v7 uint = 2
    var v8 bool = (!(510 <= int32(p0)))
    p0 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 163
}

func main() {
    v1 := []uint32{2147483647, 2147483648}
    var v2 string = ("foo bar" + "a")
    var v3 float32 = ((1.5 + (1e10 / -2.25)) + (((1.5 * (((3.14159 / 0.1) * float32(1e10)) + 0.1)) - (0.0 + ((0.0 + 3.14159) * (0.0 - (3.14159 * -2.25))))) / ((1.5 - (float32(1.5) * ((1e10 / 1.0) / 1.0))) / 1e10)))
    var v9 float32 = 0.0
    var v10 string = "foo bar"
    v10 = ("" + v2)
    var v11 uint8 = 42
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 100}
            var v14 int8 = 1
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS1{F0: 42}
                v9 -= 1e10
                fmt.Println(v14, v2, fi15)
                _ = fi15
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        v1 = append(v1, 388)
        v17 := GenS1{F0: 127}
        _ = v17
    } else {
        if (0 >= 0) {
            for fi18 := 0; fi18 < 8; fi18++ {
                fmt.Println(v11, fi18)
                if len(v1) > 0 {
                    v1[0] = 42
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    if (!(v11 < (v11 * (v11 * v11)))) {
                        v10 = v10
                    }
                    _ = fi19
                }
                _ = fi18
            }
            fmt.Println(v9, v2, v9)
        }
        if (true || true) {
            if false {
                v10 = v2
                v20 := GenS2{F0: 0, F1: 1}
                switch 32767 {
                case -990:
                    v21 := GenS1{F0: 127}
                    _ = v21
                case -303:
                    switch 0 {
                    case 2147483647:
                        fmt.Println(v2, v10, v10)
                    case 9223372036854775807:
                        v10 = "foo bar"
                    default:
                        if true {
                            v22 := GenS3{F0: 1000, F1: 42}
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            _ = v22
                        }
                    }
                case 284:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        if (0 > 895) {
                            v24 := GenS1{F0: -104}
                            _ = v24
                        } else {
                            var v25 int = (fi23 >> ((4294967295) & 63))
                            v26 := GenS3{F0: 42, F1: 411}
                            _ = v25
                            _ = v26
                        }
                        var v27 uint16 = 342
                        _ = fi23
                        _ = v27
                    }
                }
                _ = v20
            }
            switch v10 {
            case "a":
                if false {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        if (true || (0 != 42)) {
                            var v29 int16 = 42
                            v30 := GenS4{F0: 4294967295}
                            _ = v29
                            _ = v30
                        }
                        _ = fi28
                    }
                    v31 := GenS1{F0: 100}
                    v1 = append(v1, 65535)
                    _ = v31
                } else {
                    fmt.Println(v2)
                    var v32 bool = true
                    _ = v32
                }
            case "":
                for fi33 := 0; fi33 < 9; fi33++ {
                    v3 = v3
                    _ = fi33
                }
            }
            var v34 uint16 = 255
            _ = v34
        }
    }
    v3 = 3.14159
    var v35 int64 = 42
    v35 += 0
    switch 9223372036854775807 {
    case 506:
        _ = len(v1)
    case 128:
        var v36 int16 = 1000
        _ = v36
    default:
        v10 = v2
    }
    var v37 int8 = 100
    if false {
        switch ((((v10 + ("hello" + "\n\t")) + v10) + "\n\t") + v2) {
        case "a":
            v3 = (v3 / v9)
        case "":
            var v38 int64 = v35
            _ = v38
        case "foo bar":
            var v39 int64 = (v35 % ((-1) | 1))
            _ = v39
        default:
            fmt.Println(v35, v37, v9)
        }
        _ = len(v1)
    } else {
        for fi40 := 0; fi40 < 8; fi40++ {
            var v41 int = (fi40 & (-(fi40)))
            _ = fi40
            _ = v41
        }
        v42 := GenS2{F0: 934, F1: v35}
        for fi43 := 0; fi43 < 7; fi43++ {
            if false {
                v11 = genFn1(v35, (v35 | (v35 / ((v35) | 1))), fi43)
                v35 = (^(v35))
            } else {
                for fi44 := 0; fi44 < 16; fi44++ {
                    v11 += (v11 + (v11 * v11))
                    _ = fi44
                }
                for fi45 := 0; fi45 < 17; fi45++ {
                    switch (fi43 - (fi45 & fi45)) {
                    case -312:
                        v9 += float32((-2.25 + 1e10))
                    case -1:
                        v10 = (v2 + "foo bar")
                    default:
                        var v46 uint = 0
                        _ = v46
                    }
                    _ = fi45
                }
            }
            v47 := GenS5{F0: (((v10 + (("foo bar" + (((v10 + "") + "\n\t") + v2)) + (v10 + (v2 + v10)))) + v2) + (v10 + "a")), F1: 0}
            _ = fi43
            _ = v47
        }
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v9
}
