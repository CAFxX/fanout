package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0() uint32 {
    if (4294967295 < 2147483648) {
        var v4 int8 = -87
        _ = v4
    }
    var v5 int8 = 1
    v5 -= v5
    _ = v5
    return 174
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            if false {
                v8 := GenS1{F0: 255, F1: 42}
                v9 := GenS1{F0: 293, F1: 804}
                v3 -= v3
                _ = v8
                _ = v9
            }
            switch ("\n\t" + ("a" + "foo bar")) {
            case "":
                var v10 uint64 = 360
                _ = v10
            case "a":
                v3 = v3
            default:
                switch fi6 {
                case 802:
                    v3 += v3
                case -645:
                    if false {
                        _ = len(v1)
                        var v11 int = (-(fi7))
                        v12 := v2["a"]
                        _ = v11
                        _ = v12
                    }
                case 582:
                    var v13 uint64 = 1
                    _ = v13
                default:
                    switch "foo bar" {
                    case "a":
                        v3 += float32((0.0 + float64(v3)))
                    case "foo bar":
                        var v14 float32 = float32(float64(float32((1e10 / -2.25))))
                        _ = v14
                    }
                }
            }
            if (!(1 < -2147483648)) {
                v15 := []int{(fi7 << ((300) & 63)), (fi6 + 32768), fi7, fi6}
                if (((("\n\t" + "a") + "hello") + ("a" + "a")) > (("foo bar" + "\n\t") + "")) {
                    var v16 int8 = -128
                    _ = len(v2)
                    v17 := GenS2{F0: fi6, F1: 42}
                    v18 := GenS2{F0: (^(fi7)), F1: 32767}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                delete(v1, "")
                v19 := v2["\n\t"]
                _ = v15
                _ = v19
            } else {
                var v20 uint8 = 1
                _ = v20
            }
            _ = fi7
        }
        var v21 float32 = float32((-2.25 / 0.1))
        v22 := []int{1, 0}
        _ = fi6
        _ = v21
        _ = v22
    }
    fmt.Println(v3)
    delete(v2, "hello")
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 int8 = 1
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS3{F0: 0, F1: 819}
            _ = fi25
            _ = v26
        }
    } else {
        fmt.Println(v3)
        if false {
            var v27 uint32 = genFn0()
            v27 -= 65535
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := GenS1{F0: 255, F1: 996}
                _ = fi28
                _ = v29
            }
            _ = v27
        } else {
            v30 := v1[("\n\t" + ("foo bar" + "a"))]
            var v31 uint = 65535
            _ = v30
            _ = v31
        }
        v32 := []int{717, 0, 9223372036854775807, 1}
        _ = v32
    }
    var v33 string = "hello"
    var v34 uint = 18446744073709551615
    var v35 int = 127
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 string = "foo bar"
        var v38 int8 = -56
        v39 := []int{}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    switch -826 {
    case 541:
        var v40 uint = 65535
        _ = v40
    case -2147483648:
        if (false || (!false)) {
            var v41 int = 450
            _ = v41
        } else {
            v42 := GenS3{F0: 65535, F1: 513}
            _ = v42
        }
    case -946:
        if false {
            v3 += v3
            v3 += v3
            v33 = v33
        }
    default:
        var v43 float64 = ((-2.25 / 3.14159) - 3.14159)
        _ = v43
    }
    fmt.Println(v34, v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
