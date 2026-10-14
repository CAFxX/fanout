package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
func genFn0() int32 {
    var v5 int32 = 2147483647
    fmt.Println(v5)
    var v6 int8 = -128
    v7 := []int{9223372036854775807, 9223372036854775807, -971}
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint16 {
    if ((false || (!(-1 == 100))) && (!(-1 != -1))) {
        var v8 uint8 = 0
        _ = v8
    }
    switch 32767 {
    case 128:
        var v9 uint8 = 127
        _ = v9
    default:
        switch "\n\t" {
        case "a":
            var v10 float64 = 1.5
            _ = v10
        case "":
            var v11 uint16 = 32767
            _ = v11
        case "\n\t":
            var v12 int32 = genFn0()
            _ = v12
        }
    }
    if (true || false) {
        var v13 int = 1
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 5; fi15++ {
                v13 = (-(fi15))
                _ = fi15
            }
            switch (v13 / ((-255) | 1)) {
            case 0:
                switch -25 {
                case 672:
                    v16 := GenS1{F0: 711}
                    _ = v16
                case -32:
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v18 := GenS1{F0: 200}
                        var v19 uint32 = 48
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            case 4294967295:
                switch (fi14 / ((fi14) | 1)) {
                case -9223372036854775808:
                    fmt.Println(fi14, v13)
                case -131:
                    var v20 int64 = 1
                    _ = v20
                default:
                    v21 := GenS2{F0: 639, F1: 192, F2: 255}
                    _ = v21
                }
            default:
                v13 = v13
            }
            var v22 int8 = 127
            _ = fi14
            _ = v22
        }
        _ = v13
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS2{F0: 42, F1: 65535, F2: 4294967295}
            var v26 uint16 = 65535
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 uint32 = 2147483648
        _ = fi23
        _ = v27
    }
    return 0
}

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    var v3 uint32 = 4294967295
    v4 := make(map[int]int)
    v3 += 0
    if true {
        v28 := GenS1{F0: 2}
        v2 = (-2.25 + (float64(0.1) * 1e10))
        _ = v28
    } else {
        var v29 uint64 = 0
        _ = v29
    }
    v2 = (v2 / float64(1e10))
    switch 32767 {
    case 4294967295:
        var v30 int32 = 664
        _ = v30
    case 42:
        fmt.Println(v2, v3)
    }
    var v31 string = ("a" + "")
    v32 := v4[174]
    for fi33 := 0; fi33 < 8; fi33++ {
        v31 = v31
        var v34 int64 = 9223372036854775807
        switch (v31 + v31) {
        case "hello":
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 float64 = v2
                switch ((v31 + (v31 + "a")) + v31) {
                case "\n\t":
                    var v37 int = (fi33 * v32)
                    _ = v37
                case "a":
                    v3 += v3
                }
                v36 -= (-2.25 / (float64(1.5) / v36))
                _ = fi35
                _ = v36
            }
        case "a":
            for fi38 := 0; fi38 < 4; fi38++ {
                _ = len(v1)
                _ = fi38
            }
        case "":
            var v39 uint = 18446744073709551615
            _ = v39
        }
        _ = fi33
        _ = v34
    }
    v32 = (v32 | (v32 & v32))
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 uint64 = 73
        if (128 < 42) {
            v32 += fi40
        } else {
            fmt.Println(fi40)
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v32)
    v1 = append(v1, v31)
    if len(v1) > 0 {
        v1[0] = ("" + ((v31 + (v31 + "foo bar")) + ("a" + v31)))
    }
    if (-2147483648 > -2147483648) {
        fmt.Println(v31)
        v32 = v32
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v42 := GenS1{F0: 2}
        _ = v42
    }
    switch (v31 + "a") {
    case "hello":
        for fi43 := 0; fi43 < 5; fi43++ {
            if false {
                if false {
                    v44 := GenS2{F0: genFn0(), F1: 42, F2: 42}
                    fmt.Println(v3, v32)
                    _ = v44
                }
                var v45 int64 = 0
                var v46 string = v31
                _ = v45
                _ = v46
            } else {
                switch ("\n\t" + v31) {
                case "a":
                    if (417 != v32) {
                        fmt.Println(v31, v2, v2)
                    } else {
                        v3 -= (v3 << ((892) & 63))
                        var v47 uint = 1
                        _ = v47
                    }
                }
            }
            v4[(v32 * (-(v32)))] = (-(fi43))
            fmt.Println(v31, v31, v31)
            _ = fi43
        }
    default:
        var v48 bool = (1 >= 0)
        _ = v48
    }
    v3 = (v3 | v3)
    var v49 int32 = -1
    var v50 int32 = v49
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v49
    _ = v50
}
