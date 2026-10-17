package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 string = ("a" + ("a" + ("\n\t" + ("hello" + ("foo bar" + "foo bar")))))
    v2 := []int{127}
    var v3 uint64 = 1
    var v4 int = 128
    for fi5 := 0; fi5 < 6; fi5++ {
        if (498 <= 0) {
            _ = len(v2)
            if (!(4294967295 < 606)) {
                fmt.Println(v4)
                var v6 uint64 = v3
                v1 = (((v1 + v1) + ("foo bar" + (v1 + v1))) + v1)
                _ = v6
            }
        }
        v4 -= fi5
        _ = fi5
    }
    var v7 float64 = 0.0
    if true {
        var v8 int = v4
        switch (v8 ^ v4) {
        case -782:
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v4, v7)
                fmt.Println(v1, v7)
                if false {
                    v2 = append(v2, (fi9 << ((42) & 63)))
                    var v10 string = "foo bar"
                    _ = v10
                } else {
                    fmt.Println(v8)
                    var v11 uint64 = (v3 | (^(v3)))
                    _ = v11
                }
                _ = fi9
            }
        case -9223372036854775808:
            v12 := GenS1{F0: uint(v3), F1: 316, F2: 1}
            _ = v12
        default:
            v13 := GenS1{F0: 18446744073709551615, F1: 1, F2: -1}
            _ = v13
        }
        _ = v8
    } else {
        var v14 uint16 = 42
        for fi15 := 0; fi15 < 9; fi15++ {
            switch (v4 / ((fi15) | 1)) {
            case -906:
                v16 := GenS2{F0: fi15, F1: true, F2: 65535}
                _ = v16
            case 4294967295:
                _ = len(v2)
            }
            _ = fi15
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 string = ""
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS3{F0: v4, F1: 0}
            var v21 float64 = ((v7 + v7) - ((1.5 + v7) / 0.1))
            if false {
                switch -97 {
                case 4294967295:
                    if ((!(v21 >= float64(-2.25))) || false) {
                        v7 = (0.0 / 1.0)
                        v2 = append(v2, v4)
                        v1 = ((("hello" + "\n\t") + (v18 + v18)) + "a")
                    }
                default:
                    var v22 bool = (v18 < (("\n\t" + v18) + v18))
                    _ = v22
                }
                v1 = v1
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        if len(v2) > 0 {
            v2[0] = fi17
        }
        _ = fi17
        _ = v18
    }
    if ((false || (false && true)) && (100 > 66)) {
        fmt.Println(v7)
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int16 = 925
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            v26 := GenS4{F0: 42, F1: 303, F2: 2147483647}
            var v27 uint8 = 127
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        v4 = v4
        var v28 string = v1
        _ = v28
    }
    v7 = (v7 * 0.0)
    if false {
        var v29 int = (v4 & int(v3))
        _ = v29
    } else {
        var v30 uint = 96
        _ = v30
    }
    var v31 int64 = 657
    v2 = append(v2, v4)
    var v32 bool = (!true)
    if len(v2) > 0 {
        v2[0] = v4
    }
    var v33 int64 = 1
    var v34 uint16 = 766
    v35 := GenS4{F0: -842, F1: 65535, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v7
}
