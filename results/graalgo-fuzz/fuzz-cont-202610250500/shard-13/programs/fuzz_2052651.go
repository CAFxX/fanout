package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32, p1 bool) uint8 {
    switch "a" {
    case "hello":
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                var v9 int32 = int32(fi8)
                _ = v9
            }
            v10 := GenS1{F0: 42}
            p0 += 2147483647
            _ = fi8
            _ = v10
        }
    case "\n\t":
        p0 = p0
    default:
        var v11 uint16 = 65535
        _ = v11
    }
    var v12 float64 = ((3.14159 + 3.14159) - 1.5)
    p1 = false
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := []int{fi13, -1, (fi13 & (fi13 - 32767)), (fi13 / (((fi13 ^ fi13)) | 1))}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 1
}
func genFn1() int {
    var v15 bool = (0.1 <= 0.0)
    if true {
        fmt.Println(v15)
        var v16 int8 = 0
        _ = v16
    }
    _ = v15
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int = 1
    var v4 uint8 = 151
    var v5 float64 = (-2.25 / -2.25)
    v6 := make(map[string]string)
    var v17 uint16 = 60
    v18 := []int{(v3 & v3), v3, (v3 >> ((uint(v4)) & 63))}
    var v19 int32 = -12
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 bool = false
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 0
    v19 -= v19
    var v23 bool = true
    v5 = (1e10 / 1.0)
    v22 += (v22 % ((v22) | 1))
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int16 = 1
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(v17)
            _ = fi26
        }
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 int16 = -1
        var v29 string = (("a" + "\n\t") + (((("a" + ("\n\t" + (("\n\t" + "a") + (("\n\t" + "") + ("a" + ((("a" + "\n\t") + "foo bar") + ((("hello" + "") + "\n\t") + "a"))))))) + ("hello" + (("foo bar" + "hello") + "foo bar"))) + ("" + ("" + ("\n\t" + (("a" + "") + (("hello" + ("" + "foo bar")) + ("hello" + "\n\t"))))))) + "a"))
        _ = fi27
        _ = v28
        _ = v29
    }
    switch "a" {
    case "\n\t":
        var v30 float32 = 0.0
        _ = v30
    case "a":
        if len(v18) > 0 {
            v18[0] = v3
        }
    default:
        v31 := GenS1{F0: 1}
        _ = v31
    }
    v6["hello"] = ("\n\t" + ("\n\t" + ("" + "")))
    for fi32 := 0; fi32 < 3; fi32++ {
        if len(v18) > 0 {
            v18[0] = fi32
        }
        _ = fi32
    }
    if true {
        v18 = append(v18, v3)
        if v23 {
            for fi33 := 0; fi33 < 16; fi33++ {
                switch "a" {
                case "foo bar":
                    for fi34 := 0; fi34 < 8; fi34++ {
                        v35 := GenS2{F0: uint32(v22), F1: 2}
                        if len(v18) > 0 {
                            v18[0] = genFn1()
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                _ = fi33
            }
        }
        v4 = uint8(v22)
    } else {
        v18 = append(v18, v3)
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        switch v3 {
        case 1:
            var v37 int8 = 28
            _ = v37
        case -1:
            switch "" {
            case "foo bar":
                for fi38 := 0; fi38 < 9; fi38++ {
                    switch (fi36 | genFn1()) {
                    case 794:
                        var v39 bool = v23
                        _ = v39
                    case 128:
                        var v40 uint16 = (v17 + (v17 + 934))
                        _ = v40
                    default:
                        var v41 int64 = 809
                        _ = v41
                    }
                    v42 := v2[""]
                    _ = fi38
                    _ = v42
                }
            case "\n\t":
                v4 += (v4 / ((v4) | 1))
            }
        default:
            switch v3 {
            case -34:
                v43 := GenS1{F0: 0}
                _ = v43
            case -41:
                var v44 uint = uint(v17)
                _ = v44
            case 745:
                var v45 float64 = v5
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 19; fi46++ {
            switch int(v17) {
            case -9223372036854775808:
                fmt.Println(v5, v3, fi46)
            }
            var v47 uint32 = 2
            _ = fi46
            _ = v47
        }
        _ = fi36
    }
    if (false && (96 > 0)) {
        v4 += v4
        for fi48 := 0; fi48 < 4; fi48++ {
            v19 = -51
            v18 = append(v18, (-(fi48)))
            _ = fi48
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v4, v17, v22)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
