package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 bool, p1 uint) uint64 {
    v4 := GenS1{F0: 0, F1: 255}
    if p0 {
        v5 := []int{32768, 968, 128}
        _ = v5
    } else {
        var v6 float64 = 1.5
        fmt.Println(p1, v6)
        if (861 >= 18446744073709551615) {
            var v7 int64 = -9223372036854775808
            var v8 int64 = 746
            _ = v7
            _ = v8
        }
        _ = v6
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 962
}
func genFn1() bool {
    switch ("a" + (("hello" + "") + "foo bar")) {
    case "hello":
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int32 = 2147483647
            _ = fi10
            _ = v11
        }
    default:
        if ((!true) && (!(42 > 0))) {
            if true {
                var v12 int = 128
                var v13 float64 = 0.1
                _ = v12
                _ = v13
            } else {
                v14 := []int{55}
                _ = len(v14)
                v15 := GenS2{F0: 0, F1: 65535}
                _ = v14
                _ = v15
            }
            if false {
                v16 := GenS1{F0: 2, F1: 142}
                var v17 int8 = 1
                _ = v16
                _ = v17
            }
        }
    }
    v18 := GenS2{F0: 42, F1: 42}
    var v19 uint16 = 0
    fmt.Println(v19)
    _ = v18
    _ = v19
    return (true && (false || (32767 > 42)))
}

func main() {
    var v1 int32 = 561
    var v2 uint8 = 2
    var v3 uint32 = 65535
    if (65535 != 65535) {
        if (("" != "a") || true) {
            v1 += (v1 + (v1 + (v1 | (v1 * -710))))
        }
    } else {
        v20 := GenS3{F0: "", F1: (("hello" + (("foo bar" + "hello") + ("a" + "a"))) + "\n\t"), F2: 18446744073709551615}
        var v21 uint64 = 500
        _ = v20
        _ = v21
    }
    v2 = v2
    v1 = v1
    var v22 string = (("foo bar" + ("" + "\n\t")) + "")
    if (980 == 490) {
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := []int{fi23}
            v25 := GenS4{F0: 0, F1: v2}
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        v3 = 2147483647
        v26 := GenS3{F0: v22, F1: v22, F2: 0}
        _ = v26
    }
    v27 := []int{1, -1, -1}
    var v28 string = v22
    var v29 int64 = 1
    var v30 uint = 732
    v22 = ((("foo bar" + "\n\t") + (((((v22 + v22) + ((v22 + "a") + "hello")) + v28) + v28) + v28)) + ("a" + v22))
    if (!false) {
        if len(v27) > 0 {
            v27[0] = 1
        }
    } else {
        var v31 float32 = ((1.5 - 1.5) - (3.14159 - (float32(0.1) + (0.1 - -2.25))))
        var v32 int8 = 0
        _ = v31
        _ = v32
    }
    v28 = v22
    v27 = append(v27, -2147483648)
    v27 = append(v27, -9223372036854775808)
    v33 := GenS2{F0: -32768, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
}
