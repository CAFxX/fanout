package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 string) int16 {
    if (!false) {
        var v4 float32 = 1.5
        _ = v4
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 32767
}
func genFn1() uint16 {
    var v5 string = "hello"
    v5 = "a"
    _ = v5
    return 65535
}

func main() {
    var v1 int16 = 42
    var v2 uint32 = 1
    v3 := make(map[string]int)
    if (128 != 131) {
        if false {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 float32 = 3.14159
                if (!true) {
                    if false {
                        v8 := []int{}
                        _ = v8
                    }
                    fmt.Println(v2, v1)
                }
                var v9 int8 = -62
                _ = fi6
                _ = v7
                _ = v9
            }
        }
        v2 = 2
    }
    v10 := GenS1{F0: 127}
    v2 = 97
    v11 := v3[""]
    switch (-(v11)) {
    case -937:
        if false {
            var v12 uint32 = 2147483648
            _ = v12
        }
    default:
        fmt.Println(v2)
    }
    if false {
        if ((!true) && ("foo bar" >= "\n\t")) {
            v1 -= v1
            fmt.Println(v11, v1, v2)
        }
    } else {
        fmt.Println(v11, v11, v11)
        switch 32767 {
        case 324:
            v13 := GenS2{F0: 1, F1: 0}
            _ = v13
        case 240:
            v14 := []int{(v11 * 32768), (v11 / ((v11) | 1)), -9223372036854775808, (^(v11))}
            _ = v14
        case -161:
            v15 := GenS3{F0: -7}
            _ = v15
        default:
            for fi16 := 0; fi16 < 4; fi16++ {
                _ = len(v3)
                v1 -= v1
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int = fi17
                    var v19 int16 = (v1 ^ -749)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
        fmt.Println(v1)
    }
    fmt.Println(v11, v11)
    if true {
        if (1 > 1) {
            var v20 int32 = 42
            var v21 uint64 = 1
            _ = v20
            _ = v21
        } else {
            v1 = v1
            switch "" {
            case "hello":
                v22 := GenS3{F0: -16}
                _ = v22
            case "foo bar":
                var v23 int = v11
                _ = v23
            }
            v1 = v1
        }
        if ((!false) && true) {
            fmt.Println(v2, v1, v11)
        }
    } else {
        v2 = 963
        if false {
            var v24 uint = 103
            _ = v24
        } else {
            _ = len(v3)
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 float32 = float32(1.5)
                v26 = v26
                var v27 int32 = -362
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 string = (("hello" + ("\n\t" + "a")) + "a")
            _ = v28
        }
    }
    v3[""] = v11
    v29 := GenS4{F0: -9223372036854775808}
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := []int{fi30, fi30, 128}
        _ = fi30
        _ = v31
    }
    if (false || false) {
        switch "" {
        case "\n\t":
            var v32 float64 = 1e10
            _ = v32
        case "a":
            switch ("" + "foo bar") {
            case "foo bar":
                v33 := GenS4{F0: 528}
                _ = v33
            default:
                var v34 uint64 = 0
                _ = v34
            }
        }
        v35 := GenS4{F0: -9223372036854775808}
        v36 := GenS5{F0: 4294967295, F1: (v2 >> ((57) & 63))}
        _ = v35
        _ = v36
    } else {
        v2 -= v2
    }
    v37 := []int{v11, v11, 32767}
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    if false {
        v38 := GenS6{F0: 0}
        _ = v38
    }
    fmt.Println(v11, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v37
}
