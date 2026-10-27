package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8) uint8 {
    p1 = (p1 >> ((65535) & 63))
    v5 := []int{32767, -2147483648, int(p1), 1}
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = (1.5 / float32(float64((float32(0.1) * float32(float64(float32(1e10)))))))
        v8 += ((((((3.14159 * 3.14159) + v8) * float32(3.14159)) / (0.1 - (v8 + (0.1 - v8)))) + (0.1 + v8)) - v8)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 ^ 127)
}

func main() {
    v1 := []int{1, 4294967295, 32767, -736}
    var v2 uint64 = 423
    v3 := []int{0, 32768, -620}
    v4 := make(map[string]string)
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v4, "")
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 5; fi10++ {
                _ = len(v1)
                var v11 int = fi10
                if true {
                    v11 += (^(fi10))
                    v12 := GenS1{F0: 2, F1: 1}
                    _ = v12
                } else {
                    var v13 uint32 = 434
                    var v14 string = "foo bar"
                    var v15 int64 = -768
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    _ = len(v1)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint16 = 1
        switch fi16 {
        case -297:
            var v18 int16 = 32767
            _ = v18
        case 32768:
            if ((1 != -1) || true) {
                v19 := GenS1{F0: 2, F1: (v17 + 2)}
                if false {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if ((^(v17)) != (v17 - v17)) {
                            v2 = v2
                        }
                        var v21 uint64 = (v2 - (^(v2)))
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS1{F0: uint(v17), F1: v17}
                    switch fi16 {
                    case 128:
                        _ = len(v1)
                    case 904:
                        switch fi16 {
                        case 1:
                            if false {
                                var v23 int = (^(fi16))
                                v24 := GenS2{F0: (v23 / (((v23 ^ fi16)) | 1))}
                                _ = v23
                                _ = v24
                            }
                        default:
                            v4["hello"] = "hello"
                        }
                    case 42:
                        v25 := GenS3{F0: 1, F1: 42, F2: ""}
                        _ = v25
                    }
                    _ = v22
                } else {
                    switch 159 {
                    case 9223372036854775807:
                        var v26 uint32 = 200
                        _ = v26
                    case -695:
                        if false {
                            v27 := GenS3{F0: v2, F1: 0, F2: "foo bar"}
                            v2 = 42
                            _ = len(v4)
                            _ = v27
                        } else {
                            v1 = append(v1, fi16)
                            switch "foo bar" {
                            case "a":
                                v28 := GenS2{F0: fi16}
                                _ = v28
                            case "":
                                switch ("foo bar" + ((("a" + "foo bar") + "foo bar") + "foo bar")) {
                                case "":
                                    v17 += v17
                                case "a":
                                    v29 := GenS1{F0: 2, F1: v17}
                                    _ = v29
                                case "hello":
                                    var v30 int = (fi16 - fi16)
                                    _ = v30
                                default:
                                    delete(v4, "\n\t")
                                }
                            case "hello":
                                var v31 int16 = 1
                                _ = v31
                            default:
                                v2 -= 2
                            }
                        }
                    case 1:
                        var v32 float32 = -2.25
                        _ = v32
                    }
                }
                _ = v19
            }
        }
        _ = fi16
        _ = v17
    }
    switch 127 {
    case -1:
        v2 = (^(v2))
    case -2147483648:
        delete(v4, "hello")
    case -154:
        v1 = append(v1, -1)
    default:
        var v33 float64 = 3.14159
        _ = v33
    }
    v34 := GenS4{F0: true}
    fmt.Println(v2)
    if true {
        var v35 bool = false
        if false {
            var v36 float32 = ((float32(1.5) / 3.14159) + 1e10)
            if len(v1) > 0 {
                v1[0] = 127
            }
            _ = v36
        }
        _ = v35
    } else {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 string = "\n\t"
            _ = fi37
            _ = v38
        }
        fmt.Println(v2)
    }
    v2 += v2
    for fi39 := 0; fi39 < 3; fi39++ {
        v2 = (v2 | (v2 - (v2 >> ((1) & 63))))
        var v40 int8 = -116
        _ = fi39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
