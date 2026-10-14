package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int64 }
func genFn0(p0 string) int64 {
    var v4 uint = 42
    var v5 uint = v4
    _ = p0
    _ = v4
    _ = v5
    return 20
}
func genFn1(p0 string, p1 uint16) uint64 {
    v6 := GenS1{F0: -20}
    switch -771 {
    case -9223372036854775808:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 bool = (921 >= 2)
    v3 := []uint32{65535}
    v7 := GenS2{F0: 1}
    v8 := GenS2{F0: -125}
    if false {
        fmt.Println(v1)
        if (937 >= 4294967295) {
            v2 = false
            var v9 string = "foo bar"
            fmt.Println(v2, v9)
            _ = v9
        }
        var v10 uint8 = v1
        _ = v10
    }
    switch ("foo bar" + "\n\t") {
    case "":
        if (!v2) {
            if v2 {
                if (-9223372036854775808 > 4294967295) {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v2 = (!(("a" + "a") >= "foo bar"))
                        _ = fi11
                    }
                    v1 += (v1 % ((42) | 1))
                    fmt.Println(v2, v1)
                }
                switch 128 {
                case 2147483647:
                    var v12 uint64 = genFn1("a", 255)
                    _ = v12
                }
            } else {
                v13 := GenS3{F0: -794}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v1)
                _ = v13
            }
            var v14 uint8 = (v1 / (((v1 + (^(v1)))) | 1))
            _ = v14
        } else {
            var v15 uint = 1
            for fi16 := 0; fi16 < 1; fi16++ {
                if v2 {
                    v17 := GenS3{F0: 1}
                    fmt.Println(v2, v1, v1)
                    _ = v17
                }
                if v2 {
                    if true {
                        var v18 uint32 = 1
                        _ = v18
                    }
                } else {
                    if v2 {
                        v15 = 65535
                        var v19 uint = v15
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 int = fi16
                    _ = fi20
                    _ = v21
                }
                _ = fi16
            }
            _ = v15
        }
    }
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 1; fi22++ {
        v2 = (true || v2)
        fmt.Println(v2, v2, v1)
        _ = fi22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 uint32 = 2
    switch 32767 {
    case 32767:
        v24 := GenS4{F0: 4294967295, F1: uint16(v23), F2: (v23 + v23)}
        _ = v24
    case 0:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v25 := GenS5{F0: 2147483647, F1: 0}
        _ = v25
    }
    var v26 string = ""
    if len(v3) > 0 {
        v3[0] = v23
    }
    if v2 {
        for fi27 := 0; fi27 < 20; fi27++ {
            switch (fi27 - 4294967295) {
            case -9223372036854775808:
                v2 = (!v2)
            case 4294967295:
                v2 = (133 < -465)
            case 0:
                for fi28 := 0; fi28 < 7; fi28++ {
                    v29 := GenS6{F0: 1, F1: v1, F2: 203}
                    var v30 string = "foo bar"
                    fmt.Println(fi28, v30, v26)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi27
        }
    }
    var v31 int16 = -854
    v23 = v23
    if false {
        switch "hello" {
        case "hello":
            var v32 uint8 = (v1 & (v1 / (((v1 << ((42) & 63))) | 1)))
            _ = v32
        case "":
            v33 := GenS4{F0: 42, F1: 42, F2: v23}
            _ = v33
        case "\n\t":
            v31 -= v31
        }
        var v34 string = (v26 + v26)
        v35 := GenS7{F0: 42, F1: 0}
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v31
    _ = v7
    _ = v8
}
