package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint32) string {
    switch -914 {
    case 172:
        var v5 float32 = 0.0
        _ = v5
    default:
        var v6 int8 = -95
        _ = v6
    }
    v7 := []int{224}
    _ = p0
    _ = v7
    return ("\n\t" + ((("a" + "") + "\n\t") + ("hello" + ("hello" + "a"))))
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8, (fi8 % ((fi8) | 1))}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: (fi10 >> ((698) & 63)), F1: 42, F2: 2}
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int8 = 1
            v14 := GenS2{F0: 866, F1: v13, F2: 9223372036854775807}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = fi10
        _ = fi10
        _ = v11
        _ = v15
    }
    return 255
}

func main() {
    v1 := []int{-115, -734, -2147483648}
    var v2 int64 = 9223372036854775807
    var v3 bool = (!("" < (("hello" + ((("\n\t" + ("a" + (("a" + "hello") + ("a" + "")))) + "foo bar") + ("hello" + "foo bar"))) + "a")))
    var v4 int8 = -22
    switch "\n\t" {
    case "hello":
        for fi16 := 0; fi16 < 11; fi16++ {
            if ((1 < 18446744073709551615) && (false && v3)) {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: (fi16 << ((42) & 63)), F1: (v4 + (v4 - v4)), F2: 642}
                    var v19 uint64 = 42
                    v2 += v2
                    v3 = ("a" > "\n\t")
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 int32 = 1
                    var v22 float32 = 1.5
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v2 += v2
                var v23 uint = 416
                _ = v23
            }
            v24 := GenS3{F0: 1, F1: 42}
            _ = fi16
            _ = v24
        }
    case "":
        var v25 uint8 = 230
        _ = v25
    case "foo bar":
        v4 = (v4 & (^(v4)))
    default:
        v26 := GenS3{F0: -1, F1: 65535}
        _ = v26
    }
    v3 = (42 <= 1)
    switch ("hello" + "") {
    case "\n\t":
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS4{F0: 1, F1: v2, F2: int8(v2)}
            var v29 int8 = (v4 & int8(v2))
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v4 = v4
        _ = fi30
    }
    fmt.Println(v4)
    v31 := GenS5{F0: 863}
    _ = len(v1)
    _ = len(v1)
    if true {
        v1 = append(v1, 9223372036854775807)
        var v32 bool = v3
        _ = len(v1)
        _ = v32
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v4 = v4
        if (!((v3 && v3) || (65535 >= 4294967295))) {
            fmt.Println(v4)
            switch "\n\t" {
            case "":
                _ = len(v1)
            case "hello":
                fmt.Println(v4)
            }
        } else {
            v4 = 127
            switch -593 {
            case 32767:
                var v33 int16 = 1000
                _ = v33
            default:
                fmt.Println(v4, v4)
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
