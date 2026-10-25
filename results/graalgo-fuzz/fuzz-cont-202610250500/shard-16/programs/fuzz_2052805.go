package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
func genFn0(p0 bool, p1 int, p2 uint8) uint64 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 148, F1: -128}
        v6 := []int{317, 1}
        p2 = 0
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p0, p0)
    v7 := []int{p1, 812, p1, (p1 >> ((853) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{776, 1, 674, 4294967295}
    v3 := []int{}
    var v8 bool = (1e10 == 1e10)
    v9 := v1["a"]
    switch ("" + ("hello" + "foo bar")) {
    case "foo bar":
        v8 = v8
    }
    var v10 uint = 1
    var v11 uint8 = 56
    var v12 uint32 = 65535
    for fi13 := 0; fi13 < 12; fi13++ {
        v8 = false
        if (42 < 1) {
            var v14 string = ""
            v1[(("\n\t" + v14) + "hello")] = (^(fi13))
            _ = v14
        } else {
            v12 = (v12 * (v12 ^ 2))
            var v15 uint16 = 32767
            _ = v15
        }
        switch 128 {
        case -459:
            if v8 {
                fmt.Println(v10)
                var v16 int = fi13
                var v17 uint8 = 190
                _ = v16
                _ = v17
            } else {
                var v18 bool = (868 <= 637)
                switch (v9 / ((fi13) | 1)) {
                case -9223372036854775808:
                    var v19 uint16 = 65535
                    _ = v19
                }
                for fi20 := 0; fi20 < 6; fi20++ {
                    if (-2.25 <= ((((((1e10 * 1e10) + (((((-2.25 * float32(0.0)) + -2.25) - 1e10) / 1.0) / 3.14159)) - 1.5) + ((3.14159 - 0.1) - ((float32((3.14159 + 0.1)) * (1.5 + 1e10)) + (-2.25 + -2.25)))) / ((0.0 - 3.14159) / 3.14159)) * 0.0)) {
                        if (((true || false) && ((fi13 & 0) < -987)) || (1 > 1)) {
                            var v21 float64 = 1e10
                            var v22 int32 = 0
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS1{F0: 255, F1: -128}
                        v18 = v18
                        v9 += (fi20 ^ fi13)
                        if v18 {
                            var v24 bool = (!true)
                            _ = v24
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v23
                    } else {
                        if ((^(v11)) >= v11) {
                            v11 = (v11 | v11)
                            v10 = 255
                        } else {
                            if (!(0 != 18446744073709551615)) {
                                if (42 >= -9223372036854775808) {
                                    v18 = (!("hello" == "\n\t"))
                                } else {
                                    var v25 float32 = 3.14159
                                    var v26 float32 = 3.14159
                                    fmt.Println(v26, v25, v26)
                                    _ = v25
                                    _ = v26
                                }
                                v8 = (!true)
                            } else {
                                v11 = (^(v11))
                            }
                        }
                        fmt.Println(fi13, v8, v11)
                        _ = len(v1)
                    }
                    _ = fi20
                }
                _ = v18
            }
        case 2147483647:
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 int64 = 42
                _ = fi27
                _ = v28
            }
        case 4294967295:
            var v29 uint64 = 961
            _ = v29
        }
        _ = fi13
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS1{F0: 32767, F1: -9}
        if (1.5 >= 0.0) {
            switch "hello" {
            case "":
                var v32 int64 = 0
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                if (1 < 295) {
                    switch "foo bar" {
                    case "\n\t":
                        v34 := GenS2{F0: v9, F1: v11}
                        _ = v34
                    case "a":
                        fmt.Println(v8, v10, fi30)
                    default:
                        var v35 float32 = 0.0
                        _ = v35
                    }
                    var v36 bool = v8
                    _ = v36
                } else {
                    v37 := GenS2{F0: 32768, F1: v11}
                    var v38 string = "a"
                    var v39 float64 = 1e10
                    _ = v37
                    _ = v38
                    _ = v39
                }
                var v40 float64 = -2.25
                _ = fi33
                _ = v40
            }
        } else {
            switch "\n\t" {
            case "hello":
                var v41 float64 = 0.1
                _ = v41
            case "\n\t":
                if v8 {
                    v1[""] = (v9 ^ (v9 << ((v10) & 63)))
                    if len(v2) > 0 {
                        v2[0] = (v12 + v12)
                    }
                    v42 := GenS3{F0: -1, F1: -760, F2: -1}
                    _ = v42
                }
            }
            var v43 int8 = 0
            fmt.Println(v12, v43)
            _ = v43
        }
        for fi44 := 0; fi44 < 19; fi44++ {
            v45 := v1[""]
            _ = fi44
            _ = v45
        }
        _ = fi30
        _ = v31
    }
    v10 += 0
    var v46 int8 = -1
    var v47 float32 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v8
    _ = v9
}
