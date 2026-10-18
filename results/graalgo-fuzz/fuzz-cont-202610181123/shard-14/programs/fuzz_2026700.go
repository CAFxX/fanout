package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int; F2 uint }
func genFn0() bool {
    var v7 int = -421
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = 195
        v9 = (v9 << ((4294967295) & 63))
        _ = fi8
        _ = v9
    }
    v10 := []int{(^(v7)), (^(v7))}
    var v11 float32 = (((0.1 - 1.5) + (3.14159 / 3.14159)) * ((3.14159 - 0.1) * 3.14159))
    _ = v10
    _ = v11
    _ = v7
    return true
}
func genFn1(p0 string, p1 int32, p2 int8) bool {
    fmt.Println(p1, p0, p2)
    var v12 int16 = -1
    v13 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return true
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 float32 = (1.5 * ((((0.0 / -2.25) * (3.14159 + float32(1e10))) - ((1.5 / -2.25) - 3.14159)) - float32(1e10)))
    var v4 int = -9223372036854775808
    v5 := []string{"a", "a", "foo bar", ""}
    var v6 int8 = -128
    var v14 uint8 = 194
    for fi15 := 0; fi15 < 4; fi15++ {
        _ = len(v5)
        switch int(v14) {
        case 304:
            for fi16 := 0; fi16 < 19; fi16++ {
                if (!(9223372036854775807 < 1)) {
                    if true {
                        var v17 int16 = -32768
                        _ = v17
                    } else {
                        v5 = append(v5, "foo bar")
                        v5 = append(v5, "a")
                    }
                    var v18 uint = 2
                    _ = v18
                }
                _ = fi16
            }
        case 677:
            var v19 string = "foo bar"
            _ = v19
        }
        _ = fi15
    }
    v20 := GenS1{F0: 167}
    v21 := GenS2{F0: 1000}
    var v22 string = ("a" + "\n\t")
    if true {
        if (!(v4 < (v4 % ((127) | 1)))) {
            var v23 uint64 = 18446744073709551615
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := v2[("a" + (((v22 + ("a" + ("" + "a"))) + v22) + "foo bar"))]
                _ = fi24
                _ = v25
            }
            if true {
                v5 = append(v5, v22)
                fmt.Println(v14)
                v3 -= float32((0.0 * (0.0 + 1e10)))
            }
            _ = v23
        } else {
            v14 -= (v14 << ((803) & 63))
        }
        if false {
            _ = len(v5)
            var v26 uint32 = 0
            _ = v26
        }
        var v27 float64 = (0.1 + (1.5 - -2.25))
        _ = v27
    }
    var v28 uint64 = uint64(v14)
    if false {
        v14 -= 128
        v29 := GenS1{F0: v28}
        _ = v29
    } else {
        if false {
            switch v4 {
            case -1:
                v30 := GenS3{F0: false}
                _ = v30
            case 0:
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 float32 = (v3 - v3)
                    v33 := GenS4{F0: 29, F1: fi31, F2: 384}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            default:
                if false {
                    if false {
                        switch v22 {
                        case "a":
                            for fi34 := 0; fi34 < 8; fi34++ {
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    if false {
                                        v3 = v3
                                    } else {
                                        switch v4 {
                                        case -517:
                                            v22 = (v22 + v22)
                                        case 32768:
                                            fmt.Println(v3, fi34, v22)
                                            v36 := v2["\n\t"]
                                            _ = v36
                                        }
                                        v37 := GenS2{F0: -1}
                                        _ = v37
                                    }
                                    _ = fi35
                                }
                                _ = fi34
                            }
                        }
                        _ = len(v5)
                        var v38 int8 = (v6 << ((2) & 63))
                        _ = v38
                    } else {
                        var v39 bool = false
                        _ = v39
                    }
                }
            }
            switch v22 {
            case "hello":
                var v40 uint8 = v14
                _ = v40
            case "a":
                v14 = (v14 ^ (v14 * v14))
            case "foo bar":
                if true {
                    delete(v1, 32768)
                    var v41 string = (v22 + v22)
                    fmt.Println(v3, v28, v22)
                    _ = v41
                }
            }
        } else {
            v4 -= v4
        }
        v4 -= (-(v4))
    }
    var v42 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
