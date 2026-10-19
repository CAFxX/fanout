package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64, p1 int8) bool {
    v7 := GenS1{F0: 712}
    if (((255 != 255) && (!(127 <= 140))) || true) {
        v8 := []int{127, 32767, 9223372036854775807}
        var v9 int32 = 606
        _ = v8
        _ = v9
    } else {
        switch "foo bar" {
        case "":
            v10 := []int{}
            _ = v10
        case "foo bar":
            p0 -= (^(p0))
        case "hello":
            v11 := []int{1, 32767}
            _ = v11
        }
        v12 := GenS1{F0: 4294967295}
        if (false && false) {
            fmt.Println(p0, p0)
        } else {
            switch 925 {
            case 2147483647:
                p0 = p0
            case 9223372036854775807:
                fmt.Println(p0, p1)
            case 0:
                fmt.Println(p1, p1)
            default:
                var v13 int16 = 32767
                _ = v13
            }
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
        _ = v12
    }
    switch ((("a" + "hello") + "a") + (("\n\t" + "\n\t") + "")) {
    case "\n\t":
        fmt.Println(p0)
    case "":
        if true {
            var v15 uint16 = 2
            for fi16 := 0; fi16 < 4; fi16++ {
                p1 += (p1 - (p1 | (p1 % ((31) | 1))))
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS1{F0: 474}
                    v19 := []int{-9223372036854775808}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            v20 := GenS1{F0: 42}
            _ = v15
            _ = v20
        }
    case "a":
        var v21 uint64 = 18446744073709551615
        _ = v21
    default:
        var v22 int16 = 1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 uint8, p1 uint32) int32 {
    p1 = p1
    v23 := GenS1{F0: 18446744073709551615}
    v24 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return -2147483648
}

func main() {
    var v1 uint = 255
    var v2 uint16 = 255
    var v3 int32 = 42
    var v4 uint8 = 0
    var v5 uint64 = 18446744073709551615
    v6 := []int{128}
    var v25 uint32 = 294
    v26 := GenS2{F0: genFn1(54, (v25 >> ((v1) & 63)))}
    var v27 uint8 = 42
    _ = len(v6)
    if (true || (0 != (^(v5)))) {
        var v28 int16 = -1
        if ("foo bar" > "foo bar") {
            if true {
                v2 = 118
                v27 = (v4 >> (((v1 ^ 42)) & 63))
            } else {
                v1 += uint(v25)
                var v29 float32 = 1e10
                v30 := GenS1{F0: (v5 ^ 2)}
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 20; fi31++ {
                v3 = (v3 % ((v3) | 1))
                v2 = v2
                var v32 int16 = (^(v28))
                _ = fi31
                _ = v32
            }
        }
        fmt.Println(v5, v27)
        _ = v28
    } else {
        if true {
            var v33 uint16 = 255
            _ = v33
        }
        v34 := GenS1{F0: (v5 >> (((v1 - v1)) & 63))}
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v3 = (v3 % ((-2147483648) | 1))
        if (!(fi35 == fi35)) {
            if true {
                v4 = (^(v27))
                v36 := GenS1{F0: v5}
                var v37 int32 = v3
                _ = v36
                _ = v37
            }
            v38 := GenS2{F0: (v3 * -1)}
            v27 += v4
            _ = v38
        }
        v39 := GenS3{F0: 184, F1: (v2 | 65535)}
        _ = fi35
        _ = v39
    }
    switch "" {
    case "foo bar":
        for fi40 := 0; fi40 < 19; fi40++ {
            if len(v6) > 0 {
                v6[0] = fi40
            }
            _ = fi40
        }
    case "hello":
        if (v5 == (^(v5))) {
            if false {
                if true {
                    fmt.Println(v4)
                    var v41 uint32 = 42
                    var v42 int = 42
                    _ = v41
                    _ = v42
                }
            }
            if ((v5 | (v5 >> ((v1) & 63))) <= v5) {
                v3 -= (v3 - (v3 ^ v3))
                var v43 int32 = v3
                _ = v43
            }
            v44 := GenS3{F0: genFn1((v27 | v27), (v25 + 2)), F1: 1}
            _ = v44
        }
    case "\n\t":
        v45 := GenS4{F0: (v25 & 1)}
        _ = v45
    }
    var v46 uint8 = uint8(v2)
    var v47 uint = v1
    v48 := GenS5{F0: -155}
    var v49 uint32 = (^(v25))
    for fi50 := 0; fi50 < 14; fi50++ {
        var v51 string = ("a" + (("a" + "\n\t") + "hello"))
        var v52 float64 = 0.0
        var v53 uint16 = v2
        _ = fi50
        _ = v51
        _ = v52
        _ = v53
    }
    v2 = (v2 + (v2 | 2))
    v27 = (v46 >> (((v1 & (v47 << ((v47) & 63)))) & 63))
    if false {
        v54 := GenS2{F0: v3}
        v2 -= v2
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v54
    }
    v27 += 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
