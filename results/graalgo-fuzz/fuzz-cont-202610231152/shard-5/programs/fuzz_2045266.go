package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0() uint8 {
    var v7 uint8 = 62
    v7 = v7
    v8 := []int{}
    v7 = 0
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint16 = 65535
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return (v7 % (((v7 >> ((160) & 63))) | 1))
}
func genFn1(p0 int32, p1 int64, p2 int32) int32 {
    switch 128 {
    case 127:
        p0 = p2
    case -1:
        if (!((1e10 * (-2.25 + 3.14159)) != (((1e10 * 3.14159) * float32(1e10)) + (0.0 + float32(float64((0.0 / (1e10 / -2.25)))))))) {
            v11 := []int{380, 2147483647}
            var v12 string = ""
            _ = v11
            _ = v12
        }
    default:
        var v13 int8 = 127
        _ = v13
    }
    v14 := []int{2147483647, -2147483648, int(p2)}
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v15 int64 = (^(p1))
        _ = v15
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p0 ^ p0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -1
    var v3 int16 = 1
    var v4 float32 = -2.25
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]uint32)
    var v16 uint16 = 42
    v2 -= 1
    if false {
        switch -962 {
        case 1:
            v17 := GenS1{F0: (v2 | v2)}
            _ = v17
        case -9223372036854775808:
            delete(v1, 4294967295)
        case 32768:
            switch (("a" + ("" + "foo bar")) + "\n\t") {
            case "hello":
                v18 := []int{9223372036854775807, 412}
                _ = v18
            case "":
                v19 := v5[""]
                _ = v19
            default:
                v20 := GenS1{F0: (v2 + 120)}
                _ = v20
            }
        default:
            var v21 string = ("" + ((("a" + ((((("" + "hello") + "") + (("foo bar" + "") + "a")) + "hello") + "foo bar")) + ("\n\t" + "a")) + "hello"))
            _ = v21
        }
        v22 := []int{42, 139, 2147483647}
        _ = v22
    } else {
        v23 := []int{42, 42, int(v2)}
        var v24 int32 = 2147483647
        fmt.Println(v4, v24, v16)
        _ = v23
        _ = v24
    }
    var v25 string = ("foo bar" + "foo bar")
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS2{F0: (^(v3)), F1: 2147483647}
        switch v25 {
        case "":
            var v28 bool = false
            _ = v28
        default:
            v25 = ("a" + "foo bar")
        }
        v2 -= -1
        _ = fi26
        _ = v27
    }
    var v29 bool = false
    switch 80 {
    case -721:
        if (0.1 <= (0.0 / v4)) {
            v30 := v5[v25]
            _ = v30
        }
    case 128:
        for fi31 := 0; fi31 < 17; fi31++ {
            delete(v1, (fi31 - (^(fi31))))
            v32 := []int{-2147483648, fi31}
            _ = fi31
            _ = v32
        }
    default:
        switch ((v25 + (v25 + v25)) + (v25 + ((("" + v25) + "") + v25))) {
        case "foo bar":
            v33 := []int{0, -181}
            _ = v33
        case "a":
            if true {
                var v34 uint64 = 42
                _ = v34
            }
        case "hello":
            for fi35 := 0; fi35 < 18; fi35++ {
                v36 := []int{fi35, (fi35 * 2147483647), fi35, -569}
                v2 = (v2 / (((v2 * -1)) | 1))
                _ = fi35
                _ = v36
            }
        }
    }
    fmt.Println(v16, v16, v16)
    var v37 float32 = v4
    v38 := GenS2{F0: 32767, F1: -602}
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 string = "\n\t"
            _ = fi40
            _ = v41
        }
        switch (fi39 + (^(fi39))) {
        case 42:
            v42 := []int{fi39, fi39}
            _ = v42
        default:
            v43 := GenS1{F0: v2}
            _ = v43
        }
        if (v29 || v29) {
            var v44 string = ("hello" + "\n\t")
            _ = v44
        }
        _ = fi39
    }
    var v45 uint = 18446744073709551615
    switch "a" {
    case "a":
        v46 := GenS1{F0: (v2 * (v2 + -128))}
        _ = v46
    case "hello":
        fmt.Println(v4, v4)
    }
    v37 = 0.0
    if (v29 && ((!v29) || (!(v45 > 255)))) {
        if (-2147483648 < 128) {
            v47 := GenS2{F0: (v3 * -192), F1: 2147483647}
            for fi48 := 0; fi48 < 19; fi48++ {
                var v49 int32 = 2147483647
                var v50 uint64 = 763
                fmt.Println(v16, fi48, v16)
                _ = fi48
                _ = v49
                _ = v50
            }
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
