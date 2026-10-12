package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8, p1 int64, p2 uint8) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: "a", F1: p1}
        if ((!(false && (false || true))) && ((p0 < p0) && ((float32(0.1) < 0.0) && ((true && (false && true)) || true)))) {
            switch "a" {
            case "hello":
                p1 = 0
            }
            p0 = p0
            v8 := []int{}
            _ = v8
        } else {
            p1 += p1
        }
        var v9 int32 = -608
        _ = fi6
        _ = v7
        _ = v9
    }
    v10 := GenS1{F0: "hello", F1: -1}
    v11 := GenS2{F0: 18446744073709551615}
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 4294967295
}
func genFn1(p0 uint16, p1 int32) int {
    p1 = p1
    var v13 uint16 = p0
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS3{F0: 142, F1: true}
        _ = fi14
        _ = v15
    }
    p0 += 1
    v16 := []int{-863, -1, -2147483648}
    _ = p0
    _ = p1
    _ = v13
    _ = v16
    return 32767
}

func main() {
    var v1 bool = (-714 != -2147483648)
    var v2 bool = v1
    var v3 int16 = -32768
    var v4 string = "foo bar"
    v5 := make(map[string]uint32)
    var v17 uint8 = 128
    var v18 uint8 = (v17 % ((42) | 1))
    if (!(-518 == 0)) {
        var v19 int8 = 73
        _ = v19
    } else {
        v1 = false
        v4 = (("\n\t" + ((v4 + "a") + v4)) + "hello")
        switch genFn1(42, 2147483647) {
        case 32768:
            switch -9223372036854775808 {
            case 0:
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int = genFn1(2, 1)
                    var v22 uint = 4294967295
                    fmt.Println(fi20, v17, v4)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case 32767:
                var v23 int32 = 42
                _ = v23
            }
        case 409:
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS1{F0: v4, F1: -9223372036854775808}
                v26 := []int{(fi24 % (((^(fi24))) | 1)), 0, (^(fi24)), fi24}
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 string = v4
        for fi29 := 0; fi29 < 19; fi29++ {
            switch fi29 {
            case -563:
                var v30 uint64 = 949
                _ = v30
            case 9223372036854775807:
                switch 727 {
                case 0:
                    var v31 uint32 = 0
                    _ = v31
                case 1:
                    fmt.Println(v4, v4)
                case 32768:
                    v32 := GenS4{F0: 94}
                    _ = v32
                }
            case 4294967295:
                v4 = ((v28 + "\n\t") + "\n\t")
            }
            var v33 float64 = (3.14159 + 0.0)
            _ = fi29
            _ = v33
        }
        _ = fi27
        _ = v28
    }
    var v34 bool = (3.14159 <= -2.25)
    switch 32767 {
    case 32767:
        switch ((((((v4 + v4) + v4) + v4) + v4) + "hello") + ((v4 + v4) + "")) {
        case "a":
            switch -799 {
            case 0:
                var v35 int = 128
                _ = v35
            default:
                v36 := GenS1{F0: "\n\t", F1: -820}
                _ = v36
            }
        default:
            v37 := GenS5{F0: 1000}
            _ = v37
        }
    case 1:
        delete(v5, v4)
    case 0:
        var v38 uint16 = 255
        _ = v38
    }
    if true {
        v18 += v18
        var v39 int16 = (-(v3))
        _ = v39
    } else {
        var v40 int64 = -248
        _ = v40
    }
    v41 := GenS1{F0: "a", F1: -1}
    fmt.Println(v3)
    var v42 float32 = 0.1
    var v43 uint16 = 334
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
