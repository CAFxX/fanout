package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }
func genFn0(p0 bool, p1 int8) int32 {
    p0 = (9223372036854775807 < -1)
    if true {
        p1 = (p1 / ((-128) | 1))
        p1 = (p1 + p1)
        fmt.Println(p0, p1)
    }
    v7 := []int{-1, 183, -421}
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1() uint32 {
    var v8 string = ("\n\t" + "a")
    if true {
        var v9 float32 = 1e10
        v9 -= float32(((0.0 * -2.25) + (1.5 - ((0.0 / (((-2.25 + 0.1) * ((((0.1 * 3.14159) + float64(v9)) / 3.14159) + 1e10)) * 3.14159)) - -2.25))))
        v8 = v8
        _ = v9
    }
    switch 32768 {
    case 1:
        for fi10 := 0; fi10 < 10; fi10++ {
            switch (fi10 % (((fi10 & -122)) | 1)) {
            case 840:
                fmt.Println(fi10, v8)
            case -718:
                switch (v8 + "foo bar") {
                case "hello":
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := []int{-2147483648, fi11, fi10, 128}
                        _ = fi11
                        _ = v12
                    }
                }
            default:
                v13 := GenS1{F0: 69}
                _ = v13
            }
            _ = fi10
        }
    case 32768:
        if (!false) {
            var v14 uint = 42
            _ = v14
        }
    case 32767:
        var v15 int64 = 9223372036854775807
        _ = v15
    default:
        var v16 int8 = 42
        _ = v16
    }
    v8 = "a"
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 float32 = ((-2.25 * (float32((1e10 * 1e10)) - 0.1)) - -2.25)
    var v5 uint8 = 127
    v6 := make(map[string]bool)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int = fi17
        v19 := []int{4294967295, 9223372036854775807}
        for fi20 := 0; fi20 < 4; fi20++ {
            if len(v19) > 0 {
                v19[0] = 2147483647
            }
            _ = fi20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v21 int8 = -1
    fmt.Println(v1, v1, v5)
    if (-9223372036854775808 == 0) {
        v22 := []int{}
        var v23 string = ("hello" + ("hello" + "\n\t"))
        _ = v22
        _ = v23
    } else {
        v4 -= v4
    }
    var v24 bool = (true || false)
    v25 := GenS1{F0: 32767}
    var v26 int8 = (v21 & v21)
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 uint16 = 0
        _ = fi27
        _ = v28
    }
    v26 = -128
    v29 := []int{2147483647, -1}
    v30 := GenS2{F0: 255, F1: -1, F2: 1}
    v31 := GenS1{F0: 37}
    if v24 {
        if ((v24 && true) || (v24 || false)) {
            v4 += 0.0
            var v32 uint = 0
            v33 := GenS3{F0: 32767, F1: -1, F2: 32767}
            _ = v32
            _ = v33
        } else {
            var v34 uint32 = 2
            fmt.Println(v4, v5)
            for fi35 := 0; fi35 < 5; fi35++ {
                if true {
                    v26 -= (v21 << ((18446744073709551615) & 63))
                    v26 += 37
                    for fi36 := 0; fi36 < 5; fi36++ {
                        if false {
                            var v37 float32 = 3.14159
                            v4 = ((v37 / (float32(float64(v4)) / (-2.25 - v37))) - v37)
                            var v38 uint32 = (^(v34))
                            _ = v37
                            _ = v38
                        } else {
                            v2["a"] = (v34 << ((0) & 63))
                            switch (fi35 | fi36) {
                            case -9223372036854775808:
                                var v39 uint64 = v1
                                fmt.Println(v1, fi36, fi36)
                                var v40 uint8 = (v5 / (((^(v5))) | 1))
                                _ = v39
                                _ = v40
                            case -352:
                                v41 := GenS4{F0: 255, F1: -9223372036854775808, F2: 32767}
                                _ = v41
                            default:
                                var v42 float32 = 3.14159
                                _ = v42
                            }
                        }
                        _ = fi36
                    }
                }
                _ = fi35
            }
            _ = v34
        }
        v43 := v2["foo bar"]
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
