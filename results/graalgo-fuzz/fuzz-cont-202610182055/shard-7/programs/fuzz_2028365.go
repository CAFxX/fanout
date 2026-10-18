package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64, p2 int8) int32 {
    if false {
        switch "hello" {
        case "":
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint = uint(p1)
                _ = fi7
                _ = v8
            }
        case "foo bar":
            for fi9 := 0; fi9 < 7; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    if false {
                        fmt.Println(p2, p1, p2)
                    }
                    switch "" {
                    case "":
                        var v11 uint32 = 2147483648
                        _ = v11
                    case "hello":
                        var v12 uint64 = p0
                        _ = v12
                    default:
                        p2 = p2
                    }
                    switch 127 {
                    case -9223372036854775808:
                        var v13 uint32 = 2147483648
                        _ = v13
                    case -716:
                        var v14 float32 = (-2.25 + 0.0)
                        _ = v14
                    case 9223372036854775807:
                        p2 = (p2 % ((p2) | 1))
                    }
                    _ = fi10
                }
                if false {
                    p1 = 116
                    v15 := []int{(fi9 >> ((255) & 63)), fi9, (fi9 / ((int(p2)) | 1))}
                    if (((3.14159 * (-2.25 * (((float32((0.1 + 3.14159)) / 1e10) * (-2.25 - 0.0)) - float32(-2.25)))) - 3.14159) != 0.1) {
                        _ = len(v15)
                    } else {
                        if true {
                            p0 = p0
                            v15 = append(v15, fi9)
                        } else {
                            for fi16 := 0; fi16 < 4; fi16++ {
                                p1 -= 1
                                p2 = 0
                                var v17 uint64 = (p1 + p1)
                                _ = fi16
                                _ = v17
                            }
                            var v18 int32 = -108
                            var v19 uint8 = 141
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = v15
                } else {
                    if false {
                        v20 := []int{fi9, (fi9 ^ fi9), (fi9 & (fi9 & fi9))}
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 4; fi21++ {
                        var v22 int64 = -1
                        var v23 int8 = p2
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    p1 += 0
                }
                _ = fi9
            }
        }
        var v24 int32 = -910
        switch "a" {
        case "":
            p1 = 18446744073709551615
        case "foo bar":
            v24 = -216
        case "hello":
            v25 := GenS1{F0: 1, F1: 42}
            _ = v25
        default:
            p1 -= p0
        }
        _ = v24
    }
    var v26 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 184
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 int = 0
    var v6 int = v5
    for fi27 := 0; fi27 < 18; fi27++ {
        v6 += (v5 | v5)
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v6, v2)
            _ = fi28
        }
        _ = fi27
    }
    var v29 uint = 255
    fmt.Println(v2)
    fmt.Println(v5)
    _ = len(v4)
    var v30 int32 = -2147483648
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 string = "a"
        if (276 != (v30 ^ v30)) {
            v33 := []int{int(v29)}
            _ = v33
        } else {
            v34 := GenS1{F0: 78, F1: 1}
            _ = v34
        }
        v35 := []int{}
        _ = fi31
        _ = v32
        _ = v35
    }
    var v36 int64 = 0
    if (!false) {
        v37 := GenS2{F0: 1}
        v2 = 2147483647
        v38 := GenS1{F0: 32767, F1: 171}
        _ = v37
        _ = v38
    }
    v39 := GenS1{F0: 1000, F1: v36}
    var v40 uint16 = 255
    _ = len(v1)
    for fi41 := 0; fi41 < 17; fi41++ {
        if (!v3) {
            v2 = (v2 * v2)
            v5 -= (^(fi41))
            v42 := []int{(fi41 - (v6 & (v5 % ((v5) | 1)))), (^(v6)), v6}
            _ = v42
        } else {
            v40 += 65535
        }
        fmt.Println(v36, v3)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
