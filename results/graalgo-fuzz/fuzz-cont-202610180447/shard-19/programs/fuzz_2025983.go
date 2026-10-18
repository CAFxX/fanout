package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int16; F1 int8; F2 bool }

func main() {
    v1 := []int{436, 42}
    var v2 uint64 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    if false {
        if (518 < 2147483648) {
            _ = len(v1)
            fmt.Println(v2)
            var v5 int16 = 371
            _ = v5
        }
    }
    if len(v1) > 0 {
        v1[0] = -194
    }
    _ = len(v1)
    v2 = v2
    v6 := GenS1{F0: 906}
    v2 -= v2
    v2 = v2
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 1
        _ = fi7
        _ = v8
    }
    var v9 float32 = 1e10
    var v10 uint32 = 2
    v2 = v2
    if false {
        for fi11 := 0; fi11 < 8; fi11++ {
            switch ("" + "a") {
            case "hello":
                fmt.Println(v9, fi11)
            case "foo bar":
                switch (^(fi11)) {
                case 1:
                    delete(v4, "hello")
                default:
                    if (42 <= 2147483647) {
                        var v12 uint8 = 255
                        var v13 int32 = -2147483648
                        v13 += (v13 % ((1) | 1))
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(v2, v10, v2)
                        var v14 int16 = 488
                        _ = v14
                    }
                }
            default:
                v15 := GenS1{F0: 524}
                _ = v15
            }
            fmt.Println(v10, v9)
            _ = fi11
        }
        fmt.Println(v2, v9, v2)
        for fi16 := 0; fi16 < 18; fi16++ {
            switch ((("a" + "foo bar") + "a") + "foo bar") {
            case "foo bar":
                if (1 != 127) {
                    fmt.Println(v9, v9, v10)
                }
            case "":
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS1{F0: 2}
                    fmt.Println(v2)
                    _ = fi17
                    _ = v18
                }
            case "hello":
                var v19 int16 = -424
                _ = v19
            }
            var v20 uint32 = uint32(v2)
            _ = fi16
            _ = v20
        }
    } else {
        if (uint16(v2) < 0) {
            fmt.Println(v2, v2)
        } else {
            if (!true) {
                v2 -= (v2 - (v2 / ((4294967295) | 1)))
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 int8 = -128
                    v22 += (v22 - (^(v22)))
                    _ = fi21
                    _ = v22
                }
                switch ("\n\t" + "") {
                case "":
                    v10 = v10
                case "hello":
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        if true {
                            v10 = v10
                            v23 := GenS2{F0: 2, F1: 0}
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 65535, F1: 119, F2: 0}
                            v25 := GenS4{F0: -809, F1: 4294967295}
                            v26 := GenS5{F0: 42, F1: 42}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                default:
                    var v27 bool = false
                    _ = v27
                }
            }
        }
    }
    var v28 float32 = 1e10
    var v29 int64 = 1
    var v30 uint = 42
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            v4["\n\t"] = true
            _ = fi32
        }
        v30 = (^(v30))
        for fi33 := 0; fi33 < 8; fi33++ {
            v28 = float32(((float64(v9) - (float64(0.1) * 1e10)) - (1e10 - (((3.14159 * 3.14159) + (0.1 / 1.0)) * float64(v28)))))
            v34 := GenS6{F0: 42, F1: -13, F2: false}
            v35 := GenS4{F0: 42, F1: (v30 * v30)}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v6
    _ = v9
}
