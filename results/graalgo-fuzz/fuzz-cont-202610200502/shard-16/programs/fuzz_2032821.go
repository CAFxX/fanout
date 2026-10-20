package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 bool }
func genFn0(p0 bool) uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 = true
            _ = fi5
        }
        var v6 float32 = (float32(3.14159) / 0.1)
        var v7 uint32 = 0
        _ = fi4
        _ = v6
        _ = v7
    }
    var v8 string = ""
    for fi9 := 0; fi9 < 19; fi9++ {
        p0 = false
        var v10 float64 = (3.14159 + 0.0)
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v10, p0, v8)
        case "a":
            switch fi9 {
            case 6:
                if (("foo bar" + v8) < (v8 + v8)) {
                    fmt.Println(fi9, v10)
                    v8 = ((v8 + v8) + v8)
                }
            case -9223372036854775808:
                for fi11 := 0; fi11 < 13; fi11++ {
                    fmt.Println(fi9, v8)
                    _ = fi11
                }
            }
        case "hello":
            var v12 int16 = -309
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 float64 = (1.5 / float64(float32(1.5)))
    _ = p0
    _ = v13
    _ = v8
    return 65535
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 bool = true
    var v14 float32 = 0.1
    if v3 {
        v15 := GenS1{F0: ("a" + "foo bar")}
        if len(v1) > 0 {
            v1[0] = 74
        }
        _ = v15
    } else {
        switch 1 {
        case -9223372036854775808:
            if true {
                v3 = v3
                if len(v1) > 0 {
                    v1[0] = 128
                }
                switch 32767 {
                case 0:
                    if v3 {
                        v3 = (!(18446744073709551615 <= 0))
                        if ((!true) || ((("\n\t" + ("\n\t" + "hello")) + "a") <= "hello")) {
                            var v16 float64 = 0.0
                            fmt.Println(v16, v16)
                            _ = v16
                        }
                    } else {
                        if v3 {
                            v17 := GenS1{F0: ("a" + "hello")}
                            var v18 bool = (true || (float32(3.14159) >= v14))
                            _ = v17
                            _ = v18
                        }
                    }
                default:
                    v14 -= float32(3.14159)
                }
            } else {
                var v19 bool = true
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int64 = -1
                    for fi22 := 0; fi22 < 14; fi22++ {
                        fmt.Println(fi20, v19)
                        _ = fi22
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        case 127:
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 bool = v3
                _ = fi23
                _ = v24
            }
        default:
            v2[-790] = false
        }
        v25 := GenS1{F0: "hello"}
        for fi26 := 0; fi26 < 5; fi26++ {
            if v3 {
                fmt.Println(fi26)
                var v27 float32 = (v14 - v14)
                _ = v27
            }
            var v28 float64 = (3.14159 - 3.14159)
            _ = fi26
            _ = v28
        }
        _ = v25
    }
    _ = len(v2)
    _ = len(v1)
    switch 0 {
    case -182:
        var v29 int16 = -451
        _ = v29
    }
    v14 += v14
    var v30 uint64 = 1
    switch ("" + ("hello" + "a")) {
    case "foo bar":
        switch "a" {
        case "\n\t":
            var v31 uint64 = (^(v30))
            _ = v31
        case "hello":
            fmt.Println(v14)
        }
    case "hello":
        var v32 string = ((("foo bar" + "hello") + "hello") + (("a" + ("\n\t" + "foo bar")) + "a"))
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        switch "a" {
        case "a":
            v34 := GenS2{F0: 43, F1: 687, F2: 970}
            _ = v34
        case "\n\t":
            v35 := GenS3{F0: ("\n\t" < "a")}
            _ = v35
        default:
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint = 2
                _ = fi36
                _ = v37
            }
        }
        var v38 uint = 255
        var v39 int = 32767
        _ = fi33
        _ = v38
        _ = v39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v30 -= (v30 | 4294967295)
        switch "" {
        case "hello":
            for fi41 := 0; fi41 < 15; fi41++ {
                var v42 bool = true
                if true {
                    v14 = v14
                    _ = len(v2)
                    _ = len(v2)
                } else {
                    v43 := GenS3{F0: (127 == 1)}
                    fmt.Println(v42, v30)
                    v2[fi41] = (!false)
                    _ = v43
                }
                var v44 bool = v3
                _ = fi41
                _ = v42
                _ = v44
            }
        }
        _ = fi40
    }
    var v45 float64 = ((0.1 - -2.25) / -2.25)
    fmt.Println(v3, v3, v3)
    v46 := GenS1{F0: ((("\n\t" + "a") + "\n\t") + "\n\t")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v45
    _ = v46
}
