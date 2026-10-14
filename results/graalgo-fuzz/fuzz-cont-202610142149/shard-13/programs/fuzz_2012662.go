package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 string = ("foo bar" + (("hello" + "hello") + "a"))
    var v5 uint16 = 702
    fmt.Println(v4, v4)
    if ((-361 < -421) && (false || false)) {
        v6 := GenS1{F0: -32768, F1: 510, F2: 131}
        _ = v6
    }
    if true {
        v7 := GenS1{F0: 32767, F1: 42, F2: -599}
        fmt.Println(v5, v4)
        fmt.Println(v5, v4)
        _ = v7
    } else {
        var v8 int8 = -1
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint = 255
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v4, v4)
            if false {
                v8 -= (v8 & v8)
                v8 += (v8 * -1)
            }
            _ = fi11
        }
        _ = v8
    }
    var v12 int8 = 64
    var v13 int32 = 2147483647
    v4 = "a"
    switch ("\n\t" + ((("foo bar" + (v4 + v4)) + v4) + ("\n\t" + (((v4 + "foo bar") + "a") + "\n\t")))) {
    case "\n\t":
        if true {
            var v14 uint16 = 32767
            var v15 float64 = 0.0
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: v13, F1: -32768}
            v17 := []int{}
            if len(v17) > 0 {
                v17[0] = 787
            }
            _ = v16
            _ = v17
        }
    case "":
        switch 2147483647 {
        case 537:
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(v4, fi18)
                for fi19 := 0; fi19 < 16; fi19++ {
                    v1[((((v4 + v4) + ("a" + "\n\t")) + v4) + v4)] = 2
                    v20 := GenS1{F0: 32767, F1: 65535, F2: -1}
                    _ = fi19
                    _ = v20
                }
                var v21 uint = 776
                _ = fi18
                _ = v21
            }
        case -361:
            if ((v13 | -2147483648) != v13) {
                var v22 int8 = (v12 % ((v12) | 1))
                _ = v22
            }
        default:
            var v23 int16 = int16(v13)
            _ = v23
        }
    case "a":
        v24 := []int{-624, 2147483647}
        _ = v24
    default:
        if true {
            v25 := GenS2{F0: (v13 ^ (v13 * v13)), F1: 1}
            _ = v25
        } else {
            var v26 int32 = -2147483648
            v27 := []int{32768, -829, -988, -9223372036854775808}
            _ = v26
            _ = v27
        }
    }
    v28 := []int{}
    v4 = "a"
    v5 = (v5 / (((v5 * v5)) | 1))
    if true {
        v29 := GenS1{F0: 1000, F1: 995, F2: 756}
        if (144 == 2) {
            if len(v28) > 0 {
                v28[0] = -1
            }
        } else {
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS1{F0: 0, F1: 42, F2: 128}
                var v32 float32 = 1e10
                _ = fi30
                _ = v31
                _ = v32
            }
            v12 = (v12 >> ((1) & 63))
            var v33 uint = 4294967295
            _ = v33
        }
        v13 += v13
        _ = v29
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    switch "" {
    case "hello":
        v34 := v3[2147483648]
        _ = v34
    case "":
        v35 := GenS1{F0: 198, F1: 42, F2: 32767}
        _ = v35
    case "a":
        v36 := GenS3{F0: 2147483647, F1: (v5 << ((2) & 63))}
        _ = v36
    default:
        if (false && true) {
            for fi37 := 0; fi37 < 8; fi37++ {
                fmt.Println(v5, v13, fi37)
                fmt.Println(v5, v5)
                _ = fi37
            }
            if false {
                for fi38 := 0; fi38 < 13; fi38++ {
                    fmt.Println(v4)
                    _ = fi38
                }
            } else {
                var v39 uint = 42
                var v40 int16 = 887
                _ = v39
                _ = v40
            }
            v4 = v4
        }
    }
    v4 = "foo bar"
    var v41 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
