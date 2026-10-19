package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 uint, p2 uint16) int32 {
    var v6 uint = (p1 + 255)
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(p0, v6, p1)
        _ = fi8
    }
    switch 4294967295 {
    case 865:
        var v9 int8 = 1
        _ = v9
    case -1:
        switch "a" {
        case "foo bar":
            var v10 uint32 = 65535
            _ = v10
        case "hello":
            var v11 int16 = 679
            _ = v11
        case "a":
            var v12 uint = p1
            _ = v12
        }
    default:
        v13 := GenS1{F0: p2}
        _ = v13
    }
    var v14 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{724, 2, 980}
    v3 := []string{}
    v4 := []uint32{2147483647}
    v5 := make(map[string]uint32)
    v15 := GenS2{F0: false, F1: 9223372036854775807, F2: true}
    var v16 float32 = 1e10
    if (1 >= 4294967295) {
        v16 -= ((v16 * v16) + v16)
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            if true {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v16 = v16
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        _ = len(v5)
                    case "a":
                        var v19 string = ""
                        _ = v19
                    }
                    v20 := GenS1{F0: 876}
                    var v21 int64 = 92
                    var v22 uint64 = 2
                    _ = fi18
                    _ = v20
                    _ = v21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: 42}
                var v24 int = (fi17 | fi17)
                _ = v23
                _ = v24
            }
            if false {
                v25 := GenS3{F0: 42}
                _ = v25
            } else {
                var v26 int32 = 495
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if false {
                        if ((!true) || true) {
                            var v27 int64 = 752
                            v28 := v5[""]
                            _ = v27
                            _ = v28
                        }
                        v26 += (v26 + v26)
                    } else {
                        if (!false) {
                            var v29 uint8 = 205
                            _ = v29
                        }
                    }
                } else {
                    v30 := GenS1{F0: 42}
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v1 += v1
                        _ = fi31
                    }
                    _ = v30
                }
                var v32 string = ("a" + (("foo bar" + "") + "hello"))
                _ = v26
                _ = v32
            }
            _ = fi17
        }
    }
    if (-128 >= -22) {
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 bool = (!true)
            _ = fi33
            _ = v34
        }
        var v35 uint16 = 42
        _ = v35
    } else {
        if true {
            v16 -= 0.0
            v36 := GenS2{F0: ((0 < 70) || (genFn0(0, 556, 854) != 2147483647)), F1: 42, F2: false}
            _ = v36
        }
        v16 = float32(((3.14159 / 1e10) * 1e10))
    }
    var v37 uint64 = 0
    v37 += (v37 & v37)
    switch 633 {
    case 0:
        if true {
            _ = len(v5)
            v1 = 0
            v16 -= v16
        } else {
            v37 += v37
            var v38 bool = false
            v39 := GenS1{F0: 938}
            _ = v38
            _ = v39
        }
    case -9223372036854775808:
        v5["\n\t"] = 2
    case -1:
        fmt.Println(v37, v37)
    }
    delete(v5, "\n\t")
    if false {
        v1 -= v1
        _ = len(v4)
    }
    v1 = -701
    fmt.Println(v37)
    if (false && (true || false)) {
        fmt.Println(v1, v37, v37)
    } else {
        var v40 int64 = 42
        fmt.Println(v16, v1)
        _ = v40
    }
    var v41 int = 697
    v2 = append(v2, 4294967295)
    fmt.Println(v37)
    v37 = v37
    var v42 uint64 = v37
    for fi43 := 0; fi43 < 5; fi43++ {
        for fi44 := 0; fi44 < 1; fi44++ {
            v16 -= (-2.25 * (v16 - ((float32(1.5) * (v16 - (3.14159 - (v16 + -2.25)))) - (3.14159 * v16))))
            v16 += -2.25
            _ = fi44
        }
        switch "a" {
        case "\n\t":
            v45 := v5["\n\t"]
            _ = v45
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
