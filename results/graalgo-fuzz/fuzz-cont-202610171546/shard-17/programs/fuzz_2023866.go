package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 bool = ((p0 / ((p0) | 1)) >= p0)
            v9 = v9
            var v10 uint = 67
            if ((!(int8(fi7) == 0)) || (1000 <= 42)) {
                v11 := GenS1{F0: 65535, F1: 493, F2: v9}
                fmt.Println(fi8)
                _ = v11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        switch ((((((("a" + (("a" + "a") + "\n\t")) + (((("a" + "foo bar") + "") + "foo bar") + "hello")) + "") + ("a" + (("\n\t" + ("\n\t" + (("" + "") + ("foo bar" + "")))) + "\n\t"))) + "hello") + "a") + "\n\t") {
        case "":
            p0 = (p0 + (p0 << ((4294967295) & 63)))
        case "\n\t":
            var v12 int16 = 1000
            _ = v12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 string = ("" + ("hello" + "hello"))
        var v15 int32 = 0
        for fi16 := 0; fi16 < 15; fi16++ {
            p0 -= p0
            var v17 uint32 = 65535
            var v18 uint = 0
            var v19 int32 = v15
            fmt.Println(fi13, fi16, v17)
            v17 -= v17
            v18 = (^(v18))
            v15 = 42
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    p0 = (p0 - (p0 + (p0 + (p0 | (p0 << ((631) & 63))))))
    v20 := []int{146, 1, -727, -944}
    _ = p0
    _ = v20
    return 1
}

func main() {
    var v1 bool = (42 == 1)
    v2 := make(map[uint32]int)
    v3 := []string{("\n\t" + "\n\t")}
    var v4 string = "foo bar"
    v5 := make(map[uint32]uint32)
    var v6 uint = 110
    var v21 uint8 = 128
    var v22 uint = 0
    delete(v2, 65535)
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS1{F0: 439, F1: 0, F2: true}
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 float32 = (-2.25 / 1.5)
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    v22 = v22
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS2{F0: 203}
        v29 := GenS3{F0: v4, F1: 32767, F2: 65535}
        v22 = v6
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v4, v1, v21)
    v22 -= uint(v21)
    v4 = "a"
    var v30 uint16 = 255
    if len(v3) > 0 {
        v3[0] = ((v4 + ((v4 + "hello") + (v4 + v4))) + v4)
    }
    var v31 uint8 = v21
    _ = len(v2)
    v21 = v31
    fmt.Println(v22, v6, v22)
    if ((-1 >= 1) && (!(v1 || false))) {
        var v32 int = 217
        if v1 {
            var v33 uint32 = 778
            v34 := GenS4{F0: 9223372036854775807, F1: 32767}
            _ = len(v3)
            _ = v33
            _ = v34
        } else {
            v21 -= (v31 - (v21 % ((v21) | 1)))
            var v35 uint = (v22 % (((v22 * v6)) | 1))
            var v36 uint8 = (v21 ^ 98)
            _ = v35
            _ = v36
        }
        _ = v32
    } else {
        var v37 int = 32767
        for fi38 := 0; fi38 < 12; fi38++ {
            v22 -= 0
            _ = fi38
        }
        _ = v37
    }
    v21 -= 154
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
