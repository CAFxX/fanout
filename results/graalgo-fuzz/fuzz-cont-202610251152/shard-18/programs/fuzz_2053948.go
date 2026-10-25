package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = (644 > 905)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 127}
    if v1 {
        var v5 float32 = ((-2.25 * 3.14159) + -2.25)
        _ = v5
    }
    v6 := GenS1{F0: 55}
    var v7 bool = (true || (v1 && v1))
    v8 := []int{127}
    switch ("" + ("hello" + ("foo bar" + "\n\t"))) {
    case "hello":
        if true {
            switch 127 {
            case -406:
                var v9 int64 = 533
                _ = v9
            case -670:
                v3[42] = ("a" + ("\n\t" + ""))
            case 19:
                switch "foo bar" {
                case "a":
                    var v10 int8 = -17
                    _ = v10
                case "foo bar":
                    fmt.Println(v1, v1)
                case "hello":
                    v11 := GenS1{F0: 218}
                    _ = v11
                }
            default:
                v1 = v7
            }
            if (!(v1 && (439 >= 2))) {
                v1 = (42 != 1)
                v7 = (3.14159 <= ((0.1 + 0.1) - float32(((float64(1e10) / 3.14159) * 1.5))))
                fmt.Println(v7, v7)
            }
            v12 := GenS1{F0: 183}
            _ = v12
        }
    case "foo bar":
        if true {
            v7 = (127 < 1)
            for fi13 := 0; fi13 < 19; fi13++ {
                v7 = (42 == 0)
                _ = fi13
            }
            var v14 uint8 = 104
            _ = v14
        }
    case "\n\t":
        v7 = (!true)
    default:
        var v15 uint32 = 2147483648
        _ = v15
    }
    var v16 bool = v1
    _ = len(v8)
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint16 = 1
            _ = fi18
            _ = v19
        }
        var v20 int = (fi17 >> ((1) & 63))
        _ = fi17
        _ = v20
    }
    if v1 {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint = 42
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 int16 = 42
                var v25 uint64 = 751
                v25 -= v25
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 bool = (((-2.25 + (float32((0.1 - -2.25)) / 1e10)) / 0.1) < 3.14159)
            _ = fi21
            _ = v22
            _ = v26
        }
    }
    if true {
        v7 = ((0 > 286) && true)
    } else {
        v16 = v16
    }
    var v27 int64 = 294
    v28 := GenS1{F0: 0}
    v8 = append(v8, 943)
    if v7 {
        v1 = false
        v2["a"] = 27
        for fi29 := 0; fi29 < 8; fi29++ {
            if (!(42 >= v27)) {
                _ = len(v2)
                var v30 uint8 = 255
                _ = v30
            }
            v1 = (-112 != 1)
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
