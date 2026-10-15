package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint32; F2 uint32 }

func main() {
    var v1 int64 = 42
    var v2 bool = false
    var v3 int8 = 100
    var v4 uint16 = 507
    var v5 int = 356
    var v6 int8 = v3
    switch ((("hello" + "") + ((("hello" + "") + "hello") + "\n\t")) + ("a" + (("" + ("" + ("hello" + ""))) + "a"))) {
    case "hello":
        v3 = 0
    case "foo bar":
        v7 := []int{-469, (v5 + v5)}
        _ = v7
    }
    switch 32767 {
    case 592:
        v8 := GenS1{F0: v6, F1: 1, F2: (("foo bar" + "foo bar") + "hello")}
        _ = v8
    }
    var v9 int64 = (-(v1))
    for fi10 := 0; fi10 < 3; fi10++ {
        if ((v4 | v4) < v4) {
            var v11 int64 = v1
            v12 := GenS2{F0: 2}
            v3 -= v6
            _ = v11
            _ = v12
        } else {
            v6 = 0
            fmt.Println(fi10)
        }
        var v13 bool = v2
        _ = fi10
        _ = v13
    }
    var v14 int32 = int32(v5)
    v1 = v9
    fmt.Println(v1)
    if v2 {
        v4 -= (v4 ^ v4)
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS3{F0: 127, F1: (v4 - v4)}
            _ = fi15
            _ = v16
        }
    }
    switch (v5 / (((^(v5))) | 1)) {
    case 32768:
        v5 += 1
    case 32767:
        fmt.Println(v14, v3)
    case 543:
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS4{F0: (^(v3)), F1: fi18}
                v20 := []int{v5, (fi18 & int(v4)), 2147483647, 128}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    default:
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := []int{fi21}
            var v23 uint32 = 587
            switch "foo bar" {
            case "a":
                for fi24 := 0; fi24 < 17; fi24++ {
                    if (((v6 >> ((572) & 63)) < (v6 << ((0) & 63))) && v2) {
                        var v25 uint16 = v4
                        _ = v25
                    }
                    if (!true) {
                        fmt.Println(v1)
                    } else {
                        var v26 int32 = v14
                        fmt.Println(v5)
                        var v27 uint16 = (v4 + v4)
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
            case "foo bar":
                v28 := GenS3{F0: 255, F1: (v4 + uint16(v23))}
                _ = v28
            case "":
                var v29 float64 = -2.25
                _ = v29
            default:
                var v30 uint8 = 42
                _ = v30
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := []int{(-(v5)), 42}
        for fi33 := 0; fi33 < 8; fi33++ {
            v9 -= v1
            _ = len(v32)
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if false {
            var v35 int8 = (v3 | (v3 | v6))
            _ = v35
        } else {
            for fi36 := 0; fi36 < 7; fi36++ {
                switch ("\n\t" + ((("\n\t" + ("a" + "foo bar")) + ("hello" + ("\n\t" + "foo bar"))) + ("\n\t" + "\n\t"))) {
                case "a":
                    v2 = v2
                case "hello":
                    v37 := []int{v5, fi34, -1, (fi34 * 194)}
                    _ = v37
                }
                v38 := GenS5{F0: 335}
                if true {
                    fmt.Println(v5, fi36)
                    if (0 <= 1) {
                        switch (fi34 - (fi34 % ((v5) | 1))) {
                        case 128:
                            v3 = v3
                        case 9223372036854775807:
                            var v39 int8 = (v3 * (v6 >> ((4294967295) & 63)))
                            _ = v39
                        }
                        var v40 bool = false
                        var v41 float64 = (0.0 / (3.14159 + ((0.1 - -2.25) - 3.14159)))
                        _ = v40
                        _ = v41
                    }
                } else {
                    var v42 int16 = 1
                    _ = v42
                }
                v43 := GenS3{F0: 127, F1: v4}
                if (!v2) {
                    switch "hello" {
                    case "foo bar":
                        v44 := []int{(fi36 & fi36), 330, (fi36 % (((v5 + 42)) | 1))}
                        _ = v44
                    case "a":
                        fmt.Println(fi34, v6, v1)
                    default:
                        var v45 float64 = 0.0
                        _ = v45
                    }
                    var v46 int64 = (v1 + (-(v1)))
                    _ = v46
                } else {
                    v47 := []int{(v5 % (((fi34 + (fi34 ^ (^(fi36))))) | 1))}
                    v48 := GenS1{F0: v3, F1: 255, F2: "\n\t"}
                    var v49 uint32 = 138
                    _ = v47
                    _ = v48
                    _ = v49
                }
                v50 := GenS6{F0: (v3 != -128), F1: 1, F2: 42}
                v51 := []int{fi36, (fi36 + 2147483647), (v5 & v5)}
                if len(v51) > 0 {
                    v51[0] = (fi36 & fi36)
                }
                _ = fi36
                _ = v38
                _ = v43
                _ = v50
                _ = v51
            }
            v1 = v1
        }
        v52 := []int{71, (^(fi34))}
        _ = fi34
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
