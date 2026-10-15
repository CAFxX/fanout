package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool, p1 int) bool {
    if (18446744073709551615 > 1) {
        if p0 {
            fmt.Println(p1)
            var v5 uint16 = 255
            _ = v5
        }
        var v6 uint16 = 0
        var v7 uint32 = 2147483648
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 1000}
        v9 := []int{p1, 1, (p1 - p1)}
        _ = v8
        _ = v9
    }
    p0 = p0
    fmt.Println(p0)
    p0 = false
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := []int{989, 1, 9223372036854775807, 4294967295}
    v2 := make(map[string]int)
    v3 := []int{0, 2147483647, 0}
    var v4 uint16 = 0
    if false {
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(fi10, v4)
            v11 := GenS1{F0: 32767}
            _ = fi10
            _ = v11
        }
        var v12 uint64 = 42
        _ = v12
    } else {
        switch (("\n\t" + "a") + "\n\t") {
        case "hello":
            var v13 uint64 = 996
            _ = v13
        case "\n\t":
            fmt.Println(v4)
        case "":
            v4 -= v4
        default:
            _ = len(v2)
        }
        v4 -= (v4 & (v4 & 328))
        v4 += (v4 | 42)
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = 874
        v4 -= (v4 - v4)
        fmt.Println(v15, fi14)
        _ = fi14
        _ = v15
    }
    var v16 int16 = 1
    v16 = v16
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float32 = -2.25
        delete(v2, ("" + "foo bar"))
        _ = fi17
        _ = v18
    }
    if (v4 == (v4 / ((v4) | 1))) {
        var v19 int = 835
        var v20 string = ""
        fmt.Println(v19, v16, v16)
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            v2["hello"] = fi21
            _ = fi21
        }
        switch "foo bar" {
        case "":
            v4 = 2
        case "foo bar":
            for fi22 := 0; fi22 < 19; fi22++ {
                v3 = append(v3, (^(fi22)))
                fmt.Println(v16)
                if (-272 > 42) {
                    v23 := GenS1{F0: (v16 / ((24) | 1))}
                    var v24 uint64 = 2
                    _ = v23
                    _ = v24
                } else {
                    v4 = v4
                    for fi25 := 0; fi25 < 11; fi25++ {
                        if true {
                            var v26 float64 = 3.14159
                            _ = v26
                        } else {
                            v16 -= (v16 % ((v16) | 1))
                            if len(v3) > 0 {
                                v3[0] = 32767
                            }
                        }
                        fmt.Println(fi22)
                        if true {
                            _ = len(v2)
                            var v27 int32 = 588
                            _ = v27
                        } else {
                            switch fi25 {
                            case 4294967295:
                                switch (fi22 ^ 32767) {
                                case -2147483648:
                                    var v28 int8 = 113
                                    _ = v28
                                default:
                                    v4 = v4
                                }
                            default:
                                if true {
                                    v4 = (^(v4))
                                    fmt.Println(v4, v4)
                                } else {
                                    var v29 int8 = 0
                                    _ = v29
                                }
                            }
                            v30 := GenS1{F0: v16}
                            v3 = append(v3, 42)
                            _ = v30
                        }
                        _ = fi25
                    }
                    var v31 uint8 = 127
                    _ = v31
                }
                _ = fi22
            }
        case "a":
            v32 := GenS1{F0: (v16 & v16)}
            _ = v32
        default:
            var v33 int = -9223372036854775808
            _ = v33
        }
    }
    if (0 != 112) {
        v16 = (v16 - -32768)
        v16 -= (^(v16))
    }
    v34 := GenS1{F0: v16}
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 uint = 659
        _ = fi35
        _ = v36
    }
    var v37 uint = 892
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
}
