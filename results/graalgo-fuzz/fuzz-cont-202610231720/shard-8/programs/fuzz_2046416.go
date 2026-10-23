package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16) uint32 {
    p0 = (p0 * (p0 + p0))
    var v5 uint16 = p0
    var v6 float64 = ((float64((0.0 + 1e10)) + 3.14159) / (1e10 + -2.25))
    v7 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 uint16) uint8 {
    v8 := []int{}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 164
        v11 := GenS1{F0: fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = p0
    p0 = p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    return (p0 & (p0 / ((uint8(p1)) | 1)))
}

func main() {
    v1 := []string{}
    v2 := []string{"\n\t", ((("a" + ("" + "a")) + "") + ("foo bar" + "")), ("" + "")}
    v3 := []string{""}
    var v4 float64 = (3.14159 - (-2.25 * (1.5 * (0.0 - (1e10 / float64(3.14159))))))
    var v12 uint16 = 32767
    if len(v3) > 0 {
        v3[0] = (("" + "hello") + (("\n\t" + "foo bar") + (((("foo bar" + ("\n\t" + "\n\t")) + ("" + "\n\t")) + "") + "foo bar")))
    }
    v12 += v12
    v2 = append(v2, "\n\t")
    for fi13 := 0; fi13 < 5; fi13++ {
        switch ("" + ("\n\t" + (("a" + ((((("a" + "foo bar") + "\n\t") + "foo bar") + "foo bar") + ("\n\t" + ""))) + "hello"))) {
        case "a":
            v14 := GenS2{F0: 1, F1: v12}
            _ = v14
        case "":
            v4 = (v4 + v4)
        default:
            v15 := GenS2{F0: genFn0((v12 >> ((1) & 63))), F1: 2}
            _ = v15
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float64 = (1.5 - (1.5 + (3.14159 + -2.25)))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                v18 := GenS1{F0: 127}
                _ = v18
            } else {
                v19 := GenS3{F0: 1, F1: genFn0((v12 + (v12 & 255))), F2: 1}
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v4, v12, v4)
        _ = fi13
    }
    v4 += float64(1e10)
    var v20 uint8 = 127
    v4 -= v4
    v21 := GenS4{F0: false}
    v2 = append(v2, "\n\t")
    var v22 int = -2147483648
    if ((false || false) && (!(float64(1e10) > (v4 * v4)))) {
        var v23 uint32 = 4294967295
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = v23
    } else {
        v4 -= ((v4 + float64(float32((((0.0 * v4) / ((v4 * ((float64(0.1) + v4) + 3.14159)) + (0.1 - ((1.5 - (0.0 * (float64(((float32(v4) * 1e10) / 1.0)) - v4))) * (1e10 + (v4 * (float64(1.5) - v4))))))) - ((v4 * v4) + float64((-2.25 * -2.25))))))) - v4)
        v22 -= (^(v22))
        v4 += v4
    }
    v4 = v4
    switch "a" {
    case "":
        var v24 string = "hello"
        _ = v24
    default:
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int8 = 127
            v27 := GenS5{F0: ((-2.25 * ((3.14159 * float32(v4)) * -2.25)) >= 0.1)}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    if ((v20 + (^(v20))) != 42) {
        fmt.Println(v20, v4)
        if false {
            if true {
                v22 = 127
                v28 := GenS5{F0: ((false || false) || ((!false) || (!(((!("foo bar" > "hello")) && false) && (!true)))))}
                _ = v28
            }
        } else {
            v29 := GenS5{F0: true}
            _ = v29
        }
    } else {
        var v30 int64 = 42
        v31 := GenS2{F0: 932, F1: (v12 ^ v12)}
        _ = v30
        _ = v31
    }
    v32 := GenS5{F0: (-128 != -128)}
    var v33 int64 = 1
    fmt.Println(v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
