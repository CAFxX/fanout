package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 bool, p1 bool) bool {
    var v5 uint = 837
    p0 = ((v5 >> (((^(v5))) & 63)) < 4294967295)
    v6 := []int{-9223372036854775808, 231, -1, 128}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (0 >= 255)
}
func genFn1() int16 {
    var v7 int8 = 127
    var v8 int8 = 0
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 = (v8 * v7)
        _ = fi9
    }
    var v10 string = ("a" + "foo bar")
    _ = v10
    _ = v7
    _ = v8
    return 787
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]int)
    v3 := []uint32{(v1 & v1)}
    v4 := make(map[string]int)
    v11 := GenS1{F0: 9223372036854775807}
    v12 := GenS1{F0: 0}
    fmt.Println(v1)
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            v2["\n\t"] = fi13
            switch "\n\t" {
            case "hello":
                v1 = 734
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                var v14 uint8 = 128
                _ = v14
            }
            if len(v3) > 0 {
                v3[0] = v1
            }
        }
        if len(v3) > 0 {
            v3[0] = v1
        }
        var v15 bool = false
        _ = fi13
        _ = v15
    }
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 17; fi17++ {
        v1 = v1
        v18 := GenS1{F0: -9223372036854775808}
        var v19 int = fi17
        _ = fi17
        _ = v18
        _ = v19
    }
    switch "" {
    case "\n\t":
        v1 += (^(v1))
    case "":
        var v20 string = ("\n\t" + "hello")
        _ = v20
    default:
        v21 := GenS1{F0: 9223372036854775807}
        _ = v21
    }
    switch 498 {
    case 9223372036854775807:
        fmt.Println(v1, v16)
    case 776:
        v1 = (^(v1))
    case 2147483647:
        switch 9223372036854775807 {
        case 0:
            v3 = append(v3, 2)
        case 372:
            var v22 string = "hello"
            _ = v22
        default:
            v16 = 0.0
        }
    default:
        v1 = (v1 / ((655) | 1))
    }
    switch (("foo bar" + ("\n\t" + (("\n\t" + "hello") + "\n\t"))) + "") {
    case "foo bar":
        v3 = append(v3, 549)
    case "a":
        v23 := GenS2{F0: (^(v1)), F1: 1}
        _ = v23
    default:
        v16 = v16
    }
    var v24 int8 = 127
    switch 42 {
    case 127:
        for fi25 := 0; fi25 < 4; fi25++ {
            if (!(264 != 42)) {
                var v26 int = (fi25 - (fi25 + fi25))
                _ = v26
            } else {
                if true {
                    v27 := GenS1{F0: 9223372036854775807}
                    fmt.Println(v16, v24)
                    for fi28 := 0; fi28 < 13; fi28++ {
                        v1 -= v1
                        _ = fi28
                    }
                    _ = v27
                } else {
                    fmt.Println(v16, v24)
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := v2[("hello" + "foo bar")]
                    fmt.Println(v24, v1)
                    var v31 uint64 = 833
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                fmt.Println(v16)
            }
            v32 := GenS3{F0: ((42 <= 127) && (true || ((!(42 >= -305)) && false))), F1: (v24 | (v24 ^ v24))}
            v1 = (v1 - v1)
            _ = fi25
            _ = v32
        }
    case 0:
        _ = len(v3)
    case -717:
        switch -1 {
        case 4294967295:
            for fi33 := 0; fi33 < 5; fi33++ {
                v3 = append(v3, 2)
                v24 = (^(v24))
                _ = fi33
            }
        }
    default:
        if (!false) {
            var v34 float32 = 1e10
            var v35 float32 = float32((float64((v34 / float32(3.14159))) + 1e10))
            var v36 int = 42
            _ = v34
            _ = v35
            _ = v36
        } else {
            var v37 float32 = (1e10 * 0.0)
            _ = v37
        }
    }
    var v38 float64 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
}
