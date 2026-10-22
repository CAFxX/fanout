package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 uint16) bool {
    for fi4 := 0; fi4 < 6; fi4++ {
        p0 -= 49
        if false {
            p0 = p0
        }
        _ = fi4
    }
    if true {
        for fi5 := 0; fi5 < 19; fi5++ {
            p0 -= 0
            var v6 int32 = 0
            fmt.Println(p0)
            _ = fi5
            _ = v6
        }
    }
    _ = p0
    return (!true)
}
func genFn1(p0 int8, p1 int16) int {
    var v7 int16 = (p1 | p1)
    if true {
        if false {
            var v8 uint16 = 0
            _ = v8
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(p0, p1)
                switch ("\n\t" + ("foo bar" + "")) {
                case "foo bar":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v7 = p1
                        fmt.Println(p1, p0)
                        switch (fi10 << ((2) & 63)) {
                        case 128:
                            var v11 int16 = (v7 / ((v7) | 1))
                            _ = v11
                        case -9223372036854775808:
                            v12 := []int{fi9, (^(fi9))}
                            _ = v12
                        case 279:
                            var v13 bool = true
                            _ = v13
                        default:
                            var v14 float32 = 3.14159
                            _ = v14
                        }
                        p1 += (v7 - (p1 - (p1 + (v7 >> ((963) & 63)))))
                        _ = fi10
                    }
                case "":
                    for fi15 := 0; fi15 < 20; fi15++ {
                        p0 += (p0 + 1)
                        _ = fi15
                    }
                case "\n\t":
                    v16 := []int{fi9, fi9, fi9, (^(fi9))}
                    _ = v16
                }
                if ((!(((((0.1 - float64(float32(0.0))) + (1.5 + 0.0)) / 1.0) / (1e10 + (-2.25 - (-2.25 * 1.5)))) != 3.14159)) || (0 >= (^(p1)))) {
                    v17 := GenS1{F0: 882}
                    if ((4294967295 < 941) || ((("foo bar" + "foo bar") + "") < "")) {
                        var v18 bool = (!true)
                        _ = v18
                    } else {
                        fmt.Println(v7)
                        var v19 int16 = -246
                        _ = v19
                    }
                    _ = v17
                }
                _ = fi9
            }
            v20 := GenS2{F0: 25}
            _ = v20
        }
        p1 = -1
    } else {
        if (100 == 100) {
            v21 := GenS3{F0: -1, F1: "foo bar"}
            var v22 uint8 = 128
            v7 = (v7 << ((255) & 63))
            _ = v21
            _ = v22
        }
    }
    v23 := []int{}
    var v24 int16 = (p1 | -418)
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 46
    v2 := []string{"hello", "hello", "hello", "hello"}
    v3 := []string{"foo bar", "a", ("\n\t" + ("a" + ""))}
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(v1, v1)
        v1 = 0
        fmt.Println(fi25)
        _ = fi25
    }
    switch -808 {
    case 634:
        v1 += (v1 ^ v1)
    case -9223372036854775808:
        v1 -= (^(v1))
    default:
        var v26 uint8 = v1
        _ = v26
    }
    fmt.Println(v1)
    if false {
        fmt.Println(v1)
    }
    if false {
        var v27 uint = 65535
        v1 += (v1 + (v1 + v1))
        for fi28 := 0; fi28 < 16; fi28++ {
            v3 = append(v3, "hello")
            _ = fi28
        }
        _ = v27
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        switch ("a" + "\n\t") {
        case "":
            var v30 uint32 = 685
            _ = v30
        case "hello":
            v1 = (v1 * v1)
        case "a":
            if (0 >= 2147483647) {
                var v31 uint64 = 0
                v32 := GenS4{F0: 4294967295}
                _ = v31
                _ = v32
            } else {
                v1 += (v1 | v1)
                var v33 int64 = 534
                _ = v33
            }
        }
        v34 := GenS5{F0: 825, F1: -1, F2: 0}
        v1 = v1
        _ = fi29
        _ = v34
    }
    v1 = v1
    v1 -= 255
    var v35 float64 = 0.0
    switch 127 {
    case 32768:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    case 9223372036854775807:
        for fi36 := 0; fi36 < 18; fi36++ {
            v1 = v1
            for fi37 := 0; fi37 < 1; fi37++ {
                fmt.Println(v35)
                _ = len(v2)
                _ = fi37
            }
            switch fi36 {
            case 2147483647:
                var v38 uint32 = 2147483648
                _ = v38
            default:
                var v39 int64 = 1
                _ = v39
            }
            _ = fi36
        }
    default:
        var v40 int = -1
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
