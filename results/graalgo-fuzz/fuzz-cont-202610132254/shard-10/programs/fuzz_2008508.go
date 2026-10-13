package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint }

func main() {
    var v1 bool = false
    v2 := []int{-9223372036854775808}
    var v3 uint32 = 0
    var v4 string = "hello"
    if true {
        var v5 float32 = 3.14159
        if false {
            var v6 int64 = -879
            v1 = false
            switch 1 {
            case -9223372036854775808:
                if v1 {
                    var v7 int32 = 42
                    if (!v1) {
                        for fi8 := 0; fi8 < 17; fi8++ {
                            v6 -= v6
                            _ = fi8
                        }
                        for fi9 := 0; fi9 < 10; fi9++ {
                            var v10 string = v4
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        v11 := GenS1{F0: (v6 >> ((745) & 63)), F1: 18446744073709551615, F2: 1}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v6)
                        _ = v11
                    }
                    _ = v7
                }
            case 42:
                switch v4 {
                case "\n\t":
                    var v12 int8 = 100
                    _ = v12
                case "":
                    var v13 bool = v1
                    _ = v13
                case "a":
                    v5 += v5
                }
            }
            _ = v6
        }
        _ = v5
    }
    if v1 {
        for fi14 := 0; fi14 < 10; fi14++ {
            switch (v4 + "foo bar") {
            case "a":
                _ = len(v2)
            case "":
                var v15 int8 = 0
                _ = v15
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if v1 {
                v16 := GenS1{F0: 673, F1: 18446744073709551615, F2: 627}
                _ = v16
            }
            fmt.Println(fi14, v1)
            _ = fi14
        }
        switch (("hello" + v4) + "a") {
        case "foo bar":
            var v17 uint16 = 65535
            _ = v17
        case "a":
            var v18 int8 = -1
            _ = v18
        case "\n\t":
            _ = len(v2)
        default:
            if ((1 != 609) || v1) {
                v4 = "a"
                var v19 int = 32767
                _ = v19
            }
        }
        v4 = v4
    }
    v1 = false
    var v20 float32 = 0.0
    fmt.Println(v4, v3, v20)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -336 {
    case 1:
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 string = v4
            if v1 {
                var v23 uint64 = 42
                _ = v23
            }
            var v24 int64 = 397
            _ = fi21
            _ = v22
            _ = v24
        }
    case 105:
        var v25 int = 576
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            if v1 {
                if (42 < -1) {
                    v28 := GenS2{F0: 733, F1: 851}
                    _ = v28
                }
            } else {
                var v29 uint64 = 925
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v29
            }
            v30 := GenS2{F0: -400, F1: int16(fi26)}
            _ = fi27
            _ = v30
        }
        _ = fi26
    }
    var v31 bool = v1
    switch 32768 {
    case 139:
        v20 = v20
    case 2147483647:
        var v32 bool = v31
        _ = v32
    }
    var v33 int32 = -525
    var v34 int = 4294967295
    v35 := GenS2{F0: -505, F1: -1}
    var v36 uint64 = 453
    if (v33 != (v33 >> ((18446744073709551615) & 63))) {
        var v37 int8 = 102
        v38 := GenS2{F0: 1000, F1: 527}
        var v39 int32 = 268
        _ = v37
        _ = v38
        _ = v39
    } else {
        var v40 bool = (!v1)
        _ = v40
    }
    v41 := GenS3{F0: (^(v3)), F1: 4294967295}
    v33 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
}
