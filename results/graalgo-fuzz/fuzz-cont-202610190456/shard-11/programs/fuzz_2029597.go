package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 int) uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float64 = (1.5 / 3.14159)
        _ = fi7
        _ = v8
    }
    var v9 string = ("\n\t" + ("foo bar" + "a"))
    _ = p0
    _ = p1
    _ = v9
    return 4294967295
}
func genFn1(p0 int16, p1 uint32) string {
    fmt.Println(p1, p1)
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(p1, p0, fi10)
        _ = fi10
    }
    var v11 float64 = 0.1
    _ = p0
    _ = p1
    _ = v11
    return ("a" + "hello")
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]int)
    var v3 uint8 = 119
    v4 := make(map[string]string)
    var v5 uint64 = 4294967295
    v6 := make(map[string]int)
    v12 := GenS1{F0: 558, F1: 116, F2: true}
    if true {
        fmt.Println(v3)
        switch -46 {
        case 128:
            var v13 uint64 = v5
            _ = v13
        case 1:
            v14 := []int{-1, 914, 128, -1}
            _ = v14
        case -551:
            var v15 int16 = 1
            _ = v15
        }
    }
    var v16 float32 = float32(float64(v1))
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := []int{32768, (fi17 & -459), (fi17 * -9223372036854775808), fi17}
        _ = fi17
        _ = v18
    }
    if (2147483648 >= 1) {
        var v19 int16 = -1
        _ = v19
    } else {
        v20 := v2[""]
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint16 = 255
            v5 = uint64(v22)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    v23 := []int{-9223372036854775808}
    var v24 uint = 4294967295
    var v25 int32 = 263
    var v26 uint64 = v5
    fmt.Println(v16, v16)
    fmt.Println(v3, v25, v1)
    var v27 int32 = (v25 & 1)
    for fi28 := 0; fi28 < 1; fi28++ {
        v25 -= -2147483648
        _ = fi28
    }
    var v29 uint = (^(v24))
    if false {
        for fi30 := 0; fi30 < 9; fi30++ {
            delete(v2, "foo bar")
            if false {
                v31 := v2[(("" + "foo bar") + "foo bar")]
                if (9223372036854775807 != 725) {
                    switch 42 {
                    case 34:
                        if (v5 > (v5 * v26)) {
                            v23 = append(v23, fi30)
                        } else {
                            v27 = -2147483648
                            var v32 float32 = v16
                            _ = v32
                        }
                    case 111:
                        var v33 uint16 = 32767
                        _ = v33
                    case -9223372036854775808:
                        _ = len(v23)
                    }
                    for fi34 := 0; fi34 < 2; fi34++ {
                        switch (fi34 | fi30) {
                        case 1:
                            fmt.Println(fi30, v27, v27)
                        case -9223372036854775808:
                            v5 = (v5 % (((^(v5))) | 1))
                        }
                        _ = fi34
                    }
                } else {
                    fmt.Println(v27, v29, v1)
                    v1 -= v1
                    v35 := GenS2{F0: 65535}
                    _ = v35
                }
                var v36 int16 = 11
                _ = v31
                _ = v36
            }
            v37 := GenS3{F0: 32767}
            _ = fi30
            _ = v37
        }
        fmt.Println(v1, v1)
        var v38 float64 = (-2.25 * 1e10)
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
