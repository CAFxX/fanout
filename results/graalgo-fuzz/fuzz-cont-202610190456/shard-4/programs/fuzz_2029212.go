package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int, p1 uint32) int8 {
    switch "\n\t" {
    case "":
        v5 := []int{(p0 & (p0 % (((p0 >> ((901) & 63))) | 1))), p0, 32768, 42}
        _ = v5
    case "foo bar":
        p0 = (p0 & p0)
    case "hello":
        v6 := GenS1{F0: -1}
        _ = v6
    default:
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint = 734
            _ = fi7
            _ = v8
        }
    }
    if false {
        v9 := []int{}
        for fi10 := 0; fi10 < 5; fi10++ {
            if true {
                var v11 uint32 = (^(p1))
                _ = v11
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi10
        }
        _ = v9
    }
    var v12 int8 = 42
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{p0, p0}
        v15 := GenS1{F0: 554}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    return -128
}
func genFn1(p0 uint32) uint8 {
    for fi16 := 0; fi16 < 16; fi16++ {
        if (!(!false)) {
            fmt.Println(fi16)
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(fi16, fi17)
                _ = fi17
            }
        } else {
            switch "a" {
            case "foo bar":
                var v18 int16 = 0
                _ = v18
            default:
                v19 := []int{}
                _ = v19
            }
            var v20 uint16 = 255
            _ = v20
        }
        var v21 float64 = 1e10
        v22 := GenS2{F0: -692, F1: 105, F2: 2}
        _ = fi16
        _ = v21
        _ = v22
    }
    if true {
        if (!true) {
            p0 -= 2147483648
        }
        var v23 uint16 = 42
        _ = v23
    } else {
        v24 := GenS3{F0: 42, F1: 104}
        for fi25 := 0; fi25 < 19; fi25++ {
            p0 -= 4294967295
            switch fi25 {
            case -9223372036854775808:
                p0 -= p0
            case 32767:
                for fi26 := 0; fi26 < 14; fi26++ {
                    p0 -= (p0 ^ (p0 * p0))
                    var v27 uint8 = 127
                    v28 := GenS4{F0: 42}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = fi25
        }
        v29 := GenS5{F0: p0, F1: ("a" + "\n\t")}
        _ = v24
        _ = v29
    }
    p0 = p0
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            switch -2147483648 {
            case 127:
                switch 9223372036854775807 {
                case -2147483648:
                    switch (fi30 << ((4294967295) & 63)) {
                    case 127:
                        fmt.Println(fi30)
                    case 32768:
                        p0 = p0
                    }
                    switch (^(fi30)) {
                    case 128:
                        p0 += 259
                    case -2147483648:
                        var v32 uint64 = 4294967295
                        _ = v32
                    }
                }
            case -1:
                v33 := []int{fi30}
                _ = v33
            default:
                var v34 int8 = -108
                _ = v34
            }
            _ = fi31
        }
        switch fi30 {
        case 0:
            v35 := []int{}
            _ = v35
        case -1:
            fmt.Println(p0, fi30)
        case 127:
            v36 := GenS2{F0: -67, F1: -477, F2: 128}
            _ = v36
        }
        p0 = 42
        _ = fi30
    }
    v37 := GenS5{F0: (p0 | (p0 >> ((18446744073709551615) & 63))), F1: "\n\t"}
    _ = p0
    _ = v37
    return 127
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int = 4294967295
    var v3 int64 = 42
    v4 := []string{"\n\t", ("a" + "hello")}
    var v38 float64 = 3.14159
    switch ("hello" + "") {
    case "\n\t":
        v38 -= (-2.25 * v38)
    }
    v39 := GenS6{F0: 42}
    switch ("\n\t" + "a") {
    case "hello":
        v4 = append(v4, ("foo bar" + "hello"))
    case "a":
        fmt.Println(v3, v38)
    default:
        var v40 uint = 18446744073709551615
        _ = v40
    }
    switch (v2 / ((32767) | 1)) {
    case -1:
        var v41 int8 = -108
        _ = v41
    case 42:
        switch "a" {
        case "\n\t":
            var v42 int32 = -176
            _ = v42
        }
    case -454:
        v3 = (^(v3))
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        for fi44 := 0; fi44 < 17; fi44++ {
            v38 -= v38
            v45 := GenS1{F0: -32768}
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        fmt.Println(v3)
        v47 := GenS7{F0: 42}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi46
        _ = v47
    }
    if true {
        v3 = v3
        if len(v4) > 0 {
            v4[0] = (((("foo bar" + ("a" + "a")) + "a") + "\n\t") + "\n\t")
        }
        v2 -= v2
    }
    var v48 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v48
}
