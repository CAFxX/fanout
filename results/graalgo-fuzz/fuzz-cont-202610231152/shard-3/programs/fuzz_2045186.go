package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int8, p1 int8) bool {
    for fi4 := 0; fi4 < 5; fi4++ {
        switch (^(fi4)) {
        case -331:
            var v5 float32 = 3.14159
            _ = v5
        case -9223372036854775808:
            switch (("\n\t" + "hello") + "") {
            case "hello":
                var v6 int8 = p0
                _ = v6
            default:
                fmt.Println(fi4, p1)
            }
        case -1:
            var v7 int8 = (p0 - (p0 / ((p1) | 1)))
            _ = v7
        }
        v8 := []int{(fi4 + fi4)}
        _ = fi4
        _ = v8
    }
    v9 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v9
    return true
}
func genFn1(p0 uint32, p1 int64) int {
    switch (("hello" + ("\n\t" + "")) + ("" + ("\n\t" + (("foo bar" + "\n\t") + "\n\t")))) {
    case "a":
        for fi10 := 0; fi10 < 3; fi10++ {
            if (true || (0.1 >= 1e10)) {
                var v11 int64 = (-(p1))
                v12 := GenS1{F0: ((3.14159 * ((1e10 / 1.5) / (1.5 / 1.0))) != 0.0)}
                v13 := GenS1{F0: (806 == 1)}
                _ = v11
                _ = v12
                _ = v13
            }
            p0 += 2147483648
            if (809 != 18446744073709551615) {
                var v14 bool = (100 < 0)
                if v14 {
                    v15 := GenS1{F0: true}
                    _ = v15
                }
                var v16 uint64 = 0
                _ = v14
                _ = v16
            }
            _ = fi10
        }
    case "foo bar":
        if false {
            fmt.Println(p1, p0)
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 float64 = ((float64((((3.14159 + 1.5) + 1.5) + (1.5 - float32(3.14159)))) - 0.1) + (float64(1e10) / 1e10))
        var v19 float64 = ((v18 - (v18 + (v18 - float64(0.1)))) / float64((0.1 + float32((float64((0.1 * ((-2.25 - float32(v18)) * (3.14159 - 0.0)))) / v18)))))
        var v20 int64 = 0
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    p1 = (p1 + (p1 ^ (p1 ^ (-(p1)))))
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int32 = 42
    var v2 float64 = (-2.25 + 3.14159)
    var v3 uint8 = 201
    switch "" {
    case "hello":
        if false {
            v21 := GenS1{F0: false}
            v3 = (^(v3))
            _ = v21
        }
    case "":
        var v22 bool = (255 <= 725)
        _ = v22
    case "a":
        if (-564 < 1) {
            v23 := GenS2{F0: 201, F1: (v1 * (^(v1))), F2: 127}
            fmt.Println(v2, v2)
            v24 := GenS1{F0: ((9223372036854775807 != 1) && true)}
            _ = v23
            _ = v24
        } else {
            v25 := []int{42, -2147483648}
            var v26 uint8 = v3
            _ = v25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(v3, v2, fi27)
        _ = fi27
    }
    v28 := []int{-953, 9223372036854775807, -331}
    if false {
        v2 += ((v2 + v2) * -2.25)
    } else {
        if true {
            for fi29 := 0; fi29 < 17; fi29++ {
                _ = len(v28)
                for fi30 := 0; fi30 < 3; fi30++ {
                    if true {
                        var v31 uint32 = 2147483648
                        var v32 uint = 0
                        _ = v31
                        _ = v32
                    }
                    var v33 uint32 = 4294967295
                    if true {
                        switch (-(fi30)) {
                        case -1:
                            var v34 int8 = 117
                            _ = v34
                        case -346:
                            _ = len(v28)
                        case 9223372036854775807:
                            v35 := GenS3{F0: 1000, F1: 100}
                            v1 = v1
                            _ = v35
                        default:
                            v36 := GenS1{F0: (42 == 19)}
                            _ = v36
                        }
                    } else {
                        v3 -= v3
                        var v37 int64 = -1
                        v38 := GenS4{F0: -308, F1: ("" + ("a" + "a"))}
                        _ = v37
                        _ = v38
                    }
                    v33 -= v33
                    _ = fi30
                    _ = v33
                }
                v2 -= (float64((0.1 - 0.0)) * v2)
                _ = fi29
            }
        }
        if (!(!(!true))) {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            v39 := GenS5{F0: 2}
            var v40 float64 = v2
            _ = v39
            _ = v40
        }
        v1 = 42
    }
    fmt.Println(v2)
    var v41 float32 = -2.25
    v28 = append(v28, int(v1))
    if true {
        for fi42 := 0; fi42 < 14; fi42++ {
            for fi43 := 0; fi43 < 9; fi43++ {
                var v44 float64 = float64(float32(v2))
                _ = len(v28)
                if (18446744073709551615 > 2) {
                    v1 -= (v1 << ((65535) & 63))
                } else {
                    var v45 uint32 = 2147483647
                    fmt.Println(v3)
                    var v46 uint16 = 0
                    _ = v45
                    _ = v46
                }
                v44 = ((v44 / 0.1) * 3.14159)
                _ = fi43
                _ = v44
            }
            _ = fi42
        }
        var v47 int16 = -32768
        v47 = v47
        _ = v47
    }
    fmt.Println(v3, v3, v41)
    v3 -= (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v41
}
