package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint16 }
func genFn0() uint16 {
    v7 := GenS1{F0: (!false), F1: 41, F2: "foo bar"}
    v8 := []int{-1, 42, 121, 0}
    var v9 uint8 = 0
    if true {
        for fi10 := 0; fi10 < 13; fi10++ {
            switch "\n\t" {
            case "foo bar":
                if len(v8) > 0 {
                    v8[0] = 4294967295
                }
            case "":
                if ((("hello" + ("hello" + (("hello" + ("" + "")) + ""))) + "\n\t") <= (("hello" + ("a" + "foo bar")) + "")) {
                    if true {
                        v9 = 127
                        v9 += v9
                    }
                    var v11 uint64 = 2
                    _ = v11
                }
            case "hello":
                if true {
                    v9 += (v9 % ((128) | 1))
                    if len(v8) > 0 {
                        v8[0] = (fi10 & 686)
                    }
                }
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int = fi12
                switch "a" {
                case "foo bar":
                    if len(v8) > 0 {
                        v8[0] = 127
                    }
                case "\n\t":
                    fmt.Println(fi12, v9, fi10)
                }
                if (2 >= 381) {
                    var v14 int32 = -2147483648
                    var v15 int = 1
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            if len(v8) > 0 {
                v8[0] = -550
            }
            _ = fi10
        }
        var v16 int = 32767
        v17 := GenS1{F0: (!false), F1: 65535, F2: (("hello" + "") + ((("foo bar" + "a") + "\n\t") + "foo bar"))}
        _ = v16
        _ = v17
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint {
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            if true {
                v20 := []int{127, (fi18 - -9223372036854775808), fi19}
                _ = v20
            }
            v21 := GenS2{F0: 42, F1: 42}
            fmt.Println(fi19, fi18)
            _ = fi19
            _ = v21
        }
        switch fi18 {
        case 32768:
            switch (fi18 ^ (^(fi18))) {
            case -43:
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint = 943
                    _ = fi22
                    _ = v23
                }
            case 32768:
                var v24 uint64 = 882
                _ = v24
            case 4294967295:
                for fi25 := 0; fi25 < 10; fi25++ {
                    if ((((602 >= 18446744073709551615) || (99 >= 0)) || (false && false)) && (128 < 119)) {
                        var v26 uint32 = 509
                        var v27 float64 = 3.14159
                        v26 += (v26 & 2147483647)
                        _ = v26
                        _ = v27
                    } else {
                        var v28 uint64 = 2
                        _ = v28
                    }
                    _ = fi25
                }
            }
        case -2147483648:
            var v29 int32 = 0
            _ = v29
        }
        _ = fi18
    }
    v30 := GenS3{F0: false, F1: -1}
    v31 := []int{-550}
    var v32 uint64 = 305
    switch 128 {
    case 9223372036854775807:
        switch "hello" {
        case "a":
            var v33 float64 = (1e10 / (0.0 - -2.25))
            _ = v33
        case "\n\t":
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 int32 = 853
                fmt.Println(v35, fi34, fi34)
                _ = fi34
                _ = v35
            }
        case "foo bar":
            if true {
                switch 285 {
                case 1:
                    var v36 int16 = 0
                    _ = v36
                default:
                    v32 = v32
                }
                var v37 float64 = 1.5
                _ = v37
            } else {
                v32 += v32
            }
        }
    case 2147483647:
        v32 = 638
    case 4294967295:
        _ = len(v31)
    }
    _ = v30
    _ = v31
    _ = v32
    return 263
}

func main() {
    v1 := []int{}
    var v2 uint8 = 128
    var v3 int16 = 32767
    v4 := []int{-189, -912}
    var v5 int32 = -1
    v6 := []uint32{2147483648, 134, 0, 0}
    v5 += v5
    if (565 <= (v3 ^ -32768)) {
        v38 := GenS4{F0: (v2 | (^(v2))), F1: (v2 & (v2 << ((1) & 63))), F2: (v2 & 173)}
        var v39 float64 = 3.14159
        _ = v38
        _ = v39
    }
    if true {
        if (float64((-2.25 + (-2.25 + 3.14159))) <= 3.14159) {
            var v40 uint = uint(v2)
            switch 2147483647 {
            case -147:
                if true {
                    v41 := GenS5{F0: 0, F1: genFn0()}
                    var v42 int64 = 1
                    _ = len(v6)
                    _ = v41
                    _ = v42
                }
            case 2147483647:
                var v43 int32 = int32(v2)
                _ = v43
            case 4294967295:
                var v44 float64 = (-2.25 / (1.5 / 0.1))
                _ = v44
            }
            _ = v40
        }
        var v45 uint32 = 65535
        var v46 bool = (!((!false) && ((-9223372036854775808 < 1) || false)))
        _ = v45
        _ = v46
    }
    var v47 uint16 = 32767
    v48 := GenS6{F0: v47}
    var v49 int16 = (^(v3))
    switch 128 {
    case 741:
        for fi50 := 0; fi50 < 12; fi50++ {
            if true {
                fmt.Println(v47)
                fmt.Println(v2)
            } else {
                var v51 int16 = (v3 - v49)
                _ = v51
            }
            var v52 bool = (!(true && (65535 == 4294967295)))
            var v53 float64 = 0.0
            _ = fi50
            _ = v52
            _ = v53
        }
    case 128:
        var v54 float32 = (((0.1 / 1.0) + float32(-2.25)) + (-2.25 / (-2.25 + 1.5)))
        _ = v54
    }
    var v55 int32 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v55
    _ = v6
}
