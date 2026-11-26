package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 int16, p2 uint16) int64 {
    v6 := GenS1{F0: 0}
    if true {
        if true {
            var v7 uint = 255
            var v8 string = (("foo bar" + ("hello" + ("hello" + "hello"))) + "foo bar")
            _ = v7
            _ = v8
        } else {
            p1 += (p0 / ((p1) | 1))
        }
        fmt.Println(p0)
    }
    v9 := GenS2{F0: 0, F1: 42}
    switch "hello" {
    case "foo bar":
        v10 := []int{32767, 42}
        _ = v10
    case "":
        p1 += (p1 / (((p1 * p0)) | 1))
    default:
        p0 += -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 int64) int64 {
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int64 = p0
        for fi13 := 0; fi13 < 8; fi13++ {
            v12 = -709
            var v14 int16 = 32767
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 uint32 = 556
    switch 127 {
    case -485:
        v16 := GenS1{F0: 127}
        _ = v16
    }
    switch (("a" + "\n\t") + ("" + "foo bar")) {
    case "foo bar":
        v17 := []int{42, 509, -9223372036854775808}
        _ = v17
    case "a":
        for fi18 := 0; fi18 < 12; fi18++ {
            v15 += v15
            _ = fi18
        }
    }
    _ = p0
    _ = v15
    return p0
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 string = ("\n\t" + ("\n\t" + "\n\t"))
    var v3 uint = 65535
    var v4 string = v2
    var v5 int32 = 682
    fmt.Println(v5, v2, v2)
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int8 = 127
        v21 := GenS3{F0: uint16(v3)}
        v1 = append(v1, v4)
        _ = fi19
        _ = v20
        _ = v21
    }
    v4 = v2
    var v22 uint = (v3 >> (((^(v3))) & 63))
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            switch fi23 {
            case 32767:
                var v24 uint8 = 42
                _ = v24
            case 1:
                var v25 int16 = 697
                _ = v25
            case 500:
                v4 = ("" + "a")
            default:
                var v26 int8 = 127
                _ = v26
            }
            _ = fi23
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v5 -= v5
        v5 = (v5 / ((1) | 1))
        if ((4294967295 <= 4294967295) && (((-1 >= genFn0(1, 42, 111)) || (9223372036854775807 == genFn1(300))) || false)) {
            if ((v3 >= (v3 & 913)) && true) {
                v4 = "foo bar"
                v4 = (v2 + v2)
            } else {
                fmt.Println(v22, v22, v2)
                if len(v1) > 0 {
                    v1[0] = v2
                }
            }
            v28 := GenS4{F0: (v22 ^ (v22 - 1))}
            _ = v28
        } else {
            if ((true || true) && false) {
                var v29 uint64 = 42
                _ = v29
            }
            fmt.Println(v2, fi27)
        }
        _ = fi27
    }
    var v30 uint = (v3 << ((731) & 63))
    v31 := GenS3{F0: 65535}
    fmt.Println(v5, v3, v5)
    if true {
        var v32 float64 = 1.5
        _ = v32
    } else {
        if (260 < 1) {
            switch v4 {
            case "hello":
                v3 += 65535
            default:
                var v33 float64 = 0.1
                _ = v33
            }
            var v34 uint = (v22 << (((v30 >> (((v30 % ((v22) | 1))) & 63))) & 63))
            _ = v34
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
