package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 int, p1 uint8) uint8 {
    var v4 float64 = 3.14159
    if true {
        v5 := []int{p0, (p0 * p0), (-(p0)), p0}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    }
    v6 := GenS1{F0: true}
    fmt.Println(p0, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 1
}
func genFn1(p0 int16) uint32 {
    switch "foo bar" {
    case "foo bar":
        if false {
            var v7 uint16 = 967
            v8 := GenS1{F0: (0 >= 255)}
            _ = v7
            _ = v8
        } else {
            fmt.Println(p0)
            fmt.Println(p0)
            var v9 int8 = -128
            _ = v9
        }
    }
    if true {
        p0 = (p0 - (p0 - 0))
    }
    var v10 int32 = -2147483648
    var v11 float32 = (1e10 / (0.1 + (1e10 - 1.5)))
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int64 = 307
    v2 := []int{}
    var v3 uint64 = 42
    var v12 int8 = 60
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int64 = v1
        switch (("hello" + (("a" + "a") + "hello")) + ("" + ("a" + "hello"))) {
        case "a":
            fmt.Println(v14)
        default:
            var v15 uint32 = genFn1(1000)
            _ = v15
        }
        fmt.Println(v3)
        _ = fi13
        _ = v14
    }
    if false {
        var v16 float64 = 1.5
        var v17 int32 = -851
        if true {
            if false {
                var v18 int8 = v12
                for fi19 := 0; fi19 < 19; fi19++ {
                    switch fi19 {
                    case 0:
                        fmt.Println(v3, v18, v17)
                    default:
                        var v20 uint = 18446744073709551615
                        _ = v20
                    }
                    _ = len(v2)
                    _ = len(v2)
                    _ = fi19
                }
                _ = len(v2)
                _ = v18
            } else {
                if (((("foo bar" + "") + ("\n\t" + (("a" + "\n\t") + ""))) + ("a" + "")) <= "foo bar") {
                    var v21 int16 = 344
                    _ = v21
                }
            }
            if false {
                _ = len(v2)
            }
            if len(v2) > 0 {
                v2[0] = 678
            }
        }
        _ = v16
        _ = v17
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            if (2 != 2) {
                v3 -= v3
            } else {
                var v23 uint16 = 1
                if (1 > 2147483647) {
                    fmt.Println(v3, v1, v23)
                }
                fmt.Println(fi22, v3, v12)
                _ = v23
            }
            if true {
                v12 -= (v12 << ((4294967295) & 63))
            }
            v2 = append(v2, fi22)
            _ = fi22
        }
        fmt.Println(v3)
        var v24 int64 = -9223372036854775808
        _ = v24
    }
    v25 := GenS1{F0: (1e10 >= (float32(((1.5 + 0.0) + (float64((float32(1.5) / (0.0 + 1.5))) - 0.0))) - ((0.1 / 0.1) * 1.5)))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS2{F0: 1000, F1: 41}
    fmt.Println(v3, v3)
    fmt.Println(v1, v12)
    switch 920 {
    case 42:
        var v27 int32 = -2147483648
        _ = v27
    case -528:
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS3{F0: -75, F1: fi29}
                var v31 int32 = 1
                var v32 int16 = 0
                fmt.Println(v1, v12)
                v3 = 528
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi28
        }
    case -9223372036854775808:
        v1 -= (v1 << ((802) & 63))
    }
    fmt.Println(v3)
    if ((v3 * (v3 % ((v3) | 1))) != v3) {
        var v33 int8 = 1
        v3 = (v3 - (^(v3)))
        v3 = v3
        _ = v33
    }
    v34 := GenS2{F0: 32767, F1: 2}
    v12 += (v12 >> ((255) & 63))
    var v35 uint = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((v35 * (v35 - v35)) < (v35 + (v35 / (((v35 | 65535)) | 1)))) {
        switch 548 {
        case 358:
            v36 := GenS3{F0: -32768, F1: -2147483648}
            _ = v36
        case 1:
            var v37 uint16 = 42
            _ = v37
        case 32768:
            v38 := GenS2{F0: 32767, F1: 42}
            _ = v38
        default:
            for fi39 := 0; fi39 < 1; fi39++ {
                _ = len(v2)
                _ = fi39
            }
        }
        _ = len(v2)
    } else {
        v40 := GenS1{F0: false}
        var v41 string = "foo bar"
        var v42 string = (v41 + v41)
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println(v3, v1)
    var v43 uint8 = 113
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v43
}
