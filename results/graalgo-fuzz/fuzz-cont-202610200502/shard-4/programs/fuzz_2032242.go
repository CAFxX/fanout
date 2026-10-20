package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0() uint32 {
    if true {
        var v4 int16 = 0
        v5 := []int{2147483647, -1, 4294967295, 4294967295}
        var v6 uint16 = 547
        _ = v4
        _ = v5
        _ = v6
    } else {
        var v7 float32 = (((1e10 * (((0.1 / 1e10) + ((3.14159 + 0.0) * float32(1e10))) - (float32(-2.25) * 1e10))) - (0.0 * (0.0 * -2.25))) - 3.14159)
        _ = v7
    }
    v8 := GenS1{F0: 100, F1: 569, F2: 127}
    switch "" {
    case "\n\t":
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
    case "a":
        v10 := GenS1{F0: 127, F1: 2, F2: 2}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(fi11)
        var v12 int16 = 1
        _ = fi11
        _ = v12
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if false {
            var v13 uint8 = 127
            _ = v13
        } else {
            var v14 uint16 = 1
            v14 += (^(v14))
            _ = v14
        }
    case "foo bar":
        v15 := []int{128}
        _ = v15
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint32 = 110
            var v18 int64 = 736
            v19 := GenS1{F0: -128, F1: 1, F2: 248}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    }
    _ = v8
    return 1
}
func genFn1(p0 int16) int16 {
    var v20 uint = 0
    for fi21 := 0; fi21 < 15; fi21++ {
        v20 = v20
        _ = fi21
    }
    var v22 int16 = (p0 | (p0 ^ (p0 | p0)))
    v23 := GenS2{F0: (v20 + (v20 & v20))}
    _ = p0
    _ = v20
    _ = v22
    _ = v23
    return (p0 + (^(v22)))
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 1.5
    v3 := make(map[string]uint32)
    var v24 int = -1
    var v25 string = ("a" + "a")
    v2 -= v2
    switch "\n\t" {
    case "hello":
        switch (v25 + ("foo bar" + v25)) {
        case "foo bar":
            v2 = float64(0.1)
        case "\n\t":
            if true {
                fmt.Println(v25, v1, v2)
                v2 += ((0.0 * v2) * v2)
            }
        }
    case "foo bar":
        if true {
            var v26 uint32 = 2147483648
            _ = v26
        }
    }
    v27 := GenS2{F0: 340}
    v28 := []int{(v24 * v24), (v24 / ((v24) | 1))}
    var v29 int32 = -2147483648
    var v30 int = v24
    if (v25 != v25) {
        var v31 uint16 = 65535
        if false {
            delete(v3, "foo bar")
            v32 := GenS3{F0: -1, F1: 65535}
            v33 := GenS1{F0: -1, F1: 1, F2: 2}
            _ = v32
            _ = v33
        }
        var v34 int16 = v1
        _ = v31
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 uint64 = 4294967295
            v38 := GenS1{F0: 42, F1: v37, F2: 0}
            v2 += v2
            switch fi35 {
            case 2147483647:
                v39 := GenS1{F0: 42, F1: (v37 - 2), F2: 255}
                _ = v39
            default:
                fmt.Println(v37, v25, fi36)
            }
            v24 = v30
            v40 := GenS1{F0: 59, F1: 698, F2: 42}
            _ = fi36
            _ = v37
            _ = v38
            _ = v40
        }
        _ = fi35
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        v29 = (v29 & (v29 - v29))
        if false {
            for fi42 := 0; fi42 < 19; fi42++ {
                if (42 != 42) {
                    switch (^(v30)) {
                    case -556:
                        var v43 uint8 = 220
                        _ = v43
                    case -2147483648:
                        fmt.Println(v30, v2)
                    default:
                        v25 = "foo bar"
                    }
                }
                if (!true) {
                    var v44 uint64 = 4294967295
                    _ = len(v3)
                    v45 := GenS3{F0: 282, F1: 2}
                    _ = v44
                    _ = v45
                }
                v24 += (fi41 % (((fi41 & v24)) | 1))
                _ = fi42
            }
            for fi46 := 0; fi46 < 2; fi46++ {
                v47 := v3[((v25 + v25) + (("hello" + v25) + v25))]
                var v48 int8 = 100
                _ = fi46
                _ = v47
                _ = v48
            }
            v24 -= (^(v24))
        }
        if (true && false) {
            var v49 bool = ((1 != 1) && true)
            _ = v49
        }
        _ = fi41
    }
    var v50 int16 = (v1 & v1)
    v2 += 1.5
    if false {
        var v51 int = -2147483648
        _ = v51
    }
    for fi52 := 0; fi52 < 2; fi52++ {
        var v53 uint8 = 163
        v54 := GenS4{F0: 880, F1: (^(v53)), F2: (v29 | (v29 & v29))}
        _ = fi52
        _ = v53
        _ = v54
    }
    var v55 int = v30
    var v56 string = v25
    if false {
        var v57 float32 = float32((v2 + ((0.1 * 0.1) + (v2 - float64(-2.25)))))
        if true {
            v29 += (v29 % ((v29) | 1))
        } else {
            var v58 float32 = v57
            var v59 int8 = 0
            v24 = 410
            _ = v58
            _ = v59
        }
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v50
    _ = v55
    _ = v56
}
