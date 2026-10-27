package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0() int {
    var v7 uint16 = 32767
    var v8 bool = false
    var v9 int8 = -128
    if (950 < 1) {
        v8 = v8
        for fi10 := 0; fi10 < 9; fi10++ {
            v8 = ((0.0 <= 3.14159) && (208 < -622))
            v11 := []int{(fi10 / ((188) | 1)), 42, fi10, fi10}
            _ = fi10
            _ = v11
        }
        v12 := []int{}
        _ = v12
    }
    v13 := GenS1{F0: (v7 - (v7 * (^(v7)))), F1: 1}
    _ = v13
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = (3.14159 / 1.0)
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = v2
    var v4 uint32 = 4294967295
    var v5 bool = ((!(0 > 42)) && ("\n\t" == "\n\t"))
    var v6 bool = v5
    switch "hello" {
    case "a":
        switch "" {
        case "foo bar":
            v14 := GenS1{F0: 1, F1: -1}
            _ = v14
        case "\n\t":
            v4 = 4294967295
        case "hello":
            v1 = v1
        }
    case "":
        fmt.Println(v3, v5)
    }
    v15 := []int{}
    switch ("" + "a") {
    case "a":
        if (1.5 <= 0.0) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v16 uint = 1
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint16 = 903
            fmt.Println(v3)
            _ = fi17
            _ = v18
        }
    }
    var v19 int64 = -324
    var v20 float32 = -2.25
    var v21 uint = uint(v4)
    v21 = v21
    var v22 uint8 = 137
    if v6 {
        for fi23 := 0; fi23 < 19; fi23++ {
            v2 = v2
            var v24 uint16 = 42
            for fi25 := 0; fi25 < 13; fi25++ {
                v15 = append(v15, (fi23 % ((genFn0()) | 1)))
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        v26 := GenS2{F0: v19}
        v27 := GenS1{F0: 42, F1: 1}
        _ = v26
        _ = v27
    } else {
        var v28 int32 = 1
        v29 := GenS1{F0: 181, F1: -2147483648}
        switch "" {
        case "\n\t":
            v5 = v5
        case "":
            if (false && v6) {
                var v30 string = ((((("foo bar" + "foo bar") + "") + ("foo bar" + "\n\t")) + "foo bar") + ("\n\t" + ("a" + "\n\t")))
                var v31 float64 = v1
                if len(v15) > 0 {
                    v15[0] = -75
                }
                _ = v30
                _ = v31
            }
        }
        _ = v28
        _ = v29
    }
    var v32 uint32 = (v4 & 416)
    fmt.Println(v1)
    _ = len(v15)
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            v35 := GenS1{F0: uint16(v22), F1: 852}
            _ = fi34
            _ = v35
        }
        v36 := GenS3{F0: 255}
        _ = fi33
        _ = v36
    }
    v20 += ((v20 / float32(v1)) / ((1.5 + v20) + 0.1))
    if v6 {
        if true {
            fmt.Println(v2, v20, v3)
            v4 -= v4
            fmt.Println(v4, v1, v19)
        }
        switch 42 {
        case -995:
            v37 := GenS2{F0: v19}
            _ = v37
        case -9223372036854775808:
            v15 = append(v15, int(v21))
        default:
            var v38 uint8 = v22
            _ = v38
        }
        if (!true) {
            if true {
                switch (("foo bar" + "a") + ("a" + "a")) {
                case "":
                    var v39 uint64 = (v3 << ((2) & 63))
                    _ = v39
                }
                var v40 float32 = v20
                fmt.Println(v20, v4, v5)
                _ = v40
            }
            var v41 float32 = 0.1
            _ = v41
        } else {
            var v42 bool = (v6 || (42 == int32(v2)))
            _ = v42
        }
    }
    var v43 float32 = v20
    switch ("\n\t" + ("" + "a")) {
    case "\n\t":
        fmt.Println(v20, v43, v1)
    case "":
        if false {
            v19 -= (v19 - (v19 >> (((v21 & 4294967295)) & 63)))
        } else {
            if false {
                v6 = v5
                fmt.Println(v20)
            } else {
                fmt.Println(v3)
                v32 -= 2
                v4 = (v4 - (v4 / ((1) | 1)))
            }
            if false {
                for fi44 := 0; fi44 < 3; fi44++ {
                    v45 := GenS3{F0: (v22 % (((v22 % ((128) | 1))) | 1))}
                    _ = fi44
                    _ = v45
                }
                var v46 bool = v6
                v19 += v19
                _ = v46
            } else {
                v32 = 2147483648
            }
        }
    case "hello":
        v47 := GenS3{F0: (v22 % (((v22 | (^(v22)))) | 1))}
        _ = v47
    default:
        v3 = uint64(v21)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
