package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint16) string {
    switch (((("\n\t" + "\n\t") + "") + "hello") + ("\n\t" + ("foo bar" + "\n\t"))) {
    case "hello":
        var v5 int = -924
        _ = v5
    case "a":
        switch "foo bar" {
        case "a":
            p1 += (p0 - 2)
        case "foo bar":
            p0 += p1
        case "":
            var v6 uint16 = (p1 / (((^(p0))) | 1))
            _ = v6
        }
    case "\n\t":
        var v7 int16 = -448
        _ = v7
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            if true {
                fmt.Println(fi8)
            } else {
                var v9 uint8 = 186
                _ = v9
            }
            _ = fi8
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        p1 -= (p1 & (p0 + (p1 / ((0) | 1))))
        var v11 uint64 = 0
        _ = fi10
        _ = v11
    }
    if (true || (485 <= 211)) {
        if (!true) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 string = ""
                var v14 string = "a"
                for fi15 := 0; fi15 < 8; fi15++ {
                    v16 := GenS1{F0: "", F1: 0}
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v17 uint32 = 4294967295
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS1{F0: (((("a" + (("foo bar" + "\n\t") + "a")) + ("a" + "a")) + (("\n\t" + ("a" + "hello")) + ("\n\t" + ""))) + ("hello" + "\n\t")), F1: -1}
        p1 = p1
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS2{F0: 229}
            p1 = (p1 * p0)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v22 := []int{}
    _ = p0
    _ = p1
    _ = v22
    return "foo bar"
}

func main() {
    var v1 bool = false
    var v2 int = 128
    v3 := make(map[string]int)
    var v4 uint8 = 255
    switch int(v4) {
    case 127:
        fmt.Println(v1, v4)
    case 32767:
        switch -9223372036854775808 {
        case 407:
            var v23 uint16 = 0
            _ = v23
        case 1:
            switch 32767 {
            case 9223372036854775807:
                delete(v3, "\n\t")
            case 31:
                if v1 {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v1 = true
                        v25 := GenS3{F0: -1}
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v27 := []int{(v2 ^ (fi26 | fi26))}
                            var v28 uint = 42
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = fi24
                        _ = v25
                    }
                }
            case 127:
                var v29 uint32 = 2
                _ = v29
            }
        case 0:
            v2 -= (v2 * (v2 * v2))
        }
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 bool = (!((v4 / (((v4 | v4)) | 1)) <= (^(v4))))
        _ = fi30
        _ = v31
    }
    var v32 bool = (32767 >= -756)
    v2 -= (v2 << ((731) & 63))
    var v33 float64 = 0.1
    switch "hello" {
    case "hello":
        if v32 {
            v34 := GenS1{F0: ("hello" + ("" + "foo bar")), F1: 0}
            for fi35 := 0; fi35 < 11; fi35++ {
                v32 = true
                var v36 int32 = -423
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            v32 = false
            if true {
                fmt.Println(v4, v2, v1)
                var v37 int = (v2 / ((-1) | 1))
                _ = v37
            }
        }
    case "foo bar":
        fmt.Println(v2, v4)
    }
    v38 := []int{(v2 * (v2 << ((876) & 63))), (v2 ^ 1), 42, 2147483647}
    v39 := GenS2{F0: 2147483647}
    if true {
        if (607 <= 65535) {
            v40 := GenS1{F0: "foo bar", F1: 457}
            _ = v40
        } else {
            v4 -= v4
            v1 = v1
            v41 := GenS2{F0: 0}
            _ = v41
        }
    } else {
        var v42 int = 32767
        v33 = 0.0
        var v43 uint32 = uint32(v4)
        _ = v42
        _ = v43
    }
    v2 = v2
    v38 = append(v38, (v2 ^ (v2 * 32768)))
    if (float32(float64(3.14159)) == float32((((v33 * v33) / v33) / (3.14159 * float64(1.5))))) {
        switch -2147483648 {
        case 63:
            var v44 string = "a"
            _ = v44
        }
    } else {
        if len(v38) > 0 {
            v38[0] = -145
        }
        var v45 uint32 = 946
        _ = v45
    }
    var v46 float32 = 3.14159
    var v47 bool = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v47
}
