package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 uint; F1 int; F2 int }
type GenS5 struct { F0 uint; F1 int16; F2 string }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 bool; F1 int64; F2 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint16, p1 bool, p2 uint16) int32 {
    if (255 == 32767) {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := []int{4294967295}
            _ = fi5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 2, F1: 782, F2: 1}
            fmt.Println(p2)
            _ = fi7
            _ = v8
        }
        v9 := GenS2{F0: "hello", F1: -106, F2: -1}
        _ = v9
    }
    p2 = (^(p0))
    p0 = p0
    v10 := GenS3{F0: 42, F1: 2, F2: p1}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{(^(fi11))}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -337
}
func genFn1() string {
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 float32 = 0.0
        v15 := GenS1{F0: 0, F1: 65535, F2: -1}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    _ = v16
    return ""
}

func main() {
    var v1 int = -2147483648
    v2 := []int{v1, (v1 | v1), -9223372036854775808}
    var v3 string = "\n\t"
    var v4 float32 = 3.14159
    var v17 bool = false
    var v18 float64 = -2.25
    v19 := GenS3{F0: 42, F1: 874, F2: (false && (!(2147483647 != 2147483647)))}
    v3 = v3
    v20 := GenS4{F0: 505, F1: v1, F2: v1}
    var v21 uint8 = 2
    var v22 int = v1
    v17 = v17
    if len(v2) > 0 {
        v2[0] = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        if (v3 > "hello") {
            v24 := GenS5{F0: 255, F1: 0, F2: v3}
            v4 = float32(float64((v4 - (-2.25 + (v4 * v4)))))
            v25 := GenS4{F0: 18446744073709551615, F1: 0, F2: 4294967295}
            _ = v24
            _ = v25
        }
        v26 := GenS6{F0: (32767 >= 42), F1: 344}
        _ = fi23
        _ = v26
    }
    v3 = "a"
    var v27 int = v1
    var v28 float64 = -2.25
    if false {
        for fi29 := 0; fi29 < 15; fi29++ {
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS7{F0: ((fi30 - 1) < -354), F1: 241, F2: ((!(!(!v17))) && true)}
                if (v21 == (^(v21))) {
                    v32 := GenS8{F0: true}
                    _ = len(v2)
                    var v33 int16 = 32767
                    _ = v32
                    _ = v33
                }
                switch v22 {
                case 32768:
                    fmt.Println(v27)
                case -378:
                    switch v3 {
                    case "hello":
                        v34 := GenS7{F0: v17, F1: 9223372036854775807, F2: (v17 && false)}
                        _ = v34
                    }
                }
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    } else {
        if (v17 || ("hello" >= v3)) {
            switch v22 {
            case -2147483648:
                var v35 uint8 = 128
                _ = v35
            }
        } else {
            if (4294967295 >= 0) {
                if ((0.1 - (v28 + 1e10)) == v18) {
                    var v36 float32 = (v4 / 1.5)
                    fmt.Println(v27, v3)
                    _ = v36
                } else {
                    fmt.Println(v4)
                    v2 = append(v2, v1)
                    var v37 string = (v3 + ("" + ("\n\t" + "foo bar")))
                    _ = v37
                }
            }
            v18 = v28
            var v38 int64 = 9223372036854775807
            _ = v38
        }
    }
    switch ("hello" + "\n\t") {
    case "\n\t":
        fmt.Println(v28, v27, v4)
    default:
        if (v21 < v21) {
            v3 = v3
        } else {
            _ = len(v2)
            v39 := GenS9{F0: 255, F1: 18446744073709551615}
            v40 := GenS5{F0: 507, F1: -1, F2: "a"}
            _ = v39
            _ = v40
        }
    }
    fmt.Println(v1)
    fmt.Println(v27, v27)
    var v41 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
}
