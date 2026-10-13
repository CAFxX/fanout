package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8) bool {
    var v5 int64 = 972
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 17; fi6++ {
        if false {
            if false {
                v5 = 0
            }
        }
        _ = fi6
    }
    var v7 int8 = p0
    v8 := GenS1{F0: ("hello" + "\n\t"), F1: 42, F2: -413}
    _ = p0
    _ = v5
    _ = v7
    _ = v8
    return (true && true)
}
func genFn1(p0 bool) int {
    switch 32768 {
    case 2147483647:
        p0 = (!p0)
    case 127:
        fmt.Println(p0)
    case 934:
        var v9 int8 = -28
        _ = v9
    default:
        switch (("\n\t" + ("foo bar" + "\n\t")) + "") {
        case "hello":
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int = (fi10 % ((fi10) | 1))
                _ = fi10
                _ = v11
            }
        default:
            var v12 int = 42
            _ = v12
        }
    }
    if p0 {
        for fi13 := 0; fi13 < 6; fi13++ {
            p0 = p0
            var v14 float64 = 1.5
            _ = fi13
            _ = v14
        }
        var v15 uint16 = 724
        _ = v15
    }
    switch -1 {
    case 128:
        var v16 uint64 = 4294967295
        _ = v16
    case -2147483648:
        v17 := []int{32767}
        _ = v17
    default:
        var v18 float32 = 1e10
        _ = v18
    }
    _ = p0
    return -615
}

func main() {
    var v1 uint = 42
    v2 := []int{}
    v3 := []string{"a", "\n\t", ("foo bar" + "a"), "hello"}
    var v4 string = (("" + "a") + "\n\t")
    fmt.Println(v1, v1)
    _ = len(v2)
    v4 = v4
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v4 = "a"
            _ = fi20
        }
        if true {
            v1 -= 0
            _ = len(v3)
        } else {
            var v21 int64 = 0
            switch (^(fi19)) {
            case 1:
                fmt.Println(v21, v4)
            case 821:
                v21 = (v21 % (((v21 >> ((v1) & 63))) | 1))
            case -1:
                if true {
                    v22 := GenS1{F0: v4, F1: 457, F2: 42}
                    v2 = append(v2, (^(fi19)))
                    _ = v22
                } else {
                    v1 = 4294967295
                }
            }
            _ = v21
        }
        _ = fi19
    }
    var v23 uint64 = 0
    v23 = (v23 >> (((v1 / ((619) | 1))) & 63))
    for fi24 := 0; fi24 < 3; fi24++ {
        v1 += (v1 / (((^(v1))) | 1))
        _ = fi24
    }
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v25 := GenS1{F0: (("\n\t" + (v4 + (((v4 + v4) + "foo bar") + v4))) + (v4 + v4)), F1: v23, F2: 42}
    v3 = append(v3, v4)
    v4 = v4
    v4 = (v4 + v4)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println(v23)
    switch 1 {
    case 361:
        v23 += 18446744073709551615
    case 785:
        v1 += (v1 / (((^(v1))) | 1))
    case -706:
        if (float64((float32(3.14159) / 1e10)) != 1.5) {
            v23 = v23
            switch 32767 {
            case 9223372036854775807:
                var v26 int = 4294967295
                _ = v26
            case 127:
                switch -2147483648 {
                case 4294967295:
                    v1 -= (v1 << ((1) & 63))
                }
            case -729:
                v23 -= 1
            default:
                var v27 uint = 18446744073709551615
                _ = v27
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 int = genFn1(true)
                _ = fi28
                _ = v29
            }
        }
    }
    if (-917 > 480) {
        v3 = append(v3, "\n\t")
        if (2 < 2147483647) {
            v30 := GenS2{F0: 42, F1: true, F2: 42}
            v4 = (v4 + ((v4 + v4) + v4))
            _ = v30
        } else {
            var v31 float32 = 0.1
            v32 := GenS3{F0: v23}
            _ = v31
            _ = v32
        }
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
