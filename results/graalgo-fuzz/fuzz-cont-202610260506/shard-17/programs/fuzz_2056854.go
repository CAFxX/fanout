package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    v4 := []int{1, 9223372036854775807, 4294967295}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 float64 = -2.25
        fmt.Println(v6, fi5)
        v6 = (v6 / 3.14159)
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 2
    _ = v4
    _ = v7
    return 1
}
func genFn1(p0 uint) uint {
    v8 := []int{0, 0, 1, -1}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 964
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{19, 4294967295, 2}
    var v3 int64 = 0
    v3 += genFn0()
    v9 := GenS1{F0: 1}
    if len(v2) > 0 {
        v2[0] = 884
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint64 = 1
                v1 = -2.25
                var v13 float64 = float64((3.14159 / (float32((0.1 * 1e10)) * ((v1 + v1) * (float32(1e10) + (v1 * 3.14159))))))
                switch fi10 {
                case 42:
                    v12 += (v12 / ((2) | 1))
                default:
                    var v14 float32 = 0.1
                    _ = v14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            if true {
                fmt.Println(v3, v1)
                v3 = genFn0()
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if (0 >= 143) {
                    if (1.5 < 0.0) {
                        v1 += v1
                    }
                    var v16 int16 = 1
                    v17 := GenS2{F0: 1}
                    _ = v16
                    _ = v17
                } else {
                    if false {
                        var v18 bool = (true && true)
                        if false {
                            v3 = (v3 ^ v3)
                        } else {
                            var v19 float64 = 0.1
                            var v20 uint64 = 42
                            switch "hello" {
                            case "a":
                                v21 := GenS3{F0: 4294967295}
                                _ = v21
                            default:
                                if (0 == 1) {
                                    fmt.Println(fi10)
                                    if len(v2) > 0 {
                                        v2[0] = 525
                                    }
                                    fmt.Println(fi10, v20)
                                }
                                v20 = v20
                                v3 = v3
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = v18
                    }
                    var v22 uint32 = 41
                    _ = v22
                }
            }
            switch ("hello" + (("a" + "foo bar") + ("\n\t" + ("hello" + "foo bar")))) {
            case "":
                v2 = append(v2, 4294967295)
            case "hello":
                v1 = v1
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    switch (fi10 / (((fi15 * fi15)) | 1)) {
                    case -308:
                        v2 = append(v2, 0)
                    default:
                        var v23 uint = 2
                        _ = v23
                    }
                default:
                    v1 = v1
                    v3 += v3
                    var v24 uint = 65535
                    _ = v24
                }
            default:
                fmt.Println(v1)
            }
            _ = fi15
        }
        _ = fi10
    }
    var v25 uint16 = 32767
    switch 760 {
    case 9223372036854775807:
        var v26 uint32 = 0
        _ = v26
    case -1:
        v25 = (v25 ^ 2)
    case 471:
        v3 = 42
    }
    var v27 string = ""
    var v28 float32 = (v1 / v1)
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v29 int16 = -914
    if (!(!(false && false))) {
        var v30 uint32 = 4294967295
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v9
}
