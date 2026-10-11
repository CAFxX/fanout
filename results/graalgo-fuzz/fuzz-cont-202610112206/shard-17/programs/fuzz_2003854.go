package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint32 = 65535
    var v2 int32 = 0
    v3 := make(map[string]int)
    if false {
        if true {
            v1 = (^(v1))
        }
        _ = len(v3)
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 uint32 = (^(v1))
            v5 += v5
            switch "\n\t" {
            case "a":
                var v6 string = "hello"
                _ = v6
            case "hello":
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 uint64 = 42
                    var v9 uint32 = (v1 * v5)
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            default:
                fmt.Println(fi4, v2)
            }
            _ = fi4
            _ = v5
        }
    }
    switch 32767 {
    case -2147483648:
        v10 := GenS1{F0: 4294967295}
        _ = v10
    case 128:
        var v11 bool = false
        _ = v11
    }
    _ = len(v3)
    v12 := []int{32768}
    v2 = v2
    var v13 uint = 1
    var v14 string = "hello"
    v15 := GenS1{F0: v1}
    switch -14 {
    case -186:
        v16 := v3[v14]
        _ = v16
    case -1:
        var v17 float32 = 0.1
        _ = v17
    case -212:
        v2 += (-(v2))
    default:
        v13 += v13
    }
    fmt.Println(v14)
    var v18 uint16 = 2
    var v19 uint64 = 1
    var v20 string = v14
    if true {
        var v21 uint64 = v19
        _ = v21
    } else {
        if (false && false) {
            fmt.Println(v1, v18, v1)
            if (v13 != v13) {
                if false {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := GenS2{F0: (^(v1)), F1: (fi22 << ((v13) & 63))}
                        _ = fi22
                        _ = v23
                    }
                }
                v24 := GenS3{F0: 32767, F1: 45}
                for fi25 := 0; fi25 < 4; fi25++ {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        fmt.Println(v13, v2, fi26)
                        if len(v12) > 0 {
                            v12[0] = 2147483647
                        }
                        _ = fi26
                    }
                    switch (^(fi25)) {
                    case 128:
                        var v27 int = fi25
                        _ = v27
                    case 310:
                        v28 := GenS3{F0: 32767, F1: -569}
                        _ = v28
                    default:
                        var v29 uint = (v13 & (^(v13)))
                        _ = v29
                    }
                    var v30 int32 = (v2 + -1)
                    _ = fi25
                    _ = v30
                }
                _ = v24
            } else {
                for fi31 := 0; fi31 < 7; fi31++ {
                    if true {
                        var v32 int64 = -9223372036854775808
                        _ = v32
                    } else {
                        var v33 int = fi31
                        v34 := GenS4{F0: -9223372036854775808}
                        if false {
                            switch 2147483647 {
                            case -874:
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    v33 += fi31
                                    v1 -= v1
                                    _ = fi35
                                }
                            case 9223372036854775807:
                                v12 = append(v12, (fi31 * (v33 >> (((v13 * 578)) & 63))))
                            default:
                                switch ("foo bar" + v14) {
                                case "foo bar":
                                    var v36 uint = (v13 - (^(v13)))
                                    _ = v36
                                }
                            }
                            var v37 int16 = 664
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 16; fi38++ {
                                v39 := v3[((((v14 + v14) + "\n\t") + v14) + (((v20 + "hello") + (v20 + "a")) + "a"))]
                                fmt.Println(v18)
                                fmt.Println(v14)
                                v33 -= v39
                                var v40 uint = (v13 >> ((v13) & 63))
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                            var v41 uint32 = (v1 & (^(v1)))
                            if (-128 >= int8(v2)) {
                                var v42 bool = (!false)
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    _ = len(v3)
                                    _ = fi43
                                }
                                _ = len(v12)
                                _ = v42
                            } else {
                                var v44 uint = v13
                                if true {
                                    var v45 float64 = float64(float32((float64(1.5) + 0.0)))
                                    _ = v45
                                }
                                fmt.Println(v13)
                                _ = v44
                            }
                            _ = v41
                        }
                        _ = v33
                        _ = v34
                    }
                    v13 += v13
                    v19 = (v19 * v19)
                    _ = fi31
                }
            }
            v14 = v20
        } else {
            var v46 float32 = (0.0 - 3.14159)
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
