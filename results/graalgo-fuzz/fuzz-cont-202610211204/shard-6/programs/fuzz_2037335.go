package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint; F2 string }
func genFn0(p0 int16, p1 int32) uint {
    p1 = (p1 * (p1 | 1))
    if (42 < 1) {
        var v6 string = ("foo bar" + "\n\t")
        v7 := []int{}
        _ = v6
        _ = v7
    }
    var v8 uint16 = 0
    var v9 uint8 = 1
    var v10 int16 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int32, p1 int16) uint8 {
    var v11 uint16 = 32767
    var v12 int16 = p1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 144
}

func main() {
    v1 := []int{}
    var v2 int32 = 2147483647
    var v3 string = ("" + "\n\t")
    v4 := make(map[string]uint32)
    var v5 int8 = 127
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            if true {
                v2 -= v2
            } else {
                if ((1e10 * ((-2.25 - 0.0) / 1.0)) <= (0.1 * (1.5 + (1e10 - 1e10)))) {
                    if (!(((255 > 25) || false) || false)) {
                        v3 = v3
                        v15 := GenS1{F0: ("\n\t" + (v3 + v3))}
                        _ = v15
                    }
                }
                var v16 bool = false
                _ = v16
            }
            _ = fi14
        }
        var v17 uint64 = 732
        _ = fi13
        _ = v17
    }
    v3 = v3
    if (true && (1 <= 1)) {
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(fi18, fi18, fi18)
            switch (v3 + v3) {
            case "foo bar":
                var v19 float64 = 1e10
                _ = v19
            case "a":
                if (489 >= 65535) {
                    var v20 uint32 = 338
                    v21 := GenS2{F0: 0, F1: 255, F2: v3}
                    _ = v20
                    _ = v21
                } else {
                    v3 = "hello"
                }
            case "hello":
                var v22 int32 = (v2 >> ((42) & 63))
                _ = v22
            default:
                v2 -= v2
            }
            v5 = 42
            _ = fi18
        }
    }
    var v23 uint8 = genFn1(2147483647, 803)
    delete(v4, v3)
    var v24 bool = (false || true)
    _ = len(v4)
    v5 += (v5 * v5)
    if v24 {
        switch "a" {
        case "foo bar":
            for fi25 := 0; fi25 < 3; fi25++ {
                switch ((v3 + "hello") + ("a" + ("" + v3))) {
                case "":
                    v3 = v3
                case "foo bar":
                    if false {
                        var v26 uint16 = 124
                        v23 = (v23 % (((v23 >> ((255) & 63))) | 1))
                        v3 = v3
                        _ = v26
                    }
                default:
                    var v27 int8 = -1
                    _ = v27
                }
                if v24 {
                    v1 = append(v1, (fi25 % ((fi25) | 1)))
                    for fi28 := 0; fi28 < 1; fi28++ {
                        var v29 int = fi28
                        v30 := GenS1{F0: "a"}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    v23 += genFn1(v2, int16(v5))
                }
                v31 := GenS2{F0: 81, F1: 2, F2: v3}
                _ = fi25
                _ = v31
            }
        case "hello":
            var v32 int8 = (v5 % ((v5) | 1))
            _ = v32
        default:
            var v33 float32 = 3.14159
            _ = v33
        }
        if (2 <= genFn0(42, (v2 % (((v2 * 124)) | 1)))) {
            if (!v24) {
                _ = len(v4)
                _ = len(v4)
            }
        }
    }
    var v34 uint16 = 0
    switch 128 {
    case 2147483647:
        for fi35 := 0; fi35 < 9; fi35++ {
            if (!(uint(v23) >= 255)) {
                var v36 int64 = 0
                _ = v36
            }
            v37 := GenS1{F0: ("foo bar" + "")}
            var v38 string = v3
            _ = fi35
            _ = v37
            _ = v38
        }
    case 1:
        v4[(v3 + v3)] = 2147483647
    case 32768:
        var v39 string = ((v3 + v3) + "\n\t")
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
