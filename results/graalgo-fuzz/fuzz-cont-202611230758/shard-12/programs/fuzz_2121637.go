package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool) int64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{(fi6 << ((255) & 63)), (-(fi6))}
        var v8 float32 = (0.0 - ((1.5 * float32(((0.1 + (0.1 + -2.25)) / 1.5))) / 1e10))
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = (true || true)
    _ = p0
    return -9223372036854775808
}
func genFn1(p0 int8, p1 uint8) uint {
    p0 = -1
    v9 := GenS1{F0: -18}
    switch 311 {
    case 269:
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(p1, p0)
            _ = fi10
        }
    default:
        if (!true) {
            var v11 uint64 = 42
            _ = v11
        } else {
            v12 := []int{32768, 42}
            var v13 float64 = 0.0
            _ = v12
            _ = v13
        }
    }
    fmt.Println(p0)
    v14 := GenS1{F0: -999}
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return 65535
}

func main() {
    v1 := []uint32{1, 2147483648, 4294967295}
    var v2 int32 = -1
    var v3 float32 = 0.0
    var v4 int64 = -1
    v5 := []uint32{156, 42, 2147483648}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 := GenS1{F0: 0}
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v2)
        v17 := GenS1{F0: (fi16 >> ((605) & 63))}
        v18 := GenS2{F0: 237}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v2)
    switch "" {
    case "hello":
        var v19 int32 = v2
        _ = v19
    case "a":
        v4 += v4
    case "":
        switch ("hello" + "hello") {
        case "":
            var v20 int16 = 42
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float64 = 3.14159
                _ = fi21
                _ = v22
            }
        case "\n\t":
            v1 = append(v1, 65535)
        }
    }
    var v23 float32 = 1.5
    if false {
        var v24 float64 = 0.1
        v23 = (v3 * 0.0)
        var v25 uint8 = 93
        _ = v24
        _ = v25
    } else {
        switch 127 {
        case 932:
            v26 := GenS3{F0: 212}
            _ = v26
        }
        switch -18 {
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 1:
            for fi27 := 0; fi27 < 5; fi27++ {
                for fi28 := 0; fi28 < 15; fi28++ {
                    switch ("a" + ("" + "a")) {
                    case "\n\t":
                        v4 += genFn0((365 >= 165))
                    case "a":
                        switch "a" {
                        case "a":
                            fmt.Println(fi28, v23, v2)
                        case "\n\t":
                            v23 = v3
                        case "":
                            if false {
                                v29 := GenS1{F0: fi27}
                                var v30 uint16 = 32767
                                v4 -= (^(v4))
                                _ = v29
                                _ = v30
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 2
                                }
                                switch fi28 {
                                case 1:
                                    v31 := GenS2{F0: 2147483648}
                                    _ = v31
                                }
                            }
                        }
                    }
                    switch fi28 {
                    case 273:
                        fmt.Println(fi28, fi27)
                    case 42:
                        var v32 string = "\n\t"
                        _ = v32
                    case 419:
                        v4 = 42
                        if true {
                            fmt.Println(v4)
                        } else {
                            var v33 string = "foo bar"
                            _ = v33
                        }
                        var v34 string = "foo bar"
                        _ = v34
                    }
                    _ = fi28
                }
                _ = fi27
            }
        case 32767:
            for fi35 := 0; fi35 < 10; fi35++ {
                v2 += v2
                switch "hello" {
                case "":
                    var v36 uint32 = 2
                    _ = v36
                }
                _ = fi35
            }
        }
        var v37 int8 = -81
        _ = v37
    }
    v3 -= (-2.25 - 0.1)
    v1 = append(v1, 453)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
