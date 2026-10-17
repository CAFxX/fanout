package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint16, p1 string, p2 uint32) string {
    v5 := []int{-145, -2147483648, 128, 127}
    var v6 uint8 = 1
    var v7 int16 = 1000
    var v8 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (((("\n\t" + p1) + "foo bar") + "\n\t") + (p1 + p1))
}
func genFn1(p0 int, p1 int8, p2 int) uint32 {
    v9 := GenS1{F0: 1, F1: (^(p1))}
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 bool = true
            v12 := []int{33, 4294967295, (p2 >> ((18446744073709551615) & 63))}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: (p2 << ((836) & 63)), F1: 0, F2: (p2 >> ((541) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return 2147483648
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{522}
    var v3 float64 = 1e10
    var v4 int = 32767
    var v14 uint8 = 127
    var v15 uint8 = (^(v14))
    switch (((((((("\n\t" + "foo bar") + "") + ("hello" + "hello")) + "a") + ("" + "a")) + ("hello" + (("\n\t" + "") + ((("foo bar" + ("a" + "\n\t")) + "\n\t") + "a")))) + "") + "\n\t") {
    case "hello":
        if (v3 >= v3) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    if ("" == ("\n\t" + "\n\t")) {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            if true {
                switch (fi16 - (^(v4))) {
                case 408:
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
            }
            switch "hello" {
            case "":
                _ = len(v2)
            }
            v17 := GenS1{F0: 215, F1: 0}
            _ = fi16
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 int32 = -1
            var v20 uint64 = v1
            var v21 bool = true
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        var v22 uint16 = 65535
        _ = v22
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 string = ""
        var v25 string = v24
        v14 -= v15
        _ = fi23
        _ = v24
        _ = v25
    }
    _ = len(v2)
    fmt.Println(v15)
    if ((v14 - (v15 | 0)) <= 127) {
        switch "a" {
        case "hello":
            switch v4 {
            case -883:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 4294967295:
                v2 = append(v2, 65535)
            default:
                var v26 string = ("\n\t" + ((("" + ("\n\t" + "hello")) + "hello") + "hello"))
                _ = v26
            }
        }
    } else {
        switch int(v15) {
        case 726:
            if len(v2) > 0 {
                v2[0] = 65535
            }
        }
        _ = len(v2)
        v27 := GenS3{F0: 42, F1: 42}
        _ = v27
    }
    fmt.Println(v4, v1, v3)
    v28 := GenS4{F0: 796}
    if false {
        switch (-(v4)) {
        case 896:
            v4 += (v4 - v4)
        case 4294967295:
            v15 = v15
        case 295:
            v4 -= (-(v4))
        }
        v15 = (^(v15))
    } else {
        v29 := GenS5{F0: (^(v4)), F1: 0, F2: -1}
        var v30 uint = 65535
        _ = v29
        _ = v30
    }
    if false {
        var v31 int = v4
        if (255 <= 2) {
            v32 := GenS6{F0: genFn1(v4, 100, (-(v31))), F1: int16(v31)}
            if len(v2) > 0 {
                v2[0] = 264
            }
            v14 -= v14
            _ = v32
        }
        var v33 int = v4
        _ = v31
        _ = v33
    } else {
        v4 -= 16
        v14 += 1
    }
    var v34 uint64 = 42
    v1 -= (v34 ^ (v34 << ((42) & 63)))
    fmt.Println(v4, v14)
    var v35 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
