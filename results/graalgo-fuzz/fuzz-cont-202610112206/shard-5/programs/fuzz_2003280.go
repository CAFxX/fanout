package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0(p0 string) string {
    v5 := []int{}
    p0 = p0
    var v6 int = 0
    var v7 uint32 = 2
    if ((^(v6)) != v6) {
        for fi8 := 0; fi8 < 1; fi8++ {
            if len(v5) > 0 {
                v5[0] = (fi8 + 834)
            }
            v7 -= v7
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi8
        }
        if (!false) {
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 float32 = 0.1
                    _ = fi10
                    _ = v11
                }
                if len(v5) > 0 {
                    v5[0] = 544
                }
                v12 := GenS1{F0: (v7 << ((2) & 63))}
                _ = fi9
                _ = v12
            }
            switch p0 {
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                v6 -= (v6 | v6)
            }
        } else {
            v5 = append(v5, v6)
            v13 := GenS1{F0: 186}
            _ = v13
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                switch (p0 + p0) {
                case "\n\t":
                    var v15 int8 = -31
                    _ = v15
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    fmt.Println(fi14, fi14)
                }
            } else {
                _ = len(v5)
                switch fi14 {
                case 4294967295:
                    v5 = append(v5, -1)
                case 162:
                    fmt.Println(v6, p0, p0)
                default:
                    v6 -= (v6 - v6)
                }
                p0 = p0
            }
            p0 = p0
            _ = fi14
        }
    } else {
        _ = len(v5)
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint8) bool {
    if false {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0)
        case "\n\t":
            p0 += p0
        default:
            p0 -= (p0 % (((p0 | p0)) | 1))
        }
    } else {
        var v16 bool = (975 < 4294967295)
        _ = v16
    }
    fmt.Println(p0)
    _ = p0
    return true
}

func main() {
    v1 := []int{4294967295, 0, 0}
    var v2 uint64 = 2
    var v3 uint64 = v2
    v4 := make(map[int]bool)
    v3 = v2
    var v17 int32 = 42
    var v18 uint = 65535
    v3 += v3
    v17 = (v17 - (v17 + 0))
    var v19 float32 = (-2.25 / 3.14159)
    var v20 string = "foo bar"
    v20 = v20
    for fi21 := 0; fi21 < 17; fi21++ {
        if false {
            if false {
                if (!true) {
                    v22 := GenS1{F0: 2}
                    var v23 bool = (!(!(42 == uint32(v18))))
                    _ = v22
                    _ = v23
                }
                _ = len(v4)
                v24 := v4[fi21]
                _ = v24
            } else {
                if true {
                    switch (fi21 * (fi21 + (fi21 + (-(fi21))))) {
                    case 104:
                        fmt.Println(v2)
                    case 42:
                        if (!false) {
                            var v25 int8 = 0
                            _ = v25
                        }
                    case 0:
                        _ = len(v4)
                    }
                }
                if false {
                    var v26 bool = false
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v28 := GenS1{F0: 65535}
                        v29 := GenS2{F0: 1, F1: 255}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    if v26 {
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 int32 = v17
                            var v32 uint8 = 127
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        fmt.Println(v3)
                        v20 = (v20 + ((v20 + "\n\t") + v20))
                    }
                    _ = v26
                } else {
                    if (-647 == 32767) {
                        v33 := v4[(fi21 / ((-9223372036854775808) | 1))]
                        for fi34 := 0; fi34 < 18; fi34++ {
                            v35 := GenS2{F0: 0, F1: 1}
                            v19 += 1.5
                            _ = fi34
                            _ = v35
                        }
                        _ = v33
                    } else {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            fmt.Println(v3, fi36)
                            v17 -= (v17 | 42)
                            v37 := GenS1{F0: 4294967295}
                            _ = fi36
                            _ = v37
                        }
                        var v38 uint8 = 0
                        _ = v38
                    }
                }
                switch "a" {
                case "foo bar":
                    var v39 uint8 = 230
                    _ = v39
                case "\n\t":
                    var v40 int32 = v17
                    _ = v40
                default:
                    v18 += (v18 << ((v18) & 63))
                }
            }
        } else {
            fmt.Println(v3, v2, v18)
            v18 = v18
        }
        var v41 int = fi21
        _ = fi21
        _ = v41
    }
    var v42 uint16 = 255
    v43 := GenS3{F0: 42, F1: 127}
    var v44 int = -1
    v17 = v17
    if (1 != 1) {
        var v45 string = ("a" + v20)
        var v46 float32 = (v19 + float32((0.1 / 1e10)))
        _ = v45
        _ = v46
    }
    switch v44 {
    case 127:
        v2 -= 18446744073709551615
    case -1:
        v4[v44] = (true && true)
    case 1:
        v47 := GenS3{F0: 146, F1: 255}
        _ = v47
    default:
        v20 = v20
    }
    var v48 float64 = (0.1 - ((float64(v19) * (float64((float32((3.14159 + float64(v19))) - v19)) + 1.5)) + 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v48
}
