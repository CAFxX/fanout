package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8) string {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{(fi7 ^ 4294967295)}
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    var v9 int16 = 391
    p0 -= (p0 << ((156) & 63))
    var v10 float64 = 0.0
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v9
    return ("hello" + ("\n\t" + "hello"))
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 4294967295
    var v3 float32 = 0.1
    v4 := make(map[string]uint32)
    var v5 int64 = -1
    v6 := []uint32{2, 93}
    for fi11 := 0; fi11 < 2; fi11++ {
        v2 = v2
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    var v13 float64 = 3.14159
    fmt.Println(v5, v13)
    switch -839 {
    case 4294967295:
        if false {
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v3)
                _ = fi14
            }
            if false {
                _ = len(v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v15 float64 = v13
            _ = v15
        }
    case 554:
        if (!false) {
            if false {
                var v16 uint16 = 1
                if (!false) {
                    fmt.Println(v2)
                    var v17 uint8 = 2
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        if (fi18 > (fi18 + (fi18 + fi18))) {
                            switch "" {
                            case "a":
                                v5 += (v5 - (v5 - (v5 + (v5 * int64(v16)))))
                            case "foo bar":
                                for fi19 := 0; fi19 < 7; fi19++ {
                                    v13 += v13
                                    if true {
                                        switch fi18 {
                                        case 407:
                                            fmt.Println(fi19)
                                        case 32768:
                                            v16 = (v16 ^ v16)
                                        }
                                        v20 := GenS1{F0: -522, F1: true, F2: 911}
                                        _ = v20
                                    } else {
                                        switch fi18 {
                                        case 42:
                                            fmt.Println(v5, fi18)
                                        case -2147483648:
                                            var v21 float64 = v13
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: 1, F1: false, F2: 4294967295}
                                        _ = v22
                                    }
                                    _ = fi19
                                }
                            }
                        } else {
                            v16 = v16
                            v2 -= (v2 - v2)
                            v23 := GenS2{F0: (-(v5)), F1: "\n\t", F2: 127}
                            _ = v23
                        }
                        _ = fi18
                    }
                    v4[""] = 0
                }
                _ = v16
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println(v13)
    var v24 int32 = -1
    var v25 uint8 = 127
    v26 := GenS3{F0: 1, F1: 305}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
