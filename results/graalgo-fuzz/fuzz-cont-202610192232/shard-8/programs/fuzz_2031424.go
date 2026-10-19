package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 string; F2 int }
func genFn0(p0 uint16, p1 uint16) uint8 {
    if false {
        if false {
            v7 := GenS1{F0: 42}
            for fi8 := 0; fi8 < 6; fi8++ {
                fmt.Println(p0, p0, p1)
                v9 := []int{}
                for fi10 := 0; fi10 < 19; fi10++ {
                    if (2147483647 >= 657) {
                        p0 += (p1 * (p0 * p0))
                    }
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
            p0 = (p0 >> ((2) & 63))
            _ = v7
        }
        p1 = (p0 ^ p1)
        switch (("hello" + (("foo bar" + (("foo bar" + "foo bar") + "a")) + ("\n\t" + "foo bar"))) + "") {
        case "foo bar":
            p1 = p1
        default:
            switch (("" + ("foo bar" + ("foo bar" + "hello"))) + "") {
            case "foo bar":
                var v11 uint16 = 0
                _ = v11
            case "hello":
                fmt.Println(p0, p1)
            case "":
                p0 += 65535
            }
        }
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        switch (fi12 & (fi12 & fi12)) {
        case 0:
            p0 += (^(p0))
        case 128:
            for fi13 := 0; fi13 < 8; fi13++ {
                p1 = 778
                switch (^(fi13)) {
                case 4294967295:
                    var v14 int16 = 164
                    _ = v14
                case 784:
                    fmt.Println(p0, p1, p0)
                default:
                    var v15 float32 = float32(0.1)
                    _ = v15
                }
                var v16 float64 = float64(float32(0.0))
                _ = fi13
                _ = v16
            }
        case 42:
            fmt.Println(p1, p1, fi12)
        default:
            v17 := []int{318, fi12}
            _ = v17
        }
        v18 := GenS1{F0: (fi12 | fi12)}
        _ = fi12
        _ = v18
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint8 = 127
    var v3 int16 = -472
    var v4 uint = 18446744073709551615
    v5 := make(map[uint32]bool)
    var v6 uint64 = 2
    v19 := GenS1{F0: 9223372036854775807}
    var v20 uint64 = v6
    v21 := GenS1{F0: -813}
    if (1 >= 2147483647) {
        if false {
            for fi22 := 0; fi22 < 15; fi22++ {
                for fi23 := 0; fi23 < 3; fi23++ {
                    v1 = append(v1, 849)
                    fmt.Println(v20)
                    _ = fi23
                }
                fmt.Println(v6, v3, v20)
                v20 = v6
                _ = fi22
            }
            v24 := GenS1{F0: 997}
            fmt.Println(v4)
            _ = v24
        } else {
            if false {
                var v25 int = 509
                v26 := GenS1{F0: (v25 | (-(v25)))}
                _ = v25
                _ = v26
            } else {
                var v27 uint = v4
                _ = v27
            }
            if true {
                var v28 uint = v4
                _ = v28
            } else {
                if (!(!true)) {
                    _ = len(v1)
                    v20 -= v6
                }
            }
        }
        v1 = append(v1, 980)
        var v29 int16 = v3
        _ = v29
    } else {
        var v30 float64 = ((((3.14159 + 1e10) * 0.1) * (-2.25 + (0.1 / (3.14159 + -2.25)))) / 1.5)
        fmt.Println(v6)
        var v31 int16 = (v3 - int16(v6))
        _ = v30
        _ = v31
    }
    switch "hello" {
    case "":
        v20 -= (v6 << ((v4) & 63))
    case "foo bar":
        v32 := v5[2147483648]
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 uint8 = (^(v2))
        for fi35 := 0; fi35 < 8; fi35++ {
            v20 = (v6 + (v6 - v6))
            v34 += v34
            _ = fi35
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v3)
    v20 = (v6 % (((v20 ^ 669)) | 1))
    if ((42 >= -1) && true) {
        var v36 int32 = 2147483647
        _ = v36
    } else {
        v20 = v20
        delete(v5, 2147483647)
    }
    v37 := GenS2{F0: uint8(v20), F1: "\n\t", F2: 2147483647}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
